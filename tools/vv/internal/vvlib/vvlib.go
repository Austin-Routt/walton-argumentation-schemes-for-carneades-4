// Package vvlib holds the checks of the verification and validation suite for
// the Walton scheme library. Every check drives the unmodified Carneades 4
// engine (and, through it, GoCHR) and reports a Result; the tests and the
// schemecheck report tool only decide what to do with the results.
package vvlib

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/carneades/carneades-4/src/engine/caes"
	cyaml "github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
	"github.com/carneades/carneades-4/src/engine/terms"
	"github.com/carneades/carneades-4/src/engine/validation"
	chr "github.com/hfried/GoCHR/src/engine/CHR"
)

// BudgetMarker is the line the engine as shipped printed on standard output
// when it stopped after the maximum number of rule applications, its only
// sign of it. Since E-BUDGET (PLAN.md 1.2) Infer returns an error wrapping
// caes.ErrLimitReached instead; LimitReached accepts both.
const BudgetMarker = "inference engine interrupted"

// LimitReached tells whether a run stopped at a limit of the engine.
func LimitReached(err error, stdout string) bool {
	return errors.Is(err, caes.ErrLimitReached) || strings.Contains(stdout, BudgetMarker)
}

// InferTimeout bounds one inference run. The engine always stops after
// caes.MAXRULEAPPS rule applications, so this is only a safety net.
const InferTimeout = 90 * time.Second

// Result is the outcome of one check.
type Result struct {
	File   string `json:"file"`
	Scheme string `json:"scheme"`
	Check  string `json:"check"`
	Pass   bool   `json:"pass"`
	Detail string `json:"detail,omitempty"`
}

func (r Result) Key() string { return r.File + "|" + r.Scheme + "|" + r.Check }

// ---------------------------------------------------------------- repository

// RepoRoot returns the scheme repository: $WALTON_REPO, or the directory two
// levels above this module (the suite lives in tools/vv).
func RepoRoot() string {
	if r := os.Getenv("WALTON_REPO"); r != "" {
		return r
	}
	wd, _ := os.Getwd()
	for d := wd; d != filepath.Dir(d); d = filepath.Dir(d) {
		if _, err := os.Stat(filepath.Join(d, "walton.yml")); err == nil {
			return d
		}
	}
	return wd
}

// SchemeFiles lists walton.yml and the category files, repository-relative.
func SchemeFiles(root string) []string {
	out := []string{"walton.yml"}
	m, _ := filepath.Glob(filepath.Join(root, "[A-Z]_*", "*.yml"))
	sort.Strings(m)
	for _, p := range m {
		rel, _ := filepath.Rel(root, p)
		out = append(out, filepath.ToSlash(rel))
	}
	return out
}

// ---------------------------------------------------------------- engine state

var basicNames = map[string]bool{"linked": true, "convergent": true, "cumulative": true, "factorized": true}

// ResetEngineGlobals removes what earlier yaml.Import calls left behind in the
// engine's package-level maps. yaml.Import aliases caes.BasicSchemes and
// caes.BasicWeighingFunctions and writes every imported scheme into them, so
// without this reset one file's schemes would be visible while checking the
// next file (see engine test "yaml-import-keeps-globals-clean").
func ResetEngineGlobals() {
	for k := range caes.BasicSchemes {
		if !basicNames[k] {
			delete(caes.BasicSchemes, k)
		}
	}
	for k := range caes.BasicWeighingFunctions {
		if !basicNames[k] {
			delete(caes.BasicWeighingFunctions, k)
		}
	}
}

// Capture redirects the process's stdout and stderr, which is where the engine
// reports most of its problems.
type Capture struct {
	oldOut, oldErr *os.File
	fo, fe         *os.File
	stopped        bool
}

func StartCapture() *Capture {
	c := &Capture{oldOut: os.Stdout, oldErr: os.Stderr}
	c.fo, _ = os.CreateTemp("", "vv-out-*")
	c.fe, _ = os.CreateTemp("", "vv-err-*")
	os.Stdout, os.Stderr = c.fo, c.fe
	return c
}

