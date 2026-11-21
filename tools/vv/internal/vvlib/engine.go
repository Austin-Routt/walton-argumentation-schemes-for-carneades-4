package vvlib

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/carneades/carneades-4/src/engine/caes"
	"github.com/carneades/carneades-4/src/engine/caes/encoding/dot"
	cyaml "github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
	"github.com/carneades/carneades-4/src/engine/validation"
)

// Engine checks state what the engine has to do for the scheme library to be
// usable in practice. Each one is a small, self-contained experiment against
// the unmodified engine; none of them depends on the scheme files.

const engine = "(engine)"

func epass(c string) Result { return Result{Scheme: engine, Check: c, Pass: true} }
func efail(c, d string, a ...any) Result {
	return Result{Scheme: engine, Check: c, Detail: fmt.Sprintf(d, a...)}
}
func eres(c string, ok bool, d string, a ...any) Result {
	if ok {
		return epass(c)
	}
	return efail(c, d, a...)
}

// Eval imports a YAML argument graph, constructs arguments and labels it, the
// way `carneades eval` does after validation.
type Eval struct {
	AG        *caes.ArgGraph
	L         caes.Labelling
	ImportErr error
	InferErr  error
	TimedOut  bool
	Panic     any
	Stdout    string
	Stderr    string
}

func EvalYAML(text string) *Eval {
	ResetEngineGlobals()
	e := &Eval{}
	c := StartCapture()
	defer func() {
		o, s := c.Stop()
		e.Stdout += o
		e.Stderr += s
	}()
	e.AG, e.ImportErr = cyaml.Import(strings.NewReader(text))
	if e.ImportErr != nil {
		return e
	}
	done := make(chan error, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				e.Panic = p
				done <- fmt.Errorf("panic: %v", p)
			}
		}()
		done <- e.AG.Infer()
	}()
	select {
	case e.InferErr = <-done:
	case <-time.After(InferTimeout):
		e.TimedOut = true
		Taint()
		return e
	}
	e.L = e.AG.GroundedLabelling()
	return e
}

func (e *Eval) Label(id string) string {
	if e.AG == nil || e.L == nil {
		return "missing"
	}
	s, ok := e.AG.Statements[Norm(id)]
	if !ok {
		return "missing"
	}
	return e.L[s].String()
}

func (e *Eval) usable() bool {
	return e.ImportErr == nil && e.InferErr == nil && !e.TimedOut && e.AG != nil
}

func (e *Eval) why() string {
	switch {
	case e.ImportErr != nil:
		return "import: " + e.ImportErr.Error()
	case e.TimedOut:
		return "inference did not return"
	case e.InferErr != nil:
		return "infer: " + e.InferErr.Error()
	}
	return ""
}

// MutualRecursion: two schemes that derive each other's premise, followed by an
// unrelated scheme. With a store that kept one copy of each fact this would
// stop after two rule applications.
const MutualRecursion = `
language:
  p/1: "%s is p."
  q/1: "%s is q."
  r/1: "%s is r."
  s/1: "%s is s."
argument_schemes:
  - id: fwd
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
  - id: back
    variables: [X]
    premises: [q(X)]
    conclusions: [p(X)]
  - id: later
    variables: [X]
    premises: [r(X)]
    conclusions: [s(X)]
assumptions: [p(a), r(a)]
`

// Counter never terminates, whatever the store semantics, and stays small: the
// scheme replaces n(X) by n(X+1), so the store holds one number at a time.
const Counter = `
language:
  n/1: "%s is a number."
argument_schemes:
  - id: count
    variables: []
    deletions: [n(X)]
    conclusions: [n(X+1)]
assumptions: [n(0)]
`

// inferOnly imports a graph and constructs arguments without labelling it.
func inferOnly(text string) (stdout, stderr string, err error, timedOut bool) {
	ResetEngineGlobals()
	c := StartCapture()
	ag, ierr := cyaml.Import(strings.NewReader(text))
	if ierr != nil {
		o, e := c.Stop()
		return o, e, ierr, false
	}
	done := make(chan error, 1)
	go func() { done <- ag.Infer() }()
	select {
	case err = <-done:
	case <-time.After(InferTimeout):
		timedOut = true
		Taint()
	}
	stdout, stderr = c.Stop()
	ResetEngineGlobals()
	return
}

const baseTheory = `
language:
  p/1: "%s is p."
  q/1: "%s is q."
  r/1: "%s is r."
  p/2: "%s p %s."
  q/2: "%s q %s."
`

