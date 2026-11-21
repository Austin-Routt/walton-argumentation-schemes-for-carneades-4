package caes

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/terms"
)

// A graph of a few thousand statements under the negation issue scheme, which
// the library always carries, builds its issues: issue construction must not
// end as limit reached on a graph the unbounded scan completed.
func TestIssueConstructionScales(t *testing.T) {
	ag := negationGraph(3000)
	if err := ag.Infer(); err != nil {
		t.Fatalf("3,000 statements under the negation issue scheme: %v", err)
	}
	if got := len(ag.Issues); got != 750 {
		t.Errorf("want 750 issues (one for each statement assumed with a contrary argument), got %d", got)
	}
}

// Looking positions up by key gives the same issues as the scan of all
// statements, for the negation scheme and for schemes whose later patterns
// are not ground (which still scan).
func TestIssueLookupMatchesScan(t *testing.T) {
	schemes := map[string]IssueScheme{
		"negation": {"S", "¬S"},
		"pair":     {"accepts(X)", "rejects(X)", "abstains(X)"},
		"open":     {"leads(X)", "trails(Y)"},
	}
	build := func(lookup bool) string {
		saved := issueLookup
		issueLookup = lookup
		defer func() { issueLookup = saved }()
		ag := NewArgGraph()
		for name, s := range schemes {
			s := s
			ag.Theory.IssueSchemes[name] = &s
		}
		ids := []string{"p(a)", "¬p(a)", "p(b)", "q(c)", "¬q(d)",
			"accepts(m)", "rejects(m)", "abstains(n)", "rejects(n)", "accepts(o)",
			"leads(x)", "trails(y)", "r([a,b])", "¬r([a, b])", "s(f(g(h)))", "¬s(f(g(h)))",
			// statements with variables take the scan; one bound to itself
			// once looped in substitution
			"u(Y,Y)", "¬u(Y,Y)", "u(Z,W)", "accepts(V)", "rejects(V)"}
		// and a larger, irregular graph: 600 statements of four shapes
		for i := 0; i < 600; i++ {
			k := (i * 7919) % 97
			switch i % 4 {
			case 0:
				ids = append(ids, fmt.Sprintf("p(k%d)", k))
			case 1:
				ids = append(ids, fmt.Sprintf("¬p(k%d)", k))
			case 2:
				ids = append(ids, fmt.Sprintf("accepts(k%d)", k%31))
			default:
				ids = append(ids, fmt.Sprintf("rejects(k%d)", k%37))
			}
		}
		for _, id := range ids {
			id = terms.Normalize(id)
			ag.Statements[id] = &Statement{Id: id, Text: id}
		}
		if err := ag.materializeIssues(context.Background()); err != nil {
			t.Fatalf("lookup=%v: %v", lookup, err)
		}
		return describeIssues(ag)
	}
	scan, lookup := build(false), build(true)
	if scan != lookup {
		t.Errorf("issues differ\nscan:\n%s\nlookup:\n%s", scan, lookup)
	}
	if !strings.Contains(lookup, "p(a) ¬p(a)") || !strings.Contains(lookup, "accepts(m) rejects(m)") || !strings.Contains(lookup, "r([a, b]) ¬r([a, b])") {
		t.Errorf("expected issues missing:\n%s", lookup)
	}
}

// negationGraph: n statements p(i), every fourth with an argument for its
// negation, so that 3,000 statements give 750 issues.
func negationGraph(n int) *ArgGraph {
	ag := NewArgGraph()
	neg := IssueScheme{"S", "¬S"}
	ag.Theory.IssueSchemes["negation"] = &neg
	for i := 0; i < n; i++ {
		id := fmt.Sprintf("p(c%d)", i)
		ag.Statements[id] = &Statement{Id: id, Text: id}
		if i%4 == 0 {
			nid := "¬" + id
			ag.Statements[nid] = &Statement{Id: nid, Text: nid}
		}
	}
	return ag
}

func describeIssues(ag *ArgGraph) string {
	lines := []string{}
	for _, i := range ag.Issues {
		ps := []string{}
		for _, p := range i.Positions {
			ps = append(ps, p.Id)
		}
		sort.Strings(ps)
		lines = append(lines, strings.Join(ps, " "))
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}
