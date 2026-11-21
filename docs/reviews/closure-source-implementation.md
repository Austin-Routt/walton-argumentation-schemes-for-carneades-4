# Source-closure implementation and verification record

Date: 2026-11-21. Implementation author: `closure_source_changes`.
Baseline public repository: `e6d2d6f427c6ec0447cae514810c256903b15779`.
This is an implementation report for independent review. It grants no
source approvals and is not an integrated release qualification.

## Evidence inspected

The supplied *Argumentation Schemes* PDF has SHA-256
`d010348014e8b80fa561b68e5fac732cce0b3905fb0b1d522dd6fd2578a69634`.
Printed page numbers are used here; the supplied PDF's page ordinal is ten
higher. Source text and the rendered printed pages were checked for the
parenthetical alternatives and directional wording. Relevant pages are
312 (4.5–4.7), 316–317 (9.2, 10.2), 319–320 (16), 322 (20), 325
(22.6.1–3), 327 (26.2) and 332 (33.1–33.3).

The source-grounded test cases instantiate the printed patterns with
hand-chosen examples and adverse facts. Their conventional `/book` names
do not assert that the book prints those particular example stories.
The earlier books in the compendium's bibliographies are not certified as
independently inspected by this review.

## Decisions and changes

### Consequence likelihood and vagueness

Commit `746b5ab2f267033e052185f4cee9fbcaae5a6529` harmonizes entry 33.3
with the shared likelihood-adequacy predicate of 33.1–33.2. The common CQ
does not mandate the former categorical strong-likelihood threshold.
Old positive, negative and questioned threshold judgments require manual
reassessment; the migration is not a blind predicate rename.

The implementer's entry 16 CQ3 wording correctly concerns whether the
vagueness is a problem for the dialogue. The unchanged formula has an
explicit static evidence burden: required precision evidence is a premise,
affirmative harmlessness is an exception, merely asking the exception does
not prove it, and evidence about another argument/property cannot defeat
this argument. This is independent concurrence with that wording and an
authored test/contract proposal; a separate reviewer must approve the
complete interpretation. See `docs/source-closure-contracts.md`.

### Part/whole coverage and necessary versus sufficient conditions

Commit `611e0c7d81501538ac8cb9750b748681456e78f8` adds six records.
Commit `04d586c9c4272ae0a3c2a0b7490a8ecb84b24e88` then applies the
independent reviewer's requirement that the historical three-part positive
division form also require an explicit sufficient-direction warrant.

| Entry | New record | Distinction made executable |
| --- | --- | --- |
| 9.2 | `part_in_whole_physical` | Physical part, separate from taxonomic species |
| 9.2 | `part_in_whole_species_degree` | Lesser degree for species on a stated common scale |
| 9.2 | `part_in_whole_physical_degree` | Lesser degree for a physical part on that scale |
| 9.2 | `part_in_whole_smaller_degree` | Smaller-part alternative with explicit relevant size evidence |
| 10.2 | `whole_into_parts_selected` | Any selected member of an explicit finite partition, with a sufficient warrant |
| 10.2 | `whole_into_parts_negative` | Necessary only-if direction plus a supported universal negative over the complete partition |

The aggregate contract does not hard-code three parts, infer a complete set
from a partial list, infer universal negatives from absent positive facts,
or reverse the source's “only if.” The application supplies the finite,
nonempty, complete partition and its evidence. The engine checks consistent
identifiers and rule behavior; it cannot certify the external truth of that
completeness assertion. The positive direction is a disclosed application
adaptation with an ordinary sufficient warrant.

### Remaining printed alternatives

The following thirteen records complete the assigned parenthetical branch
work. Exact formulas, source mapping and example names are in
`docs/reviews/closure-source-proposals.json`; the contracts are in
`docs/source-branch-contracts.md`.

| Entry | New record | Distinction made executable |
| --- | --- | --- |
| 4.5 | `moral_justification_actual_goal` | Named respondent's actual goal |
| 4.5 | `moral_justification_ought_goal` | Ought-goal without asserting an actual goal |
| 4.6 | `moral_excuse_action` | Named speaker and the particular action norm |
| 4.6 | `moral_excuse_policy` | Named speaker and policy acceptance, without asserting policy truth |
| 4.7 | `snob_appeal_acceptance` | Acceptance of a proposition/policy |
| 4.7 | `snob_appeal_property` | Acquisition of a property held by all group members |
| 4.7 | `snob_appeal_purchase` | Purchase of a kind of object possessed by all group members |
| 20 | `sacrifice_small` | Small sacrifice with an explicit small-value conditional |
| 22.6.1 | `end_and_means_bad` | Bad end and bad means |
| 22.6.1 | `end_and_means_less_bad` | Qualitatively less bad, with an explicit adverse baseline |
| 22.6.2 | `end_and_means_reverse` | An end reclassified as a means in one evaluation context |
| 22.6.3 | `end_and_means_difficult` | Explicit not-easy evidence and the bad-end alternative |
| 26.2 | `negative_practical_true` | Opposite seriousness ranking and the true-presumption choice |