// EngineChecks runs the API-level engine checks.
func EngineChecks() []Result {
	out := []Result{}

	// --- what works today and has to keep working -------------------------
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s1
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
    exceptions: [r(X)]
assumptions: [p(a), p(b), r(b)]
`)
		ok := e.usable() && e.Label("q(a)") == "in" && e.Label("q(b)") == "out"
		out = append(out, eres("baseline-scheme-with-exception", ok,
			"expected q(a) in and q(b) out (undercut by r(b)); got q(a) %s, q(b) %s %s", e.Label("q(a)"), e.Label("q(b)"), e.why()))
	}
	{
		// the generic negation issue used by upstream's self-defeat example
		e := EvalYAML(baseTheory + `
issue_schemes:
  negation: ["¬Q", "Q"]
argument_schemes:
  - id: pro
    variables: [X]
    premises: [p(X)]
    conclusions: [r(X)]
  - id: con
    variables: [X]
    premises: [q(X)]
    conclusions: ["¬r(X)"]
assumptions: [p(a), q(a), p(b)]
`)
		ok := e.usable() && e.Label("r(a)") == "out" && e.Label("¬r(a)") == "out" && e.Label("r(b)") == "in"
		out = append(out, eres("baseline-negation-issue", ok,
			"expected r(a) and ¬r(a) both out (equal weight), r(b) in; got r(a) %s, ¬r(a) %s, r(b) %s %s",
			e.Label("r(a)"), e.Label("¬r(a)"), e.Label("r(b)"), e.why()))
	}
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s1
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
assumptions: [p(a)]
`)
		ok := false
		detail := e.why()
		if e.usable() {
			e.AG.ApplyLabelling(e.L)
			var buf bytes.Buffer
			if err := dot.Export(&buf, e.AG); err != nil {
				detail = "dot.Export: " + err.Error()
			} else if _, err := exec.LookPath("dot"); err != nil {
				ok = strings.Contains(buf.String(), "digraph")
				detail = "no digraph in the DOT output"
			} else {
				cmd := exec.Command("dot", "-Tsvg")
				cmd.Stdin = &buf
				svg, err := cmd.CombinedOutput()
				ok = err == nil && bytes.Contains(svg, []byte("<svg"))
				detail = fmt.Sprintf("Graphviz rejected the DOT output: %v %.200s", err, svg)
			}
		}
		out = append(out, eres("baseline-dot-export-renders", ok, "%s", detail))
	}

	// --- global state -------------------------------------------------------
	{
		ResetEngineGlobals()
		c := StartCapture()
		_, err := cyaml.Import(strings.NewReader(baseTheory + "argument_schemes:\n  - id: s1\n    variables: [X]\n    premises: [p(X)]\n    conclusions: [q(X)]\n"))
		c.Stop()
		extra := []string{}
		for k := range caes.BasicSchemes {
			if !basicNames[k] {
				extra = append(extra, k)
			}
		}
		ResetEngineGlobals()
		out = append(out, eres("yaml-import-keeps-globals-clean", err == nil && len(extra) == 0,
			"after importing a file, caes.BasicSchemes also contains %v: imports leak into one another and are not safe to run concurrently", extra))
	}

	// --- the engine's validator and arguments ---------------------------------
	argFile := func(params string) string {
		return baseTheory + `
statements:
  p(a,b): A p B.
  q(a,b): A q B.
argument_schemes:
  - id: s2
    variables: [X,Y]
    premises: ["p(X,Y)"]
    conclusions: ["q(X,Y)"]
arguments:
  a1:
    scheme: s2
    parameters: ` + params + `
    premises: ["p(a,b)"]
    conclusion: q(a,b)
`
	}
	validate := func(text string, cleanGlobals bool) (probs []validation.Problem, panicked any, err error) {
		ResetEngineGlobals()
		c := StartCapture()
		defer c.Stop()
		ag, err := cyaml.Import(strings.NewReader(text))
		if err != nil {
			return nil, nil, err
		}
		if cleanGlobals {
			ResetEngineGlobals()
		}
		defer func() { panicked = recover() }()
		probs = validation.Validate(ag)
		return
	}
	{
		probs, p, err := validate(argFile("[a]"), false)
		found := false
		for _, pr := range probs {
			if pr.Category == validation.ARGUMENT {
				found = true
			}
		}
		out = append(out, eres("validator-checks-arguments-against-schemes", err == nil && p == nil && found,
			"an argument with one parameter for a two-variable scheme passes validation (problems: %d, panic: %v, import error: %v)", len(probs), p, err))
	}
	{
		probs, p, err := validate(argFile("[a, b]"), true)
		n := 0
		for _, pr := range probs {
			if pr.Category == validation.ARGUMENT {
				n++
			}
		}
		out = append(out, eres("validator-accepts-correct-argument", err == nil && p == nil && n == 0,
			"validating a correctly instantiated argument (with the leaked global scheme table cleared): panic %v, %d argument problem(s), import error %v", p, n, err))
	}
	{
		probs, p, err := validate(baseTheory+`
statements:
  p(a): A is p.
  q(b): B is q.
argument_schemes:
  - id: s3
    variables: [X]
    premises: ["p(X)"]
    conclusions: ["q(X)"]
arguments:
  a1:
    scheme: s3
    parameters: [a]
    premises: ["p(a)"]
    conclusion: q(b)
`, true)
		found := false
		for _, pr := range probs {
			if pr.Category == validation.ARGUMENT {
				found = true
			}
		}
		out = append(out, eres("validator-enforces-shared-bindings", err == nil && p == nil && found,
			"scheme p(X) => q(X): a hand-written argument with premise p(a) and conclusion q(b) is not reported (argument problems found: %v, panic: %v, import error: %v)", found, p, err))
	}

	// --- round trips -----------------------------------------------------------
	twoParam := baseTheory + `
argument_schemes:
  - id: s2
    variables: [X,Y]
    premises: ["p(X,Y)"]
    conclusions: ["q(X,Y)"]
assumptions: ["p(a,b)"]
`
	{
		e := EvalYAML(twoParam)
		ok := false
		detail := e.why()
		if e.usable() {
			e.AG.ApplyLabelling(e.L)
			var buf bytes.Buffer
			c := StartCapture()
			cyaml.Export(&buf, e.AG)
			c.Stop()
			ResetEngineGlobals() // a new process starts with clean globals
			c = StartCapture()
			_, err := cyaml.Import(bytes.NewReader(buf.Bytes()))
			c.Stop()
			ResetEngineGlobals()
			ok = err == nil
			if err != nil {
				detail = "the engine cannot read the YAML it wrote: " + strings.TrimSpace(err.Error())
			}
		}
		out = append(out, eres("yaml-export-can-be-reimported", ok, "%s", detail))
	}
	{
		e := EvalYAML(twoParam)
		ok := false
		detail := e.why()
		if e.usable() {
			n1 := len(e.AG.Arguments)
			c := StartCapture()
			e.AG.Infer()
			c.Stop()
			n2 := len(e.AG.Arguments)
			ok = n1 == n2
			detail = fmt.Sprintf("a second Infer on the same graph raises the number of arguments from %d to %d", n1, n2)
		}
		out = append(out, eres("infer-is-idempotent", ok, "%s", detail))
	}

	// --- terms -------------------------------------------------------------------
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s3
    variables: [P,S]
    premises: ["p(P,S)"]
    conclusions: ["q(P,S)"]
