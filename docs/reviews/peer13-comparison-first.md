# Independent comparison with PR13: book entries 1–27

Peer commit: `96daf6c0545f7cb84fe5e30cfbffdde5d6f6c0d1`. Its substantive review inspected the earlier `49ad425` implementation. Our integration baseline is `a708add511f52fb3603d92ba368875ef3d3f2c42`; the verified reconciliation patch is `71a0d993fab795beeb18897dfc5017205ff08bd6`.

## Decision

The peer review is valuable and adds corrections our earlier review missed: **three wording errors, two safe shared-statement opportunities, and the remaining content error in the original negative-practical reference case**. These are accepted and repaired in the isolated patch. It also independently confirms several earlier findings already repaired in our branch. Its broader approval count does not resolve the distinct remaining questions about burden policy, source variants, context and original-source fidelity. Preserve the current **70 approvals and 50 holds** while integrating the concrete corrections and peer history.

For these 63 assigned records the peer approves 47 and holds 16; our final disposition approves 32 and holds 31. Thirty are approved by both, 14 are held by both, 17 are approved only by the peer, and 2 were held by the peer for wording now corrected in our patch. These counts compare judgments about different representations; they are not votes about one frozen artifact.

## Accepted new findings and corrections

| Record or pair | Accepted evidence | Action |
|---|---|---|
| `popular_opinion` | Printed p311 CQ2 doubts whether A is true, even if generally accepted. | Correct good-reasons/truth wording; preserve its scoped exception role and refresh approval fingerprint. |
| `rhetorical_oppositions_descriptive` | Printed p318 explicitly reverses the two participants and includes contrariwise. | Correct the warrant language; the existing conclusion already had correct argument order. |
| `two_person_practical` | Printed pp97,325 say sees, not sets. | Correct wording and preserve adviser-to-agent communication; retain its separate semantic hold. |
| `from_example` / `illustration` | Printed pp314–315 use the same individual/F/G conjunction. | Migrate `same_properties/3` to `has_properties/3`; add actual two-scheme chaining and wrong-individual/property controls. |
| `interaction_act_person_1/2` | Printed p321 uses the same person/judgment-value statement as conclusion and premise. | Migrate `person_judgment_value/2` to `person_value/2`; add two-scheme chaining and wrong-person/value controls. This pair is entry 18, not positive/negative value in entry 19. |
| `negative_practical/book` | The already repaired negative-presumption rule was still paired with an incoherent loaded/accident/delay illustration. | Change the illustrative proposition to unloaded and identify it as an application of the printed schema. Assuming unloaded falsely risks an accident; assuming not-unloaded falsely causes delay. |

The supplied PDF pages were visually checked; printed page plus 10 gives the PDF page. The first-half addendum records the corrected interpretations, migrations and approval-fingerprint scope. Five fingerprints change; four already approved entries are refreshed, and the held two-person record acquires no approval.

## Verification consequence of legitimate chaining

The new shared conjunction revealed an invalid old test assumption. A generalization withdrawn from stipulated facts may still be derived by `from_example`, so its absence from the input does not imply absence from the full theory. Questioning that induction default affects acceptance but still permits structural argument construction, so it cannot isolate an argument-construction check.

The generic premise-necessity and reference-case withdrawal controls therefore isolate the actual selected schema, with an isolated successful positive control. The ordinary full-library run, expected arguments, grounded labels and exception checks remain in their original theory context. A synthetic regression verifies that full-theory rederivation is valid while a genuinely unnecessary premise and an absent positive control are still rejected. The adverse illustration case targets the intended flies argument rather than forbidding a legitimate independently derived sings argument. The new explicit two-scheme cases preserve the interoperability assertion. All intermediate red logs are retained. This changes a documented oracle contract; it does not waive a failed rule or blindly replace expected labels.

## Earlier findings independently corroborated

The peer confirms the example-dependent generalization defect, practical agent leakage, each-one sufficient wording, mistaken part/whole defeater, negative-practical polarity, precision/certainty typo, and whole-to-parts coverage limits. The first six were already repaired or explicitly held in our branch; the negative-practical fixture needed the additional correction described above. The peer also agrees that `verbal_classification` can be certified against the available compendium/code origin because upstream cited no separate work.

