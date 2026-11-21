package vvlib

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// The ledger (known_issues.yml) lists the checks that are known to fail today
// and the issue each failure belongs to. It turns the suite into a ratchet:
// a failure that is not in the ledger is a regression, and a ledger entry that
// stops failing has to be removed. Nothing in the code stops a later change
// from adding entries or removing checks; min_checks makes the second visible,
// and both need the same review as a change to the engine.

type Issue struct {
	Id       string `yaml:"id"`
	Where    string `yaml:"where"`    // engine, library or both
	Severity string `yaml:"severity"` // blocker, major, minor
	Title    string `yaml:"title"`
}

type Expected struct {
	Issue  string   `yaml:"issue"`
	File   string   `yaml:"file"`   // "" = every file that contains the scheme
	Scheme string   `yaml:"scheme"` // scheme id, "(file)" or "(engine)"
	Checks []string `yaml:"checks"` // exact names; a trailing * matches a prefix
	// Fingerprint, when set, is a text the failure's detail must contain: a
	// listed check that fails for another reason is reported as a new failure
	// (PLAN.md 0.3).
	Fingerprint string `yaml:"fingerprint"`
}

type Ledger struct {
	// MinChecks is the number of checks a complete run has to produce. A run
	// with fewer means that a check was removed or skipped.
	MinChecks int        `yaml:"min_checks"`
	Issues    []Issue    `yaml:"issues"`
	Expected  []Expected `yaml:"expected"`
	used      map[int]bool
}

// ResultKeys returns the complete, sorted inventory. Repeated identities are
// invalid even when their outcomes agree: otherwise a count can hide a missing
// check, or a map can silently overwrite a different outcome.
func ResultKeys(results []Result) ([]string, error) {
	keys := make([]string, len(results))
	for i, r := range results {
		keys[i] = r.Key()
	}
	return uniqueCheckKeys(keys)
}

// ParseExpectedChecks validates the one-key-per-line manifest, including before
// it is regenerated. An empty manifest is allowed for initial bootstrapping.
func ParseExpectedChecks(text string) ([]string, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return []string{}, nil
	}
	keys := strings.Split(text, "\n")
	for i := range keys {
		keys[i] = strings.TrimSuffix(keys[i], "\r")
	}
	return uniqueCheckKeys(keys)
}

func uniqueCheckKeys(keys []string) ([]string, error) {
	seen := map[string]int{}
	for i, key := range keys {
		if first, ok := seen[key]; ok {
			return nil, fmt.Errorf("duplicate check key %q (entries %d and %d)", key, first+1, i+1)
		}
		seen[key] = i
	}
	sort.Strings(keys)
	return keys, nil
}

func LoadLedger(path string) (*Ledger, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	l := &Ledger{used: map[int]bool{}}
	if err := yaml.UnmarshalStrict(b, l); err != nil {
		return nil, err
	}
	ids := map[string]bool{}
	for _, i := range l.Issues {
		ids[i.Id] = true
	}
	for n, e := range l.Expected {
		if !ids[e.Issue] {
			return nil, fmt.Errorf("expected entry %d refers to unknown issue %q", n+1, e.Issue)
		}
	}
	return l, nil
}

func match(pattern, s string) bool {
	if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix(s, strings.TrimSuffix(pattern, "*"))
	}
	return pattern == s
}

// Matches reports whether a ledger entry covers a result. The scheme "*"
// stands for every scheme of the file (not for "(file)" or "(engine)").
func (e Expected) Matches(r Result) bool {
	if e.File != "" && e.File != r.File {
		return false
	}
	if e.Scheme == "*" {
		if strings.HasPrefix(r.Scheme, "(") {
			return false
		}
	} else if e.Scheme != r.Scheme {
		return false
	}
	for _, c := range e.Checks {
		if match(c, r.Check) {
			return true
		}
	}
	return false
}

func (l *Ledger) find(r Result) int {
	for n, e := range l.Expected {
		if e.Matches(r) {
			return n
		}
	}
	return -1
}

// Lookup returns the id of the issue a failing result is filed under, or "".
func (l *Ledger) Lookup(r Result) string {
	if n := l.find(r); n >= 0 {
		return l.Expected[n].Issue
	}
	return ""
}

// Verdict sorts results into what the ledger expects and what it does not.
type Verdict struct {
	Passed int
	Known  []Result // fail, listed
	New    []Result // fail, not listed: a regression or an undiscovered defect
	Fixed  []Result // pass, but listed: the ledger is out of date
}

func (l *Ledger) Judge(results []Result) Verdict {
	v := Verdict{}
	for _, r := range results {
		n := l.find(r)
		switch {
		case r.Pass && n < 0:
			v.Passed++
		case !r.Pass && n >= 0 && l.Expected[n].Fingerprint != "" && !strings.Contains(r.Detail, l.Expected[n].Fingerprint):
			// listed, but failing for another reason than the one recorded
			r.Detail = fmt.Sprintf("fails for another reason than %q, the fingerprint in known_issues.yml: %s", l.Expected[n].Fingerprint, r.Detail)
			v.New = append(v.New, r)
		case !r.Pass && n >= 0:
			v.Known = append(v.Known, r)
		case !r.Pass:
			v.New = append(v.New, r)
		default:
			v.Fixed = append(v.Fixed, r)
		}
	}
	return v
}

// ByIssue groups known failures by issue id.
func (l *Ledger) ByIssue(known []Result) map[string][]Result {
	out := map[string][]Result{}
	for _, r := range known {
		id := l.Lookup(r)
		out[id] = append(out[id], r)
	}
	for _, rs := range out {
		sort.Slice(rs, func(i, j int) bool { return rs[i].Key() < rs[j].Key() })
	}
	return out
}
