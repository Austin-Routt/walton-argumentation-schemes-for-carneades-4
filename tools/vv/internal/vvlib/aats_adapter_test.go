package vvlib

// This boundary test deliberately uses authored expectations. It reads only
// facts from the Python adapter, never its predicted argument identifiers or
// labels, and evaluates the current library in the actual vendored engine.
import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const aatsSamePairModel = `schema: walton-aats/v1
name: same_pair
source: independent adapter boundary regression
agents:
  alice: [act_a, act_b]
joint_actions:
  ja: {alice: act_a}
  jb: {alice: act_b}
audiences:
  cautious: [safety, benefit]
situations:
  - name: start
    source: same state pair has the same effects for both actions
    agent: alice
    values: {alice: [safety, benefit]}
    initial: s0
    states:
      s0: [risk]
      s1: [goal]
      s1_alias: [goal]
    transitions:
      - from: s0
        by: ja
        to: s1
        promotes: {benefit: [goal]}
        demotes: {safety: [risk]}
      - from: s0
        by: jb
        to: s1
        promotes: {benefit: [goal]}
        demotes: {safety: [risk]}
`

func aatsAdapterFacts(t *testing.T, model string) ([]byte, error) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "model.yml")
	if err := os.WriteFile(path, []byte(model), 0600); err != nil {
		t.Fatal(err)
	}
	code := `import json, sys
sys.path.insert(0, sys.argv[1])
import aats
try:
    print(json.dumps(aats.Model(sys.argv[2]).situations[0].facts()))
except aats.ModelError as exc:
    print(str(exc), file=sys.stderr)
    sys.exit(2)
`
	return exec.Command("python3", "-c", code, filepath.Join(RepoRoot(), "tools"), path).CombinedOutput()
}

func TestAATSSameStatePairActualEngine(t *testing.T) {
	data, err := aatsAdapterFacts(t, aatsSamePairModel)
	if err != nil {
		t.Fatalf("adapter: %v: %s", err, data)
	}
	var facts []string
	if err := json.Unmarshal(data, &facts); err != nil {
		t.Fatal(err)
	}
	c := Case{
		Name:         "aats/same-state-pair",
		Schemes:      []string{"practical_reasoning_as2", "practical_reasoning_as2_demotion"},
		ExactSchemes: []string{"practical_reasoning_as2", "practical_reasoning_as2_demotion"},
		Assumptions:  facts,
		Arguments: []string{
			"practical_reasoning_as2(alice,act_a,ja,s0,s1,goal,benefit)",
			"practical_reasoning_as2(alice,act_b,jb,s0,s1,goal,benefit)",
			"practical_reasoning_as2_demotion(alice,act_a,ja,s0,s1,risk,safety)",
			"practical_reasoning_as2_demotion(alice,act_b,jb,s0,s1,risk,safety)",
		},
		Out: []string{"should_perform(alice,act_a)", "should_perform(alice,act_b)",
			"¬should_perform(alice,act_a)", "¬should_perform(alice,act_b)"},
	}
	if problems := RunCase(RepoRoot(), c); len(problems) != 0 {
		t.Fatalf("actual engine disagrees with authored AS2 expectations: %v", problems)
	}
}

func TestAATSRejectsConflictingStatePairBeforeEngine(t *testing.T) {
	// Remove only jb's promotion. The pre-fix adapter accepted this model
	// and emitted facts which gave jb ja's promotion in the engine.
	model := strings.Replace(aatsSamePairModel,
		"by: jb\n        to: s1\n        promotes: {benefit: [goal]}\n",
		"by: jb\n        to: s1\n", 1)
	data, err := aatsAdapterFacts(t, model)
	if err == nil || !strings.Contains(string(data), "inconsistent effects") {
		t.Fatalf("inconsistent state-pair effects must be rejected before graph generation: err=%v output=%s", err, data)
	}
}
