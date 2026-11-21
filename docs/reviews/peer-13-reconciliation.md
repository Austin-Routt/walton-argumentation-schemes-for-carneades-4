# Reconciliation of PR #13 and the integrated independent review

Peer revision: `96daf6c0545f7cb84fe5e30cfbffdde5d6f6c0d1`, including PR #13's review and metadata commits. Integrated comparison baseline: `a708add511f52fb3603d92ba368875ef3d3f2c42`. This document records the comparison before the subsequent wording corrections prompted by the peer review. The active source ledger and current fingerprinted registry remain the authority for final dispositions.

## Merge decision

Preserve PR #13 in the Git ancestry and retain its original review, ledger, plan and release assessment as historical evidence. Adopt its explicit `RecordReview` fields (`by`, `date`, `decision`, `note`) and its stage-consistency checks. Retain the integrated ledger's source contracts, stages, question counts and question kinds, especially `source_questions`, positive `support`, and unresolved `alternative` mappings. An older review's stage promotions cannot replace later evidence-based holds.

PR #13's review is useful independent evidence. Its 7,994-check and historical R1 reruns concern the revisions it names; they do not certify the current integration or dispose of engine counterexamples found afterward. The review's release recommendation must therefore remain attributed to its original revision, rather than becoming the current release verdict merely through a merge.

Its reviewer/date fields document attribution; they cannot prove a reviewer was independent or performed a comparison. The semantic approval registry supplies a separate integrity property: the present representation must be the representation approved. Both are needed, and neither is source-author endorsement or a proof about all inputs.

## Dispositions reconciled

| Comparison at the stated baselines | Records |
|---|---:|
| PR #13 approved | 90 |
| Integrated review approved | 70 |
| Approved by both | 65 |
| Integrated approval after a peer hold | 5 |
| Peer approval retained as an integrated hold | 25 |
| Integrated records below interpretation approval | 50 |

The five integrated approvals absent from the peer's approved set are `abductive_action_to_character`, `abductive_character_to_action`, `popular_opinion`, `retroductive`, and `rhetorical_oppositions_descriptive`. Approval is conditional on the current source form and recorded scope; the peer's wording findings require correction and renewed acceptance of affected fingerprints before final integration.

The 25 peer-approved rows that remain held are:

`abductive_backward`, `alternative_normative`, `cause_to_effect`, `circumstantial_ad_hominem`, `end_and_means_1`, `end_and_means_2`, `end_and_means_3`, `from_antimodel`, `moral_justification_ad_populum`, `moral_justification_ad_populum_excuse`, `negative_normal_expectations`, `negative_value`, `perception`, `plea_for_excuse`, `popular_practice`, `positive_value`, `practical_inference`, `rule_mental_world`, `rule_physical_world_2`, `sacrifice`, `sign`, `snob_appeal_ad_populum`, `sunk_costs`, `value_based_practical_reasoning`, and `verbal_classification_arbitrariness`.

These holds have individual source, scope, burden or original-source reasons in `record-dispositions.json`; they are not an automatic rejection of the peer's work. For example, a passing isolated record does not establish correct separation of several agents, candidate explanations, observers, institutions or languages. The source-stage decision must match the actual promised scope.

The metadata reconciliation preserves 70 `example-passed`, 48 `transcription-checked`, and two `mapped` rows. Its decisions are 70 `approved`, 45 `hold`, and five `open`. The open cases are `position_to_know`, `expert_opinion`, `definition_to_verbal_classification`, `practical_reasoning2`, and `exeption_to_rule`: three pure own-original-source obligations and two unreviewed mappings. The remaining holds have known interpretation or implementation qualifications; several also retain unavailable-original-source obligations.

## Atkinson and Bench-Capon: exact explanation of the two counts

The source is the supplied 2007 paper, SHA-256 `dad4124c052c48b5c82de02494aac6971de113bb8452ae9fc6e764aa35e43cbc`. The relevant passages were reread for this reconciliation: the question inventory on pp. 858–859, the distinction between factual/model questions and alternative arguments on pp. 859 and 861–862, and the explicit positive treatment of CQ10 on p. 862. The full original comparison and source-based examples are in `atkinson-2007.md` and `atkinson-2007-decisions.json`.

