# Final consolidated independent source dispositions

Review date: 2026-11-21.

The current integrated ledger has **70 records at `example-passed`, 48 at `transcription-checked`, and 2 at `mapped`**. The 70 approved records comprise **65 book-origin and 5 inherited records**. The other **50 records remain held**. Every actual promoted ID matches an explicit independent adjudication; no unsupported promotion or unreviewed drift in the 114 book-mapped formulas was found.

This is a source and interpretation disposition, not a replacement for the complete commit-matched engine/library release gate. Approval covers the declared executable interpretation and documented scope. It does not claim that every historical source alternative, numeric probability, dialogue state, or cited predecessor has been implemented or independently read.

## Final correction and provenance

The initial reconciliation matched 69 intended approvals. Final scrutiny found one false blocker in the first-half review: `verbal_classification` did not cite a separate original work. Direct comparison with the upstream code found identical variables, premises and conclusion; the added `doubtful_classification(F,G)` implements the supplied compendium's second question and is already documented in migration. Its declared compendium/code-origin interpretation is therefore approved, bringing the total to 70. The unseen 2006 book is not claimed as verified. The dated provenance addendum preserves and corrects the earlier review history.

The same distinction changes the description of the `cause_to_effect` hold: the outstanding issue is causal-context scope, not an unidentified unavailable original. The original sunk-cost article has now been accessed by root as author-provided web text with printed-page markers; the legacy form remains a qualified adaptation because it differs from the printed precommitment formulation. The Atkinson (2007) source has been fully supplied and reviewed; its hold concerns incomplete interpretation/model coverage, not access.

## Counts and source-question boundaries

| Disposition | Records |
|---|---:|
| book form approved; own-original-source hold | 3 |
| qualified; interpretation approval withheld | 44 |
| approved within declared scope | 70 |
| blocked as exact mapped book formulation | 1 |
| mapping-only; source interpretation unreviewed | 2 |

There are **206 printed-question occurrences counted per record**, across **118 records with known declared-source counts**. **53 known counts are zero**. Two counts remain **unknown**, for `practical_reasoning2` and `exeption_to_rule`; unknown is not zero. This per-record count repeats a source question when two variants share it and is not a count of unique book questions.

There are **15 additional `n: 0` hooks**, openly separated from the printed questions. In particular, `abduction` has `source_questions: 0` and a recorded `more_coherent_explanation(T,S)` exception at `n: 0`. `credible_source` has zero printed questions and three inherited policy exceptions. A zero printed count does not mean the implementation contains no assumptions or defeaters.

The **eight nonstandard blank operational mappings** are Atkinson CQ5,6,7,11,12,14,15,17. The **eleven additional `standard` rows without atoms** describe application-level strength/proof/burden decisions and are not counted as the same kind of implementation blank. Their acceptability remains record-specific: an issue's proof standard does not by itself allocate a dialogue burden. A nonblank mapping can also be semantically partial, as Atkinson CQ8-9 demonstrate. The still-unread originals may reveal additional questions or distinctions.

For a book-mapped inherited record, its declared count is the compendium count. The original-source status and count are separately reported below and in the CSV. The supplied book's bibliography does not prove that an earlier work was inspected.

## Snapshot and audit coverage

Repository HEAD at capture: `15a5e155647f0df098c62aa507ad1a3822fc03dc`. Integration edits may be uncommitted; the actual reviewed content is fixed by these digests:

- Source ledger SHA256: `6cf5bfcacca1a0629ba88a289a8de16ba40ec9eff0a4b392d21a69757e71e545`.
- `walton.yml` SHA256: `197c68f249b26012e846fe5a84a53eb6e782511c4d2b8ac87fab8029653d1d47`.
- Supplied book SHA256: `d010348014e8b80fa561b68e5fac732cce0b3905fb0b1d522dd6fd2578a69634`.

All 120 IDs are unique and accounted for. All known CQ inventories have contiguous complete numbering; this mechanical condition is kept separate from semantic completeness. Variables, premises, assumptions, exceptions and conclusions of all 114 book-mapped records match the independently adjudicated repaired forms. The 31.4 counterevidence wording now explicitly requires evidence sufficient to undermine the qualitative prediction; an isolated atypical act is not automatically a refutation of a likely tendency.

The first-half review inspected 63 records, the second-half review 51, the original Atkinson review one, and the remaining upstream-only dispositions are supported by their specific code/original-source comparisons or retained mapping-only status. Root independently reviewed the original appearance and credible-source articles and the available sunk-cost article text. Those access methods are attributed rather than represented as this consolidator's independent PDF inspection.

## PR13 follow-up and preserved review history

Remote PR13 supplied useful independent findings about an earlier implementation. The reconciliation accepted three wording corrections (`popular_opinion`, `rhetorical_oppositions_descriptive`, `two_person_practical`), two first-half shared-statement opportunities, a remaining error in the original negative-practical application fixture, and the perception/memory shared reasonable-belief conclusion. It also restored the inherited generic biased wording. These concrete corrections do not justify replacing the final record-specific holds with the peer's broader approval count.

The only formula differences between the pre-PR13 adjudications and the current library are the independently approved canonicalizations in `illustration` (`same_properties` to `has_properties`), `interaction_act_person_2` (`person_judgment_value` to `person_value`), and `perception` (`reasonable_case` to `reasonable_believe`). Original adjudication JSONs are preserved as `adjudication-pre-peer13.json` and `per-record-audit-pre-peer13.json`; the active reviewed-form fields were updated only for those three identities, with dated source/evidence notes. The record-by-record sections below append the other wording, fixture and policy clarifications and the new example IDs. All 114 book-mapped formulas match those explicitly reviewed representations, with no additional drift.

The new conjunction sharing permits legitimate rederivation of a withdrawn stipulation. Premise-necessity controls now isolate the actual selected schema after its own successful positive control; the ordinary full-library graph, argument/label assertions, exception checks and explicit two-scheme cases remain in their original contexts. This documented oracle correction preserves the distinction between structural argument construction and acceptance. Intermediate failing logs and the synthetic control for genuinely unnecessary premises are retained in the review evidence.

The complete peer comparisons are `peer13-comparison-first.md/json` (all 63 first-half records), `peer13-comparison-second.md/json` (the second-half scope), and the integrated `docs/reviews/pr13-reconciliation.md`. Exact aliases, preserved-input hashes, per-record follow-up notes and evidence references are embedded in the companion JSON's `peer13_followup` metadata and in `peer13-followup-2026-11-21.json`. The final source populations remain 70 example-passed, 48 transcription-checked and two mapped. A fresh clean acceptance run must use the final merged commit rather than the earlier candidate hashes.

## Complete record-by-record dispositions

Every record below includes its actual current stage, approval/hold reason, access obligation, concrete next action, source locator, CQ provenance and example evidence. The companion CSV has the same complete 120-row information; the JSON adds exact machine-readable audit metadata.

### 1. `position_to_know`

**Entry:** 1. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** book form approved; own-original-source hold. **Book-form decision:** approve.

