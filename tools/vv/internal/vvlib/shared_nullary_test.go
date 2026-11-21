package vvlib

import (
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
)

func nullaryControl() *Library {
	ag := caes.NewArgGraph()
	ag.Theory.Language["do_obligator/0"] = "One must not do what is prohibited, and must do what is obligatory."
	ag.Theory.ArgSchemes = []*caes.Scheme{
		{Id: "regulative_rule", Premises: []string{"do_obligator"}},
		{Id: "regulative_rule_prohibition", Premises: []string{"do_obligator"}},
	}
	return &Library{Rel: "walton.yml", AG: ag}
}

func TestReviewedSharedNullaryLibraryControl(t *testing.T) {
	lib, err := LoadLibrary(RepoRoot(), "walton.yml")
	if err != nil {
		t.Fatal(err)
	}
	if result := SharedNullaryCheck(lib); !result.Pass {
		t.Fatalf("reviewed master library rejected: %s", result.Detail)
	}
	if result := SharedNullaryCheck(nullaryControl()); !result.Pass {
		t.Fatalf("explicit reviewed-warrant control rejected: %s", result.Detail)
	}
}

func TestSharedNullaryMutations(t *testing.T) {
	cases := []struct {
		name, problem string
		mutate        func(*Library)
	}{
		{"unexpected-sharing", "unreviewed shared nullary case_context/0", func(lib *Library) {
			lib.Schemes()[0].Assumptions = []string{"case_context"}
			lib.Schemes()[1].Exceptions = []string{"case_context"}
		}},
		{"sharing-through-negation", "unreviewed shared nullary case_context/0", func(lib *Library) {
			lib.Schemes()[0].Premises = append(lib.Schemes()[0].Premises, "case_context")
			lib.Schemes()[1].Conclusions = []string{"¬case_context"}
		}},
		{"new-consumer", "reviewed consumers", func(lib *Library) {
			lib.AG.Theory.ArgSchemes = append(lib.Schemes(), &caes.Scheme{Id: "unreviewed", Premises: []string{"do_obligator"}})
		}},
		{"removed-consumer", "reviewed consumers", func(lib *Library) {
			lib.Schemes()[1].Premises = []string{"different_warrant"}
		}},
		{"renamed-consumer", "reviewed consumers", func(lib *Library) {
			lib.Schemes()[1].Id = "unreviewed"
		}},
		{"changed-wording", "universal-warrant wording changed", func(lib *Library) {
			lib.AG.Theory.Language["do_obligator/0"] = "The current agent must perform this particular action."
		}},
		{"negated-approved-warrant", "must be a positive premise", func(lib *Library) {
			lib.Schemes()[1].Premises = []string{"¬do_obligator"}
		}},
		{"approved-warrant-becomes-default", "must be a positive premise", func(lib *Library) {
			lib.Schemes()[1].Premises = nil
			lib.Schemes()[1].Assumptions = []string{"do_obligator"}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lib := nullaryControl()
			tc.mutate(lib)
			result := SharedNullaryCheck(lib)
			if result.Pass || !strings.Contains(result.Detail, tc.problem) {
				t.Fatalf("mutation not rejected for the intended reason: %+v", result)
			}
		})
	}
}

func TestSharedNullaryIgnoresNonsharingAndDifferentArities(t *testing.T) {
	lib := nullaryControl()
	lib.Schemes()[0].Assumptions = []string{"local_warrant", "local_warrant"}
	lib.Schemes()[0].Conclusions = []string{"A", "holds(shared_constant)"}
	lib.Schemes()[1].Conclusions = []string{"¬A", "holds(shared_constant)"}
	lib.AG.Theory.ArgSchemes = append(lib.Schemes(), &caes.Scheme{Id: "parameterized", Premises: []string{"do_obligator(X)"}})
	if result := SharedNullaryCheck(lib); !result.Pass {
		t.Fatalf("local occurrences, variables, term arguments or another arity counted as shared nullaries: %s", result.Detail)
	}
}