The peer approved the old `from_model` while acknowledging that its third premise retained only the conditional consequent. That acknowledgment does not rebut our source-backed restoration of the full model-specific warrant. Retain the repair and its wrong-model adverse case. Likewise, our compendium-supported definition CQ2 mapping can remain implemented while the separately cited original is still held; an unavailable original does not prevent a clearly attributed available-book adaptation.

## Why the wider peer approvals do not replace the final holds

The 17 peer-approved/final-held records are listed individually below. Some disagreements are about declared coverage rather than a claim that the executable branch is intrinsically wrong: great-sacrifice versus all alternatives, positive end/means branches, and a binary recommendation reading. Others involve substantive unclosed scope or burden questions: mixed agents or contexts, incomplete type separation, dialogue challenge versus evidential exception, and degree/strength policies. Existing positive examples do not discriminate all of those alternatives. The peer adds no new primary-source comparison or adversarial evidence closing them. A later explicit bounded contract and corresponding review can close a specific hold; importing an approved label alone cannot.

Both reviewers accessed the sunk-cost article and agree its printed p489 scheme differs from the legacy quantity/feasibility form. The remaining disagreement is whether to certify the legacy composite adaptation under its selected policies. Keep our qualified adaptation decision and describe the access accurately; do not call the article unavailable.

## Other cross-cutting peer observations

- The full-library all-parts conclusion is expected while the divisional-heredity default remains undisputed. It must not be described as evidence that every arbitrary property transfers to parts. Usage guidance should expose scheme selection, whole/context identity and questioning. This agrees with the bounded approval already recorded for composition/division.
- The peer identifies additional moral-bad-person and reasonable-belief splits outside this first-half pair assignment. They require the separate second-half source comparison; they are not silently treated as fixed here.
- Restoring upstream's general `biased/1` wording and correcting historical connectivity counts are legitimate provenance/documentation concerns for the integration owner. Neither provides a reason to override first-half semantic holds.
- The peer review field is useful historical evidence. It refers to an earlier implementation and must coexist with the final representation-bound approval registry and corrected dispositions.
- This source comparison makes no new certification of historical R1 engine behavior. The independent runtime findings and final clean merged-commit gates remain separate evidence.

## Complete 63-record comparison

### `position_to_know` — entry 1

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** No material first-half disagreement beyond the recorded scope/source qualifications.

**Peer note:** identical to upstream's file and compared with its entry in the compendium; the book it cites is not available: 5.5

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Book pages:** 309.

### `expert_opinion` — entry 2

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** No material first-half disagreement beyond the recorded scope/source qualifications.

**Peer note:** compared with entry 2 and with pp. 374-375, which have the assumption credible(W) added in 5.5; identical to upstream's file otherwise. Its own source, Appeal to Expert Opinion (1997), is not available: 5.5

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Book pages:** 310, 374, 375.

### `witness_testimony` — entry 3

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** No material first-half disagreement beyond the recorded scope/source qualifications.

**Peer note:** compared with entry 3; identical to upstream's file but for the order of in_domain's arguments, which its template requires. Its own source (ICAIL 2003) was not found in an open copy: 5.5

**Next concrete action:** Adopt and test case-specific witness/account identifiers for bias and sincerity; compare both separately cited witness sources before original-source certification.

**Book pages:** 310.

### `popular_opinion` — entry 4

**Peer:** transcription-checked (hold). **Final:** example-passed. **Comparison:** peer holds; final approves after correction.

**Adjudication:** New missed wording defect accepted and repaired: CQ2 doubts truth of A, not its general truth; use good reasons, printed p311. Retain approval only with the corrected template and reviewed fingerprint.

**Peer note:** wording: the exception's template says 'doubting %s is generally true'; the print's second question is about doubting that it is true

**Book pages:** 311.

### `pop_scheme` — entry 4.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 311.

### `pop_scheme_acceptance` — entry 4.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 311.

### `position_to_know_ad_populum` — entry 4.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 311.

### `expert_opinion_ad_populum` — entry 4.3

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 312.

### `deliberation_ad_populum` — entry 4.4

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 312.

### `moral_justification_ad_populum` — entry 4.5

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** Peer approves while noting the omitted (or should be) alternative. Preserve the contextual/participant qualification and disclose that omitted source alternative; the approval vote does not implement it.

