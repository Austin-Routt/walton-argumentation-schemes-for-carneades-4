package vvlib

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/carneades/carneades-4/src/engine/terms"
)

// SharedNullaryCheck closes the shared-0ary parity item (PLAN.md 5.4).
// Cross-scheme nullaries require an explicit reviewed global meaning. The
// universal warrant of entry 55.3 (p. 343) is intentionally a positive premise
// of its obligation and prohibition variants; see CHANGELOG.md, 3.2 and 5.4,
// and their source-ledger rows. Other nullaries may remain local to one scheme.
func SharedNullaryCheck(lib *Library) Result {
	approved := map[string]struct {
		Consumers []string
		Template  string
	}{
		"do_obligator/0": {
			Consumers: []string{"regulative_rule", "regulative_rule_prohibition"},
			Template:  "One must not do what is prohibited, and must do what is obligatory.",
		},
	}
	uses := map[string]map[string]bool{}
	problems := []string{}
	for _, scheme := range lib.Schemes() {
		for _, field := range []struct {
			Name       string
			Statements []string
		}{{"premises", scheme.Premises}, {"assumptions", scheme.Assumptions},
			{"exceptions", scheme.Exceptions}, {"conclusions", scheme.Conclusions}} {
			for _, statement := range field.Statements {
				t, ok := parse(statement)
				if !ok {
					problems = append(problems, fmt.Sprintf("%s %s: cannot parse %q", scheme.Id, field.Name, statement))
					continue
				}
				negative := false
				for {
					c, ok := t.(terms.Compound)
					if !ok || c.Functor != "¬" || len(c.Args) != 1 {
						break
					}
					negative, t = !negative, c.Args[0]
				}
				name, predicate := terms.Functor(t)
				if !predicate || terms.Arity(t) != 0 {
					continue // variables and nested term arguments are not nullary statements
				}
				key := name + "/0"
				if uses[key] == nil {
					uses[key] = map[string]bool{}
				}
				uses[key][scheme.Id] = true
				if _, reviewed := approved[key]; reviewed && (negative || field.Name != "premises") {
					problems = append(problems, fmt.Sprintf("%s: reviewed global warrant %s must be a positive premise", scheme.Id, key))
				}
			}
		}
	}
	consumers := func(key string) []string {
		ids := []string{}
		for id := range uses[key] {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		return ids
	}
	for key, ids := range uses {
		if _, reviewed := approved[key]; len(ids) > 1 && !reviewed {
			problems = append(problems, fmt.Sprintf("unreviewed shared nullary %s used by %v", key, consumers(key)))
		}
	}
	for key, approval := range approved {
		if got := consumers(key); !reflect.DeepEqual(got, approval.Consumers) {
			problems = append(problems, fmt.Sprintf("%s: reviewed consumers %v changed to %v", key, approval.Consumers, got))
		}
		if lib.AG.Theory.Language[key] != approval.Template {
			problems = append(problems, fmt.Sprintf("%s: reviewed universal-warrant wording changed", key))
		}
	}
	sort.Strings(problems)
	return res(lib.Rel, "(file)", "static/shared-0ary", problems)
}
