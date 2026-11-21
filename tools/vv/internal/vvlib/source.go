package vvlib

// The source ledger (source_ledger.yml) maps the entries of the compendium in
// Walton, Reed & Macagno (2008), ch. 9, to the schemes of walton.yml.

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carneades/carneades-4/src/engine/caes"
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
	// QuestionKinds are the four sorts of PLAN.md 5.1, and standard: a
	// question about the burden or the standard of proof, which Carneades
	// sets for an issue in the graph, not in a scheme.
	QuestionKinds = map[string]bool{"premise": true, "assumption": true, "exception": true, "rebuttal": true, "standard": true}
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
		// PLAN.md 5.9: every record has a title and a source; a record filed
		// under an entry names the compendium, the entry and its page first,
		// and the row's cites holds the rest of its source.
		src, title := "", ""
		if v, ok := s.Metadata["source"]; ok && v != nil {
			src = strings.Join(strings.Fields(fmt.Sprint(v)), " ")
		}
		if v, ok := s.Metadata["title"]; ok && v != nil {
			title = strings.TrimSpace(fmt.Sprint(v))
		}
		if title == "" {
			probs = append(probs, "the record has no meta.title")
		}
		want := src
		if r.Entry != "" {
			std := fmt.Sprintf("%s. Argumentation Schemes. Cambridge University Press, ch. 9, entry %s, p. %d.", CompendiumSource, r.Entry, entries[r.Entry].Page)
			if !strings.HasPrefix(src, std) {
				probs = append(probs, fmt.Sprintf("meta.source should begin with %q", std))
			}
			want = strings.TrimSpace(strings.TrimPrefix(src, std))
		} else if src == "" {
			probs = append(probs, "the record has no meta.source")
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
		probs = append(probs, questionProblems(r, s, entries[r.Entry].Questions, stage >= StageIndex("questions-sorted"))...)
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

// questionProblems checks a row's sorted questions (PLAN.md 5.1) against its
// record: an assumption is one of the record's assumptions, an exception one
// of its exceptions, a premise (when it names one) one of its premises, and a
// rebuttal the negation of one of its conclusions, weighed against it in the
// negation issue. Numbers run from 1 to the count the compendium prints; 0
// marks a statement the record models although the book prints no question
// for it. Once the row is at questions-sorted, every printed question is
// sorted and every assumption and exception of the record is accounted for.
func questionProblems(r RecordRow, s *caes.Scheme, printed int, complete bool) []string {
	probs := []string{}
	in := func(list []string, st string) bool {
		for _, x := range list {
			if Norm(x) == Norm(st) {
				return true
			}
		}
		return false
	}
	negations := []string{}
	for _, c := range s.Conclusions {
		n := Norm(c)
		if strings.HasPrefix(n, "¬") {
			negations = append(negations, strings.TrimPrefix(n, "¬"))
		} else {
			negations = append(negations, "¬"+n)
		}
	}
	seen := map[int]bool{}
	named := map[string]bool{}
	for _, q := range r.Questions {
		switch {
		case q.N < 0 || (printed >= 0 && r.Origin == "book" && q.N > printed):
			probs = append(probs, fmt.Sprintf("question %d: the compendium prints %d", q.N, printed))
		case q.N > 0 && seen[q.N]:
			probs = append(probs, fmt.Sprintf("question %d is sorted twice", q.N))
		}
		seen[q.N] = true
		ok := true
		if q.As == "" && q.Kind != "premise" && q.Kind != "standard" {
			// sorted, not modelled yet (a gap of PLAN.md 5.2)
			if complete {
				probs = append(probs, fmt.Sprintf("question %d, %s, is not modelled", q.N, article(q.Kind)))
			}
			seen[q.N] = true
			continue
		}
		switch q.Kind {
		case "premise":
			ok = q.As == "" || in(s.Premises, q.As)
		case "assumption":
			ok = in(s.Assumptions, q.As)
		case "exception":
			ok = in(s.Exceptions, q.As)
		case "rebuttal":
			ok = in(negations, q.As)
		case "standard":
			ok = q.As == ""
		default:
			probs = append(probs, fmt.Sprintf("question %d: unknown kind %q", q.N, q.Kind))
			continue
		}
		if !ok {
			probs = append(probs, fmt.Sprintf("question %d, %s, is modelled by %q, which is not one of the record's %ss", q.N, article(q.Kind), q.As, map[string]string{"premise": "premise", "assumption": "assumption", "exception": "exception", "rebuttal": "negated conclusion"}[q.Kind]))
		}
		if q.As != "" {
			named[Norm(q.As)] = true
		}
	}
	if complete {
		for n := 1; n <= printed; n++ {
			if !seen[n] {
				probs = append(probs, fmt.Sprintf("question %d of %d is not sorted", n, printed))
			}
		}
		for _, st := range append(append([]string{}, s.Assumptions...), s.Exceptions...) {
			if !named[Norm(st)] {
				probs = append(probs, fmt.Sprintf("%s is modelled but no question names it (use n: 0 for a statement the book prints no question for)", st))
			}
		}
	}
	return probs
}

// article puts "a" or "an" before a kind of question.
func article(kind string) string {
	if strings.ContainsAny(kind[:1], "aeiou") {
		return "an " + kind
	}
	return "a " + kind
}

// MigrationChecks compares tools/migration_r2.yml, the old-to-new table of
// release R2 (PLAN.md D4, 5.6), with the library: a row's new statement is in
// its record, and its old statement, where the row replaces or removes one,
// is not.
func MigrationChecks(lib *Library) []Result {
	const file = "tools/migration_r2.yml"
	b, err := os.ReadFile(filepath.Join(RepoRoot(), filepath.FromSlash(file)))
	if err != nil {
		return []Result{fail(file, "(file)", "migration/readable", "%v", err)}
	}
	var t struct {
		Changes []struct {
			Record string `yaml:"record"`
			Kind   string `yaml:"kind"`
			Old    string `yaml:"old"`
			New    string `yaml:"new"`
			Plan   string `yaml:"plan"`
			Page   int    `yaml:"page"`
			Note   string `yaml:"note"`
		} `yaml:"changes"`
	}
	if err := yaml.UnmarshalStrict(b, &t); err != nil {
		return []Result{fail(file, "(file)", "migration/readable", "%v", err)}
	}
	out := []Result{pass(file, "(file)", "migration/readable")}
	for _, c := range t.Changes {
		probs := []string{}
		s := lib.Scheme(c.Record)
		if s == nil {
			probs = append(probs, "no such record")
		} else {
			all := map[string]bool{}
			for _, l := range [][]string{s.Premises, s.Assumptions, s.Exceptions, s.Conclusions} {
				for _, st := range l {
					all[Norm(st)] = true
				}
			}
			switch c.Kind {
			case "added", "premise":
				if c.Old != "" || c.New == "" || !all[Norm(c.New)] {
					probs = append(probs, fmt.Sprintf("added: %q should be in the record, with no old statement", c.New))
				}
			case "removed":
				if c.New != "" || c.Old == "" || all[Norm(c.Old)] {
					probs = append(probs, fmt.Sprintf("removed: %q should no longer be in the record", c.Old))
				}
			case "arguments", "arity", "replaced":
				if c.Old == "" || c.New == "" || all[Norm(c.Old)] || !all[Norm(c.New)] {
					probs = append(probs, fmt.Sprintf("%s: %q should have replaced %q in the record", c.Kind, c.New, c.Old))
				}
			default:
				probs = append(probs, fmt.Sprintf("unknown kind %q", c.Kind))
			}
		}
		if c.Plan == "" || c.Page == 0 {
			probs = append(probs, "a row needs its plan item and page")
		}
		st := c.New
		if st == "" {
			st = c.Old
		}
		out = append(out, res(file, c.Record, "migration/"+c.Kind+" "+st, probs))
	}
	return out
}
