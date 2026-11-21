# Independent review of book entries 1–27

**Review date:** 2026-11-21. **Baseline repository:** `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`.

## Decision

The book-form interpretations in this part of the library were not all ready for unconditional approval. The independent review identified concrete semantic discrepancies that ordinary successful engine runs did not reveal. I reproduced seven source-expectation failures on the unchanged baseline, repaired the supported discrepancies in an isolated copy, and supplied eleven new positive/adverse reference cases. The seven reproductions concern six topics: generalization identity, model conditional, degree-qualified property, personal practical reasoning (two records), negative-practical choice and the missing definition question.

After the repairs, **35 of the 63 assigned book-form readings are approved and 28 remain qualified**. Four of the 35 approved book forms are inherited records whose independently cited originals are a separate verification obligation. **31 book-origin records are candidates to advance to `example-passed` once the integrated suite passes.** The patch does not silently advance source-ledger stages. The complete per-record judgments appear below and in `adjudication.json` and `adjudication.csv`.

“Qualified” is not a euphemism for a failing engine. It means that a restricted source branch, contextual assumption, dialogue burden decision or inherited-source difference must be explicit before the record is described as an unqualified implementation of the source. Several qualifications can be resolved through precise contractual scope and a targeted example; others require another schema or additional parameters. Passing existing `/book` examples cannot decide those questions automatically.

## Scope and evidence

This review covers all records mapped to book entries 1 through 27, excluding inherited `practical_reasoning1`, whose Atkinson and Bench-Capon source is assigned to another reviewer. There are 63 records: 55 book-origin and eight upstream-origin records. I compared the actual premises, conclusions, assumptions, exceptions, language templates, source-ledger classifications and existing `/book` reference cases against the supplied book, *Argumentation Schemes* by Walton, Reed and Macagno (2008), printed pp. 309–328. Narrative pp. 95–96, 113 and the book's Carneades discussion pp. 374–375 were consulted where they explain a reading.

The supplied PDF's printed page p is PDF page p+10. Its SHA-256 is `d010348014e8b80fa561b68e5fac732cce0b3905fb0b1d522dd6fd2578a69634`. Decisive passages were visually inspected on printed pp. 96, 314, 315, 316, 319, 323, 324 and 327; the relevant rendered pages are in `qa/`. The full forms and source-ledger rows used for comparison are preserved in `inventory.json`. No claim is made that inaccessible separately cited originals were read, that every conceivable argument graph was tested, or that source conformity implies the philosophical validity of every source pattern.

## What was found and repaired

### A general conclusion must have a general identity

The example scheme's source concludes a general relationship from property F to property G. The baseline encoded its conclusion as `both_properties(example,F,G)`. Two witnesses therefore supported two different propositions despite their statements expressing the same generalization. The original reference case repeated that representation and passed.

The correction uses `usually_properties(F,G)`, already the rule premise of the illustration scheme. Individual examples remain distinct supporting arguments. A new case checks two examples supporting one generalized proposition and then feeding an illustration. This is an appropriate further use of the shared-vocabulary policy: statement identity reflects the proposition, while supporting argument parameters preserve provenance. The book's strength question remains a qualification: an isolated exception is not automatically a refutation of a defeasible generalization. (Book p. 314; PDF p. 324.)

### The model scheme needs its conditional warrant

The source's third premise links the named model, action and imitators conditionally. The baseline required only `worth_imitated(action,imitators)`, which already asserts the consequent and drops the model from the warrant. The ledger itself acknowledged that truncation while calling the record a transcription.

The correction introduces `imitation_warrant(model,action,imitators)` with the complete conditional meaning. A positive case supplies the source-faithful warrant; an adverse control proves that another model's warrant does not instantiate the first model's argument. Migration is manual because the old statement alone does not reveal its model. (Book p. 315; PDF p. 325.)

### Being less beautiful does not mean being not beautiful

The part-in-whole source states a property conclusion with a parenthetical lower-degree qualification. The baseline turned that qualification into an exception. Adding the fact that a rose has a lower degree of beauty than its whole therefore made the argument for the rose's having the beauty property evaluate out.

The correction removes that undercutting role and retains the positive property branch. The degree statement remains available for graphs that assert it. This intentionally adds one unused-template warning; it does not invent a new executable degree calculus. The species/part alternatives and degree reasoning remain declared scope qualifications. (Book p. 316; PDF p. 326.)

### Practical reasoning must keep the same deliberating person

The baseline necessary and sufficient practical schemes named a person in `goal_of_agent` and `should_do`, but their selection, feasibility and preference premises still used unscoped first-person statements. In a two-person graph, Ann's chosen means and deliberation facts gave Bob the same recommendation merely because Bob shared the goal. The engine supplied Bob's feasibility default; it did not need evidence that Bob made Ann's choice.

The correction names the agent on all personal deliberation premises, the personal acceptability assumption and the alternative-means exception in these two records. The new cases check that Ann receives the supported recommendation, Bob does not inherit it, and Bob's alternative does not defeat Ann's argument. These are source-driven role controls, not just generic grounding tests. (Book pp. 95–96, 323–324; PDF pp. 105–106, 333–334.)

The sufficient-condition template also retained “at least one” after an earlier fix changed “necessary” to “sufficient.” The narrative schema on p. 96 says each member is sufficient. The repaired wording follows that narrative rather than the apparent copied necessary-schema line on p. 323.

### Negative practical reasoning must select a polarity

When the loss from falsely assuming A true is worse than the converse error, the source selects action on the presumption that A is false. The baseline conclusion template said that A is “not true or true,” so it encoded no choice at all.

