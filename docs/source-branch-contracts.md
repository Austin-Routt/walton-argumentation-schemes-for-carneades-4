# Executable alternatives in the source compendium


Current approval status (2026-11-21): implementation-stage statements below saying review is pending are historical. The new compendium variants now have non-author acceptance in `reviews/closure-current-register.md` and individual acceptance reports. Inherited-original holds remain separate; the contracts are retained.

This review reads Walton, Reed and Macagno (2008), printed pp.312, 322, 325
and 327 of the supplied book. It does not claim independent inspection of
the earlier works listed in those entries' bibliographies. The records model
presumptive argument patterns. A generated argument does not certify a
premise's truth or make the pattern a sound universal rule of practical or
moral reasoning.

The book often prints alternatives in parentheses. Here a separate record
selects each meaningful branch; it does not mix a premise from one branch
with a conclusion from another. The named-agent forms make the printed
speaker or respondent explicit so multiple people can occur in one graph.
That is a declared application interpretation. Action terms such as
`accept(policy)`, `acquire(property)` and `buy(object_kind)` keep the intended
roles distinct. The engine does not independently verify the external type
or truth of a user-supplied term.

All new records and revised interpretations await independent approval.
The implementation author has not advanced their source stages or renewed
any approval fingerprints.

## Entries 4.5–4.7: goal, excuse and snob alternatives (p.312)

### Moral justification: actual goal and ought-goal

`moral_justification_actual_goal(R,G,P)` requires both the good group's
acceptance of policy P and R's actual goal of being good or belonging to G.
It concludes `should_do(R,accept(P))`. The separately named
`moral_justification_ought_goal` uses `ought_goal_good_group(R,G)` instead.
Neither branch proves that the other goal statement holds. In particular,
the normative claim that someone ought to have a goal does not establish
that the person actually has it. The group and respondent must agree across
the premises.

The historical `moral_justification_ad_populum` remains its actual-goal
acceptance branch for one designated respondent. Its `should_accept(P)`
conclusion contains no person identifier. It is not automatically equated
with every named person's obligation to accept P; the named forms are used
for graphs involving several respondents. P is a policy or proposition,
not an arbitrary physical object.

### Excuse: doing an action and accepting a policy

`moral_excuse_action(S,G,X)` needs universal group practice of X, the
acceptability of that same action as a norm for G, and speaker S's membership
in G. Its conclusion is `morally_justified(S,X)`. The policy alternative,
`moral_excuse_policy(S,G,A)`, needs the group's acceptance of A, its policy
norm, and the same speaker's membership. It concludes
`morally_justified(S,accept(A))`. That conclusion does not prove A itself.
An action norm is not a policy norm, and the practice of one group cannot be
combined with membership in an unrelated group.

Speaker membership uses `person_group_member/2`. The existing
`group_member/2` also includes ideas, habits, products and methods in entry
21.2, so it cannot by itself establish that the member is a speaking person.
The new person-specific predicate prevents those broad membership facts from
producing moral excuses for objects.

The historical combined record remains a single designated speaker's
disjunctive action/policy formulation. Its paired X and A identify the same
application of the source alternative, and its combined norm statement is
supplied as a whole. It does not produce either of the separately typed
conclusions above. Applications needing to distinguish whether the excuse
is for X or for acceptance of A should use the separate records; no automatic
disjunct selection or universal moral justification is claimed for the
legacy combined conclusion.

### Snob appeal: acceptance, property and purchase

The three named records are `snob_appeal_acceptance`,
`snob_appeal_property` and `snob_appeal_purchase`. Each needs an elite group,
the particular respondent's prestige goal and that respondent's matching
membership warrant. Their first premises respectively say that every group
member accepts a proposition, has a property, or possesses an object of a
specified kind. Their conclusions respectively recommend accepting,
acquiring or buying. No rule treats an object's possession as acceptance of
that object as a proposition. A group collectively having a property does
not establish that all its members have it.

The historical `snob_appeal_ad_populum` keeps its identifier and parameter
order but now additionally requires `group_accepted(G,A)`. This ordinary
premise establishes the acceptance clause actually used by its conclusion;
the inherited disjunctive `property_or_object` statement by itself is
insufficient. Its membership conditional already includes acceptance of A
as a sufficient path to membership. The explicit acceptance premise avoids
using only the property/possession clauses to recommend an unrelated belief.
The added requirement is a disclosed source-branch repair, not a new numbered
critical question, and saved graphs need a justified witness. The migration
tool flags that requirement and never invents acceptance evidence.

