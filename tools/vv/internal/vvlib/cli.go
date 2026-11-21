package vvlib

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
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
	return c.RunInterrupted(stdin, 0, args...)
}

// RunInterrupted is Run, with an interrupt sent to the command after the
// given time (none when it is 0).
func (c *CLI) RunInterrupted(stdin string, after time.Duration, args ...string) CLIRun {
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
	if after > 0 {
		go func() {
			time.Sleep(after)
			cmd.Process.Signal(os.Interrupt)
		}()
	}
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
	// docs/contracts.md, contract 3: invalid input exits with status 1 and prints no evaluation
	out = append(out, eres("cli-eval-signals-refusal", r.Exit == 1 && strings.TrimSpace(r.Stdout) == "",
		"`carneades eval` on a file with a validation problem should print nothing and exit with status 1 (invalid input); it exited with %d (stderr: %s)", r.Exit, firstLine(r.Stderr)))

	r = cli.Run(broken, "check")
	out = append(out, eres("cli-check-exit-status", r.Exit == 1,
		"`carneades check` should exit with status 1 (invalid input) when it reports a problem; it exited with %d (%s)", r.Exit, firstLine(r.Stderr)))

	// docs/contracts.md, contracts 3 and 4: a graph that stipulates a statement
	// and its negation is an inconsistent context, reported with status 3 and
	// no evaluation; one that stipulates only the statement evaluates.
	bad := cli.Run("statements:\n  p: P is the case.\nassumptions: [p, \"¬p\"]\n", "eval", "-t", "yaml")
	good := cli.Run("statements:\n  p: P is the case.\nassumptions: [p]\n", "eval", "-t", "yaml")
	out = append(out, eres("cli-reports-inconsistent-context",
		bad.Exit == 3 && strings.TrimSpace(bad.Stdout) == "" && good.Exit == 0 && strings.TrimSpace(good.Stdout) != "",
		"assuming p and ¬p should end with status 3 (inconsistent context) and no evaluation, and assuming p alone with status 0: got %d (stdout %d bytes) and %d", bad.Exit, len(bad.Stdout), good.Exit))

	// docs/contracts.md, contract 3: an interrupt cancels the evaluation, which
	// ends with status 5 and prints nothing. The graph has 300 facts and a
	// scheme over pairs of them: 90,000 arguments, seconds of work.
	{
		var b strings.Builder
		b.WriteString("language:\n  p/1: \"%s is p.\"\n  q/2: \"%s q %s.\"\nargument_schemes:\n  - id: pair\n    variables: [X, Y]\n    premises:\n      - p(X)\n      - p(Y)\n    conclusions:\n      - q(X,Y)\nassumptions:\n")
		for i := 0; i < 300; i++ {
			fmt.Fprintf(&b, "  - p(c%d)\n", i)
		}
		r := cli.RunInterrupted(b.String(), 300*time.Millisecond, "eval", "-t", "yaml")
		out = append(out, eres("cli-interrupt-is-cancelled", r.Exit == 5 && strings.TrimSpace(r.Stdout) == "" && strings.Contains(r.Stderr, "cancelled"),
			"`carneades eval` interrupted 0.3 s into a run of several seconds should print nothing and exit with status 5 (cancelled); it exited with %d, %d bytes on stdout (stderr: %s)", r.Exit, len(r.Stdout), firstLine(r.Stderr)))
	}

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

		// every example upstream ships passes `check`, but the two that are
		// refused by design: chr-leq is outside the profile (contract 1) and
		// paraconsistency assumes p and ¬p (contract 4)
		want := map[string]int{"chr-leq.yml": 2, "paraconsistency.yml": 3}
		files, _ := filepath.Glob(filepath.Join(dir, "examples", "AGs", "YAML", "*.yml"))
		probs := []string{}
		for _, f := range files {
			b, _ := os.ReadFile(f)
			r := cli.Run(string(b), "check")
			if r.Exit != want[filepath.Base(f)] {
				probs = append(probs, fmt.Sprintf("%s: exit %d, want %d (%s)", filepath.Base(f), r.Exit, want[filepath.Base(f)], firstLine(r.Stderr)))
			}
		}
		out = append(out, eres("cli-upstream-examples-check", len(files) > 0 && len(probs) == 0, "%s", strings.Join(probs, "; ")))
	}
	return out
}

