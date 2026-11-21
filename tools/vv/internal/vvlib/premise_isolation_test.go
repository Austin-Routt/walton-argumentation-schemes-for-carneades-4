package vvlib

import (
	"strings"
	"testing"

	cyaml "github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
)

// A withdrawn stipulation is not necessarily missing from the full closure.
// These controls retain full-theory composition while checking the selected
// rule's own premise necessity, and still reject a truly superfluous premise.
func TestPremiseNecessityAllowsOtherSchemeRederivation(t *testing.T) {
	data := []byte(`language:
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
`)
	c := Case{Name: "selected/control", Graph: string(data), Assumptions: []string{"p(a)", "q(a)"}, Arguments: []string{"selected(a)"}, In: []string{"r(a)"}}
	full := c
	full.Assumptions = []string{"p(a)"}
	if probs := RunCase("", full); len(probs) != 0 {
		t.Fatalf("the full theory must rederive Q and construct selected: %v", probs)
	}
	if probs := variantProblems(data, c); len(probs) != 0 {
		t.Fatalf("isolated reference-case withdrawals: %v", probs)
	}
	ag, err := cyaml.Import(strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	s := ag.Theory.ArgSchemes[0]
	if probs := premiseProblems(ag.Theory.Language, ag.Theory.ArgSchemes, s, NewInstance(s), ""); len(probs) != 0 {
		t.Fatalf("isolated generic withdrawals: %v", probs)
	}
	broken := []byte(strings.Replace(string(data), "premises: [p(X), q(X)]", "premises: [p(X)]", 1))
	if probs := variantProblems(broken, c); !strings.Contains(strings.Join(probs, "\n"), "without q(a) the argument is still built") {
		t.Fatalf("a genuinely unnecessary reference premise must still fail: %v", probs)
	}
	missing := c
	missing.Assumptions = []string{"q(a)"}
	if probs := variantProblems(data, missing); !strings.Contains(strings.Join(probs, "\n"), "isolated positive control") {
		t.Fatalf("an absent positive control must not count as successful withdrawal: %v", probs)
	}
}
