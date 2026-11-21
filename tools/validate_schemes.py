#!/usr/bin/env python3
"""Validate the Walton argumentation-scheme YAML files for Carneades 4.

Usage:  python3 tools/validate_schemes.py [repo_root] [--json]
Exit status 0 if there are no ERRORs (WARNINGs are reported but allowed).

Checks follow Carneades 4 semantics (github.com/carneades/carneades-4):
  * every file parses as YAML and has an 'argument_schemes' list;
  * every scheme has a unique id, only keys Carneades understands, list-valued
    premises/conclusions/assumptions/exceptions/variables, and at least one premise
    and one conclusion;
  * every statement is a well-formed predicate: [¬]name or [¬]name(arg, ...);
  * every variable used is declared, and every declared variable is used;
  * every predicate used has a language template keyed name/arity (0-ary predicates
    need name/0, which is how Carneades looks them up), with exactly as many %s/%v
    verbs as the arity (Carneades fills templates with Go's Sprintf);
  * the category files agree exactly with walton.yml, every scheme in walton.yml is
    in exactly one category file, and every category scheme is in walton.yml;
  * regression tests pin the direction of schemes whose logic was once wrong;
  * the engine's own rules for the supported profile (PLAN.md 3.10, docs/contracts.md)
    are errors: every variable occurs in a premise (range restriction); is, true and
    false are no predicate names; no assumption repeats a premise; a template has
    as many % signs as its predicate has arguments (the engine counts every %);
    deletions and guards are outside the profile; no YAML key occurs twice;
  * a conclusion that repeats a premise, and a cycle of schemes that derive each
    other's premises, are warnings;
  * one instance per repaired or corrected scheme, taken from the book, with the
    instantiated statements written out by hand, so that a flipped negation or
    two swapped arguments fail.
Requires PyYAML.
"""
import sys, os, re, glob, json

try:
    import yaml
except ImportError:  # pragma: no cover
    sys.exit("PyYAML is required: pip install pyyaml")

class _UniqueKeyLoader(yaml.SafeLoader):
    """A SafeLoader that records every mapping key that occurs twice: YAML itself
    keeps the last one silently."""
    duplicates = None

def _construct_mapping(loader, node, deep=False):
    seen = {}
    for key_node, _ in node.value:
        try:
            key = loader.construct_object(key_node, deep=deep)
            hash(key)
        except Exception:
            continue
        if key in seen:
            loader.duplicates.append((key, seen[key] + 1, key_node.start_mark.line + 1))
        else:
            seen[key] = key_node.start_mark.line
    return yaml.SafeLoader.construct_mapping(loader, node, deep=deep)

_UniqueKeyLoader.add_constructor(yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG, _construct_mapping)

def load_unique(text):
    """Load YAML; return (data, [(key, first line, second line)] for duplicate keys)."""
    loader = _UniqueKeyLoader(text)
    loader.duplicates = []
    try:
        return loader.get_single_data(), loader.duplicates
    finally:
        loader.dispose()

RESERVED = {"is", "true", "false"}          # docs/contracts.md, contract 7
WEIGHING_FUNCTIONS = {"linked", "convergent", "cumulative", "factorized"}
VARIABLE = re.compile(r"[A-Z][A-Za-z0-9_]*")

ALLOWED_KEYS = {"id", "meta", "variables", "premises", "conclusions", "assumptions",
                "exceptions", "weight", "roles", "deletions", "guards"}
LIST_FIELDS = ("premises", "conclusions", "assumptions", "exceptions")

# ------------------------------------------------------------------ strict term grammar
TOK = re.compile(r"\s*(?:(¬)|([a-z][A-Za-z0-9_]*)|([A-Z_][A-Za-z0-9_]*)|(-?\d+(?:\.\d+)?)|([(),]))")

class TermError(ValueError):
    pass

def parse_statement(s):
    """Return (negated, functor, arity, variables) or raise TermError."""
    toks, pos = [], 0
    s = s.strip()
    while pos < len(s):
        m = TOK.match(s, pos)
        if not m or m.end() == pos:
            raise TermError(f"unexpected character {s[pos]!r} at position {pos}")
        toks.append(m.groups()); pos = m.end()
    i = 0
    neg = False
    if i < len(toks) and toks[i][0]:
        neg = True; i += 1
    def term(i):
        if i >= len(toks): raise TermError("unexpected end of statement")
        neg_, atom, var, num, punct = toks[i]
        if neg_:                                   # ¬ may negate an argument, e.g. opposite_action(A,S,¬X)
            t, j = term(i + 1)
            return ("neg", "¬", [t]), j
        if var: return ("var", var, []), i + 1
        if num: return ("num", num, []), i + 1
        if atom:
            if i + 1 < len(toks) and toks[i + 1][4] == "(":
                args, j = [], i + 2
                while True:
                    a, j = term(j); args.append(a)
                    if j < len(toks) and toks[j][4] == ",": j += 1; continue
                    if j < len(toks) and toks[j][4] == ")": return ("cmp", atom, args), j + 1
                    raise TermError(f"expected ',' or ')' in arguments of {atom}")
            return ("atom", atom, []), i + 1
        raise TermError(f"unexpected token {[t for t in toks[i] if t][0]!r}")
    # A statement is a predicate, or a bare schema variable standing for a whole statement
    # (expert_opinion concludes S). A parenthesised variable such as (X) is rejected.
    if i >= len(toks) or not (toks[i][1] or toks[i][2]):
        raise TermError("statement must be a lower-case predicate or a schema variable")
    t, j = term(i)
    if j != len(toks):
        raise TermError("trailing tokens after the statement")
    vs = set()
    def walk(t):
        if t[0] == "var": vs.add(t[1])
        for a in t[2]: walk(a)
    walk(t)
    if t[0] == "var":
        return neg, None, 0, vs                    # schema variable: no functor, no template needed
    return neg, t[1], len(t[2]), vs


