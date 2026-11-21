package vvlib

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/carneades/carneades-4/src/engine/caes"
	"gopkg.in/yaml.v2"
)

// CLI checks treat the `carneades` command as a black box and start a new
// process for every run, which is how the library is used in practice: no
// state survives from one file to the next, and the validation gate in front
// of `carneades eval` is part of what is tested.

type CLI struct{ Bin string }

// FindCLI returns the carneades binary: $CARNEADES_BIN, or one built from the
// engine sources in third_party/ at the repository root.
func FindCLI() (*CLI, error) {
	if b := os.Getenv("CARNEADES_BIN"); b != "" {
		if _, err := os.Stat(b); err != nil {
			return nil, err
		}
		return &CLI{Bin: b}, nil
	}
	tmp, err := os.MkdirTemp("", "vv-cli-*")
	if err != nil {
		return nil, err
	}
	bin := filepath.Join(tmp, "carneades")
	cmd := exec.Command("go", "build", "-o", bin, "github.com/carneades/carneades-4/src/cmd/carneades")
	if out, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("building the carneades command: %v\n%s", err, out)
	}
	return &CLI{Bin: bin}, nil
}

type CLIRun struct {
	Stdout, Stderr string
	Exit           int
	TimedOut       bool
}

func (c *CLI) Run(stdin string, args ...string) CLIRun {
	cmd := exec.Command(c.Bin, args...)
	cmd.Stdin = strings.NewReader(stdin)
	var o, e bytes.Buffer
	cmd.Stdout, cmd.Stderr = &o, &e
	r := CLIRun{}
	if err := cmd.Start(); err != nil {
		r.Exit = -1
		r.Stderr = err.Error()
		return r
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if ee, ok := err.(*exec.ExitError); ok {
			r.Exit = ee.ExitCode()
		} else if err != nil {
			r.Exit = -1
		}
	case <-time.After(InferTimeout):
		cmd.Process.Kill()
		r.TimedOut = true
		r.Exit = -1
	}
	r.Stdout, r.Stderr = o.String(), e.String()
	return r
}

// labels reads the statement labels from the YAML that `carneades eval -t yaml` prints.
func labels(out string) (map[string]string, error) {
	var doc struct {
		Statements map[string]interface{} `yaml:"statements"`
	}
	if err := yaml.Unmarshal([]byte(out), &doc); err != nil {
		return nil, err
	}
	l := map[string]string{}
	for id, v := range doc.Statements {
		l[id] = "undecided"
		if m, ok := v.(map[interface{}]interface{}); ok {
			if lab, ok := m["label"].(string); ok {
				l[id] = lab
			}
		}
	}
	return l, nil
}

// CLIFileChecks: does the command accept the file as it is?
func CLIFileChecks(cli *CLI, lib *Library) []Result {
	out := []Result{}
	r := cli.Run(lib.Text, "check")
	problems := strings.TrimSpace(r.Stderr)
	n := 0
	if problems != "" {
		n = len(strings.Split(problems, "\n"))
	}
	out = append(out, res(lib.Rel, "(file)", "cli/check-clean", func() []string {
		if n == 0 {
			return nil
		}
		return []string{fmt.Sprintf("`carneades check` reports %d problem line(s); first: %s", n, firstLine(problems))}
	}()))
	r = cli.Run(lib.Text, "eval", "-t", "yaml")
	if strings.TrimSpace(r.Stdout) == "" {
		out = append(out, fail(lib.Rel, "(file)", "cli/eval-produces-output",
			"`carneades eval` prints nothing and exits with status %d: any validation problem makes it skip inference, evaluation and export", r.Exit))
	} else {
		out = append(out, pass(lib.Rel, "(file)", "cli/eval-produces-output"))
	}
	return out
}

// Scenario is the scheme file followed by the ground premises of one scheme as assumptions.
func Scenario(lib *Library, in *Instance) string {
	var b strings.Builder
	b.WriteString(lib.Text)
	b.WriteString("\nassumptions:\n")
	for _, p := range in.Premises {
		b.WriteString("  - '" + p + "'\n")
	}
	return b.String()
}

