package vvlib

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	cyaml "github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
)

func caseTestCLI(t *testing.T) *CLI {
	t.Helper()
	cli, err := FindCLI()
	if err != nil {
		t.Fatal(err)
	}
	if os.Getenv("CARNEADES_BIN") == "" {
		t.Cleanup(func() { os.RemoveAll(filepath.Dir(cli.Bin)) })
	}
	return cli
}

func TestCaseRoundTripPreservesExistingContext(t *testing.T) {
	cli := caseTestCLI(t)
	c := Case{
		Name: "context/append-not-overwrite",
		Graph: `language:
  p/0: P
  q/0: Q
  r/0: R
  d/0: D
  e/0: E
  blocked_d/0: Blocked D
  blocked_e/0: Blocked E
argument_schemes:
  - id: joint
    premises: [p, q]
    conclusions: [r]
  - id: default_d
    premises: [p]
    assumptions: [d]
    conclusions: [blocked_d]
  - id: default_e
    premises: [q]
    assumptions: [e]
    conclusions: [blocked_e]
assumptions: [p]
questioned: [d]
`,
		Assumptions: []string{"q"},
		Questioned:  []string{"e"},
		Arguments:   []string{"joint()"},
		In:          []string{"p", "q", "r"},
		NotIn:       []string{"blocked_d", "blocked_e"},
	}
	if probs := RunCase(RepoRoot(), c); len(probs) != 0 {
		t.Fatalf("API context: %v", probs)
	}
	if probs := CaseRoundTrip(cli, RepoRoot(), c); len(probs) != 0 {
		t.Fatalf("CLI context: %v", probs)
	}
}

func TestCaseFactsPromoteExistingDefaults(t *testing.T) {
	cli := caseTestCLI(t)
	for _, fact := range []string{"p", "¬¬p"} {
		c := Case{
			Name: "context/stipulate-existing-default",
			Graph: `language:
  p/0: P
  q/0: Q
  r/0: R
  d/0: D
  blocked/0: Blocked
argument_schemes:
  - id: joint
    premises: [p, q]
    conclusions: [r]
  - id: still_default
    premises: [q]
    assumptions: [d]
    conclusions: [blocked]
defaults: [p, d]
questioned: [p, d]
`,
			Assumptions: []string{fact, "q"},
			Arguments:   []string{"joint()"},
			In:          []string{"p", "q", "r"},
			NotIn:       []string{"d", "blocked"},
		}
		if probs := RunCase(RepoRoot(), c); len(probs) != 0 {
			t.Fatalf("API stipulation %q: %v", fact, probs)
		}
		if probs := CaseRoundTrip(cli, RepoRoot(), c); len(probs) != 0 {
			t.Fatalf("CLI stipulation %q: %v", fact, probs)
		}
		if probs := variantProblems([]byte(c.Graph), c); len(probs) != 0 {
			t.Fatalf("case variants %q: %v", fact, probs)
		}
	}
}

const selectionCaseLibrary = `language:
  p/1: "%s has P"
  q/1: "%s has Q"
  r/1: "%s has R"
argument_schemes:
  - id: selected
    variables: [X]
    premises: [p(X), q(X)]
    conclusions: [r(X)]
  - id: derive
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
`