The corrected negative branch concludes `act_on_presumption_false(A)`. An adverse case reverses the seriousness ordering and verifies that this negative branch no longer fires. A future true-presumption branch needs a separate explicit representation; it is not claimed to be implemented by this patch. The change is intentionally a manual semantic migration, not a blanket identifier substitution. (Book p. 327; PDF p. 337.)

### The definition question is available in the supplied book

The author update deferred the merely stipulative or biased definition question because the inherited record's original reference was inaccessible. The supplied compendium itself prints that question. It is therefore possible to implement a clearly attributed book-supported interpretation while retaining the original-source verification hold.

The correction adds `doubtful_definition(D,G)` as an exception, with a relevant-definition defeat case and a control for an unrelated definition. The record remains upstream-origin; this addition does not pretend that the cited 2006 book was inspected. The source also has inconsistent bound variables in its classification premise; the interpretation retains the coherent definition-to-class reading rather than concluding that all objects have the class. (Book p. 319; PDF p. 329.)

Finally, the vagueness record's exception wording is corrected from certainty to precision. Those are different demands, and the printed question concerns precision. Its existing identifier is retained to avoid an unnecessary extra rename. (Book pp. 319–320.)

## Remaining decisions and release consequences

**Independent approval must remain substantive.** A test can prove that an encoded scheme fires, is undercut or round-trips as specified. It cannot prove the encoded conditional, actor, scope, quantifier or chosen polarity matches the book. The baseline failures above are direct evidence of that distinction. Advancing every row solely because its current tests pass would erase the purpose of the interpretation stage.

**Keep static evaluation distinct from full dialogue behavior.** On p. 96, asking any of five practical questions can temporarily defeat an argument until an adequate answer is given. Several current mappings instead require affirmative contrary evidence. That can be an explicit static-evaluation interpretation, but it is not automatically the complete dialogue policy in the source. The plan needs a decision that states which behavior the released application promises and source-derived cases for questioning, answering, withdrawal and independent support.

**Keep proof standards distinct from burden allocation.** A proof standard compares evidential sufficiency. It does not alone designate the ultimate probandum or who must establish it. The ignorance-family ledger currently labels both questions “standard”; the application contract must retain their distinction.

**Name the scope of represented alternatives.** Three-part positive division is not arbitrary-n division with every negative variant. Good/easy/great branches do not implement bad/difficult/small branches. An opaque action-or-policy predicate does not by itself supply independently typed action and policy conclusions. The record-level ledger identifies each such case so it can be scoped or expanded deliberately.

**Require consistent contextual identity when sharing vocabulary.** `has_property(entity,property)` is appropriate where the entity and property denote the same things. Bare proposition conclusions correctly separate the proposition from the strength of its support. `should_do(agent,act)` is appropriate for named-agent recommendations only when the premises bind that same agent. Implicit speaker, audience, case and value-context records need a declared current-context convention or explicit context parameters when combined across cases. Likewise, ordinary deontic logic distinguishes not-required(A) from required(not-A); a two-position practical recommendation convention should be stated as such.

**Separate a book match from an original-source match.** Four inherited records in this assignment now have approved book-form readings. That does not verify all details of their own cited originals. The provenance ledger should record the basis of the approval instead of using one green label to imply multiple sources were inspected.

The next useful round is targeted implementation and acceptance of these named scope/burden decisions, followed by the integrated clean suite. A further open-ended cycle of reviewers restating the plan is unnecessary. Stable release approval should wait until the full library's remaining qualified and inherited-source rows are explicitly resolved; this limited audit alone cannot certify all 120 records.

## Files and reproducibility

- `first-half-source-repairs.patch`: isolated source corrections, reference cases, category regeneration, migration rows and review documentation.
- `adjudication.json`: every assigned row, before/after logical forms, original ledger claims, all questions, source-page locators, existing example reviewed and new examples.
- `adjudication.csv`: compact comparison table with every per-record judgment.
- `adjudication-summary.json`: exact counts and the 31 book-origin promotion candidates.
- `baseline/`: seven input graphs, complete YAML outputs, stderr and independently stated expected-result failures.
- `final-new-cases-source.log` and `final-new-cases-source-results.jsonl`: final bounded suite execution; the broader diagnostic run is retained separately as `cases-source.log`.
- `validator-final.log`, `validator-tests.log`, `migration-tests.log`: validator and migration evidence.
- `qa/`: source page renders used for decisive visual comparisons.

The source archive, engine and public release are not changed by this patch. The parent review owns integrated engine verification, the complete expected-check inventory and the release decision. This review provides complete source adjudication for its assigned scope.

## Verification of the source repair patch

The final targeted run uses Go 1.27.4 and the baseline engine binary against the patched library. It passes **630 of 630 checks**, covering all new source reference cases through the Go API and CLI round trip, plus all source/migration checks. The log is `final-new-cases-source.log`, with every result in `final-new-cases-source-results.jsonl`.

Before adding the final definition-CQ cases, the combined full `TestCases|TestSource` run passed 1,225 of 1,225 checks. The later 1,231-check run passed 1,229; its only failures were the two new fixtures listing definition-scheme parameters in the wrong order. Their argument expectations were corrected to the scheme's existing `(object,class,definition)` parameter order, and both pass in the final 630-check run. The earlier diagnostic run is retained as `cases-source.log`; it is not the final acceptance result. The parent review owns the complete integrated engine/library run and expected-check inventory.

The validator reports zero errors and eleven warnings. The extra warning is the intentionally retained lower-degree template that no longer undercuts property possession. The validator's five mutation/control tests and migration tool's seven tests pass. Generated category files match the master; the patch passes `git diff --check`.

