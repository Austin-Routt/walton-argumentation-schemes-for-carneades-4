# Validation entry protocol

Status: adopted next-stage protocol; preparation artifacts are supplied, but
no human annotation study or new benchmark campaign is claimed complete.
This protocol applies after the exact library 0.4.1 / engine walton.8
candidate passes the acceptance gates in `../PLAN.md`. A full immutable
commit and input hashes are sufficient; a published tag is optional.

## Objectives and first milestone

Use Kialo as the first application corpus for structural ingestion and human
interpretation. Develop independent finite CAES/CHR references in parallel.
The first milestone is a reliable annotation procedure, explicit supported
input contract and reviewed reference scenarios. It is not a single accuracy
score for the entire argumentation framework.

| Question | Reference needed | What an initial result can support |
|---|---|---|
| Did ingestion preserve the captured graph? | Immutable captured fields, explicit view policy, separately reviewed preservation checks | Structural fidelity for that input profile |
| Does the engine implement its formal contract? | Independent finite evaluator and hand-derived, reviewed counterexamples | Correctness on the stated bounded domain |
| Does the formalization preserve an argument's meaning? | Independent human source annotation, adjudication and retained alternatives | Sampled interpretation fidelity and feasibility |
| Does the system run within practical limits? | Correctness-qualified workloads, fixed resource limits and run records | A measured operating envelope on the stated hardware |

Kialo data can contribute to several rows, but each requires its own
reference. A result from one row cannot substitute for another.

## P0: bind the experiment before measurement

Record the accepted implementation commit/tree; library, engine, adapter and
runner versions/hashes; data snapshot and view; task; supported profile;
independent-reference version; expected outcome definition; sampling/split
policy; and resource limits. State the unit of analysis and intended claim.
Do not use a moving `master` branch or let a tag hide an unreviewed change.

Separate invalid, unsupported, inconsistent, limited, cancelled and internal
error outcomes from a semantic `out` or `undecided` result. Retain every
failure/refusal and its eligibility reason. A safely refused conversion is
useful robustness evidence, not a successful interoperability result.

The source-stage gate covers source interpretations. It does not validate
the truth of an input statement, adequacy of evidence, annotation decisions,
external predicates, corpus representativeness or engine capacity.

## P1: use a current, explicitly unadjudicated pilot packet

The previous package contains one minimized discussion derivative, twelve
machine-proposed reasoning occurrences, source-only blank packets and four
authored conditional examples. Preserve those historical files and hashes.
Their old 143-record catalog and `fcccbba`/walton.5 pins must not be presented
as current source status.

Create a new context derivative with the bundled tool. Run this block in a
clean dedicated checkout whose `HEAD` is the same `FULL_ACCEPTED_COMMIT`:
the source-approval checker reads the worktree, while the refresher reads
immutable Git objects. The final acceptance record must identify that same
commit and show the source gate passed there.

```sh
python3 tools/check_source_approvals.py --repo . --require-stable
python3 tools/validation/refresh_kialo_context.py \
  --repo . --ref FULL_ACCEPTED_COMMIT --expect-commit FULL_ACCEPTED_COMMIT \
  --source-annotations /path/to/historical/kialo/annotations \
  --graph /path/to/historical/kialo/data/graph.json \
  --output /path/to/new-current-context --self-test
```

Replace both commit placeholders with the full accepted SHA. To verify a
saved derivative, use the same arguments with `--check` in place of
`--output`. The tool reads committed Git objects even if its checkout is
dirty, records the actual pin/versions/hashes, and refuses a mismatched
`--expect-commit`. It never overwrites a nonempty output directory or the
historical packet.

This bounded tool checks the supplied twelve-draft/fourteen-candidate
profile, the current 149-record inventory, matching ledger/registry review
stamps, unchanged candidate formulas/glosses, source text and development
exposure. The repository source gate separately recomputes approval
fingerprints. The tool's deterministic controls do not create human labels,
reassess the machine interpretations or execute the engine. A changed
candidate meaning or inventory requires a deliberate protocol revision.

The supplied six-occurrence packet is a purposive cost/ambiguity exercise
from the exposed discussion: `u01`, `u04`, `u07`, `u09`, `u11` and `u12`.
It contains five distinct child claims in one family, deliberately including
two placements of one claim. It has no population-level sampling meaning.

The copied `reviews/validation-entry-independent-review.md` and
`validation-pilot-proposal.json` preserve their original evidence and
package-relative artifact pointers. Those pointers refer to the contributor
archive's historical Kialo material and pilot preparation, and do not all
resolve in a standalone repository checkout. Use the archive's `START_HERE.md`
and `REPRODUCE.md` for its current directory map. The proposal's null
`experiment_commit` is intentional; the final archive supplies the separately
pin-bound context and preparation record.

## P2: start three bounded tracks in parallel

### Structural ingestion

Preserve source discussion, claim and occurrence identifiers; exact text and
version; parent/target identity; platform relation; link status; external
cross-links; and the selected view's archived/missing-state policy. Preserve
export navigation separately from reasoning direction. A canonical claim
can occur in multiple argument contexts, and those occurrences must remain
distinct even when they share text.

Require zero unexplained supported-field losses, polarity reversals or
occurrence merges in the finite accepted fixtures. Refuse unsupported
structures explicitly and produce a loss/coverage report. The captured
active-plus-archive union is not automatically one historical debate state.
Graph preservation does not prove any relation is a valid inference.

