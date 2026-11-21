package caes

import (
	"context"
	"errors"
	"testing"
)

func TestReviewAlreadyCancelledEmptyTheory(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	for i := 0; i < 30; i++ {
		ag := NewArgGraph()
		ag.Assumptions = []string{"p(a)"}
		if err := ag.InferContext(ctx); !errors.Is(err, ErrCancelled) {
			t.Fatalf("attempt %d: cancelled empty theory: %v", i, err)
		}
		if len(ag.Statements) != 0 {
			t.Fatal("cancelled run changed the graph")
		}
	}
}

func TestReviewGeneratedIssueRejectsAcceptedPositions(t *testing.T) {
	for _, explicit := range []bool{false, true} {
		ag := NewArgGraph()
		p := IssueScheme{"accepts(X)", "rejects(X)"}
		ag.Theory.IssueSchemes["choice"] = &p
		ag.Assumptions = []string{"accepts(proposal)", "rejects(proposal)"}
		if explicit {
			for _, id := range ag.Assumptions {
				ag.Statements[id] = &Statement{Id: id, Text: id}
			}
		}
		before := len(ag.Statements)
		if err := ag.Infer(); !errors.Is(err, ErrInconsistentContext) {
			t.Errorf("explicit=%v: want inconsistent-context, got %v", explicit, err)
		}
		if len(ag.Statements) != before || len(ag.Issues) != 0 {
			t.Errorf("explicit=%v: failed run changed the graph", explicit)
		}
	}
}

func TestReviewInvalidIssueLeavesGraphUnchanged(t *testing.T) {
	ag := NewArgGraph()
	p := IssueScheme{"p(X)", "q("}
	ag.Theory.IssueSchemes["bad"] = &p
	ag.Assumptions = []string{"p(a)"}
	if err := ag.Infer(); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("want invalid-input, got %v", err)
	}
	if len(ag.Statements) != 0 || len(ag.Issues) != 0 {
		t.Fatal("invalid issue scheme changed the graph")
	}
}

func TestReviewGeneratedIssuePreservesDefaultPriority(t *testing.T) {
	for _, bothDefaults := range []bool{false, true} {
		ag := NewArgGraph()
		p := IssueScheme{"accepts(X)", "rejects(X)"}
		ag.Theory.IssueSchemes["choice"] = &p
		ag.Assumptions = []string{"accepts(proposal)", "rejects(proposal)"}
		ag.Defaults = map[string]bool{"accepts(proposal)": true}
		if bothDefaults {
			ag.Defaults["rejects(proposal)"] = true
		}
		if err := ag.Infer(); err != nil {
			t.Fatalf("default conflict is not inconsistent stipulation: %v", err)
		}
		l := ag.GroundedLabelling()
		if l[ag.Statements["accepts(proposal)"]] == In {
			t.Error("default remained accepted")
		}
		if got := l[ag.Statements["rejects(proposal)"]] == In; got == bothDefaults {
			t.Errorf("contrary accepted=%v with bothDefaults=%v", got, bothDefaults)
		}
	}
}

func TestReviewEmptyTheoryRejectsMalformedAssumption(t *testing.T) {
	ag := NewArgGraph()
	ag.Assumptions = []string{"p("}
	if err := ag.Infer(); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("want invalid-input: %v", err)
	}
	if len(ag.Statements) != 0 {
		t.Fatal("invalid run committed a statement")
	}
}