## Complete per-record adjudication

Each section states the source reading, question/burden interpretation, remaining action and stage recommendation. An approval is bounded by its stated source form; no finite example suite proves correctness for every input.

### Entry 1 — `position_to_know`

**Decision after patch:** approve. **Origin:** upstream. **Source pages:** 309. **Stage recommendation:** `transcription-checked`.

Domain D links W's access and proposition S; source and assertion roles are coherent. Bare S preserves truth-claim identity.

CQ1 and CQ3 are premises, honesty is an opponent-supported exception. This is a reasonable static reading of the three book questions.

None for book form; its own cited Legal Argumentation and Evidence p.46 remains outside this book-only approval.

**Example reviewed:** `position_to_know/book`; expected accepted conclusion(s): `red_light_run`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 2 — `expert_opinion`

**Decision after patch:** approve. **Origin:** upstream. **Source pages:** 310, 374, 375. **Stage recommendation:** `transcription-checked`.

Expert W, proposition S and domain D stay linked. Bare S is the correct object proposition, without a second plausibility fact.

All six printed questions are represented; credibility as a default is additionally corroborated by the book's own Carneades model, pp.374-375.

Its separately cited original source remains subject to the inherited-source worklist.

**Example reviewed:** `expert_opinion/book`; expected accepted conclusion(s): `liable(acme)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 3 — `witness_testimony`

**Decision after patch:** qualified. **Origin:** upstream. **Source pages:** 310. **Stage recommendation:** `transcription-checked`.

W's sincerity is represented by believes(W,S) plus assertion; this is a defensible interpretation, not an exact quotation of knowing truth. A domain fact supplies access.

Internal consistency defaults; conflicting facts, other witnesses, bias and implausibility are exceptions. Global biased(W) needs case-specific use, because the question concerns the account.

Make explicit that W denotes this witness in the case and that bias is relevant to this account; own-source verification is separate.

**Example reviewed:** `witness_testimony/book`; expected accepted conclusion(s): `red_light_run`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 4 — `popular_opinion`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 311. **Stage recommendation:** `example-passed`.

The output remains a reason in favor of A, not an assertion that A is true. General acceptance and the acceptance-to-reason warrant are distinct.

CQ1 is challenged via the general-acceptance premise; CQ2 is an exception. Both printed questions are accounted for.

No semantic blocker for the stated static interpretation.

**Example reviewed:** `popular_opinion/book`; expected accepted conclusion(s): `reason_in_favor(earth_round)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.1 — `pop_scheme`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 311. **Stage recommendation:** `example-passed`.

The assertive variant concludes A itself. A may be a signed proposition, so accepting a negative proposition does not require an extra truth-wrapper.

No printed CQs in this numbered variant; underlying group acceptance is an explicit premise.

Approve the assertive variant separately from the advice variant.

**Example reviewed:** `pop_scheme/book`; expected accepted conclusion(s): `earth_round`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.1 — `pop_scheme_acceptance`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 311. **Stage recommendation:** `example-passed`.

The advisory variant concludes that the implicit respondent should accept A. It no longer asserts both truth and acceptance in one inference.

No printed CQs; this is the expressly separated acceptance reading.

Scope is accepting the proposition supplied as A; rejecting an A is not automatically interchangeable with accepting its negation in every dialogue.

**Example reviewed:** `pop_scheme_acceptance/book`; expected accepted conclusion(s): `should_accept(earth_round)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.2 — `position_to_know_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 311. **Stage recommendation:** `example-passed`.

G is the same group in consensus and epistemic-position premises; A is the same conclusion proposition. Replacing plausibly_true(A) with A is appropriate for a defeasible argument.

No printed CQs for this variant; premise support carries the challenge path.

No semantic blocker for this source form.

**Example reviewed:** `position_to_know_ad_populum/book`; expected accepted conclusion(s): `safe_beach`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.3 — `expert_opinion_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `example-passed`.

The expert group and consensus proposition are preserved. The source itself does not bind the proposition to a separate domain variable, so this review does not invent one.

No printed CQs in this variant; a stronger domain-sensitive version could be added under a distinct interpretation.

No semantic blocker for the printed form.

**Example reviewed:** `expert_opinion_ad_populum/book`; expected accepted conclusion(s): `vaccines_safe`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.4 — `deliberation_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `example-passed`.

The deliberating group is also the group accepting A; the output is A, with defeasibility represented by argument evaluation.

No printed CQs in this variant.

No semantic blocker for this source form.

**Example reviewed:** `deliberation_ad_populum/book`; expected accepted conclusion(s): `guilty_verdict`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.5 — `moral_justification_ad_populum`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `transcription-checked`.

Good-group policy acceptance and the respondent's goal support should_accept(P); the positive good-group branch is represented.

No printed CQs; respondent identity remains implicit, as in the text.

Document one implicit respondent per graph or encode recipient identity when combining different audiences.

**Example reviewed:** `moral_justification_ad_populum/book`; expected accepted conclusion(s): `should_accept(recycling)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 4.6 — `moral_justification_ad_populum_excuse`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `transcription-checked`.

The record retains action and policy alternatives in composite predicates. It does not separately infer morally_acceptable(action) or acceptable(policy).

No printed CQs. The speaker is implicit in group_member/1 and the conclusion.

Declare this as an opaque combined action/policy reading; split typed action and policy variants before claiming independent interoperability of each branch.

**Example reviewed:** `moral_justification_ad_populum_excuse/book`; expected accepted conclusion(s): `morally_acceptable_action(speeding,speeding_ok)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 4.7 — `snob_appeal_ad_populum`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `transcription-checked`.