def issue_template_keys(issue_schemes):
    """Language dependencies of fixed or enumerated issue patterns."""
    if issue_schemes is None:
        return set()
    if not isinstance(issue_schemes, dict):
        raise TermError("issue_schemes must be a mapping")
    keys = set()
    for sid, patterns in issue_schemes.items():
        if not isinstance(patterns, list):
            raise TermError(f"issue {sid}: patterns must be a list")
        for pattern in patterns:
            if pattern == "...":
                continue
            if not isinstance(pattern, str):
                raise TermError(f"issue {sid}: pattern must be text")
            _, functor, arity, _ = parse_statement(pattern)
            if functor is not None:
                keys.add(f"{functor}/{arity}")
    return keys

VERB = re.compile(r"%(?:%|(\[\d+\])?([sv])|.?)")

def template_args(tpl, arity):
    """Simulate how Go's Sprintf consumes arguments for templates using %s, %v and %[n]s.
    Return (malformed_verbs, problems): problems lists arguments out of range or never used."""
    bad, used, argnum = [], set(), 0
    for m in VERB.finditer(tpl):
        v = m.group(0)
        if v == "%%": continue
        if not m.group(2):
            bad.append(v); continue
        if m.group(1): argnum = int(m.group(1)[1:-1]) - 1
        used.add(argnum); argnum += 1
    probs = []
    over = sorted(i + 1 for i in used if i < 0 or i >= arity)
    if over: probs.append(f"refers to argument(s) {over} but the arity is {arity} (Go prints %!s(MISSING))")
    unused = sorted(i + 1 for i in set(range(arity)) - used)
    if unused: probs.append(f"never prints argument(s) {unused} of {arity}")
    return bad, probs

# ------------------------------------------------------------------ helpers
class Report:
    def __init__(self): self.items = []
    def add(self, sev, code, where, msg): self.items.append(dict(severity=sev, code=code, where=where, message=msg))
    def errors(self): return [x for x in self.items if x["severity"] == "ERROR"]

def raw_duplicate_keys(path, section):
    """YAML silently keeps the last of duplicate mapping keys; find them in the raw text."""
    seen, dups, inside = {}, [], False
    with open(path, encoding="utf-8") as f:
        lines = f.readlines()
    for n, line in enumerate(lines, 1):
        if re.match(rf"^{section}:\s*$", line): inside = True; continue
        if inside and re.match(r"^\S", line): break
        m = re.match(r"^  ([^\s#][^:]*?):\s", line) if inside else None
        if m:
            k = m.group(1).strip().strip("'\"")
            if k in seen: dups.append((k, seen[k], n))
            else: seen[k] = n
    return dups

def scheme_statements(x):
    for f in LIST_FIELDS:
        v = x.get(f)
        if isinstance(v, list):
            for s in v: yield f, s

