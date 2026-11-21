package vvlib

import (
	"strings"
	"testing"
)

// PLAN.md 5.6: the fixture check looks for old identifiers only where
// identifiers stand. An old name used as an ordinary word in the text of a
// template or a statement ("if every part has it") is not one.
func TestIdentifierText(t *testing.T) {
	src := strings.Join([]string{
		"# has is mentioned in a comment",
		"language:",
		`  compositionally_hereditary/2: "if every part has it, the whole has it."`,
		`  has/2: "%s has %s."`,
		"statements:",
		`  has(car,heavy): "car has heavy"`,
		`  "group_accepted(g,has(a,b))": "has, in text"`,
		"argument_schemes:",
		"  - id: group_1",
		"    conclusions:",
		"      - has(G,Q)",
		"assumptions:",
		"  - has(car,heavy)",
	}, "\n")
	got := identifierText(src)
	for _, keep := range []string{"  has/2\n", "  has(car,heavy)\n", `  "group_accepted(g,has(a,b))"` + "\n", "      - has(G,Q)\n", "  - has(car,heavy)\n"} {
		if !strings.Contains(got, keep) {
			t.Errorf("lost %q in\n%s", keep, got)
		}
	}
	for _, drop := range []string{"comment", "every part", "%s has %s", "car has heavy", "has, in text"} {
		if strings.Contains(got, drop) {
			t.Errorf("kept text %q in\n%s", drop, got)
		}
	}
}
