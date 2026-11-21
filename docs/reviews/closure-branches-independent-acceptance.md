# Independent acceptance of seventeen source-branch interpretations

Date: 2026-11-21. Reviewer: ChatGPT independent closure scope reviewer.
Reviewed source implementation: `7b443b6c78c23dec3bd8b2928ffb358875bda307`
in the source worktree, integrated as `22bfa8d` and inspected in the combined
tree at `65f0f01`. This reviewer did not author those branch formulas or their
source-case fixtures.

## Decision and evidence boundary

**Approve seventeen declared compendium interpretations at `example-passed`:**
eleven new records and six existing parent records. Each has its own exact
semantic fingerprint and existing example references in
`independent-approvals.json`. The approvals concern the supplied 2008
compendium and the disclosed executable interpretation, not endorsement by
the source authors or independent certification of every earlier work in the
compendium's bibliography.

I read all twenty-one relevant proposal rows, formulas, display templates,
ledger contracts, and positive/adverse fixtures. I compared printed pp.312,
322,325,327 with the supplied book and visually inspected all four page
facsimiles. Images are retained as `scope-audit/book-322.png`,
`book-332.png`, `book-335.png`, and `book-337.png`; these filenames use PDF page
numbers, ten above the corresponding printed pages. The source did not gain
new numbered critical questions merely because an application warrant or
ordinary premise is now explicit.

The remaining four interpretations are not silently included in this
approval. The two practical-presumption records are assigned to the
integration reviewer for the final dialogue/burden contract. I authored the
narrow evaluative-template clarification for the two importance records, so
the integration reviewer independently assesses those two as well. See
`closure-importance-clarification.md`; I grant no approval to my own change.

## Record-by-record acceptance

| Record | Printed source | Accepted interpretation and discriminating evidence |
|---|---|---|
| `moral_justification_actual_goal` | 4.5, p.312 | The respondent actually has the goal. The same group appears in the acceptance warrant and goal; an unrelated group does not support this recommendation. Named respondent and `accept(P)` retain the source roles. |
| `moral_justification_ought_goal` | 4.5, p.312 | The parenthetical normative goal is an explicit premise. Its conclusion recommends policy acceptance without manufacturing an actual goal. Missing and wrong-group premises block the instance. |
| `moral_justification_ad_populum` | 4.5, p.312 | The unchanged legacy actual-goal branch is valid for one designated respondent. Its positive example and premise-removal controls pass. It is not equated with every named person's obligation; the two executable named variants supply the previously absent branch/role coverage. |
| `moral_excuse_action` | 4.6, p.312 | Universal practice, the same action norm, and the speaker's person-specific membership support the action excuse. Another group, a policy-only premise, or membership of a product cannot substitute. |
| `moral_excuse_policy` | 4.6, p.312 | Policy acceptance and its norm support the same named person's accepting action. A different policy norm is insufficient. A selected live-record control verifies that this record does not itself prove the policy proposition. |
| `moral_justification_ad_populum_excuse` | 4.6, p.312 | The legacy combined disjunctive statement remains a single designated speaker's application. Its conclusion remains combined; it does not select one typed alternative. The two explicit action/policy variants now execute the distinction that this combined record leaves implicit. |
| `snob_appeal_acceptance` | 4.7, p.312 | Universal group acceptance, elite status, this respondent's prestige goal, and the matching acceptance-to-membership conditional support acceptance. A purchase conditional does not answer this warrant. |
| `snob_appeal_property` | 4.7, p.312 | The first premise concerns every member's property, not a collective property of the group. The acquisition warrant and recommendation share respondent, property, and group. The conclusion does not assert that acquisition has already occurred. |
| `snob_appeal_purchase` | 4.7, p.312 | The object-kind and purchase branch is distinct from accepting a proposition. Another person's membership conditional cannot supply this person's missing warrant. The Rolex-kind example neither asserts an object as a proposition nor recommends accepting that object. |
| `snob_appeal_ad_populum` | 4.7, p.312 | The additional ordinary `group_accepted(G,A)` premise justifies selecting the acceptance branch. The inherited disjunctive property/possession premise alone no longer suffices. The repaired book example uses a proposition, while the separate purchase record handles objects. |
| `sacrifice` | 20, p.322 | The existing great-sacrifice branch retains its explicit conditional and qualitative great-value conclusion. Its positive example and missing-premise controls pass. The separately executable small branch now supplies the printed alternative without using an invalid converse. |
| `sacrifice_small` | 20, p.322 | The printed small alternative is an explicitly conditional application interpretation: a same-evaluation small-sacrifice warrant, actual sacrifice, and smallness evidence are required. A great-sacrifice warrant, a different object's warrant, or absence of evidence of greatness is insufficient. |
| `end_and_means_1` | 22.6.1, p.325 | The original good-end/good-means argument matches one printed branch. New bad and less-bad records supply the previously absent alternatives; this record does not claim to calculate a universal utility ordering. |
| `end_and_means_bad` | 22.6.1, p.325 | The bad-end parenthesis yields a defeasible bad-means assessment. Badness of a different end does not transfer to this means. |
| `end_and_means_less_bad` | 22.6.1, p.325 | The slash-separated less-bad alternative is a qualitative assessment relative to the same good purpose. An explicit adverse baseline makes the comparison meaningful. The selected record neither concludes goodness nor negates badness; another good purpose cannot supply this one. |
| `end_and_means_3` | 22.6.3, p.325 | The easy-means/good-end branch remains distinct. The new difficult branch executes the printed negative alternative; neither branch infers easiness or difficulty from missing information. |
| `end_and_means_difficult` | 22.6.3, p.325 | Explicit not-easy evidence for the same means supports a bad-end argument. Missing easiness evidence or difficulty of another means is insufficient. |