Group, prestige goal and membership warrant are present, but acceptance, property acquisition and object purchase are compressed. Only should_accept(A) is concluded.

No printed CQs. P and O occur in the membership warrant but do not select a conclusion branch.

The Rolex reference case treats an object as A in accept(A); replace with a clearly typed acceptance example, and add property/purchase variants if those alternatives are claimed as covered.

**Example reviewed:** `snob_appeal_ad_populum/book`; expected accepted conclusion(s): `should_accept(rolex)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 4.8 — `appeal_to_vanity_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 312. **Stage recommendation:** `example-passed`.

The respondent's action X, not admired group G or proposition A, is the conclusion object. Separating the second-person ought from a first-person practical goal is correct.

No printed CQs. The source uses one implicit respondent.

Approve for one current respondent; do not silently treat the implicit audience as every named agent.

**Example reviewed:** `appeal_to_vanity_ad_populum/book`; expected accepted conclusion(s): `ought_to_bring_about(buy_brand_x)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.9 — `rhetoric_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 313. **Stage recommendation:** `example-passed`.

Consensus, valued membership and contrasting out-group acceptance refer to the same G and A. The conclusion is acceptance by the implicit respondent.

No printed CQs.

One current respondent is part of the source reading.

**Example reviewed:** `rhetoric_ad_populum/book`; expected accepted conclusion(s): `should_accept(leadership_honest)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.10 — `common_folk_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 313. **Stage recommendation:** `example-passed`.

Speaker S is explicit in background(S) and being_credible(S); removing the old zero-argument forms prevents unrelated speakers from sharing these facts.

No printed CQs. The audience remains the current audience of the case.

Approve the speaker-bound, single-audience interpretation.

**Example reviewed:** `common_folk_ad_populum/book`; expected accepted conclusion(s): `being_credible(senator)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 4.11 — `common_folks_group_ad_populum`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 313. **Stage recommendation:** `example-passed`.

The speaker belongs to the same background group as the respondent; accept_saying(S) retains the speaker identity.

No printed CQs. Group G is consistently bound.

Approve for the current respondent; the predicate is an endorsement of this speaker, not all of that person's future utterances across cases.

**Example reviewed:** `common_folks_group_ad_populum/book`; expected accepted conclusion(s): `accept_saying(senator)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 5 — `popular_practice`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 314. **Stage recommendation:** `transcription-checked`.

The popular-practice premise and current-case conclusion have the right roles. The minor template is the consequent of an anaphoric warrant, rather than a self-contained conditional.

CQ1 is a premise challenge; CQ2 is represented by justified grounds as a default assumption, a defensible interpretation.

Clarify the warrant template and state that each graph has one current case; cases with different circumstances require distinct contextual practice terms.

**Example reviewed:** `popular_practice/book`; expected accepted conclusion(s): `acceptable_in_case(tipping)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 6.1 — `from_example`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 314. **Stage recommendation:** `transcription-checked`.

REPAIRED: the generalization now depends on F and G only, using usually_properties(F,G). The example a remains in argument parameters, allowing convergent support and feeding illustration.

CQ1 premise, CQ2/CQ3 defaults, CQ5 exception. CQ4 is a supported challenge to the generalization in its negation issue; one exceptional case does not by itself refute a defeasible generalization.

Keep this CQ4 policy explicit. Generic grounded examples are insufficient to calibrate how strong a generalization is; a policy-level strength example remains desirable.

**Example reviewed:** `from_example/book`; expected accepted conclusion(s): `usually_properties(bird,flies)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

**New independent cases:** `from_example/independent-two-examples-one-generalization`.

### Entry 6.2 — `illustration`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 315. **Stage recommendation:** `example-passed`.

The rule premise and validity conclusion depend on F and G rather than on the illustrating individual. The particular case supplies an instance of both properties.

No printed CQs. Quantification over x in the rule is represented as an opaque rule proposition, not an inference that enumerates all x.

Approve this source-form interpretation; the new two-example case also verifies composition from example to illustration.

**Example reviewed:** `illustration/book`; expected accepted conclusion(s): `valid_rule(bird,flies)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 6.3 — `from_model`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 315. **Stage recommendation:** `example-passed`.

REPAIRED: the third premise is now imitation_warrant(model,action,imitators), preserving the complete conditional and its model. The imitators are the recipients of should_do.

No printed CQs. The new adverse case prevents a warrant about one model from licensing a different model's argument.

No remaining book-form blocker after the patch and positive/adverse cases.

**Example reviewed:** `from_model/book`; expected accepted conclusion(s): `should_do(students,fasting)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

**New independent cases:** `from_model/independent-conditional-warrant`.

### Entry 6.4 — `from_antimodel`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 315. **Stage recommendation:** `transcription-checked`.

The model, act and would-be imitators are consistently bound. The implementation concludes negated should_do, while the text says should not do.

No printed CQs. Under a two-position recommendation policy these can be treated as opposing recommendations, but ordinary deontic logic distinguishes not obligatory from obligatory not.

Document the binary recommendation convention or introduce an explicit refrain-from-action predicate before claiming full deontic equivalence.