Additional repairs and controls accompany these records:

- The old snob record requires `group_accepted(G,A)` as ordinary evidence of
  the acceptance branch it concludes. A merely disjunctive ownership/property
  statement cannot justify accepting a proposition. Its old Rolex-as-a-
  proposition example is replaced by an actual proposition.
- The two named excuse forms use `person_group_member/2`. The general
  `group_member/2` also includes products, ideas and methods, which cannot
  alone supply a speaking person. An adverse case checks that collision.
- Both entry 22.6.2 forms name evaluation context C in the reclassification,
  warrant and conclusion. The warrant is an ordinary premise. Old graphs
  need a manually supplied context; the migrator invents none.
- Small sacrifice does not follow by invalidly taking the converse of the
  great-sacrifice conditional. Less-bad remains distinct from good and does
  not compute numerical utility. Not-easy requires explicit negative evidence.
- Entry 26.2's three CQs have an explicit static evidence, ultimate-issue and
  proof-standard contract. Two current-library worked graphs show a declared
  binary PE issue, with either one ranking or conflicting rankings.

## Changes made during critical verification

Two initially proposed negative label expectations failed because the whole
library's separate `pop_scheme` independently proves A from
`group_accepted(G,A)`. The new excuse/acceptance records themselves only
conclude action recommendations. The tests now retain full-library positive
integration and select the live target record for the distinct claim that
it does not itself prove A. They use the independently authored optional
`schemes` support from harness commit
`2f5a0e686eabfb19384e649a9d49fe262909875b`, which also checks all CLI
label, argument and issue expectations. No copied source rule or weakened
negative expectation replaces the selected-record test.

The first fixed global presumption issue exposed another real defect:
category and selected-graph generators copied issue patterns but omitted
language dependencies used only by those patterns. Actual engine validation
failed in two selected categories. The generators now retain these
dependencies, Python validation checks them, and a synthetic mutation test
proves a missing declaration is rejected.

Independent review then identified a semantic problem with the proposed
global presumption issue itself. A true choice and a negated false choice can
be compatible; automatically merging overlapping choice and negation issues
would treat them as exclusive alternatives in the engine's single-issue
model. The global issue was therefore withdrawn. The worked examples use
an explicitly scoped two-positive-choice protocol, and another case proves
the general library does not impose the unsafe global merge. This avoids
invalidating every source fingerprint for an unjustified global rule.

## Focused verification

The final focused runs use Go 1.24.13, the baseline vendored engines, and
the independently supplied complete-case-oracle harness. Full integrated
engine testing on both supported Go versions remains the parent review's
responsibility. The expected check manifest was not rewritten by this
implementation author.

| Focused verification | Result |
| --- | --- |
| All assigned source cases plus the historical book cases; API and CLI | 404/404 checks passed |
| Source ledger and migration mappings | 733/733 checks passed |
| Generic library checks in the three categories containing the new thirteen records | 1,791/1,791 checks passed |
| Four current generated usage examples, including both presumption protocols | 8/8 checks passed |
| Graph generator Python suite, including every current record and missing-issue-language mutation | 5 tests passed |
| Saved-graph migration Python suite, including no invented context/acceptance witness | 13 tests passed |
| Python scheme validator | 0 errors, 11 warnings |
| Whitespace/diff check | Passed |

The additional warning is the bad-end-to-bad-means propagation cycle.
It is reported openly rather than suppressed; the generic inference and
reinference checks pass. Warning count is not an acceptance proof.

Raw logs, JSONL result identities, the earlier failing controls and rendered
source evidence are preserved in the accompanying source-audit evidence.
The final logs are `branches-final-cases`, `branches-source-final`,
`branches-generic-final`, `branches-final-docs`, `branches-generator-final`,
`branches-migration-final` and `branches-validator-final`.

## Approval and integration boundary

This child snapshot has 139 records: the original 120 plus six part/whole
forms and thirteen other source alternatives. This is not a claim about the
eventual integrated record count. New variants and revised contracts remain
at `transcription-checked` with an explicit author hold.

The author ran the approval integrity check without changing the registry.
It reports precisely the expected stale approval for
`negative_consequences_reasoning`, whose likelihood semantics changed in the
first commit. That stale fingerprint must be renewed by an independent
reviewer after review. No global issue change remains, so unrelated approvals
are not made stale by a new global choice rule.

The next step is independent source approval of the authored interpretations
and final integrated verification. None of the source contracts claims
automatic advancement from a passing test, an implemented dialogue protocol,
unseen-source certification, or authority for this author to approve these
changes.