// MigrationFixtureChecks migrates each graph of cases/migration, written for
// release R1, with tools/migrate_r2.py, and evaluates the graph before and
// after with the command (PLAN.md 5.6: graph fixtures migrated in a test).
// The migrated graph uses none of the old identifiers, both evaluations
// complete, and every statement has the same label under its new name.
// identifierText keeps the parts of a graph file that hold identifiers: it
// drops comments and the natural-language values of language: and
// statements:, where an old name may stand as an ordinary word ("has").
func identifierText(src string) string {
	var b strings.Builder
	section := ""
	keyed := regexp.MustCompile(`^(\s+("(?:[^"\\]|\\.)*"|'[^']*'|[^:"']+?))\s*:`)
	for _, ln := range strings.Split(src, "\n") {
		t := strings.TrimSpace(ln)
		if strings.HasPrefix(t, "#") {
			continue
		}
		if m := regexp.MustCompile(`^([\w-]+)\s*:`).FindStringSubmatch(ln); m != nil {
			section = m[1]
		} else if section == "language" || section == "statements" {
			if m := keyed.FindStringSubmatch(ln); m != nil {
				ln = m[1]
			}
		}
		b.WriteString(ln)
		b.WriteString("\n")
	}
	return b.String()
}

func MigrationFixtureChecks(cli *CLI) []Result {
	root := RepoRoot()
	b, err := os.ReadFile(filepath.Join(root, "tools", "migration_r2.yml"))
	if err != nil {
		return []Result{fail("tools/migration_r2.yml", "(file)", "migration/fixtures", "%v", err)}
	}
	var t struct {
		Changes []struct {
			Kind string `yaml:"kind"`
			Old  string `yaml:"old"`
			New  string `yaml:"new"`
		} `yaml:"changes"`
	}
	if err := yaml.Unmarshal(b, &t); err != nil {
		return []Result{fail("tools/migration_r2.yml", "(file)", "migration/fixtures", "%v", err)}
	}
	type ren struct {
		rx  *regexp.Regexp
		new string
	}
	rens := []ren{}
	for _, c := range t.Changes {
		o, n := c.Old, c.New
		switch c.Kind {
		case "rename-predicate":
			o, n = strings.Split(o, "/")[0], strings.Split(n, "/")[0]
		case "rename-scheme":
		default:
			continue
		}
		rens = append(rens, ren{regexp.MustCompile(`(^|[^\w-])` + regexp.QuoteMeta(o) + `($|[^\w-])`), n})
	}
	rename := func(s string) string {
		for _, r := range rens {
			s = r.rx.ReplaceAllString(s, "${1}"+r.new+"${2}")
		}
		return s
	}
	files, _ := filepath.Glob(filepath.Join("cases", "migration", "*.yml"))
	sort.Strings(files)
	out := []Result{}
	for _, f := range files {
		name := filepath.ToSlash(f)
		src, err := os.ReadFile(f)
		if err != nil {
			out = append(out, fail(name, "(file)", "migration/fixture", "%v", err))
			continue
		}
		cmd := exec.Command("python3", "-I", filepath.Join(root, "tools", "migrate_r2.py"), "--stdout", f)
		var stdout, stderr bytes.Buffer
		cmd.Stdout, cmd.Stderr = &stdout, &stderr
		if err := cmd.Run(); err != nil {
			out = append(out, fail(name, "(file)", "migration/fixture", "migrate_r2.py: %v %s", err, firstLine(stderr.String())))
			continue
		}
		migrated := stdout.String()
		probs := []string{}
		if ids := identifierText(migrated); rename(ids) != ids {
			probs = append(probs, "the migrated graph still uses an old identifier")
		}
		before, after := cli.Run(string(src), "eval", "-t", "yaml"), cli.Run(migrated, "eval", "-t", "yaml")
		if before.Exit != 0 || after.Exit != 0 {
			probs = append(probs, fmt.Sprintf("evaluation ended with %d before and %d after (%s)", before.Exit, after.Exit, firstLine(after.Stderr)))
		} else {
			l1, err1 := labels(before.Stdout)
			l2, err2 := labels(after.Stdout)
			if err1 != nil || err2 != nil {
				probs = append(probs, "an output is not YAML")
			}
			if len(l1) != len(l2) {
				probs = append(probs, fmt.Sprintf("%d statements before and %d after", len(l1), len(l2)))
			}
			for id, l := range l1 {
				if l2[Norm(rename(id))] != l && l2[rename(id)] != l {
					probs = append(probs, fmt.Sprintf("%s was %s and %s is %s", id, l, rename(id), l2[rename(id)]))
				}
			}
		}
		sort.Strings(probs)
		out = append(out, res(name, "(file)", "migration/fixture", probs))
	}
	if len(files) == 0 {
		out = append(out, fail("cases/migration", "(file)", "migration/fixture", "no fixture"))
	}
	return out
}