**Approval or hold reason:** Domain D links W's access and proposition S; source and assertion roles are coherent. Bare S preserves truth-claim identity. CQ1 and CQ3 are premises, honesty is an opponent-supported exception. This is a reasonable static reading of the three book questions. Book-form approval does not verify its separately cited original; preserve the own-original-source gate.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Legal Argumentation and Evidence, The Pennsylvania State University Press, University Park, 2002, p.46.

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 1, printed pp. 309 (PDF pages +10). Separately cited original: Douglas Walton, Legal Argumentation and Evidence, The Pennsylvania State University Press, University Park, 2002, p.46.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: position_to_know(W,D); CQ2 exception: dishonest(W); CQ3 premise: asserts(W,S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** position_to_know/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 2. `expert_opinion`

**Entry:** 2. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** book form approved; own-original-source hold. **Book-form decision:** approve.

**Approval or hold reason:** Expert W, proposition S and domain D stay linked. Bare S is the correct object proposition, without a second plausibility fact. All six printed questions are represented; credibility as a default is additionally corroborated by the book's own Carneades model, pp.374-375. Book-form approval does not verify its separately cited original; preserve the own-original-source gate.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Appeal to Expert Opinion, The Pennsylvania University Press, University Park, Albany, 1997, p.211-225.

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 2, printed pp. 310, 374, 375 (PDF pages +10). Separately cited original: Douglas Walton, Appeal to Expert Opinion, The Pennsylvania University Press, University Park, Albany, 1997, p.211-225.

**Declared-source CQ count:** 6. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 assumption: credible(W); CQ2 premise: in_domain(S,D); CQ3 premise: asserts(W,S); CQ4 exception: untrustworthy(W); CQ5 exception: inconsistent_with_other_experts(S); CQ6 assumption: based_on_evidence(asserts(W,S))

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** expert_opinion/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 3. `witness_testimony`

**Entry:** 3. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** W's sincerity is represented by believes(W,S) plus assertion; this is a defensible interpretation, not an exact quotation of knowing truth. A domain fact supplies access. Internal consistency defaults; conflicting facts, other witnesses, bias and implausibility are exceptions. Global biased(W) needs case-specific use, because the question concerns the account. Remaining scope: Make explicit that W denotes this witness in the case and that bias is relevant to this account; own-source verification is separate. Follow-up: PR13 follow-up restores the generic upstream biased/1 wording instead of the added Arguer restriction. This does not resolve the case/account scope or separately cited original-source holds.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Henry Prakken, Chris Reed, Argumentation Schemes and Generalisations in Reasoning about Evidence, Proceedings of the 9th International Conference on Artificial Intelligence and Law, Edinburgh, 2003. New York: ACM Press 2003, pp. 35. Douglas Walton, Witness Testimony Evidence, Cambridge University Press, 2008.

**Next concrete action:** Adopt and test case-specific witness/account identifiers for bias and sincerity; compare both separately cited witness sources before original-source certification.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 3, printed pp. 310 (PDF pages +10). Separately cited original: Douglas Walton, Henry Prakken, Chris Reed, Argumentation Schemes and Generalisations in Reasoning about Evidence, Proceedings of the 9th International Conference on Artificial Intelligence and Law, Edinburgh, 2003. New York: ACM Press 2003, pp. 35. Douglas Walton, Witness Testimony Evidence, Cambridge University Press, 2008.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 assumption: internally_consistent(S); CQ2 exception: inconsistent_with_known_facts(S); CQ3 exception: inconsistent_with_other_witnesses(S); CQ4 exception: biased(W); CQ5 exception: implausible(S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** witness_testimony/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-second-followup.md.

### 4. `popular_opinion`

**Entry:** 4. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The output remains a reason in favor of A, not an assertion that A is true. General acceptance and the acceptance-to-reason warrant are distinct. CQ1 is challenged via the general-acceptance premise; CQ2 is an exception. Both printed questions are accounted for. Follow-up: PR13 follow-up corrected an error missed by the first independent review: printed p311 CQ2 asks for good reasons to doubt that A is true, not that A is generally true. The template now expresses that exact truth question; its exception role and approved scope are retained.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4, printed pp. 311 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation. New York: Cambridge University Press, 2006, p. 91. Douglas Walton, Informal Logic. New York: Cambridge University Press, 1989, p. 89. Douglas Walton, Appeal to Popular Opinion. University Park: Pennsylvania State University Press, 1999, pp. 223-226.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: generally_true(A); CQ2 exception: doubting_general(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** popular_opinion/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 5. `pop_scheme`

**Entry:** 4.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The assertive variant concludes A itself. A may be a signed proposition, so accepting a negative proposition does not require an extra truth-wrapper. No printed CQs in this numbered variant; underlying group acceptance is an explicit premise.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.1, printed pp. 311 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** pop_scheme/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 6. `pop_scheme_acceptance`

**Entry:** 4.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The advisory variant concludes that the implicit respondent should accept A. It no longer asserts both truth and acceptance in one inference. No printed CQs; this is the expressly separated acceptance reading.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.1, printed pp. 311 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** pop_scheme_acceptance/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 7. `position_to_know_ad_populum`

**Entry:** 4.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** G is the same group in consensus and epistemic-position premises; A is the same conclusion proposition. Replacing plausibly_true(A) with A is appropriate for a defeasible argument. No printed CQs for this variant; premise support carries the challenge path.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.2, printed pp. 311 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** position_to_know_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 8. `expert_opinion_ad_populum`

**Entry:** 4.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The expert group and consensus proposition are preserved. The source itself does not bind the proposition to a separate domain variable, so this review does not invent one. No printed CQs in this variant; a stronger domain-sensitive version could be added under a distinct interpretation.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.3, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** expert_opinion_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 9. `deliberation_ad_populum`

**Entry:** 4.4. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The deliberating group is also the group accepting A; the output is A, with defeasibility represented by argument evaluation. No printed CQs in this variant.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.4, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** deliberation_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 10. `moral_justification_ad_populum`

**Entry:** 4.5. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Good-group policy acceptance and the respondent's goal support should_accept(P); the positive good-group branch is represented. No printed CQs; respondent identity remains implicit, as in the text. Remaining scope: Document one implicit respondent per graph or encode recipient identity when combining different audiences. Follow-up: PR13 additionally notes that the source's (or should be) alternative is omitted from premise2. That omission must be stated in the selected reading; peer approval alone does not supply the missing branch or close the participant/context qualification.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Apply and verify the documented one-respondent contract, or carry respondent identity for mixed audiences; explicitly declare the selected branch and the omitted source alternative (or should be), with a source-grounded example of the chosen interpretation.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.5, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** moral_justification_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; peer13-comparison-first.md.

### 11. `moral_justification_ad_populum_excuse`

**Entry:** 4.6. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The record retains action and policy alternatives in composite predicates. It does not separately infer morally_acceptable(action) or acceptable(policy). No printed CQs. The speaker is implicit in group_member/1 and the conclusion. Remaining scope: Declare this as an opaque combined action/policy reading; split typed action and policy variants before claiming independent interoperability of each branch.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Choose the bounded combined action/policy reading explicitly; if both are offered independently, add typed variants and separate examples.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.6, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** moral_justification_ad_populum_excuse/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 12. `snob_appeal_ad_populum`

**Entry:** 4.7. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Group, prestige goal and membership warrant are present, but acceptance, property acquisition and object purchase are compressed. Only should_accept(A) is concluded. No printed CQs. P and O occur in the membership warrant but do not select a conclusion branch. Remaining scope: The Rolex reference case treats an object as A in accept(A); replace with a clearly typed acceptance example, and add property/purchase variants if those alternatives are claimed as covered.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Replace the object-as-proposition Rolex acceptance fixture with a typed acceptance proposition; add separately typed property/purchase variants if those source alternatives are claimed.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.7, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** snob_appeal_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 13. `appeal_to_vanity_ad_populum`

**Entry:** 4.8. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The respondent's action X, not admired group G or proposition A, is the conclusion object. Separating the second-person ought from a first-person practical goal is correct. No printed CQs. The source uses one implicit respondent.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.8, printed pp. 312 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** appeal_to_vanity_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 14. `rhetoric_ad_populum`

**Entry:** 4.9. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Consensus, valued membership and contrasting out-group acceptance refer to the same G and A. The conclusion is acceptance by the implicit respondent. No printed CQs.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.9, printed pp. 313 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rhetoric_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 15. `common_folk_ad_populum`

**Entry:** 4.10. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Speaker S is explicit in background(S) and being_credible(S); removing the old zero-argument forms prevents unrelated speakers from sharing these facts. No printed CQs. The audience remains the current audience of the case.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.10, printed pp. 313 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** common_folk_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 16. `common_folks_group_ad_populum`

**Entry:** 4.11. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The speaker belongs to the same background group as the respondent; accept_saying(S) retains the speaker identity. No printed CQs. Group G is consistently bound.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 4.11, printed pp. 313 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** common_folks_group_ad_populum/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 17. `popular_practice`

**Entry:** 5. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The popular-practice premise and current-case conclusion have the right roles. The minor template is the consequent of an anaphoric warrant, rather than a self-contained conditional. CQ1 is a premise challenge; CQ2 is represented by justified grounds as a default assumption, a defensible interpretation. Remaining scope: Clarify the warrant template and state that each graph has one current case; cases with different circumstances require distinct contextual practice terms.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Clarify the anaphoric warrant and require context-qualified practices; add a two-context adverse case demonstrating that one practice context cannot establish another.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 5, printed pp. 314 (PDF pages +10). Separately cited original: Douglas Walton , Fundamentals of Critical Argumentation. New York: Cambridge University Press, 2006, p. 93.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: popular_practice(A); CQ2 assumption: have_grounds_think(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** popular_practice/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 18. `from_example`

**Entry:** 6.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED: the generalization now depends on F and G only, using usually_properties(F,G). The example a remains in argument parameters, allowing convergent support and feeding illustration. CQ1 premise, CQ2/CQ3 defaults, CQ5 exception. CQ4 is a supported challenge to the generalization in its negation issue; one exceptional case does not by itself refute a defeasible generalization. Remaining scope: Keep this CQ4 policy explicit. Generic grounded examples are insufficient to calibrate how strong a generalization is; a policy-level strength example remains desirable. Follow-up: The PR13 follow-up additionally unifies illustration's same_properties/3 with this record's has_properties/3, because the same individual and two properties are asserted. The new composed case uses one conjunction fact to support this induction and then illustration. This interoperability improvement does not resolve the independent generalization-strength policy hold.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Specify how the book's exception question affects the strength of a usually-generalization; add a policy-level example where one counterexample does not automatically refute a statistical generalization.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 6.1, printed pp. 314 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995, p. 135.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: has_properties(A,P1,P2); CQ2 assumption: supports_generalization(A,P1,P2); CQ3 assumption: typical_example(A); CQ4 rebuttal: ¬usually_properties(P1,P2); CQ5 exception: impair_general(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** from_example/book; from_example/independent-two-examples-one-generalization; illustration/peer13-shared-conjunction-chains-from-example

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 19. `illustration`

**Entry:** 6.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The rule premise and validity conclusion depend on F and G rather than on the illustrating individual. The particular case supplies an instance of both properties. No printed CQs. Quantification over x in the rule is represented as an opaque rule proposition, not an inference that enumerates all x. Follow-up: PR13 follow-up confirms the individual/F/G conjunction is identical in entries6.1 and6.2. same_properties/3 is now canonically has_properties/3, with migration and an actual two-scheme chain plus wrong-individual/property controls. Generic premise-withdrawal controls isolate the actual selected schema after a positive control; full-library positive, label and exception checks retain their context, so valid rederivation is not forbidden.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 6.2, printed pp. 315 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 357-368.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** illustration/book; illustration/peer13-shared-conjunction-chains-from-example

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 20. `from_model`

**Entry:** 6.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** REPAIRED: the third premise is now imitation_warrant(model,action,imitators), preserving the complete conditional and its model. The imitators are the recipients of should_do. No printed CQs. The new adverse case prevents a warrant about one model from licensing a different model's argument. Follow-up: PR13 approved the older form while noting it retained only the conditional consequent. The independent source-supported model-specific conditional warrant repair and wrong-model case remain required; peer approval of an earlier representation is not a reason to revert them.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 6.3, printed pp. 315 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 357-368.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** from_model/book; from_model/independent-conditional-warrant

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 21. `from_antimodel`

**Entry:** 6.4. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The model, act and would-be imitators are consistently bound. The implementation concludes negated should_do, while the text says should not do. No printed CQs. Under a two-position recommendation policy these can be treated as opposing recommendations, but ordinary deontic logic distinguishes not obligatory from obligatory not. Remaining scope: Document the binary recommendation convention or introduce an explicit refrain-from-action predicate before claiming full deontic equivalence.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Adopt and test the binary recommendation convention, or introduce an explicit refrain-from-action conclusion that distinguishes ought-not from not-ought.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 6.4, printed pp. 315 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 357-368.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** from_antimodel/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 22. `analogy`

**Entry:** 7. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The source and current cases are distinguished by an implicit current-case convention. The conclusion is A in that current case. CQ1 relevant differences and CQ2 base premise are present. CQ3 in the book asks about another similar contrary case, whereas more_on_point imposes a stronger relevance condition. Remaining scope: Own-source form may justify the inherited variant, but this is not yet a complete equivalence to the book CQ3; keep that scope difference explicit.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Fundamentals of Critical Argumentation, Cambridge University Press, New York 2006, p. 96-97.

**Next concrete action:** Compare the cited 2006 source; decide whether the inherited more-on-point defeater is retained as a variant or supplemented by the book's similar contrary-case CQ3, with an adverse case.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 7, printed pp. 315 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation, Cambridge University Press, New York 2006, p. 96-97.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 exception: relevant_differences(C); CQ2 premise: in_case(A,C); CQ3 exception: more_on_point(A,C)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** analogy/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 23. `practical_analogy_positive`

**Entry:** 8.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The action X is preserved while the situation changes from A1 to similar A2. The conclusion remains a normative assessment in A2. No printed CQs in these practical analogy variants.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 8.1, printed pp. 316 (PDF pages +10). Separately cited original: Douglas Walton, Informal Logic. New York: Cambridge University Press, 1989. p.257.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** practical_analogy_positive/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 24. `practical_analogy_negative`

**Entry:** 8.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The negative assessment of doing X transfers from A1 to similar A2; it does not change the action or assert the opposite descriptive proposition. No printed CQs.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 8.2, printed pp. 316 (PDF pages +10). Separately cited original: Douglas Walton, Informal Logic. New York: Cambridge University Press, 1989. p.257.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** practical_analogy_negative/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 25. `composition_generic`

**Entry:** 9.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** All parts have the property; the conclusion attributes it to the whole. Sharing has_property with other schemes is correct when object and property identifiers denote the same things. The compositional-heredity default follows the direction established in the narrative; the compendium parenthesis reverses the direction and is appropriately treated as a supported correction.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 9.1, printed pp. 316, 113 (PDF pages +10). Separately cited original: Douglas Walton, Informal Logic. New York: Cambridge University Press, 1989, p.130.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: compositionally_hereditary(A,P)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** composition_generic/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 26. `part_in_whole`

**Entry:** 9.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED: lower degree is no longer an exception defeating possession of the property. The positive species/part-to-property branch concludes has_property(y,A). No printed CQ exists here. The removed exception was a conclusion parenthesis incorrectly promoted to a defeat condition. Remaining scope: The current predicate species_of represents the species reading; physical part and comparative degree variants need explicit scope or separate records. Do not claim the degree parenthesis has been fully computed.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Declare the repaired positive species branch as the supported variant; add separate physical-part and degree variants only if those alternatives are included in release coverage.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 9.2, printed pp. 316 (PDF pages +10). Separately cited original: Ch. Perelman and L.Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp.231-233,234-241.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** part_in_whole/book; part_in_whole/independent-lower-degree-is-not-a-defeater

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 27. `division_generic`

**Entry:** 10.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The whole has a property and, defeasibly, all parts have it. The direction is distinct from composition even though they can form an engine-safe cycle. The divisional-heredity default is the substantive warrant; it is challengeable. The source's reversed explanatory parenthesis is corrected using narrative context. Follow-up: PR13's full-library observation is correct: an undisputed divisional-heredity default can transfer a stipulated whole-property fact to all parts. This is the declared defeasible policy, not evidence of universal heredity. Usage guidance must retain whole/context identity, deliberate scheme selection and questioning conventions.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 10.1, printed pp. 317, 113 (PDF pages +10). Separately cited original: Douglas Walton, Argument: The Logic of Fallacies, Toronto: McGraw-Hill Ryerson, 1982, pp.206-208.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: divisionally_hereditary(A,P)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** division_generic/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/using.md; peer13-comparison-first.md.

### 28. `whole_into_parts`

**Entry:** 10.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The implementation fixes three named parts and derives the whole's property from X1 plus the stated whole/part warrant. The source allows x1 through xn and alternative negative branches. No printed CQs. The source itself says only if; treating this as a presumptive argumentative pattern does not establish a deductively valid sufficient condition. Remaining scope: Declare the positive, three-part restricted instance. Arbitrary cardinality, other selected parts, negative branches and the only-if interpretation remain unimplemented or unresolved.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Declare the positive, three-part instance explicitly and test that contract; add or exclude arbitrary cardinality, selected-part, negative and only-if readings separately.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 10.2, printed pp. 317 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 231-233,234-241.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** whole_into_parts/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 29. `opposition_descriptive_1`

**Entry:** 11.1a. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The contrary-property warrant and asserted property concern the same subject A; the conclusion is inability to present the opposite at the same time. No printed CQs. Modal cannot-present is kept distinct from a plain negated observation.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 11.1a, printed pp. 317 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** opposition_descriptive_1/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 30. `opposition_descriptive_2`

**Entry:** 11.1b. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The negative property premise instantiates the antecedent of the stated opposition warrant; the conclusion names the paired property of the same subject. No printed CQs. The warrant, not absence of any arbitrary property, licenses the inference.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 11.1b, printed pp. 318 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** opposition_descriptive_2/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 31. `opposition_descriptive_3`

**Entry:** 11.1c. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Complementary-or-contrary alternatives and absence of one yield the required other predicate for the same subject and time context. No printed CQs. This is the declared paired-property interpretation, not unrestricted closed-world reasoning.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 11.1c, printed pp. 318 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** opposition_descriptive_3/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 32. `opposition_normative`

**Entry:** 11.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The corrected minor premise instantiates the first evaluation in the major warrant; the conclusion excludes the opposite evaluation for the same subject. No printed CQs. The source's inconsistent symbol is corrected coherently rather than reproduced as a swapped role.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 11.2, printed pp. 318 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** opposition_normative/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 33. `rhetorical_oppositions_normative`

**Entry:** 12.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Entity and treatment pairs retain the correct direction: the current required treatment triggers the warranted contrary treatment for the contrary entity. No printed CQs.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 12.1, printed pp. 318 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rhetorical_oppositions_normative/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 34. `rhetorical_oppositions_descriptive`

**Entry:** 12.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A's relation to B and the paired converse relation yield B's converse relation to A. Roles are reversed exactly once. No printed CQs. The one implemented direction is warranted by the source's bidirectional description. Follow-up: PR13 follow-up corrected the ambiguous major-premise template against printed p318: the latter participant bears the converse relation to the former, and contrariwise. The original conclusion already had the correct reversed arguments. Approval applies to the corrected template.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 12.2, printed pp. 318 (PDF pages +10).

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rhetorical_oppositions_descriptive/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 35. `alternative_cognitive`

**Entry:** 13.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The chosen alternatives are propositions A and B; negated A and the alternative warrant yield B. Bare propositions allow this output to feed other schemes. No printed CQs. The rule remains a defeasible argument rather than a modal theorem about all possible worlds.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 13.1, printed pp. 318 (PDF pages +10). Separately cited original: Manfred Kienpointner, Alltagslogik, Struktur und Funktion von Argumentationsmustern. Stuttgart-Bad Cannstatt: Frommann-Holzboog, 1992, p.306.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** alternative_cognitive/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 36. `alternative_normative`

**Entry:** 13.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** A and B remain the same normative alternatives. Required A yields negated required B under the supplied alternative warrant. No printed CQs. This reading requires mutually exclusive normative alternatives; an inclusive A-or-B would not suffice. Remaining scope: Document exclusivity and distinguish not-required(B) from required(not-B).

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Specify and test an exclusive-alternatives recommendation policy; preserve the distinction between absence of obligation and an obligation to refrain.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 13.2, printed pp. 318 (PDF pages +10). Separately cited original: Manfred Kienpointner, Alltagslogik, Struktur und Funktion von Argumentationsmustern. Stuttgart-Bad Cannstatt: Frommann-Holzboog, 1992, p.306.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** alternative_normative/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 37. `verbal_classification`

**Entry:** 14. **Origin:** upstream. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Object O classified under F and F subsumed by G yield O under G. Classification hierarchy and object positions are coherent. CQ1 is the object-class premise; CQ2 is the doubtful-classification exception. Final provenance correction: upstream code has the same O/F/G variables, premises and conclusion, with no cited external source. The added doubtful_classification(F,G) is precisely the compendium's CQ2 and has an existing migration row. The earlier generic unavailable-original hold was mistaken; declared compendium/code-origin approval is now complete, without claiming to have inspected Walton (2006).

**Source-access obligation:** No outstanding access obligation for the declared compendium interpretation. Upstream code had no source citation; the 2006 work in the book's bibliography was not inspected and is not claimed as an independently verified transcription.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 14, printed pp. 319 (PDF pages +10). Upstream code comparison: variables/premises/conclusion unchanged; added book-CQ2 exception documented in tools/migration_r2.yml.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Inherited code cited no separate work; its current declared source is the supplied compendium. No unseen external original is certified. No distinct external-source CQ inventory is claimed; use the compendium count for the declared interpretation.

**Positive-numbered CQ mappings:** CQ1 premise: instance(O,F); CQ2 exception: doubtful_classification(F,G)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** verbal_classification/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; final-disposition/provenance-addendum-2026-11-21.md.

### 38. `definition_to_verbal_classification`

**Entry:** 15. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** book form approved; own-original-source hold. **Book-form decision:** approve.

**Approval or hold reason:** REPAIRED: the definition satisfaction and definition-to-class link yield the correct object classification; the complete book CQ set is represented. The source's inconsistent x/a in its universal premise is read coherently, not as all objects having G. CQ1 inadequacy remains an exception. CQ2 now has doubtful_definition(D,G), explicitly for a merely stipulative or biased definition subject to doubt. Positive and wrong-definition controls distinguish relevant from unrelated challenges. Book-form approval does not verify its separately cited original; preserve the own-original-source gate.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Fundamentals of Critical Argumentation, Cambridge University Press, New York 2006, p. 129.

**Next concrete action:** Obtain and compare the separately cited original formulation and questions; reconcile any differences explicitly before promoting this inherited record.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 15, printed pp. 319 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation, Cambridge University Press, New York 2006, p. 129.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 exception: inadequate_definition(D,G); CQ2 exception: doubtful_definition(D,G)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** definition_to_verbal_classification/book; definition_to_verbal_classification/independent-biased-definition-defeats; definition_to_verbal_classification/independent-other-definition-does-not-defeat

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 39. `verbal_classification_vagueness`

**Entry:** 16. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED wording: the counterclaim concerns precision, not certainty. The argument and property roles in the vagueness/precision premises are coherent. CQ1/CQ2 are premises; CQ3 is an opposing exception asserting that precision is not required here. This is an interpretation of a why-question, not the only possible burden policy. Remaining scope: Approve the content correction, but document when an unanswered precision challenge defeats the argument versus requiring an affirmative exception.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Adopt an explicit challenge burden: determine whether an unanswered precision challenge suspends the argument or requires affirmative contrary evidence, and test the chosen behavior.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 16, printed pp. 319, 320 (PDF pages +10). Separately cited original: Douglas Walton, Argumentation Schemes for Presumptive Reasoning. Mahwah, NJ.: Erlbaum, 1996, pp.102-103.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: precision_needed(A); CQ2 premise: property_vagueness(F,A); CQ3 exception: certainty_no_need(A,F)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** verbal_classification_vagueness/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 40. `verbal_classification_arbitrariness`

**Entry:** 17. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The arbitrary property, relevant argument and need for a nonarbitrary definition retain their roles; the conclusion rejects that argument as deficient. CQ1/CQ2 are premises; CQ3 harmless arbitrariness is an exception. The burden choice is a project interpretation. Remaining scope: Keep that burden policy explicit; no extra premise should simply restate the conclusion.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Adopt and test the burden for an unanswered arbitrariness challenge, without adding a premise that merely restates the desired conclusion.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 17, printed pp. 320 (PDF pages +10). Separately cited original: Douglas Walton, Argumentation Schemes for Presumptive Reasoning. Mahwah, NJ.: Erlbaum, 1996, pp.104-105.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: needed_nonarbitrary(A,F); CQ2 premise: arbitrary_property(F,A); CQ3 exception: arbitrariness_harmless(A,F)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** verbal_classification_arbitrariness/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 41. `interaction_act_person_1`

**Entry:** 18.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A person's act with value V supports attributing V to that person. Agent, act and value remain separate. No printed CQs. Follow-up: PR13 follow-up confirms that this record's person_value/2 conclusion and entry18.2's person_judgment_value/2 premise express the same person/judgment-value relation on p321. The consumer is canonicalized to person_value/2; this producer formula is unchanged. A two-scheme case checks chaining without mixing people or values.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 18.1, printed pp. 320 (PDF pages +10). Separately cited original: Ch. Perelman, L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 296-305.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** interaction_act_person_1/book; interaction_act_person_2/peer13-shared-person-valuation-chains

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 42. `interaction_act_person_2`

**Entry:** 18.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A prior judgment of the person's value plus a stated value-to-act evaluation warrant yields the evaluation of that person's act. No printed CQs. The generic warrant is intentionally not tied to one individual unless instantiated through the first premise. Follow-up: PR13 follow-up canonically replaces person_judgment_value/2 with person_value/2: printed p321 uses the same person/judgment-value assertion as entry18.1's conclusion. Migration and an actual two-scheme case preserve the evaluation-context convention and reject wrong-person/wrong-value instances.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 18.2, printed pp. 320, 321 (PDF pages +10). Separately cited original: Ch. Perelman, L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 296-305.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** interaction_act_person_2/book; interaction_act_person_2/peer13-shared-person-valuation-chains

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 43. `positive_value`

**Entry:** 19.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Agent A's judgment and its effect on A's goal lead to a reason for retaining the goal. The conclusion omits A. No printed CQs. Two people with an identical goal term but different evaluations would share the conclusion's identity. Remaining scope: Require goal terms to identify the owner's goal instance or add the agent to the conclusion before mixed-agent value comparisons.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 19.1, printed pp. 321 (PDF pages +10). Separately cited original: Trevor Bench-Capon, "Persuasion in Practical Argument Using Value-based Argumentation Frameworks." Journal of Logic and Computation 13 (2003), pp.429­448. Trevor Bench-Capon, "Agreeing to Differ: Modelling Persuasive Dialogue between Parties without a Consensus about Values." Informal Logic 22 (2003a), pp. 231 -245.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** positive_value/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 44. `negative_value`

**Entry:** 19.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The corrected conclusion retracts commitment, rather than retaining it. As with the positive record, the conclusion omits the evaluating agent. No printed CQs. Remaining scope: Owner-scoped goal identifiers or an explicit agent argument are necessary for multi-agent use; this is separate from the already corrected sign.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use owner-qualified goal terms or add an owner argument to the resulting commitment; add a mixed-agent adverse case independently of the repaired polarity.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 19.2, printed pp. 321 (PDF pages +10). Separately cited original: Trevor Bench-Capon, "Persuasion in Practical Argument Using Value-based Argumentation Frameworks." Journal of Logic and Computation 13 (2003), pp.429­448. Trevor Bench-Capon, "Agreeing to Differ: Modelling Persuasive Dialogue between Parties without a Consensus about Values." Informal Logic 22 (2003a), pp. 231 -245.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_value/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 45. `sacrifice`

**Entry:** 20. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The person, sacrifice and its value are coherent in the great-sacrifice branch. The current record does not encode the small-sacrifice alternative printed parenthetically. No printed CQs. Remaining scope: Approve only the great-sacrifice variant; add or explicitly exclude the small-sacrifice branch from coverage claims.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Release only the documented great-sacrifice variant, or add the source's small-sacrifice alternative with separate semantics and examples.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 20, printed pp. 322 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 248-254.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** sacrifice/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 46. `group_1`

**Entry:** 21.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A member's quality plus an explicit transfer warrant yields the group's quality. The has_property(group,quality) identity is the same property statement used elsewhere. No printed CQs. The warrant, not membership alone, supports the transfer.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 21.1, printed pp. 321, 322 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 321-326.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** group_1/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 47. `group_2`

**Entry:** 21.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A group's quality, membership and explicit downward transfer warrant yield the member's quality. The shared property predicate preserves entity roles. No printed CQs. The membership and transfer premises prevent an unqualified automatic group-to-person inference.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 21.2, printed pp. 322 (PDF pages +10). Separately cited original: Ch. Perelman and L. Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation. Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 321-326.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** group_2/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 48. `practical_inference`

**Entry:** 22.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Agent P's goal and means-to-end support should_do(P,A); explicit agent binding makes the shared practical conclusion appropriate. All five printed questions have a ledger route, but other/better means and bad consequences are opponent-proven exceptions. Narrative p.96 says merely asking any of the five questions can provisionally defeat until answered. Remaining scope: A static evidential exception policy is not the book's full dialogue burden behavior. Keep interpretation approval qualified until that mismatch is resolved or explicitly contracted.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Choose and document static evidential exceptions versus the source's question-induced temporary suspension; test the selected contract and avoid claiming full dialogue behavior if only static evaluation is supported.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.1, printed pp. 95, 96, 323 (PDF pages +10). Separately cited original: Trevor Bench-Capon, "Persuasion in Practical Argument Using Value-based Argumentation Frameworks." Journal of Logic and Computation 13 (2003), pp. 429-448. Douglas Walton, Practical Reasoning. Savage, Md.: Rowman and Littlefield, 1990. Douglas Walton, "Actions and Inconsistency: The Closure Problem of Practical Reasoning. "In Contemporary Action Theory, vol.1 , ed. Ghita Holmstrom-Hintikka and Raimo Tuomela. Dordrecht: Kluwer, 1997, p. 164. Douglas Walton, Slippery Slope Arguments. Newport News, Va.: Vale Press, 1992, pp. 89-90.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: conflicting_goals(P,G); CQ2 exception: more_efficient_alternative(A,G); CQ3 exception: more_efficient_alternative(A,G); CQ4 assumption: possible_bring_about(P,A); CQ5 exception: incompatible_consequences(A,G)

**Additional `n: 0` hooks:** assumption: possible_goal(G)

**Example evidence:** practical_inference/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 49. `practical_necessary_condition`

**Entry:** 22.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED: P now scopes all personal selection, consideration, feasibility and preference premises, plus personal acceptability and alternatives. A second agent sharing a goal cannot inherit another's deliberation. CQ coverage exists; exceptions versus questions retain the practical dialogue-burden discrepancy described for 22.1. Remaining scope: New positive and two-agent adverse cases pass after the patch. Resolve question-induced suspension policy before unqualified source approval.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the repaired agent scope and passing two-agent controls; resolve and test the source's question-induced suspension policy before unqualified interpretation approval.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.2, printed pp. 95, 96, 323, 324 (PDF pages +10). Separately cited original: Trevor Bench-Capon, "Persuasion in Practical Argument Using Value-based Argumentation Frameworks." Journal of Logic and Computation 13 (2003), pp. 429-448. Douglas Walton, Practical Reasoning. Savage, Md.: Rowman and Littlefield, 1990. Douglas Walton, "Actions and Inconsistency: The Closure Problem of Practical Reasoning. "In Contemporary Action Theory, vol.1 , ed. Ghita Holmstrom-Hintikka and Raimo Tuomela. Dordrecht: Kluwer, 1997, p. 164. Douglas Walton, Slippery Slope Arguments. Newport News, Va.: Vale Press, 1992, pp. 89-90.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: alternative_means(P,A,X); CQ2 assumption: acceptable_action(P,X); CQ3 assumption: possible_bring_about(P,X); CQ4 exception: side_effect_action(P,X); CQ5 exception: conflicting_goals(P,A)

**Additional `n: 0` hooks:** assumption: possible_goal(A)

**Example evidence:** practical_necessary_condition/book; practical_necessary_condition/independent-one-agents-selection-does-not-bind-another; practical_necessary_condition/independent-another-agents-alternative-does-not-defeat

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 50. `practical_sufficient_condition`

**Entry:** 22.3. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED: all personal premises bind the same P; the alternatives template now says each member is sufficient, matching p.96 rather than merely one member. Selected member X is distinct from set S. Same five question routes and extra feasibility defaults as 22.2; dialogue-burden qualification remains. Remaining scope: New positive and two-agent adverse cases pass after the patch. Document the narrative-supported correction and resolve provisional challenge behavior.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the narrative-supported each-one sufficient reading and repaired agent scope; resolve and test provisional challenge behavior before unqualified interpretation approval.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.3, printed pp. 96, 323, 324 (PDF pages +10). Separately cited original: Trevor Bench-Capon, "Persuasion in Practical Argument Using Value-based Argumentation Frameworks." Journal of Logic and Computation 13 (2003), pp. 429-448. Douglas Walton, Practical Reasoning. Savage, Md.: Rowman and Littlefield, 1990. Douglas Walton, "Actions and Inconsistency: The Closure Problem of Practical Reasoning. "In Contemporary Action Theory, vol.1 , ed. Ghita Holmstrom-Hintikka and Raimo Tuomela. Dordrecht: Kluwer, 1997, p. 164. Douglas Walton, Slippery Slope Arguments. Newport News, Va.: Vale Press, 1992, pp. 89-90.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: alternative_means(P,A,X); CQ2 assumption: acceptable_action(P,X); CQ3 assumption: possible_bring_about(P,X); CQ4 exception: side_effect_action(P,X); CQ5 exception: conflicting_goals(P,A)

**Additional `n: 0` hooks:** assumption: possible_goal(A)

**Example evidence:** practical_sufficient_condition/book; practical_sufficient_condition/independent-one-agents-selection-does-not-bind-another; practical_sufficient_condition/independent-another-agents-alternative-does-not-defeat

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 51. `value_based_practical_reasoning`

**Entry:** 22.4. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Agent P, goal G, value set V and act A are connected through the means premise; practical conclusion should_do(P,A) is coherent. All seven printed CQs are mapped. Value-based preference and alternative facts omit P and rely on goal/action terms carrying context; practical challenge burden remains a policy interpretation. Remaining scope: Use owner/context-specific goals and value sets, and settle the explicit challenge policy before full mixed-agent semantic approval.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Require owner/context-specific goals and value sets; settle the challenge burden and add a mixed-agent/context adverse example.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.4, printed pp. 95, 96, 324 (PDF pages +10).

**Declared-source CQ count:** 7. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: conflicting_goals(P,G); CQ2 premise: goal_supported_by_values(G,V); CQ3 exception: alternative_means(G,A); CQ4 exception: more_efficient_alternative(A,G); CQ5 exception: better_alternative_by_values(A,V); CQ6 assumption: possible_bring_about(P,A); CQ7 exception: worse_consequences(A,G)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** value_based_practical_reasoning/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 52. `practical_goal`

**Entry:** 22.5. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The same person's goal G and action X's contribution to that goal yield should_do(P,X). The named agent makes vocabulary sharing appropriate. No printed CQs in this variant. This is a defeasible reason to act, not a proof that every contributing action must be jointly performed.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.5, printed pp. 325 (PDF pages +10). Separately cited original: Bart Verheij, "Dialectical Argumentation with Argumentation Schemes: An Approach to Legal Logic." Artificial Intelligence and Law 11 (2003), p. 169.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** practical_goal/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 53. `end_and_means_1`

**Entry:** 22.6.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The record represents the positive branch: X is a means to good end Y, so X is good. The source also prints less-bad and bad alternatives. No printed CQs. Evaluation is context-sensitive; the source is presumptive, not unrestricted moral closure. Remaining scope: Approve only the explicitly positive branch; missing parenthetical variants cannot be counted as tested.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Declare the positive branch as the released variant; add or explicitly exclude the source's parenthetical alternatives and test any added branch independently.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.6.1, printed pp. 325 (PDF pages +10). Separately cited original: Chaim Perelman and Lucie Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 273-278.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** end_and_means_1/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 54. `end_and_means_2`

**Entry:** 22.6.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The ends-importance warrant and means-to-end premise preserve the positive branch. The record's importance warrant is zero-argument and therefore global within a graph. No printed CQs. Remaining scope: Limit this to one evaluation context and the positive branch; opposite valuation and mixed-context warrants need separate representation.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Declare one evaluation context and the positive branch; add scoped warrants for mixed contexts and separate opposite-valuation examples if those are claimed.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.6.2, printed pp. 325 (PDF pages +10). Separately cited original: Chaim Perelman and Lucie Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 273-278.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** end_and_means_2/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 55. `end_and_means_3`

**Entry:** 22.6.3. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Ease of achieving the means supports goodness of the end in the positive branch. Difficult-means/bad-end is not represented. No printed CQs. Remaining scope: Approve only the easy/good branch and record the omitted alternative explicitly.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Declare the easy/good branch as the released variant; add or explicitly exclude the source's alternative branch.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 22.6.3, printed pp. 325 (PDF pages +10). Separately cited original: Chaim Perelman and Lucie Olbrechts-Tyteca, The New Rhetoric: A Treatise on Argumentation Notre Dame, Ind.: University of Notre Dame Press, 1969, pp. 273-278.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** end_and_means_3/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 56. `two_person_practical`

**Entry:** 23. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** A states an intention to B; B supplies necessity information about action Y for A's intention X; the conclusion directs A, not B, to act. All five printed questions are sorted. Alternative means and additional actions omit the deliberating agent and assume one current case. Remaining scope: Correct speaker/recipient roles are approved, but practical dialogue burdens and cross-agent availability of alternatives need an explicit contract or scoped facts. Follow-up: PR13 follow-up corrected sets to sees using visually checked printed pp97,325, preserving the adviser-to-agent communication. This closes the wording defect; its separate dialogue and alternative-action scope qualifications remain.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Keep the repaired speaker/recipient roles; define practical-dialogue burdens and scope alternative actions to the deliberating parties, with an adverse multi-party example.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 23, printed pp. 325, 326 (PDF pages +10). Separately cited original: Douglas Walton, Appeal to Expert Opinion. University Park: Pennsylvania State University Press. 1997 p. 163.

**Declared-source CQ count:** 5. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: conflicting_goals(A,X); CQ2 exception: alternative_means(X,Y); CQ3 exception: side_effect_other_act(Y,A); CQ4 assumption: possible_bring_about(A,Y); CQ5 exception: other_actions_required(X,Y)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** two_person_practical/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 57. `waste`

**Entry:** 24. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Agent A's stopping would waste efforts toward goal X; avoiding that waste supports continuing toward X. Feasibility concerns bringing about X. CQ1 is the possible-goal default; CQ2 is reassessment against continuing, rather than an automatic rule that sunk costs justify every continuation.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 24, printed pp. 326 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, Tuscaloosa and London: University of Alabama Press, 1995, p. 157. Douglas Walton "The Sunk Cost Fallacy or Argument from Waste" Argumentation 16 (2002), p.488.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: possible_goal(X); CQ2 exception: reassessment_against_continuing(A,X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** waste/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 58. `sunk_costs`

**Entry:** 25. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The legacy quantity/feasibility form is not the article's printed p. 489 precommitment scheme or the compendium form. Root has now inspected the article and confirmed that distinction. Its quantity/waste, feasibility and future-loss policies draw on separate discussion and practical-inference questions. Keep it a qualified adaptation, distinct from the approved book-form sunk_costs_precommitment; missing article access is no longer a hold reason.

**Source-access obligation:** No missing-article blocker: root accessed the original article as web text with printed-page markers. No successful PDF retrieval/hash/visual inspection is claimed; retain access-method provenance and resolve adaptation policy.

**Next concrete action:** Retain the legacy quantity/feasibility record as a documented adaptation, distinct from sunk_costs_precommitment; settle the selected practical-question and future-loss policy rather than treating the legacy form as the article's printed p. 489 scheme.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 25, printed pp. 327 (PDF pages +10). Separately cited original: Douglas Walton, ‘The Sunk Costs Fallacy or Argument from Waste’, Argumentation, 16, 2002, p. 489 Root checked author article text: https://www.academia.edu/13186555/The_Sunk_Costs_Fallacy_or_Argument_from_Waste ; printed pp. 488,489,493.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Root inspected the original author's article text with printed-page markers: p. 488 waste, p. 489 precommitment, p. 493 practical-inference discussion. Legacy record is an adaptation, not the printed p. 489 form. The article was accessed; a record-specific CQ inventory for this composite legacy adaptation is not established. Compendium entry 25 has zero printed CQs; that is not a claim that the article discusses no questions.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** assumption: feasible(A); exception: future_losses_outweigh(A)

**Example evidence:** sunk_costs/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 59. `sunk_costs_precommitment`

**Entry:** 25. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The choice at T2, prior commitment through T1, acting agent P and chosen A retain the roles of both printed premises; should_do(P,A) is the proper practical conclusion. No printed CQs for this numbered form. Earlier/later ordering is part of the precommitted relation's meaning; arbitrary timestamps are not separately reasoned about.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations. This book-form approval does not certify complete practical-question coverage of the entire 2002 article.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 25, printed pp. 327 (PDF pages +10). Separately cited original: Douglas Walton, ‘The Sunk Costs Fallacy or Argument from Waste’, Argumentation 16, 2002, p. 489.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** sunk_costs_precommitment/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 60. `ignorance`

**Entry:** 26. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** Would-be-known(S) and not-known(S) yield not-S. This is the correct negative conclusion and is not a closed-world inference from mere absence. CQ1 investigation is an exception; CQ2 dialogue burden and CQ3 evidential strength are both marked standard. A proof-standard threshold does not itself decide the dialogue's probandum or burden holder. Remaining scope: Own-source approval remains separate; specify how dialogue burden/probandum is supplied by an application before claiming all CQs are operationally implemented.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 150

**Next concrete action:** Compare the cited 1995 source and specify application-supplied dialogue burden/probandum independently of the issue proof standard; add examples for the chosen burden contract.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 26, printed pp. 327 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 150

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 exception: uninvestigated(S); CQ2 standard: [no atom: application-level standard]; CQ3 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** ignorance/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 61. `negative_normal_expectations`

**Entry:** 26.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The warranted normal expectation of A and explicit not-A support a non-normal situation, identified by A in the current encoding. CQ1 investigation is represented. CQ2/CQ3 have the same burden/proof-standard distinction as ignorance. Remaining scope: Document one expectation context or add a situation identifier; distinguish dialogue burden allocation from strength of proof.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use one expectation context or add a situation identifier; specify dialogue burden allocation separately from required proof strength and add a mixed-context adverse example.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 26.1, printed pp. 327 (PDF pages +10). Separately cited original: Douglas Walton, Arguments from Ignorance. University Park: Pennsylvania State University Press, 1996, pp. 84,86. Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995. p. 150.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: uninvestigated(A); CQ2 standard: [no atom: application-level standard]; CQ3 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_normal_expectations/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 62. `negative_practical`

**Entry:** 26.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** REPAIRED: B>C now yields acting on the false presumption, explicitly. The opposite C>B branch is not silently collapsed into the same ambiguous action. CQ1 is an evidence-progress default. CQ2 burden allocation and CQ3 strength still require distinct external dialogue information. Remaining scope: Negative branch is source-correct and tested; declare the missing true-presumption branch and resolve dialogue metadata before unqualified approval. Follow-up: PR13 correctly found that the original /book fixture still paired loaded with reversed accident/delay consequences after the formula repair. It now uses unloaded: assuming unloaded falsely risks an accident; assuming not-unloaded falsely costs delay. The source locator explicitly identifies this as an illustrative application of p327, not a printed example. The source-correct negative branch remains held for the existing branch/burden scope.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the source-correct negative-presumption branch; explicitly exclude or add the true-presumption branch and supply a separate dialogue-burden/probandum contract.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 26.2, printed pp. 327 (PDF pages +10). Separately cited original: Douglas Walton, Arguments from Ignorance. University Park: Pennsylvania State University Press, 1996, pp. 84,86. Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995. p. 150.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: sufficient_evidence(A); CQ2 standard: [no atom: application-level standard]; CQ3 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_practical/book; negative_practical/independent-false-presumption-selected; negative_practical/independent-opposite-ranking-does-not-select-false

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md; docs/reviews/peer13-first-half-addendum.md; peer13-comparison-first.md.

### 63. `epistemic_ignorance`

**Entry:** 27. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** All five premises are present, including the bivalence of domain D. The proposition A and knowledge base K retain their membership roles; the conclusion is not-A. No printed CQs for this separate schema. The awkward predicate name true_propositions_contained has negative natural-language content, but it is an opaque fact, not a computed double negation.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 27, printed pp. 328 (PDF pages +10). Separately cited original: Douglas Walton , "Nonfallacious Arguments from Ignorance." American Philosophical Quarterly 29 (1992), p. 386.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** epistemic_ignorance/book

**Review reference:** book-first-audit/adjudication.json; first-half-independent-review.md.

### 64. `cause_to_effect`

**Entry:** 28. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The two-premise causal inference and effect conclusion correspond to entry 28. The inherited form makes one occurred event and causal relation explicit; the printed may/might modality is implemented as defeasible support, not a probability value. CQ1–CQ2 challenge the supplied causal generalization; CQ3 is interference(E1). The latter is broad: it must mean an interference with this causal application, not every possible effect of E1. The actual outstanding issue is single-causal-application scope. No distinct original work was cited by the inherited code; the compendium is the declared source.

**Source-access obligation:** No separate external work is cited by the inherited code. The supplied compendium was inspected; the remaining hold is causal-context interpretation, not acquisition of an unidentified book.

**Next concrete action:** Bind interference to the particular causal application, or adopt and test a single-causal-context contract; do not invent an unavailable original source that the inherited code never cited.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 28, printed pp. 328 (PDF pages +10).

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Inherited code cited no separate work; its current declared source is the supplied compendium. No unseen external original is certified. No distinct external-source CQ inventory is claimed; use the compendium count for the declared interpretation.

**Positive-numbered CQ mappings:** CQ1 premise: causes(E1,E2); CQ2 premise: causes(E1,E2); CQ3 exception: interference(E1)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** cause_to_effect/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 65. `correlation_to_cause`

**Entry:** 29. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The direction correlation(E1,E2) → causes(E1,E2) follows entry 29. It is defeasible, not a demonstration that correlation alone establishes causation. CQ1 is the correlation premise; CQ3 is a common-cause exception. CQ2 is implemented by a challengeable explanatory theory. That is a selected evidential policy for non-coincidence, not a literal equivalence between the two. Remaining scope: Original cited 1995 work not independently read. Preserve the explanatory-theory policy qualification and require an application to justify causal evidence.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 142.

**Next concrete action:** Compare the cited 1995 source; retain an explicit explanatory-theory policy and require causal evidence, with a case distinguishing correlation alone from a justified causal application.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 29, printed pp. 328, 329 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 142.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: correlated(E1,E2); CQ2 assumption: explanatory_theory(E1,E2); CQ3 exception: causes_both(E1,E2)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** correlation_to_cause/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 66. `sign`

**Entry:** 30. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The sign and indicated event are correctly separated in the ordinary premises and conclusion. The comparative alternative defeater identifies the sign alone. CQ1 is pair-scoped strong_correlation(A,B). CQ2 more_reliable(A) compresses the competing-event comparison; with several candidate events for one sign it may suppress every candidate, including an independently better one. Remaining scope: Before unrestricted multi-interpretation use, scope the comparator to the candidate (A,B) or impose and test a one-candidate-per-sign case contract. Keep below blanket example-passed approval.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Scope the competing-interpretation comparator to candidate (A,B), or enforce and test one candidate per sign context; include a wrong-candidate adverse case.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 30, printed pp. 329 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation. New York: Cambridge University Press, 2006, pp. 113-114.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: strong_correlation(A,B); CQ2 exception: more_reliable(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** sign/book; sign/two-signs; sign/a-sign-of-a-sign

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 67. `abductive_backward`

**Entry:** 31.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The selected account, data set and best-explanation conclusion follow 31.1. The available-accounts set A is named, but A1 is not explicitly linked to membership in that set. CQ1 is sufficient_account(A1,D); CQ2 is the comparative best-account premise. CQ3–CQ4 share inquiry_incomplete(D), an explicit aggregate policy about whether to conclude now. Remaining scope: Document and enforce one candidate universe per D or add an explicit membership/comparison link before approving unrestricted multi-inquiry use.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Add candidate membership/comparison links for each inquiry D, or enforce and test one candidate universe per D before unrestricted multi-inquiry use.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 31.1, printed pp. 329, 330 (PDF pages +10).

**Declared-source CQ count:** 4. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: sufficient_account(A1,D); CQ2 premise: successful_account(A1,D); CQ3 exception: inquiry_incomplete(D); CQ4 exception: inquiry_incomplete(D)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** abductive_backward/book; abductive_backward/positive; abductive_backward/adverse-inquiry-incomplete

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 68. `abductive_forward`

**Entry:** 31.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. Printed pp. 329–330 distinguish diagram G1 from conclusion C1. The old unary successful_diagram allowed the diagram itself to be concluded as a hypothesis. Selection now names G and G1 and an explicit diagram_conclusion(G1,C1) supplies the conclusion. CQ1 remains a challengeable account-adequacy aggregate; CQ2 is selection of the strongest diagram within G; CQ3–CQ4 retain the inquiry-completeness aggregate. These are declared defeasible interpretations, not numeric graph-strength calculations.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 31.2, printed pp. 329, 330 (PDF pages +10).

**Declared-source CQ count:** 4. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: sufficient_account(C1,D); CQ2 premise: successful_diagram(G,G1); CQ3 exception: inquiry_incomplete(D); CQ4 exception: inquiry_incomplete(D)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** abductive_forward/book; abductive_forward/independent-diagram-and-conclusion; abductive_forward/independent-wrong-diagram-set; abductive_forward/independent-wrong-conclusion-link; abductive_forward/positive; abductive_forward/adverse-account-questioned; abductive_forward/adverse-inquiry-incomplete

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 69. `abductive_action_to_character`

**Entry:** 31.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. The person A, performed action X and particular quality Q are explicit. The typed conclusion character_quality(A,Q) preserves what kind of property the source asserts; its additional has_property projection makes the trait available to generic trait reasoning. CQ1 is the named quality in the action premise. CQ2 is quality_defined(Q); CQ3 is action_fits_quality(A,X,Q). A challenge concerning one quality or action does not suspend another.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 31.3, printed pp. 330 (PDF pages +10).

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: action_character(A,X,Q); CQ2 assumption: quality_defined(Q); CQ3 assumption: action_fits_quality(A,X,Q)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** abductive_action_to_character/book; abductive_action_to_character/independent-quality-and-prediction-chain; abductive_action_to_character/independent-fit-is-action-and-quality-specific; abductive_action_to_character/independent-quality-definition

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 70. `abductive_character_to_action`

**Entry:** 31.4. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. The ordinary premise requires a typed character_quality(A,Q). The conclusion remains the source conditional prediction about future conduct with that quality; it does not identify any specific past actor/event. The quality is explicit in CQ1 and CQ2. The counterexample defeater is scoped to (A,Q). The definition remains a defeasible default that can be challenged. Integration clarification resolves the cross-review concern: the defeater denotes relevant counterevidence sufficient to undermine the qualitative prediction, not an isolated atypical act automatically refuting a statistical tendency; evidence adequacy is external.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 31.4, printed pp. 330 (PDF pages +10).

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: character_quality(A,Q); CQ2 assumption: quality_defined(Q); CQ3 exception: not_character_action(A,Q)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** abductive_character_to_action/book; abductive_character_to_action/independent-counterexample-is-quality-specific; abductive_character_to_action/independent-ordinary-property-is-not-character

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 71. `retroductive`

**Entry:** 31.5. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. Event E is now present in candidate-agent and event-quality premises, the conclusion, the fit question and reference class. The source distinguishes this backward attribution from forward character prediction. CQ1 identifies the quality through event_quality(E,Q); CQ2 concerns that Q; CQ3 concerns A/E/Q; CQ4 concerns other agents who could have brought about this E and possess Q.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 31.5, printed pp. 331 (PDF pages +10).

**Declared-source CQ count:** 4. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: event_quality(E,Q); CQ2 assumption: quality_defined(Q); CQ3 assumption: action_fits_quality(A,E,Q); CQ4 assumption: reference_class(A,E,Q)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** retroductive/book; retroductive/independent-event-and-trait-link; retroductive/independent-wrong-event; retroductive/independent-wrong-quality; retroductive/independent-reference-class-event-scope; retroductive/positive; retroductive/adverse-fit-questioned; retroductive/two-agents

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 72. `evidence_hypothesis_verification`

**Entry:** 32.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The conditional H→observed A and observed-A premise support H in the printed defeasible verification schema. This is intentionally defeasible confirmation, not a deductively valid affirmation-of-the-consequent rule. CQ1–CQ2 test the two ordinary premises. CQ3 is the chosen sufficient_correlation(H,A) default, representing exclusion/adequacy of competing explanations when challenged.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 32.1, printed pp. 331, 332 (PDF pages +10). Separately cited original: Douglas Walton, Argumentation Schemes for Presumptive Reasoning. Mahwah, N.J.: Erlbaum, 1996, pp. 67-70.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: hypothesis_event(H,A); CQ2 premise: event_true(A); CQ3 assumption: sufficient_correlation(H,A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** evidence_hypothesis_verification/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 73. `evidence_hypothesis_falsification`

**Entry:** 32.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** H and event A remain in the same order as verification. An observed-false event supports explicit ¬H, matching 32.2. The book prints the shared alternative-explanation question after both variants. Keeping sufficient_correlation(H,A) in the same order is an explicit shared-CQ interpretation; the textual question is awkward for falsification and is not silently rewritten as an erratum.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 32.2, printed pp. 331, 332 (PDF pages +10). Separately cited original: Douglas Walton, Argumentation Schemes for Presumptive Reasoning. Mahwah, N.J.: Erlbaum, 1996, pp. 67-70.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: hypothesis_event(H,A); CQ2 premise: event_false(A); CQ3 assumption: sufficient_correlation(H,A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** evidence_hypothesis_falsification/book; evidence_hypothesis_falsification/positive; evidence_hypothesis_falsification/adverse-correlation-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 74. `positive_consequences`

**Entry:** 33.1. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The inherited action/outcome form supports should_be_performed(A), corresponding to 33.1. This review adds a book-supported likelihood-adequacy default; the previous missing-CQ status was not forced by lack of the 1995 source. CQ1 is now consequence_likelihood_adequate(A,G), meaning adequate assessment for the claimed argument strength, not necessarily strongly likely. CQ2 remains the supported consequence premise. CQ3 is a rebuttal for ¬should_be_performed(A). Remaining scope: The compendium-supported hook closes this operational gap. Preserve the original cited-source hold: the 1995 work was not independently read.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, pp. 155-156.

**Next concrete action:** Retain the compendium-supported adequacy hook without imposing a fixed numerical threshold; compare the original cited 1995 form before original-source approval.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 33.1, printed pp. 332, 333 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, pp. 155-156.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 assumption: consequence_likelihood_adequate(A,G); CQ2 premise: brings_about(A,G); CQ3 rebuttal: ¬should_be_performed(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** positive_consequences/book; positive_consequences/independent-likelihood-questioned; positive_consequences/independent-likelihood-backed; positive_consequences/independent-wrong-action-questioned; positive_consequences/independent-wrong-outcome-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 75. `negative_consequences`

**Entry:** 33.2. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The inherited action/outcome form supports ¬should_be_performed(A), corresponding to 33.2. The new likelihood default repairs the available-book omission without claiming to have seen the older cited book. CQ1 is the same action/outcome adequacy assessment used by 33.1. CQ2 remains brings_about(A,G), and CQ3 permits an opposing positive-consequence argument in one issue. No fixed probability threshold is introduced. Remaining scope: The original cited 1995 work remains unverified; this is an explicit adaptation of the inherited schema from the supplied compendium.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, pp. 155-156.

**Next concrete action:** Retain the compendium-supported adequacy hook and scoped inherited adaptation; compare the original cited 1995 form before original-source approval.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 33.2, printed pp. 332, 333 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, pp. 155-156.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 assumption: consequence_likelihood_adequate(A,G); CQ2 premise: brings_about(A,G); CQ3 rebuttal: should_be_performed(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_consequences/book; negative_consequences/independent-likelihood-questioned; negative_consequences/independent-likelihood-backed; negative_consequences/independent-wrong-action-questioned; negative_consequences/independent-wrong-outcome-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 76. `negative_consequences_reasoning`

**Entry:** 33.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Agent P and action A are in the correct roles. The selected bring-about reading concludes ¬should_do(P,A), so a practical recommendation for the same agent/action is its opponent in the negation issue. CQ1 is outcome_likely(A,B); CQ2 is the agent-conditioned consequence premise; CQ3 is a counterargument supporting should_do(P,A). The printed parenthetical alternative concerning not bringing A about is not claimed as a separately implemented variant.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 33.3, printed pp. 332, 333 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995. pp. 155-156. Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 123.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: outcome_likely(A,B); CQ2 premise: agent_bring_about(P,A,B); CQ3 rebuttal: should_do(P,A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_consequences_reasoning/book; negative_consequences_reasoning/positive; negative_consequences_reasoning/adverse-likelihood-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 77. `negative_consequences_prudential`

**Entry:** 33.4. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The contemplated omission of X and its bad consequence yield reconsideration and a prudential recommendation to do X. This conclusion is intentionally more specific than an unconditional practical ought. Likelihood and evidential CQ roles are retained. other_consequences(X) is a broad defeater for overlooked consequences, not a separately weighted opposite conclusion; this is an explicit conservative policy.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 33.4, printed pp. 332, 333 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995. pp. 155- 156. Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 123.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: strongly_occur(X,C); CQ2 premise: bad_consequence(X,C); CQ3 exception: other_consequences(X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** negative_consequences_prudential/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 78. `pragmatic_alternative`

**Entry:** 34. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The disjunctive practical premise and bad outcome yield the respondent-directed ought. Shared ought_to_bring_about has the same anonymous respondent reading as the threat/fear/danger family. The compendium prints no separate CQs here. The two ordinary premises must be supported; no extra question is asserted to be printed.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 34, printed pp. 333 (PDF pages +10). Separately cited original: Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 142.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** pragmatic_alternative/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 79. `threat`

**Entry:** 35. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The three source roles remain distinct: adverse consequence of the respondent action, the proponent being able to produce the consequence, and the expressed intention. The conclusion opposes bringing about the action. No separate CQs are printed for 35. The representation is a schema of threat argumentation and does not endorse the threat or establish a general moral prohibition.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 35, printed pp. 333 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995, p. 157.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** threat/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 80. `ad_baculum`

**Entry:** 35.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The disjunctive threat variant is distinct from entry 35: failure to do the requested action triggers the threatened bad outcome, and its conclusion recommends that action. No separate CQs are printed. The threatening conditional is one ordinary premise, not an inferred causal fact.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 35.1, printed pp. 333 (PDF pages +10). Separately cited original: Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 140.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** ad_baculum/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 81. `fear`

**Entry:** 36. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The omitted action leads to a bad event; prevention is recommended; the selected action is the only way to prevent it. The conclusion follows the source direction. No separate CQs are printed. The prevention recommendation and sole-means relation remain ordinary premises; they are not inferred solely from fear.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 36, printed pp. 333, 334 (PDF pages +10). Separately cited original: Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 22.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** fear/book; fear/against-danger

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 82. `danger`

**Entry:** 37. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** A dangerous consequence of doing X supports the on-balance recommendation not to do X. Explicit negation of ought_to_bring_about is the selected practical opposition policy. No separate CQs are printed. The badness/danger and consequence premises remain distinct.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 37, printed pp. 334 (PDF pages +10). Separately cited original: Douglas Walton, Scare Tactics. Dordrecht: Kluwer, 2000, p. 173.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** danger/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 83. `help`

**Entry:** 38. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Beneficiary A, potential helper B and action X are correctly separated. The conclusion should_do(B,X) is about the helper, not the beneficiary. CQ1 is whether the proposed concrete action really helps A. CQ2 and CQ3 restate the ability and acceptable-cost ordinary premises; duplicating them as defaults/exceptions was properly removed.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 38, printed pp. 334 (PDF pages +10). Separately cited original: Douglas Walton, Appeal to Pity. Albany: State University of New York Press, 1997, pp. 104,155.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: action_will_help(X,A); CQ2 premise: can_carry_out(B,X); CQ3 premise: not_too_costly(B,X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** help/book; help/positive; help/adverse-help-questioned; help/adverse-premise-missing; help/one-conclusion-with-distress

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 84. `distress`

**Entry:** 39. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The relief premise explicitly identifies the distressed person A, the relieving actor B and action X. The shared conclusion concerns B doing X. CQ1 is the distress premise; CQ2 is the relief relation; CQ3 is a challengeable possibility assumption; CQ4 is the actor/action side-effect exception.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 39, printed pp. 334, 335 (PDF pages +10). Separately cited original: Douglas Walton, Appeal to Pity. Albany: State University of New York Press, 1991, pp. 105,155.

**Declared-source CQ count:** 4. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: distress(A); CQ2 premise: relieve(B,X,A); CQ3 assumption: possible_bring_about(B,X); CQ4 exception: great_side_effects(B,X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** distress/book; distress/positive; distress/adverse-someone-elses-distress; distress/adverse-side-effects; distress/possibility-questioned-for-practical_inference-too

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 85. `commitment1`

**Entry:** 40a. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Evidence of commitment to X and the commitment link to Y support commitment to Y. It does not assert that Y itself is true. CQ1’s positive part restates the commitment premise; contrary evidence is an exception. CQ2’s linkage exception names the same person, X and Y.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 40a, printed pp. 335 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation. Cambridge: Cambridge University Press, 2006, pp. 117- 118. Douglas Walton, "The Sunk Costs Fallacy or Argument from Waste." Argumentation 16 (2002), p. 490.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: commitment_contrary_evidence(A,X); CQ2 exception: commitment_linkage_exception(A,X,Y)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** commitment1/book; commitment1/positive; commitment1/adverse-linkage-exception

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 86. `commitment2`

**Entry:** 40b. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The second version preserves its dialogue-specific conditional: a prior commitment to X yields a presumptive commitment when Y becomes an issue. The distinct conclusion was correctly retained. The commitment premise and contrary-evidence exception handle CQ1. CQ2 is the same person/X/Y linkage exception as version 1.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 40b, printed pp. 335 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation. Cambridge: Cambridge University Press, 2006, pp. 117- 118. Douglas Walton, "The Sunk Costs Fallacy or Argument from Waste." Argumentation 16 (2002), p. 490.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: commitment_contrary_evidence(A,X); CQ2 exception: commitment_linkage_exception(A,X,Y)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** commitment2/book; commitment2/positive; commitment2/adverse-contrary-evidence

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 87. `ethotic1`

**Entry:** 41. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The inherited positive-ethos record concludes S itself, whereas the compendium describes increased plausibility/acceptability. Defeasible support for S can be a declared evidential policy, but it is not a literal transcription of the graded conclusion. Moral character is an ordinary premise, relevance is a default, and claimed weight belongs to the issue’s evidential policy. The default does not provide a numerical weight calibration. Remaining scope: Retain the declared-strength and unavailable-1995-source qualifications. A stable claim of exact compendium coverage needs a distinct graded formulation or an explicitly accepted argument-weight policy.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 152.

**Next concrete action:** Compare the cited 1995 source and decide explicitly whether bare-proposition support is the accepted strength policy; otherwise add a graded-plausibility variant and examples.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 41, printed pp. 336 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 152.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: good_moral_character(P); CQ2 assumption: character_is_relevant(S); CQ3 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** ethotic1/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 88. `ethotic2`

**Entry:** 41. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** blocked as exact mapped book formulation. **Book-form decision:** block.

**Approval or hold reason:** The inherited negative-ethos record concludes explicit ¬S. Reduced plausibility or rejection of a speaker’s assertion does not in general establish the contradictory proposition. This is a substantive strength/polarity difference from the supplied compendium. Relevance and issue proof standards cannot by themselves make the inference from discrediting testimony to ¬S source-equivalent. The source asks whether the claimed weight is warranted. Remaining scope: Do not approve as exact entry 41. Retain as an explicitly qualified legacy variant outside a fully source-certified release, or replace/add a scoped reduced-plausibility formulation and cases. The cited 1995 work remains unverified.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 152.

**Next concrete action:** Keep outside exact entry 41 certification; add or replace with a scoped reduced-plausibility formulation and examples, or explicitly retain the legacy negation variant outside the certified subset; compare the 1995 source.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 41, printed pp. 336 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 152.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: bad_moral_character(P); CQ2 assumption: character_is_relevant(S); CQ3 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** ethotic2/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 89. `ad_hominem`

**Entry:** 42. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The source’s stated generic conclusion is strong rejection of the argument. The record identifies the attacked person, argument and conclusion; it does not infer the negation of that conclusion. CQ1 restates the character-attack premise (also discussed by the book at p. 374). CQ2 is argument-relevant character. CQ3 is implemented conservatively by an exception for stronger supporting evidence. This is a selected strong-rejection policy, not numerical downweighting. Follow-up: PR13 follow-up explicitly retains bad_character/1 separately from morally_bad/1 in entry47. The source does not justify automatic equivalence of generic character criticism and the more specific moral judgment.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 42, printed pp. 336 (PDF pages +10). Separately cited original: Douglas Walton , Fundamentals of Critical Argumentation. Cambridge: Cambridge University Press, 2006, p. 123.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: bad_character(A); CQ2 assumption: character_relevant(S); CQ3 exception: stronger_than_character(S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** ad_hominem/book; ad_hominem/positive; ad_hominem/adverse-another-persons-argument; ad_hominem/adverse-another-argument-of-the-same-person; ad_hominem/exception; ad_hominem/questioned-relevance; ad_hominem/questioning-the-removed-assumption-is-idle; ad_hominem/two-arguments; ad_hominem/one-conclusion-with-pragmatic_inconsistency

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; docs/reviews/peer13-second-followup.md.

### 90. `pragmatic_inconsistency`

**Entry:** 43. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The advocate, argument and proposition are explicit. The contrary action is represented with the proposition’s explicit negation, and the conclusion rejects the argument rather than declaring its conclusion false. CQ1 is strong advocacy/commitment; CQ2 is the contrary-action ordinary premise; CQ3 is relevance of the inconsistency to this person and argument.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 43, printed pp. 336, 337 (PDF pages +10). Separately cited original: Douglas Walton , Ad Hominem Arguments. Tuscaloosa: University of Alabama Press, 1998, pp. 218, 251.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: strong_advocate(A,S,X); CQ2 premise: opposite_action(A,S,¬X); CQ3 assumption: inconsistency_relevant(A,S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** pragmatic_inconsistency/book; pragmatic_inconsistency/positive; pragmatic_inconsistency/adverse-relevance-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 91. `inconsistent_commitment`

**Entry:** 44. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Evidence of commitment and opposed commitment supports inconsistency of the person’s commitments. The record no longer defeats itself merely because its second premise is present. CQ1 and CQ2 test the two evidence premises. CQ3 is the challengeable commitments_conflict(A,S) relation, so apparent opposition alone need not suffice once challenged.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 44, printed pp. 337 (PDF pages +10). Separately cited original: Douglas Walton, Fundamentals of Critical Argumentation. Cambridge: Cambridge University Press, 2006, pp. 120-121.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: initial_commitment(A,S); CQ2 premise: opposed_commitment(A,S); CQ3 assumption: commitments_conflict(A,S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** inconsistent_commitment/book; inconsistent_commitment/positive; inconsistent_commitment/the-evidence-of-premise-2-does-not-undercut; inconsistent_commitment/adverse-conflict-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 92. `circumstantial_ad_hominem`

**Entry:** 45. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The practical commitment correctly opposes the argument’s conclusion C, and the final claim is weaker decreased plausibility of argument S. However, credibility_questioned(A) and character_issue(A) omit S. CQ1 and the explained-inconsistency exception are correctly proposition/argument-scoped. CQ3 is still person-only; one argument’s dependence on credibility can be shared unintentionally with another argument by that person. CQ4 selects the weaker conclusion. Remaining scope: Before unrestricted use, bind character_issue and credibility_questioned to (A,S), with adverse wrong-argument controls and migration, or enforce one attacked argument per person/context.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Scope character_issue and credibility_questioned to person and attacked argument, or enforce one attacked argument per person/context; add wrong-argument controls and a migration entry.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 45, printed pp. 337, 338 (PDF pages +10). Separately cited original: Douglas Walton, Ad Hominem Arguments. Tuscaloosa: University of Alabama Press, 1998, pp. 255-256. Douglas Walton, Fundamentals of Critical Argumentation. Cambridge: Cambridge University Press, 2006, pp. 125-126.

**Declared-source CQ count:** 4. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: committed_opposite(A,C); CQ2 exception: inconsistency_explained(A,S); CQ3 assumption: character_issue(A); CQ4 standard: [no atom: application-level standard]

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** circumstantial_ad_hominem/book; circumstantial_ad_hominem/positive; circumstantial_ad_hominem/adverse-opposite-of-another-proposition; circumstantial_ad_hominem/adverse-inconsistency-explained

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 93. `bias`

**Entry:** 46. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The corrected conclusion is a comparative likelihood of incomplete consideration, not a categorical claim that the argument is false or that no contrary evidence was considered. CQ1 is relevance of the bias to this dialogue/conclusion; CQ2 is the ordinary bias premise. The same person is used in the general premise and application. Follow-up: PR13 follow-up restores the generic upstream biased/1 wording, which also fits the arguer of this entry. The logical formula and source interpretation are unchanged.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 46, printed pp. 338 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London : University of Alabama Press, 1995, p. 153.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: bias_relevant(A,C); CQ2 premise: biased(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** bias/book; bias/positive; bias/adverse-another-persons-premise; bias/adverse-dialogue-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; docs/reviews/peer13-second-followup.md.

### 94. `bias_ad_hominem`

**Entry:** 47. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The printed bad-bias/honest-participation sequence ends in reduced credibility for argument S. The record retains the printed morally-bad premise as an explicit premise rather than deriving it from any bias whatsoever. The evidence of bias and its inappropriate kind are ordinary premises. No extra printed CQs are invented, and normal permissible bias does not satisfy the inappropriate-bias premise by definition. Follow-up: PR13 follow-up explicitly retains morally_bad/1 separately from generic bad_character/1: moral badness is a more specific source judgment, not an automatic equivalence to every character criticism. An application may supply a justified bridge; the generic library does not assert one.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 47, printed pp. 338, 339 (PDF pages +10). Separately cited original: Douglas Walton, Ad Hominem Arguments. Tuscaloosa: University of Alabama Press, 1998, p.255.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: proponent_biased(A,S); CQ2 premise: biased_not_take_part(A,D,S)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** bias_ad_hominem/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; docs/reviews/peer13-second-followup.md.

### 95. `gradualism`

**Entry:** 48. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The initial proposition, an intervening implication sequence, and acceptance of the endpoint conditional yield acceptance of the final proposition. The sequence predicate is an aggregate witness with endpoints, not an algorithm that searches every chain. No CQs are printed. proper_sequence(P1,P4) is transparently marked n=0 as an additional executable-order assumption.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 48, printed pp. 339 (PDF pages +10). Separately cited original: Douglas Walton, Argumentation Schemes for Presumptive Reasoning. Mahwah, N.J.: Erlbaum, 1996, p. 96.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** assumption: proper_sequence(P1,P4)

**Example evidence:** gradualism/book; gradualism/two-sequences

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 96. `slippery_slope_base_case`

**Entry:** 49. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** This is one part of the inherited decomposition. It concludes negative consequences of an action, whereas entry 49’s full argument concludes that the initial action should not be brought about. The action/event link carries the chain-link CQs in this decomposition. No separate first-step presumption or normative bridge appears in this record. Remaining scope: Preserve as a useful legacy component with original-source hold. Do not count the pair alone as exact completed coverage of entry 49’s full scheme.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Slippery Slope Arguments, Vale Press, Newport News, 1999, pp. 93, 95.

**Next concrete action:** Compare the cited 1999 pages; retain the base component explicitly and provide the full entry 49 formulation if complete compendium coverage is claimed.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 49, printed pp. 339, 340 (PDF pages +10). Separately cited original: Douglas Walton, Slippery Slope Arguments, Vale Press, Newport News, 1999, pp. 93, 95.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: would_realize(A,E); CQ2 premise: would_realize(A,E); CQ3 premise: would_realize(A,E)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** slippery_slope_base_case/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 97. `slippery_slope_inductive_step`

**Entry:** 49. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The inherited recursive rule propagates horrible costs backwards along a causal link. It is a component of a slope construction, not the whole printed scheme. The link-strength, missing-step and weakest-link questions can be raised by contesting individual supplied links. The engine does not search for omitted intermediate events or calculate probabilities. Remaining scope: Retain inherited-component and unavailable-original-source qualifications, together with the entry-49 coverage limitation.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, Slippery Slope Arguments, Vale Press, Newport News, 1999, pp. 93, 95.

**Next concrete action:** Compare the cited 1999 pages; retain the inductive component explicitly and provide the full entry 49 formulation if complete compendium coverage is claimed.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 49, printed pp. 339, 340 (PDF pages +10). Separately cited original: Douglas Walton, Slippery Slope Arguments, Vale Press, Newport News, 1999, pp. 93, 95.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: causes(E1,E2); CQ2 premise: causes(E1,E2); CQ3 premise: causes(E1,E2)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** slippery_slope_inductive_step/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 98. `precedent_slippery_slope`

**Entry:** 50. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The source genuinely prints the initial case C0 in the intolerable-outcome premise and CQ3, even after describing a sequence ending at Cn. The existing C/R intolerability relation is faithful to that print. CQ1–CQ2 test the precedent-series premise; CQ3 tests the stated intolerable exception premise. No intuitive C0→Cn correction was made.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 50, printed pp. 340 (PDF pages +10).

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 premise: precedent_series(C,R,N); CQ2 premise: precedent_series(C,R,N); CQ3 premise: intolerable_exception(C,R)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** precedent_slippery_slope/book; precedent_slippery_slope/positive; precedent_slippery_slope/adverse-tolerable

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 99. `sorites_slippery_slope`

**Entry:** 51. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The clearly held starting property, successor conditional and reapplication sequence support the possibility-like endpoint claim may_have(A,P). Property P is now carried throughout the sequence. No separate CQs are printed. The whole chain is an ordinary aggregate premise whose witness must preserve property and endpoints.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 51, printed pp. 340, 341 (PDF pages +10). Separately cited original: Douglas Walton, "The Argument of the Beard." Informal Logic 18 (1996), p. 151.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** sorites_slippery_slope/book; sorites_slippery_slope/positive; sorites_slippery_slope/adverse-sequence-for-another-property; sorites_slippery_slope/adverse-sequence-between-other-individuals; sorites_slippery_slope/conditional-shared-with-verbal_slippery_slope

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 100. `verbal_slippery_slope`

**Entry:** 52. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The same property P is preserved across the original individual, neighbor and generalized endpoint. The conclusion says one cannot truly assert the endpoint has P; it is distinct from ordinary ¬has_property. No separate CQs are printed. Vague-neighbor and extreme-case claims are ordinary premises; an application must justify the intervening sequence and vagueness warrant.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 52, printed pp. 341 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995, p. 160.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** verbal_slippery_slope/book; verbal_slippery_slope/positive; verbal_slippery_slope/adverse-another-property-in-premise-2

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 101. `full_slippery_slope`

**Entry:** 53. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The current aggregate formula follows the printed 53: initial presumption, linked sequence and social climate accepting successive steps, concluding initial nonacceptability. The printed display does not separately state a bad endpoint; this review does not silently add one. CQ1–CQ2 concern strength and required proof standard, which remain application-level choices. CQ3 is the social-climate ordinary premise; duplicating it as an assumption was properly removed.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 53, printed pp. 341 (PDF pages +10). Separately cited original: Douglas Walton, Slippery Slope Arguments. Newport News, Va.: Vale Press, 1992, pp. 200, 203.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 standard: [no atom: application-level standard]; CQ2 standard: [no atom: application-level standard]; CQ3 premise: accept_all_steps(C0)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** full_slippery_slope/book; full_slippery_slope/positive; full_slippery_slope/questioning-the-removed-assumption-is-idle; full_slippery_slope/two-series

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 102. `rule_physical_world_1`

**Entry:** 54.1. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. The warrant explicitly requires referential coextension. The old whale/mammal fixture gave only class inclusion and nevertheless produced synonymy. The additional premise makes the warrant’s missing antecedent reviewable. The nonnecessary-truth rebuttal factor remains an undercutter scoped to W/D. The universal synonymy warrant is still explicitly supplied; the review does not assert it as an independently established linguistic theory.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 54.1, printed pp. 342 (PDF pages +10).

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: not_necessarily_true(W,D)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rule_physical_world_1/book; rule_physical_world_1/independent-coextension; rule_physical_world_1/independent-inclusion-is-insufficient; rule_physical_world_1/independent-other-pair-is-insufficient; rule_physical_world_1/positive; rule_physical_world_1/rebuttal-factor; rule_physical_world_1/two-claims

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 103. `rule_physical_world_2`

**Entry:** 54.2. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The institutional rulebook, warrant and changed-rule defeater match 54.2. The conclusion counts_as(D,W) omits institution I. The only printed rebuttal factor is correctly scoped to I/D/W. The rulebook-reporting warrant remains global as a universal warrant, not a claim that every particular rulebook is current. Remaining scope: Use one institution or institution-qualified ground terms. A generally multi-institution release should carry I in the counts-as statement and corresponding consumers.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use institution-qualified ground terms or carry institution I in the counts-as statement and its consumers; add a two-institution adverse example.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 54.2, printed pp. 342 (PDF pages +10).

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: rule_changed(I,D,W)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rule_physical_world_2/book; rule_physical_world_2/two-rules

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 104. `rule_mental_world`

**Entry:** 54.3. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The source language L is present in the speakers’ premise but omitted from means_same(W,D). The counterexample without_correcting(W) is also broader than the particular language/definition. The printed rebuttal concerns competent usage without correct understanding of the referent. The current W-only hook captures a single lexical context, not arbitrary languages or competing definitions. Remaining scope: Retain a single-language/definition context qualification or make language and definition explicit before unrestricted composition across languages.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Enforce one language/definition context or represent those identities explicitly; add a mixed-language/definition adverse case.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 54.3, printed pp. 342 (PDF pages +10). Separately cited original: Wayne Grennan, Informal Logic, Montreal: McGill-Queen's University Press, 1997, p. 196.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: without_correcting(W)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rule_mental_world/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 105. `established_rule`

**Entry:** 55.1. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The inherited schema remains rule-and-conclusion application, rather than 55.1’s actor/action form. This review adds the available-book conflicting/overriding-rule and exceptional/excusing-case hooks. CQ1 remains the asserted rule content; CQ2 and CQ3 now name R and its particular conclusion/application C. The valid(R) assumption remains explicitly marked as an additional n=0 policy. Remaining scope: The two previously unmodeled book questions are operationally closed. Original 1995 source verification and exact actor/action compendium-form coverage remain qualified.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 147.

**Next concrete action:** Retain the repaired scoped book questions; compare the original 1995 source and distinguish the inherited rule/conclusion form from the compendium actor/action form with explicit coverage claims.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 55.1, printed pp. 343 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 147.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** CQ1 premise: has_conclusion(R,C); CQ2 exception: conflicting_or_overriding_rule(R,C); CQ3 exception: rule_exception_or_excuse(R,C)

**Additional `n: 0` hooks:** assumption: valid(R)

**Example evidence:** established_rule/book; established_rule/independent-conflicting_or_overriding_rule-same-application; established_rule/independent-conflicting_or_overriding_rule-wrong-rule; established_rule/independent-conflicting_or_overriding_rule-wrong-conclusion; established_rule/independent-rule_exception_or_excuse-same-application; established_rule/independent-rule_exception_or_excuse-wrong-rule; established_rule/independent-rule_exception_or_excuse-wrong-conclusion

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 106. `rule`

**Entry:** 55.2. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Entry 55.2 prints alternative evaluation/conduct wording. The combined evaluation_justified(E,C) atom is read as the selected evaluative/conduct judgement described by the input, not as two independent propositions that are both automatically established. No separate CQs are printed. Both the rule conditional and particular case are ordinary premises.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 55.2, printed pp. 343 (PDF pages +10). Separately cited original: Peter Jan Schellens and Menno De Jong, "Argumentation Schemes in Persuasive Brochures. "Argumentation 18, 2004, p. 311.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** rule/book

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 107. `regulative_rule`

**Entry:** 55.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The source warrant supports obligatory action → must do. The printed parenthetical pairing reverses obligation/prohibition in the displayed conclusion; the current variant follows the warrant, with addressee scope explicit. The source rebuttal factor is an adequate excuse or overriding duty, correctly actor/action-scoped.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 55.3, printed pp. 343 (PDF pages +10). Separately cited original: Wayne Grennan, Informal Logic. Montreal: McGill-Queen's University Press, 1997, p. 169.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: excuse_or_overriding_duty(A,X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** regulative_rule/book; regulative_rule/obligation; regulative_rule/adverse-not-an-addressee; regulative_rule/excuse

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 108. `regulative_rule_prohibition`

**Entry:** 55.3. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The prohibition variant concludes must_not_do(A,X), a duty not to act. It does not merely deny a duty to act, and its addressee premise prevents introducing an unbound actor. The same actor/action excuse-or-overriding-duty factor is appropriate for the printed combined warrant.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 55.3, printed pp. 343 (PDF pages +10). Separately cited original: (55.3, read for a prohibition). Wayne Grennan, Informal Logic. Montreal: McGill-Queen's University Press, 1997, p. 169.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: excuse_or_overriding_duty(A,X)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** regulative_rule_prohibition/book; regulative_rule_prohibition/prohibition; regulative_rule_prohibition/adverse-obligation-is-no-prohibition

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 109. `exceptional`

**Entry:** 56. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired interpretation. The three questions are conditional: applicable recognition or evidence for an unrecognized case, with comparable-case support additionally required when borderline. They are not a flat OR. The aggregate exception_backed(C) remains a defeasible default. After challenge, the witness contract requires an explicit recognition status, borderline status and the applicable evidence/comparison route. Case C includes its governing rule/context.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 56, printed pp. 344 (PDF pages +10). Separately cited original: Douglas Walton, A Pragmatic Theory of Fallacy. Tuscaloosa and London: University of Alabama Press, 1995, p. 148.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: exception_backed(A); CQ2 assumption: exception_backed(A); CQ3 assumption: exception_backed(A)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** exceptional/book; exceptional/independent-recognized-clear-route; exceptional/independent-unrecognized-evidenced-route; exceptional/independent-recognized-borderline-route; exceptional/independent-unrecognized-borderline-route; exceptional/independent-recognition-does-not-bypass-borderline-comparison; exceptional/independent-comparison-does-not-bypass-nonapplicability-evidence; exceptional/independent-comparisons-for-another-case-do-not-suffice; exceptional/independent-unknown-status-does-not-default-to-clear; exceptional/positive; exceptional/adverse-rule-for-another-case; exceptional/adverse-exception-questioned; exceptional/questioned-and-backed

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md.

### 110. `from_precedent`

**Entry:** 57. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The legitimate counterexample has property F and explicit ¬G under the stated R/F/G general rule. The conclusion recognizes an exception for R/A and thereby calls for qualification; it does not silently rewrite the rule store. CQ1 is rule inapplicability to the case; CQ2 is legitimacy of the counterexample; CQ3 is an already recognized exception requiring no rule change. Each carries the relevant R/A scope.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 57, printed pp. 344 (PDF pages +10).

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: rule_inapplicable_to(R,A); CQ2 assumption: legitimate_case(A,R); CQ3 exception: recognized_exception_category(A,R)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** from_precedent/book; from_precedent/positive; from_precedent/adverse-the-case-has-g; from_precedent/recognized-category; from_precedent/legitimacy-questioned

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 111. `plea_for_excuse`

**Entry:** 58. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The special circumstances now identify the committed action. However the pleader remains implicit and excuse(E)/not_excuse(E) omit the governing rule, action and sanction application. CQ1–CQ2 are a recognized or specially justified excuse aggregate. CQ3 is a future-precedent problem for an unrecognized excuse. The aggregate must encode this conditional meaning for the particular plea. Remaining scope: Use case-specific excuse and sanction terms, or introduce explicit pleader/case scope before approving unrestricted multi-plea graphs.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use case-specific excuse/sanction terms or introduce explicit pleader/case arguments; add independent-plea controls before general multi-plea approval.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 58, printed pp. 345 (PDF pages +10). Separately cited original: Douglas Walton, Appeal to Pity. Albany: State University of New York Press, 1997, pp. 154-156.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 assumption: excuse(E); CQ2 assumption: excuse(E); CQ3 exception: not_excuse(E)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** plea_for_excuse/book; plea_for_excuse/positive; plea_for_excuse/adverse-excuse-for-another-action; plea_for_excuse/adverse-precedent-problem

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 112. `perception`

**Entry:** 59.1. **Origin:** book. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** qualified.

**Approval or hold reason:** The perceiver P and perceptible property I follow 59.1, with a prima-facie evidential warrant. The defeater not_indicator(I) omits P and perceptual context. The source undercutter asks about the circumstances in which the image is unreliable. A generic property name alone cannot distinguish one observer’s misleading circumstance from another observer’s reliable one. Remaining scope: Use a context-qualified I or separate graphs; for general multi-observer use, scope the unreliability defeater to observer/context and add independent-support controls like those now used for memory. Follow-up: PR13 follow-up visually compared pp345–346 and canonicalized reasonable_case/1 to memory's reasonable_believe/1 for the same case-qualified proposition. Five new cases establish independent supports, survival of either undefeated support, defeat of both, and protection against another case's memory. The not_indicator(I) observer/context qualification remains; no stage is promoted. The migrator rejects colliding old/canonical keyed entries rather than silently choosing their metadata or labels.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Use context-qualified perception claims or explicit observer/context arguments; add independent-observer controls for the unreliability defeater.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 59.1, printed pp. 345 (PDF pages +10). Separately cited original: John Pollock, Cognitive Carpentry. Cambridge, Mass.: MIT Press, 1995, p. 53.

**Declared-source CQ count:** 1. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: not_indicator(I)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** perception/book; perception/peer13-converges-with-memory; perception/peer13-unreliable-image-leaves-memory; memory/peer13-false-basis-leaves-perception; memory/peer13-both-supports-defeated; perception/peer13-other-case-memory-does-not-rescue

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; docs/reviews/peer13-second-followup.md; peer13-comparison-second.md.

### 113. `appearance`

**Entry:** 59.2. **Origin:** upstream. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** The visible classification premise, defeasible category conclusion, and both critical questions match pp. 345–346. The root reviewer independently checked the original 2006 journal PDF pp. 324–325; this reviewer separately checked the supplied compendium. misleading_appearance(O,C) undercuts misleading appearance in the proposed classification. more_justifiable_category(O,C) is a bounded existential/comparative defeater: evidence must show a better alternative category for this O/C claim, not merely some other category.

**Source-access obligation:** No outstanding source-access obligation within the declared scope: the original article was independently checked by root; retain the stated interpretation/policy qualifications.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 59.2, printed pp. 345, 346 (PDF pages +10). Separately cited original: Douglas Walton, ‘Argument from Appearance: A New Argumentation Scheme’, Logique et Analyse 195, 2006, pp. 324-325.

**Declared-source CQ count:** 2. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Root independently inspected the original journal PDF pp. 324-325; second-half reviewer independently inspected the compendium pp. 345-346. Two original-source CQs independently compared.

**Positive-numbered CQ mappings:** CQ1 exception: misleading_appearance(O,C); CQ2 exception: more_justifiable_category(O,C)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** appearance/book; appearance/positive; appearance/adverse-misleading; appearance/adverse-another-category

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md.

### 114. `memory`

**Entry:** 60. **Origin:** book. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** approve.

**Approval or hold reason:** Repaired. The belief-history undercutters identify recaller P and remembered proposition M. A false basis or absence of an original belief belonging to Ann does not defeat Bob’s independent recall of M. All three printed undercutters are modeled as exceptions. CQ1–CQ2 now have P/M scope; CQ3 already had it. The source’s unusual wording of CQ2 is preserved rather than silently reinterpreted. Follow-up: The PR13 follow-up makes perception's reasonable-belief conclusion identical to this record's existing conclusion for the same case-qualified proposition. This record's formula and scoped recaller defeaters remain unchanged. Five composed cases separately defeat each source, both sources and an unrelated case's memory.

**Source-access obligation:** The supplied compendium was directly inspected. No further book acquisition is needed for this declared book interpretation; follow the stated semantic action if held. Bibliographical predecessors are not silently claimed as read.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Walton, Reed and Macagno (2008), Argumentation Schemes, supplied complete PDF, entry 60, printed pp. 346 (PDF pages +10). Separately cited original: J.L. Pollock, "Defeasible Reasoning." Cognitive Science 11 (1987), pp. 481-518.

**Declared-source CQ count:** 3. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Not a separate original-source claim: approval concerns the directly inspected supplied compendium. Not applicable to this book-origin claim.

**Positive-numbered CQ mappings:** CQ1 exception: based_on_false_belief(P,M); CQ2 exception: not_originally_believed(P,M); CQ3 exception: recaller_doubts(P,M)

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** memory/book; memory/independent-false-basis-same-recaller; memory/independent-false-basis-other-recaller; memory/independent-support-survives; memory/independent-both-bases-defeated; memory/independent-original-belief-same-recaller; memory/independent-original-belief-other-recaller; memory/positive; memory/adverse-false-basis; memory/adverse-not-originally-believed; memory/adverse-recaller-doubts; perception/peer13-converges-with-memory; perception/peer13-unreliable-image-leaves-memory; memory/peer13-false-basis-leaves-perception; memory/peer13-both-supports-defeated; perception/peer13-other-case-memory-does-not-rescue

**Review reference:** book-second-audit/per-record-audit.json; independent-source-report.md; docs/source-contracts-second-half.md; book-first-audit/second-half-cross-review.md; docs/reviews/peer13-second-followup.md; peer13-comparison-second.md.

### 115. `abduction`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Its only declared source is the upstream code. The record's premises, conclusion and exception are compared against that code, and the explicit application reference case is available. The source prints no CQ list; more_coherent_explanation(T,S) remains transparently recorded as an additional n0 exception.

**Source-access obligation:** No outstanding access obligation: upstream code is the sole declared source and was compared statement by statement.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Carneades 4, examples/AGs/YAML/walton.yml (upstream, commit d19d543), where the record cites no source.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Verified against its only declared source, the upstream code record; no external theoretical source is cited. Known zero printed CQ list; implementation hooks are separately disclosed.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** exception: more_coherent_explanation(T,S)

**Example evidence:** abduction/book

**Review reference:** tools/vv/source_ledger.yml; tools/vv/internal/vvlib/source.go source/record-row; docs/reviews/own-source-review.md.

### 116. `credible_source`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Root independently compared Wyner (2016) p. 127: the three premises and conclusion match. The source prints no CQs. biased(W), dishonest(W), and other_credible_sources_disagree(S) are openly retained upstream application exceptions, not falsely attributed printed questions. Approval covers that declared adaptation. Follow-up: PR13 follow-up restores the inherited generic biased/1 wording rather than adding an Arguer restriction. The proposition/exception remains unchanged, and the reviewed fingerprint is refreshed for its language contract.

**Source-access obligation:** No outstanding source-access obligation within the declared scope: the original article was independently checked by root; retain the stated interpretation/policy qualifications.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Wyner, A., A functional perspective on argumentation schemes. Argument and Computation, vol. 7, pp. 113-133. 2016, p. 127.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Root independently checked Wyner (2016) p. 127; the printed premises/conclusion match. Three inherited exceptions are additional application policy. Known zero printed CQs on the source formulation; this does not erase the three n0 exceptions.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** exception: biased(W); exception: dishonest(W); exception: other_credible_sources_disagree(S)

**Example evidence:** credible_source/book

**Review reference:** tools/vv/source_ledger.yml; docs/reviews/own-source-review.md (Wyner 2016 p. 127); docs/reviews/peer13-second-followup.md.

### 117. `defeasible_modus_ponens`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `example-passed`. **Disposition:** approved within declared scope. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Its only declared source is the upstream code. The implication premise and antecedent produce the consequent under the inherited defeasible scheme; statement-by-statement provenance comparison and the explicit reference case support this narrow source claim. The code prints no CQ list.

**Source-access obligation:** No outstanding access obligation: upstream code is the sole declared source and was compared statement by statement.

**Next concrete action:** Retain the recorded interpretation and source provenance; include only with a passing complete commit-matched engine/library gate. Re-review semantic or burden changes rather than regenerating expectations.

**Source reference:** Carneades 4, examples/AGs/YAML/walton.yml (upstream, commit d19d543), where the record cites no source.

**Declared-source CQ count:** 0. Known zero printed questions in the declared source formulation; additional policy hooks remain separate.

**Own-original-source status:** Verified against its only declared source, the upstream code record; no external theoretical source is cited. Known zero printed CQ list; implementation hooks are separately disclosed.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** defeasible_modus_ponens/book

**Review reference:** tools/vv/source_ledger.yml; tools/vv/internal/vvlib/source.go source/record-row; docs/reviews/own-source-review.md.

### 118. `exeption_to_rule`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `mapped`. **Disposition:** mapping-only; source interpretation unreviewed. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Mapping-only. The inherited code record was named precedent and is now named exeption_to_rule. Its cited 1995 p. 148 source and the basis for this naming/interpretation have not been independently adjudicated. CQ count is unknown; a generic case1/rule1 fixture does not supply source approval.

**Source-access obligation:** Obtain and independently compare the separately cited source: Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 148 Also retrieve any preserved 2019 correspondence if it is used to justify renaming upstream precedent as an exception.

**Next concrete action:** Obtain the cited 1995 p. 148 and any preserved 2019 correspondence underlying the rename; adjudicate whether this is precedent or exception, inventory actual source CQs, and replace the generic reference example with a sourced interpretation.

**Source reference:** Douglas Walton, A Pragmatic Theory of Fallacy, The University of Alabama Press, Tuscaloosa and London, 1995, p. 148

**Declared-source CQ count:** unknown. Unknown: source not inventoried; not zero.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** exeption_to_rule/book (generic fixture; not source approval)

**Review reference:** tools/vv/source_ledger.yml; upstream precedent record; current generic fixture.

### 119. `practical_reasoning1`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `transcription-checked`. **Disposition:** qualified; interpretation approval withheld. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Original-source comparison is complete, but the record is a flat AS1 abstraction with inherited defaults and unconditional exceptions. The paper has 17 questions. Eight mappings remain incomplete (5,6,7,11,12,14,15,17); CQ8-9 are partial policy treatments; CQ10 supplies positive support, not an automatic undercutter. The source-corrected farmer Arg8 and Arg6 reference cases do not establish full AS2/AATS fidelity.

**Source-access obligation:** No outstanding access obligation: the complete supplied 20-page paper was read. Remaining work is model/CQ interpretation and implementation, not obtaining the paper.

**Next concrete action:** Preserve the legacy flat AS1 record; implement a separate explicitly scoped AATS variant/adapter or documented model-validation contract covering CQ5,6,7,11,12,14,15,17, plus agent/transition/value and CQ8-9 policy differences; retain CQ10 as support.

**Source reference:** Atkinson, K., and Bench-Capon, T. J. M. Practical reasoning as presumptive argumentation using action based alternating transition systems. Artificial Intelligence 171, 10-15 (2007), 855–874.

**Declared-source CQ count:** 17. Known printed-question count; complete numbering does not imply complete operational or semantic coverage.

**Own-original-source status:** Complete supplied 2007 article independently read; source comparison complete, full interpretation approval withheld. CQ1-16 on p. 858 and CQ17 on p. 859; formal treatments pp. 861-862.

**Positive-numbered CQ mappings:** CQ1 premise: current_circumstances(S1); CQ2 premise: would_bring_about(A,S1,S2); CQ3 premise: would_be_realized(G,S2); CQ4 premise: would_promote_value(G,V); CQ5 alternative: [incomplete operational mapping]; CQ6 alternative: [incomplete operational mapping]; CQ7 alternative: [incomplete operational mapping]; CQ8 exception: side_effects(A,S1); CQ9 exception: side_effects(A,S1); CQ10 support: should_be_performed(A); CQ11 alternative: [incomplete operational mapping]; CQ12 assumption: [incomplete operational mapping]; CQ13 assumption: possible(A); CQ14 assumption: [incomplete operational mapping]; CQ15 assumption: [incomplete operational mapping]; CQ16 assumption: legitimate_value(V); CQ17 assumption: [incomplete operational mapping]

**Additional `n: 0` hooks:** assumption: worthy_goal(G); exception: bring_about_more_effectively(S2,A); exception: realize_more_effectively(G,A); exception: promote_more_effectively(V,A)

**Example evidence:** practical_reasoning1/book; tools/vv/cases/atkinson_2007.yml (12 dedicated cases)

**Review reference:** docs/reviews/atkinson-2007.md; docs/reviews/atkinson-2007-decisions.json.

### 120. `practical_reasoning2`

**Entry:** upstream-only. **Origin:** upstream. **Current stage:** `mapped`. **Disposition:** mapping-only; source interpretation unreviewed. **Book-form decision:** not a compendium mapping.

**Approval or hold reason:** Mapping-only. The 2015 source has not been inspected, its actual CQ count has not been inventoried, and the generic Hal/insulin fixture does not establish that source's intended interpretation. Source correctness and burdens remain unapproved.

**Source-access obligation:** Obtain and independently compare the separately cited source: Walton, Douglas (2015). Goal-Based Reasoning for Argumentation. Cambridge University Press.

**Next concrete action:** Obtain Walton (2015), locate and compare the actual instrumental schema and all its questions, inventory the printed CQ count, then review burdens and a genuinely source-grounded reference example.

**Source reference:** Walton, Douglas (2015). Goal-Based Reasoning for Argumentation. Cambridge University Press.

**Declared-source CQ count:** unknown. Unknown: source not inventoried; not zero.

**Own-original-source status:** Own separately cited original work has not been independently certified in this review; any supplied-compendium comparison is recorded separately. Unknown for the own original until it is inspected; do not replace this with the compendium count or zero.

**Positive-numbered CQ mappings:** No printed positive-numbered list in this declared formulation, or source inventory still unknown as stated above.

**Additional `n: 0` hooks:** None recorded.

**Example evidence:** practical_reasoning2/book (generic fixture; not source approval)

**Review reference:** tools/vv/source_ledger.yml; inherited code and current generic fixture.

