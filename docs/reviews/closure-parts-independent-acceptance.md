# Independent acceptance of entries 9.2, 10.2 and 16

**Reviewed implementation:** source-review commits `611e0c7d81501538ac8cb9750b748681456e78f8` and `04d586c9c4272ae0a3c2a0b7490a8ecb84b24e88`. **Reviewer:** ChatGPT independent closure scope reviewer, 2026-11-21. The reviewer did not author these formulas or their cases. The supplied book’s printed pp.316–317 were rendered and visually checked, including the exact only-if direction; printed pp.319–320 had also been visually checked. Every new formula, language template and expected case outcome was inspected independently.

The decision includes the four physical/degree forms and two general partition forms actually implemented. It does not grant coverage by renaming the two old narrow schemes. The former three-part positive division rule was additionally repaired: a necessary condition by itself cannot license its positive conclusion.

| Record | Decision and reason | Evidence examined |
|---|---|---|
| `part_in_whole` | Approve the existing positive species branch now that the physical and degree alternatives have their own executable records. The taxonomic relation does not silently become physical parthood. | Existing book example and physical-is-not-species control. |
| `part_in_whole_physical` | Approve the separately stated physical-part branch of p.316. It is a defeasible source pattern, not a certification that every empirical property always transfers to every part. | Gold-fragment constructed example and species-is-not-physical control. |
| `part_in_whole_species_degree` | Approve the qualitative lesser-degree source reading with an explicitly supplied common property scale. Possession and lower degree are compatible conclusions; no numerical magnitude is computed. | Positive species degree, absence of comparability, and no negation-of-possession check. |
| `part_in_whole_physical_degree` | Approve the analogous physical-part reading with the same explicit scale qualification. | Positive regional example and a scale belonging to a different property. |
| `part_in_whole_smaller_degree` | Approve the smaller-part reading. Both a relevant smaller-size claim and a common degree scale are required; neither is inferred from missing evidence. | Positive smaller-part example and missing relative-size premise control. |
| `whole_into_parts` | Approve as a disclosed strengthened application interpretation of p.317: its positive conclusion now needs an ordinary sufficient-direction premise. The printed only-if statement is retained as necessity, not silently reworded as sufficiency. | Updated book application; legacy only-if without sufficiency refuses inference; a different selected part’s warrant is insufficient. |
| `whole_into_parts_selected` | Approve the arbitrary finite partition form with explicit membership and a sufficient aggregate warrant. The template unambiguously says that existence of at least one property-bearing member is sufficient. It does not select an unknown special member or reverse only-if. | Seven-member and two-member cases; necessary-is-not-sufficient, wrong partition and wrong whole controls. |
| `whole_into_parts_negative` | Approve the negative source branch using the necessary warrant and affirmative universal-negative evidence over the complete same partition. This implements the stated negative condition without a closed-world inference. | Positive negative-branch case; missing universal evidence, one negative member, another partition, and sufficiency-in-place-of-necessity controls. |
| `verbal_classification_vagueness` | Approve the implementer’s problem-for-this-dialogue wording and the documented static burden. Supported warrant, need and property inadequacy are required. Supported harmlessness defeats the argument; merely questioning that exception does not prove it. | Six cases: positive, missing precision requirement, supported harmlessness, question-only exception, wrong argument and wrong property. |

The opaque set identifier denotes a complete finite nonempty partition supplied by the application. The engine does not count members to establish completeness. The two-member fixture need not enumerate its other member because the complete-set premise is explicitly supplied; the seven-member fixture demonstrates that any listed selected member can act as the witness. Neither fixture proves external completeness of a real-world partition.

On entry 9.2 the general qualitative-degree and the explicitly smaller-part readings can generate overlapping arguments if an application loads both and supplies both sets of premises. The contract identifies them as alternative readings; overlap is not additional independent measurement. On entry 10.2 the supplied source contains the directional tension between its only-if premise and positive conclusion. Explicit sufficiency is a visible executable strengthening, while the negative branch retains the printed necessary direction. This is an honest interpretation approval, not a claim to repair the printed book or certify the separately referenced Perelman text.

## Independent execution and exact approval

In the independent clone, the two implementation commits were combined with the prior reviewed context changes. The new cases were run with Go 1.24.13, local pinned engine modules, `GOTOOLCHAIN=local`, `GOFLAGS=-mod=mod` and `VV_STRICT=1`:

```sh
go test -count=1 -v -run '^TestCases/cases/(closure_parts|closure_vagueness)'
```

**54 API/CLI checks passed, zero known failures, zero unexpected failures, zero stale ledger entries.** The execution log `parts-vagueness-independent.log` has SHA-256 `576a467a07434f75ed1f9fbcf843a3cfd3e8649b10c6f790e5324d53f710e79f`.

Each of these nine records receives its own current semantic fingerprint, explicit case references and independent review stamp. The six additions plus three formerly held records raise this clone’s approval count from 83/120 to 92/126. That is a stage-accounting result for this reviewed state, not the final count after further additions. Full acceptance of the combined deliverable remains a separate integration gate.
