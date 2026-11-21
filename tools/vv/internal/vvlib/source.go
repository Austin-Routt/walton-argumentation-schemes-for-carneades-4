package vvlib

// The source ledger (source_ledger.yml) maps the entries of the compendium in
// Walton, Reed & Macagno (2008), ch. 9, to the schemes of walton.yml.

import (
	"fmt"
	"os"
	"sort"

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
}

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
}

// Fields is the number of assumption and exception statements the record has.
// It is not the number of critical questions the record answers: one question
// can need several statements, or none when it only restates a premise.
func (r SourceRow) Fields() int { return r.Assumptions + r.Exceptions }

// SourceRows lists one row per scheme, in the order of the compendium, followed
// by the upstream-only records (entry "-").
func SourceRows(l *SourceLedger, lib *Library) []SourceRow {
	rows := []SourceRow{}
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
