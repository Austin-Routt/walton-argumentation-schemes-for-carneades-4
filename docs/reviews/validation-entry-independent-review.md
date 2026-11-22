# Beginning validation after framework convergence

Independent planning review, 22 November 2026 (UTC). The inspected public repository was `d9a25857277140a5d3f10aaf1dfa4ade0f15b91f`, library 0.4.1 / engine 4.3.1+walton.7. The parent reviewer owns acceptance of that integration and any subsequent correction. This document does not certify the server patch, rerun the engine, or begin a benchmark campaign.

## Recommendation

**Begin the validation effort with a short protocol-and-reference stage, and use Kialo as the first application-data pilot alongside independent formal-engine work.** Kialo should not be the sole correctness benchmark, nor should thousands of captured graphs be translated and evaluated before their semantics are defined. The current V0–V6 and K0–K6 plans already support this sequence. Their historical implementation statuses need an explicit current binding, not another general review round. [E1–E4]

The useful immediate output is a versioned pilot protocol with a frozen implementation pin, a structural data contract, source-only annotation packets, an independent scenario reference, explicit exclusions and a run-record schema. Initial data auditing and annotation-method development can proceed while a bounded CAES/CHR reference evaluator is developed. Claims about conditional engine accuracy wait for that independent reference; claims about natural-argument fidelity wait for human annotation and adjudication.

A remote tag is useful release bookkeeping but is not a prerequisite for reproducible experiments: a full immutable Git commit plus file hashes supplies the experimental pin. Complete the non-author review of the integration changes before publishing a tag described as independently accepted. Do not let a blocked tag push delay protocol preparation, and do not call a new tag evidence of validation.

## What the existing Kialo material actually provides

The supplied example is one captured English-language discussion, “Does God exist?”, discussion 2629. The minimized derivative preserves 8,319 claim records and 10,947 location records. Its current thesis-directed view has 3,554 claims and 3,690 relations. Its all-active graph has 7,286 endpoint claims and 7,098 relations; 3,847 relations are explicitly archived, and two locations are structural bookkeeping. These scopes are different views, not independent samples. The active graph is reported as a DAG with 446 weak components; the active-plus-archive union contains cycles and is not established as a single historical state. [E4]

The package already supplies twelve machine-proposed reasoning occurrences from eleven distinct child claims, a source-only blank review packet, a draft annotation manual, split rules, structural checks and four small authored native scenarios. **There are zero human-adjudicated semantic reference items.** The four scenarios exercise a two-premise rule using explicit stipulations; they do not establish a faithful automatic interpretation of the poll claim from which they were motivated. The reported 97% agreement does not entail the scheme's universal “everybody” premise. [E4–E7]

This material can support structural preservation checks, importer robustness, annotation training, hypothesis formation and conditional examples. It cannot yet estimate scheme frequencies across Kialo, establish cross-topic interpretation accuracy, calibrate argument weights, verify external cited claims, reconstruct complete dialogue history, or establish performance across the owner's collection. Thousands of nodes in one discussion do not supply thousands of independent discussions. [E2–E7]

The current discussion, future versions of it, and confirmed copies remain development-exposed. Its twelve machine proposals may be used for training and transparent critique. They must not enter a sealed final evaluation set, and they do not count as twelve independently sampled calibration families.

## A concrete preflight correction: the old pilot is still pinned to the old library

The historical package was correctly preserved rather than silently rewritten. Its active use now requires a new derivative. The annotation builder hardcodes commit `fcccbba53337321566caf4951eaa8addf1037c0b` and engine `4.3.1+walton.5`; supplying a newer `--library` path does not update those strings. Its default JSON register is `closure-current-register.json`. That same-named file in the latest repository intentionally describes the earlier 143-record review; the complete current source acceptance covers 149 records. [E8–E10]