**Example reviewed:** `from_antimodel/book`; expected accepted conclusion(s): `¬should_do(citizens,lying)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 7 — `analogy`

**Decision after patch:** qualified. **Origin:** upstream. **Source pages:** 315. **Stage recommendation:** `transcription-checked`.

The source and current cases are distinguished by an implicit current-case convention. The conclusion is A in that current case.

CQ1 relevant differences and CQ2 base premise are present. CQ3 in the book asks about another similar contrary case, whereas more_on_point imposes a stronger relevance condition.

Own-source form may justify the inherited variant, but this is not yet a complete equivalence to the book CQ3; keep that scope difference explicit.

**Example reviewed:** `analogy/book`; expected accepted conclusion(s): `liable`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 8.1 — `practical_analogy_positive`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 316. **Stage recommendation:** `example-passed`.

The action X is preserved while the situation changes from A1 to similar A2. The conclusion remains a normative assessment in A2.

No printed CQs in these practical analogy variants.

No semantic blocker for the positive source form.

**Example reviewed:** `practical_analogy_positive/book`; expected accepted conclusion(s): `carry_out_right(case2,mediate)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 8.2 — `practical_analogy_negative`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 316. **Stage recommendation:** `example-passed`.

The negative assessment of doing X transfers from A1 to similar A2; it does not change the action or assert the opposite descriptive proposition.

No printed CQs.

No semantic blocker for the negative source form.

**Example reviewed:** `practical_analogy_negative/book`; expected accepted conclusion(s): `carry_out_wrong(case2,strike)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 9.1 — `composition_generic`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 316, 113. **Stage recommendation:** `example-passed`.

All parts have the property; the conclusion attributes it to the whole. Sharing has_property with other schemes is correct when object and property identifiers denote the same things.

The compositional-heredity default follows the direction established in the narrative; the compendium parenthesis reverses the direction and is appropriately treated as a supported correction.

This is defeasible composition, not a universally valid transfer of every property. The hereditary default is question-sensitive.

**Example reviewed:** `composition_generic/book`; expected accepted conclusion(s): `has_property(car,heavy)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 9.2 — `part_in_whole`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 316. **Stage recommendation:** `transcription-checked`.

REPAIRED: lower degree is no longer an exception defeating possession of the property. The positive species/part-to-property branch concludes has_property(y,A).

No printed CQ exists here. The removed exception was a conclusion parenthesis incorrectly promoted to a defeat condition.

The current predicate species_of represents the species reading; physical part and comparative degree variants need explicit scope or separate records. Do not claim the degree parenthesis has been fully computed.

**Example reviewed:** `part_in_whole/book`; expected accepted conclusion(s): `has_property(rose,beautiful)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

**New independent cases:** `part_in_whole/independent-lower-degree-is-not-a-defeater`.

### Entry 10.1 — `division_generic`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 317, 113. **Stage recommendation:** `example-passed`.

The whole has a property and, defeasibly, all parts have it. The direction is distinct from composition even though they can form an engine-safe cycle.

The divisional-heredity default is the substantive warrant; it is challengeable. The source's reversed explanatory parenthesis is corrected using narrative context.

Approve with the same property-identity and case-context convention as composition.

**Example reviewed:** `division_generic/book`; expected accepted conclusion(s): `all_parts_property(car,heavy)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 10.2 — `whole_into_parts`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 317. **Stage recommendation:** `transcription-checked`.

The implementation fixes three named parts and derives the whole's property from X1 plus the stated whole/part warrant. The source allows x1 through xn and alternative negative branches.

No printed CQs. The source itself says only if; treating this as a presumptive argumentative pattern does not establish a deductively valid sufficient condition.

Declare the positive, three-part restricted instance. Arbitrary cardinality, other selected parts, negative branches and the only-if interpretation remain unimplemented or unresolved.

**Example reviewed:** `whole_into_parts/book`; expected accepted conclusion(s): `has_property(team,strong)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 11.1a — `opposition_descriptive_1`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 317. **Stage recommendation:** `example-passed`.

The contrary-property warrant and asserted property concern the same subject A; the conclusion is inability to present the opposite at the same time.

No printed CQs. Modal cannot-present is kept distinct from a plain negated observation.

No semantic blocker for the printed variant.

**Example reviewed:** `opposition_descriptive_1/book`; expected accepted conclusion(s): `cannot_present_sametime(wall,black)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 11.1b — `opposition_descriptive_2`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

The negative property premise instantiates the antecedent of the stated opposition warrant; the conclusion names the paired property of the same subject.

No printed CQs. The warrant, not absence of any arbitrary property, licenses the inference.

No semantic blocker for this source-form record.

**Example reviewed:** `opposition_descriptive_2/book`; expected accepted conclusion(s): `must_present_sametime(number,odd)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 11.1c — `opposition_descriptive_3`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

Complementary-or-contrary alternatives and absence of one yield the required other predicate for the same subject and time context.

No printed CQs. This is the declared paired-property interpretation, not unrestricted closed-world reasoning.

No semantic blocker for the stated source form.

**Example reviewed:** `opposition_descriptive_3/book`; expected accepted conclusion(s): `must_present_sametime(light,off)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 11.2 — `opposition_normative`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

The corrected minor premise instantiates the first evaluation in the major warrant; the conclusion excludes the opposite evaluation for the same subject.

No printed CQs. The source's inconsistent symbol is corrected coherently rather than reproduced as a swapped role.

No semantic blocker for this documented correction.

**Example reviewed:** `opposition_normative/book`; expected accepted conclusion(s): `cannot_receive_evaluation(act,bad)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 12.1 — `rhetorical_oppositions_normative`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

Entity and treatment pairs retain the correct direction: the current required treatment triggers the warranted contrary treatment for the contrary entity.

No printed CQs.

No semantic blocker; source-form relational variables are not conflated with properties of the same entity.

**Example reviewed:** `rhetorical_oppositions_normative/book`; expected accepted conclusion(s): `contrary_treatment_required(enemies,severity)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 12.2 — `rhetorical_oppositions_descriptive`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

A's relation to B and the paired converse relation yield B's converse relation to A. Roles are reversed exactly once.