assumptions:
  - 'p(joe,"it rains")'
`)
		ok := false
		detail := e.why()
		if e.usable() {
			got := []string{}
			for id, st := range e.AG.Statements {
				got = append(got, id)
				if strings.HasPrefix(id, "q(joe,") && strings.Contains(id, "it rains") && e.L[st] == caes.In {
					ok = true
				}
			}
			sort.Strings(got)
			detail = fmt.Sprintf("the assumption p(joe,\"it rains\") is read as %q and no conclusion about \"it rains\" is derived; statements: %v", e.AG.Assumptions, got)
		}
		out = append(out, eres("string-constants-survive", ok, "%s", detail))
	}
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s4
    variables: [P,D]
    premises: ["p(P,D)"]
    conclusions: ["q(P,D)"]
assumptions:
  - p(joe,1990-05-17)
`)
		got := []string{}
		if e.usable() {
			for id := range e.AG.Statements {
				if strings.HasPrefix(id, "q(") {
					got = append(got, id+" "+e.Label(id))
				}
			}
		}
		ok := e.usable() && e.Label("q(joe,1990-05-17)") == "in"
		out = append(out, eres("hyphenated-constants-are-not-computed", ok,
			"from p(joe,1990-05-17) the engine derives %v: the rule engine evaluates 1990-05-17 as a subtraction %s", got, e.why()))
	}

	// --- rule engine ---------------------------------------------------------------
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s6
    variables: [X]
    premises: ["¬p(X)"]
    conclusions: [r(X)]
assumptions: ["¬p(a)"]
`)
		out = append(out, eres("negated-premises-match", e.usable() && e.Label("r(a)") == "in",
			"a scheme with the premise ¬p(X) does not fire on the assumption ¬p(a): r(a) is %s %s", e.Label("r(a)"), e.why()))
	}
	{
		// A negation that reaches a rule through a statement variable. Both other
		// reviews found that a fix which rewrites ¬ in the scheme text alone
		// misses this path.
		e := EvalYAML(`
language:
  says/2: "%s says that %s."
  p/1: "%s is p."
  r/1: "%s is r."