The read-only comparison in `catalog_readiness_audit.json` checked all fourteen scheme records referenced by the pilot. Their executable fields are unchanged in the inspected current library. One metadata citation changed. Four records still carry old open/hold statuses in the historical pilot catalog although the current ledger approves them: `expert_opinion`, `ignorance`, `positive_consequences` and `witness_testimony`. This is a provenance/catalog rebasing task, not a newly disputed source interpretation.

Before compiling current scenarios or presenting the pilot as current, create a versioned catalog and metadata derivative from the final accepted commit, its `walton.yml`, the live source ledger, the current source-acceptance register and current fingerprint registry. Parameterize and verify the code commit and engine/library versions instead of changing only paths. Revise the manual's obsolete 143-record and source-hold statements. Preserve the original packet and its historical hashes unchanged. A source record becoming approved does not automatically approve its proposed application to a Kialo occurrence.

This step should fail when the catalog, source status, commit, formulas or version metadata do not agree. It should retain the source text, occurrence keys, development-exposure status, proposed interpretations and zero-human-reference status. The existing validator can test preservation; it is not a semantic oracle or a human adjudicator.

### Preparation completed in this review

The bounded context refresher is now implemented in `../../pilot/refresh_pilot_context.py`, outside the framework repository. The included `../../pilot/context-d9a2585/` directory is a completed derivative of the inspected public commit. It reads the library, live ledger, current `independent-approvals.json` registry and engine version from immutable Git objects; a dirty worktree cannot be mislabeled as a clean commit. The script accepts explicit `--repo`, `--ref` and optional exact `--expect-commit`, so the parent reviewer can regenerate a separately named derivative after accepting any final integration commit.

The completed preflight checked the matching 149-record inventories and ledger/registry review stamps, bound the fourteen candidate records and twelve machine drafts to actual versions/hashes, refreshed the four outdated source stages and two obsolete hold statements, and preserved all semantic proposals and human-review fields. It copied the twelve-unit blank packet byte-for-byte and prepared the six-unit source-only subset. Human gold remains zero; annotation semantics and engine behavior remain unevaluated.

The preflight also exposed an overinclusive historical catalog extraction: it matched predicate names without arities and included an unused `relevant_differences/2` template. The actual candidate `analogy` record uses `relevant_differences/1`. The new catalog selects the used signatures, records exclusions, and refuses changed candidate formulas or changed used glosses. This narrows a metadata extraction; it does not alter an argumentation scheme or adjudicate a corpus interpretation.

Eleven deliberate mismatch controls were rejected, covering stale commit/version bindings, a registry mismatch, a missing current record, changed candidate formulas/glosses, source corruption, false human-gold/test promotion, a leaked machine proposal in a blank packet and catalog corruption. The output was independently reloaded and compared with a deterministic reconstruction. These are bounded metadata/provenance controls, not another framework test count. Approval fingerprints are copied as recorded attestations, not recomputed by this helper; the repository's source-approval gate remains a separate baseline check. Details and reproduction commands are in `../../pilot/README.md`.

## First milestone: a bounded, reviewable pilot

Retain the already agreed **60-dossier feasibility budget: 12 calibration, 12 development and 36 sealed evaluation dossiers**. The prior plan prefers one anchor dossier from each of sixty independent discussion/duplicate families; if that is infeasible it proposes at least thirty families with all items from a family kept together. These are planning allocations, not a demonstrated statistical sample-size requirement. The collection must be inventoried before claiming that this allocation is achievable. Keep several source families in the wider human-reference programme; sixty Kialo items alone would not establish performance on unrestricted prose or other domains. [E2–E3, E6]

Begin more narrowly with the previously proposed six-unit cost and ambiguity exercise. A source-only training packet is included in this review for `kialo2629-u01`, `u04`, `u07`, `u09`, `u11` and `u12`. These six occurrences contain five distinct child claims in one discussion family. The final two deliberately retain one claim's different placements. This is a purposive calibration exercise; no prevalence or generalization estimate should be calculated from it.