func TestCaseSchemeSelection(t *testing.T) {
	cli := caseTestCLI(t)
	root := t.TempDir()
	path := filepath.Join(root, "live.yml")
	if err := os.WriteFile(path, []byte(selectionCaseLibrary), 0600); err != nil {
		t.Fatal(err)
	}
	base := Case{Name: "selection/control", Library: "live.yml", Assumptions: []string{"p(a)"}}
	for _, run := range []struct {
		name string
		fn   func(Case) []string
	}{
		{"api", func(c Case) []string { return RunCase(root, c) }},
		{"cli", func(c Case) []string { return CaseRoundTrip(cli, root, c) }},
	} {
		t.Run(run.name, func(t *testing.T) {
			full := base
			full.Arguments = []string{"derive(a)", "selected(a)"}
			full.In = []string{"r(a)"}
			if probs := run.fn(full); len(probs) != 0 {
				t.Fatalf("full-library composition: %v", probs)
			}
			selected := base
			selected.Schemes = []string{"selected"}
			selected.NoArguments = []string{"derive", "selected"}
			selected.NotIn = []string{"q(a)", "r(a)"}
			if probs := run.fn(selected); len(probs) != 0 {
				t.Fatalf("selection must exclude independent support: %v", probs)
			}
			positive := selected
			positive.Assumptions = []string{"p(a)", "q(a)"}
			positive.Arguments = []string{"selected(a)"}
			positive.NoArguments = []string{"derive"}
			positive.In = []string{"r(a)"}
			positive.NotIn = nil
			positive.Variants = true
			if probs := run.fn(positive); len(probs) != 0 {
				t.Fatalf("selected positive control: %v", probs)
			}
			empty := selected
			empty.Schemes = []string{}
			if probs := run.fn(empty); len(probs) != 0 {
				t.Fatalf("explicit empty selection: %v", probs)
			}
			for _, bad := range []struct {
				ids  []string
				want string
			}{
				{[]string{"missing"}, "unknown selected scheme missing"},
				{[]string{"selected", "selected"}, "duplicate selected scheme selected"},
			} {
				c := base
				c.Schemes = bad.ids
				if probs := run.fn(c); !strings.Contains(strings.Join(probs, "\n"), bad.want) {
					t.Errorf("invalid selection %v was not rejected: %v", bad.ids, probs)
				}
			}
			// Changing the selected record in the file must change this case's
			// result: selection must not use a copied implementation fixture.
			mutated := strings.Replace(selectionCaseLibrary, "premises: [p(X), q(X)]", "premises: [p(X)]", 1)
			if err := os.WriteFile(path, []byte(mutated), 0600); err != nil {
				t.Fatal(err)
			}
			probs := run.fn(selected)
			if err := os.WriteFile(path, []byte(selectionCaseLibrary), 0600); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(strings.Join(probs, "\n"), "an argument of selected was built") {
				t.Fatalf("mutation of the live selected rule was hidden: %v", probs)
			}
		})
	}
	casePath := filepath.Join(root, "cases.yml")
	if err := os.WriteFile(casePath, []byte("cases:\n- name: selected\n  schemes: [selected]\n- name: none\n  schemes: []\n- name: all\n"), 0600); err != nil {
		t.Fatal(err)
	}
	cases, err := LoadCases(casePath)
	if err != nil || len(cases) != 3 || len(cases[0].Schemes) != 1 || cases[1].Schemes == nil || cases[2].Schemes != nil {
		t.Fatalf("selection YAML must distinguish explicit empty from omitted: cases=%v err=%v", cases, err)
	}
}

func TestCaseSelectionRebuildsIndex(t *testing.T) {
	ResetEngineGlobals()
	defer ResetEngineGlobals()
	ag, err := cyaml.Import(strings.NewReader(selectionCaseLibrary))
	if err != nil {
		t.Fatal(err)
	}
	old := ag.Theory
	selected := old.ArgSchemes[0]
	if err := selectCaseSchemes(ag, []string{"selected"}); err != nil {
		t.Fatal(err)
	}
	if ag.Theory == old || len(ag.Theory.ArgSchemes) != 1 || ag.Theory.ArgSchemes[0] != selected {
		t.Fatal("selection must use a fresh theory and the original imported record")
	}
	ag.InstantiateScheme("derive", []string{"a"})
	if len(ag.Arguments) != 0 {
		t.Fatal("excluded rule survived in the private scheme index")
	}
	ag.InstantiateScheme("selected", []string{"a"})
	if len(ag.Arguments) != 1 {
		t.Fatal("selected rule is missing from the fresh scheme index")
	}
}

