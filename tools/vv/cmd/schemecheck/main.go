// schemecheck runs every scheme of the repository's scheme files through the
// Carneades 4 engine and prints what it finds.
//
//	go run ./cmd/schemecheck                     summary per file and every failing check
//	go run ./cmd/schemecheck -format json        all results, one JSON object per line
//	go run ./cmd/schemecheck -format md          matrix: one row per scheme of walton.yml
//	go run ./cmd/schemecheck -format source      CSV: each scheme next to its entry in the book's compendium
//	go run ./cmd/schemecheck -file walton.yml    one file only
//
// Exit status 1 if a check fails that is not in the ledger of known issues
// (known_issues.yml), or if a ledger entry no longer fails.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"waltonvv/internal/vvlib"
)

func main() {
	repo := flag.String("repo", vvlib.RepoRoot(), "root of the scheme repository")
	only := flag.String("file", "", "check only this file (repository-relative)")
	format := flag.String("format", "text", "text, json, md or source")
	ledgerPath := flag.String("ledger", "known_issues.yml", "ledger of known issues ('' for none)")
	sourcePath := flag.String("source", "source_ledger.yml", "map from the book's compendium to the schemes")
	flag.Parse()

	if *format == "source" {
		source(*repo, *sourcePath)
		return
	}

	var ledger *vvlib.Ledger
	if *ledgerPath != "" {
		var err error
		ledger, err = vvlib.LoadLedger(*ledgerPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ledger:", err)
			os.Exit(2)
		}
	}

	files := vvlib.SchemeFiles(*repo)
	if *only != "" {
		files = []string{*only}
	}
	all := []vvlib.Result{}
	libs := map[string]*vvlib.Library{}
	for _, f := range files {
		lib, err := vvlib.LoadLibrary(*repo, f)
		if err != nil {
			all = append(all, vvlib.Result{File: f, Scheme: "(file)", Check: "import", Detail: err.Error()})
			continue
		}
		all = append(all, vvlib.Result{File: f, Scheme: "(file)", Check: "import", Pass: true})
		libs[f] = lib
		all = append(all, vvlib.CheckLibrary(lib)...)
	}

	switch *format {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		for _, r := range all {
			enc.Encode(r)
		}
	case "md":
		markdown(all, libs, files)
	default:
		text(all, libs, files, ledger)
	}

	if ledger != nil {
		v := ledger.Judge(all)
		if *format == "text" {
			fmt.Printf("\nledger: %d passed, %d known failures, %d new failures, %d ledger entries that no longer fail\n",
				v.Passed, len(v.Known), len(v.New), len(v.Fixed))
			for _, r := range v.New {
				fmt.Printf("  NEW    %s: %s\n", r.Key(), r.Detail)
			}
			for _, r := range v.Fixed {
				fmt.Printf("  FIXED  %s (remove it from the ledger)\n", r.Key())
			}
		}
		if len(v.New)+len(v.Fixed) > 0 {
			os.Exit(1)
		}
	}
}

// source prints, as CSV, every scheme of walton.yml next to its entry in the
// compendium: the questions the book prints and the ones the scheme models.
func source(repo, path string) {
	l, err := vvlib.LoadSourceLedger(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "source ledger:", err)
		os.Exit(2)
	}
	lib, err := vvlib.LoadLibrary(repo, "walton.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "walton.yml:", err)
		os.Exit(2)
	}
	rows := vvlib.SourceRows(l, lib)
	fmt.Println("entry,page,questions_in_book,list_shared_with_siblings,scheme,assumptions,exceptions,fields,status,note")
	for _, r := range rows {
		page := ""
		if r.Page > 0 {
			page = fmt.Sprint(r.Page)
		}
		fmt.Printf("%s,%s,%d,%v,%s,%d,%d,%d,%s,%q\n", r.Entry, page, r.Questions, r.Shared, r.Scheme, r.Assumptions, r.Exceptions, r.Fields(), r.Status, r.Note)
	}
	fmt.Fprintln(os.Stderr, vvlib.SourceSummary(rows))
}

func text(all []vvlib.Result, libs map[string]*vvlib.Library, files []string, ledger *vvlib.Ledger) {
	for _, f := range files {
		schemes := map[string]bool{}
		failing := map[string]bool{}
		n, bad := 0, 0
		for _, r := range all {
			if r.File != f {
				continue
			}
			n++
			if r.Scheme != "(file)" {
				schemes[r.Scheme] = true
			}
			if !r.Pass {
				bad++
				if r.Scheme != "(file)" {
					failing[r.Scheme] = true
				}
			}
		}
		fmt.Printf("%s: %d schemes, %d checks, %d failing checks in %d schemes\n", f, len(schemes), n, bad, len(failing))
	}
	fmt.Println()
	for _, r := range all {
		if r.Pass {
			continue
		}
		tag := ""
		if ledger != nil {
			if id := ledger.Lookup(r); id != "" {
				tag = " [" + id + "]"
			}
		}
		fmt.Printf("FAIL %s%s\n     %s\n", r.Key(), tag, r.Detail)
	}
	if lib, ok := libs["walton.yml"]; ok {
		recursion(lib)
	}
}

func recursion(lib *vvlib.Library) {
	ids := []string{}
	for _, s := range lib.Schemes() {
		ids = append(ids, s.Id)
	}
	feeds := vvlib.Feeds(lib.Schemes())
	fmt.Println("\nSchemes that can re-trigger themselves or each other (cycles of the feeds graph, walton.yml):")
	for _, c := range vvlib.Cycles(feeds, ids) {
		fmt.Printf("  %s\n", strings.Join(c, " <-> "))
	}
	ca, pa := vvlib.Wildcards(lib.Schemes())
	fmt.Printf("Schemes that conclude a bare statement variable (can feed any scheme): %s\n", strings.Join(ca, ", "))
	fmt.Printf("Schemes with a bare statement variable as a premise (fed by any statement): %s\n", strings.Join(pa, ", "))
}

// markdown prints one row per scheme of walton.yml with the outcome of the
// checks that decide whether the scheme is usable.
func markdown(all []vvlib.Result, libs map[string]*vvlib.Library, files []string) {
	cols := []string{"engine-validator", "chr-translation", "iso/fires", "iso/ground", "iso/conclusions-in", "iso/terminates", "full/terminates", "iso/exceptions-undercut", "full/exceptions-undercut", "iso/premises-needed", "full/premises-needed", "iso/texts"}
	head := []string{"validator", "rule", "fires", "ground", "concl. in", "stops", "stops (full)", "CQs", "CQs (full)", "premises", "premises (full)", "texts"}
	for _, f := range files {
		lib, ok := libs[f]
		if !ok {
			continue
		}
		by := map[string]map[string]vvlib.Result{}
		for _, r := range all {
			if r.File == f {
				if by[r.Scheme] == nil {
					by[r.Scheme] = map[string]vvlib.Result{}
				}
				by[r.Scheme][r.Check] = r
			}
		}
		fmt.Printf("### %s\n\n| scheme | %s |\n|---|%s\n", f, strings.Join(head, " | "), strings.Repeat("---|", len(head)))
		ids := []string{}
		for _, s := range lib.Schemes() {
			ids = append(ids, s.Id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			row := []string{}
			for _, c := range cols {
				r, ok := by[id][c]
				switch {
				case !ok:
					row = append(row, "–")
				case r.Pass:
					row = append(row, "ok")
				default:
					row = append(row, "**FAIL**")
				}
			}
			fmt.Printf("| `%s` | %s |\n", id, strings.Join(row, " | "))
		}
		fmt.Println()
	}
}