// DocExampleChecks (PLAN.md Phase 6, "document how the library is used"): each
// graph in docs/examples is what tools/graph_from_library.py makes from the
// library today for the schemes it names, followed by a case; the command
// evaluates it, and every label its tests: section expects is the label the
// command gives.
func DocExampleChecks(cli *CLI) []Result {
	root := RepoRoot()
	files, _ := filepath.Glob(filepath.Join(root, "docs", "examples", "*.yml"))
	sort.Strings(files)
	out := []Result{}
	for _, f := range files {
		name := "docs/examples/" + filepath.Base(f)
		src, err := os.ReadFile(f)
		if err != nil {
			out = append(out, fail(name, "(file)", "docs/example-current", "%v", err))
			continue
		}
		var doc struct {
			Meta    map[string]interface{} `yaml:"meta"`
			Schemes []struct {
				Id string `yaml:"id"`
			} `yaml:"argument_schemes"`
			Tests map[string][]string `yaml:"tests"`
		}
		if err := yaml.Unmarshal(src, &doc); err != nil {
			out = append(out, fail(name, "(file)", "docs/example-current", "not YAML: %v", err))
			continue
		}
		args := []string{"-I", filepath.Join(root, "tools", "graph_from_library.py")}
		for _, s := range doc.Schemes {
			args = append(args, s.Id)
		}
		if t, ok := doc.Meta["title"].(string); ok {
			args = append(args, "--title", t)
		}
		cmd := exec.Command("python3", args...)
		var gen, stderr bytes.Buffer
		cmd.Stdout, cmd.Stderr = &gen, &stderr
		probs := []string{}
		if err := cmd.Run(); err != nil {
			probs = append(probs, fmt.Sprintf("graph_from_library.py: %v %s", err, firstLine(stderr.String())))
		} else if !strings.HasPrefix(string(src), gen.String()) {
			probs = append(probs, "the example does not begin with what tools/graph_from_library.py makes from the library today: regenerate its schemes and language")
		}
		out = append(out, res(name, "(file)", "docs/example-current", probs))

		probs = []string{}
		r := cli.Run(string(src), "eval", "-t", "yaml")
		if r.Exit != 0 {
			probs = append(probs, fmt.Sprintf("evaluation ended with %d (%s)", r.Exit, firstLine(r.Stderr)))
		} else if l, err := labels(r.Stdout); err != nil {
			probs = append(probs, "the output is not YAML")
		} else {
			n := 0
			for want, ids := range doc.Tests {
				for _, id := range ids {
					n++
					if got := l[Norm(id)]; got != want && l[id] != want {
						probs = append(probs, fmt.Sprintf("%s is %s, the example expects %s", id, got, want))
					}
				}
			}
			if n == 0 {
				probs = append(probs, "the example expects no label: add tests")
			}
		}
		sort.Strings(probs)
		out = append(out, res(name, "(file)", "docs/example-labels", probs))
	}
	if len(files) == 0 {
		out = append(out, fail("docs/examples", "(file)", "docs/example-labels", "no example"))
	}
	return out
}
