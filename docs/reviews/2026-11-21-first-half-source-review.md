# Independent source review: book entries 1–27

Reviewed baseline: `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`.
Source: the supplied *Argumentation Schemes* (Walton, Reed and Macagno, 2008),
chapter 9, printed pp. 309–328, with narrative pp. 95–96, 113 and the book's
Carneades discussion pp. 374–375 where relevant. Printed pages have PDF page
numbers ten greater. This review covers 63 mapped records, excluding inherited
`practical_reasoning1`, which has a separate review of its own source.

Passing a graph generated from a scheme does not establish that the scheme says
what its source says. The independent review found the following discrepancies
and added eleven reference cases in `tools/vv/cases/independent_first.yml`.

| Record | Source finding | Correction |
|---|---|---|
| `from_example` | The conclusion on p. 314 generalizes from properties F to G; the example individual is not part of the generalized proposition. | Use `usually_properties(F,G)`, already the rule premise of `illustration`; keep the example in the supporting argument's parameters. Two distinct examples now support one proposition. |
| `from_model` | Premise 3 on p. 315 is a conditional tied to the same model. | Restore `imitation_warrant(model,action,imitators)`; do not require only its already-detached consequent. An adverse case rejects another model's warrant. |
| `part_in_whole` | The parenthetical lower degree on p. 316 is part of the conclusion; the entry prints no critical question. | Stop treating `smaller_part_property` as an exception defeating possession of the property. Preserve the positive property branch. |
| `practical_necessary_condition`, `practical_sufficient_condition` | Every I/me in pp. 95–96, 323–324 denotes the same deliberating agent. Unscoped selection/practicality facts made another agent with the same goal inherit the recommendation. | Add the agent to all personal deliberation premises, acceptability assumptions and alternative-means exceptions. Verify positive and two-agent adverse cases. |
| `practical_sufficient_condition` | Narrative p. 96 says each alternative is sufficient, whereas the compendium's p. 323 alternatives line has the necessary-schema wording. | Use the narrative-supported each-one sufficient reading; the old template had fixed necessary to sufficient but retained at-least-one. |
| `negative_practical` | When B is more serious than C, p. 327 chooses acting on the presumption that A is false. | Use `act_on_presumption_false(A)`; the opposite ranking does not trigger this negative branch. |
| `definition_to_verbal_classification` | CQ2 on p. 319 asks about a merely stipulative or biased definition subject to doubt. | Add `doubtful_definition(D,G)` with book provenance and relevant/other-definition controls. The separately cited 2006 original remains unverified. |
| `verbal_classification_vagueness` | The question concerns precision, not certainty. | Correct the counterclaim's wording; preserve its identifier for compatibility. |

All changes to statement shape are listed in `tools/migration_r2.yml`. Agent
identity and conditional models cannot be recovered from the affected old
statement alone, so those rows require manual review. They are not automatic
renames. Existing generated category files and reference cases have been updated.
The degree-comparison template is retained so existing graphs can state a lower
degree without inadvertently defeating property possession; it now produces an
intentional unused-template warning.

## Review limits that still affect source approval

The patch does not blanket-promote source-ledger stages. These distinctions
remain necessary:

- A book-form match does not verify the separately cited original of one of the
  eight inherited records in this assignment. A secondhand citation is not a
  substitute for inspecting that original.
- The book's p. 96 practical-reasoning discussion says asking any of its five
  questions can temporarily defeat the argument until adequately answered.
  Several current mappings require a supported contrary exception. That is a
  defensible static evaluation policy only if stated as an interpretation; it
  does not implement the complete dialogue burden behavior.
- An issue's proof standard determines comparative evidential sufficiency. It
  does not itself specify who holds the dialogue burden or what the ultimate
  probandum is. The ignorance-family CQ2 and CQ3 must retain that distinction.
- `definition_to_verbal_classification` now represents the book's biased/stipulative
  definition question from p. 319. This is a clearly attributed book-supported
  interpretation; it does not claim access to the separately cited 2006 original.
- Source alternatives are not all executable branches. `whole_into_parts` is a
  positive instance with three parts; `sacrifice` and the end/means records use
  positive branches; `negative_practical` now explicitly selects its negative
  branch. Snob appeal and the excuse subtype combine action, policy or object
  readings. These limitations must remain visible in coverage claims.
- Where the source has one current case, speaker, respondent or value context,
  graph authors need distinct context-specific terms when combining cases.
  Named-agent conclusions such as `should_do` need named-agent premises too.
- Under a binary recommendation policy, `should_do` and its negation are
  competing positions. This does not make not-obligatory(A) identical to
  obligatory(not-A) in unrestricted deontic logic.

The independent row adjudication distinguishes 35 book-form approvals and 28
qualified readings after these repairs. Of the 35 approvals,
four are inherited records awaiting their own-source verification; 31 book-origin
records are candidates for `example-passed` after the integrated suite passes.
Qualified readings can advance once their explicit scope and burden decisions are
adopted and tested. They must not be advanced merely because a current `/book`
case passes.