func (c *Capture) Stop() (stdout, stderr string) {
	if c.stopped {
		return "", ""
	}
	c.stopped = true
	os.Stdout, os.Stderr = c.oldOut, c.oldErr
	read := func(f *os.File) string {
		f.Seek(0, 0)
		b, _ := io.ReadAll(f)
		f.Close()
		os.Remove(f.Name())
		return string(b)
	}
	return read(c.fo), read(c.fe)
}

// ---------------------------------------------------------------- library

// Library is one scheme file as the engine's own YAML importer sees it.
type Library struct {
	Rel    string
	Path   string
	Text   string // file contents
	AG     *caes.ArgGraph
	Stdout string
	Stderr string
}

func LoadLibrary(root, rel string) (*Library, error) {
	p := filepath.Join(root, filepath.FromSlash(rel))
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	ResetEngineGlobals()
	lib := &Library{Rel: rel, Path: p, Text: string(b)}
	c := StartCapture()
	lib.AG, err = cyaml.Import(strings.NewReader(lib.Text))
	lib.Stdout, lib.Stderr = c.Stop()
	ResetEngineGlobals()
	return lib, err
}

func (l *Library) Schemes() []*caes.Scheme { return l.AG.Theory.ArgSchemes }

func (l *Library) Scheme(id string) *caes.Scheme {
	for _, s := range l.Schemes() {
		if s.Id == id {
			return s
		}
	}
	return nil
}

// ---------------------------------------------------------------- instances

// Instance is a ground instantiation of a scheme: every declared variable gets
// its own constant, and variables that stand for a whole statement (the S of
// "expert E asserts S, therefore S") get a small proposition.
type Instance struct {
	Scheme      *caes.Scheme
	Values      map[string]string
	Bound       map[string]bool // variable occurs in a premise or deletion
	Unbound     []string
	Premises    []string
	Assumptions []string
	Exceptions  []string
	Conclusions []string
	Synthetic   map[string]string // language templates for the synthetic propositions
	bindings    terms.Bindings
}

// namesOnly reports whether a term is built from atoms alone.
func namesOnly(t terms.Term) bool {
	switch t.Type() {
	case terms.AtomType:
		return true
	case terms.CompoundType:
		c, ok := t.(terms.Compound)
		if !ok {
			return false
		}
		for _, a := range c.Args {
			if !namesOnly(a) {
				return false
			}
		}
		return true
	}
	return false
}

func parse(s string) (terms.Term, bool) {
	c := StartCapture()
	t, ok := terms.ReadString(s)
	c.Stop()
	return t, ok
}

// Norm normalises a term the way the engine does (terms.Normalize) without
// letting parse errors reach the terminal.
func Norm(s string) string {
	t, ok := parse(s)
	if !ok {
		return s
	}
	return t.String()
}

func propositionalVars(s *caes.Scheme) map[string]bool {
	out := map[string]bool{}
	for _, l := range [][]string{s.Premises, s.Conclusions, s.Assumptions, s.Exceptions, s.Deletions} {
		for _, a := range l {
			t, ok := parse(a)
			if !ok {
				continue
			}
			if t.Type() == terms.VariableType {
				out[t.String()] = true
			}
			if c, isC := t.(terms.Compound); isC && c.Functor == "¬" && len(c.Args) == 1 && c.Args[0].Type() == terms.VariableType {
				out[c.Args[0].String()] = true
			}
		}
	}
	return out
}

func varsOf(list ...[]string) map[string]bool {
	out := map[string]bool{}
	for _, l := range list {
		for _, a := range l {
			if t, ok := parse(a); ok {
				for _, v := range t.OccurVars() {
					out[v.Name] = true
				}
			}
		}
	}
	return out
}

func NewInstance(s *caes.Scheme) *Instance {
	in := &Instance{Scheme: s, Values: map[string]string{}, Synthetic: map[string]string{}}
	props := propositionalVars(s)
	in.Bound = varsOf(s.Premises, s.Deletions)
	for _, v := range s.Variables {
		val := "k_" + strings.ToLower(v)
		if props[v] {
			f := "claim_" + strings.ToLower(v)
			val = f + "(x)"
			in.Synthetic[f+"/1"] = "Claim " + strings.ToLower(v) + " holds for %s."
		}
		in.Values[v] = val
		t, _ := parse(val)
		in.bindings = terms.AddBinding(terms.NewVariable(v), t, in.bindings)
		if !in.Bound[v] {
			in.Unbound = append(in.Unbound, v)
		}
	}
	in.Premises = in.ground(s.Premises)
	in.Assumptions = in.ground(s.Assumptions)
	in.Exceptions = in.ground(s.Exceptions)
	in.Conclusions = in.ground(s.Conclusions)
	return in
}