No printed CQs. The one implemented direction is warranted by the source's bidirectional description.

No semantic blocker for this direction of the relational pattern.

**Example reviewed:** `rhetorical_oppositions_descriptive/book`; expected accepted conclusion(s): `converse_relation(bob,learns_from,ann)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 13.1 — `alternative_cognitive`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `example-passed`.

The chosen alternatives are propositions A and B; negated A and the alternative warrant yield B. Bare propositions allow this output to feed other schemes.

No printed CQs. The rule remains a defeasible argument rather than a modal theorem about all possible worlds.

No semantic blocker for the stated alternative warrant.

**Example reviewed:** `alternative_cognitive/book`; expected accepted conclusion(s): `snow`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 13.2 — `alternative_normative`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 318. **Stage recommendation:** `transcription-checked`.

A and B remain the same normative alternatives. Required A yields negated required B under the supplied alternative warrant.

No printed CQs. This reading requires mutually exclusive normative alternatives; an inclusive A-or-B would not suffice.

Document exclusivity and distinguish not-required(B) from required(not-B).

**Example reviewed:** `alternative_normative/book`; expected accepted conclusion(s): `¬is_required(go)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 14 — `verbal_classification`

**Decision after patch:** approve. **Origin:** upstream. **Source pages:** 319. **Stage recommendation:** `transcription-checked`.

Object O classified under F and F subsumed by G yield O under G. Classification hierarchy and object positions are coherent.

CQ1 is the object-class premise; CQ2 is the doubtful-classification exception.

Book form approved; its own cited original remains subject to the inherited-source worklist.

**Example reviewed:** `verbal_classification/book`; expected accepted conclusion(s): `instance(rex,mammal)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 15 — `definition_to_verbal_classification`

**Decision after patch:** approve. **Origin:** upstream. **Source pages:** 319. **Stage recommendation:** `transcription-checked`.

REPAIRED: the definition satisfaction and definition-to-class link yield the correct object classification; the complete book CQ set is represented. The source's inconsistent x/a in its universal premise is read coherently, not as all objects having G.

CQ1 inadequacy remains an exception. CQ2 now has doubtful_definition(D,G), explicitly for a merely stipulative or biased definition subject to doubt. Positive and wrong-definition controls distinguish relevant from unrelated challenges.

Book-form interpretation approved; this addition is sourced to p.319 and does not verify the independently cited 2006 original. Inherited-source stage remains held.

**Example reviewed:** `definition_to_verbal_classification/book`; expected accepted conclusion(s): `instance(rex,dog)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

**New independent cases:** `definition_to_verbal_classification/independent-biased-definition-defeats`, `definition_to_verbal_classification/independent-other-definition-does-not-defeat`.

### Entry 16 — `verbal_classification_vagueness`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 319, 320. **Stage recommendation:** `transcription-checked`.

REPAIRED wording: the counterclaim concerns precision, not certainty. The argument and property roles in the vagueness/precision premises are coherent.

CQ1/CQ2 are premises; CQ3 is an opposing exception asserting that precision is not required here. This is an interpretation of a why-question, not the only possible burden policy.

Approve the content correction, but document when an unanswered precision challenge defeats the argument versus requiring an affirmative exception.

**Example reviewed:** `verbal_classification_vagueness/book`; expected accepted conclusion(s): `rejected_as_deficient(arg1)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 17 — `verbal_classification_arbitrariness`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 320. **Stage recommendation:** `transcription-checked`.

The arbitrary property, relevant argument and need for a nonarbitrary definition retain their roles; the conclusion rejects that argument as deficient.

CQ1/CQ2 are premises; CQ3 harmless arbitrariness is an exception. The burden choice is a project interpretation.

Keep that burden policy explicit; no extra premise should simply restate the conclusion.

**Example reviewed:** `verbal_classification_arbitrariness/book`; expected accepted conclusion(s): `rejected_as_deficient(arg1)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 18.1 — `interaction_act_person_1`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 320. **Stage recommendation:** `example-passed`.

A person's act with value V supports attributing V to that person. Agent, act and value remain separate.

No printed CQs.

No semantic blocker for this source-form presumptive transfer.

**Example reviewed:** `interaction_act_person_1/book`; expected accepted conclusion(s): `person_value(ann,brave)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 18.2 — `interaction_act_person_2`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 320, 321. **Stage recommendation:** `example-passed`.

A prior judgment of the person's value plus a stated value-to-act evaluation warrant yields the evaluation of that person's act.

No printed CQs. The generic warrant is intentionally not tied to one individual unless instantiated through the first premise.

No semantic blocker for the stated relational pattern.

**Example reviewed:** `interaction_act_person_2/book`; expected accepted conclusion(s): `person_acts_value(ann,report,truthful)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 19.1 — `positive_value`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 321. **Stage recommendation:** `transcription-checked`.

Agent A's judgment and its effect on A's goal lead to a reason for retaining the goal. The conclusion omits A.

No printed CQs. Two people with an identical goal term but different evaluations would share the conclusion's identity.

Require goal terms to identify the owner's goal instance or add the agent to the conclusion before mixed-agent value comparisons.

**Example reviewed:** `positive_value/book`; expected accepted conclusion(s): `reason_retaining_goal(life,survival)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 19.2 — `negative_value`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 321. **Stage recommendation:** `transcription-checked`.

The corrected conclusion retracts commitment, rather than retaining it. As with the positive record, the conclusion omits the evaluating agent.

No printed CQs.

Owner-scoped goal identifiers or an explicit agent argument are necessary for multi-agent use; this is separate from the already corrected sign.

**Example reviewed:** `negative_value/book`; expected accepted conclusion(s): `reason_retracting_goal(theft,survival)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 20 — `sacrifice`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 322. **Stage recommendation:** `transcription-checked`.