argument_schemes:
  - id: s6a
    variables: [W,S]
    premises: ["says(W,S)"]
    conclusions: [S]
  - id: s6b
    variables: [X]
    premises: ["¬p(X)"]
    conclusions: [r(X)]
assumptions: ["says(ann,¬p(a))"]
`)
		out = append(out, eres("negation-through-statement-variable", e.usable() && e.Label("r(a)") == "in",
			"says(ann,¬p(a)) yields ¬p(a) through a statement variable (¬p(a) is %s), and the scheme with the premise ¬p(X) does not fire on it: r(a) is %s %s", e.Label("¬p(a)"), e.Label("r(a)"), e.why()))
	}
	{
		// A user predicate that happens to be called neg is an ordinary predicate:
		// it must not be taken for the engine's own encoding of ¬.
		e := EvalYAML(`
language:
  neg/1: "%s is negative."
  p/1: "%s is p."
  r/1: "%s is r."
argument_schemes:
  - id: s6c
    variables: [X]
    premises: ["¬p(X)"]
    conclusions: [r(X)]
assumptions: ["neg(p(a))"]
`)
		n := 0
		if e.usable() {
			for _, a := range e.AG.Arguments {
				if a.Scheme != nil && a.Scheme.Id == "s6c" {
					n++
				}
			}
		}
		out = append(out, eres("baseline-user-predicate-neg-is-ordinary", e.usable() && n == 0,
			"the fact neg(p(a)), a statement about a user predicate neg/1, makes the scheme with the premise ¬p(X) construct %d argument(s) %s", n, e.why()))
	}
	{
		e := EvalYAML(`
language:
  p/2: "%s p %s."
  is/2: "%s is %s."
argument_schemes:
  - id: s7
    variables: [X,Y]
    premises: ["p(X,Y)"]
    conclusions: ["is(X,Y)"]
assumptions: ["p(a,b)"]
`)
		out = append(out, eres("is-as-predicate-in-conclusion", e.usable() && e.Label("is(a,b)") == "in",
			"a scheme concluding is(X,Y) constructs nothing: is(a,b) is %s %s", e.Label("is(a,b)"), e.why()))
	}
	loop := MutualRecursion
	{
		e := EvalYAML(loop)
		exhausted := LimitReached(e.InferErr, e.Stdout)
		out = append(out, eres("rederivation-terminates", e.usable() && !exhausted,
			"two schemes that derive each other's premise run until the limit of %d rule applications %s", caes.MAXRULEAPPS, e.why()))
		out = append(out, eres("no-scheme-is-starved", e.usable() && e.Label("s(a)") == "in",
			"a scheme listed after two mutually recursive schemes never gets its turn: from r(a) it should conclude s(a), which is %s", e.Label("s(a)")))
	}
	// Theories that can run away are evaluated in a process of their own, with a
	// time limit and a memory limit (isolate.go).
	// A counter: genuinely infinite, whatever the store semantics, and small.
	out = append(out, Isolated("counter", []string{"budget-exhaustion-is-reported", "engine-keeps-stdout-clean"}, 60*time.Second)...)
	// A term that grows with every rule application: the limit on rule
	// applications bounds neither the time nor the memory this takes.
	out = append(out, Isolated("growing-term", []string{"growing-term-is-bounded"}, 20*time.Second)...)
	// A fact with the functor of a premise and another arity.
	out = append(out, Isolated("wrong-arity", []string{"wrong-arity-fact-is-harmless"}, 30*time.Second)...)
	{
		e := EvalYAML(`
language:
  implies/2: "%s implies %s."
argument_schemes:
  - id: mp2
    variables: [A, B]
    premises: [A, "implies(A,B)"]
    conclusions: [B]
