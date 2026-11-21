# Independent checks of inherited-source provenance

Review date: 2026-11-21. Baseline: `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`. This record distinguishes inspection of a source, agreement with its printed form, and acceptance of an executable interpretation. The final source ledger and complete per-record dispositions are authoritative for stage counts.

## abduction

The inherited record names the original Carneades YAML as its only source. The original file is available inside the SHA-256-locked `carneades-4.zip`; comparison is performed against that archive, not an evolving online copy. Variables, premises, assumptions, exceptions and conclusions match the code source. The `more_coherent_explanation(T,S)` exception is explicitly recorded as an additional `n: 0` implementation hook, because that source prints no numbered critical-question list. `source_questions: 0` is therefore verified for this declared code source. The existing `abduction/book` case and adverse controls exercise it. Approved as a transcription of the named code source, not as a claim that it is identical to every philosophical account of abduction.

## defeasible_modus_ponens

The same archive-source method applies. The code's variables, premises and conclusion match. It prints no numbered question list, so zero is known for this declared source. The source comparison includes a deliberate altered-record mutation, which the gate rejects. The reference example and premise controls pass. Approved as a transcription of the named code source.

## credible_source

Wyner's *A functional perspective on argumentation schemes* (2016), printed p.127, has the source/domain, assertion and proposition/domain premises with the proposition as conclusion. Its discussion on p.130 leaves critical-question treatment for further work. The three exceptions in this library are inherited Carneades additions, not a three-question list supplied by Wyner. Approval covers that declared extension with one application context; it does not attribute the exceptions to the article. The zero printed-question count is specific to that displayed source form. Source: https://journals.sagepub.com/doi/pdf/10.3233/AAC-160010 .

## appearance

Walton's 2006 article, printed pp.324–325, supports the appearance inference and the two stated challenges. The library's scoped reliability/contrary-appearance hooks and the source reviewer’s examples are acceptable for the declared interpretation. The source was inspected through the journal PDF, also downloaded and hashed in the evidence inventory. Source: https://www.logiqueetanalyse.be/archive/issues87-220/LA195/06walton.pdf .

## verbal_classification

The earlier first-half report mistakenly described an unavailable separately cited original as its blocker. The current record cites the supplied compendium; the original Carneades record has no separate source citation. The independent reconciliation compared the original code: variables `O,F,G`, premises `instance(O,F)` and `subclass(F,G)`, and conclusion `instance(O,G)` agree. The additional `doubtful_classification(F,G)` exception explicitly models compendium CQ2, printed p.319, and is in the migration table. The supplied book, code origin and Rex/dog/mammal example were reviewed. Approval is for this declared compendium interpretation, with code provenance; it does not claim examination of a separate 2006 book. The final review corrects this false blocker and promotes the record.

## sunk_costs and sunk_costs_precommitment

The published article’s text was accessible through the author-paper copy at https://www.academia.edu/13186555/The_Sunk_Costs_Fallacy_or_Argument_from_Waste . Printed p.488 presents waste reasoning; p.489 presents choice at a later time and prior commitment as the precommitment form. This confirms the author’s page-489 claim and the separate compendium record. Printed p.493 also discusses practical-reasoning questions; “no printed questions in compendium entry 25” must not become “the article has no critical questions.” The inherited `sunk_costs` form remains a qualified adaptation. The approved `sunk_costs_precommitment` record is scoped to the supplied compendium. PDF download from another mirror returned 403; this check used accessible article text with printed page markers, not a claimed binary-PDF inspection. Publisher metadata: https://link.springer.com/article/10.1023/A:1021108016075 .

## practical_reasoning1

The supplied Atkinson and Bench-Capon PDF resolves source access completely. Its SHA-256 is `dad4124c052c48b5c82de02494aac6971de113bb8452ae9fc6e764aa35e43cbc`. The full-paper review, all 17 question dispositions and corrected source examples are in `atkinson-2007.md` and `atkinson-2007-decisions.json`. The record stays at `transcription-checked`: eight mappings remain incomplete, some existing alternatives/value predicates are stronger or coarser than the source, and the flat record does not implement full AS2 transition/value semantics. This is a content hold, not an access hold.

## practical_reasoning2

An official Cambridge excerpt of *Goal-Based Reasoning for Argumentation* (2015) provides introductory goal/means discussion on pp.2–3 but does not verify the record's complete formulation and question treatment. Source: https://assets.cambridge.org/97811071/19048/excerpt/9781107119048_excerpt.pdf . The limited preview is insufficient to advance this record; it remains mapped with an unknown own-source question count.

## Remaining originals and correspondence

The other inherited records' exact obligations appear individually in the complete 120-record disposition register. Relevant separately cited works include *Fundamentals of Critical Argumentation* (2006), *A Pragmatic Theory of Fallacy* (1995), *Appeal to Expert Opinion* (1997), *Legal Argumentation and Evidence* (2002), *Slippery Slope Arguments* (the cited edition must be reconciled), and the witness-testimony sources. No claim is made to have verified inaccessible pages in those works. Some records have additional contextual or modality defects even if an original becomes available; acquisition alone will not approve them.

The repository's historical account of the September 2019 correspondence remains provenance reported by the repository, not independently inspected correspondence. Preserve `exeption_to_rule` and its hold; do not silently identify it with the compendium’s separately implemented `from_precedent`.

## Approval scope

Inherited origin does not automatically create an obligation to acquire a nonexistent separate source. Conversely, reading the compendium does not automatically verify a separately claimed original. The five approved inherited records are `abduction`, `defeasible_modus_ponens`, `credible_source`, `appearance`, and `verbal_classification`. Together with 65 approved book-origin records they give70 approvals. Every approval is bound to the exact reviewed semantics by `independent-approvals.json`; the remaining50 stay below interpretation approval.