**Peer note:** the template of premise 2 leaves out the print's '(or should be)'

**Next concrete action:** State and test a one-respondent contract, or carry respondent identity in the acceptance and conclusion predicates before mixed-audience use.

**Book pages:** 312.

### `moral_justification_ad_populum_excuse` — entry 4.6

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Choose the bounded combined action/policy reading explicitly; if both are offered independently, add typed variants and separate examples.

**Next concrete action:** Choose the bounded combined action/policy reading explicitly; if both are offered independently, add typed variants and separate examples.

**Book pages:** 312.

### `snob_appeal_ad_populum` — entry 4.7

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Replace the object-as-proposition Rolex acceptance fixture with a typed acceptance proposition; add separately typed property/purchase variants if those source alternatives are claimed.

**Next concrete action:** Replace the object-as-proposition Rolex acceptance fixture with a typed acceptance proposition; add separately typed property/purchase variants if those source alternatives are claimed.

**Book pages:** 312.

### `appeal_to_vanity_ad_populum` — entry 4.8

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 312.

### `rhetoric_ad_populum` — entry 4.9

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 313.

### `common_folk_ad_populum` — entry 4.10

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 313.

### `common_folks_group_ad_populum` — entry 4.11

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 313.

### `popular_practice` — entry 5

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Clarify the anaphoric warrant and require context-qualified practices; add a two-context adverse case demonstrating that one practice context cannot establish another.

**Next concrete action:** Clarify the anaphoric warrant and require context-qualified practices; add a two-context adverse case demonstrating that one practice context cannot establish another.

**Book pages:** 314.

### `from_example` — entry 6.1

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** The example-dependent conclusion was already repaired to usually_properties(F,G), with separate example provenance. Retain the unresolved strength-policy qualification; the new shared conjunction additionally improves legitimate composition.

**Peer note:** the conclusion both_properties(A,P1,P2) carries the example A, which the print's generalization does not: two examples do not support one generalization, and the conclusion cannot be 6.2's first premise. 6.2 was corrected for the same thing (5.1)

**Next concrete action:** Specify how the book's exception question affects the strength of a usually-generalization; add a policy-level example where one counterexample does not automatically refute a statistical generalization.

**Book pages:** 314.

### `illustration` — entry 6.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer identified a missed safe shared statement: same_properties/3 and has_properties/3 have the same individual/property/property roles. Canonicalize to has_properties and add a two-scheme positive/adverse case. Its existing approved interpretation is retained with a new fingerprint.

**Peer note:** same_properties/3 says what 6.1's has_properties/3 says, under another name

**Book pages:** 315.

### `from_model` — entry 6.3

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer approved the old record despite explicitly noting its third premise retained only the conditional consequent. Keep our independently source-supported conditional-warrant repair and wrong-model control; the peer approval is not evidence to revert it.

**Book pages:** 315.

### `from_antimodel` — entry 6.4

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Adopt and test the binary recommendation convention, or introduce an explicit refrain-from-action conclusion that distinguishes ought-not from not-ought.

**Peer note:** approved as an interpretation: 'Qs should not do A' is the negation of should_do, so that 6.3 and 6.4 meet in one issue (5.6); the claim is changed from transcription

**Next concrete action:** Adopt and test the binary recommendation convention, or introduce an explicit refrain-from-action conclusion that distinguishes ought-not from not-ought.

**Book pages:** 315.

### `analogy` — entry 7

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** No material first-half disagreement beyond the recorded scope/source qualifications.

**Peer note:** identical to upstream's file and compared with its entry in the compendium; the book it cites is not available: 5.5

**Next concrete action:** Compare the cited 2006 source; decide whether the inherited more-on-point defeater is retained as a variant or supplemented by the book's similar contrary-case CQ3, with an adverse case.

**Book pages:** 315.

### `practical_analogy_positive` — entry 8.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 316.

### `practical_analogy_negative` — entry 8.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 316.

### `composition_generic` — entry 9.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer notes broad default composition/division behavior. Preserve approval within the stated same-object/property context and challenge-sensitive hereditary policy; do not imply every arbitrary property really transfers.

**Book pages:** 316, 113.