assumptions: ["implies(p,q)", p]
`)
		out = append(out, eres("premise-order-does-not-matter", e.usable() && e.Label("q") == "in",
			"a scheme with the premises A, implies(A,B) constructs nothing from implies(p,q) and p, although the same scheme with the premises in the other order does: q is %s %s", e.Label("q"), e.why()))
	}
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s8
    variables: [A, A1]
    premises: [p(A1)]
    conclusions: [q(A)]
assumptions: [p(a)]
`)
		out = append(out, eres("unbound-variable-is-not-captured", e.usable() && e.Label("q(a)") != "in",
			"scheme variables [A, A1], premise p(A1), conclusion q(A): from p(a) the engine concludes q(a), because the fresh name it gives the unbound A is A1 %s", e.why()))
	}

	// --- arguments written by hand that name a scheme ---------------------------------
	expert := `
language:
  expert/2: "%s is an expert in the %s domain."
  in_domain/2: "%s is in the domain of %s."
  asserts/2: "%s asserts that %s is true."
  untrustworthy/1: "%s is untrustworthy."
  safe/1: "%s is safe."
argument_schemes:
  - id: expert_opinion
    variables: [W,D,S]
    premises:
      - expert(W,D)
      - in_domain(S,D)
      - asserts(W,S)
    exceptions:
      - untrustworthy(W)
    conclusions:
      - S
`
	{
		e := EvalYAML(expert + `
statements:
  expert(ann,vaccines): Ann is an expert on vaccines.
  in_domain(safe(vaccine),vaccines): Vaccine safety is a question about vaccines.
  asserts(ann,safe(vaccine)): Ann says the vaccine is safe.
  safe(vaccine): The vaccine is safe.
  untrustworthy(ann): Ann is untrustworthy.
arguments:
  a1:
    scheme: expert_opinion
    parameters: [ann, vaccines, safe(vaccine)]
    premises:
      - expert(ann,vaccines)
      - in_domain(safe(vaccine),vaccines)
      - asserts(ann,safe(vaccine))
    conclusion: safe(vaccine)
assumptions:
  - expert(ann,vaccines)
  - in_domain(safe(vaccine),vaccines)
  - asserts(ann,safe(vaccine))
  - untrustworthy(ann)
`)
		n := 0
		if e.usable() {
			for _, a := range e.AG.Arguments {
				if a.Scheme != nil && a.Scheme.Id == "expert_opinion" {
					n++
				}
			}
		}
		out = append(out, eres("explicit-argument-gets-critical-questions", e.usable() && e.Label("safe(vaccine)") == "out",
			"an argument written out in the file with `scheme: expert_opinion` gets no undercutter, so the exception untrustworthy(ann) does not touch it: safe(vaccine) is %s (the graph now holds %d expert_opinion arguments for the same parameters) %s",
			e.Label("safe(vaccine)"), n, e.why()))
	}
	{
		e := EvalYAML(expert + `
arguments:
  a1:
    scheme: expert_opinion
    parameters: [ann, vaccines, safe(vaccine)]
`)
		out = append(out, eres("argument-from-scheme-and-parameters", e.usable(),
			"USAGE.md says a scheme named by an argument \"can also be used to infer conclusion or premises\"; an argument with only scheme and parameters is rejected (%s)", strings.TrimSpace(e.why())))
	}

	// --- evaluation --------------------------------------------------------------------
	{
		e := EvalYAML(`
statements:
  p: P is the case.
  q: Q is the case.
issues:
  i1:
    positions: [p, q]
assumptions: [p, q]
`)
		ok := e.usable() && e.AG.Inconsistent()
		out = append(out, eres("inconsistent-assumptions-detected", ok,
			"both positions of an issue are assumed and ArgGraph.Inconsistent() answers false %s", e.why()))
	}
	{
		ResetEngineGlobals()
		c := StartCapture()
		ag, err := cyaml.Import(strings.NewReader(`
statements:
  p: P is the case.
  q: Q is the case.
issues:
  i1:
    positions: [p, q]
assumptions: [p, q]
`))
		before, after, withdrawn := false, false, true
		if err == nil {
			before = ag.Inconsistent()
			ag.GroundedLabelling()
			after = ag.Inconsistent()
			ag.Assumptions = []string{"p"}
			withdrawn = ag.Inconsistent()
		}
		c.Stop()
		ResetEngineGlobals()
		out = append(out, eres("inconsistent-follows-current-assumptions", err == nil && before && after && !withdrawn,
			"two positions of one issue are assumed. Inconsistent() answers %v before labelling and %v after it (both should be true), and %v once one of the two assumptions has been withdrawn (should be false): the answer comes from a cache that labelling fills, not from the assumptions as they are (import error: %v)", before, after, withdrawn, err))
	}
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s9
    variables: [X]
    roles: [major]
    premises: [p(X)]
    conclusions: [q(X)]
assumptions: [p(a)]
`)
		role := "(no argument)"
		if e.usable() {
			for _, a := range e.AG.Arguments {
				if a.Scheme != nil && a.Scheme.Id == "s9" && len(a.Premises) > 0 {
					role = a.Premises[0].Role
				}
			}
		}
		out = append(out, eres("scheme-roles-are-imported", role == "major",
			"USAGE.md documents `roles` for schemes, but the importer has no field for it: the premise role is %q", role))
	}

	// --- found by the two other reviews of 2026-11-20, confirmed and added here ---------
	// Critical questions modelled as assumptions. The book (ch. 11, sec. 5,
	// pp. 373-375) has an assumption hold until the other side calls it into
	// question; from then on it needs support like an ordinary premise. The
	// checks below state that contract. `questioned:` is the key proposed for
	// it in decision D2 of the plan; the engine does not know it yet.
	questionable := `
language:
  p/1: "%s is p."
  a/1: "%s is a."
  c/1: "%s is c."
  e/1: "%s is e."
