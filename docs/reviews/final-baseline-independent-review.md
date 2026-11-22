# Independent integration review before benchmarking

Review date: 2026-11-21 America/Chicago; supporting runs and some attestations
are dated 2026-11-22 UTC. Public starting point:
`98a2335d2c9909cf597ae85a9f359f20d092ae6c` (PR21).

## Decision and scope

Accept the corrected **library 0.4.1 / engine 4.3.1+walton.6 candidate as the
framework baseline for the next validation effort**, conditional on the exact
commit's clean two-toolchain gates in the accompanying archive. Retain all
149 library records within their declared source interpretations. Do not
identify the original PR21 tree, the old R1 pin, or a passing source-stage
counter as the corrected baseline or a complete stable 1.0.0 release.

This review closes specified source, interpretation, input-boundary and
integration decisions. It does not claim formal proof of arbitrary inputs,
original-author endorsement, human annotation gold, empirical argument
adequacy, general VAF semantics, full format conformance or benchmark-scale
capacity. Those require the separate validation programme. No tag, remote
release, upstream submission or author correspondence is asserted here.

## What was independently established at the starting point

The current GitHub source was retrieved and isolated from the earlier working
checkout. GitHub's API reported successful Go 1.24.13 and Go 1.27.4 jobs and a
successful result-comparison job for this exact PR21 commit. A fresh detached
checkout independently reproduced **10,428 / 10,428 registered checks** on
Go 1.27.4, with no known or unexpected failures. The accompanying log records
the full repository/Python/harness run; this was not an independent rerun of
every historical source review or every prospective benchmark.

The 149-record census, 120 book-origin / 24 inherited / 5 publication split,
and validator's zero errors / ten warnings are correct. The warnings are
five unused inherited templates and five dependency-cycle groups. The source
ledger's nine final-closure approvals explicitly admitted self-authorship.
The old checker correctly verified matching hashes but did not satisfy D14
merely by passing those rows. One additional D20 source-contract change also
needed a non-author decision.

An initial source-archive attempt encountered local Go VCS-status errors and
was stopped; it is not presented as a successful acceptance run. The clean
detached checkout supplies the reproduced baseline result. Both attempts are
retained in the evidence rather than erasing the unsuccessful attempt.

## Independent source closure

The separate source reviewer inspected the compendium and the supplied 1995,
2000, 2002 and 2015 works at the relevant passages. The five publication
records were inspected separately against the 2007 paper, including rendered
formal definitions and figures. `final-source-independent-review.md`,
`final-atkinson-independent-review.md` and `final-source-acceptance.json`
record exact scope and fingerprints.

The principal source decisions are supported:

* `practical_reasoning1` is the 2015 configured value-based Carneades scheme,
  not a claim to implement every 2007 question. Its three assumptions and
  four exceptions agree with that source configuration.
* `practical_reasoning2` correctly distinguishes possibility of the goal from
  possibility of the action and restores the broader instrumental side-effect
  question. Both changes remain manual semantic migrations.
* `exeption_to_rule` is an explicitly declared reading of Case 80a on the
  1995 p. 148, separate from the displayed rule-revision conclusion of the
  precedent scheme. Preserve its historical id and provenance; do not infer
  author approval of today's executable semantics from a private letter's
  existence or the 2019 README.
* The omission branch of 33.3 is defensible from Scare Tactics pp. 139–140.
  Its local recommendation and omission-specific likelihood must not become
  a global equivalence between deontic operators.
* The inherited quantity/feasibility `sunk_costs` reading belongs under
  compendium entry 24's waste interpretation; entry 25's printed precommitment
  form stays separate. The original D20 reviewer authored that reclassification;
  a fresh independent judgment now closes that narrower gap.

Three authoritative source locators are corrected: action 33.3 to Scare
Tactics p. 140, omission 33.3 to pp. 139–140, and disjunctive threat 35.1 to
p. 142. The compendium's shared original references remain in provenance
notes. This is citation disambiguation, not a claimed publisher-confirmed
erratum. The compendium does include a conditional threat at entry 35; the
claim that it included none was incorrect. Smart Car prose now identifies
the environmental bridge E as the premise supported by V1, rather than the
factual fuel-consumption premise B.

D20 remains the owner's source-of-record decision. Unseen expert-opinion,
slippery-slope and witness originals remain unverified research; their absence
does not silently invalidate a clearly scoped compendium product. The scope
does not claim that those originals were compared or must agree.

## Approval integrity and review independence

The candidate replaces the nine disclosed self-stamps and the sunk-cost
source-contract self-stamp with documented non-author acceptance. Two further
citation-only refreshes retain their executable semantics. Eight earlier
repairs had genuine later non-author acceptance but still named their original
review/implementation team; the current attribution identifies the documented
second reviewer. Other judgments are retained historical acceptance, not
misrepresented as 149 newly repeated line-by-line reviews.

Registry v2 requires a Boolean `independent_of_author` and a nonempty
`independence_basis`, matched in ledger and registry alongside reviewer, date
and decision. It still cannot authenticate an identity or prove a statement
true. The metadata helper restores missing metadata from valid accepted
stamps, refuses explicit withdrawals or attribution conflicts, and preserves
the existing reviewer. It cannot manufacture approval by refreshing a hash.

