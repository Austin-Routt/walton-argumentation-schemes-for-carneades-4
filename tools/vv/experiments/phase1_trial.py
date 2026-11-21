#!/usr/bin/env python3
"""Trial edits to the scheme library, to measure a plan before anyone commits to
it. Run it on a COPY of the repository.

Usage:  python3 phase1_trial.py <copy_of_repo_root> [--final | --converged | --keep-pair]

  (no option)   library-only route: every engine limit is worked around in the
                scheme files (range restriction, true/1, is/2, negated premises
                replaced by positive predicates, the composition/division pair
                split, hyphenated ids renamed).
  --keep-pair   the same, but composition_generic and division_generic keep
                concluding each other's premise.
  --converged   the route of the converged plan: only what is wrong in the
                scheme files themselves is changed, following the book where it
                says something (see README.md here). Negated premises, the
                composition/division pair and the scheme ids stay as they are;
                they are left to the engine.
  --final       the same nine schemes as --converged, with the edits of the
                final plan. Where a statement the library already has can do
                the binding, it is used (advocates_conclusion/3, has/2); the
                regulative rule and the pop scheme follow the second review's
                book-based patch; the sorites sequence names its property, as
                the third review asked. The third review proposed no edits of
                its own; the conditions it set for them are in the plan.

It edits walton.yml, and in the first two modes categories.yml and the
regression table of tools/validate_schemes.py, in place. Regenerate the category
files afterwards with the repository's own tools/build_category_files.py. See
README.md in this directory for the commands and for what was measured.

New predicate names and texts are PLACEHOLDERS chosen to test the mechanics.
What each scheme should say is a content decision for the library's owner.
The script stops with an AssertionError, before writing anything, if the library
no longer looks the way it did at commit 7ad5617.
"""
import re
import sys

root = sys.argv[1]
opts = sys.argv[2:]
FINAL = "--final" in opts
CONVERGED = "--converged" in opts
# --keep-pair leaves composition_generic and division_generic as they are: each
# concludes the other's premise, which is what the book's two schemes do
# (9.1 and 10.1, pp. 316-317) and what the regression table of
# tools/validate_schemes.py requires. The loop between them then needs the
# engine patch in this directory.
KEEP_PAIR = "--keep-pair" in opts or CONVERGED or FINAL

path = root + "/walton.yml"
lines = open(path, encoding="utf-8").read().split("\n")

lang_start = next(i for i, l in enumerate(lines) if re.match(r"^language:\s*$", l))
schemes_start = next(i for i, l in enumerate(lines) if re.match(r"^argument_schemes:\s*$", l))
head, lang, rest = lines[:lang_start + 1], lines[lang_start + 1:schemes_start], lines[schemes_start:]


# ---------------------------------------------------------------- language
def key_of(line):
    m = re.match(r"^  ([^\s#][^:]*?):\s", line)
    return m.group(1).strip().strip("'\"") if m else None


def lang_remove(key):
    hits = [i for i, l in enumerate(lang) if key_of(l) == key]
    assert len(hits) == 1, (key, hits)
    del lang[hits[0]]


def lang_add(key, text):
    assert all(key_of(l) != key for l in lang), key
    new = '  %s: "%s"' % (key, text)
    for i, l in enumerate(lang):
        k = key_of(l)
        if k is not None and k > key:
            lang.insert(i, new)
            return
    lang.append(new)


# ---------------------------------------------------------------- schemes
starts = [i for i, l in enumerate(rest) if re.match(r"^  - id:\s*\S", l)]


def block_range(sid):
    for n, i in enumerate(starts):
        if re.match(r"^  - id:\s*%s\s*$" % re.escape(sid), rest[i]):
            return i, (starts[n + 1] if n + 1 < len(starts) else len(rest))
    raise SystemExit("no scheme " + sid)


def find_stmt(sid, stmt):
    a, b = block_range(sid)
    hits = [i for i in range(a, b) if re.match(r"^\s*-\s*%s\s*$" % re.escape(stmt), rest[i])]
    assert len(hits) == 1, (sid, stmt, hits)
    return hits[0]


def reindex():
    global starts
    starts = [i for i, l in enumerate(rest) if re.match(r"^  - id:\s*\S", l)]


def replace_stmt(sid, old, new):
    i = find_stmt(sid, old)
    rest[i] = rest[i].replace(old, new)


def add_after(sid, anchor, new):
    i = find_stmt(sid, anchor)
    indent = re.match(r"^(\s*-\s*)", rest[i]).group(1)
    rest.insert(i + 1, indent + new)
    reindex()


