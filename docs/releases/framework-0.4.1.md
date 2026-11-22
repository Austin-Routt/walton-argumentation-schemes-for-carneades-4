# Framework baseline preview: library 0.4.1 / engine 4.3.1+walton.8

Engine walton.7 added the web import profile at public PR22 `d9a2585`.
Engine walton.8 completes a subsequent independent review of the actual
routes: both now validate authored arguments, propagate inference failures
and request cancellation, and return non-success HTTP statuses without
exporting a failed graph. Eleven live D20 reviewer attributions are also
corrected without changing a library formula or semantic fingerprint.
The latest review archive contains separate acceptance evidence for this
exact candidate; the previous walton.6 archive and public walton.7 CI remain
historical evidence. See `../reviews/post-integration-readiness-review.md`.

This candidate follows PR21 `98a2335d2c9909cf597ae85a9f359f20d092ae6c` and
incorporates the independent pre-benchmark closure review. Publish only the
exact commit identified by the clean two-toolchain acceptance evidence in the
review archive. This document does not assert that a remote tag, pull request
or GitHub release has been created.

Independently review the changes and finish those exact-commit gates before
tagging this preview. A full commit hash is sufficient for starting the
protocol and pilot-preparation work; a remote tag is not a prerequisite.

## Why this revision is needed

PR21's 149 recorded approvals included nine acknowledged self-reviews, and a
tenth source-contract change needed non-author acceptance. Its 10,428 passing
registered checks were reproducible but did not cover a real AATS effect
aliasing counterexample or the known importer losses. Its claim that every
figure/text discrepancy left the chosen action unchanged was incorrect.

## Changes

* Retain all 149 source interpretations within declared scope: 120 book-origin,
  24 inherited and five publication records. Renew the nine self-stamps and
  sunk-cost classification through non-author source review; preserve honest
  historical attribution, correct three source locators and example prose.
* Require explicit non-author attestations in approval-registry v2. Refuse
  stale fingerprints, malformed attestations and attempts to overwrite an
  explicit withdrawal through metadata reconciliation.
* Repair AATS state-pair effect consistency, duplicate-input loss and
  problem-formulation diagnostics. Permit distinct states with identical
  proposition sets. Test the Figure 5 choice sensitivity and the actual
  emitted-fact/engine boundary independently.
* Declare the external AATS comparison policy and the extra model required for
  AS1 comparisons. This is not a general VAF solver or an automatic model
  reconstruction system.
* Reject lossy native YAML/AIF input shapes, preserve AIF topology and its
  explicit projection policy, and require opt-in for experimental legacy XML
  converters. Correct the identified CAF defaults, AGXML topology and LKIF
  exception-link faults. Restrict Dung TGF to unlabelled attacks and repair its
  test fixture/equality weaknesses.
* Align both CAES HTTP routes with authored-argument validation and typed
  inference failures, including request cancellation. Add actual-handler
  tests and independently reviewed live-server controls. This does not
  certify hosting, the Dung route or output-stream failure recovery.
* Rebase Kialo pilot context from immutable Git objects and the live source
  registry. Preserve original drafts, source-only packets and zero-human-gold
  status; use Kialo as the first application pilot alongside formal references.

## Compatibility and source scope

No scheme id or executable library formula is changed in this revision. The
three corrected metadata citations have renewed fingerprints. Category files
and worked-example prefixes are regenerated. The earlier semantic migrations
remain manual where evidence or role interpretation is required; specifically,
do not globally rename `possible/1`, `side_effects/2` or case-difference
predicates that still have valid uses in other records.

Inputs previously accepted through silent field loss can now be rejected.
Legacy CAF/AGXML/LKIF CLI and server use needs `-experimental-import`. AIF import is the
specific projection in `../import-profiles.md`, not a generic format-fidelity
claim. AATS models whose joint actions disagree about effects for the same
state pair must be corrected before generating a graph.

The HTTP routes now return non-2xx errors for invalid or failed evaluations;
the old form route sometimes returned an error page or failed graph as HTTP
200, and the raw route previously evaluated invalid authored arguments. This
is an intentional compatibility correction. Restricted AIF I-node IDs must
also satisfy native ground-formula validation; arbitrary numeric identifiers
need a future explicit mapping, as documented in the import profile.

The compendium remains the source of record under D20. The newly supplied
originals have been compared for the relevant passages; inaccessible older
originals remain unverified research, not undisclosed release conditions.
Source acceptance is technical, AI-assisted non-author review within scope;
it is not original-author or institutional certification.

## Acceptance and next work

The archive records the exact frozen candidate, both pinned Go versions
(1.24.13 and 1.27.4), native build/vet/unit/race gates, strict registered
inventory and check-by-check comparison. Counts alone do not establish
semantic or empirical validity. Source gate L checks all 149 live records;
it is not the complete stable-release gate.

Use this preview as the baseline for the separate verification, validation and
benchmarking programme. Stable 1.0.0 waits for its agreed evidence and a fresh
commit-matched release run. The original root ZIPs remain unchanged provenance
inputs; the patched source snapshot and separately named patched-engine assets
are the executable candidate. D22 remains in force: do not send the prepared
upstream patches or author memorandum.