argument_schemes:
  - id: s10
    variables: [X]
    premises: [p(X)]
    assumptions: [a(X)]
    conclusions: [c(X)]
  - id: s11
    variables: [X]
    premises: [e(X)]
    conclusions: [a(X)]
`
	{
		e := EvalYAML(questionable + "assumptions: [p(x)]\n")
		out = append(out, eres("baseline-undisputed-assumption-holds", e.usable() && e.Label("a(x)") == "in" && e.Label("c(x)") == "in",
			"a scheme assumption nobody has questioned should hold: a(x) is %s, c(x) is %s %s", e.Label("a(x)"), e.Label("c(x)"), e.why()))
	}
	{
		e := EvalYAML(questionable + "assumptions: [p(x)]\nquestioned: [a(x)]\n")
		out = append(out, eres("questioned-assumption-is-not-assumed", e.usable() && e.Label("a(x)") != "in" && e.Label("c(x)") != "in",
			"a(x) is listed as questioned and nothing supports it, yet a(x) is %s and the scheme's conclusion c(x) is %s. Every scheme assumption is assumed when the scheme is instantiated, and assumed statements are labelled in before any argument is looked at: a file has no way to call a critical question written under `assumptions:` into question %s",
			e.Label("a(x)"), e.Label("c(x)"), e.why()))
	}
	{
		e := EvalYAML(questionable + "assumptions: [p(x), e(x)]\nquestioned: [a(x)]\n")
		out = append(out, eres("questioned-assumption-holds-when-supported", e.usable() && e.Label("a(x)") == "in" && e.Label("c(x)") == "in",
			"a(x) is questioned and an argument from e(x) supports it, so the scheme's conclusion should hold: a(x) is %s, c(x) is %s %s", e.Label("a(x)"), e.Label("c(x)"), e.why()))
	}
	{
		// Through the Go API an assumption can be withdrawn by hand. That has to
		// survive the next Infer; today it does only when the argument has one parameter.
		lab := func(ag *caes.ArgGraph, id string) string {
			l := ag.GroundedLabelling()
			if st, ok := ag.Statements[Norm(id)]; ok {
				return l[st].String()
			}
			return "missing"
		}
		withdraw := func(text, assumption, conclusion string) (first, second string, problem any) {
			ResetEngineGlobals()
			c := StartCapture()
			defer func() {
				if p := recover(); p != nil {
					problem = p
				}
				c.Stop()
				ResetEngineGlobals()
			}()
			ag, err := cyaml.Import(strings.NewReader(text))
			if err != nil {
				return "", "", err
			}
			if err := ag.Infer(); err != nil {
				return "", "", err
			}
			first = lab(ag, conclusion)
			keep := []string{}
			for _, k := range ag.Assumptions {
				if Norm(k) != Norm(assumption) {
					keep = append(keep, k)
				}
			}
			ag.Assumptions = keep
			if err := ag.Infer(); err != nil {
				return first, "", err
			}
			return first, lab(ag, conclusion), nil
		}
		f1, s1, p1 := withdraw(questionable+"assumptions: [p(x)]\n", "a(x)", "c(x)")
		out = append(out, eres("baseline-withdrawn-assumption-stays-withdrawn", p1 == nil && f1 == "in" && s1 != "in",
			"one parameter: c(x) should be in, and no longer in once a(x) is withdrawn and Infer has run again; got %s then %s (problem: %v)", f1, s1, p1))
		f2, s2, p2 := withdraw(baseTheory+`
argument_schemes:
  - id: s13
    variables: [X, Y]
    premises: ["p(X,Y)"]
    assumptions: ["r(X)"]
    conclusions: ["q(X,Y)"]
