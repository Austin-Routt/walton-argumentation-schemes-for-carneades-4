// Verification and validation of the Carneades 4 engine snapshot and the
// Walton scheme library, run as `go test` (see README.md).
//
// Every check states behaviour that practical use needs. Checks that are known
// to fail are listed, with the issue they belong to, in known_issues.yml:
//
//   - a listed check that fails is reported as KNOWN and does not fail the run;
//   - a check that fails without being listed fails the run (a regression, or
//     a defect nobody had seen yet);
//   - a listed check that passes fails the run too, so the ledger gets cleaned
//     up when something is fixed.
//
// VV_STRICT=1 turns every known failure into a test failure; that run is green
// only when the ledger is empty.
package waltonvv

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"

	"waltonvv/internal/vvlib"
)

var (
	ledger  *vvlib.Ledger
	mu      sync.Mutex
	collect []vvlib.Result
)

func TestMain(m *testing.M) {
	// A probe worker: run one contained check and leave (internal/vvlib/isolate.go).
	if p := os.Getenv(vvlib.ProbeEnv); p != "" {
		os.Exit(vvlib.RunProbeChild(p))
	}
	var err error
	ledger, err = vvlib.LoadLedger("known_issues.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "known_issues.yml:", err)
		os.Exit(2)
	}
	code := m.Run()
	// Also applies to focused runs that do not select TestLedgerIsUsed.
	if _, err := vvlib.ResultKeys(collect); err != nil {
		fmt.Fprintln(os.Stderr, "result inventory:", err)
		code = 1
	}
	summary()
	os.Exit(code)
}

// judge applies the ledger to a batch of results inside the current test.
func judge(t *testing.T, results []vvlib.Result) {
	t.Helper()
	mu.Lock()
	collect = append(collect, results...)
	mu.Unlock()
	if vvlib.Tainted() {
		// An inference run timed out inside this process. The engine cannot be
		// stopped from outside: it is still running and still using the output
		// streams and the engine's package-level state. Nothing measured after
		// this point could be trusted, so the run ends here, as a failure.
		summary()
		fmt.Fprintln(os.Stderr, "\nSTOPPED: an inference run did not return in time and cannot be cancelled; the checks after it were not run.")
		os.Exit(1)
	}
	strict := os.Getenv("VV_STRICT") != ""
	v := ledger.Judge(results)
	for _, r := range v.New {
		t.Errorf("FAIL  %s\n      %s", r.Key(), r.Detail)
	}
	for _, r := range v.Fixed {
		t.Errorf("FIXED %s passes but known_issues.yml still lists it under %s: remove the entry", r.Key(), ledger.Lookup(r))
	}
	for _, r := range v.Known {
		if strict {
			t.Errorf("KNOWN %s [%s]\n      %s", r.Key(), ledger.Lookup(r), r.Detail)
		} else if testing.Verbose() {
			t.Logf("known %s [%s]", r.Key(), ledger.Lookup(r))
		}
	}
}

