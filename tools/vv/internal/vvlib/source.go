package vvlib

// The source ledger (source_ledger.yml) maps the entries of the compendium in
// Walton, Reed & Macagno (2008), ch. 9, to the schemes of walton.yml.

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// SourceEntry is one scheme of the compendium.
type SourceEntry struct {
	Entry     string   `yaml:"entry"`
	Page      int      `yaml:"page"`
	Questions int      `yaml:"questions"`
	Shared    bool     `yaml:"shared"`
	Schemes   []string `yaml:"schemes"`
	// Status is empty when the records listed are meant to be this entry's
	// scheme, and "other-content" when the record listed under the entry holds
	// a different scheme (the note says which). The ledger records where a
	// scheme is filed; it does not certify that a record says what the book says.
	Status string `yaml:"status"`
	Note   string `yaml:"note"`
}

type SourceLedger struct {
	Compendium   []SourceEntry `yaml:"compendium"`
	UpstreamOnly []string      `yaml:"upstream_only"`
	// Names records each predicate name the plan introduced, with its page
	// and how it was chosen (D6).
	Names []struct {
		Name  string `yaml:"name"`
		Page  int    `yaml:"page"`
		Entry string `yaml:"entry"`
		How   string `yaml:"how"`
		From  string `yaml:"from"`
	} `yaml:"names"`
	// Citations records citations of records that differ from the
	// compendium's, without changing them (PLAN.md 3.11).
	Citations []struct {
		Scheme     string `yaml:"scheme"`
		Record     string `yaml:"record"`
		Compendium string `yaml:"compendium"`
	} `yaml:"citations"`
	// Records has one row per record of walton.yml (PLAN.md 0.6).
	Records []RecordRow `yaml:"records"`
}

// RecordRow is the source ledger's row for one record (PLAN.md 0.6).
type RecordRow struct {
	Record    string           `yaml:"record"`
	Origin    string           `yaml:"origin"`
	Upstream  string           `yaml:"upstream"`
	Entry     string           `yaml:"entry"`
	Cites     string           `yaml:"cites"`
	Claim     string           `yaml:"claim"`
	Stage     string           `yaml:"stage"`
	Questions []RecordQuestion `yaml:"questions"`
	Changes   []RecordChange   `yaml:"changes"`
	Note      string           `yaml:"note"`
}

// RecordQuestion is one printed critical question, sorted (PLAN.md 5.1).
type RecordQuestion struct {
	N    int    `yaml:"n"`
	Kind string `yaml:"kind"` // premise, assumption, exception, rebuttal
	As   string `yaml:"as"`   // the statement that models it ("" for premise)
	Note string `yaml:"note"`
}

// RecordChange is one change made to a record, with the claim it makes.
type RecordChange struct {
	Plan  string `yaml:"plan"`
	Page  int    `yaml:"page"`
	Claim string `yaml:"claim"`
}

// The kinds of claim a row or a change makes, and the stages of a row in
// order (PLAN.md 0.6).
var (
	Claims = map[string]bool{"transcription": true, "correction": true, "interpretation": true, "policy": true, "unchecked": true}
	Stages = []string{"mapped", "transcription-checked", "interpretation-approved", "questions-sorted", "example-passed"}
	// QuestionKinds are the four sorts of PLAN.md 5.1.
	QuestionKinds = map[string]bool{"premise": true, "assumption": true, "exception": true, "rebuttal": true}
)

// StageIndex is the position of a stage in Stages, or -1.
func StageIndex(s string) int {
	for i, x := range Stages {
		if x == s {
			return i
		}
	}
	return -1
}

// CompendiumSource is how walton.yml cites the compendium itself.
const CompendiumSource = "Walton, Douglas and Reed, Chris and Macagno, Fabrizio (2008)"

func LoadSourceLedger(path string) (*SourceLedger, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	l := &SourceLedger{}
	if err := yaml.UnmarshalStrict(b, l); err != nil {
		return nil, err
	}
	for _, e := range l.Compendium {
		if e.Status != "" && e.Status != OtherContent {
			return nil, fmt.Errorf("entry %s: unknown status %q", e.Entry, e.Status)
		}
	}
	return l, nil
}

// OtherContent marks an entry whose record holds a different scheme.
const OtherContent = "other-content"