func (in *Instance) ground(list []string) []string {
	out := []string{}
	for _, a := range list {
		t, ok := parse(a)
		if !ok {
			out = append(out, a)
			continue
		}
		out = append(out, terms.Substitute(t, in.bindings).String())
	}
	return out
}

// ExpectedParams are the values the engine should report for the scheme's
// variables, in declaration order; "" marks a variable no premise binds.
func (in *Instance) ExpectedParams() []string {
	out := []string{}
	for _, v := range in.Scheme.Variables {
		if in.Bound[v] {
			out = append(out, Norm(in.Values[v]))
		} else {
			out = append(out, "")
		}
	}
	return out
}

// ---------------------------------------------------------------- inference

// Run is one call of ArgGraph.Infer followed by grounded labelling.
type Run struct {
	AG       *caes.ArgGraph
	L        caes.Labelling
	Err      error
	TimedOut bool
	Stdout   string
	Stderr   string
	Elapsed  time.Duration
}

func (r *Run) BudgetExhausted() bool { return LimitReached(r.Err, r.Stdout) }

// Infer builds an argument graph from a language, a set of schemes and ground
// assumptions, and runs the engine's argument construction and evaluation.
func Infer(lang caes.Language, extra map[string]string, schemes []*caes.Scheme, assumptions []string) *Run {
	ResetEngineGlobals()
	ag := caes.NewArgGraph()
	th := caes.NewTheory()
	for k, v := range lang {
		th.Language[k] = v
	}
	for k, v := range extra {
		th.Language[k] = v
	}
	th.ArgSchemes = schemes
	th.InitSchemeIndex()
	ag.Theory = th
	ag.Assumptions = append([]string{}, assumptions...)

	r := &Run{AG: ag}
	done := make(chan error, 1)
	c := StartCapture()
	start := time.Now()
	go func() { done <- ag.Infer() }()
	select {
	case r.Err = <-done:
	case <-time.After(InferTimeout):
		r.TimedOut = true
		Taint()
	}
	r.Elapsed = time.Since(start)
	if r.TimedOut {
		r.Stdout, r.Stderr = c.Stop()
		return r
	}
	r.L = ag.GroundedLabelling()
	r.Stdout, r.Stderr = c.Stop()
	return r
}

// ArgsOf returns the arguments built from the scheme with this id, sorted by id.
func (r *Run) ArgsOf(id string) []*caes.Argument {
	out := []*caes.Argument{}
	for _, a := range r.AG.Arguments {
		if a.Scheme != nil && a.Scheme.Id == id {
			out = append(out, a)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Id < out[j].Id })
	return out
}

func (r *Run) Statement(id string) *caes.Statement {
	if s, ok := r.AG.Statements[Norm(id)]; ok {
		return s
	}
	return nil
}

func (r *Run) Label(id string) string {
	s := r.Statement(id)
	if s == nil {
		return "missing"
	}
	return r.L[s].String()
}