assumptions: ["p(a,b)"]
`, "r(a)", "q(a,b)")
		out = append(out, eres("withdrawn-assumption-survives-second-infer", p2 == nil && f2 == "in" && s2 != "in",
			"two parameters: q(a,b) is %s, and after r(a) is withdrawn and Infer runs again it is %s (problem: %v). The second Infer does not recognise the argument it built before, instantiates the scheme again and assumes r(a) again", f2, s2, p2))
	}
	{
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s12
    variables: [X]
    premises: [p(X)]
    exceptions: [r(X)]
    conclusions: [q(X)]
assumptions: [p(a)]
`)
		ok := false
		detail := e.why()
		if e.usable() {
			n1 := len(e.AG.Arguments)
			var panicked any
			var err error
			func() {
				defer func() { panicked = recover() }()
				c := StartCapture()
				defer c.Stop()
				err = e.AG.Infer()
			}()
			n2 := len(e.AG.Arguments)
			ok = panicked == nil && err == nil && n1 == n2
			detail = fmt.Sprintf("a second Infer on a graph whose scheme has an exception: panic %v, error %v, arguments %d then %d (the undercutter built for the exception has no scheme, and Infer reads its scheme id)", panicked, err, n1, n2)
		}
		out = append(out, eres("infer-twice-with-exception", ok, "%s", detail))
	}
	{
		// one parameter that is itself a statement with two arguments
		e := EvalYAML(`
language:
  asserts/2: "%s asserts %s."
  liable/2: "%s is liable to %s."
argument_schemes:
  - id: s13
    variables: [W,S]
    premises:
      - "asserts(W,S)"
    conclusions:
      - S
assumptions:
  - "asserts(w,liable(acme,smith))"
tests:
  in:
    - "liable(acme,smith)"
`)
		okParams, okTests := false, false
		dParams, dTests := e.why(), e.why()
		if e.usable() {
			before, wantTests := "", len(e.AG.ExpectedLabeling)
			for _, a := range e.AG.Arguments {
				if a.Scheme != nil && a.Scheme.Id == "s13" {
					before = strings.Join(a.Parameters, " | ")
				}
			}
			e.AG.ApplyLabelling(e.L)
			var buf bytes.Buffer
			c := StartCapture()
			cyaml.Export(&buf, e.AG)
			// no reset here: until the export carries the theory, only the scheme
			// table left behind by the first import lets the file be read at all
			ag2, err := cyaml.Import(bytes.NewReader(buf.Bytes()))
			c.Stop()
			ResetEngineGlobals()
			if err != nil {
				dParams = "the exported file cannot be read back: " + firstLine(err.Error())
				dTests = dParams
			} else {
				after := ""
				for _, a := range ag2.Arguments {
					if a.Scheme != nil && a.Scheme.Id == "s13" {
						after = strings.Join(a.Parameters, " | ")
					}
				}
				okParams = before != "" && before == after
				dParams = fmt.Sprintf("the argument's parameters [%s] come back as [%s]: the export writes them into a YAML flow list without quotes, so a comma inside a parameter splits it", before, after)
				okTests = wantTests > 0 && len(ag2.ExpectedLabeling) == wantTests
				dTests = fmt.Sprintf("%d expected label(s) under `tests:` before the export, %d after reading the export back: the export writes them under `labels:`, a key the importer does not read", wantTests, len(ag2.ExpectedLabeling))
			}
		}
		out = append(out, eres("export-keeps-compound-parameters", okParams, "%s", dParams))
		out = append(out, eres("export-keeps-expected-labels", okTests, "%s", dTests))
	}
	{
		// an issue that gains a position when inference runs a second time
		e := EvalYAML(IssueGrowth)
		ok := false
		detail := e.why()
		if e.usable() {
			e.AG.Assumptions = append(e.AG.Assumptions, "says(cy,car,green)")
			var panicked any
			func() {
				defer func() { panicked = recover() }()
				c := StartCapture()
				defer c.Stop()
				e.AG.Infer()
			}()
			st := e.AG.Statements[Norm("colour(car,green)")]
			listed := false
			if st != nil && st.Issue != nil {
				for _, p := range st.Issue.Positions {
					if p == st {
						listed = true
					}
				}
			}
			ok = panicked == nil && st != nil && st.Issue != nil && listed
			detail = fmt.Sprintf("a third position joins an existing issue on a second Infer (panic: %v): the new statement points at the issue (%v) but the issue does not list it among its positions (%v). Labelling such a graph does not return: the issue is resolved again and again without ever labelling the new statement",
				panicked, st != nil && st.Issue != nil, listed)
		}
		out = append(out, eres("issue-extension-keeps-positions", ok, "%s", detail))
	}
	// Labelling the same graph. On the engine as shipped it never returns, so
	// the probe runs in a process of its own and is stopped after ten seconds.
	out = append(out, Isolated("issue-labelling", []string{"labelling-returns-after-issue-gains-position"}, 10*time.Second)...)
	{
		// a premise the term parser cannot read
		e := EvalYAML(baseTheory + `
argument_schemes:
  - id: s15
    variables: [X]
    premises:
      - p(X)
      - "broken("
    conclusions:
      - q(X)
assumptions: [p(a)]
`)
		reported := e.ImportErr != nil || e.InferErr != nil
		n := 0
		if e.AG != nil {
			n = len(e.AG.Arguments)
		}
		out = append(out, eres("malformed-statement-is-an-error", reported,
			"a scheme with the unreadable premise `broken(` is imported and run without any error: %d argument(s) constructed, q(a) is %s. A statement the parser cannot read is dropped with a line on stderr, GoCHR accepts malformed rule text, and Carneades discards the error AddRule returns",
			n, e.Label("q(a)")))
	}
	return out
}