| CQ group | PR #13 ledger | Integrated adjudication |
|---|---|---|
| 1–4 | Mapped premises | Direct AS1 premises; full AS2 model details are outside the flat abstraction |
| 5–7 | Mapped to stronger “more effectively” exceptions | Partial legacy policies; the source questions ask about alternatives without requiring greater effectiveness, and comparison remains incomplete |
| 8–9 | Both mapped to one side-effect exception | Partial legacy policy; same-value and other-value demotion are merged and defeat is unconditional, while the paper compares negative arguments and value preferences |
| 10 | Blank exception with a note that exception may not fit | Positive support for the same action through another value/goal instantiation, with dedicated argument-local cases |
| 11, 12, 14, 15, 17 | Blank mappings | Still lack separate operational mappings or a verified model-construction contract |
| 13, 16 | Defaults for action possibility and value legitimacy | Direct AS1 abstractions under explicitly inherited default/questioning policies |

Thus PR #13's “11 modelled” means **11 nonblank mapping fields**: CQ1–9, CQ13 and CQ16. Five of these are the partial legacy policies CQ5–9. Its six blanks are CQ10–12, CQ14–15 and CQ17. This count is reproducible, but it is not evidence of eleven complete source-faithful implementations.

The integrated classification is **seven direct AS1 treatments** (1–4, 10, 13, 16), **five partial legacy treatments** (5–9), and **five questions without separate mappings** (11, 12, 14, 15, 17). CQ5–7 remain blank in the completeness ledger and their stronger legacy exceptions are explicitly retained as additional `n: 0` policies. The result is **eight incomplete ledger mappings**: 5, 6, 7, 11, 12, 14, 15 and 17. CQ8–9 remain nonblank but qualified, so even eight blank mappings are not a complete measure of the remaining semantic differences.

CQ10 is the substantive difference in the executable assessment. The paper describes a positive side effect supporting performance of the action, not an automatic reason to defeat it (p. 862). The current engine can represent two supporting instances for the same action; the farmer Arg6-based tests establish that scoped capability. Treating CQ10 as an unmapped exception would discard working positive support and encode the wrong direction if an exception were later added mechanically.

Both reviews correctly reject equating goal worthiness with goal realizability and retain the separate CQ17 concerning other agents' participation. The integrated `practical_reasoning1` hold is therefore appropriate: source access and transcription comparison are complete, but full-source interpretation is not. No additional broad review round is needed to settle this count. The bounded remaining task is the explicit AATS adapter/model contract and alternative/value-comparison work already recorded.

## Metadata integrity and application

Keep `by`, `date`, and `decision` outside the semantic fingerprint. They describe the review act, not the scheme's representation. Each current registry approval instead snapshots those three ledger fields and must match them exactly. A changed attribution, date or decision fails the gate even when the semantic hash is unchanged. A held decision cannot authorize a registry approval merely because both copies say `hold`. A withdrawn stage also cannot retain an active approving review.

`tools/apply_source_review_metadata.py` derives metadata from the existing, accepted disposition register and current approval registry. It does not create an approval, promote a stage, or recalculate semantic hashes. It refuses a changed approved scheme until that change has been independently adjudicated and the accepted fingerprint updated. Its default invocation is read-only; `--write` updates only ledger review fields and approval review snapshots. The attribution names the integrated ChatGPT source-review team, never the peer reviewer whose historical decisions are archived separately.

The helper was applied to an isolated copy of the 120-row baseline, then the modified data were restored. It produced 70 approvals, 45 holds and five opens, left all 120 semantic fingerprints unchanged, preserved every non-review ledger field, passed the approval audit, and was idempotent. Nineteen synthetic approval tests passed, including the new review-attribution and metadata-reconciliation cases. This metadata patch adds no registered engine/library checks and makes no new full-suite or two-toolchain execution claim.