The legacy book case now accepts `luxury_watches_are_desirable`, a
proposition, instead of accepting `rolex`, an object. The separate purchase
case concludes `should_do(ann,buy(rolex_watch))` and explicitly refuses both
`should_accept(rolex_watch)` and truth of the object identifier.

The compendium prints no separate numbered CQ list for entries 4.5–4.7.
The positive cases test every ordinary premise by removal, and adverse cases
test wrong groups, wrong agents and mismatched branch warrants.

Whole-library composition also includes the distinct `pop_scheme`, whose
premise `group_accepted(G,A)` supports A. Therefore a full-library graph can
contain an argument for A even though the selected excuse or snob record
concludes only an action recommendation. The fixtures keep full-library
positive integration and separately select the live target record for the
negative control that it does not itself prove A. No copied scheme formula
or weakened absence expectation substitutes for that distinction.

## Entry 20: small sacrifice without an invalid converse (p.322)

The old `sacrifice` is the great-sacrifice branch. The book also parenthesizes
small sacrifice and small value, but its printed conditional expressly
describes great sacrifice and greater value; “vice versa” does not by itself
license small sacrifice implying small value. `sacrifice_small` therefore
requires an explicit ordinary application warrant,
`small_sacrifice_warrant(S,A,V)`, as well as the sacrifice and evidence that
it was small. Its conclusion is `small_sacrifice_value(A,V)`.

This is a declared conditional interpretation of the printed small branch.
No rule derives its warrant from the great-sacrifice conditional, an absence
of evidence of great sacrifice, or a warrant about another object. The value
and sacrifice identifiers refer to the same evaluation. The conclusion is
qualitative; no numerical value, inverse proportionality, or total ordering
of all sacrifices is computed. The explicit warrant is an application
premise, not a numbered source CQ.

## Entry 22.6: badness, lesser badness and both reclassifications (p.325)

### Variant 1: good, bad and less bad

`end_and_means_1` keeps the good-end/good-means branch.
`end_and_means_bad` represents the bad-end/bad-means parenthesis.
`end_and_means_less_bad` represents a means that is less bad insofar as it
serves a good end. That last form requires `bad(X)` as an explicit application
precondition, making the adverse baseline implicit in “less bad” visible.
`less_bad_as_means(X,Y)` compares evaluation of this means with and without
its good purpose; it is qualitative and tied to Y. It is not a synonym for
`good(X)`, a denial of `bad(X)`, or a number computed from utilities.

Loading all alternatives can construct both the good-means argument and the
less-bad argument when their premises hold. They retain separate provenance.
Applications choosing one interpretation should select the corresponding
record with `graph_from_library.py`. This is not a hidden preference for the
stronger valuation. The source itself gives alternatives separated by a
slash; it does not supply a utility calculus or an automatic preference rule.

### Variant 2: evaluation context and reverse reclassification

`end_and_means_2(X,C)` now uses the same context C in `means_end(X,C)`,
`more_important_ends(C)` and `ends_important(X,C)`. The importance comparison
is an ordinary premise supplied for this evaluation, never a global default.
The reverse record, `end_and_means_reverse(X,C)`, instead requires
`end_means(X,C)` and concludes `¬ends_important(X,C)` under the same warrant.

These conclusions are comparative evaluative assessments. The positive term
says that X warrants assessment as highly important in C; its explicit
negation is a supported rejection of that assessment. This is the declared
executable reading of the source's “should (should not) be regarded” language,
given the ordinary premise comparing the importance of ends and means. It is
not an obligation to perform or refrain from an actor's action, and it does
not identify not-obligatory with obligatory-not. The reverse conclusion needs
its own `end_means(X,C)` premise and the same-context importance warrant.
Failure to establish the positive assessment alone does not yield the
negative assessment. The shared display template states the evaluative
meaning explicitly; no new global issue or deontic bridge is installed.

Contrary reclassifications within the same context produce opposed
arguments in the ordinary negation issue. Equal support does not establish
either position under the default preponderance standard. Contrary
reclassifications in different contexts do not oppose or lend warrants to
one another. Saved one-argument forms need a deliberate context assignment;
the migrator leaves them for review rather than inventing a context.