## Interpretation details that affect the verdict

The named-agent forms are intentional adaptations of printed first- and
second-person roles. They do not reinterpret those pronouns as arbitrary
objects. `person_group_member/2` is necessary because the older generic
`group_member/2` also covers ideas, habits, products, and methods in entry21.2.
The adverse product-member fixture tests that concrete role distinction.

The small-sacrifice warrant and the prior-badness premise in the less-bad
branch are disclosed application premises. They are defensible ways to make
the source's compressed alternatives executable. They are not represented
as extra numbered source questions, a numerical valuation function, or
conclusions derivable from lack of contrary evidence.

The complete library also contains independent argument patterns. For
example, `pop_scheme` can support a proposition from universal group
acceptance, even when a snob or excuse argument concludes only an acceptance
action. Therefore the review checks both full-library positive integration
and selected **live canonical records** for the narrower negative claim that
a particular record does not itself prove the proposition. No source formula
was copied into a special test-only substitute. Likewise, good and less-bad
arguments can coexist with distinct provenance if both alternatives are
loaded; the less-bad record alone does not become the good branch.

The seven legacy forms considered here are deliberately not a substitute for
implementing omitted branches. The additional actual/ought, action/policy,
acceptance/property/purchase, small, bad/less-bad, reverse, and difficult forms
are present and have concrete executable cases. Six legacy parent records
are approved in this report; the seventh, `end_and_means_2`, is independently
reviewed by integration after my evaluative clarification.

## Independent execution and discovered integration fault

The first selected rerun exercised `book.yml`, all49 source-branch cases,
the independently authored second-book case file, and the worked documents
with the corrected API/CLI expectation oracle. **All408 registered source-case
checks passed.** The documents produced7/8 passes: the old practical-reasoning
example had a stale generated prefix after predicate changes. The combined
result was415/416, and that failure is retained rather than hidden. The
integration reviewer separately regenerated that example and the other
current examples; this report does not claim to have performed their final
integrated rerun.

After the two importance templates were clarified, the exact `book.yml`,
`closure_source_branches.yml`, and new `closure_assessment.yml` selection
passed **340/340 API/CLI checks**. These are evidence for the declared source
cases, not the final whole-repository matrix result.

| Evidence log | SHA-256 |
|---|---|
| `scope-audit/branch-independent.log` | `ca4cadfeaa5fa0f46cf3e19067936387a9e7ab7e619c61eac3cf034e7be11cba` |
| `scope-audit/branch-independent-assessment.log` | `b2605cc94bb33926c4b39ede507d389b398dbf96d19f83f06c3cdf8168881f23` |

Both used the local Go1.24.13 toolchain, `GOTOOLCHAIN=local`,
`GOFLAGS=-mod=mod`, and strict known-issue handling. The exact source ledger
and approval registry are checked after these seventeen promotions. Final
clean-copy verification, the second toolchain, release artifacts, and
publication decisions remain the integration reviewer's responsibility.

After these promotions the exact approval-integrity checker passes for all
143 records, with117 approvals present in this worktree. The source/migration
gate also passes **777/777 registered checks**. Its log is
`scope-audit/branch-approval-source-gate.log`, SHA-256
`2edae1a3094cefdb013be567e2f08e2d2965e455d3d1f1901fb1d2af080581a3`.
These intermediate totals precede the integration reviewer's remaining
eleven practical, position-to-know, ignorance, and importance decisions.