def remove_stmt(sid, field, stmt):
    """Remove one statement from a list field; drop the field if it becomes empty."""
    a, b = block_range(sid)
    f = next(i for i in range(a, b) if re.match(r"^    %s:\s*$" % field, rest[i]))
    items = []
    j = f + 1
    while j < b and re.match(r"^\s*-\s", rest[j]):
        items.append(j)
        j += 1
    hits = [i for i in items if re.match(r"^\s*-\s*%s\s*$" % re.escape(stmt), rest[i])]
    assert len(hits) == 1, (sid, field, stmt)
    if len(items) == 1:
        del rest[f:f + 2]
    else:
        del rest[hits[0]]
    reindex()


def set_variables(sid, old, new):
    a, b = block_range(sid)
    hits = [i for i in range(a, b) if re.match(r"^    variables:\s*\[%s\]\s*$" % re.escape(old), rest[i])]
    assert len(hits) == 1, (sid, old)
    rest[hits[0]] = rest[hits[0]].replace("[" + old + "]", "[" + new + "]")


def rename_id(old, new):
    a, _ = block_range(old)
    rest[a] = rest[a].replace(old, new)


RENAMED = {}
table_edits = []

if FINAL:
    # ---- X-RANGE: bind what the book leaves implicit, and nothing else ----
    # Generic ad hominem (42, p. 336): the attacked argument appears only in the
    # conclusion; the sibling schemes 43 and 45 open with a premise saying that
    # the person advocates it. The library already has that statement for the
    # siblings, with the argument's conclusion as a third place, so no new
    # predicate is needed. (The second review first proposed a new two-place
    # advocates/2 and withdrew it in its third round for this reason.)
    set_variables("ad_hominem", "A,S", "A,S,C")
    add_after("ad_hominem", "bad_character(A)", "advocates_conclusion(A,S,C)")

    # Practical reasoning from a necessary and from a sufficient condition
    # (22.2, 22.3, pp. 323-324): the premises are in the first person, so the
    # goal premise names the agent; the last question asks about other goals
    # without naming one, so that statement loses its place for it.
    for sid in ("practical_necessary_condition", "practical_sufficient_condition"):
        set_variables(sid, "X,A,A2,P", "X,A,P")
        replace_stmt(sid, "goal_is_action(A)", "goal_of_agent(P,A)")
        replace_stmt(sid, "conflict_goals(P,A2,A)", "conflicting_goals(P,A)")
    # New names, not the old ones at a new arity: a fact left over at the old
    # arity can stall a rule until the limit (issue E-MATCH).
    lang_remove("goal_is_action/1")
    lang_add("goal_of_agent/2", "The goal of %s is to bring about %s.")
    lang_remove("conflict_goals/3")
    lang_add("conflicting_goals/2", "%s has other goals that might conflict with realizing %s.")

    # Regulative rule (55.3, p. 343): the agent appears in the conclusion and in
    # the rebuttal factor only. For "S must do A" to follow from "A is
    # obligatory", S has to be someone the rule applies to: that instantiation
    # is the premise added here. The excuse is S's, so it gets its subject; the
    # warrant is general and stays without one.
    add_after("regulative_rule", "do_obligator", "addressee_of(A,X)")
    lang_add("addressee_of/2", "%s is subject to the rule concerning %s.")
    replace_stmt("regulative_rule", "excuse", "excuse_or_overriding_duty(A,X)")
    lang_remove("excuse/0")
    lang_add("excuse_or_overriding_duty/2", "%s has an adequate excuse, or an overriding duty, for not doing %s.")

    # Sorites slippery slope (51, pp. 340-341): the sequence of sub-arguments
    # is what ties the conclusion's individual to the clear case. It gets its
    # two ends and the property it carries, so that a sequence for another
    # property, or between other individuals, does not serve (third review).
    replace_stmt("sorites_slippery_slope", "reapplication_sequence", "reapplication_sequence(A2,A,P)")
    lang_remove("reapplication_sequence/0")
    lang_add("reapplication_sequence/3", "A sequence of modus ponens subarguments leads from %s, in the clear area, through the grey area to %s, for the property %s.")

    # ---- L-TRUE: pop scheme (4.1, p. 311) ----
    # The book gives the conclusion two wordings, "A is true" or "you should
    # accept A". The first is the statement A itself, as in expert_opinion.
    replace_stmt("pop_scheme", "true(A)", "A")
    lang_remove("true/1")

    # ---- X-IS: group and its members, variant 2 (21.2, p. 322) ----
    # "M is Q" says of the member what variant 1 says of the group with has/2.
    # The premises of the two variants use other statements, so no cycle arises.
    replace_stmt("group_2", "is(M,Q)", "has(M,Q)")
    # is/2 works as a premise today; it is renamed because `is` is an operator
    # of the rule engine. The new name is a PLACEHOLDER.
    replace_stmt("rule_physical_world_1", "is(W,D)", "is_a(W,D)")
    lang_remove("is/2")
    lang_add("is_a/2", "%s is %s.")

    # ---- Argument for an exceptional case (56, p. 344) ----
    # The major premise is the general rule, instantiated for the case of the
    # minor premise and the conclusion. The assumption that repeats the minor
    # premise goes; the scheme's three questions are conditional on one another
    # and are left to the content phase.
    set_variables("exeptional", "A,X", "A")
    replace_stmt("exeptional", "rule_can_waived(X)", "rule_can_waived(A)")
    remove_stmt("exeptional", "assumptions", "case_exception(A)")