def check_file(path, rel, rep):
    try:
        with open(path, encoding="utf-8") as f:
            d, dups = load_unique(f.read())
        for k, a, b in dups:
            rep.add("ERROR", "YAML_DUPLICATE_KEY", rel, f"key {k!r} at lines {a} and {b}; YAML keeps only the last")
    except yaml.YAMLError as e:
        mark = getattr(e, "problem_mark", None)
        loc = f" (line {mark.line + 1}, column {mark.column + 1})" if mark else ""
        rep.add("ERROR", "YAML_PARSE", rel, f"{getattr(e, 'problem', str(e))}{loc}")
        return None
    if not isinstance(d, dict) or not isinstance(d.get("argument_schemes"), list):
        rep.add("ERROR", "NO_SCHEMES", rel, "missing top-level 'argument_schemes' list"); return None
    lang = d.get("language") or {}
    if not isinstance(lang, dict):
        rep.add("ERROR", "LANGUAGE_NOT_MAPPING", rel, "'language' must be a mapping"); lang = {}
    for k, a, b in raw_duplicate_keys(path, "language"):
        rep.add("ERROR", "LANGUAGE_DUPLICATE_KEY", rel, f"template {k!r} defined at lines {a} and {b}; YAML keeps only the last")
    for k, v in lang.items():
        if not re.fullmatch(r"¬?[a-z][A-Za-z0-9_]*/\d+", str(k)):
            rep.add("ERROR", "LANGUAGE_KEY_FORMAT", rel, f"key {k!r} is not name/arity; Carneades never finds it (0-ary predicates need name/0)")
            continue
        name, arity = str(k).split("/")
        if name.lstrip("¬") in RESERVED:
            rep.add("ERROR", "RESERVED_PREDICATE", rel, f"{k}: {name} is a reserved name and cannot be a predicate (contract 7)")
        bad, probs = template_args(str(v), int(arity))
        if bad: rep.add("ERROR", "TEMPLATE_BAD_VERB", rel, f"{k}: malformed format verb(s) {bad} in {v!r}")
        for p in probs: rep.add("ERROR", "TEMPLATE_ARITY", rel, f"{k}: template {p}: {v!r}")
        if str(v).count("%") != int(arity):
            rep.add("ERROR", "TEMPLATE_PERCENT_COUNT", rel, f"{k}: the engine's validator counts every % sign and finds {str(v).count('%')} for arity {arity}: {v!r}")
    ids, used = {}, set()
    for i, x in enumerate(d["argument_schemes"]):
        where = f"{rel}#{i + 1}"
        if not isinstance(x, dict):
            rep.add("ERROR", "SCHEME_NOT_MAPPING", where, f"entry is {type(x).__name__}, not a mapping: {str(x)[:60]!r}"); continue
        sid = x.get("id")
        if not isinstance(sid, str) or not sid.strip():
            rep.add("ERROR", "ID_MISSING", where, "scheme has no id"); continue
        where = f"{rel}:{sid}"
        if sid in ids: rep.add("ERROR", "ID_DUPLICATE", where, f"id also used by entry {ids[sid] + 1}; Carneades keeps only the last")
        ids[sid] = i
        for k in x:
            if k not in ALLOWED_KEYS: rep.add("ERROR", "UNKNOWN_KEY", where, f"key {k!r} is not read by Carneades (did you mean 'assumptions'?)" if k == "assumption" else f"key {k!r} is not read by Carneades")
        meta = x.get("meta")
        if not isinstance(meta, dict) or not meta.get("title"): rep.add("WARNING", "META_TITLE_MISSING", where, "no meta.title")
        elif not str(meta.get("source") or "").strip(): rep.add("WARNING", "META_SOURCE_MISSING", where, "no meta.source")
        for f in LIST_FIELDS + ("variables", "deletions", "guards", "roles"):
            if f in x and x[f] is not None and not isinstance(x[f], list):
                rep.add("ERROR", "FIELD_NOT_LIST", where, f"{f!r} is a {type(x[f]).__name__} ({str(x[f])[:50]!r}), not a list")
        for f in ("deletions", "guards"):
            if x.get(f):
                rep.add("ERROR", "UNSUPPORTED", where, f"{f!r}: CHR deletions and guards are outside the supported profile (PLAN.md D3); the engine refuses the scheme")
        if isinstance(x.get("variables"), list):
            names = x["variables"]
            for v in names:
                if not isinstance(v, str) or not VARIABLE.fullmatch(v):
                    rep.add("ERROR", "VARIABLE_NAME", where, f"{v!r} is not a variable name (a capital letter, then letters, digits or _)")
            hashable = [v for v in names if isinstance(v, str)]
            for v in sorted({v for v in hashable if hashable.count(v) > 1}):
                rep.add("ERROR", "VARIABLE_DUPLICATE", where, f"variable {v} is declared more than once")
        if "weight" in x:
            w = x["weight"]
            if not (isinstance(w, dict) or (isinstance(w, str) and w in WEIGHING_FUNCTIONS)):
                rep.add("ERROR", "WEIGHT_UNKNOWN", where, f"weight {w!r} is neither a basic weighing function {sorted(WEIGHING_FUNCTIONS)} nor a mapping")
        if "roles" in x and isinstance(x["roles"], list):
            if not all(isinstance(r, str) for r in x["roles"]) or len(x["roles"]) > len(x.get("premises") or []):
                rep.add("ERROR", "ROLES", where, "roles must be names, at most one per premise")
        if not (isinstance(x.get("premises"), list) and x["premises"]): rep.add("ERROR", "NO_PREMISES", where, "scheme has no premises")
        if not (isinstance(x.get("conclusions"), list) and x["conclusions"]): rep.add("ERROR", "NO_CONCLUSIONS", where, "scheme has no conclusions")
        declared = {v for v in (x.get("variables") or []) if isinstance(v, str)} if isinstance(x.get("variables"), list) else set()
        allvars, premvars = set(), set()
        norm = lambda s: str(s).replace(" ", "")
        premises = {norm(s) for s in x.get("premises") or []} if isinstance(x.get("premises"), list) else set()
        for f, s in scheme_statements(x):
            try:
                neg, fn, ar, vs = parse_statement(str(s))
            except TermError as e:
                rep.add("ERROR", "STATEMENT_SYNTAX", where, f"{f}: {s!r}: {e}"); continue
            allvars |= vs
            if f == "premises": premvars |= vs
            if f == "assumptions" and norm(s) in premises:
                rep.add("ERROR", "ASSUMPTION_IS_PREMISE", where, f"assumption {s!r} repeats a premise: a stipulated premise can never be questioned, and the rule feeds itself")
            if f == "conclusions" and norm(s) in premises:
                rep.add("WARNING", "CONCLUSION_IS_PREMISE", where, f"conclusion {s!r} repeats a premise")
            if fn is None:                               # schema variable
                if f in ("assumptions", "exceptions"):
                    rep.add("ERROR", "VARIABLE_STATEMENT", where, f"{f}: {s!r}: an assumption or exception may not be a bare variable (the engine's validator rejects it)")
                continue
            if fn in RESERVED:
                rep.add("ERROR", "RESERVED_PREDICATE", where, f"{f}: {s!r}: {fn} is a reserved name and cannot be a predicate (contract 7)")
            key = f"{fn}/{ar}"; used.add(key)
            if key not in lang:
                rep.add("ERROR", "TEMPLATE_MISSING", where, f"{f}: {s!r} needs language template {key!r} in this file")
        for v in sorted(allvars - declared): rep.add("ERROR", "VARIABLE_UNDECLARED", where, f"variable {v} is used but not declared in 'variables'")
        for v in sorted(declared - allvars): rep.add("WARNING", "VARIABLE_UNUSED", where, f"variable {v} is declared but never used")
        unbound = sorted((allvars & declared) - premvars)
        if unbound: rep.add("ERROR", "RANGE_RESTRICTION", where, f"variable(s) {unbound} occur only outside the premises: the engine's validator rejects the scheme, and an argument from it holds statements with variables (contract 1)")
    try:
        issue_used = issue_template_keys(d.get("issue_schemes"))
    except TermError as error:
        rep.add("ERROR", "ISSUE_PATTERN_SYNTAX", rel, str(error))
    else:
        used |= issue_used
        for key in sorted(issue_used - set(lang)):
            rep.add("ERROR", "ISSUE_TEMPLATE_MISSING", rel, f"issue pattern needs language template {key!r} in this file")
    for k in sorted(set(lang) - used):
        rep.add("WARNING", "TEMPLATE_UNUSED", rel, f"template {k!r} is not used by any scheme in this file")
    return d

