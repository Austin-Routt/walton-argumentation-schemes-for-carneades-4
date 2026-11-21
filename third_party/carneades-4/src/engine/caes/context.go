package caes

import (
	"errors"
	"github.com/carneades/carneades-4/src/engine/terms"
)

// Inconsistent reports contradictory accepted context, including issue
// schemes not yet materialized in the graph. CheckContext also reports
// malformed issue patterns. Neither method changes the graph.
func (ag *ArgGraph) Inconsistent() bool {
	return errors.Is(ag.CheckContext(), ErrInconsistentContext)
}

// CheckContext validates issue patterns and tests the currently accepted
// context. A private graph preserves the origin and questioning of defaults,
// so a default yields to a stipulated contrary instead of being counted as a
// contradictory stipulation. No argument construction or labelling is needed.
func (ag *ArgGraph) CheckContext() error {
	for _, a := range ag.Assumptions {
		if t, ok := terms.ReadString(a); !ok || t == nil {
			return NewRunError(ErrInvalidInput, "cannot parse assumption %q", a)
		}
	}
	if ag.inconsistentPositions() {
		return NewRunError(ErrInconsistentContext, "the graph assumes two positions of one issue, or a statement and its negation")
	}
	if ag.Theory == nil {
		return nil
	}
	for name, patterns := range ag.Theory.IssueSchemes {
		if patterns == nil {
			return NewRunError(ErrInvalidInput, "issue scheme %s has no patterns", name)
		}
		if _, err := readIssuePatterns(*patterns); err != nil {
			return NewRunError(ErrInvalidInput, "issue scheme %s: %v", name, err)
		}
		probe := NewArgGraph()
		probe.Assumptions = ag.Assumptions
		probe.Defaults = ag.Defaults
		probe.Questioned = ag.Questioned
		for _, a := range ag.Assumptions {
			probe.Statements[a] = &Statement{Id: a, Text: a}
		}
		if err := probe.makeIssue(name, *patterns); err != nil {
			return NewRunError(ErrInvalidInput, "issue scheme %s: %v", name, err)
		}
		if probe.inconsistentPositions() {
			return NewRunError(ErrInconsistentContext, "accepted context contains two positions of issue scheme %s", name)
		}
	}
	return nil
}
