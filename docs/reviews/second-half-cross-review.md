# Bounded independent cross-review of the second-half source repair

**Reviewed commit:** `9f96692387857014157820232946be7bd58a3c02`, in `book-second-work`.

**Scope:** Entries 31.2–31.5, 54.1, 56 and 60, plus the consequence-CQ interpretation requested by the parent. I inspected actual scheme terms, language templates, source-ledger claims, the new reference-case expectations and `docs/source-contracts-second-half.md`. The decisive book passages on printed pp. 330, 331, 342, 344 and 346 were visually checked; the start of entry 31.2 on p. 329 and consequence questions on pp. 332–333 were also read. This is a bounded review of the named repairs, not a second certification of every record in entries 28–60.

The two new independent reference-case files were independently rerun using the baseline engine against the second-half repaired library: **86 of 86 API/CLI round-trip checks pass**. See `second-half-cross.log` and `second-half-cross-results.jsonl`. The source judgments below do not follow merely from that result; they concern whether the expectations express the source or a defensible stated interpretation.

## Decisions

| Entry | Decision | Source-based reason and controls |
|---|---|---|
| 31.2 | Accept the stated interpretation | The book distinguishes diagram G1 from its conclusion C1. `successful_diagram(G,G1)` identifies the selected diagram in the named set; `diagram_conclusion(G1,C1)` supplies the missing relation. Positive, wrong-set and wrong-diagram-link controls now express those roles correctly. The output is the conclusion proposition's plausibility, not the diagram's. Adequacy and inquiry-completeness remain declared qualitative aggregates, not computed numerical rankings. |
| 31.3 | Accept | The source names an agent, a performed action and a particular character quality. The new A/X/Q parameters retain all three. `character_quality(A,Q)` preserves the type; `has_property(A,Q)` is a justified generic projection. The action-fit and quality-definition questions remain separately scoped. The chaining and wrong-quality/agent controls are meaningful. |
| 31.4 | Accept role/type repair, with the qualification below | Requiring `character_quality(A,Q)` prevents arbitrary ordinary properties such as a rose's redness from triggering character prediction. The conclusion retains the source's conditional future tendency and does not falsely identify a past actor. The wrong-quality and ordinary-property controls are appropriate. |
| 31.5 | Accept the stated interpretation | The observed event E is retained in the candidate premise, event-quality premise, attribution conclusion and relevant questions. Different event and different quality controls reject the correct substitutions. A character-derived generic property may feed this record because the event-quality premise independently identifies Q as the relevant character quality; generic property alone does not suffice. |
| 54.1 | Accept as an explicit executable interpretation | The printed warrant is conditional on referential coextension. Mere class inclusion does not establish that antecedent. Requiring `referentially_coextensive(W,D)` is a well-supported completion of the warrant, and the documentation correctly avoids calling it a publisher-confirmed erratum. Whale/mammal and unrelated-pair controls are substantive. Approval does not endorse the universal philosophical truth of the source's synonymy warrant. |
| 60 | Accept recaller-specific history | The source's original-belief history belongs to the person recalling M. Adding P to the first two undercutters prevents Ann's defect from automatically defeating Bob's independent recall. Same-person defeat, other-person non-defeat, convergent support and both-supports-defeated cases correctly test that distinction. |
| 56 | Accept the explicitly declared conditional witness policy | The source's second question applies when an exception is unrecognized; the third applies when the case is borderline. The four witness routes preserve both conditions. Their negative statuses are structural explicit negations, not absence of facts. Recognition cannot bypass borderline comparisons; comparison alone cannot substitute for non-applicability evidence in an unrecognized case. Wrong-case and unknown-status controls are correct. The core aggregate remains a defeasible default until challenged, and the documentation distinguishes witness rules from book-printed schemes. |

## Narrow qualification for 31.4

`not_character_action(A,Q)` currently says that A has actions that do not fit the defined quality. The ordinary conclusion predicts that future conduct is *likely* to fit the trait. A single isolated counterexample does not mathematically refute a statistical tendency. Treating such counterevidence as an automatic undercutter is therefore a conservative qualitative policy, not a probability theorem and not an exact equivalence supplied by CQ3.

The role repair should be retained. To make the approval scope fully explicit, either document that conservative undercutter policy or word the predicate as relevant counterevidence sufficient to defeat this particular trait prediction. This is a focused interpretation clarification; no actor/quality mismatch remains in the inspected repair.

## Consequence likelihood

The second-half patch's `consequence_likelihood_adequate(A,G)` is consistent with the safe qualitative policy recommended in the engine cross-review: adequacy of the likelihood assessment for this argument is separate from empirical support for the causal premise and from opposing consequences. It does not impose an undocumented universal numerical cutoff or equate all uncertainty with “strongly likely.” The separately cited 1995 original remains explicitly unverified. That is the correct provenance distinction.

No further broad reviewer round is needed for the named role and scope repairs. The parent should integrate them, preserve their declared interpretation boundaries, resolve the narrow 31.4 wording/policy note, and run the final unified acceptance suite.