The person, sacrifice and its value are coherent in the great-sacrifice branch. The current record does not encode the small-sacrifice alternative printed parenthetically.

No printed CQs.

Approve only the great-sacrifice variant; add or explicitly exclude the small-sacrifice branch from coverage claims.

**Example reviewed:** `sacrifice/book`; expected accepted conclusion(s): `sacrifice_value(degree,worth)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 21.1 — `group_1`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 321, 322. **Stage recommendation:** `example-passed`.

A member's quality plus an explicit transfer warrant yields the group's quality. The has_property(group,quality) identity is the same property statement used elsewhere.

No printed CQs. The warrant, not membership alone, supports the transfer.

No semantic blocker when terms identify the same quality and group in the same case.

**Example reviewed:** `group_1/book`; expected accepted conclusion(s): `has_property(team,fast)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 21.2 — `group_2`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 322. **Stage recommendation:** `example-passed`.

A group's quality, membership and explicit downward transfer warrant yield the member's quality. The shared property predicate preserves entity roles.

No printed CQs. The membership and transfer premises prevent an unqualified automatic group-to-person inference.

No semantic blocker for the stated source form.

**Example reviewed:** `group_2/book`; expected accepted conclusion(s): `has_property(ann,wise)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 22.1 — `practical_inference`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 95, 96, 323. **Stage recommendation:** `transcription-checked`.

Agent P's goal and means-to-end support should_do(P,A); explicit agent binding makes the shared practical conclusion appropriate.

All five printed questions have a ledger route, but other/better means and bad consequences are opponent-proven exceptions. Narrative p.96 says merely asking any of the five questions can provisionally defeat until answered.

A static evidential exception policy is not the book's full dialogue burden behavior. Keep interpretation approval qualified until that mismatch is resolved or explicitly contracted.

**Example reviewed:** `practical_inference/book`; expected accepted conclusion(s): `should_do(me,exercise)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 22.2 — `practical_necessary_condition`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 95, 96, 323, 324. **Stage recommendation:** `transcription-checked`.

REPAIRED: P now scopes all personal selection, consideration, feasibility and preference premises, plus personal acceptability and alternatives. A second agent sharing a goal cannot inherit another's deliberation.

CQ coverage exists; exceptions versus questions retain the practical dialogue-burden discrepancy described for 22.1.

New positive and two-agent adverse cases pass after the patch. Resolve question-induced suspension policy before unqualified source approval.

**Example reviewed:** `practical_necessary_condition/book`; expected accepted conclusion(s): `should_do(me,study)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

**New independent cases:** `practical_necessary_condition/independent-one-agents-selection-does-not-bind-another`, `practical_necessary_condition/independent-another-agents-alternative-does-not-defeat`.

### Entry 22.3 — `practical_sufficient_condition`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 96, 323, 324. **Stage recommendation:** `transcription-checked`.

REPAIRED: all personal premises bind the same P; the alternatives template now says each member is sufficient, matching p.96 rather than merely one member. Selected member X is distinct from set S.

Same five question routes and extra feasibility defaults as 22.2; dialogue-burden qualification remains.

New positive and two-agent adverse cases pass after the patch. Document the narrative-supported correction and resolve provisional challenge behavior.

**Example reviewed:** `practical_sufficient_condition/book`; expected accepted conclusion(s): `should_do(me,study)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

**New independent cases:** `practical_sufficient_condition/independent-one-agents-selection-does-not-bind-another`, `practical_sufficient_condition/independent-another-agents-alternative-does-not-defeat`.

### Entry 22.4 — `value_based_practical_reasoning`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 95, 96, 324. **Stage recommendation:** `transcription-checked`.

Agent P, goal G, value set V and act A are connected through the means premise; practical conclusion should_do(P,A) is coherent.

All seven printed CQs are mapped. Value-based preference and alternative facts omit P and rely on goal/action terms carrying context; practical challenge burden remains a policy interpretation.

Use owner/context-specific goals and value sets, and settle the explicit challenge policy before full mixed-agent semantic approval.

**Example reviewed:** `value_based_practical_reasoning/book`; expected accepted conclusion(s): `should_do(hal,take_insulin)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 22.5 — `practical_goal`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 325. **Stage recommendation:** `example-passed`.

The same person's goal G and action X's contribution to that goal yield should_do(P,X). The named agent makes vocabulary sharing appropriate.

No printed CQs in this variant. This is a defeasible reason to act, not a proof that every contributing action must be jointly performed.

No semantic blocker for the source-form record.

**Example reviewed:** `practical_goal/book`; expected accepted conclusion(s): `should_do(ann,exercise)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 22.6.1 — `end_and_means_1`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 325. **Stage recommendation:** `transcription-checked`.

The record represents the positive branch: X is a means to good end Y, so X is good. The source also prints less-bad and bad alternatives.

No printed CQs. Evaluation is context-sensitive; the source is presumptive, not unrestricted moral closure.

Approve only the explicitly positive branch; missing parenthetical variants cannot be counted as tested.

**Example reviewed:** `end_and_means_1/book`; expected accepted conclusion(s): `good(training)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 22.6.2 — `end_and_means_2`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 325. **Stage recommendation:** `transcription-checked`.

The ends-importance warrant and means-to-end premise preserve the positive branch. The record's importance warrant is zero-argument and therefore global within a graph.

No printed CQs.

Limit this to one evaluation context and the positive branch; opposite valuation and mixed-context warrants need separate representation.

**Example reviewed:** `end_and_means_2/book`; expected accepted conclusion(s): `ends_important(money)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 22.6.3 — `end_and_means_3`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 325. **Stage recommendation:** `transcription-checked`.

