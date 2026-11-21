# First-half reconciliation of independent PR13 findings

Date: 2026-11-21. Peer review: `96daf6c0545f7cb84fe5e30cfbffdde5d6f6c0d1` (PR13, reviewing the earlier `49ad425` implementation). Integration baseline: `a708add511f52fb3603d92ba368875ef3d3f2c42`.

This reconciliation accepts new source-supported findings without substituting the peer's approval count for record-specific judgment. The supplied book's printed pp. 97, 311, 314, 315, 318, 321 and 325 were visually inspected. The earlier first-half review missed three wording errors and two safe shared-statement opportunities. They are corrected here. The canonical source-stage population remains 70 approved and 50 held; `two_person_practical` retains its separate scope/burden hold.

## Three wording corrections

| Record | Source finding | Corrected interpretation |
|---|---|---|
| `popular_opinion` | Entry 4 CQ2, p. 311, asks for good reasons to doubt A's truth, even if A is generally accepted. It does not ask whether A is generally true. | `doubting_general(A)` now reads as good reasons to doubt that A is true. The identifier and opponent-supported exception role remain unchanged. |
| `rhetorical_oppositions_descriptive` | Entry 12.2, p. 318, explicitly puts Y in the converse relation with X, and adds the reverse direction. | The warrant template now identifies both participants in the reversed order and includes "contrariwise". The existing conclusion already had the correct argument order. |
| `two_person_practical` | Entry 23, p. 325, and the narrative form on p. 97 say that the expert sees the situation, not that the expert sets it. | The template says "sees" and preserves the adviser's communication to the agent. Its B/Y/X/A argument roles remain unchanged. |

These are corrections to the language contract, not new argument rules. The two already approved records remain approved for their corrected source-supported interpretations. The two-person record remains held for its independently identified dialogue and alternative-action scope questions. All three changes have wording entries in the migration table. Repeated participants are expressed with unambiguous "latter/former" and "adviser" references because the current importer requires exactly arity-many format verbs; an attempted repeated-index template was rejected during pre-commit validation and corrected.

## Two safe statement unifications

### Example and illustration

Entry 6.1, p. 314, says that a particular individual has F and G. Entry 6.2, p. 315, uses the same individual/property/property assertion as its second premise. `same_properties(Y,F,G)` did not denote similarity between two individuals or a different relation: its three arguments have exactly the roles of `has_properties(Y,F,G)`.

`illustration` now consumes `has_properties(Y,F,G)`. The migration table renames `same_properties/3` to `has_properties/3`. The existing `has_properties` template serves both records. A two-scheme case supplies only one conjunction fact: `from_example` produces the shared generalization and `illustration` consumes that generalization and the same conjunction fact. Wrong-individual and wrong-property argument instances must remain absent. This does not settle `from_example`'s separate strength-policy hold.

The new interoperability also changes two old test assumptions. Withdrawing a stipulated generalization does not remove it from the full theory's closure: another scheme can derive it. Questioning the induction default alone does not prevent structural argument construction, even when it prevents acceptance. The generic `premiseProblems` check and reference-case `variantProblems` therefore isolate the actual selected schema for their premise-withdrawal controls, after a successful isolated positive control. Ordinary full-library runs, their expected arguments and grounded labels, exception checks, and the new chaining cases retain the full theory. The older adverse-properties fixture now forbids the specific flies argument, preserving its original binding test while allowing an independently induced sings argument.

This is a documented correction to the oracle contract, applied generically rather than as an illustration waiver. The synthetic regression proves full-theory rederivation remains valid, isolated withdrawal succeeds, a truly superfluous reference premise is still rejected, and a missing isolated positive control cannot count as a successful test. The original red logs and the unsuccessful default-questioning attempt are retained in the review evidence. No rule is copied into a fixed test fixture merely to bypass the current library.

### Interaction of act and person

Entry 18.1, p. 321, concludes that person P is V. Entry 18.2 begins with person P is V1, explicitly identifying V1 as a judgment value. This is the same person/valuation relation. The pair is `interaction_act_person_1/2`, not the positive/negative value pair in entry 19.