The detailed selection rationale belongs to the annotation lead, not the initial source-only packet: quantifier and missing-warrant distinctions, a possible explanation versus an established defeater, analogy versus ignorance, a practical recommendation versus a truth claim, and the same text's occurrence-dependent target. These are diagnostic contrasts already present in the historical packet, not accepted labels for annotators to copy.

Two trained people should independently annotate the six source-only units before seeing model proposals or each other's answers. Capture whether either person previously saw the proposals. An annotation lead can then measure time, context requests, disagreement, unresolved roles and unjustified premise insertion; an appropriately qualified third person adjudicates disagreements while preserving permissible alternatives. If contributors already know these examples, use them transparently as training and select unexposed cases for subsequent blind work. No second AI pass should be reported as an independent human annotation.

The immediate milestone is a workable manual and an honest cost/ambiguity report. Completing six exercises does not establish reliability of a 149-label taxonomy. Before using sealed evaluation items, freeze the revised manual, the relevant task definitions, the system/prompt versions and the scoring procedure. Any later correction to the manual or references creates a new benchmark version and a visible contamination/correction record.

## Four tracks, four different references

| Track | First bounded deliverable | Independent reference | What a pass would establish |
|---|---|---|---|
| Structural ingestion | Versioned capture manifest and loss report for the supplied derivative, followed by additional eligible captures | Immutable captured fields, explicit scope definitions and separately reviewed preservation checks | Correct preservation/refusal for the declared structural profile |
| Formal implementation | Small native CAES/CHR cases and a bounded independent evaluator; keep appropriate Dung controls separate | Frozen semantic definitions, hand-derived counterexamples and separately implemented finite reference procedures | Correct answers on the named finite supported domain |
| Human interpretation | Six-unit calibration, then the 12/12/36 feasibility programme | Two blind initial human annotations, qualified adjudication and retained alternatives | Feasibility and sampled fidelity of source-to-model interpretation |
| Capacity | Small correctness-qualified workload ladders after inputs and answers are eligible | The same fixed expected answers plus recorded resource/process outcomes | Measured behavior over the stated workload and machine, not universal capacity |

Kialo should lead the structural and human-interpretation tracks. It can supply the shapes of capacity workloads, but a raw graph is not automatically a meaningful CAES workload. It can motivate formal scenarios, but the expected answers must be derived from the explicitly authored scenario, not from votes, pro/con labels or the engine's own output. ICCMA belongs to matching abstract-formal tasks after the protocol adapter is tested; it does not supply CAES defaults, critical-question oracles or Kialo scheme annotations. [E1–E3]

### Formal reference work should start in parallel

Choose a deliberately small, declared native profile and derive its expected behavior from a written contract. The first cases should cover linked premises, repeated-premise/set behavior, supported exceptions, target-specific undercutting, rebuttal/issues, defaults and competing stipulations, and typed failures. Keep supported proof standards and weighing functions explicit. If the first reference evaluator covers only a narrower subset, say so and retain separate cases for unsupported parts.

The independent evaluator must not call Carneades or GoCHR for its expected labels, copy their generated argument identifiers, or share the key inference/evaluation routines. The recent circular AATS comparison demonstrates why this matters. A hand trace can serve a small case, but it must be reviewed independently of the implementation's output. Deliberately wrong answers, missing arguments and extra arguments should make the checker fail.

Use metamorphic checks only with stated preconditions. Bijective renaming and input ordering should preserve the relevant logical result; an unrelated component requires an actually disjoint interaction scope. Adding a fact is not a general monotonicity property when defaults, exceptions or competing issues are present. Counting a metamorphic relation as independent evidence requires a justified relation, not merely agreement between two outputs.

Expand the four existing Kialo-motivated scenarios before making a broad application claim. At minimum, add separately specified families for linked versus convergent premises, wrong-instance binding, target-specific undercutting, rebuttal, questioned default versus established exception, and an alternative admissible interpretation. A positive and discriminating negative case for each family is a coverage checklist, not a statistical minimum or a claim that these are already implemented.

### Do not make an automatic Kialo-to-Walton converter the first milestone