const sourceFile = "source_ledger.yml"

// SourceChecks compares the source ledger with the master scheme file: every
// scheme has exactly one place in the ledger, every id in the ledger exists,
// and every entry of the compendium has a record that is meant to be that
// entry. This is a check of the inventory. Whether a record's premises,
// questions and conclusion say what the book says is not checked here.
func SourceChecks(l *SourceLedger, lib *Library) []Result {
	out := []Result{}
	where := map[string][]string{}
	for _, e := range l.Compendium {
		for _, id := range e.Schemes {
			where[id] = append(where[id], "entry "+e.Entry)
		}
	}
	for _, id := range l.UpstreamOnly {
		where[id] = append(where[id], "upstream_only")
	}
	inLib := map[string]bool{}
	for _, s := range lib.Schemes() {
		inLib[s.Id] = true
		switch n := len(where[s.Id]); {
		case n == 1:
			out = append(out, pass(lib.Rel, s.Id, "source/mapped"))
		case n == 0:
			out = append(out, fail(lib.Rel, s.Id, "source/mapped", "the scheme has no place in %s: add it to the compendium entry it encodes, or to upstream_only", sourceFile))
		default:
			out = append(out, fail(lib.Rel, s.Id, "source/mapped", "the scheme is listed %d times in %s: %v", n, sourceFile, where[s.Id]))
		}
	}
	ids := []string{}
	for id := range where {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		if !inLib[id] {
			out = append(out, fail(sourceFile, id, "source/scheme-exists", "%s names a scheme that %s does not contain", sourceFile, lib.Rel))
		}
	}
	// every recorded name is a predicate of the library's language, with a
	// page and a way it was chosen
	for _, n := range l.Names {
		_, ok := lib.AG.Theory.Language[n.Name]
		switch {
		case !ok:
			out = append(out, fail(sourceFile, n.Name, "source/name-in-language", "%s records %s, which the language of %s does not define", sourceFile, n.Name, lib.Rel))
		case n.Page == 0 || (n.How != "reused" && n.How != "trial" && n.How != "book"):
			out = append(out, fail(sourceFile, n.Name, "source/name-in-language", "the row needs a page and how (reused, trial or book)"))
		default:
			out = append(out, pass(sourceFile, n.Name, "source/name-in-language"))
		}
	}
	for _, c := range l.Citations {
		if !inLib[c.Scheme] {
			out = append(out, fail(sourceFile, c.Scheme, "source/scheme-exists", "a citation row names a scheme that %s does not contain", lib.Rel))
		}
	}
	out = append(out, RecordChecks(l, lib)...)
	for _, e := range l.Compendium {
		name := "entry-" + e.Entry
		switch {
		case len(e.Schemes) == 0:
			out = append(out, fail(sourceFile, name, "source/has-record", "compendium entry %s (p. %d) has no record in %s", e.Entry, e.Page, lib.Rel))
		case e.Status == OtherContent:
			out = append(out, fail(sourceFile, name, "source/has-record", "compendium entry %s (p. %d) is filed under %v, a record with other content, so the scheme the compendium prints has no record in %s", e.Entry, e.Page, e.Schemes, lib.Rel))
		default:
			out = append(out, pass(sourceFile, name, "source/has-record"))
		}
	}
	return out
}

// SourceRow compares what the compendium prints for an entry with what the
// scheme models.
type SourceRow struct {
	Entry       string
	Page        int
	Questions   int
	Shared      bool
	Scheme      string
	Assumptions int
	Exceptions  int
	Status      string
	Note        string
	// from the record's row (PLAN.md 0.6)
	Origin, Claim, Stage string
	Sorted               int // questions sorted
}

// Fields is the number of assumption and exception statements the record has.
// It is not the number of critical questions the record answers: one question
// can need several statements, or none when it only restates a premise.
func (r SourceRow) Fields() int { return r.Assumptions + r.Exceptions }