elif CONVERGED:
    # ---- X-RANGE: bind every variable, following the book where it speaks ----
    # Generic ad hominem (42, p. 336): the argument appears only in the
    # conclusion. Its sibling schemes 43 and 45 (pp. 336-337) open with a premise
    # saying that the person advocates the argument; the library already has
    # that statement, with the argument's conclusion as a third place.
    set_variables("ad_hominem", "A,S", "A,S,C")
    add_after("ad_hominem", "bad_character(A)", "advocates_conclusion(A,S,C)")

    # Regulative rule (55.3, p. 343): the agent appears only in the conclusion
    # and in the rebuttal. The book gives no premise for it. PLACEHOLDER.
    add_after("regulative_rule", "obligatory(X)", "rule_applies_to(X,A)")
    lang_add("rule_applies_to/2", "The rule about %s applies to %s.")

    # Practical reasoning, necessary and sufficient condition (22.2 and 22.3,
    # pp. 323-324): the premises are in the first person; the questions name the
    # agent and "goals other than" the goal. The agent is bound by the goal
    # premise; the other goals are not named.
    for sid in ("practical_necessary_condition", "practical_sufficient_condition"):
        set_variables(sid, "X,A,A2,P", "X,A,P")
        replace_stmt(sid, "goal_is_action(A)", "goal_of_agent(P,A)")
        replace_stmt(sid, "conflict_goals(P,A2,A)", "conflicting_goals(P,A)")
    lang_remove("goal_is_action/1")
    lang_add("goal_of_agent/2", "The goal of %s is to bring about %s.")
    lang_remove("conflict_goals/3")
    lang_add("conflicting_goals/2", "%s has some other goals that conflict with %s.")

    # Sorites slippery slope (51, pp. 340-341): the individual of the conclusion
    # is tied to the clear case only by the sequence of sub-arguments. Giving
    # that premise its two ends binds it, and removes a statement without
    # variables.
    replace_stmt("sorites_slippery_slope", "reapplication_sequence", "reapplication_sequence(A2,A)")
    lang_remove("reapplication_sequence/0")
    lang_add("reapplication_sequence/2", "There is a sequence of modus ponens subarguments from %s to %s, linking premises and conclusions from the clear area through the grey area.")

    # ---- L-TRUE ----
    replace_stmt("pop_scheme", "true(A)", "holds_true(A)")
    lang_remove("true/1")
    lang_add("holds_true/1", "%s is true.")

    # ---- X-IS: group and its members, variant 2 (21.2, p. 322) ----
    # "M is Q" there says what variant 1 says of the group with has/2.
    replace_stmt("group_2", "is(M,Q)", "has(M,Q)")
    replace_stmt("rule_physical_world_1", "is(W,D)", "is_a(W,D)")
    lang_remove("is/2")
    lang_add("is_a/2", "%s is %s.")

    # ---- Argument for an exceptional case (56, p. 344) ----
    # The major premise is a general rule and the minor premise and conclusion
    # are about one case; the library gave the major premise a variable of its
    # own. The assumption repeats the minor premise (CQ1); the book calls a
    # question that only asks whether a premise is true redundant (p. 374).
    set_variables("exeptional", "A,X", "A")
    replace_stmt("exeptional", "rule_can_waived(X)", "rule_can_waived(A)")
    remove_stmt("exeptional", "assumptions", "case_exception(A)")
