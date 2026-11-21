package vvlib

import (
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
)

// These fixtures use the actual inherited record, selected from walton.yml.
// Isolating it prevents practical_reasoning2's independent support for the
// same action from hiding a failure of practical_reasoning1. They verify the
// legacy AS1 abstraction, not full AS2 or the paper's preference resolution.
func atkinsonGraph(t *testing.T, facts, questioned []string) (*caes.ArgGraph, caes.Labelling) {
	t.Helper()
	lib, err := LoadLibrary(RepoRoot(), "walton.yml")
	if err != nil {
		t.Fatal(err)
	}
	s := lib.Scheme("practical_reasoning1")
	if s == nil {
		t.Fatal("no practical_reasoning1")
	}
	ag := caes.NewArgGraph()
	ag.Theory.Language = lib.AG.Theory.Language
	ag.Theory.WeighingFunctions = lib.AG.Theory.WeighingFunctions
	ag.Theory.IssueSchemes = lib.AG.Theory.IssueSchemes
	ag.Theory.ArgSchemes = []*caes.Scheme{s}
	ag.Theory.InitSchemeIndex()
	ag.Assumptions = append([]string{}, facts...)
	ag.Questioned = append([]string{}, questioned...)
	ResetEngineGlobals()
	cap := StartCapture()
	err = ag.Infer()
	cap.Stop()
	ResetEngineGlobals()
	if err != nil {
		t.Fatalf("Infer must complete before judging the case: %v", err)
	}
	return ag, ag.GroundedLabelling()
}

func atkinsonArgument(t *testing.T, ag *caes.ArgGraph, call string) *caes.Argument {
	t.Helper()
	for _, a := range ag.Arguments {
		if a != nil && a.Scheme != nil && Norm(a.Scheme.Id+"("+strings.Join(a.Parameters, ",")+")") == Norm(call) {
			return a
		}
	}
	t.Fatalf("no argument %s", call)
	return nil
}

// Source: Atkinson and Bench-Capon (2007), p. 866, Arg8, abstracted to AS1.
// The three default burdens and the aggregate side-effect exception are
// explicitly inherited policies, not claimed to be prescribed by the paper.
func TestAtkinsonInheritedQuestionScope(t *testing.T) {
	base := []string{
		"current_circumstances(q0)",
		"would_bring_about(row_chicken,q0,q9)",
		"would_be_realized(chicken_right,q9)",
		"would_promote_value(chicken_right,progress)",
	}
	cases := []struct {
		name       string
		facts      []string
		questioned []string
		positive   bool
		undercut   bool
	}{
		{name: "positive", positive: true},
		{name: "cq13-questioned-action", questioned: []string{"possible(row_chicken)"}},
		{name: "cq16-questioned-value", questioned: []string{"legitimate_value(progress)"}},
		{name: "inherited-worthiness-policy", questioned: []string{"worthy_goal(chicken_right)"}},
		{name: "other-action-questioned", questioned: []string{"possible(row_dog)"}, positive: true},
		{name: "other-value-questioned", questioned: []string{"legitimate_value(friendship)"}, positive: true},
		{name: "inherited-side-effect-policy", facts: []string{"side_effects(row_chicken,q0)"}, undercut: true},
		{name: "side-effect-other-state", facts: []string{"side_effects(row_chicken,q9)"}, positive: true},
		{name: "side-effect-other-action", facts: []string{"side_effects(row_dog,q0)"}, positive: true},
		{name: "questioned-action-supported", facts: []string{"possible(row_chicken)"}, questioned: []string{"possible(row_chicken)"}, positive: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			facts := append(append([]string{}, base...), tc.facts...)
			ag, l := atkinsonGraph(t, facts, tc.questioned)
			a := atkinsonArgument(t, ag, "practical_reasoning1(row_chicken,q0,q9,chicken_right,progress)")
			if got := a.GetWeight(l) > 0; got != tc.positive {
				t.Errorf("argument positive weight = %v; want %v", got, tc.positive)
			}
			if got := l[a.Conclusion] == caes.In; got != tc.positive {
				t.Errorf("isolated conclusion In = %v; want %v", got, tc.positive)
			}
			if got := a.Undercut(l) == caes.In; got != tc.undercut {
				t.Errorf("undercutter In = %v; want %v; questioning is not an asserted exception", got, tc.undercut)
			}
		})
	}
}

// CQ10 is explicitly supportive (p. 862). Arg6 (p. 866) supplies P and F.
// Its q0 -> q7 transition is stipulated here: this is the initial argument,
// not the paper's later decision after CQ17 and value preferences are applied.
func TestAtkinsonCQ10AddsSupport(t *testing.T) {
	facts := []string{
		"current_circumstances(q0)",
		"would_bring_about(row_dog,q0,q7)",
		"would_be_realized(dog_right,q7)",
		"would_promote_value(dog_right,progress)",
		"would_promote_value(dog_right,friendship)",
	}
	for _, questionFriendship := range []bool{false, true} {
		name := "both-values-supported"
		var questioned []string
		if questionFriendship {
			name = "only-additional-value-questioned"
			questioned = []string{"legitimate_value(friendship)"}
		}
		t.Run(name, func(t *testing.T) {
			ag, l := atkinsonGraph(t, facts, questioned)
			progress := atkinsonArgument(t, ag, "practical_reasoning1(row_dog,q0,q7,dog_right,progress)")
			friendship := atkinsonArgument(t, ag, "practical_reasoning1(row_dog,q0,q7,dog_right,friendship)")
			if progress.Id == friendship.Id {
				t.Fatal("different value justifications were merged")
			}
			if progress.GetWeight(l) <= 0 || l[progress.Conclusion] != caes.In {
				t.Fatal("additional positive value defeated the original argument")
			}
			if got := friendship.GetWeight(l) > 0; got == questionFriendship {
				t.Errorf("additional argument positive weight = %v; questioned = %v", got, questionFriendship)
			}
			if progress.Undercut(l) == caes.In || friendship.Undercut(l) == caes.In {
				t.Fatal("a positive side effect or a bare question was made an undercutter")
			}
		})
	}
}
