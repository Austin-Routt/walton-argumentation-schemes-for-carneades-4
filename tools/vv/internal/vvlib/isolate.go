package vvlib

// Containment for checks that can hang or exhaust memory.
//
// A goroutine cannot be stopped from outside. A check that waits for the engine
// with a timeout therefore leaves the engine running when the timeout fires,
// and that worker goes on using memory, standard output and the engine's
// package-level state while later checks run. Two measures deal with it:
//
//   - the probes that are known to be able to run away (a theory that never
//     terminates, a term that keeps growing, a rule the matcher loops on) run in
//     a process of their own, with a time limit and a memory limit, and the
//     process is killed when either is exceeded;
//   - every other check that times out marks the run as tainted, and the test
//     binary stops instead of going on in a state it cannot vouch for.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync/atomic"
	"time"
)

var tainted atomic.Bool

// Taint records that an in-process worker was abandoned after a timeout.
func Taint() { tainted.Store(true) }

// Tainted reports whether an abandoned worker may still be running.
func Tainted() bool { return tainted.Load() }

// ProbeEnv names the environment variable that turns the test binary into a
// probe worker.
const ProbeEnv = "VV_PROBE"

const probeMarker = "VV-PROBE-RESULT "

// ProbeMemKB caps the address space of a probe process (kilobytes, as for
// `ulimit -v`). The probes need a few megabytes when the engine behaves.
const ProbeMemKB = 2000000

// Probes are the checks that run in a process of their own.
var Probes = map[string]func() []Result{
	"counter":         probeCounter,
	"growing-term":    probeGrowingTerm,
	"wrong-arity":     probeWrongArity,
	"issue-labelling": probeIssueLabelling,
}

// RunProbeChild runs one probe and prints its results; it is the whole job of
// a probe worker. The return value is the process exit status.
func RunProbeChild(name string) int {
	f, ok := Probes[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown probe %q\n", name)
		return 2
	}
	b, err := json.Marshal(f())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	fmt.Println(probeMarker + string(b))
	return 0
}

// Isolated runs a probe in a new process and returns its results. If the
// process is killed or dies, every check the probe was to report fails, with
// the reason.
func Isolated(name string, checks []string, timeout time.Duration) []Result {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	script := fmt.Sprintf(`ulimit -v %d 2>/dev/null; exec "$0" -test.run '^$'`, ProbeMemKB)
	cmd := exec.CommandContext(ctx, "sh", "-c", script, os.Args[0])
	cmd.Env = append(os.Environ(), ProbeEnv+"="+name)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	start := time.Now()
	err := cmd.Run()
	for _, line := range strings.Split(stdout.String(), "\n") {
		if strings.HasPrefix(line, probeMarker) {
			var res []Result
			if json.Unmarshal([]byte(strings.TrimPrefix(line, probeMarker)), &res) == nil {
				return res
			}
		}
	}
	why := ""
	switch {
	case ctx.Err() == context.DeadlineExceeded:
		why = fmt.Sprintf("the engine did not finish within %v and the process was stopped", timeout)
	case strings.Contains(stderr.String(), "out of memory") || strings.Contains(stderr.String(), "cannot allocate memory"):
		why = fmt.Sprintf("the engine ran out of memory under a cap of %d MB after %v", ProbeMemKB/1000, time.Since(start).Round(time.Second))
	default:
		why = fmt.Sprintf("the probe process ended without a result (%v): %s", err, firstLine(stderr.String()))
	}
	out := []Result{}
	for _, c := range checks {
		out = append(out, efail(c, "%s", why))
	}
	return out
}

// GrowingTerm builds a larger term with every rule application. Its store never
// repeats a fact, so no duplicate check ends it; only a limit can.
const GrowingTerm = `
language:
  t/1: "%s is a term."
argument_schemes:
  - id: grow
    variables: [X]
    premises: [t(X)]
    conclusions: [t(f(X))]
assumptions: [t(a)]
`

// WrongArity is modus ponens over statements, with one fact that uses the
// premise's functor at another arity.
const WrongArity = `
language:
  implies/2: "%s implies %s."
argument_schemes:
  - id: mp
    variables: [A, B]
    premises: ["implies(A,B)", A]
    conclusions: [B]
assumptions: ["implies(p,q)", p, "implies(x,y,z)"]
`

// IssueGrowth has an issue with two positions. Assuming a third statement and
// running Infer again adds a third position to it.
const IssueGrowth = `
language:
  colour/2: "The colour of %s is %s."
  says/3: "%s says %s is %s."
issue_schemes:
  colour: ["colour(O,C1)", "...", "colour(O,Cn)"]
argument_schemes:
  - id: s14
    variables: [W,O,C]
    premises:
      - "says(W,O,C)"
    conclusions:
      - "colour(O,C)"
assumptions:
  - "says(ann,car,red)"
  - "says(bob,car,blue)"
`

// probeIssueLabelling labels a graph whose issue gained a position on a second
// Infer. If labelling does not return, the parent stops this process and
// reports the check as failed.
func probeIssueLabelling() []Result {
	const name = "labelling-returns-after-issue-gains-position"
	e := EvalYAML(IssueGrowth)
	if !e.usable() {
		return []Result{efail(name, "the graph could not be set up %s", e.why())}
	}
	e.AG.Assumptions = append(e.AG.Assumptions, "says(cy,car,green)")
	c := StartCapture()
	var panicked any
	func() {
		defer func() { panicked = recover() }()
		e.AG.Infer()
		e.AG.GroundedLabelling()
	}()
	c.Stop()
	return []Result{eres(name, panicked == nil, "labelling the graph after a third position joined its issue panics: %v", panicked)}
}

func probeCounter() []Result {
	o, _, err, timedOut := inferOnly(Counter)
	out := []Result{}
	if timedOut {
		out = append(out, efail("budget-exhaustion-is-reported", "a scheme that counts upwards for ever did not come back within %v", InferTimeout))
	} else {
		out = append(out, eres("budget-exhaustion-is-reported", err != nil,
			"the rule engine gave up after its limit of rule applications and Infer returned no error"))
	}
	out = append(out, eres("engine-keeps-stdout-clean", strings.TrimSpace(o) == "",
		"the engine wrote to standard output, where the CLI writes its result: %q", firstLine(o)))
	return out
}

func probeGrowingTerm() []Result {
	_, _, err, timedOut := inferOnly(GrowingTerm)
	return []Result{eres("growing-term-is-bounded", !timedOut && err != nil,
		"a scheme that builds an ever larger term came back without an error (timed out: %v)", timedOut)}
}

func probeWrongArity() []Result {
	e := EvalYAML(WrongArity)
	exhausted := LimitReached(e.InferErr, e.Stdout)
	return []Result{eres("wrong-arity-fact-is-harmless", e.usable() && !exhausted && e.Label("q") == "in",
		"modus ponens over implies(p,q) and p, with a stray fact implies(x,y,z) in the graph: rule limit hit: %v, q is %s %s",
		exhausted, e.Label("q"), e.why())}
}