func summary() {
	v := ledger.Judge(collect)
	fmt.Printf("\n=== V&V summary: %d checks: %d pass, %d known failures, %d unexpected failures, %d stale ledger entries ===\n",
		len(collect), v.Passed, len(v.Known), len(v.New), len(v.Fixed))
	by := ledger.ByIssue(v.Known)
	ids := []string{}
	for id := range by {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	title := map[string]vvlib.Issue{}
	for _, i := range ledger.Issues {
		title[i.Id] = i
	}
	for _, id := range ids {
		i := title[id]
		fmt.Printf("  %-10s %-8s %-7s %4d failing checks  %s\n", id, i.Where, i.Severity, len(by[id]), i.Title)
	}
	if out := os.Getenv("VV_RESULTS"); out != "" {
		f, err := os.Create(out)
		if err == nil {
			enc := json.NewEncoder(f)
			for _, r := range collect {
				enc.Encode(struct {
					vvlib.Result
					Issue string `json:"issue,omitempty"`
				}{r, func() string {
					if r.Pass {
						return ""
					}
					return ledger.Lookup(r)
				}()})
			}
			f.Close()
		}
	}
}

// TestEngine checks the engine through its Go API, independently of the scheme files.
func TestEngine(t *testing.T) {
	for _, r := range append(vvlib.EngineChecks(), vvlib.VersionChecks()...) {
		r := r
		t.Run(r.Check, func(t *testing.T) { judge(t, []vvlib.Result{r}) })
	}
}

// TestEngineToolchain builds, vets and tests the Carneades module as shipped.
// It needs the Go toolchain to resolve the engine's own dependencies; skip it
// with VV_SKIP_TOOLCHAIN=1 where that is not possible.
func TestEngineToolchain(t *testing.T) {
	if os.Getenv("VV_SKIP_TOOLCHAIN") != "" {
		t.Skip("VV_SKIP_TOOLCHAIN is set")
	}
	for _, r := range vvlib.ToolchainChecks() {
		r := r
		t.Run(r.Check, func(t *testing.T) { judge(t, []vvlib.Result{r}) })
	}
}

// TestLibrary runs every scheme of every scheme file through the engine.
func TestLibrary(t *testing.T) {
	root := vvlib.RepoRoot()
	for _, f := range vvlib.SchemeFiles(root) {
		f := f
		t.Run(f, func(t *testing.T) {
			lib, err := vvlib.LoadLibrary(root, f)
			if err != nil {
				judge(t, []vvlib.Result{{File: f, Scheme: "(file)", Check: "import", Detail: err.Error()}})
				return
			}
			judge(t, []vvlib.Result{{File: f, Scheme: "(file)", Check: "import", Pass: true}})
			by := map[string][]vvlib.Result{}
			order := []string{}
			for _, r := range vvlib.CheckLibrary(lib) {
				if _, seen := by[r.Scheme]; !seen {
					order = append(order, r.Scheme)
				}
				by[r.Scheme] = append(by[r.Scheme], r)
			}
			for _, s := range order {
				s := s
				t.Run(s, func(t *testing.T) { judge(t, by[s]) })
			}
		})
	}
}

// TestSource compares walton.yml with the source ledger: every scheme has one
// place in the book's compendium (or is marked as taken over from upstream),
// and every entry of the compendium has a record.
// TestCases runs the reference cases of tools/vv/cases: graphs written by
// hand against the library, with the arguments and labels each must produce.
func TestCases(t *testing.T) {
	root := vvlib.RepoRoot()
	cli, err := vvlib.FindCLI()
	if err != nil {
		t.Fatalf("no carneades command: %v", err)
	}
	for _, f := range vvlib.CaseFiles() {
		f := f
		t.Run(f, func(t *testing.T) {
			judge(t, vvlib.CaseChecks(root, f))
			// PLAN.md 4.5: evaluated by the command, written, and read back
			// in a new process
			judge(t, vvlib.CaseRoundTrips(cli, root, f))
		})
	}
}

// TestProperties runs the property tests of PLAN.md 4.5.
func TestProperties(t *testing.T) {
	judge(t, vvlib.PropertyChecks(vvlib.RepoRoot()))
}

func TestSource(t *testing.T) {
	l, err := vvlib.LoadSourceLedger("source_ledger.yml")
	if err != nil {
		t.Fatalf("source_ledger.yml: %v", err)
	}
	lib, err := vvlib.LoadLibrary(vvlib.RepoRoot(), "walton.yml")
	if err != nil {
		t.Skipf("walton.yml does not import (reported by TestLibrary): %v", err)
	}
	judge(t, vvlib.SourceChecks(l, lib))
	judge(t, vvlib.MigrationChecks(lib))
}

// TestCLI drives the carneades command, one new process per run.
func TestCLI(t *testing.T) {
	cli, err := vvlib.FindCLI()
	if err != nil {
		t.Fatalf("no carneades command: %v", err)
	}
	t.Run("command", func(t *testing.T) {
		for _, r := range vvlib.CLIEngineChecks(cli) {
			r := r
			t.Run(r.Check, func(t *testing.T) { judge(t, []vvlib.Result{r}) })
		}
	})
	t.Run("migration", func(t *testing.T) { judge(t, vvlib.MigrationFixtureChecks(cli)) })
	t.Run("docs", func(t *testing.T) { judge(t, vvlib.DocExampleChecks(cli)) })
	root := vvlib.RepoRoot()
	for _, f := range vvlib.SchemeFiles(root) {
		f := f
		t.Run(f, func(t *testing.T) {
			lib, err := vvlib.LoadLibrary(root, f)
			if err != nil {
				t.Skipf("import fails (reported by TestLibrary): %v", err)
			}
			judge(t, vvlib.CLIFileChecks(cli, lib))
			for _, s := range lib.Schemes() {
				s := s
				t.Run(s.Id, func(t *testing.T) { judge(t, []vvlib.Result{vvlib.CLISchemeCheck(cli, lib, s)}) })
			}
		})
	}
}

// TestLedgerIsUsed fails when known_issues.yml lists something no check produced
// (a typo in a scheme id or check name would otherwise hide for ever), and when
// fewer checks ran than min_checks says (a check that was removed or skipped
// would otherwise go unnoticed, because a passing check leaves no entry).
func TestLedgerIsUsed(t *testing.T) {
	keys, err := vvlib.ResultKeys(collect)
	if err != nil {
		t.Fatal(err)
	}
	if os.Getenv("VV_SKIP_TOOLCHAIN") != "" {
		t.Skip("not all checks ran")
	}
	if ledger.MinChecks > 0 && len(collect) < ledger.MinChecks {
		t.Errorf("%d checks ran, and min_checks in known_issues.yml is %d: a check was removed or skipped. If that is intended, lower min_checks in the same change", len(collect), ledger.MinChecks)
	}
	// PLAN.md 0.3: the checks a complete run produces are listed by name in
	// expected_checks.txt, so a check that disappears, or a new one, is seen
	// by name. VV_UPDATE_CHECKS=1 rewrites the list, in the change that adds
	// or removes checks, where a reviewer sees the difference.
	seen := map[string]bool{}
	for _, key := range keys {
		seen[key] = true
	}
	updating := os.Getenv("VV_UPDATE_CHECKS") != ""
	b, err := os.ReadFile("expected_checks.txt")
	if err != nil && !(updating && os.IsNotExist(err)) {
		t.Fatalf("expected_checks.txt: %v (run with VV_UPDATE_CHECKS=1 to write it)", err)
	}
	expected, err := vvlib.ParseExpectedChecks(string(b))
	if err != nil {
		t.Fatalf("expected_checks.txt: %v", err)
	}
	if updating {
		if err := os.WriteFile("expected_checks.txt", []byte(strings.Join(keys, "\n")+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	want := map[string]bool{}
	for _, k := range expected {
		want[k] = true
	}
	missing, added := []string{}, []string{}
	for k := range want {
		if !seen[k] {
			missing = append(missing, k)
		}
	}
	for _, k := range keys {
		if !want[k] {
			added = append(added, k)
		}
	}
	sort.Strings(missing)
	sort.Strings(added)
	if len(missing) > 0 {
		t.Errorf("%d listed check(s) did not run, e.g. %v: a check was removed or skipped. If that is intended, rewrite expected_checks.txt with VV_UPDATE_CHECKS=1 in the same change", len(missing), firstFew(missing))
	}
	if len(added) > 0 {
		t.Errorf("%d check(s) ran that expected_checks.txt does not list, e.g. %v: rewrite it with VV_UPDATE_CHECKS=1 in the change that adds them", len(added), firstFew(added))
	}
	for n, e := range ledger.Expected {
		for _, c := range e.Checks {
			one := vvlib.Expected{Issue: e.Issue, File: e.File, Scheme: e.Scheme, Checks: []string{c}}
			hit := false
			for _, r := range collect {
				if one.Matches(r) {
					hit = true
					break
				}
			}
			if !hit {
				t.Errorf("known_issues.yml entry %d (file %q, scheme %q, check %q) matches no check that ran", n+1, e.File, e.Scheme, c)
			}
		}
	}
}

// firstFew returns at most five names, for a message.
func firstFew(s []string) []string {
	if len(s) > 5 {
		return s[:5]
	}
	return s
}
