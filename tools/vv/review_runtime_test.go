package waltonvv

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/carneades/carneades-4/src/engine/caes"
	"github.com/carneades/carneades-4/src/engine/terms"
	chr "github.com/hfried/GoCHR/src/engine/CHR"
	"waltonvv/internal/vvlib"
)

const generatedOpposition = `
language:
  accepts/1: "The context accepts %s."
  rejects/1: "The context rejects %s."
issue_schemes:
  choice: [accepts(X), rejects(X)]
assumptions: [accepts(proposal), rejects(proposal)]
`

// These independently stated counterexamples supplement the positive scheme
// suite. The native engine tests exercise exact boundaries and reset behavior.
func TestIndependentRuntime(t *testing.T) {
	check := func(name string, pass bool, detail string) {
		judge(t, []vvlib.Result{{Scheme: "(engine)", Check: "independent/" + name, Pass: pass, Detail: detail}})
	}
	cli, err := vvlib.FindCLI()
	if err != nil {
		t.Fatal(err)
	}
	for _, declared := range []bool{false, true} {
		input := generatedOpposition
		if declared {
			input += "statements:\n  accepts(proposal): Accepted.\n  rejects(proposal): Rejected.\n"
		}
		for _, op := range []string{"check", "eval"} {
			args := []string{op}
			if op == "eval" {
				args = append(args, "-t", "yaml")
			}
			r := cli.Run(input, args...)
			check(fmt.Sprintf("%s-generated-opposition-declared-%v", op, declared), r.Exit == 3 && !r.TimedOut && strings.TrimSpace(r.Stdout) == "", fmt.Sprintf("exit=%d output=%q stderr=%q", r.Exit, r.Stdout, r.Stderr))
		}
	}
	{
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		ok := true
		for i := 0; i < 30; i++ {
			ag := caes.NewArgGraph()
			ag.Assumptions = []string{"p(a)"}
			ok = ok && errors.Is(ag.InferContext(ctx), caes.ErrCancelled) && len(ag.Statements) == 0
		}
		check("cancelled-empty-theory", ok, "a pre-cancelled request must not commit even when there are no schemes")
	}
	{
		rs := chr.MakeRuleStore()
		rs.Unique = true
		if err := rs.AddRule("join", []string{"p(X)", "q(X)"}, nil, nil, []string{"r(X)"}); err != nil {
			t.Fatal(err)
		}
		done := make(chan struct{})
		close(done)
		rs.Done = done
		ok, _, err := rs.Infer([]string{"p(a)"}, 10)
		check("cancelled-unmatched-join", !ok && errors.Is(err, chr.ErrCancelled), fmt.Sprint(err))
		rs.Done = nil
		rs.MaxMatchSteps = 200
		facts := []string{}
		for i := 0; i < 100; i++ {
			facts = append(facts, fmt.Sprintf("p(a%d)", i), fmt.Sprintf("q(b%d)", i))
		}
		ok, _, err = rs.Infer(facts, 10)
		check("unsuccessful-join-budget", !ok && errors.Is(err, chr.ErrLimit), fmt.Sprint(err))
	}
	{
		e := vvlib.EvalYAML(`
language:
  person/1: "%s is a person."
  pair/2: "%s pairs with %s."
argument_schemes:
  - id: pair_rule
    variables: [X,Y]
    premises: [person(X), person(Y)]
    conclusions: ["pair(X,Y)"]
assumptions: [person(alex)]
`)
		check("logical-conjunction-reuses-fact", e.InferErr == nil && e.ImportErr == nil && e.Label("pair(alex,alex)") == "in", fmt.Sprintf("import=%v infer=%v label=%s", e.ImportErr, e.InferErr, e.Label("pair(alex,alex)")))
	}
	for i, s := range []string{"p(", "p(a,)", "f(g(", "p([a,b)"} {
		_, ok := terms.ReadString(s)
		check(fmt.Sprintf("reject-truncated-term-%d", i+1), !ok, s)
	}
	{
		ag := caes.NewArgGraph()
		ag.Assumptions = []string{"p("}
		err := ag.Infer()
		check("invalid-empty-theory-is-atomic", errors.Is(err, caes.ErrInvalidInput) && len(ag.Statements) == 0, fmt.Sprint(err))
	}
}