func firstLine(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		s = s[:i]
	}
	if len(s) > 160 {
		s = s[:160] + "…"
	}
	return s
}

// ---------------------------------------------------------------- toolchain checks

// EngineDir is the Carneades source tree: third_party/carneades-4 at the
// repository root.
func EngineDir() (string, error) {
	wd, _ := os.Getwd()
	for d := wd; d != filepath.Dir(d); d = filepath.Dir(d) {
		p := filepath.Join(d, "third_party", "carneades-4")
		if _, err := os.Stat(filepath.Join(p, "go.mod")); err == nil {
			return p, nil
		}
	}
	return "", fmt.Errorf("third_party/carneades-4 not found above the working directory")
}

// engineModfile writes a go.mod for building the Carneades module itself
// against the GoCHR snapshot of this repository. $VV_GOMOD_EXTRA may name a
// file with further directives (used where module downloads are restricted).
func engineModfile(dir string) (string, error) {
	b, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return "", err
	}
	gochr, _ := filepath.Abs(filepath.Join(dir, "..", "GoCHR"))
	mod := string(b) + "\nreplace github.com/hfried/GoCHR => " + gochr + "\n"
	if extra := os.Getenv("VV_GOMOD_EXTRA"); extra != "" {
		x, err := os.ReadFile(extra)
		if err != nil {
			return "", err
		}
		mod += "\n" + string(x) + "\n"
	}
	tmp, err := os.MkdirTemp("", "vv-mod-*")
	if err != nil {
		return "", err
	}
	sum, _ := os.ReadFile(filepath.Join(dir, "go.sum"))
	os.WriteFile(filepath.Join(tmp, "go.sum"), sum, 0o644)
	p := filepath.Join(tmp, "go.mod")
	return p, os.WriteFile(p, []byte(mod), 0o644)
}

func goCmd(dir, modfile string, args ...string) (string, error) {
	full := append([]string{args[0], "-modfile=" + modfile}, args[1:]...)
	cmd := exec.Command("go", full...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod")
	b, err := cmd.CombinedOutput()
	return string(b), err
}

// ToolchainChecks builds, vets and tests the Carneades module as shipped.
func ToolchainChecks() []Result {
	dir, err := EngineDir()
	if err != nil {
		return []Result{efail("build-all-packages", "%v", err)}
	}
	mod, err := engineModfile(dir)
	if err != nil {
		return []Result{efail("build-all-packages", "%v", err)}
	}
	defer os.RemoveAll(filepath.Dir(mod))
	out := []Result{}

	o, err := goCmd(dir, mod, "build", "./...")
	bad := []string{}
	for _, l := range strings.Split(o, "\n") {
		if strings.HasPrefix(l, "# ") {
			bad = append(bad, strings.TrimPrefix(strings.TrimPrefix(l, "# "), "github.com/carneades/carneades-4/"))
		}
	}
	out = append(out, eres("build-all-packages", err == nil, "packages that do not compile: %s", strings.Join(bad, ", ")))

	pkgs := []string{}
	lo, _ := goCmd(dir, mod, "list", "./...")
	for _, p := range strings.Fields(lo) {
		skip := false
		for _, b := range bad {
			if strings.HasSuffix(p, b) {
				skip = true
			}
		}
		if !skip && strings.HasPrefix(p, "github.com/") {
			pkgs = append(pkgs, p)
		}
	}
	o, err = goCmd(dir, mod, append([]string{"vet"}, pkgs...)...)
	n := 0
	for _, l := range strings.Split(o, "\n") {
		if strings.Contains(l, ".go:") {
			n++
		}
	}
	out = append(out, eres("go-vet-clean", err == nil, "go vet reports %d finding(s) in the packages that compile", n))

	o, err = goCmd(dir, mod, append([]string{"test", "-count=1"}, pkgs...)...)
	failed := []string{}
	for _, l := range strings.Split(o, "\n") {
		if strings.HasPrefix(l, "FAIL\t") {
			failed = append(failed, strings.TrimPrefix(strings.Fields(l)[1], "github.com/carneades/carneades-4/"))
		}
	}
	out = append(out, eres("upstream-tests-pass-as-shipped", err == nil,
		"`go test ./...` fails in: %s (vet findings stop two packages; the CouchDB test aborts the rest unless a database listens on 127.0.0.1:5984)", strings.Join(failed, ", ")))

	o, err = goCmd(dir, mod, append([]string{"test", "-count=1", "-vet=off", "-skip", "TestCouchDB"}, pkgs...)...)
	out = append(out, eres("baseline-upstream-tests-without-couchdb", err == nil, "upstream tests fail even without vet and CouchDB: %.300s", o))
	return out
}