else:
    # X-RANGE: bind every variable --------------------------------------------
    add_after("ad_hominem", "bad_character(A)", "puts_forward(A,S)")
    lang_add("puts_forward/2", "%s puts forward argument %s.")

    add_after("regulative_rule", "obligatory(X)", "rule_applies_to(X,A)")
    lang_add("rule_applies_to/2", "The rule about %s applies to %s.")

    for sid in ("practical_necessary_condition", "practical_sufficient_condition"):
        set_variables(sid, "X,A,A2,P", "X,A,P")
        add_after(sid, "goal_is_action(A)", "acting_agent(P)")
        replace_stmt(sid, "conflict_goals(P,A2,A)", "conflicting_goals(P,A)")
    lang_add("acting_agent/1", "%s is the agent who deliberates.")
    lang_add("conflicting_goals/2", "%s has some other goals that conflict with %s.")
    lang_remove("conflict_goals/3")

    add_after("sorites_slippery_slope", "all_has(A2,P,A1)", "sequence_reaches(A1,A)")
    lang_add("sequence_reaches/2", "The sequence that starts with %s reaches %s.")

    # L-TRUE ---------------------------------------------------------------------
    replace_stmt("pop_scheme", "true(A)", "holds_true(A)")
    lang_remove("true/1")
    lang_add("holds_true/1", "%s is true.")

    # X-NEG: positive predicates instead of negated premises ----------------------
    replace_stmt("from_antimodel", "¬prestigious_individual(A,B)", "not_prestigious_individual(A,B)")
    lang_add("not_prestigious_individual/2", "Individual %s is not prestigious for -admired by- individuals %s.")
    for sid in ("opposition-descriptive_2", "opposition-descriptive_3"):
        replace_stmt(sid, "¬presents(A,P)", "not_presents(A,P)")
    lang_add("not_presents/2", "%s does not present %s.")
    replace_stmt("ignorance", "¬known(S)", "not_known(S)")
    lang_add("not_known/1", "%s is not known to be true.")
    lang_remove("known/1")  # no scheme uses it any more

    # X-IS -----------------------------------------------------------------------
    replace_stmt("group_2", "is(M,Q)", "is_a(M,Q)")
    replace_stmt("rule_physical_world_1", "is(W,D)", "is_a(W,D)")
    lang_remove("is/2")
    lang_add("is_a/2", "%s is %s.")

    # X-LOOP ---------------------------------------------------------------------
    remove_stmt("exeptional", "assumptions", "case_exception(A)")
    if not KEEP_PAIR:
        replace_stmt("composition_generic", "has_property(A,P)", "whole_has_property(A,P)")
        lang_add("whole_has_property/2", "%s as a whole has the property %s.")
        table_edits.append(
            ('fns(x, "premises") == ["all_parts_propert"] and fns(x, "conclusions") == ["has_property"]',
             'fns(x, "premises") == ["all_parts_propert"] and fns(x, "conclusions") == ["whole_has_property"]'))

    # L-ID -----------------------------------------------------------------------
    RENAMED = {
        "opposition-descriptive_1": "opposition_descriptive_1",
        "opposition-descriptive_2": "opposition_descriptive_2",
        "opposition-descriptive_3": "opposition_descriptive_3",
        "opposition-normative": "opposition_normative",
    }
    for old, new in RENAMED.items():
        rename_id(old, new)
    # tools/validate_schemes.py names this id in its regression table
    table_edits.append(('("opposition-normative", lambda x:', '("opposition_normative", lambda x:'))

cpath = root + "/categories.yml"
c = open(cpath, encoding="utf-8").read()
for old, new in RENAMED.items():
    assert len(re.findall(r"(?<![\w-])%s(?![\w-])" % re.escape(old), c)) == 1, old
    c = re.sub(r"(?<![\w-])%s(?![\w-])" % re.escape(old), new, c)

vpath = root + "/tools/validate_schemes.py"
v = open(vpath, encoding="utf-8").read()
for old, new in table_edits:
    assert v.count(old) == 1, old
    v = v.replace(old, new)

# Every edit matched; only now are the files written.
open(path, "w", encoding="utf-8").write("\n".join(head + lang + rest))
open(cpath, "w", encoding="utf-8").write(c)
open(vpath, "w", encoding="utf-8").write(v)
print("trial edits applied to", root, "(final route)" if FINAL else "(converged route)" if CONVERGED else "(library-only route)")