`interaction_act_person_2` now consumes `person_value(P,V1)`. The migration table renames `person_judgment_value/2` to `person_value/2`. A two-scheme case derives a person's valuation from an act and feeds it into the second scheme with the value-conditioned warrant. Wrong-person and wrong-value conclusions and argument instances remain absent. This preserves the existing convention that the valuation terms refer to the same evaluation context; it does not equate two different agents' valuations merely because their words match.

## Approval fingerprints and review evidence

Five record fingerprints change: `popular_opinion`, `rhetorical_oppositions_descriptive`, `two_person_practical`, `illustration`, and `interaction_act_person_2`. Four already approved records receive reviewed replacement fingerprints and this addendum as evidence. No approval is created for the held two-person record. Source stages and the 70/50 population do not change.

The two new source-grounded cases are:

- `illustration/peer13-shared-conjunction-chains-from-example`.
- `interaction_act_person_2/peer13-shared-person-valuation-chains`.

Existing cases are migrated to the shared statement names, including the earlier independent two-examples case. The migrator's existing chain/merge test also checks both new aliases. Category files are regenerated from the master.

The integration owner must regenerate the complete expected-check inventory from a full run after combining peer metadata and this patch. A selected run must not replace that full inventory. The earlier clean acceptance at `a708add` is historical once this patch is applied; release evidence must refer to the final integrated commit.

## Findings already resolved or still held

The peer independently confirms problems already repaired in the integration branch: the example-dependent generalization, agent leakage in necessary/sufficient practical reasoning, the sufficient schema's each-one wording, the mistaken part/whole exception, negative-practical polarity and its independent adverse case, and the precision/certainty wording. It also confirms the remaining limitation of the three-part whole-to-parts record.

The peer's separate objection to the original `negative_practical/book` fixture is also accepted. The earlier repair changed the rule and added a sound independent safe/unsafe example but left the original `loaded` proposition paired with reversed firearm consequences. This reconciliation changes that illustrative proposition to `unloaded`: treating an actually loaded firearm as unloaded risks an accident; treating an actually unloaded firearm as not-unloaded costs delay. With accident more serious than delay, `act_on_presumption_false(unloaded)` is the coherent negative-presumption conclusion. This corrects example content, not the formula or its held source stage. The source locator explicitly calls it an application of the printed schema rather than claiming the book prints this example.

Peer approval does not close our remaining findings about quantifier/branch coverage, evidence strength, dialogue challenge burdens, agent/context scoping, or source-form mismatch. The peer's approval of the old `from_model` coexists with a ledger note acknowledging that its warrant omitted the conditional antecedent; the integration branch's source-supported conditional repair must be retained. The peer's broader source approvals and its historical review metadata are evidence of that review, not authority to overwrite the final disposition of a different representation.

The generic division observation is valid: with its divisional-heredity default undisputed, an asserted whole-property fact can support an all-parts conclusion. That is the declared defeasible rule, not a proof that every property is hereditary. Whole-library loading requires the stated context/identity and challenge conventions; the application must select appropriate schemes or challenge the hereditary default. The observation should be visible in usage guidance and must not be described as unrestricted semantic safety of arbitrary whole-library facts.

## Focused verification before integration

On Go 1.24.13, the selected `TestLibrary`, `TestCases` and `TestSource` run
passed **7,861 registered checks**, with 7,861 unique `(file, scheme, check)`
keys, zero known/unexpected failures and zero stale ledger entries. This covers
all master/category schema checks, including illustration in both contexts.
The new premise-isolation regression passed. The scheme validator reports
0 errors and 11 warnings; all 5 validator tests and all 8 migration tests pass.
Category-file parity and approval integrity for all 120 records pass.

This selected verification is not the final complete release run. No complete
expected-check inventory was regenerated from it. The integration owner must
run the full gate and both toolchains on the final merged commit.