The first adapter should preserve structure and expose unsupported or missing fields. The next compiler should consume reviewed annotation/scenario records with explicit assumptions and formal policies. It should not infer that Pro means a valid support argument, Con means literal negation, siblings are conjunctive premises, or a poll/impact score is an argument weight.

Preserve at least captured text, interpretation, formal scenario and conditional output as separate objects. For every premise, retain whether it is explicit, a justified paraphrase, an implicit reconstruction, externally evidenced, unknown or contradicted. Preserve no-match, insufficient-context and multiple-defensible-analysis outcomes. The absence of a modeled objection is not evidence that a default is warranted.

For critical questions, distinguish applicability, whether the question was raised, whether an answer exists, the evidence supporting that answer and the resulting modeled role. Asking a question does not prove an exception; proving an exception does not by itself settle every burden or proof-standard decision. [E3, E6]

## Split and sampling policy

Inventory captures before selection. Record archive and derivative hashes, discussion ID, timestamp/version lineage, schema, language, topic, graph scopes, completeness and rights evidence. “Thousands captured” remains an owner-reported collection size until unique files and discussion families are counted.

Split by discussion/version lineage and confirmed copied-argument or substantive duplicate families **before** sampling dossiers. Do not randomly divide edges of a single discussion across development and evaluation. Reused claim text retains every occurrence; split grouping and semantic identity are different operations. Preserve cross-discussion references and review large connected duplicate groups so a generic phrase or common website domain does not accidentally join the entire collection.

Use a natural-distribution sample with a declared sampling frame and inclusion probabilities when making population estimates. Keep a separate purposive challenge sample for rare or failure-prone constructions. Choose observable strata such as topic, size, depth, reused claims, pro/con ratio and completeness without deriving the answer key from system predictions. Report challenge results separately from prevalence estimates. Never force all 149 scheme records into Kialo examples to claim natural coverage.

The present discussion family is development-exposed in its entirety. A later snapshot is not automatically an unseen test. A later-time evaluation must either exclude earlier versions and copies from development, or explicitly define itself as an updating task with that prior access. Public text may also have unknown model pretraining exposure; a fresh project split does not establish that a language model has never encountered the discussion.

For uncertainty, the independent clusters are discussion/duplicate families, not graph edges. Report both family and occurrence counts. One discussion cannot produce a defensible across-topic confidence interval regardless of how many local nodes it contains. With a small feasibility sample, report raw counts and uncertainty without ranking systems on tiny differences. [E2–E6]

## Acceptance criteria to freeze before results are seen

### Engineering gates

For supported structural fields, require zero unexplained loss, reversal, incorrect merge or silent acceptance of unsupported semantics. Intentional minimization or excluded fields belong in the declared profile and loss report. Safe refusal is a successful negative-input test, but it is not successful interoperability of that source format.

For each supported deterministic finite scenario, require exact agreement with the independently frozen expected labels and argument/issue/assumption obligations. Require zero incorrect answers and zero unresolved oracle disagreements in the suite supporting the specific release claim. A disputed reference cannot silently disappear from the denominator: adjudicate it, narrow the claim explicitly, or leave that claim unpassed.

Keep process outcome, semantic answer, answer correctness and task eligibility separate. A timeout or resource refusal is not a logical “no.” An unsupported task is excluded from supported-task accuracy but remains visible in capability coverage. A task that needs an answer by a deadline still fails its availability requirement when the answer is absent.

The zero-error requirement is an engineering gate for a fixed finite regression/reference suite. It is not proof of arbitrary-input correctness, and it is not a sensible claim that a natural-language system must achieve literal zero ambiguity on all real arguments.

### Human-reference and interpretation gates

Before a reference is designated adjudicated, require exact source binding, recorded independent initial judgments, a documented resolution or permitted alternatives, and visibility of every reconstructed premise and formal policy. Unresolved items retain their status and count; they are not relabelled correct or made unique to simplify scoring.