# ------------------------------------------------------------------ regression tests (direction of corrected schemes)
# Each entry pins a scheme whose logic once contradicted its source, so the error cannot return.
# "WRM 2008" is Walton, Reed & Macagno, Argumentation Schemes, Cambridge University Press, 2008.
def regression(master, rep):
    S = {x["id"]: x for x in master["argument_schemes"] if isinstance(x, dict) and isinstance(x.get("id"), str)}
    def fns(x, f):
        out = []
        for s in x.get(f) or []:
            try: out.append(parse_statement(str(s))[1])
            except TermError: out.append(None)
        return out
    def args(s):
        m = re.search(r"\((.*)\)", str(s)); return [a.strip() for a in m.group(1).split(",")] if m else []
    def stmt(x, f, name):
        return next((s for s in x.get(f) or [] if str(s).lstrip("¬").startswith(name + "(")), None)
    checks = [
        ("division_generic", lambda x: fns(x, "premises") == ["has_property"] and fns(x, "conclusions") == ["all_parts_property"]
            and fns(x, "assumptions") == ["divisionally_hereditary"], "division runs whole -> parts and asks whether the property passes to the parts (WRM 2008, sec. 10.1, p. 317)"),
        ("composition_generic", lambda x: fns(x, "premises") == ["all_parts_property"] and fns(x, "conclusions") == ["has_property"]
            and fns(x, "assumptions") == ["compositionally_hereditary"], "composition runs parts -> whole and asks whether the whole inherits the property (sec. 9.1, p. 316)"),
        ("evidence_hypothesis_falsification", lambda x: [str(c).replace(" ", "") for c in x["conclusions"]] == ["¬H"],
            "falsification concludes that the hypothesis H is false (sec. 32.2, pp. 331-332)"),
        ("from_example", lambda x: fns(x, "premises") == ["has_properties"] and fns(x, "conclusions") == ["usually_properties"],
            "argument from example runs from the particular case to the generalization (sec. 6.1, p. 314)"),
        ("common_folk_ad_populum", lambda x: fns(x, "premises") == ["background"] and fns(x, "conclusions") == ["being_credible"],
            "the speaker's ordinariness is the premise and greater credibility the conclusion (sec. 4.10, p. 313)"),
        ("moral_justification_ad_populum_excuse", lambda x: fns(x, "conclusions") == ["morally_acceptable_action"],
            "the conclusion is that the speaker's action is morally justified (sec. 4.6, p. 312)"),
        ("danger", lambda x: stmt(x, "premises", "occur") and args(x["conclusions"][0]) == args(stmt(x, "premises", "occur"))[:1],
            "the conclusion is not to bring about the action, not the danger (sec. 37, p. 334)"),
        ("threat", lambda x: "position_to_bring_about" in fns(x, "premises") and len(args(stmt(x, "premises", "cited_bad_consequences") or "")) == 2,
            "the proponent is in a position to bring about the cited consequence (sec. 35, p. 333)"),
        ("abductive_backward", lambda x: "successful_account" in fns(x, "premises"),
            "the chosen account must explain the data most successfully (sec. 31.1, pp. 329-330)"),
        ("opposition_normative", lambda x: "cannot_opposite_evaluation" in fns(x, "premises")
            and args(stmt(x, "premises", "receives_evaluation"))[1:] == args(stmt(x, "premises", "cannot_opposite_evaluation"))[1:2],
            "the minor premise instantiates the major premise's antecedent (sec. 11.2, p. 318)"),
        ("sacrifice", lambda x: args(stmt(x, "premises", "great_sacrifice"))[2:] == args(x["conclusions"][0])[1:],
            "the value in the conditional premise is the value concluded (sec. 20, p. 322)"),
    ]
    for sid, cond, msg in checks:
        if sid not in S:
            rep.add("ERROR", "REGRESSION", f"walton.yml:{sid}", "scheme missing"); continue
        try: ok = bool(cond(S[sid]))
        except Exception: ok = False
        if not ok: rep.add("ERROR", "REGRESSION", f"walton.yml:{sid}", msg)

# ------------------------------------------------------------------ cycles across schemes (warnings)
def cycles(master, rep):
    """Warn about schemes that derive each other's premises (or their own). The
    engine terminates on them since set semantics (PLAN.md 2.2); a cycle is still
    worth knowing about. Bare-variable statements are left out: they would join
    everything."""
    S = [x for x in master["argument_schemes"] if isinstance(x, dict) and isinstance(x.get("id"), str)]
    def keys(x, fields):
        out = set()
        for f in fields:
            for s in x.get(f) or []:
                try:
                    neg, fn, ar, _ = parse_statement(str(s))
                except TermError:
                    continue
                if fn is not None:
                    out.add(("¬" if neg else "") + f"{fn}/{ar}")
        return out
    concl = {x["id"]: keys(x, ("conclusions",)) for x in S}
    prem = {x["id"]: keys(x, ("premises",)) for x in S}
    edges = {a: sorted(b for b in prem if concl[a] & prem[b]) for a in concl}
    index, low, stack, on, sccs, counter = {}, {}, [], set(), [], [0]
    def strong(v):
        index[v] = low[v] = counter[0]; counter[0] += 1
        stack.append(v); on.add(v)
        for w in edges[v]:
            if w not in index:
                strong(w); low[v] = min(low[v], low[w])
            elif w in on:
                low[v] = min(low[v], index[w])
        if low[v] == index[v]:
            comp = []
            while True:
                w = stack.pop(); on.discard(w); comp.append(w)
                if w == v: break
            if len(comp) > 1 or v in edges[v]:
                sccs.append(sorted(comp))
    sys.setrecursionlimit(10000)
    for v in sorted(edges):
        if v not in index:
            strong(v)
    for comp in sorted(sccs):
        rep.add("WARNING", "SCHEME_CYCLE", "walton.yml", f"schemes that derive each other's premises: {comp}")

