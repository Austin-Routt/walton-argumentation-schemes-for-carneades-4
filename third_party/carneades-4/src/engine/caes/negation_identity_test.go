package caes

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/terms"
)

// A scheme that concludes ¬S for a negative S concludes S. The rule store
// reads ¬¬S as S and reasons from it (docs/contracts.md, contract 4), so the
// graph has to name the same statement. Before, the store built the next
// argument from q(a) while the graph held ¬¬q(a) in and left q(a), and all
// that followed from it, out.
func TestDerivedDoubleNegationIsTheStatement(t *testing.T) {
	ag := NewArgGraph()
	neg := IssueScheme{"S", "¬S"}
	ag.Theory.IssueSchemes["negation"] = &neg
	ag.Theory.ArgSchemes = []*Scheme{
		{Id: "deny", Variables: []string{"S"}, Premises: []string{"denies(S)"}, Conclusions: []string{"¬S"}, Weight: LinkedWeighingFunction},
		{Id: "next", Variables: []string{"X"}, Premises: []string{"q(X)"}, Conclusions: []string{"p(X)"}, Weight: LinkedWeighingFunction},
		{Id: "open", Premises: []string{"safe"}, Conclusions: []string{"opened"}, Weight: LinkedWeighingFunction},
	}
	ag.Theory.InitSchemeIndex()
	ag.Statements["q(a)"] = &Statement{Id: "q(a)", Text: "q(a)"}
	ag.Assumptions = []string{"denies(¬q(a))", "denies(¬safe)", "denies(q(b))"}
	if err := ag.Infer(); err != nil {
		t.Fatal(err)
	}
	for k := range ag.Statements {
		if strings.HasPrefix(k, "¬¬") {
			t.Errorf("the graph holds the statement %s", k)
		}
	}
	l := ag.GroundedLabelling()
	for _, id := range []string{"q(a)", "p(a)", "safe", "opened", "¬q(b)"} {
		s := ag.Statements[id]
		if s == nil {
			t.Errorf("%s is not a statement of the graph", id)
		} else if l[s] != In {
			t.Errorf("%s is labelled %v, want in", id, l[s])
		}
	}
	if len(ag.Issues) != 0 {
		t.Errorf("no statement here has its negation in the graph, yet %d issue(s) were built", len(ag.Issues))
	}
}

// The folded conclusion meets a stipulated negation in the issue of the
// statement, like any other argument for it.
func TestDerivedDoubleNegationJoinsTheIssueOfTheStatement(t *testing.T) {
	ag := NewArgGraph()
	neg := IssueScheme{"S", "¬S"}
	ag.Theory.IssueSchemes["negation"] = &neg
	ag.Theory.ArgSchemes = []*Scheme{
		{Id: "deny", Variables: []string{"S"}, Premises: []string{"denies(S)"}, Conclusions: []string{"¬S"}, Weight: LinkedWeighingFunction},
		{Id: "say", Variables: []string{"S"}, Premises: []string{"says(S)"}, Conclusions: []string{"S"}, Weight: LinkedWeighingFunction},
	}
	ag.Theory.InitSchemeIndex()
	ag.Assumptions = []string{"denies(¬q(a))", "says(¬q(a))"}
	if err := ag.Infer(); err != nil {
		t.Fatal(err)
	}
	if len(ag.Issues) != 1 {
		t.Fatalf("want the one issue of q(a) and ¬q(a), got %d issues", len(ag.Issues))
	}
	for _, issue := range ag.Issues {
		ids := []string{}
		for _, p := range issue.Positions {
			ids = append(ids, p.Id)
		}
		sort.Strings(ids)
		if got := strings.Join(ids, " "); got != "q(a) ¬q(a)" {
			t.Errorf("positions %q, want \"q(a) ¬q(a)\"", got)
		}
	}
}