### `part_in_whole` — entry 9.2

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** The false lower-degree defeater was already removed. Keep the explicit positive/species/degree-computation limitation.

**Peer note:** the exception smaller_part_property is the parenthesis of the print's conclusion ('y is less A than X ...'), a variant of the conclusion and not a defeater of 'y is A'; the print gives 9.2 no question

**Next concrete action:** Declare the repaired positive species branch as the supported variant; add separate physical-part and degree variants only if those alternatives are included in release coverage.

**Book pages:** 316.

### `division_generic` — entry 10.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** The peer's full-library all-parts observation is correct when divisional heredity remains an undisputed default. It is expected defeasible behavior, not universal validity; retain explicit whole/context and challenge conventions in usage guidance.

**Book pages:** 317, 113.

### `whole_into_parts` — entry 10.2

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** Both reviews hold the same substantive limitation: first selected part, three-part restriction and omitted negative/other-cardinality readings.

**Peer note:** only the part listed first can carry the property (has_property(X1,A)), where the print's premise 3 is about any part; the negative reading (no part is A, so X is not A) is not stated

**Next concrete action:** Declare the positive, three-part instance explicitly and test that contract; add or exclude arbitrary cardinality, selected-part, negative and only-if readings separately.

**Book pages:** 317.

### `opposition_descriptive_1` — entry 11.1a

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 317.

### `opposition_descriptive_2` — entry 11.1b

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 318.

### `opposition_descriptive_3` — entry 11.1c

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 318.

### `opposition_normative` — entry 11.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 318.

### `rhetorical_oppositions_normative` — entry 12.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 318.

### `rhetorical_oppositions_descriptive` — entry 12.2

**Peer:** transcription-checked (hold). **Final:** example-passed. **Comparison:** peer holds; final approves after correction.

**Adjudication:** New missed wording defect accepted and repaired: converse relation reverses participants and the warrant includes contrariwise, printed p318.

**Peer note:** wording: the major premise's template says 'they can be in the converse relation'; the print says that Y can be in the converse relation with X, and contrariwise

**Book pages:** 318.

### `alternative_cognitive` — entry 13.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Peer note:** the alternatives are ordered as in the print: the one denied (13.1) or required (13.2) has to be written first

**Book pages:** 318.

### `alternative_normative` — entry 13.2

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Specify and test an exclusive-alternatives recommendation policy; preserve the distinction between absence of obligation and an obligation to refrain.

**Peer note:** the alternatives are ordered as in the print: the one denied (13.1) or required (13.2) has to be written first

**Next concrete action:** Specify and test an exclusive-alternatives recommendation policy; preserve the distinction between absence of obligation and an obligation to refrain.

**Book pages:** 318.

### `verbal_classification` — entry 14

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer agrees with our corrected provenance decision: upstream cited no separate work. Declared compendium/code-origin approval is valid without pretending the unseen 2006 book was inspected.

**Peer note:** upstream's record, which cites no source, compared with its entry in the compendium; identical to upstream's file apart from the changes the row lists

**Book pages:** 319.

### `definition_to_verbal_classification` — entry 15

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** Peer still leaves the book CQ2 unmapped pending the unavailable original. Our book-supported doubtful_definition(D,G) mapping legitimately closes that operational compendium gap while preserving the separate own-original-source hold.

**Peer note:** identical to upstream's file and compared with its entry in the compendium; the book it cites is not available, and the printed questions the row lists stay unmodelled: 5.5, 5.2

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Book pages:** 319.

### `verbal_classification_vagueness` — entry 16

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** Precision rather than certainty was already corrected. Preserve the unresolved challenge-burden policy.

**Peer note:** wording: the exception's template says the property 'does not need to be certain'; the question is about precision

**Next concrete action:** Adopt an explicit challenge burden: determine whether an unanswered precision challenge suspends the argument or requires affirmative contrary evidence, and test the chosen behavior.

**Book pages:** 319, 320.

### `verbal_classification_arbitrariness` — entry 17

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Adopt and test the burden for an unanswered arbitrariness challenge, without adding a premise that merely restates the desired conclusion.

**Next concrete action:** Adopt and test the burden for an unanswered arbitrariness challenge, without adding a premise that merely restates the desired conclusion.

**Book pages:** 320.