# ------------------------------------------------------------------ instances from the book (PLAN.md 3.10)
# One instance per scheme that a changelog entry corrected or PLAN.md Phase 3
# repaired or added: the variables are bound to constants, and every statement
# the record should yield is written out by hand from the book's entry. A record
# that yields anything else fails: a flipped negation, two swapped arguments, a
# missing or extra statement. Page numbers are those of the compendium (WRM 2008).
INSTANCES = [
    ("end_and_means_2", "22.6.2, p.325; one evaluation context", {"X": "money", "C": "retirement"},
     {"premises": ["means_end(money,retirement)", "more_important_ends(retirement)"], "conclusions": ["ends_important(money,retirement)"]}),
    ("snob_appeal_ad_populum", "4.7, p.312; explicit acceptance branch", {"A": "luxury_watches_are_desirable", "G": "jet_set", "O": "watch", "P": "wealth"},
     {"premises": ["property_or_object(jet_set,luxury_watches_are_desirable)", "group_accepted(jet_set,luxury_watches_are_desirable)", "elite_group(jet_set)", "goal_prestige", "accept_property_object(luxury_watches_are_desirable,wealth,watch,jet_set)"], "conclusions": ["should_accept(luxury_watches_are_desirable)"]}),
    ("whole_into_parts", "10.2, p.317; explicit sufficient-direction adaptation", {"X": "machine", "X1": "power_unit", "X2": "control_unit", "X3": "motor", "A": "electrically_live"},
     {"premises": ["whole_of(machine,power_unit,control_unit,motor)", "part_whole(machine,electrically_live,power_unit,control_unit,motor)", "selected_part_sufficient_for_whole(machine,power_unit,electrically_live)", "has_property(power_unit,electrically_live)"], "conclusions": ["has_property(machine,electrically_live)"]}),
    # the eleven corrections of the changelog (2026-11-06)
    ("division_generic", "10.1, p. 317", {"A": "car", "P": "heavy"},
     {"premises": ["has_property(car,heavy)"], "assumptions": ["divisionally_hereditary(car,heavy)"], "conclusions": ["all_parts_property(car,heavy)"]}),
    ("composition_generic", "9.1, p. 316", {"A": "car", "P": "heavy"},
     {"premises": ["all_parts_property(car,heavy)"], "assumptions": ["compositionally_hereditary(car,heavy)"], "conclusions": ["has_property(car,heavy)"]}),
    ("evidence_hypothesis_falsification", "32.2, pp. 331-332", {"H": "h", "A": "e"},
     {"premises": ["hypothesis_event(h,e)", "event_false(e)"], "assumptions": ["sufficient_correlation(h,e)"], "conclusions": ["¬h"]}),
    ("from_example", "6.1, p. 314", {"A": "tweety", "P1": "bird", "P2": "flies"},
     {"premises": ["has_properties(tweety,bird,flies)"], "assumptions": ["typical_example(tweety)", "supports_generalization(tweety,bird,flies)"], "exceptions": ["impair_general(tweety)"], "conclusions": ["usually_properties(bird,flies)"]}),
    ("common_folk_ad_populum", "4.10, p. 313", {"S": "speaker"},
     {"premises": ["background(speaker)"], "conclusions": ["being_credible(speaker)"]}),
    ("moral_justification_ad_populum_excuse", "4.6, p. 312", {"G": "g", "X": "speaker", "A": "act"},
     {"premises": ["group_does(g,speaker,act)", "acceptable_norm(speaker,act,g)", "group_member(g)"], "conclusions": ["morally_acceptable_action(speaker,act)"]}),
    ("danger", "37, p. 334", {"X": "act", "Y": "harm"},
     {"premises": ["occur(act,harm)", "danger(harm)"], "conclusions": ["¬ought_to_bring_about(act)"]}),
    ("threat", "35, p. 333", {"X": "act", "C": "harm"},
     {"premises": ["cited_bad_consequences(act,harm)", "position_to_bring_about(harm)", "assert_bad(act,harm)"], "conclusions": ["¬ought_to_bring_about(act)"]}),
    ("abductive_backward", "31.1, pp. 329-330", {"A": "a1", "A1": "best", "D": "d"},
     {"premises": ["data_set(d)", "data_account(a1,d)", "member_of_accounts(best,a1)", "successful_account(best,a1,d)"], "assumptions": ["sufficient_account(best,d)"], "exceptions": ["inquiry_incomplete(d)"], "conclusions": ["plausible_hypothesis(best,a1,d)"]}),
    ("opposition_normative", "11.2, p. 318", {"A": "act", "X": "good", "X2": "bad"},
     {"premises": ["cannot_opposite_evaluation(act,good,bad)", "receives_evaluation(act,good)"], "conclusions": ["cannot_receive_evaluation(act,bad)"]}),
    ("sacrifice", "20, p. 322", {"A": "a", "S": "s", "V": "v"},
     {"premises": ["sacrifice_made(a,s)", "great_sacrifice(s,a,v)", "great_sacrifice_made(s,a)"], "conclusions": ["sacrifice_value(a,v)"]}),
    # the repairs of PLAN.md 3.1 to 3.7
    ("ad_hominem", "42, p. 336", {"A": "bob", "S": "arg1", "C": "tax_cut"},
     {"premises": ["bad_character(bob)", "advocates_conclusion(bob,arg1,tax_cut)"], "assumptions": ["character_relevant(arg1)"], "exceptions": ["stronger_than_character(arg1)"], "conclusions": ["argument_not_accepted(bob,arg1)"]}),
    ("regulative_rule", "55.3, p. 343", {"A": "sam", "X": "pay_tax"},
     {"premises": ["obligatory(pay_tax)", "do_obligator", "addressee_of(sam,pay_tax)"], "exceptions": ["excuse_or_overriding_duty(sam,pay_tax)"], "conclusions": ["must_do(sam,pay_tax)"]}),
    ("regulative_rule_prohibition", "55.3, p. 343", {"A": "sam", "X": "smoke"},
     {"premises": ["prohibited(smoke)", "do_obligator", "addressee_of(sam,smoke)"], "exceptions": ["excuse_or_overriding_duty(sam,smoke)"], "conclusions": ["must_not_do(sam,smoke)"]}),
    ("practical_necessary_condition", "22.2, p. 323", {"X": "b1", "A": "a", "P": "me", "S": "alts"},
     {"premises": ["goal_of_agent(me,a)", "alternative_action(me,alts,a)", "necessary_action(me,b1,alts,a)", "practicality_action(me,b1)", "more_acceptable_bringing(me,a,b1)"],
      "assumptions": ["possible_goal(a)", "acceptable_action(me,b1)", "possible_bring_about(me,b1)"],
      "exceptions": ["alternative_means(me,a,b1)", "side_effect_action(me,b1)", "conflicting_goals(me,a)"], "conclusions": ["should_do(me,b1)"]}),
    ("practical_sufficient_condition", "22.3, p. 323", {"X": "b1", "A": "a", "P": "me", "S": "alts"},
     {"premises": ["goal_of_agent(me,a)", "alternative_action_sufficient(me,alts,a)", "sufficient_action(me,b1,alts,a)", "practicality_action(me,b1)", "more_acceptable_bringing(me,a,b1)"],
      "assumptions": ["possible_goal(a)", "acceptable_action(me,b1)", "possible_bring_about(me,b1)"],
      "exceptions": ["alternative_means(me,a,b1)", "side_effect_action(me,b1)", "conflicting_goals(me,a)"], "conclusions": ["should_do(me,b1)"]}),
    ("sorites_slippery_slope", "51, pp. 340-341", {"A2": "a0", "A1": "a1", "A": "an", "P": "bald"},
     {"premises": ["beyond_contention(a0,bald)", "also_has(a0,bald,a1)", "reapplication_sequence(a0,an,bald)"], "conclusions": ["may_have(an,bald)"]}),
    ("pop_scheme", "4.1, p. 311", {"A": "claim", "G": "g"},
     {"premises": ["group_accepted(g,claim)"], "conclusions": ["claim"]}),
    ("pop_scheme_acceptance", "4.1, p. 311", {"A": "claim", "G": "g"},
     {"premises": ["group_accepted(g,claim)"], "conclusions": ["should_accept(claim)"]}),
    ("group_2", "21.2, p. 322", {"G": "g", "Q": "q", "M": "m"},
     {"premises": ["has_property(g,q)", "group_member(m,g)", "group_member_quality(g,q,m)"], "conclusions": ["has_property(m,q)"]}),
    ("rule_physical_world_1", "54.1, p. 342", {"W": "w", "D": "d"},
     {"premises": ["is_a(w,d)", "referentially_coextensive(w,d)", "synonymous"], "exceptions": ["not_necessarily_true(w,d)"], "conclusions": ["means_same(w,d)"]}),
    ("exceptional", "56, p. 344", {"A": "a"},
     {"premises": ["rule_can_waived(a)", "case_exception(a)"], "assumptions": ["exception_backed(a)"], "conclusions": ["rule_waived(a)"]}),
    # the records added by PLAN.md 3.8
    ("precedent_slippery_slope", "50, p. 340", {"C": "c0", "R": "r", "N": "cn"},
     {"premises": ["claimed_exception(c0,r)", "precedent_series(c0,r,cn)", "intolerable_exception(c0,r)"], "conclusions": ["cannot_be_exception(c0,r)"]}),
    ("from_precedent", "57, p. 344", {"R": "r", "F": "f", "G": "g", "A": "a"},
     {"premises": ["established_rule_property(r,f,g)", "has_property(a,f)", "¬has_property(a,g)"], "assumptions": ["legitimate_case(a,r)"],
      "exceptions": ["rule_inapplicable_to(r,a)", "recognized_exception_category(a,r)"], "conclusions": ["rule_exception_recognized(r,a)"]}),
    ("value_based_practical_reasoning", "22.4, p. 324", {"P": "me", "G": "g", "V": "v", "A": "a"},
     {"premises": ["goal_of_agent(me,g)", "goal_supported_by_values(me,g,v)", "necessary_or_sufficient_for_goal(me,a,g)"], "assumptions": ["possible_bring_about(me,a)"],
      "exceptions": ["conflicting_goals(me,g)", "alternative_means(me,g,a)", "more_efficient_alternative(me,a,g)", "better_alternative_by_values(me,a,v)", "worse_consequences(me,a,g)"],
      "conclusions": ["should_do(me,a)"]}),
    # the changes of PLAN.md 5.1 and 5.2
    ("popular_opinion", "4, p. 311", {"A": "tax_cut"},
     {"premises": ["generally_true(tax_cut)", "favor(tax_cut)"], "exceptions": ["doubting_general(tax_cut)"], "conclusions": ["reason_in_favor(tax_cut)"]}),
    ("verbal_classification", "14, p. 319", {"O": "rex", "F": "dog", "G": "mammal"},
     {"premises": ["instance(rex,dog)", "subclass(dog,mammal)"], "exceptions": ["doubtful_classification(dog,mammal)"], "conclusions": ["instance(rex,mammal)"]}),
    ("verbal_classification_vagueness", "16, pp. 319-320", {"A": "arg1", "F": "heap"},
     {"premises": ["vague_defined_property(arg1,heap)", "precision_needed(arg1)", "property_vagueness(heap,arg1)"], "exceptions": ["certainty_no_need(arg1,heap)"], "conclusions": ["rejected_as_deficient(arg1)"]}),
    ("verbal_classification_arbitrariness", "17, p. 320", {"A": "arg1", "F": "poverty"},
     {"premises": ["arbitrary_defined_property(arg1,poverty)", "needed_nonarbitrary(arg1,poverty)", "arbitrary_property(poverty,arg1)"], "exceptions": ["arbitrariness_harmless(arg1,poverty)"], "conclusions": ["rejected_as_deficient(arg1)"]}),
    ("two_person_practical", "23, pp. 325-326", {"A": "ann", "B": "bob", "X": "trip", "Y": "ticket"},
     {"premises": ["tells_intend(ann,trip,bob)", "situation_necessity(bob,ticket,trip,ann)"], "assumptions": ["possible_bring_about(ann,ticket)"],
      "exceptions": ["alternative_means(ann,trip,ticket)", "conflicting_goals(ann,trip)", "side_effect_other_act(ticket,ann)", "other_actions_required(ann,trip,ticket)"], "conclusions": ["should_do(ann,ticket)"]}),
    ("waste", "24, p. 326", {"A": "ann", "X": "thesis"},
     {"premises": ["stop_realize_wasted(ann,thesis)", "waste_bad_thing(ann,thesis)"], "assumptions": ["possible_goal(thesis)"], "exceptions": ["reassessment_against_continuing(ann,thesis)"], "conclusions": ["ought_to_continue(ann,thesis)"]}),
    ("negative_normal_expectations", "26.1, p. 327", {"A": "alarm_silent"},
     {"premises": ["normal_true(alarm_silent)", "¬alarm_silent"], "exceptions": ["uninvestigated(alarm_silent)"], "conclusions": ["not_normal(alarm_silent)"]}),
    ("abductive_forward", "31.2, pp. 329-330", {"D": "d", "G": "g", "C": "c", "C1": "c1", "G1": "g1"},
     {"premises": ["data_set(d)", "diagram_set(g,d,c)", "successful_diagram(g,g1)", "diagram_conclusion(g1,c1)"], "assumptions": ["sufficient_account(c1,d)"], "exceptions": ["inquiry_incomplete(d)"], "conclusions": ["plausible_conclusion(c1)"]}),
    ("abductive_action_to_character", "31.3, p. 330", {"A": "bob", "X": "kept_promise", "Q": "honest"},
     {"premises": ["action_character(bob,kept_promise,honest)"], "assumptions": ["quality_defined(honest)", "action_fits_quality(bob,kept_promise,honest)"], "conclusions": ["character_quality(bob,honest)", "has_property(bob,honest)"]}),
    ("abductive_character_to_action", "31.4, p. 330", {"A": "bob", "Q": "honest"},
     {"premises": ["character_quality(bob,honest)"], "assumptions": ["quality_defined(honest)"], "exceptions": ["not_character_action(bob,honest)"], "conclusions": ["character_action(bob,honest)"]}),
    ("retroductive", "31.5, p. 331", {"A": "ann", "Q": "greedy", "E": "missing_cash"},
     {"premises": ["agent_event(ann,missing_cash)", "event_quality(missing_cash,greedy)", "has_property(ann,greedy)"], "assumptions": ["quality_defined(greedy)", "action_fits_quality(ann,missing_cash,greedy)", "reference_class(ann,missing_cash,greedy)"], "conclusions": ["event_question(ann,missing_cash)"]}),
    ("positive_consequences", "33.1, pp. 332-333; book-supported CQ adaptation", {"A": "plant_tree", "G": "shade"},
     {"premises": ["brings_about(plant_tree,shade)", "good(shade)"], "assumptions": ["consequence_likelihood_adequate(plant_tree,shade)"], "conclusions": ["should_be_performed(plant_tree)"]}),
    ("negative_consequences", "33.2, pp. 332-333; book-supported CQ adaptation", {"A": "dump_waste", "G": "pollution"},
     {"premises": ["brings_about(dump_waste,pollution)", "bad(pollution)"], "assumptions": ["consequence_likelihood_adequate(dump_waste,pollution)"], "conclusions": ["¬should_be_performed(dump_waste)"]}),
    ("established_rule", "55.1, p. 343; book-supported CQ adaptation of inherited form", {"C": "pay_fine", "R": "traffic_rule"},
     {"premises": ["has_conclusion(traffic_rule,pay_fine)", "applicable(traffic_rule)"], "assumptions": ["valid(traffic_rule)"], "exceptions": ["conflicting_or_overriding_rule(traffic_rule,pay_fine)", "rule_exception_or_excuse(traffic_rule,pay_fine)"], "conclusions": ["pay_fine"]}),
    ("negative_consequences_reasoning", "33.3, p. 332", {"A": "tax", "B": "recession", "P": "me"},
     {"premises": ["agent_bring_about(me,tax,recession)", "bad_for_goals(me,recession)"], "assumptions": ["consequence_likelihood_adequate(tax,recession)"], "conclusions": ["¬should_do(me,tax)"]}),
    ("help", "38, p. 334", {"A": "ann", "B": "bob", "X": "lift"},
     {"premises": ["help_needed(ann,bob)", "situation_help(ann,lift,bob)", "can_carry_out(bob,lift)", "not_too_costly(bob,lift)"], "assumptions": ["action_will_help(lift,ann)"], "conclusions": ["should_do(bob,lift)"]}),
    ("distress", "39, pp. 334-335", {"B": "bob", "X": "lift", "A": "ann"},
     {"premises": ["distress(ann)", "relieve(bob,lift,ann)"], "assumptions": ["possible_bring_about(bob,lift)"], "exceptions": ["great_side_effects(bob,lift)"], "conclusions": ["should_do(bob,lift)"]}),
    ("commitment1", "40, p. 335", {"A": "ann", "X": "p", "Y": "q"},
     {"premises": ["committed(ann,p)", "also_committed(p,q)"], "exceptions": ["commitment_contrary_evidence(ann,p)", "commitment_linkage_exception(ann,p,q)"], "conclusions": ["committed_to(ann,q)"]}),
    ("commitment2", "40, p. 335", {"A": "ann", "X": "p", "Y": "q"},
     {"premises": ["also_committed_questioned(ann,p,q)", "dialog_committed(ann,p)"], "exceptions": ["commitment_contrary_evidence(ann,p)", "commitment_linkage_exception(ann,p,q)"], "conclusions": ["commitment_arises(ann,q)"]}),
    ("pragmatic_inconsistency", "43, pp. 336-337", {"A": "bob", "S": "arg1", "X": "tax_cut"},
     {"premises": ["advocates_conclusion(bob,arg1,tax_cut)", "opposite_action(bob,arg1,¬tax_cut)"], "assumptions": ["strong_advocate(bob,arg1,tax_cut)", "inconsistency_relevant(bob,arg1)"], "conclusions": ["argument_not_accepted(bob,arg1)"]}),
    ("inconsistent_commitment", "44, p. 337", {"A": "bob", "S": "tax_cut"},
     {"premises": ["initial_commitment(bob,tax_cut)", "opposed_commitment(bob,tax_cut)"], "assumptions": ["commitments_conflict(bob,tax_cut)"], "conclusions": ["inconsistent_commitment(bob)"]}),
    ("circumstantial_ad_hominem", "45, pp. 337-338", {"A": "bob", "S": "arg1", "C": "tax_cut"},
     {"premises": ["advocates_conclusion(bob,arg1,tax_cut)", "committed_opposite(bob,tax_cut)", "credibility_questioned(bob,arg1)"], "assumptions": ["character_issue(bob,arg1)"], "exceptions": ["inconsistency_explained(bob,arg1)"], "conclusions": ["plausibility_decreased(bob,arg1)"]}),
    ("bias", "46, p. 338", {"A": "bob", "C": "tax_cut"},
     {"premises": ["not_taken_both(bob,tax_cut)", "biased(bob)"], "assumptions": ["bias_relevant(bob,tax_cut)"], "conclusions": ["arguer_not_taken_both(bob,tax_cut)"]}),
    ("full_slippery_slope", "53, p. 341", {"C0": "c0", "C1": "c1", "C2": "cn"},
     {"premises": ["initial_presumption(c0)", "leads_to_next(c0,c1,cn)", "accept_all_steps(c0)"], "conclusions": ["not_acceptable_step(c0)"]}),
    ("plea_for_excuse", "58, p. 345", {"P": "ann", "C": "case1", "S": "fine", "R": "r", "T": "speeding", "E": "emergency"},
     {"premises": ["rule_require(r,speeding,fine)", "pleader_committed(ann,case1,speeding)", "special_circumstances(ann,case1,emergency,speeding)"], "assumptions": ["excuse(ann,case1,r,speeding,fine,emergency)"], "exceptions": ["not_excuse(ann,case1,r,speeding,fine,emergency)"], "conclusions": ["exempted(ann,case1,r,speeding,fine)"]}),
    ("memory", "60, p. 346", {"P": "ann", "M": "door_locked"},
     {"premises": ["recalls(ann,door_locked)", "believe_it(door_locked)"], "exceptions": ["based_on_false_belief(ann,door_locked)", "not_originally_believed(ann,door_locked)", "recaller_doubts(ann,door_locked)"], "conclusions": ["reasonable_believe(door_locked)"]}),
    ("illustration", "6.2, p. 315", {"Y": "tweety", "P1": "bird", "P2": "flies"},
     {"premises": ["usually_properties(bird,flies)", "has_properties(tweety,bird,flies)"], "conclusions": ["valid_rule(bird,flies)"]}),
    ("negative_value", "19.2, p. 321", {"V": "theft", "A": "hal", "G": "survive"},
     {"premises": ["judged_negative_value(theft,hal)", "negative_affects(theft,survive,hal)"], "conclusions": ["reason_retracting_goal(hal,theft,survive)"]}),
    ("epistemic_ignorance", "27, p. 328", {"A": "spy", "D": "staff", "K": "files"},
     {"premises": ["true_propositions_contained(staff,files)", "special_proposition(spy,files)", "special_in(spy,staff)", "special_not_in(spy,files)", "bivalent(staff)"], "conclusions": ["¬spy"]}),
    # the record added by PLAN.md 5.5
    ("sunk_costs_precommitment", "25, pp. 326-327", {"P": "me", "A": "finish", "T1": "t1", "T2": "t2"},
     {"premises": ["choice_at(t2,finish)", "precommitted(me,finish,t1,t2)"], "conclusions": ["should_do(me,finish)"]}),
    ("verbal_slippery_slope", "52, p. 341", {"A1": "a1", "A2": "a2", "A": "an", "P": "bald"},
     {"premises": ["has_property(a1,bald)", "also_has(a1,bald,a2)", "vague_property(bald,a1,an)", "¬has_property(an,bald)"], "conclusions": ["not_has_property(a1,bald)"]}),
]

