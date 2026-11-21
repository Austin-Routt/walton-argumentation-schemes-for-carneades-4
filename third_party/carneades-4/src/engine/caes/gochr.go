// Copyright © 2016 The Carneades Authors
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL
// was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.

// Wrapper for the GoCHR inference engine

package caes

import (
	"errors"
	"fmt"
	// "log"
	"strings"

	chr "github.com/hfried/GoCHR/src/engine/CHR"
)

// Translate a theory into a GoCHR rulestore. A scheme whose rule GoCHR
// refuses is left out; RuleStoreOf reports it.
func TheoryToRuleStore(t *Theory) *chr.RuleStore {
	rs, _ := RuleStoreOf(t)
	return rs
}

// RuleStoreOf translates a theory into a GoCHR rule store, one rule per
// scheme with conclusions. The error names every scheme whose rule GoCHR
// refuses, for example for a premise it cannot read (E-PARSE); those schemes
// are left out of the store.
func RuleStoreOf(t *Theory) (*chr.RuleStore, error) {
	// log.Printf("TheoryToRuleStore\n") // DEBUG
	rs := chr.MakeRuleStore()
	var errs []string
	// Argument construction is forward chaining over facts, and facts form a
	// set (docs/contracts.md, contract 2): a fact derived a second time must
	// not trigger the schemes again. Without this, schemes that re-derive a
	// fact, or derive each other's premises, run to MAXRULEAPPS.
	rs.Unique = true
	rs.MaxTermSize = MAXTERMSIZE
	for _, s := range t.ArgSchemes {
		// If the scheme has no conclusions, skip the scheme
		// and assume it only defines a weighing function but no rule
		if len(s.Conclusions) > 0 {
			// A "go" term is added to CHR rules for
			// argument schemes with no premises, since CHR requires
			// rules to have at least one term in the head.
			var premises []string
			if len(s.Premises) > 0 {
				premises = s.Premises
			} else if len(s.Deletions) == 0 {
				premises = []string{"go"}
			}
			argTerm := fmt.Sprintf("argument(%s,[%s])", s.Id, strings.Join(s.Variables, ","))
			// A new slice: appending to s.Conclusions could write into the
			// scheme's own backing array.
			conclusions := make([]string, 0, len(s.Conclusions)+len(s.Assumptions)+1)
			conclusions = append(conclusions, s.Conclusions...)
			conclusions = append(conclusions, s.Assumptions...)
			conclusions = append(conclusions, argTerm)
			// Note that the body of the rules includes the assumptions
			// and conclusions of the scheme.
			// fmt.Printf("AddRule id=%s, premises=%s, deletions=%s, guards=%s, conclusions=%s\n", s.Id, premises, s.Deletions, s.Guards, conclusions)
			if err := rs.AddRule(s.Id, premises, s.Deletions, s.Guards, conclusions); err != nil {
				errs = append(errs, fmt.Sprintf("scheme %s: %v", s.Id, err))
			}
		}
	}
	if len(errs) != 0 {
		return rs, errors.New(strings.Join(errs, "; "))
	}
	return rs, nil
}