// SourceRows lists one row per scheme, in the order of the compendium, followed
// by the upstream-only records (entry "-").
func SourceRows(l *SourceLedger, lib *Library) []SourceRow {
	rows := []SourceRow{}
	defer func() {
		byId := map[string]RecordRow{}
		for _, r := range l.Records {
			byId[r.Record] = r
		}
		for i := range rows {
			if r, ok := byId[rows[i].Scheme]; ok {
				rows[i].Origin, rows[i].Claim, rows[i].Stage, rows[i].Sorted = r.Origin, r.Claim, r.Stage, len(r.Questions)
			}
		}
	}()
	count := func(id string) (a, e int) {
		if s := lib.Scheme(id); s != nil {
			return len(s.Assumptions), len(s.Exceptions)
		}
		return 0, 0
	}
	for _, e := range l.Compendium {
		if len(e.Schemes) == 0 {
			rows = append(rows, SourceRow{Entry: e.Entry, Page: e.Page, Questions: e.Questions, Shared: e.Shared, Status: "no-record", Note: e.Note})
		}
		for _, id := range e.Schemes {
			a, x := count(id)
			st := "mapped"
			if e.Status != "" {
				st = e.Status
			}
			rows = append(rows, SourceRow{Entry: e.Entry, Page: e.Page, Questions: e.Questions, Shared: e.Shared, Scheme: id, Assumptions: a, Exceptions: x, Status: st, Note: e.Note})
		}
	}
	for _, id := range l.UpstreamOnly {
		a, x := count(id)
		rows = append(rows, SourceRow{Entry: "-", Scheme: id, Assumptions: a, Exceptions: x, Status: "upstream-only", Note: "from Carneades' own walton.yml; no entry of its own in the compendium"})
	}
	return rows
}

// SourceSummary gives the figures the plan uses to size the content review.
func SourceSummary(rows []SourceRow) string {
	entries, records, none, noneBoth, noneBook, fewer, same, more := map[string]bool{}, 0, 0, 0, 0, 0, 0, 0
	missing, other := []string{}, []string{}
	for _, r := range rows {
		if r.Entry != "-" {
			entries[r.Entry] = true
		}
		if r.Scheme == "" {
			missing = append(missing, r.Entry)
			continue
		}
		if r.Status == OtherContent {
			other = append(other, r.Entry)
		}
		records++
		m := r.Fields()
		if m == 0 {
			none++
		}
		switch {
		case r.Entry == "-":
		case m == 0 && r.Questions == 0:
			noneBoth++
		case m == 0:
			noneBook++
		case m < r.Questions:
			fewer++
		case m == r.Questions:
			same++
		default:
			more++
		}
	}
	return fmt.Sprintf("%d compendium entries, %d records; entries without a record: %v; filed under a record with other content: %v\n"+
		"records with no assumption or exception: %d; the book prints no question for %d of them and some for %d\n"+
		"records with fewer assumptions and exceptions than the book prints questions: %d; as many: %d; more: %d\n"+
		"(a count of statements, to size the review; it does not say which questions a record answers)",
		len(entries), records, missing, other, none, noneBoth, noneBook, fewer, same, more)
}