Report pre-adjudication agreement using measures appropriate to each task, together with denominators, per-class counts, confusion matrices, abstention/coverage and serious error categories such as invented premises, changed modality or wrong attack targets. Agreement is useful evidence about a procedure; two people agreeing does not itself prove the interpretation is true.

Do not invent one universal kappa/F1 threshold or claim that sixty dossiers provide statistical certification. Use calibration/development to estimate time, ambiguity, clustering and class prevalence. Then preregister use-specific quality and serious-error limits before opening the sealed evaluation. A desired confidence interval or minimum detectable difference should determine later expansion using observed clustering and prevalence. [E2–E6]

### Capacity pilot

Freeze hardware/software, warm/cold policy, concurrency, time/memory limits, seeds, repetitions and outcome policy before execution. An initial containment budget such as 30 seconds and 2 GiB per bounded case is a runner parameter, not a product latency requirement; adjust the proposed runner configuration before execution if the chosen platform cannot enforce it. Retain timeout, crash, refusal and incorrect-answer records.

A five-repeat feasibility measurement can report each observation, median and range; it should not be sold as a precise tail-latency estimate. Choose later repetition counts and workload sizes using the pilot variability and the actual product requirement. Measure correctness-qualified cases and report failure/coverage separately, rather than dropping failures and presenting only successful timing results. Until a workload and latency/memory objective is adopted, measurements characterize an envelope; they do not certify capacity.

## What can start now, and what needs additional inputs

| Can begin with the supplied material | Needs additional people, data or a later decision |
|---|---|
| Protocol registration, current catalog/provenance rebasing, structural adapter checks and blank training packets | Independent human annotations and a qualified adjudicator |
| Specification and development of bounded independent formal oracles | An actual inventory of the remaining owner-held capture collection |
| Explicit native scenarios and targeted adverse cases, with independent expected answers | Unexposed discussion families for development/evaluation separation |
| Resource-runner setup and exploratory workload design | A product workload/SLO before certifying practical capacity |
| Preparation of the 12/12/36 feasibility study | Intended-use and downstream-sharing evidence before the corresponding public corpus release |

These are concrete work packages, not reasons for another broad source-review cycle. The current readiness decision should be attached to the parent's exact accepted integration commit. It should not depend on a public tag or imply that web hosting, general import fidelity, full collection coverage or human interpretation is already validated.

## Evidence sources and work performed

This assignment read the current repository and the preserved prior package, compared the fourteen catalog records with the current library and source ledger, prepared a proposed pilot outline plus a six-unit source-only training packet, and implemented/executed the bounded context refresher described above. It did not run the engine, annotate reference labels, inspect additional captures, contact people, or perform new web/corpus research. Source-file hashes are in `SOURCE_MAP.json`; the read-only catalog comparison is in `catalog_readiness_audit.json`. The completed metadata derivative and mismatch-control results retain their own exact commit/input bindings.

| Key | File |
|---|---|
| E1 | Current repository `PLAN.md`, especially D24–D27 and Remaining sequence D–E |
| E2 | Prior package `plan/DECISIONS_AND_GATES.md` |
| E3 | Prior package `kialo/UPDATED_PLAN.md` |
| E4 | Prior package `kialo/DATASET_CARD.md` |
| E5 | Prior package `kialo/annotations/README.md`, `pilot_manifest.json`, `SPLIT_MANIFEST.json` |
| E6 | Prior package `kialo/annotations/ANNOTATION_GUIDELINES.md`, `SAMPLING_AND_SPLITS.md` |
| E7 | Prior package `kialo/cases/case_manifest.json` |
| E8 | Prior package `kialo/annotations/build_pilot.py`, `validate_pilot.py`, `library_candidate_catalog.json` |
| E9 | Current repository `docs/reviews/closure-current-register.json`, `final-source-acceptance.json`, `tools/vv/source_ledger.yml` |
| E10 | Current repository `walton.yml`; current scoped import contract `docs/import-profiles.md` |
| E11 | Current repository `docs/reviews/independent-approvals.json`; bounded preparation artifacts in `../../pilot/` |
