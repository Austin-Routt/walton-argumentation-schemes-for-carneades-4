// Copyright © 2026 The Carneades Authors
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package caes_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
	"github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
)

func infer(t *testing.T, src string) *caes.ArgGraph {
	t.Helper()
	ag, err := yaml.Import(strings.NewReader(src))
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if err := ag.Infer(); err != nil {
		t.Fatalf("Infer: %v", err)
	}
	return ag
}

// argumentsFor lists, sorted, the schemes of the arguments for a conclusion.
func argumentsFor(ag *caes.ArgGraph, conclusion string) []string {
	var ids []string
	for _, a := range ag.Arguments {
		if a != nil && a.Conclusion != nil && a.Conclusion.Id == conclusion {
			ids = append(ids, a.Scheme.Id+"("+strings.Join(a.Parameters, ",")+")")
		}
	}
	sort.Strings(ids)
	return ids
}

// Set semantics for facts (docs/contracts.md, contract 2) removes repeated
// facts, not arguments: two schemes concluding one statement give two
// arguments, and one scheme applied to two facts gives two.
func TestSetSemanticsKeepsDistinctArguments(t *testing.T) {
	ag := infer(t, `
language:
  p/1: "%s is p"
  q/1: "%s is q"
  r/1: "%s is r"
argument_schemes:
  - id: from_p
    variables: [X]
    premises:
      - p(X)
    conclusions:
      - q(X)
  - id: from_r
    variables: [X]
    premises:
      - r(X)
    conclusions:
      - q(X)
assumptions:
  - p(a)
  - r(a)
  - p(b)
`)
	if got, want := argumentsFor(ag, "q(a)"), []string{"from_p(a)", "from_r(a)"}; strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("arguments for q(a): got %v, want %v", got, want)
	}
	if got := argumentsFor(ag, "q(b)"); len(got) != 1 {
		t.Errorf("arguments for q(b): got %v, want one", got)
	}
}

// Two schemes that derive each other's premise stop after one round, each with
// its one argument, and a scheme listed after them still fires (X-LOOP).
func TestMutuallyDerivingSchemesTerminate(t *testing.T) {
	ag := infer(t, `
language:
  part/1: "%s is a part"
  whole/1: "%s is a whole"
  noted/1: "%s is noted"
argument_schemes:
  - id: composition
    variables: [X]
    premises:
      - part(X)
    conclusions:
      - whole(X)
  - id: division
    variables: [X]
    premises:
      - whole(X)
    conclusions:
      - part(X)
  - id: later
    variables: [X]
    premises:
      - part(X)
    conclusions:
      - noted(X)
assumptions:
  - part(a)
`)
	if got := argumentsFor(ag, "whole(a)"); len(got) != 1 {
		t.Errorf("arguments for whole(a): got %v, want one", got)
	}
	if got := argumentsFor(ag, "part(a)"); len(got) != 1 {
		t.Errorf("arguments for part(a): got %v, want one", got)
	}
	if got := argumentsFor(ag, "noted(a)"); len(got) != 1 {
		t.Errorf("the scheme listed last did not fire: arguments for noted(a): %v", got)
	}
}