// RecordChecks checks the rows of the source ledger, one per record
// (PLAN.md 0.6): every record has one row and every row a record; the entry
// agrees with the compendium and the citation with walton.yml; the upstream
// rows are the records upstream ships; a row's claim and stage are known, and
// its stage has what it needs.
func RecordChecks(l *SourceLedger, lib *Library) []Result {
	out := []Result{}
	entryOf := map[string]string{}
	entries := map[string]SourceEntry{}
	for _, e := range l.Compendium {
		entries[e.Entry] = e
		for _, id := range e.Schemes {
			entryOf[id] = e.Entry
		}
	}
	upstream := map[string]bool{}
	if b, err := os.ReadFile(filepath.Join(RepoRoot(), "third_party", "carneades-4", "examples", "AGs", "YAML", "walton.yml")); err == nil {
		var f struct {
			Schemes []struct {
				Id string `yaml:"id"`
			} `yaml:"argument_schemes"`
		}
		if yaml.Unmarshal(b, &f) == nil {
			for _, s := range f.Schemes {
				upstream[s.Id] = true
			}
		}
	}
	cases := map[string]bool{}
	for _, f := range CaseFiles() {
		cs, _ := LoadCases(f)
		for _, c := range cs {
			cases[c.Name] = true
		}
	}
	rows := map[string][]RecordRow{}
	for _, r := range l.Records {
		rows[r.Record] = append(rows[r.Record], r)
	}
	seenUpstream := map[string]string{}
	for _, s := range lib.Schemes() {
		rs := rows[s.Id]
		if len(rs) != 1 {
			out = append(out, fail(sourceFile, s.Id, "source/record-row", "%d rows for the record in records:, one expected", len(rs)))
			continue
		}
		r := rs[0]
		probs := []string{}
		switch r.Origin {
		case "upstream":
			u := r.Upstream
			if u == "" {
				u = r.Record
			}
			if !upstream[u] {
				probs = append(probs, fmt.Sprintf("origin upstream, but upstream's walton.yml has no record %s", u))
			} else if prev, dup := seenUpstream[u]; dup {
				probs = append(probs, fmt.Sprintf("upstream record %s is also claimed by %s", u, prev))
			}
			seenUpstream[u] = r.Record
		case "book":
			if r.Upstream != "" {
				probs = append(probs, "a book record has no upstream id")
			}
			if upstream[r.Record] {
				probs = append(probs, "upstream's walton.yml has a record with this id: origin upstream")
			}
			if r.Entry == "" {
				probs = append(probs, "a book record needs its entry")
			}
		default:
			probs = append(probs, fmt.Sprintf("origin %q is neither book nor upstream", r.Origin))
		}
		if r.Entry != entryOf[s.Id] {
			probs = append(probs, fmt.Sprintf("entry %q, but the compendium list files the record under %q", r.Entry, entryOf[s.Id]))
		}
		src := ""
		if v, ok := s.Metadata["source"]; ok && v != nil {
			src = strings.Join(strings.Fields(fmt.Sprint(v)), " ")
		}
		want := ""
		if src != "" && !strings.HasPrefix(src, CompendiumSource) {
			want = src
		}
		if r.Cites != want {
			probs = append(probs, fmt.Sprintf("cites %q, but the record cites %q", r.Cites, want))
		}
		if !Claims[r.Claim] {
			probs = append(probs, fmt.Sprintf("unknown claim %q", r.Claim))
		}
		for _, c := range r.Changes {
			if c.Plan == "" || c.Page == 0 || !Claims[c.Claim] || c.Claim == "unchecked" {
				probs = append(probs, fmt.Sprintf("change %+v needs a plan item, a page and a claim", c))
			}
		}
		stage := StageIndex(r.Stage)
		if stage < 0 {
			probs = append(probs, fmt.Sprintf("unknown stage %q", r.Stage))
		}
		if stage >= 1 && r.Claim == "unchecked" {
			probs = append(probs, "a row past mapped says what the record claims to be")
		}
		for _, q := range r.Questions {
			if !QuestionKinds[q.Kind] {
				probs = append(probs, fmt.Sprintf("question %d: unknown kind %q", q.N, q.Kind))
			} else if q.Kind != "premise" && q.As == "" {
				probs = append(probs, fmt.Sprintf("question %d is an %s: name the statement that models it", q.N, q.Kind))
			}
		}
		if stage >= StageIndex("questions-sorted") && r.Origin == "book" {
			if n := entries[r.Entry].Questions; len(r.Questions) != n {
				probs = append(probs, fmt.Sprintf("questions sorted: %d of the %d the compendium prints for entry %s", len(r.Questions), n, r.Entry))
			}
		}
		if stage >= StageIndex("example-passed") && !cases[r.Record+"/book"] {
			probs = append(probs, "example passed, but no reference case "+r.Record+"/book")
		}
		out = append(out, res(sourceFile, s.Id, "source/record-row", probs))
	}
	extra := []string{}
	for id := range rows {
		if lib.Scheme(id) == nil {
			extra = append(extra, id)
		}
	}
	missing := []string{}
	for u := range upstream {
		if _, ok := seenUpstream[u]; !ok {
			missing = append(missing, u)
		}
	}
	sort.Strings(extra)
	sort.Strings(missing)
	probs := []string{}
	if len(extra) > 0 {
		probs = append(probs, fmt.Sprintf("rows for records %s does not contain: %v", lib.Rel, extra))
	}
	if len(upstream) == 0 {
		probs = append(probs, "upstream's walton.yml cannot be read")
	}
	if len(missing) > 0 {
		probs = append(probs, fmt.Sprintf("upstream records with no row: %v", missing))
	}
	out = append(out, res(sourceFile, "(file)", "source/records-complete", probs))
	return out
}
