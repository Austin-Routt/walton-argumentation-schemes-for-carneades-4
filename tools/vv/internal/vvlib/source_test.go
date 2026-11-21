package vvlib

import (
	"strings"
	"testing"
)

// A row is at "interpretation approved" or beyond only with an approving
// review, and a review that holds a row or leaves it open says why (D14).
func TestReviewProblems(t *testing.T) {
	ok := &RecordReview{By: "a reviewer", Date: "2026-11-21", Decision: "approved"}
	cases := []struct {
		name  string
		stage string
		rev   *RecordReview
		want  string // "" for no problem, or a part of the problem's text
	}{
		{"checked, no review", "transcription-checked", nil, ""},
		{"approved, no review", "interpretation-approved", nil, "needs a review"},
		{"example passed, no review", "example-passed", nil, "needs a review"},
		{"approved with a review", "example-passed", ok, ""},
		{"approving review, stage behind", "transcription-checked", ok, "but its stage is"},
		{"hold with its reason", "transcription-checked", &RecordReview{By: "r", Date: "2026-11-21", Decision: "hold", Note: "the event has no variable"}, ""},
		{"hold without a reason", "transcription-checked", &RecordReview{By: "r", Date: "2026-11-21", Decision: "hold"}, "says in its note"},
		{"hold past approval", "example-passed", &RecordReview{By: "r", Date: "2026-11-21", Decision: "hold", Note: "x"}, "the review's decision is"},
		{"open with its reason", "mapped", &RecordReview{By: "r", Date: "2026-11-21", Decision: "open", Note: "the source is not available"}, ""},
		{"no reviewer", "example-passed", &RecordReview{Date: "2026-11-21", Decision: "approved"}, "who made it"},
		{"no date", "example-passed", &RecordReview{By: "r", Decision: "approved"}, "not a date"},
		{"a date that is none", "example-passed", &RecordReview{By: "r", Date: "21 November", Decision: "approved"}, "not a date"},
		{"unknown decision", "transcription-checked", &RecordReview{By: "r", Date: "2026-11-21", Decision: "fine"}, "unknown review decision"},
	}
	for _, c := range cases {
		got := reviewProblems(RecordRow{Record: "x", Stage: c.stage, Review: c.rev})
		switch {
		case c.want == "" && len(got) > 0:
			t.Errorf("%s: unexpected problem(s) %v", c.name, got)
		case c.want != "" && !strings.Contains(strings.Join(got, "; "), c.want):
			t.Errorf("%s: got %v, want a problem containing %q", c.name, got, c.want)
		}
	}
}
