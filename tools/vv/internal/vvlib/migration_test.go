package vvlib

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
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

// Historical predicate renames must not erase the distinction between two old
// forms that now lead to one result. Triggers also distinguish migrations, and
// row order is never part of their identity.
func TestMigrationIdentitiesUseOriginalMappingAndSurviveReordering(t *testing.T) {
	root := t.TempDir()
	t.Setenv("WALTON_REPO", root)
	if err := os.Mkdir(filepath.Join(root, "tools"), 0o755); err != nil {
		t.Fatal(err)
	}
	rows := []string{
		`- {record: s, kind: rename-predicate, old: old/1, new: fresh/1, plan: test}`,
		`- {record: s, kind: arity, old: "old(X)", trigger: "context_a(X)", new: "fresh(X,Y)", plan: test, page: 1}`,
		`- {record: s, kind: arity, old: "fresh(X)", trigger: "context_a(X)", new: "fresh(X,Y)", plan: test, page: 1}`,
		`- {record: s, kind: arity, old: "fresh(X)", trigger: "context_b(X)", new: "fresh(X,Y)", plan: test, page: 1}`,
		`- {record: s, kind: arity, old: "fresh(X)", trigger: "context_a(X)", new: "another(X,Y)", plan: test, page: 1}`,
		`- {record: s, kind: arguments, old: "fresh(X)", trigger: "context_a(X)", new: "fresh(X,Y)", plan: test, page: 1}`,
	}
	ag := caes.NewArgGraph()
	ag.Theory.Language["fresh/1"] = "%s is fresh"
	ag.Theory.ArgSchemes = []*caes.Scheme{{Id: "s", Premises: []string{"fresh(X,Y)", "another(X,Y)"}}}
	lib := &Library{AG: ag}
	check := func(rows []string) ([]string, error) {
		t.Helper()
		text := "changes:\n" + strings.Join(rows, "\n") + "\n"
		if err := os.WriteFile(filepath.Join(root, "tools", "migration_r2.yml"), []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
		results := MigrationChecks(lib)
		if len(results) != len(rows)+1 {
			t.Fatalf("lost migration checks: %d results for %d rows", len(results), len(rows))
		}
		for _, r := range results {
			if !r.Pass {
				t.Fatalf("control migration failed: %+v", r)
			}
		}
		return ResultKeys(results)
	}
	first, err := check(rows)
	if err != nil {
		t.Fatal(err)
	}
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i]
	}
	second, err := check(rows)
	if err != nil || !reflect.DeepEqual(first, second) {
		t.Fatalf("reordering altered identities: %v versus %v, %v", first, second, err)
	}
	// A genuinely repeated mapping remains an error instead of being assigned
	// a synthetic row index that conceals the duplicate.
	if _, err := check(append(rows, rows[0])); err == nil || !strings.Contains(err.Error(), "duplicate check key") {
		t.Fatalf("identical migration rows did not fail inventory validation: %v", err)
	}
}