// CLISchemeCheck: with the file as the theory and the scheme's premises as
// assumptions, does `carneades eval` construct the argument and accept its conclusion?
func CLISchemeCheck(cli *CLI, lib *Library, s *caes.Scheme) Result {
	in := NewInstance(s)
	r := cli.Run(Scenario(lib, in), "eval", "-t", "yaml")
	const check = "cli/scenario"
	switch {
	case r.TimedOut:
		return fail(lib.Rel, s.Id, check, "`carneades eval` did not finish")
	case strings.TrimSpace(r.Stdout) == "":
		return fail(lib.Rel, s.Id, check, "`carneades eval` printed nothing (exit status %d); stderr starts: %s", r.Exit, firstLine(r.Stderr))
	}
	l, err := labels(r.Stdout)
	if err != nil {
		return fail(lib.Rel, s.Id, check, "the output is not valid YAML (%v); it starts: %s", firstLine(err.Error()), firstLine(r.Stdout))
	}
	probs := []string{}
	for _, c := range in.Conclusions {
		lab, ok := l[Norm(c)]
		if !ok {
			lab = "missing"
		}
		if lab != "in" {
			probs = append(probs, fmt.Sprintf("%s is %s", c, lab))
		}
	}
	return res(lib.Rel, s.Id, check, probs)
}

// CLIEngineChecks: behaviour of the command itself.
func CLIEngineChecks(cli *CLI) []Result {
	out := []Result{}
	broken := baseTheory + "assumptions:\n  - p(X)\n" // not ground: a validation problem

	r := cli.Run(broken, "eval", "-t", "yaml")
	out = append(out, eres("cli-eval-signals-refusal", r.Exit != 0,
		"`carneades eval` refuses a file with a validation problem by printing nothing and exiting with status 0 (stderr: %s)", firstLine(r.Stderr)))

	r = cli.Run(broken, "check")
	out = append(out, eres("cli-check-exit-status", r.Exit != 0,
		"`carneades check` exits with status 0 although it reports a problem (%s), so a script or CI job cannot use it", firstLine(r.Stderr)))

	r = cli.Run(MutualRecursion, "eval", "-t", "yaml")
	_, err := labels(r.Stdout)
	out = append(out, eres("cli-output-stays-valid", err == nil && strings.TrimSpace(r.Stdout) != "",
		"when the rule engine hits its limit its message lands in the exported graph; the YAML output starts: %s", firstLine(r.Stdout)))

	first := cli.Run(baseTheory+`
argument_schemes:
  - id: s2
    variables: [X,Y]
    premises: ["p(X,Y)"]
    conclusions: ["q(X,Y)"]
assumptions: ["p(a,b)"]
`, "eval", "-t", "yaml")
	second := cli.Run(first.Stdout, "eval", "-t", "yaml")
	out = append(out, eres("cli-reads-its-own-yaml", first.Exit == 0 && second.Exit == 0 && strings.TrimSpace(second.Stdout) != "",
		"`carneades eval -t yaml` output fed back into `carneades eval` fails (exit status %d): %s", second.Exit, firstLine(second.Stderr)))

	if dir, err := EngineDir(); err == nil {
		b, _ := os.ReadFile(filepath.Join(dir, "examples", "AGs", "YAML", "walton.yml"))
		r = cli.Run(string(b), "eval", "-t", "yaml")
		l, err := labels(r.Stdout)
		ok := err == nil && l["untrustworthy(joe)"] == "in" && l["¬caused_by(global_warming,humans)"] == "out"
		out = append(out, eres("baseline-cli-upstream-walton-example", ok,
			"upstream's own Walton example no longer evaluates as before: untrustworthy(joe) %s, ¬caused_by(global_warming,humans) %s (%v)",
			l["untrustworthy(joe)"], l["¬caused_by(global_warming,humans)"], err))
	}
	return out
}