func (r *Run) NonGround() []string {
	out := []string{}
	for k := range r.AG.Statements {
		t, ok := parse(k)
		if !ok || !terms.Ground(t, nil) {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}

// ---------------------------------------------------------------- checks

var isVarName = regexp.MustCompile(`^[A-Z][A-Za-z0-9_]*$`)

func pass(f, s, c string) Result { return Result{File: f, Scheme: s, Check: c, Pass: true} }
func fail(f, s, c, d string, a ...any) Result {
	return Result{File: f, Scheme: s, Check: c, Detail: fmt.Sprintf(d, a...)}
}
func res(f, s, c string, problems []string) Result {
	if len(problems) == 0 {
		return pass(f, s, c)
	}
	return fail(f, s, c, "%s", strings.Join(problems, "; "))
}

// ValidatorProblems runs the engine's own validator on a library and groups
// what it reports: scheme problems by scheme id, everything else under "".
func ValidatorProblems(lib *Library) map[string][]string {
	ResetEngineGlobals()
	c := StartCapture()
	probs := validation.Validate(lib.AG)
	c.Stop()
	out := map[string][]string{}
	for _, p := range probs {
		msg := strings.TrimSpace(p.Description)
		if p.Expression != "" {
			msg += ": " + p.Expression
		}
		if p.Category == validation.SCHEME {
			out[p.Id] = append(out[p.Id], msg)
		} else {
			out[""] = append(out[""], p.Category.String()+": "+p.Id+": "+msg)
		}
	}
	return out
}

// RuleError translates one scheme into a GoCHR rule exactly as
// caes.TheoryToRuleStore does and returns the error GoCHR gives (the engine
// as shipped discarded it; caes.RuleStoreOf now reports it), together with
// anything GoCHR printed.
func RuleError(s *caes.Scheme) (error, string) {
	rs := chr.MakeRuleStore()
	prem := s.Premises
	if len(prem) == 0 && len(s.Deletions) == 0 {
		prem = []string{"go"}
	}
	body := append(append([]string{}, s.Conclusions...), s.Assumptions...)
	body = append(body, fmt.Sprintf("argument(%s,[%s])", s.Id, strings.Join(s.Variables, ",")))
	c := StartCapture()
	err := rs.AddRule(s.Id, prem, s.Deletions, s.Guards, body)
	o, e := c.Stop()
	return err, strings.TrimSpace(o + e)
}

// StaticChecks are properties of a scheme that can be read off its text and
// that decide whether the engine can use it at all.
func StaticChecks(file string, s *caes.Scheme) []Result {
	out := []Result{}
	id := s.Id

	// The scheme id is written into a CHR term, argument(id,[...]), and read
	// back from the store. It has to come back as it was written: a name, or
	// names joined by operators that nothing computes (opposition-normative is
	// read as a subtraction of two names and printed back unchanged). An id with
	// a number or a variable as an operand would be computed or bound on the way.
	if t, ok := parse(id); ok && t.String() == id && namesOnly(t) {
		out = append(out, pass(file, id, "static/id-round-trips"))
	} else {
		out = append(out, fail(file, id, "static/id-round-trips", "the id does not come back from the term reader as it was written, or has a number or a variable in it that the rule engine would compute or bind"))
	}

	// every variable used after the premises must be bound by a premise
	bound := varsOf(s.Premises, s.Deletions)
	unbound := []string{}
	for v := range varsOf(s.Conclusions, s.Assumptions, s.Exceptions, s.Guards) {
		if !bound[v] {
			unbound = append(unbound, v)
		}
	}
	sort.Strings(unbound)
	if len(unbound) == 0 {
		out = append(out, pass(file, id, "static/range-restricted"))
	} else {
		out = append(out, fail(file, id, "static/range-restricted", "no premise binds %v", unbound))
	}

	// GoCHR renames an unbound variable V to V1, V2, ...; a declared variable of
	// that name then captures it when the engine instantiates the scheme
	capture := []string{}
	for _, v := range unbound {
		re := regexp.MustCompile("^" + regexp.QuoteMeta(v) + `[0-9]+$`)
		for _, w := range s.Variables {
			if re.MatchString(w) {
				capture = append(capture, v+" can be captured by "+w)
			}
		}
	}
	out = append(out, res(file, id, "static/no-variable-capture", capture))

	// static/no-negated-premise is removed with PLAN.md 2.4: negated premises
	// match since then (contract 4), and the book's schemes have them.

	// functors the rule engine treats as built-ins when they appear in a rule body
	reserved := []string{}
	for _, b := range append(append([]string{}, s.Conclusions...), s.Assumptions...) {
		if t, ok := parse(b); ok {
			if c, isC := t.(terms.Compound); isC && len(c.Args) == 2 && (c.Functor == "is" || c.Functor == "=" || c.Functor == ":=") {
				reserved = append(reserved, Norm(b)+" (GoCHR reads "+c.Functor+"/2 in a rule body as an assignment)")
			}
		}
	}
	out = append(out, res(file, id, "static/no-reserved-functor", reserved))

	// a body atom equal to a premise re-triggers the rule for ever
	self := []string{}
	prem := map[string]bool{}
	for _, p := range s.Premises {
		prem[Norm(p)] = true
	}
	for _, b := range append(append([]string{}, s.Conclusions...), s.Assumptions...) {
		if prem[Norm(b)] {
			self = append(self, Norm(b)+" is both a premise and derived by the scheme")
		}
	}
	out = append(out, res(file, id, "static/not-self-feeding", self))
	return out
}

var badText = regexp.MustCompile(`%!|%s|%v|%\[`)

// DynamicChecks instantiates the scheme, lets the engine construct and
// evaluate the argument, and compares the result with what the scheme says.
// theory is the set of schemes loaded together with s; mode names the set
// ("iso" for the scheme alone, "full" for the whole file).
func DynamicChecks(file, mode string, lang caes.Language, theory []*caes.Scheme, s *caes.Scheme) ([]Result, *Run) {
	id := s.Id
	in := NewInstance(s)
	run := Infer(lang, in.Synthetic, theory, in.Premises)
	out := []Result{}
	chk := func(name string) string { return mode + "/" + name }

	if run.TimedOut {
		return []Result{fail(file, id, chk("infer-ok"), "inference did not return within %v", InferTimeout)}, run
	}
	if run.Err != nil && !run.BudgetExhausted() {
		out = append(out, fail(file, id, chk("infer-ok"), "Infer: %v", run.Err))
	} else {
		out = append(out, pass(file, id, chk("infer-ok")))
	}
	if run.BudgetExhausted() {
		out = append(out, fail(file, id, chk("terminates"), "the rule engine stopped at its limit of %d rule applications after %v", caes.MAXRULEAPPS, run.Elapsed.Round(time.Millisecond)))
	} else {
		out = append(out, pass(file, id, chk("terminates")))
	}

	// one argument per conclusion, with the parameters the premises determine
	args := run.ArgsOf(id)
	want := in.ExpectedParams()
	probs := []string{}
	if len(args) != len(s.Conclusions) {
		probs = append(probs, fmt.Sprintf("%d argument(s) constructed, %d expected (one per conclusion)", len(args), len(s.Conclusions)))
	}
	for _, a := range args {
		if len(a.Parameters) != len(want) {
			probs = append(probs, fmt.Sprintf("%s has %d parameters, the scheme declares %d", a.Id, len(a.Parameters), len(want)))
			continue
		}
		for i, w := range want {
			got := a.Parameters[i]
			if w == "" {
				continue // unbound: reported by the ground check
			}
			if Norm(got) != w {
				probs = append(probs, fmt.Sprintf("%s: parameter %s is %q, expected %q", a.Id, s.Variables[i], got, w))
			}
		}
	}
	out = append(out, res(file, id, chk("fires"), probs))

	// The checks that add an exception or leave a premise out show something
	// only if the plain run came to its end and built the scheme's argument.
	// Without that control a partial graph would count as a successful defeat.
	control := ""
	switch {
	case run.BudgetExhausted():
		control = "with every premise asserted the run stops at the rule limit"
	case run.Err != nil:
		control = "with every premise asserted Infer returns an error"
	case len(args) == 0:
		control = "with every premise asserted the scheme constructs no argument"
	}

	// A second Infer on the finished graph adds nothing and changes no label
	// (contract 2; E-IDEMP). Ported from the third review's per-scheme check
	// of a second inference (docs/parity.md).
	out = append(out, res(file, id, chk("infer-twice"), secondInferProblems(run, control)))

	if ng := run.NonGround(); len(ng) == 0 {
		out = append(out, pass(file, id, chk("ground")))
	} else {
		out = append(out, fail(file, id, chk("ground"), "statements with variables: %s", strings.Join(ng, " ")))
	}

	probs = nil
	for _, c := range in.Conclusions {
		if l := run.Label(c); l != "in" {
			probs = append(probs, fmt.Sprintf("%s is %s", c, l))
		}
	}
	out = append(out, res(file, id, chk("conclusions-in"), probs))

	if mode != "iso" {
		// With every scheme of the file loaded, an exception must still undercut
		// this scheme's argument, and a missing premise must still prevent it.
		if len(in.Exceptions) > 0 {
			out = append(out, res(file, id, chk("exceptions-undercut"), exceptionProblems(lang, theory, s, in, control)))
		}
		out = append(out, res(file, id, chk("premises-needed"), premiseProblems(lang, theory, s, in, control)))
		return out, run
	}

	// premises of the constructed argument: the scheme's premises, then its assumptions
	probs = nil
	wantPrem := []string{}
	for _, p := range append(append([]string{}, in.Premises...), in.Assumptions...) {
		if !strings.HasPrefix(p, "argument(") {
			wantPrem = append(wantPrem, Norm(p))
		}
	}
	for _, a := range args {
		got := []string{}
		for _, p := range a.Premises {
			got = append(got, p.Stmt.Id)
		}
		if strings.Join(got, " | ") != strings.Join(wantPrem, " | ") {
			probs = append(probs, fmt.Sprintf("%s has premises [%s], expected [%s]", a.Id, strings.Join(got, ", "), strings.Join(wantPrem, ", ")))
		}
	}
	for _, a := range in.Assumptions {
		if l := run.Label(a); l != "in" {
			probs = append(probs, fmt.Sprintf("assumption %s is %s", a, l))
		}
	}
	out = append(out, res(file, id, chk("premises"), probs))

	// what a reader of the exported graph sees
	probs = nil
	neg := []string{}
	for _, st := range append(append(append(append([]string{}, in.Premises...), in.Assumptions...), in.Conclusions...), in.Exceptions...) {
		stmt := run.Statement(st)
		if stmt == nil {
			continue // absence is reported by the other checks
		}
		t, _ := parse(stmt.Id)
		functor, _ := terms.Functor(t)
		if _, synthetic := in.Synthetic[functor+"/1"]; synthetic {
			continue
		}
		switch {
		case functor == "¬":
			if stmt.Text == stmt.Id {
				neg = append(neg, stmt.Id)
			}
		case strings.TrimSpace(stmt.Text) == "":
			probs = append(probs, stmt.Id+": empty text")
		case badText.MatchString(stmt.Text):
			probs = append(probs, stmt.Id+": "+stmt.Text)
		case stmt.Text == stmt.Id:
			probs = append(probs, stmt.Id+": no template applied")
		}
	}
	out = append(out, res(file, id, chk("texts"), probs))
	if len(neg) > 0 {
		out = append(out, fail(file, id, chk("negation-texts"), "printed as raw terms: %s", strings.Join(neg, " ")))
	} else {
		out = append(out, pass(file, id, chk("negation-texts")))
	}

	// critical questions modelled as exceptions: asserting one must defeat the argument
	if len(in.Exceptions) > 0 {
		out = append(out, res(file, id, chk("exceptions-undercut"), exceptionProblems(lang, theory, s, in, control)))
	}
	// every premise is needed: without it the scheme must not construct an argument
	out = append(out, res(file, id, chk("premises-needed"), premiseProblems(lang, theory, s, in, control)))
	return out, run
}

// secondInferProblems runs Infer again on the graph of a finished run.
func secondInferProblems(run *Run, control string) []string {
	if control != "" {
		return []string{"not judged: " + control}
	}
	n1 := len(run.AG.Arguments)
	before := map[string]caes.Label{}
	for st, l := range run.L {
		before[st.Id] = l
	}
	var panicked any
	var err error
	c := StartCapture()
	func() {
		defer func() { panicked = recover() }()
		err = run.AG.Infer()
	}()
	c.Stop()
	if panicked != nil {
		return []string{fmt.Sprintf("the second Infer panics: %v", panicked)}
	}
	if err != nil {
		return []string{fmt.Sprintf("the second Infer fails: %v", err)}
	}
	probs := []string{}
	if n2 := len(run.AG.Arguments); n2 != n1 {
		probs = append(probs, fmt.Sprintf("the second Infer raises the number of arguments from %d to %d", n1, n2))
	}
	for st, l := range run.AG.GroundedLabelling() {
		if b, ok := before[st.Id]; ok && b != l {
			probs = append(probs, fmt.Sprintf("%s was %s and is %s after the second Infer", st.Id, b, l))
		}
	}
	sort.Strings(probs)
	return probs
}

// exceptionProblems asserts each exception of the scheme in turn, next to the
// premises, and reports the arguments of the scheme that are not undercut.
//
// The oracle is the argument, not the conclusion: with other schemes loaded, a
// conclusion can keep the support of another argument after this one has been
// undercut, and that is how convergent arguments are meant to behave. A run in
// which the scheme constructs no argument at all proves nothing about its
// exceptions and is reported as a failure too.
func exceptionProblems(lang caes.Language, theory []*caes.Scheme, s *caes.Scheme, in *Instance, control string) []string {
	if control != "" {
		return []string{"not judged: " + control + ", so an exception has nothing to undercut"}
	}
	probs := []string{}
	for i, e := range in.Exceptions {
		r2 := Infer(lang, in.Synthetic, theory, append(append([]string{}, in.Premises...), e))
		if r2.TimedOut || r2.Err != nil {
			probs = append(probs, fmt.Sprintf("exception %d (%s): inference failed", i+1, s.Exceptions[i]))
			continue
		}
		if r2.BudgetExhausted() {
			probs = append(probs, fmt.Sprintf("exception %d (%s): the run stopped at the rule limit, so its graph is partial", i+1, s.Exceptions[i]))
			continue
		}
		args := r2.ArgsOf(s.Id)
		if len(args) == 0 {
			probs = append(probs, fmt.Sprintf("with %s accepted the scheme constructs no argument, so the exception cannot be shown to work", e))
			continue
		}
		for _, a := range args {
			if a.Undercut(r2.L) != caes.In {
				c := ""
				if a.Conclusion != nil {
					c = a.Conclusion.Id
				}
				probs = append(probs, fmt.Sprintf("with %s accepted, the argument for %s is not undercut", e, c))
			}
		}
	}
	return probs
}

// premiseProblems leaves out each premise of the scheme in turn: no argument
// of the scheme may then be constructed. Both this check and the one above
// need the plain run as a control (it finished and built the argument) and a
// finished run of their own; otherwise they report that nothing was judged. As for exceptions, the oracle is the
// scheme's own argument; whether the conclusion keeps other support is a
// property of the whole theory, not of this scheme.
func premiseProblems(lang caes.Language, theory []*caes.Scheme, s *caes.Scheme, in *Instance, control string) []string {
	if control != "" {
		return []string{"not judged: " + control + ", so leaving a premise out shows nothing"}
	}
	probs := []string{}
	for i := range in.Premises {
		rest := []string{}
		for j, p := range in.Premises {
			if j != i && Norm(p) != Norm(in.Premises[i]) {
				rest = append(rest, p)
			}
		}
		r2 := Infer(lang, in.Synthetic, theory, rest)
		if r2.TimedOut || r2.Err != nil {
			probs = append(probs, fmt.Sprintf("without premise %d (%s): inference failed", i+1, s.Premises[i]))
			continue
		}
		if r2.BudgetExhausted() {
			probs = append(probs, fmt.Sprintf("without premise %d (%s): the run stopped at the rule limit, so its graph is partial", i+1, s.Premises[i]))
			continue
		}
		if n := len(r2.ArgsOf(s.Id)); n > 0 {
			probs = append(probs, fmt.Sprintf("without premise %d (%s) the scheme still constructs %d argument(s)", i+1, s.Premises[i], n))
		}
	}
	return probs
}

// CheckLibrary runs every check on every scheme of a file.
func CheckLibrary(lib *Library) []Result {
	out := []Result{}
	vp := ValidatorProblems(lib)
	out = append(out, res(lib.Rel, "(file)", "engine-validator", vp[""]))
	lang := lib.AG.Theory.Language
	all := lib.Schemes()
	for _, s := range all {
		out = append(out, res(lib.Rel, s.Id, "engine-validator", dedup(vp[s.Id])))
		if err, printed := RuleError(s); err != nil {
			out = append(out, fail(lib.Rel, s.Id, "chr-translation", "AddRule: %v", err))
		} else if printed != "" {
			out = append(out, fail(lib.Rel, s.Id, "chr-translation", "GoCHR printed: %s", printed))
		} else {
			out = append(out, pass(lib.Rel, s.Id, "chr-translation"))
		}
		out = append(out, StaticChecks(lib.Rel, s)...)
		r, _ := DynamicChecks(lib.Rel, "iso", lang, []*caes.Scheme{s}, s)
		out = append(out, r...)
		r, _ = DynamicChecks(lib.Rel, "full", lang, all, s)
		out = append(out, r...)
	}
	return out
}

func dedup(l []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, s := range l {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

// ---------------------------------------------------------------- recursion

// Feeds reports which schemes can trigger which: an edge a -> b means some
// conclusion or assumption of a has the predicate of some premise of b. Cycles
// in this graph are where forward chaining can run away; the engine has no
// duplicate check, so a cycle that reproduces its own input never stops by
// itself. Statements that are a bare variable (the S a scheme concludes, the
// A that defeasible modus ponens takes as a premise) match anything and are
// left out here; Wildcards lists the schemes that have them.
func Feeds(schemes []*caes.Scheme) map[string][]string {
	sig := func(a string) (string, bool) {
		t, ok := parse(a)
		if !ok {
			return "", false
		}
		switch x := t.(type) {
		case terms.Compound:
			if x.Functor == "¬" {
				return "", false // negated atoms never reach a rule head
			}
			return fmt.Sprintf("%s/%d", x.Functor, len(x.Args)), true
		case terms.Atom:
			return string(x) + "/0", true
		}
		return "", false
	}
	out := map[string][]string{}
	for _, a := range schemes {
		produced := map[string]bool{}
		for _, b := range append(append([]string{}, a.Conclusions...), a.Assumptions...) {
			if s, ok := sig(b); ok {
				produced[s] = true
			}
		}
		for _, b := range schemes {
			for _, p := range b.Premises {
				if ps, ok := sig(p); ok && produced[ps] {
					out[a.Id] = append(out[a.Id], b.Id)
					break
				}
			}
		}
	}
	return out
}

// Wildcards lists the schemes that conclude a bare (possibly negated) variable
// and the schemes that take a bare variable as a premise.
func Wildcards(schemes []*caes.Scheme) (concludeAny, premiseAny []string) {
	bare := func(a string) bool {
		t, ok := parse(a)
		if !ok {
			return false
		}
		if t.Type() == terms.VariableType {
			return true
		}
		c, isC := t.(terms.Compound)
		return isC && c.Functor == "¬" && len(c.Args) == 1 && c.Args[0].Type() == terms.VariableType
	}
	for _, s := range schemes {
		for _, c := range s.Conclusions {
			if bare(c) {
				concludeAny = append(concludeAny, s.Id)
				break
			}
		}
		for _, p := range s.Premises {
			if bare(p) {
				premiseAny = append(premiseAny, s.Id)
				break
			}
		}
	}
	return
}

// Cycles returns the strongly connected components of the feeds graph that
// contain a cycle (more than one scheme, or a scheme that feeds itself).
func Cycles(feeds map[string][]string, ids []string) [][]string {
	index := 0
	idx := map[string]int{}
	low := map[string]int{}
	on := map[string]bool{}
	stack := []string{}
	out := [][]string{}
	var strong func(v string)
	strong = func(v string) {
		idx[v], low[v] = index, index
		index++
		stack = append(stack, v)
		on[v] = true
		for _, w := range feeds[v] {
			if _, seen := idx[w]; !seen {
				strong(w)
				if low[w] < low[v] {
					low[v] = low[w]
				}
			} else if on[w] && idx[w] < low[v] {
				low[v] = idx[w]
			}
		}
		if low[v] == idx[v] {
			comp := []string{}
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				on[w] = false
				comp = append(comp, w)
				if w == v {
					break
				}
			}
			self := false
			for _, w := range feeds[v] {
				if w == v {
					self = true
				}
			}
			if len(comp) > 1 || self {
				sort.Strings(comp)
				out = append(out, comp)
			}
		}
	}
	for _, v := range ids {
		if _, seen := idx[v]; !seen {
			strong(v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return strings.Join(out[i], ",") < strings.Join(out[j], ",") })
	return out
}
