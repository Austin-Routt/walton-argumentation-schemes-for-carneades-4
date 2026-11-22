# Post-integration review and validation readiness

Review date: 21 November 2026, America/Chicago. Execution timestamps use UTC.
Public baseline inspected: PR22, `d9a25857277140a5d3f10aaf1dfa4ade0f15b91f`,
tree `fcea512fd99551b453ec283158ade395262d4b7c`.
Corrected candidate: library 0.4.1 / engine 4.3.1+walton.8.

## Decision

The source interpretations and the declared AS1/AS2 scope have converged.
Another general source-opinion round is not required. The integration report is
substantially accurate, but its conclusion that only one small implementation
gap remained was too strong. This review found two additional HTTP evaluation
defects, an incomplete propagation of the D20 attribution correction and stale
Kialo preparation metadata. Those are narrow corrections, not reasons to reopen
all 149 source interpretations.

Accept this corrected preview for validation preparation after the exact
candidate passes the required clean two-toolchain gates. Complete independent
acceptance before tagging. A full Git commit plus hashes is sufficient to pin
the experiments; a remote tag is optional bookkeeping. Begin with the protocol
and independent-reference stage, using Kialo as the first application-data
pilot alongside bounded formal engine work. Do not describe that readiness as
completed real-world validation or stable version 1.0.0.

The full contributor archive supplies the final commit, source tree, clean-run
evidence, complete before/after observations, historical material and the
current pilot context. This in-repository report avoids embedding its own
containing commit hash; the external acceptance record binds the whole tree.

## What was independently confirmed

The previous review archive has exactly 115,440,579 bytes and SHA-256
`825164f8cf16c659ffbc390107f258a35ea6a56c75ad59d746a40e8160ce6ec9`.
Its 2,618 listed payload files match their hashes, and ZIP CRC checking passes.
The manifest itself is an additional archive entry. The four previous review
commits are ancestors of the public integration unchanged. GitHub's retrieved
PR and check records identify PR22 as merged at `d9a2585`, with successful Go
1.24.13, Go 1.27.4 and cross-version comparison jobs on the integrated head
and public master. Those CI facts concern the public baseline, not the later
local corrections.

The public integration changed no executable scheme content or source
fingerprint relative to the accepted prior candidate. The original
`carneades-4.zip`, `GoCHR.zip` and `sources.lock` remain unchanged provenance
inputs. The earlier concrete source/AATS/import findings and their evidence
are preserved in the prior archive and repository reports. This round audits
their integration and the remaining delta; it does not pretend to be a fresh
page-by-page rereading of every book or a new exhaustive run of all external
format examples.

The new legacy-import flag is wired correctly. CAF, AGXML and LKIF are refused
by the server profile unless explicitly opted in. Valid tiny opted-in fixtures
reach their converters and evaluate, while malformed fixtures are still
refused. The strict native YAML behavior for unknown and duplicate keys is an
intentional compatibility change, not a reason to reinstate silent loss.

## Remaining defects and their corrections

### 1. The form route ignored inference failure

At `d9a2585`, `/carneades/eval` called `ag.Infer()` without checking its returned
error. It then labelled and exported the graph. Contradictory stipulations
`p` and `¬p` returned HTTP 200 with both labelled `in`; an unsupported guard
also returned an evaluated graph. The corresponding CLI correctly returned
an error and no result output.

The corrected common evaluation boundary performs authored-graph validation,
calls `InferContext(req.Context())`, and stops before labelling/export if
either stage fails. Contradictory context and unsupported guard cases return
HTTP 422 with no result graph. This is an HTTP wrapper repair; it does not
change the engine's existing inference or error semantics.

### 2. The raw route skipped authored-argument validation

At `d9a2585`, `/carneades/eval-arg-graph` checked inference errors but did not
call `validation.Validate`. A named scheme `p(X) -> r(X)` could be attached
to an authored argument with premise `q(a)`, parameter `a`, and conclusion
`r(a)`. The route returned HTTP 200 with `r(a)` labelled `in`. A premise
`p(b)` with parameter `a` likewise exposed a wrong-instance failure.

Both routes now share the validation/inference boundary. Invalid authored
arguments return HTTP 400, the typed invalid-input outcome and no result
graph. Positive native inference and valid declared import profiles retain
their parsed output. The form route also returns a non-success transport
status for profile/validation failures that previously produced HTTP 200
error pages. CLI exit codes remain unchanged.

The precise HTTP status and scope are in `../import-profiles.md`. The
permanent tests cover real handlers, both legacy-profile settings, malformed
input, positive inference and deterministic request-context cancellation.
Cancellation plumbing is not a claim about every network-disconnection
timing or output-stream failure.

### 3. The import gate runs before conversion, not before reading the body

The integration report and code comment said the legacy upload was refused
before it was read. `FormValue` parses the multipart request first. The gate
does prevent legacy conversion and inference; it is not an upload-size or
pre-read protection. The comment and active profile documentation now say
this directly. Upload handles and temporary multipart files are cleaned up
after the form request.