def instantiate(statement, binding):
    return re.sub(r"\b[A-Z][A-Za-z0-9_]*\b", lambda m: binding.get(m.group(0), m.group(0)), str(statement).replace(" ", ""))

def instances(master, rep):
    S = {x["id"]: x for x in master["argument_schemes"] if isinstance(x, dict) and isinstance(x.get("id"), str)}
    for sid, entry, binding, want in INSTANCES:
        if sid not in S:
            rep.add("ERROR", "INSTANCE", f"walton.yml:{sid}", f"scheme missing (instance of {entry})"); continue
        x = S[sid]
        unbound = sorted(set(x.get("variables") or []) - set(binding))
        if unbound:
            rep.add("ERROR", "INSTANCE", f"walton.yml:{sid}", f"the instance of {entry} binds no value to {unbound}"); continue
        for f in LIST_FIELDS:
            got = sorted(instantiate(s, binding) for s in x.get(f) or [])
            exp = sorted(want.get(f, []))
            if got != exp:
                rep.add("ERROR", "INSTANCE", f"walton.yml:{sid}", f"{f} of the instance of {entry}: the record yields {got}, the book's reading is {exp}")

def cross_files(root, master, cats, rep):
    M = {x["id"]: x for x in master["argument_schemes"] if isinstance(x, dict) and isinstance(x.get("id"), str)}
    ML = master.get("language") or {}
    home = {}
    for rel, d in cats.items():
        for x in d["argument_schemes"]:
            if not isinstance(x, dict) or "id" not in x: continue
            home.setdefault(x["id"], []).append(rel)
            if x["id"] not in M:
                rep.add("ERROR", "MISSING_FROM_WALTON_YML", f"{rel}:{x['id']}", "scheme is in a category file but not in walton.yml")
            elif x != M[x["id"]]:
                diff = sorted(k for k in set(x) | set(M[x["id"]]) if x.get(k) != M[x["id"]].get(k))
                rep.add("ERROR", "DIFFERS_FROM_WALTON_YML", f"{rel}:{x['id']}", f"field(s) {diff} differ from walton.yml")
        for k, v in (d.get("language") or {}).items():
            if k in ML and ML[k] != v:
                rep.add("ERROR", "TEMPLATE_DIFFERS_FROM_WALTON_YML", rel, f"{k}: {v!r} vs walton.yml {ML[k]!r}")
    for sid in M:
        if sid not in home: rep.add("ERROR", "UNCATEGORIZED", f"walton.yml:{sid}", "scheme is in no category file")
    for sid, files in home.items():
        if len(set(files)) > 1: rep.add("ERROR", "MULTIPLE_CATEGORIES", sid, f"scheme appears in {sorted(set(files))}")

def main(argv):
    root = next((a for a in argv[1:] if not a.startswith("--")), os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
    rep = Report()
    master = check_file(os.path.join(root, "walton.yml"), "walton.yml", rep)
    cats = {}
    for p in sorted(glob.glob(os.path.join(root, "[A-Z]_*", "*.yml"))):
        rel = os.path.relpath(p, root).replace(os.sep, "/")
        d = check_file(p, rel, rep)
        if d is not None: cats[rel] = d
    if master is not None:
        regression(master, rep)
        instances(master, rep)
        cycles(master, rep)
        cross_files(root, master, cats, rep)
    if "--json" in argv:
        print(json.dumps(rep.items, indent=1, ensure_ascii=False))
    else:
        for it in rep.items:
            print(f"{it['severity']:7} {it['code']:32} {it['where']}: {it['message']}")
        e, w = len(rep.errors()), len(rep.items) - len(rep.errors())
        print(f"\n{e} error(s), {w} warning(s) in {1 + len(cats)} file(s).")
    return 1 if rep.errors() else 0

if __name__ == "__main__":
    sys.exit(main(sys.argv))
