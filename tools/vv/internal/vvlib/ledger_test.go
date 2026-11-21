package vvlib

import "testing"

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