// An authored statement or assumption written with a double negation is
// refused, and names the statement to write instead. The graph is unchanged.
func TestAuthoredDoubleNegationIsRefused(t *testing.T) {
	for name, build := range map[string]func(*ArgGraph){
		"statement":  func(ag *ArgGraph) { ag.Statements["¬¬q(a)"] = &Statement{Id: "¬¬q(a)", Text: "x"} },
		"assumption": func(ag *ArgGraph) { ag.Assumptions = []string{"¬¬q(a)"} },
		"atom":       func(ag *ArgGraph) { ag.Assumptions = []string{"¬¬¬rain"} },
	} {
		ag := NewArgGraph()
		neg := IssueScheme{"S", "¬S"}
		ag.Theory.IssueSchemes["negation"] = &neg
		ag.Statements["q(a)"] = &Statement{Id: "q(a)", Text: "q(a)"}
		build(ag)
		before := len(ag.Statements)
		for i, err := range []error{ag.CheckContext(), ag.Infer()} {
			if !errors.Is(err, ErrInvalidInput) || !strings.Contains(fmt.Sprint(err), "double negation") {
				t.Errorf("%s, call %d: want invalid input naming the double negation, got %v", name, i, err)
			}
		}
		if len(ag.Statements) != before || len(ag.Issues) != 0 {
			t.Errorf("%s: the refused run changed the graph", name)
		}
	}
	// a single negation, and a negation inside a term, are ordinary
	ag := NewArgGraph()
	ag.Assumptions = []string{"¬q(a)", "denies(¬¬q(a))"}
	if err := ag.Infer(); err != nil {
		t.Errorf("a negation inside a term is data, not a statement: %v", err)
	}
}

// The same graph gives the same issues on every run. An issue scheme whose
// later pattern has a variable of its own takes the first statement that
// matches, and "first" used to be the iteration order of a map: the positions
// of the issue differed from run to run.
func TestIssueConstructionIsRepeatable(t *testing.T) {
	describe := func() string {
		ag := NewArgGraph()
		open := IssueScheme{"leads(X)", "trails(Y)"}
		neg := IssueScheme{"S", "¬S"}
		ag.Theory.IssueSchemes["open"] = &open
		ag.Theory.IssueSchemes["negation"] = &neg
		for _, id := range []string{"leads(a)", "leads(b)", "leads(c)", "trails(u)", "trails(v)", "trails(w)", "trails(x)",
			"p(a)", "¬p(a)", "p(b)", "¬p(c)", "¬leads(a)", "¬trails(w)"} {
			ag.Statements[id] = &Statement{Id: id, Text: id}
		}
		if err := ag.Infer(); err != nil {
			return "error: " + err.Error()
		}
		return describeIssues(ag)
	}
	first := describe()
	for i := 0; i < 200; i++ {
		if got := describe(); got != first {
			t.Fatalf("run %d differs:\n%s\nfirst run:\n%s", i+2, got, first)
		}
	}
}

// A position is found by the normal form of its key: a graph built through
// the API, whose keys need not be in normal form, gets the issues the scan of
// all statements gave it.
func TestIssueLookupNormalizesKeys(t *testing.T) {
	build := func(lookup bool) string {
		saved := issueLookup
		issueLookup = lookup
		defer func() { issueLookup = saved }()
		ag := NewArgGraph()
		neg := IssueScheme{"S", "¬S"}
		pair := IssueScheme{"accepts(X)", "rejects(X)"}
		ag.Theory.IssueSchemes["negation"] = &neg
		ag.Theory.IssueSchemes["pair"] = &pair
		// no key of a position is in normal form: "r([a, b])", "q(c)", ...
		for _, id := range []string{"r([a,b])", "¬r([a,b])", "q( c )", "¬q( c )", "accepts( m )", "rejects( m )", "s(d)"} {
			ag.Statements[id] = &Statement{Id: id, Text: id}
		}
		if err := ag.Infer(); err != nil {
			return "error: " + err.Error()
		}
		return describeIssues(ag)
	}
	scan, lookup := build(false), build(true)
	if scan != lookup {
		t.Errorf("lookup:\n%s\nscan:\n%s", lookup, scan)
	}
	if want := "accepts( m ) rejects( m )\nq( c ) ¬q( c )\nr([a,b]) ¬r([a,b])"; lookup != want {
		t.Errorf("want the three issues\n%s\ngot:\n%s", want, lookup)
	}
}

func TestFoldNegation(t *testing.T) {
	for in, want := range map[string]string{
		"q(a)": "q(a)", "¬q(a)": "¬q(a)", "¬¬q(a)": "q(a)", "¬¬¬q(a)": "¬q(a)", "¬¬¬¬q(a)": "q(a)",
		"¬¬rain": "rain", "denies(¬¬q(a))": "denies(¬¬q(a))", "¬denies(¬¬q(a))": "¬denies(¬¬q(a))", "¬¬S": "S",
	} {
		term, ok := terms.ReadString(in)
		if !ok {
			t.Fatalf("cannot read %s", in)
		}
		if got := terms.FoldNegation(term).String(); got != terms.Normalize(want) {
			t.Errorf("FoldNegation(%s) = %s, want %s", in, got, want)
		}
	}
}