Ease of achieving the means supports goodness of the end in the positive branch. Difficult-means/bad-end is not represented.

No printed CQs.

Approve only the easy/good branch and record the omitted alternative explicitly.

**Example reviewed:** `end_and_means_3/book`; expected accepted conclusion(s): `good(health)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 23 — `two_person_practical`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 325, 326. **Stage recommendation:** `transcription-checked`.

A states an intention to B; B supplies necessity information about action Y for A's intention X; the conclusion directs A, not B, to act.

All five printed questions are sorted. Alternative means and additional actions omit the deliberating agent and assume one current case.

Correct speaker/recipient roles are approved, but practical dialogue burdens and cross-agent availability of alternatives need an explicit contract or scoped facts.

**Example reviewed:** `two_person_practical/book`; expected accepted conclusion(s): `should_do(ann,ticket)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 24 — `waste`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 326. **Stage recommendation:** `example-passed`.

Agent A's stopping would waste efforts toward goal X; avoiding that waste supports continuing toward X. Feasibility concerns bringing about X.

CQ1 is the possible-goal default; CQ2 is reassessment against continuing, rather than an automatic rule that sunk costs justify every continuation.

Approve the stated static source interpretation; a newly supported forward-looking reassessment can undercut the argument.

**Example reviewed:** `waste/book`; expected accepted conclusion(s): `ought_to_continue(ann,thesis)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 25 — `sunk_costs`

**Decision after patch:** qualified. **Origin:** upstream. **Source pages:** 327. **Stage recommendation:** `transcription-checked`.

The inherited sunk-cost quantity/feasibility form differs from the compendium's time-indexed precommitment form. Keeping both as distinct records is correct.

Its feasibility and future-loss questions belong to the inherited form, not a fabricated CQ list for the book's precommitment form.

Do not certify it as a transcription of p.327. Its cited article/original must be independently checked by the inherited-source reviewer.

**Example reviewed:** `sunk_costs/book`; expected accepted conclusion(s): `should_be_performed(concorde)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 25 — `sunk_costs_precommitment`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 327. **Stage recommendation:** `example-passed`.

The choice at T2, prior commitment through T1, acting agent P and chosen A retain the roles of both printed premises; should_do(P,A) is the proper practical conclusion.

No printed CQs for this numbered form. Earlier/later ordering is part of the precommitted relation's meaning; arbitrary timestamps are not separately reasoned about.

Approve as the separate book-form variant; graph authors must not assert precommitted for a future rather than prior commitment.

**Example reviewed:** `sunk_costs_precommitment/book`; expected accepted conclusion(s): `should_do(ann,finish_thesis)`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.

### Entry 26 — `ignorance`

**Decision after patch:** qualified. **Origin:** upstream. **Source pages:** 327. **Stage recommendation:** `transcription-checked`.

Would-be-known(S) and not-known(S) yield not-S. This is the correct negative conclusion and is not a closed-world inference from mere absence.

CQ1 investigation is an exception; CQ2 dialogue burden and CQ3 evidential strength are both marked standard. A proof-standard threshold does not itself decide the dialogue's probandum or burden holder.

Own-source approval remains separate; specify how dialogue burden/probandum is supplied by an application before claiming all CQs are operationally implemented.

**Example reviewed:** `ignorance/book`; expected accepted conclusion(s): `¬spy(bob)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 26.1 — `negative_normal_expectations`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 327. **Stage recommendation:** `transcription-checked`.

The warranted normal expectation of A and explicit not-A support a non-normal situation, identified by A in the current encoding.

CQ1 investigation is represented. CQ2/CQ3 have the same burden/proof-standard distinction as ignorance.

Document one expectation context or add a situation identifier; distinguish dialogue burden allocation from strength of proof.

**Example reviewed:** `negative_normal_expectations/book`; expected accepted conclusion(s): `not_normal(alarm_silent)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

### Entry 26.2 — `negative_practical`

**Decision after patch:** qualified. **Origin:** book. **Source pages:** 327. **Stage recommendation:** `transcription-checked`.

REPAIRED: B>C now yields acting on the false presumption, explicitly. The opposite C>B branch is not silently collapsed into the same ambiguous action.

CQ1 is an evidence-progress default. CQ2 burden allocation and CQ3 strength still require distinct external dialogue information.

Negative branch is source-correct and tested; declare the missing true-presumption branch and resolve dialogue metadata before unqualified approval.

**Example reviewed:** `negative_practical/book`; expected accepted conclusion(s): `act_on_presumption_false(loaded)`. The existing /book grounding exercises the represented branch; passing it does not resolve the specific scope/burden issue in this adjudication.

**New independent cases:** `negative_practical/independent-false-presumption-selected`, `negative_practical/independent-opposite-ranking-does-not-select-false`.

### Entry 27 — `epistemic_ignorance`

**Decision after patch:** approve. **Origin:** book. **Source pages:** 328. **Stage recommendation:** `example-passed`.

All five premises are present, including the bivalence of domain D. The proposition A and knowledge base K retain their membership roles; the conclusion is not-A.

No printed CQs for this separate schema. The awkward predicate name true_propositions_contained has negative natural-language content, but it is an opaque fact, not a computed double negation.

Approve the printed source form; no inference from an open-world database is licensed unless all five premises are asserted.

**Example reviewed:** `epistemic_ignorance/book`; expected accepted conclusion(s): `¬spy`. The existing /book grounding is consistent with the approved reading; it is not evidence for all inputs or all parenthetical source variants.
