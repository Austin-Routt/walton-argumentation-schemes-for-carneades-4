package test

import (
	"fmt"
	"github.com/carneades/carneades-4/src/engine/caes"
	"github.com/carneades/carneades-4/src/engine/caes/encoding/yaml"
	//	"log"
	"errors"
	"os"
	"testing"
)

// unsupportedExamples are the examples outside the supported profile
// (docs/contracts.md, contract 1, PLAN.md D3): chr-leq uses CHR deletions and
// guards, and Infer refuses it as unsupported instead of running it.
var unsupportedExamples = map[string]bool{"chr-leq.yml": true}

// skipUnsupported checks that Infer refuses a listed example as unsupported,
// and tells the caller to skip it.
func skipUnsupported(t *testing.T, name string) bool {
	t.Helper()
	if !unsupportedExamples[name] {
		return false
	}
	f, err := os.Open(yamlDir + name)
	if err != nil {
		t.Errorf("%s: %v", name, err)
		return true
	}
	defer f.Close()
	ag, err := yaml.Import(f)
	if err != nil {
		t.Errorf("%s: %v", name, err)
		return true
	}
	if err := ag.Infer(); !errors.Is(err, caes.ErrUnsupported) {
		t.Errorf("%s: Infer returned %v, want unsupported", name, err)
	}
	return true
}

const yamlDir = "../../examples/AGs/YAML/"
const yamlTmp = "/tmp/"
const jsonDir = "../../examples/AGs/AIF/"
const jsonTmp = "/tmp/"
const graphmlTmp = "/tmp/"

func check(t *testing.T, e error) {
	if e != nil {
		t.Errorf(e.Error())
		t.Skip("Skip Test")
	}
}

func checkLabeling(l caes.Labelling, stats map[string]*caes.Statement, expectedLbls map[string]caes.Label) error {
	errStr := ""
	for _, stat := range stats {
		calcLbl := l[stat]
		expLbl, found := expectedLbls[stat.Id]
		if found {
			if expLbl != calcLbl {
				if errStr == "" {
					errStr = fmt.Sprintf(" statement: %s, expected Label: %v, calculated Label: %v \n", stat.Id, expLbl, calcLbl)
				} else {
					errStr = errStr + fmt.Sprintf(" statement: %s, expected Label: %v, calculated Label: %v \n", stat.Id, expLbl, calcLbl)
				}
			}
		}
	}
	if errStr == "" {
		return nil
	}
	return errors.New(errStr)
}

// exportLoss lists the examples whose expected labels do not survive export
// and import: export writes weighing functions only as comments, so the
// constant weights of rebuttal, the criteria of mcda-porsche and the
// preferences of lex are lost (E-EXPORT in tools/vv/known_issues.yml,
// PLAN.md 4.1). checkRoundTrip fails for any other example that loses its
// labels, and for a listed one that keeps them, so the list cannot go stale.
var exportLoss = map[string]bool{"rebuttal.yml": true, "mcda-porsche.yml": true, "lex.yml": true}

func checkRoundTrip(t *testing.T, file string, err error) {
	t.Helper()
	switch {
	case err != nil && !exportLoss[file]:
		t.Errorf("%s: labels after export and import: %v", file, err)
	case err == nil && exportLoss[file]:
		t.Errorf("%s: export and import now keep the expected labels: remove it from exportLoss", file)
	case err != nil:
		t.Logf("%s: known export loss (E-EXPORT): %v", file, err)
	}
}

func printLabeling(l caes.Labelling) {

	for ref_stat, lbl := range l {
		fmt.Printf(" statement: %s    Label: %v\n", ref_stat.Id, lbl)
	}

}