### Variant 3: not easy is explicit negative evidence

`end_and_means_3` retains the easy-means/good-end branch.
`end_and_means_difficult` uses `¬easy_to_achieve(X)` and concludes `bad(Y)`
for the same means-to-end relation. Not observing an easy-means assertion is
not evidence that the means is difficult. Another means' difficulty also
does not transfer. These valuation patterns remain presumptive and the
source supplies no separate numbered question list for them.

## Entry 26.2: both practical presumptions and all three CQs (p.327)

`negative_practical` is the false-presumption branch: the adverse result B
of assuming A when false is more serious than C, the adverse result of
assuming not-A when true. `negative_practical_true` requires the opposite
strict ranking, C more serious than B, and concludes
`act_on_presumption_true(A)`. Neither conclusion establishes A or not-A.
Evidence of not knowing A and having to choose a presumption is explicit.
Missing knowledge facts or missing strict ordering are not completed by a
closed-world assumption. Identical consequences without a strict ordering
do not choose a branch.

Binary-choice applications explicitly declare an issue containing the two
positive choice statements for the same A. The worked examples
`docs/examples/practical_presumption.yml` and
`docs/examples/practical_presumption_conflict.yml` declare that issue and PE.
The first selects the true presumption; the second preserves both arguments
under contradictory seriousness judgments and accepts neither choice.

There is no automatic global issue scheme for these choices. Negating a
recommendation does not assert the other recommendation, and the engine's
single-issue model would merge overlapping choice and negation issues. A
true-choice position and a negated false-choice position can be compatible,
so treating every merged position as a different exclusive alternative would
misrepresent that case. The explicit binary protocol therefore uses only its
two positive choice statements; arbitrary explicit negations of those
statements are outside that issue protocol. The worked graphs make this
scope visible, while a separate control verifies that the general library
does not impose that unsafe global merge. A larger dialogue model must give
its own logically specified treatment of such negations.

The three printed CQs have the following executable/static responsibilities:

| Source question | Declared responsibility and observable behavior |
| --- | --- |
| CQ1, progress of the evidence search | `sufficient_evidence(A)` is the existing challengeable assumption. An unanswered challenge removes its default support and prevents acceptance; explicit support answers the challenge. A challenge about another question cannot do so. Search adequacy is not proof of A. |
| CQ2, ultimate probandum and who bears the burden | A binary-choice application explicitly declares the ultimate issue as these two positive practical choices for the same case-qualified A. The party advancing a choice supplies its ordinary premises; the CQ1 default has the explicitly declared challenge burden. The worked conflict graph places both branches in that declared issue. The bare library does not infer an institutional burden or silently install the binary protocol in every graph. |
| CQ3, strength of proof required | The application selects the issue's Carneades proof standard; this library uses the engine's default PE unless an explicit issue supplies another supported standard. The equal-opposition case verifies the default does not accept both equally supported choices. The engine proof-standard tests establish the available evaluation semantics. |

CQ2 and CQ3 are graph policy, not hidden extra facts or assertions that a
dialogue protocol has run. This static library does not infer which party
has an institutional burden, decide a legal standard, or simulate turn
taking. A dialogue application must supply those choices and translate its
current evidence state before evaluation. A is case-qualified when otherwise
identical propositions occur in different inquiries or decision contexts.

## Verification and review boundary

`tools/vv/cases/closure_source_branches.yml` contains hand-specified positive
and adverse cases. It loads the live library, names the expected arguments
and labels, and checks missing premises, branch types, person/group/context
isolation, explicit negative evidence, search challenges and opposed choices.
The revised historical cases are in `book.yml`. Migration tests ensure that
the new witness and evaluation context are not fabricated. The source ledger
and `docs/reviews/closure-source-proposals.json` identify every new record's
parent entry and exact formula for an independent reviewer.

Passing these tests establishes the stated executable distinctions, not that
the program has independently verified the premises of a real application.
The tests and contracts do not grant their author authority to approve the
authored interpretations.

Any copied fixed issue also requires its predicates' language
templates in every graph that copies that issue. Both category generation
and selected-graph generation now retain those dependencies even when none
of the selected argument schemes uses them. Python validation checks missing
issue-pattern language, and a mutation control removes a required template
to prove that omission is detected. The proposed global presumption issue
first exposed this generator omission; the global issue was then withdrawn
after the independent single-issue/negation review above.