The integrating reviewer authored this checker change; the separate source
reviewer challenged it with adversarial synthetic records. The first revision
could overwrite a withdrawal, and a later probe exposed integer/Boolean
equality in the registry. Both were repaired and independently rechecked.
The final controls reject these cases before writing and retain file bytes.
This is why a technical non-author review is recorded separately from the
implementer's successful unit tests.

The first complete integration run then found an error in this review's own
metadata update: the Go harness's strict ledger reader did not yet declare
the two new fields. That candidate was refused. The reader now recognizes
both fields and also requires a literal Boolean true and a nonempty string
basis for approved rows. Its ten serialized-input controls include the valid
case, missing/self/quoted/integer/null attestations, invalid bases and an
unknown field. The unsuccessful run is preserved and is superseded only by
clean acceptance of the subsequent exact commit. The separate reviewer
accepted the fix and ran its focused controls on both Go versions
(`final-go-ledger-cross-review.json`).

## AATS/AS2 implementation acceptance

The Atkinson reviewer authored the adapter repairs. The integrating reviewer
inspected their state-pair, state-identity, joint-tuple and fact-join logic;
the separate source reviewer then challenged final input boundaries and ran
the authored actual-engine regressions on both pinned Go versions. Their
cross-review is retained in the archive.

The decisive pre-fix counterexample used two joint actions with the same
source and target states but different declared effects. The adapter expected
one promotion argument; the emitted state-pair facts caused the engine to
build an additional argument. Requiring identical effects for that state
pair repairs the projection while respecting the declared state-pair model.

The paper does not require distinct proposition sets for distinct states, so
the unsupported injective-π restriction is removed. Ambiguous proposition-list
claims require explicit state ids. A legal but unnamed joint-action tuple is
distinguished from an impossible tuple: its missing transition is a CQ2
problem, not evidence that no such joint action exists. Duplicate YAML keys,
malformed collection shapes and invalid references must fail diagnostically.
Audience specifications must not turn an unknown value into an empty ranking
and an unqualified all-actions tie.

The original independence claim for the randomized comparison was too weak:
it consumed the adapter's own predicted argument identifiers. The revised
oracle joins emitted facts independently, while additional actual-engine
controls use hand-authored expected arguments. These methods supplement each
other; neither is a formal proof for every possible model.

The Figure 5 discrepancy is consequential. Interpreting its single minus-P
mark as one penalty chooses the dog for the hurried farmer, whereas the
text's two harmful goals choose the chicken. The candidate intentionally
retains the textual reconstruction and tests the sensitivity. The recorded
choice expectations are source-based reconstructions, not 32 independent
empirical observations.

AS1 and AS2 remain distinct. The comparison command needs a separately
authored AATS model; it cannot infer missing agents, joint actions, outcomes,
value effects and assurances from AS1 premises. Its worst-case goal counts,
lexicographic audience ordering and tied-choice policy are explicit project
choices. No general VAF evaluator or multi-step planner is implemented.

## Engine and import-boundary acceptance

The engine reviewer independently reproduced the old counterexamples at
PR21, then authored the repairs. The integrating reviewer inspected the
changed parsers, admission checks, topology construction, Dung equality and
tests. The final complete native and strict gates, including race tests,
bind those changes together at the exact candidate commit.

Native YAML now rejects unsupported object premises and semantic fields,
duplicate keys and invalid/null records. The minimal captured counterexample
can no longer become an unconditional argument merely by renaming its scheme.
The stricter decoder also exposed an ignored `labels:` key in the bundled
abduction example; it now uses actual `tests:` assertions.

AIF assumption selection and both issue backlinks are corrected. Unsupported
nodes/edges, duplicate keys and undeclared endpoints are refused. Its limited
projection is explicit: RA is linked support with source-scheme provenance,
binary CA becomes symmetric incompatibility, and unargued non-issue I nodes
are stipulated. This is not general AIFdb or directed-conflict equivalence.

CAF, AGXML and LKIF require explicit experimental CLI opt-in. Targeted
repairs address CAF defaults and forward undercutters, AGXML assumptions and
backlinks, and LKIF exception links and multiple-graph loss. Their remaining
general semantics are outside conversion-fidelity acceptance. The twelve
incompatible Kialo exports are explicitly refused; four previously authored
conditional native examples still pass. Refusal closes the silent-loss
defect without pretending to implement Kialo annotation or translation.

Dung's extension-family comparison is now symmetric, its odd-cycle fixture
is a real three-node cycle, and TGF rejects labelled support/malformed edges.
The exhaustive enumeration architecture and missing modern ICCMA protocol
remain next-stage capability boundaries.

## Migration, packaging and next action

The executable formulas and identifiers of all 149 records are retained in
this revision. Corrected citations change three record fingerprints. Generated
category copies and four example prefixes are refreshed. Source-backed manual
migrations preserve shared action, side-effect and case predicates until the
application supplies the correct new interpretation; the review's mixed-use
probes verify that no global rename is performed.

Original `carneades-4.zip`, `GoCHR.zip` and `sources.lock` remain provenance
inputs. They do not contain this repository's repairs. The archive supplies
the actual candidate, patch/bundle, patched engine assets and clean acceptance
evidence. Its final manifest is the authority for the exact tested commit
and file hashes; this report does not infer publication from local preparation.

After clean acceptance, proceed with the separately planned verification,
validation and benchmarking programme against this baseline. Human annotation,
corpus rights/eligibility, semantic adapters, solver capacity and external
oracles remain explicit next-stage obligations. They do not require another
general opinion round to settle the source and profile decisions closed here.