### 4. Eleven live D20 stamps still used the inaccurate attribution

The implementer correctly acknowledged three earlier contributions: the
compendium citations, `expert_opinion`'s credibility default and
`witness_testimony`'s domain-argument order. All preceded `49ad425`, the tree
examined by the historical non-author reviews. The later witness/negative
ethos repairs also have separate acceptance. The evidence exists, but the
eleven current ledger notes still said Claude wrote neither the record nor
its last change, and the matching live registry/acceptance basis did not
properly identify that evidence chain.

This review corrects the live attribution in all three places, preserves
the historical scope adjudication and names the present provenance reviewer
as retaining the already accepted interpretation. It does not misrepresent
eleven new complete source comparisons or turn older qualified judgments
into unconditional approval. A separate reviewer checked the actual
integration: all eleven changes match the approved map, the other 138 review
entries are unchanged, all 149 semantic fingerprints and example lists are
unchanged, and the source-stage gate passes. See
`integration-provenance-review.md` and
`integration-attribution-cross-review.json`.

The preceding review also retained the overbroad attribution. That was an
error in the earlier review, and correcting it in the live records is
necessary even though it does not alter the source interpretations.

### 5. The historical Kialo pilot needs a current binding

The old pilot builder hardcodes `fcccbba` and engine walton.5 and reads the
historical 143-record register. Passing it a current library path does not
make the result current. Four of its fourteen candidate records still carry
old holds although they are now accepted within declared scope. Their
executable contracts are unchanged; one citation changed. A name-only gloss
selection also carried an unused `relevant_differences/2` template alongside
the actually used arity.

`tools/validation/refresh_kialo_context.py` creates a separate derivative
from immutable Git objects at an explicit commit. It checks the live
149-record ledger/registry inventories and matching review stamps, preserves
the twelve machine proposals and original source text, selects the candidate
glosses by arity, and retains zero human gold. Its eleven deliberate mismatch
controls cover stale binding, changed formula/gloss, source corruption,
mislabelled gold and leakage into the source-only packet. Approval
fingerprints are copied attestations here; the independent repository
source gate remains required. This is preparation, not an engine benchmark
or semantic annotation validation.

## Independent acceptance and verification boundary

The server reviewer did not author commit `6647eae`, but did author this
new four-file server repair. The root reviewer separately inspected its
source, route/flag wiring and tests, then ran eight independently authored
positive/negative requests against the actual server. All eight passed.
The acceptance receipt binds the four reviewed file hashes; see
`server-boundary-non-author-acceptance.json`. The server review's own paired
matrix comprises 36 requests per binary: all 25 negative corrected requests
return the intended non-2xx outcome without a result graph, while all eleven
positive parsed graphs match the baseline. Seven matching CLI cases per
binary retain the intended behavior. Focused tests pass under both pinned
Go versions and the Go 1.27.4 race detector.

The complete release gate is still run on a frozen exact commit in clean
full clones under both Go versions. Its current registered inventory remains
10,428 distinct checks. The new HTTP tests and pilot-preflight controls are
separate checks, not additions silently folded into that number. Native
build/vet/unit/race gates, source integrity, validator output and exact
cross-version agreement are reported separately in the final evidence.

All non-author acceptances remain technical AI review. These records do not
authenticate human identities or provide institutional or original-author
certification. The HTTP work does not certify production hosting, the Dung
web route, arbitrary Graphviz behavior, request-size protection or atomic
export failure after bytes have been written.

## Readiness and next action

Use `../validation-entry.md` as the immediate protocol. Kialo is suitable
for the first structural and human-interpretation pilot. It is not an
oracle for logical correctness: Pro/Con placement, votes and impact scores
do not by themselves supply executable scheme instances, negation,
acceptable evidence weights or ground-truth conclusions.

The supplied derivative represents one discussion family, with 8,319 claim
records and 10,947 location records; its current thesis view has 3,554 claims
and 3,690 relations. The twelve draft occurrences cover eleven child claims
and remain development-exposed. The six-unit source-only exercise has five
child claims in that same family. None is human-adjudicated gold or a sealed
test case. Additional owner-held captures have not been inspected here.

First freeze task definitions, independent references, profiles, outcome
rules and leakage-aware sampling. Run structural auditing and annotation
calibration alongside independent bounded formal-oracle development. Use
the existing 12/12/36 feasibility budget only after family-level sampling
is possible; set any empirical acceptance thresholds before opening sealed
evaluation. Full ICCMA, AIFdb, other corpora, capacity characterization and
collection-scale annotation follow their own eligibility gates. Do not
repeat broad source review absent a concrete counterexample or new source.

No remote push, tag, release or outside correspondence is part of this
review. D22 remains in force. A contributor can integrate the reviewed
patches and exact bundle, rerun CI on the eventual merge commit, and then
start the bounded validation work against that accepted pin.
