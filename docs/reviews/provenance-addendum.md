# Dated addendum to the independent entries 1–27 adjudication

Date: 2026-11-21. This addendum preserves the baseline and original review history while correcting a provenance overstatement identified during the final integrated 120-record audit.

## Verbal classification: remove a false source-access blocker

The first-half adjudication approved the supplied-compendium formulation of `verbal_classification`, entry 14, printed p. 319, but used a generic inherited-record hold: “its own cited original remains subject to the inherited-source worklist.” That explanation was inaccurate for this record. The upstream code cited no separate original work. The current record cites the supplied compendium, which the independent reviewer directly inspected.

The final comparison with `third_party/carneades-4/examples/AGs/YAML/walton.yml` established the following:

| Component | Upstream code | Current reviewed interpretation |
|---|---|---|
| Variables | O, F, G | O, F, G |
| Individual premise | `instance(O,F)` | `instance(O,F)` |
| Classification premise | `subclass(F,G)` | `subclass(F,G)` |
| Conclusion | `instance(O,G)` | `instance(O,G)` |
| Exception | None | `doubtful_classification(F,G)` |

The sole behavior addition is the supplied book's CQ2 about a doubtful verbal classification. It is already disclosed as a book-supported interpretation in `tools/migration_r2.yml`. CQ1 is the individual premise. The Rex/dog/mammal reference case and its missing-premise/exception variants were included in the independent source review.

The resulting decision is **approve the declared compendium/code-origin interpretation and advance to `example-passed` when included in the passing integrated gate**. No claim is made that the 2006 work referenced by the book's bibliography was independently inspected. Upstream origin does not create an obligation to inspect a nonexistent separate code citation.

This correction raises the consolidated approved population from 69 to 70, leaving 50 held: 48 transcription-checked and 2 mapped. The approved origins are65 book and 5 inherited.

## Related provenance clarifications

`cause_to_effect` also lacked a separate source citation in the inherited code. Its real remaining qualification concerns causal-application context, especially the scope of `interference(E1)`. The supplied compendium was inspected. An unavailable unidentified original is not the reason for its hold; it remains qualified for the independently identified semantic scope.

The first-half review did not itself inspect the original 2002 sunk-cost article. Root subsequently obtained the author's article text with printed-page markers through the author-paper page on Academia, and independently checked printed pp. 488,489 and493. This confirms the distinction between the legacy quantity/waste adaptation and the article's printed precommitment scheme. `sunk_costs` therefore remains qualified for its adaptation and practical-question policy, not because the article is unavailable. The access method was web text; a binary PDF mirror refused access, so no PDF hash or visual inspection of that article is claimed. `sunk_costs_precommitment` remains approved for its directly inspected book formulation, without claiming complete CQ coverage of the whole article.

The complete Atkinson and Bench-Capon 2007 paper was supplied and independently reviewed by the original-source reviewer. The 17-question inventory includes the additional question introduced on p. 859. Eight operational mappings remain incomplete and additional policy differences remain explicit. This is an interpretation/implementation hold, not a source-access hold.

## Review boundaries retained

The other three first-half book-form approvals with separately cited originals—`position_to_know`, `expert_opinion`, and `definition_to_verbal_classification`—remain held pending those original-source comparisons. None is certified merely because its compendium formulation matches. All semantic qualifications in the original first-half adjudication remain in force unless explicitly resolved above or by the final integrated record-specific disposition.

The original adjudication files are preserved unchanged as historical evidence. `consolidated-disposition.csv`, `consolidated-disposition.json`, and `consolidated-disposition.md` contain the final reconciled 120-record result, with access obligations, concrete next actions, and source/CQ distinctions.