func TestCaseSelectionReadsWaltonLibrary(t *testing.T) {
	cli := caseTestCLI(t)
	c := Case{
		Name:        "selection/live-walton",
		Schemes:     []string{"from_example"},
		Assumptions: []string{"has_properties(tweety,bird,flies)"},
		Arguments:   []string{"from_example(tweety,bird,flies)"},
		NoArguments: []string{"illustration"},
		In:          []string{"usually_properties(bird,flies)"},
		NotIn:       []string{"valid_rule(bird,flies)"},
	}
	if probs := RunCase(RepoRoot(), c); len(probs) != 0 {
		t.Fatalf("live walton API: %v", probs)
	}
	if probs := CaseRoundTrip(cli, RepoRoot(), c); len(probs) != 0 {
		t.Fatalf("live walton CLI: %v", probs)
	}
}

// The engine output is stable in all of these mutations. Stability alone must
// not make a wrong source-case expectation pass the CLI round-trip oracle.
func TestCaseRoundTripChecksEveryExpectation(t *testing.T) {
	cli := caseTestCLI(t)
	base := Case{
		Name: "oracle/control",
		Graph: `language:
  p/1: "%s has P"
  q/1: "%s has Q"
  r/1: "%s has R"
argument_schemes:
  - id: selected
    variables: [X]
    premises: [p(X)]
    conclusions: [q(X)]
statements:
  q(a): Q
  ¬q(a): Not Q
  r(a): R
issues:
  q_issue:
    positions: [q(a), ¬q(a)]
`,
		Assumptions: []string{"p(a)"},
		Arguments:   []string{"selected(a)"},
		NoArguments: []string{"selected(b)", "absent_scheme"},
		In:          []string{"q(a)"}, Out: []string{"¬q(a)"}, NotIn: []string{"r(a)"},
		Issues: [][]string{{"q(a)", "¬q(a)"}},
	}
	for _, run := range []struct {
		name string
		fn   func(Case) []string
	}{
		{"api", func(c Case) []string { return RunCase("", c) }},
		{"cli", func(c Case) []string { return CaseRoundTrip(cli, "", c) }},
	} {
		t.Run(run.name, func(t *testing.T) {
			if probs := run.fn(base); len(probs) != 0 {
				t.Fatalf("positive control: %v", probs)
			}
			for _, tc := range []struct {
				name, want string
				mutate     func(*Case)
			}{
				{"in", "want in", func(c *Case) { c.In = []string{"r(a)"} }},
				{"not-in", "want not in", func(c *Case) { c.NotIn = []string{"q(a)"} }},
				{"out", "want out", func(c *Case) { c.Out = []string{"q(a)"} }},
				{"argument", "no argument selected(b)", func(c *Case) { c.Arguments = []string{"selected(b)"} }},
				{"no-argument", "an argument selected(a) was built", func(c *Case) { c.NoArguments = []string{"selected(a)"} }},
				{"no-scheme", "an argument of selected was built", func(c *Case) { c.NoArguments = []string{"selected"} }},
				{"issue", "no issue with the positions", func(c *Case) { c.Issues = [][]string{{"q(a)", "r(a)"}} }},
			} {
				t.Run(tc.name, func(t *testing.T) {
					c := base
					tc.mutate(&c)
					if probs := run.fn(c); !strings.Contains(strings.Join(probs, "\n"), tc.want) {
						t.Fatalf("stable wrong result was not rejected for %s: %v", tc.name, probs)
					} else if run.name == "cli" {
						for _, phase := range []string{"first output: ", "after the round trip: "} {
							found := false
							for _, problem := range probs {
								if strings.HasPrefix(problem, phase) && strings.Contains(problem, tc.want) {
									found = true
								}
							}
							if !found {
								t.Errorf("%s expectation was not checked in %s: %v", tc.name, phase, probs)
							}
						}
					}
				})
			}
		})
	}
}
