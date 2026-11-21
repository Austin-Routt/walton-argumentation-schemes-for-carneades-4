package vvlib

import (
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
)

func hasSourceProblem(probs []string, text string) bool {
	return strings.Contains(strings.Join(probs, "\n"), text)
}

// An original source's inventory is independent of a compendium entry. In
// particular, the 17th CQ in Atkinson and Bench-Capon (2007) must not disappear
// merely because practical_reasoning1 is filed under upstream_only.
func TestOriginalSourceQuestionInventory(t *testing.T) {
	n := 17
	s := &caes.Scheme{Conclusions: []string{"do(A)"}}
	base := RecordRow{Origin: "upstream", SourceQuestions: &n}
	for i := 1; i <= n; i++ {
		base.Questions = append(base.Questions, RecordQuestion{N: i, Kind: "premise"})
	}
	base.Questions[9] = RecordQuestion{N: 10, Kind: "support", As: "do(A)"}
	if probs := questionProblems(base, s, n, true); len(probs) != 0 {
		t.Fatalf("complete inventory: %v", probs)
	}
	missing := base
	missing.Questions = append([]RecordQuestion{}, base.Questions[:16]...)
	if probs := questionProblems(missing, s, n, true); !hasSourceProblem(probs, "question 17 of 17 is not sorted") {
		t.Fatalf("missing CQ17 accepted: %v", probs)
	}
	missingSupport := base
	missingSupport.Questions = append(append([]RecordQuestion{}, base.Questions[:9]...), base.Questions[10:]...)
	if probs := questionProblems(missingSupport, s, n, true); !hasSourceProblem(probs, "question 10 of 17 is not sorted") {
		t.Fatalf("missing positive CQ10 accepted: %v", probs)
	}
	extra := base
	extra.Questions = append(append([]RecordQuestion{}, base.Questions...), RecordQuestion{N: 18, Kind: "premise"})
	if probs := questionProblems(extra, s, n, false); !hasSourceProblem(probs, "question 18: the declared source prints 17") {
		t.Fatalf("out-of-range original-source CQ accepted: %v", probs)
	}
}

func TestSourceQuestionKindCheckedBeforeBlankMapping(t *testing.T) {
	s := &caes.Scheme{}
	for _, kind := range []string{"", "typo"} {
		for _, complete := range []bool{false, true} {
			r := RecordRow{Origin: "book", Questions: []RecordQuestion{{N: 1, Kind: kind}}}
			if probs := questionProblems(r, s, 1, complete); !hasSourceProblem(probs, "unknown kind") {
				t.Fatalf("kind %q, complete %v accepted: %v", kind, complete, probs)
			}
		}
	}
}

func TestSupportingQuestionPreservesPolarity(t *testing.T) {
	s := &caes.Scheme{Conclusions: []string{"do(A)"}}
	r := RecordRow{Origin: "book", Questions: []RecordQuestion{{N: 1, Kind: "support", As: "do(A)"}}}
	if probs := questionProblems(r, s, 1, true); len(probs) != 0 {
		t.Fatalf("same-conclusion support rejected: %v", probs)
	}
	r.Questions[0].As = "¬do(A)"
	if probs := questionProblems(r, s, 1, true); !hasSourceProblem(probs, "not one of the record's conclusions") {
		t.Fatalf("support accepted as a rebuttal: %v", probs)
	}
}

func TestAlternativeCannotClaimCompletenessWithoutComparison(t *testing.T) {
	s := &caes.Scheme{Conclusions: []string{"do(A)"}}
	r := RecordRow{Origin: "book", Questions: []RecordQuestion{{N: 1, Kind: "alternative"}}}
	if probs := questionProblems(r, s, 1, false); len(probs) != 0 {
		t.Fatalf("honestly unresolved alternative rejected: %v", probs)
	}
	if probs := questionProblems(r, s, 1, true); !hasSourceProblem(probs, "is not modelled") {
		t.Fatalf("unresolved alternative accepted as complete: %v", probs)
	}
	r.Questions[0].As = "do(A)"
	if probs := questionProblems(r, s, 1, false); !hasSourceProblem(probs, "needs a comparison contract") {
		t.Fatalf("a bare action conclusion accepted as a comparison: %v", probs)
	}
}

func TestUnknownOriginalQuestionCountIsNotZero(t *testing.T) {
	s := &caes.Scheme{Conclusions: []string{"do(A)"}}
	r := RecordRow{Origin: "upstream"}
	if probs := questionProblems(r, s, 0, true); !hasSourceProblem(probs, "needs source_questions") {
		t.Fatalf("unknown source inventory accepted as complete: %v", probs)
	}
	zero := 0
	r.SourceQuestions = &zero
	if probs := questionProblems(r, s, 0, true); len(probs) != 0 {
		t.Fatalf("explicitly reviewed zero-question source rejected: %v", probs)
	}
}

func TestSourceWorklistCountsPrintedQuestionsOnly(t *testing.T) {
	n := 17
	l := &SourceLedger{
		UpstreamOnly: []string{"reviewed", "unknown"},
		Records: []RecordRow{
			{Record: "reviewed", Origin: "upstream", SourceQuestions: &n,
				Questions: []RecordQuestion{{N: 1}, {N: 2}, {N: 0}, {N: 0}}},
			{Record: "unknown", Origin: "upstream"},
		},
	}
	ag := caes.NewArgGraph()
	ag.Theory.ArgSchemes = []*caes.Scheme{{Id: "reviewed"}, {Id: "unknown"}}
	rows := SourceRows(l, &Library{AG: ag})
	if len(rows) != 2 || rows[0].Questions != 17 || !rows[0].QuestionCountKnown || rows[0].Sorted != 2 {
		t.Fatalf("printed CQs and policy additions conflated: %+v", rows)
	}
	if rows[1].QuestionCountKnown {
		t.Fatal("unknown source inventory displayed as a verified zero")
	}
}
