package vvlib

import (
	"reflect"
	"strings"
	"testing"
)

// PLAN.md 0.3: a listed check that fails for another reason than the one its
// fingerprint records is a new failure, not a known one.
func TestFingerprint(t *testing.T) {
	l := &Ledger{
		Issues: []Issue{{Id: "X"}},
		Expected: []Expected{
			{Issue: "X", File: "f.yml", Scheme: "s", Checks: []string{"c"}, Fingerprint: "number of parameters"},
			{Issue: "X", File: "f.yml", Scheme: "t", Checks: []string{"c"}},
		},
	}
	v := l.Judge([]Result{
		{File: "f.yml", Scheme: "s", Check: "c", Detail: "wrong number of parameters"},
		{File: "f.yml", Scheme: "s", Check: "c", Detail: "timed out"},
		{File: "f.yml", Scheme: "t", Check: "c", Detail: "anything"},
	})
	if len(v.Known) != 2 || len(v.New) != 1 || v.New[0].Scheme != "s" {
		t.Fatalf("known %d, new %d: %+v", len(v.Known), len(v.New), v.New)
	}
}

func TestResultKeysRejectRepeatedIdentityRegardlessOfOutcome(t *testing.T) {
	first := Result{File: "f.yml", Scheme: "s", Check: "c", Pass: true}
	for _, outcome := range []bool{false, true} {
		second := first
		second.Pass = outcome
		second.Detail = "a different detail does not give a check another identity"
		keys, err := ResultKeys([]Result{first, second})
		if err == nil || keys != nil || !strings.Contains(err.Error(), "duplicate check key") {
			t.Fatalf("outcome %v accepted duplicate or partial inventory: keys=%v, err=%v", outcome, keys, err)
		}
	}
}

func TestResultKeysRetainDistinctFileSchemeAndCheck(t *testing.T) {
	rows := []Result{
		{File: "z.yml", Scheme: "s", Check: "c"},
		{File: "f.yml", Scheme: "s", Check: "c"},
		{File: "f.yml", Scheme: "other", Check: "c"},
		{File: "f.yml", Scheme: "s", Check: "other"},
	}
	keys, err := ResultKeys(rows)
	want := []string{"f.yml|other|c", "f.yml|s|c", "f.yml|s|other", "z.yml|s|c"}
	if err != nil || !reflect.DeepEqual(keys, want) || rows[0].File != "z.yml" {
		t.Fatalf("distinct keys lost, not sorted, or results reordered: %v, %v", keys, err)
	}
}

func TestExpectedChecksRejectDuplicateEntries(t *testing.T) {
	for _, text := range []string{
		"f|s|c\nf|s|c\n",
		"f|s|c\nf|other|c\nf|s|c\n",
		"f|s|c\r\nf|s|c\r\n",
	} {
		keys, err := ParseExpectedChecks(text)
		if err == nil || keys != nil || !strings.Contains(err.Error(), "duplicate check key") {
			t.Fatalf("accepted duplicate manifest %q: %v, %v", text, keys, err)
		}
	}
}

func TestExpectedChecksAcceptCompleteUnorderedInventory(t *testing.T) {
	keys, err := ParseExpectedChecks("z|s|c\r\nf|s|c\r\n")
	if err != nil || !reflect.DeepEqual(keys, []string{"f|s|c", "z|s|c"}) {
		t.Fatalf("valid manifest: %v, %v", keys, err)
	}
	if keys, err := ParseExpectedChecks(""); err != nil || len(keys) != 0 {
		t.Fatalf("empty bootstrap manifest: %v, %v", keys, err)
	}
}
