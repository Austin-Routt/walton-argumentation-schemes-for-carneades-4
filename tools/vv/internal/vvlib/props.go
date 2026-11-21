package vvlib

// Property tests (PLAN.md 4.5): properties that must hold for every input of a
// kind, checked on generated inputs with a fixed seed, so that a failure is
// reproducible.

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/carneades/carneades-4/src/engine/caes"
	"github.com/carneades/carneades-4/src/engine/terms"
	chr "github.com/hfried/GoCHR/src/engine/CHR"
	chrterms "github.com/hfried/GoCHR/src/engine/terms"
)

// genTerm writes a random statement of the supported profile: names,
// integers, quoted strings, compounds, lists and negations.
func genTerm(r *rand.Rand, depth int, top bool) string {
	names := []string{"a", "b", "joe", "k_1", "x2", "acme", "opposition-normative"}
	if !top {
		switch r.Intn(6) {
		case 0:
			return names[r.Intn(len(names))]
		case 1:
			return fmt.Sprint(r.Intn(100) - 20)
		case 2:
			strs := []string{`"it rains"`, `"1990-05-17"`, `"a, b"`, `"x"`}
			return strs[r.Intn(len(strs))]
		case 3:
			if depth > 0 {
				n := 1 + r.Intn(3)
				items := make([]string, n)
				for i := range items {
					items[i] = genTerm(r, depth-1, false)
				}
				return "[" + strings.Join(items, ",") + "]"
			}
			return names[r.Intn(len(names))]
		}
	}
	if depth <= 0 {
		return names[r.Intn(len(names)-1)]
	}
	functors := []string{"p", "has", "says", "liable", "q_2"}
	n := 1 + r.Intn(3)
	args := make([]string, n)
	for i := range args {
		args[i] = genTerm(r, depth-1, false)
	}
	t := functors[r.Intn(len(functors))] + "(" + strings.Join(args, ",") + ")"
	if r.Intn(4) == 0 {
		t = "¬" + t
	}
	return t
}

// spaceless removes the spaces outside quoted strings: the two printers
// differ only there ("[a, b]" and "[a,b]").
func spaceless(s string) string {
	var b strings.Builder
	quoted := false
	for _, c := range s {
		if c == '"' {
			quoted = !quoted
		}
		if c == ' ' && !quoted {
			continue
		}
		b.WriteRune(c)
	}
	return b.String()
}

// termsProperty: a statement means the same to Carneades' reader and to
// GoCHR's, and Carneades' printer gives back what its reader reads.
func termsProperty() Result {
	r := rand.New(rand.NewSource(20261121))
	probs := []string{}
	c := StartCapture()
	for i := 0; i < 400 && len(probs) < 10; i++ {
		s := genTerm(r, 3, true)
		ct, ok := terms.ReadString(s)
		if !ok {
			probs = append(probs, s+": Carneades' reader refuses it")
			continue
		}
		c1 := ct.String()
		if ct2, ok := terms.ReadString(c1); !ok || ct2.String() != c1 {
			probs = append(probs, fmt.Sprintf("%s: Carneades prints %s and reads that back differently", s, c1))
			continue
		}
		g, ok := chr.ParseGoalString(s)
		if !ok {
			probs = append(probs, s+": GoCHR's reader refuses it")
			continue
		}
		var gs string
		if l, isList := g.(chrterms.List); isList && len(l) == 1 {
			gs = l[0].String()
		} else {
			gs = g.String()
		}
		if spaceless(gs) != spaceless(c1) || spaceless(c1) != spaceless(s) {
			probs = append(probs, fmt.Sprintf("%s: Carneades reads %s, GoCHR %s", s, c1, gs))
		}
	}
	c.Stop()
	return eres("property-terms-mean-the-same-to-both-readers", len(probs) == 0, "%s", strings.Join(probs, "; "))
}

// renamingProperty: renaming the variables of a scheme consistently changes
// none of the arguments it constructs.
func renamingProperty(lib *Library) Result {
	probs := []string{}
	lang := lib.AG.Theory.Language
	for _, s := range lib.Schemes() {
		in := NewInstance(s)
		r1 := Infer(lang, in.Synthetic, []*caes.Scheme{s}, in.Premises)
		ren := renamed(s)
		r2 := Infer(lang, in.Synthetic, []*caes.Scheme{ren}, in.Premises)
		if r1.TimedOut || r2.TimedOut || (r1.Err == nil) != (r2.Err == nil) {
			probs = append(probs, s.Id+": the runs end differently")
			continue
		}
		a1, a2 := argumentKeys(r1, s.Id), argumentKeys(r2, s.Id)
		if strings.Join(a1, " ") != strings.Join(a2, " ") {
			probs = append(probs, fmt.Sprintf("%s: %v, renamed %v", s.Id, a1, a2))
		}
	}
	return eres("property-renaming-variables-changes-no-argument", len(probs) == 0, "%s", strings.Join(probs, "; "))
}

// renamed copies a scheme with every variable V renamed to Vr0 (so that the
// new names collide neither with the old ones nor with fresh ones).
func renamed(s *caes.Scheme) *caes.Scheme {
	m := map[string]string{}
	for _, v := range s.Variables {
		m[v] = v + "r0"
	}
	ren := func(l []string) []string {
		out := make([]string, len(l))
		for i, st := range l {
			out[i] = renameVars(st, m)
		}
		return out
	}
	c := *s
	c.Variables = ren(s.Variables)
	c.Premises, c.Assumptions, c.Exceptions = ren(s.Premises), ren(s.Assumptions), ren(s.Exceptions)
	c.Conclusions, c.Deletions, c.Guards = ren(s.Conclusions), ren(s.Deletions), ren(s.Guards)
	return &c
}

func renameVars(st string, m map[string]string) string {
	var b, word strings.Builder
	flush := func() {
		w := word.String()
		if n, ok := m[w]; ok {
			w = n
		}
		b.WriteString(w)
		word.Reset()
	}
	quoted := false
	for _, c := range st {
		if c == '"' {
			quoted = !quoted
		}
		if !quoted && (c == '_' || c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
			word.WriteRune(c)
			continue
		}
		flush()
		b.WriteRune(c)
	}
	flush()
	return b.String()
}

func argumentKeys(r *Run, id string) []string {
	out := []string{}
	for _, a := range r.ArgsOf(id) {
		c := ""
		if a.Conclusion != nil {
			c = a.Conclusion.Id
		}
		out = append(out, strings.Join(a.Parameters, ",")+" -> "+c)
	}
	sort.Strings(out)
	return out
}

// PropertyChecks runs the property tests; renaming runs over walton.yml.
func PropertyChecks(root string) []Result {
	out := []Result{termsProperty()}
	lib, err := LoadLibrary(root, "walton.yml")
	if err != nil {
		return append(out, efail("property-renaming-variables-changes-no-argument", "walton.yml: %v", err))
	}
	return append(out, renamingProperty(lib))
}