### `interaction_act_person_1` — entry 18.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer identified the same person/valuation relationship in entries18.1 and18.2. Canonicalize the second record to this record's person_value/2 and add a two-scheme case; this producer formula itself is unchanged.

**Peer note:** 18.1's conclusion person_value/2 and 18.2's first premise person_judgment_value/2 both say 'person P is V'; as two predicates they keep the pair from chaining (to decide before R2's migration)

**Book pages:** 320.

### `interaction_act_person_2` — entry 18.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Peer identified the same person/valuation relationship in entries18.1 and18.2. Rename person_judgment_value/2 to person_value/2 with migration and a two-scheme positive/adverse case; retain approval with a new fingerprint.

**Peer note:** 18.1's conclusion person_value/2 and 18.2's first premise person_judgment_value/2 both say 'person P is V'; as two predicates they keep the pair from chaining (to decide before R2's migration)

**Book pages:** 320, 321.

### `positive_value` — entry 19.1

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case.

**Next concrete action:** Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case.

**Book pages:** 321.

### `negative_value` — entry 19.2

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case independently of the repaired polarity.

**Next concrete action:** Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case independently of the repaired polarity.

**Book pages:** 321.

### `sacrifice` — entry 20

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Release only the documented great-sacrifice variant, or add the source's small-sacrifice alternative with separate semantics and examples.

**Peer note:** only the 'great' variant of the print is stated

**Next concrete action:** Release only the documented great-sacrifice variant, or add the source's small-sacrifice alternative with separate semantics and examples.

**Book pages:** 322.

### `group_1` — entry 21.1

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 321, 322.

### `group_2` — entry 21.2

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 322.

### `practical_inference` — entry 22.1

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Choose and document static evidential exceptions versus the source's question-induced temporary suspension; test the selected contract and avoid claiming full dialogue behavior if only static evaluation is supported.

**Next concrete action:** Choose and document static evidential exceptions versus the source's question-induced temporary suspension; test the selected contract and avoid claiming full dialogue behavior if only static evaluation is supported.

**Book pages:** 95, 96, 323.

### `practical_necessary_condition` — entry 22.2

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** All personal premises and relevant conditions already carry the same agent, with two-agent adverse controls. The separate question-induced suspension policy remains qualified.

**Peer note:** the alternatives, selection, practicality and side-effects premises are in the first person and carry no agent, while the goal premise and the conclusion name P: with two agents in one graph one agent's selection serves the other (probe in the review)

**Next concrete action:** Retain the repaired agent scope and passing two-agent controls; resolve and test the source's question-induced suspension policy before unqualified interpretation approval.

**Book pages:** 95, 96, 323, 324.

### `practical_sufficient_condition` — entry 22.3

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** Agent leakage and each-one sufficient wording were already repaired against narrative p96, with adverse controls. The separate question-induced suspension policy remains qualified.

**Peer note:** as 22.2; and the alternatives premise should read as p. 96 does, 'each one of [B0..Bn] is sufficient', not 'at least one of'

**Next concrete action:** Retain the narrative-supported each-one sufficient reading and repaired agent scope; resolve and test provisional challenge behavior before unqualified interpretation approval.

**Book pages:** 96, 323, 324.

### `value_based_practical_reasoning` — entry 22.4

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Require owner/context-specific goals and value sets; settle the challenge burden and add a mixed-agent/context adverse example.

**Next concrete action:** Require owner/context-specific goals and value sets; settle the challenge burden and add a mixed-agent/context adverse example.

**Book pages:** 95, 96, 324.

### `practical_goal` — entry 22.5

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 325.

### `end_and_means_1` — entry 22.6.1

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Declare the positive branch as the released variant; add or explicitly exclude the source's parenthetical alternatives and test any added branch independently.

**Next concrete action:** Declare the positive branch as the released variant; add or explicitly exclude the source's parenthetical alternatives and test any added branch independently.

**Book pages:** 325.

### `end_and_means_2` — entry 22.6.2

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Declare one evaluation context and the positive branch; add scoped warrants for mixed contexts and separate opposite-valuation examples if those are claimed.

**Next concrete action:** Declare one evaluation context and the positive branch; add scoped warrants for mixed contexts and separate opposite-valuation examples if those are claimed.

**Book pages:** 325.

### `end_and_means_3` — entry 22.6.3

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Declare the easy/good branch as the released variant; add or explicitly exclude the source's alternative branch.

**Peer note:** only the positive variant of the print is stated

**Next concrete action:** Declare the easy/good branch as the released variant; add or explicitly exclude the source's alternative branch.

**Book pages:** 325.

### `two_person_practical` — entry 23

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** New missed wording defect accepted and repaired: sees rather than sets, printed pp97,325. Separate dialogue and alternatives-scope holds remain.

**Peer note:** wording: the template of premise 2 says 'sets the situation'; the print has 'sees' (pp. 97, 325)

**Next concrete action:** Keep the repaired speaker/recipient roles; define practical-dialogue burdens and scope alternative actions to the deliberating parties, with an adverse multi-party example.

**Book pages:** 325, 326.

### `waste` — entry 24

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 326.

### `sunk_costs` — entry 25

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** Both reviews now have original-article access and agree the p489 form differs from upstream. Peer approves the legacy adaptation; our source-certification scope remains qualified for the composite quantity/feasibility/future-loss policy. This is a scope disagreement, not an unavailable-article claim.

**Peer note:** upstream's own form, kept (D5); Walton (2002) read by the reviewer: p. 489 prints the form of sunk_costs_precommitment, p. 488 the argument from waste, and no form in upstream's terms

**Next concrete action:** Retain the legacy quantity/feasibility record as a documented adaptation, distinct from sunk_costs_precommitment; settle the selected practical-question and future-loss policy rather than treating the legacy form as the article's printed p. 489 scheme.

**Book pages:** 327.

### `sunk_costs_precommitment` — entry 25

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 327.

### `ignorance` — entry 26

**Peer:** transcription-checked (open). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** No material first-half disagreement beyond the recorded scope/source qualifications.

**Peer note:** identical to upstream's file and compared with its entry in the compendium; the book it cites is not available: 5.5

**Next concrete action:** Compare the cited 1995 source and specify application-supplied dialogue burden/probandum independently of the issue proof standard; add examples for the chosen burden contract.

**Book pages:** 327.

### `negative_normal_expectations` — entry 26.1

**Peer:** example-passed (approved). **Final:** transcription-checked. **Comparison:** peer approves; final holds.

**Adjudication:** No new source evidence overturns the existing hold. Use one expectation context or add a situation identifier; specify dialogue burden allocation separately from required proof strength and add a mixed-context adverse example.

**Next concrete action:** Use one expectation context or add a situation identifier; specify dialogue burden allocation separately from required proof strength and add a mixed-context adverse example.

**Book pages:** 327.

### `negative_practical` — entry 26.2

**Peer:** transcription-checked (hold). **Final:** transcription-checked. **Comparison:** both hold.

**Adjudication:** Peer confirms the already repaired negative-presumption formula. Its separate example criticism is newly accepted: the old loaded/accident/delay fixture was still reversed; use unloaded and explicitly identify this as an application, not a printed book example. The existing branch/burden hold remains.

**Peer note:** wording: the conclusion's template says 'not true or true' and so concludes nothing; with premise 5 as stated (B more serious than C) the print concludes 'not true'. The instance in cases/book.yml has the consequences of its example the wrong way round

**Next concrete action:** Retain the source-correct negative-presumption branch; explicitly exclude or add the true-presumption branch and supply a separate dialogue-burden/probandum contract.

**Book pages:** 327.

### `epistemic_ignorance` — entry 27

**Peer:** example-passed (approved). **Final:** example-passed. **Comparison:** both approve.

**Adjudication:** Agreed within the declared source interpretation and scope.

**Book pages:** 328.


## Completed patch verification

Patch commit: `71a0d993fab795beeb18897dfc5017205ff08bd6`. Go 1.24.13 passed the selected
`TestLibrary`, `TestCases` and `TestSource` run: **7,861 checks and 7,861
unique result keys, all passing**, with zero known/unexpected failures and
zero stale ledger entries. The generic isolation regression passes; validator
output is 0 errors/11 warnings; all 5 validator tests and 8 migration tests pass.
Category parity and source-approval integrity pass. The full expected inventory
was not replaced from this partial run. Final complete merged-commit acceptance
remains the integration owner's responsibility.