Do not feed the old Kialo CAF/AGXML/LKIF/TGF exports directly into Carneades
as though their fidelity had been established. The present importer boundary
refuses known incompatible shapes. The first adapter preserves and exposes
structure. A later compiler consumes reviewed annotation/scenario records
with explicit assumptions and policies. Arbitrary external IDs also need a
reversible, collision-checked mapping into native ground statement keys.

### Independent formal references

Specify a small native profile before writing the independent evaluator.
Cover linked premises, repeated-premise set behavior, missing prerequisites,
wrong-instance binding, supported exceptions and targeted undercutters,
rebuttal/issues, questioned defaults versus established exceptions,
competing stipulations, and typed failures. State the proof standard and
weighing function in each reference case.

Expected answers must not call Carneades or GoCHR, reuse their inference
routines, or copy the adapter's predicted argument identifiers. Preserve
hand traces and independent review where used. Deliberately wrong, missing
and extra answers must fail the comparator. Require zero incorrect answers
and zero unresolved reference disagreements for the named finite suite;
this is not a theorem about every possible input.

Use bijective renaming or permutation checks with explicit preconditions.
An unrelated-component check needs a truly disjoint interaction scope.
Adding facts is not a general monotonicity law for defaults and exceptions.
Keep Dung abstract semantics and ICCMA protocol/reference checks separate
from native structured CAES reasoning and critical-question behavior.

### Human annotation calibration

Two trained people independently annotate source-only packets before seeing
machine proposals, engine results or one another's answers. Record prior
exposure. A qualified third person adjudicates disagreements and retains
defensible alternatives and unresolved cases. If contributors already know
the examples, label this activity training and use unexposed examples for
subsequent blind work. Another AI pass is not human annotation.

Train annotators with the complete declared taxonomy and neutral manual.
The fourteen-record machine-candidate catalog is a developer aid; showing
its occurrence-specific candidates during blind annotation would disclose
the model's proposals and restrict the possible answers.

Capture occurrence and target; exact proposition spans; polarity, modality,
quantifier and participant identities; linked versus convergent support;
explicit versus reconstructed premises; scheme candidates including
`no applicable scheme` or `ambiguous`; CQ/defeater target and evidence; and
the assumptions required for any executable scenario. Do not force every
Pro/Con edge into a Walton record. Do not infer a joint premise set from
siblings or turn impact/votes into evidence weights.

Measure annotation time, requested context, unresolved roles, disagreement,
unsupported premise insertion and alternative interpretations before revising
the manual. Report initial inter-annotator agreement separately from
adjudicated reference labels. Six units cannot establish reliability across
149 scheme labels. The first output is a usable procedure and cost estimate.

## P3: inventory, sample and freeze leakage-aware splits

Inspect the additional owner-held captures before making collection claims.
Track discussion/version lineage, canonical claims, substantive copies,
cross-linked material and topic/domain coverage. Assign confirmed duplicate
or copy families as a unit; do not randomly split edges of the same debate
between training and evaluation. A shared subject or citation alone is not
proof that two discussions are duplicates; audit overlap and declare the
generalization claim separately.

The currently supplied discussion, its versions and confirmed copies remain
development-exposed. The twelve old proposals and six training units are
excluded from sealed accuracy evaluation. Unknown pretraining exposure is
a limitation to disclose, not something this procedure can rule out.

Retain the existing feasibility budget of **60 new dossiers: 12 calibration,
12 development and 36 sealed evaluation**. Prefer one anchor dossier per
independent discussion/copy family; if fewer families exist, record the
clustering and revise the estimate of independent evidence. This is a
budget, not a power calculation or a claim that sixty units can validate all
scheme types. The exposed packet does not fill the sixty independent slots.

Maintain a natural-distribution sample with recorded inclusion rules and a
separate purposive challenge sample. Do not combine their accuracy into an
unweighted population estimate. Use calibration/development evidence to set
task/use-specific thresholds, scoring and any expansion targets before
opening sealed cases. Preserve rare-class counts, abstention/coverage,
accepted alternatives, unresolved rate and reconstruction errors; a single
global F1 or agreement coefficient will conceal important failures.

## P4: measure capacity after correctness and eligibility

Use supported, correctly specified workload ladders with independent expected
answers. Record the full process outcome, answer correctness, input and
reference hashes, toolchain, seed, resource cap, elapsed time, peak memory,
hardware and repetition. Do not drop failed inputs and report only successful
timings. Graph size alone does not describe rule-join, term-depth, issue or
interaction difficulty.

The proposed first bounded measurements use one worker, a 30-second wall cap,
2,048 MiB memory cap and five repetitions per case. These are containment
parameters to freeze in the runner, not a performance service objective.
Report all observations, median and range; five repeats do not support a
precise tail-latency claim. Choose later workloads and repetitions from the
pilot variability and actual intended use before capacity certification.

## P5: release evidence and decisions

Stable 1.0.0 needs results against the preregistered task-specific criteria,
reviewed semantic migrations where required, reproducible references and
datasets, and fresh complete gates on the exact release commit. Broad AIFdb
fidelity, ICCMA performance, natural-language interpretation and hosted
service operation are separate claims; include only those actually studied.

Preparation can proceed now. Human validation still needs people and
adjudication; a cross-discussion study needs additional eligible captures;
capacity certification needs a real workload objective. These are explicit
next work packages, not undisclosed failures of source convergence. D22
remains unchanged: no upstream patches or author correspondence are sent.
