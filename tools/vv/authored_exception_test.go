package waltonvv

import (
	"fmt"
	"strings"
	"testing"

	"waltonvv/internal/vvlib"
)

// A scheme id plus an arbitrary nonnil undercutter used to pass check, after
// which eval accepted q(a) despite the stipulated exception e(a). Test both
// entry points in fresh processes, with and without authored parameters.
func TestAuthoredExceptionLinksCLI(t *testing.T) {
	cli, err := vvlib.FindCLI()
	if err != nil {
		t.Fatal(err)
	}
	base := `
language:
  p/1: "%s is a premise."
  q/1: "%s is a conclusion."
  e/1: "%s is an exception."
argument_schemes:
  - id: s
    variables: [X]
    premises: [p(X)]
    exceptions: [e(X)]
    conclusions: [q(X)]
statements:
  p(a): Premise A.
  e(a): Exception A.
  q(a): Conclusion A.
  inert: An unrelated statement.
arguments:
  authored:
    scheme: s
PARAMS    premises: [p(a)]
    conclusion: q(a)
    undercutter: inert
assumptions:
  - p(a)
  - e(a)
`
	for _, tc := range []struct{ name, params string }{{"explicit", "    parameters: [a]\n"}, {"inferred", ""}} {
		for _, cmd := range []struct {
			name string
			args []string
		}{{"check", []string{"check"}}, {"eval", []string{"eval", "-t", "yaml"}}} {
			t.Run(tc.name+"/"+cmd.name, func(t *testing.T) {
				r := cli.Run(strings.Replace(base, "PARAMS", tc.params, 1), cmd.args...)
				judge(t, []vvlib.Result{{
					File: "(engine)", Scheme: "authored-" + tc.name, Check: "cli/" + cmd.name + "-exception-links",
					Pass:   !r.TimedOut && r.Exit == 1 && strings.TrimSpace(r.Stdout) == "" && strings.Contains(r.Stderr, "exceptions"),
					Detail: fmt.Sprintf("inert undercutter: want invalid input with an exception-link diagnostic and no stdout; got exit %d, stdout %q, stderr %q", r.Exit, r.Stdout, r.Stderr),
				}})
			})
		}
	}
}
