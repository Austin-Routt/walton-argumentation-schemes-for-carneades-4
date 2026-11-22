# Independent review of the HTTP evaluation boundary

## Scope, base and decision

This review examines the HTTP import and evaluation boundary added by
`6647eae994c808765af6eda1e449acf4619f1218`, as integrated into public master
`d9a25857277140a5d3f10aaf1dfa4ade0f15b91f` (tree
`fcea512fd99551b453ec283158ade395262d4b7c`). It reviews the actual server routes,
not just the new format-policy helper. It does not reapprove scheme records or
change inference semantics.

**The legacy-import gate works, but the claim that the server integration was
complete needs correction.** The form route discarded inference errors, and the
raw YAML/JSON route skipped authored-argument validation. Both defects can return
an evaluated graph where the corresponding CLI correctly refuses evaluation.
The legacy format gate itself also returned its refusal with HTTP 200. These are
result-integrity and admission defects worth closing before using the server in
validation work.

I authored the narrowly scoped repair in `src/web`, added permanent tests of the
actual handlers, and re-ran both Go toolchains. The parent reviewer must inspect
and independently accept this repair; this report is independent of the
pre-existing integration commit, but it is not an independent approval of my own
new code. No repository commit, tag, release, remote push, or outside contact was
made by this reviewer.

## Findings

### WEB-1 — The form route exported successful results after inference failed

The original `/carneades/eval` validated the imported graph but discarded the
return value of `ag.Infer()`. It then labelled and exported the graph. The engine
already has a typed error contract and atomic staging for failed inference; the
route failed to use that contract.

Two small counterexamples reproduce against a running baseline CLI server:

- Contradictory stipulated statements, `p` and `¬p`, produce HTTP 200 and an
  evaluated YAML graph with both statements labelled `in`. The same input through
  `carneades eval` returns exit status 3, `inconsistent context`, with no stdout.
- A rule containing an unsupported guard, `X == a`, produces HTTP 200 and a graph
  even though inference refuses that rule. The corresponding CLI returns exit
  status 2, `unsupported`, with no stdout.

The repair checks `InferContext(req.Context())`, returns the typed failure, and
never reaches labelling or export on that failure. Both counterexamples now
return HTTP 422 and no result graph. A positive rule with the same premise and
conclusion but no unsupported guard still derives its conclusion and labels it
`in`.

### WEB-2 — The raw route accepted invalid authored arguments

The original `/carneades/eval-arg-graph` already checked inference errors. It did
not call `validation.Validate` before inference. An imported authored argument
can refer to a named scheme while using premises inconsistent with that scheme's
instantiated form. Inference alone is not authored-argument validation.

The concrete fixture defines scheme `s`, `p(X) -> r(X)`, and an argument with
parameter `a`, premise `q(a)`, and conclusion `r(a)`. The baseline raw route
returns HTTP 200 JSON with the invalid authored argument and `r(a)` labelled
`in`. Replacing the premise with `p(b)` while retaining parameter `a` reproduces
the same failure for a wrong ground instance rather than merely a different
predicate. The CLI rejects both with exit status 1 and no stdout. The baseline
form route detects both validation errors, but returns its error page as HTTP
200.

The repair uses the same validation-and-inference function in both HTTP routes.
Both now refuse the two authored-argument fixtures with HTTP 400, `invalid input`,
and no result graph. This change may cause callers that previously depended on
unvalidated raw-API input to receive errors; it is intentional admission
consistency with the existing CLI and engine contracts.

### WEB-3 — “Before the upload is read” overstates the format gate

`req.FormValue` parses multipart data before `checkImportProfile` runs. The
format gate runs before invoking a legacy converter or evaluating the graph; it
does not establish that the uploaded bytes have not been read or spooled.
Source inspection is sufficient to distinguish these boundaries. The corrected
code comment says “before conversion.” The repair also closes the opened upload
and removes temporary multipart files after the form request.

The actual command-line flag wiring is correct. With no flag, valid and
malformed CAF, AGXML and LKIF are all refused by the profile gate. With
`server -experimental-import`, each valid tiny legacy fixture imports and
evaluates, and malformed fixtures reach and fail their converter checks. Opt-in
does not bypass validation or certify a converter's fidelity.

The old form route's refusal pages had HTTP 200. The new handler returns HTTP
422 for a disabled or unknown format and HTTP 400 for the malformed opted-in
fixtures. HTML error rendering is retained.

### OBS-1 — Restricted AIF admission also depends on native statement IDs

A restricted AIF object with one I-node named `p` is admitted and evaluates on
both versions. The otherwise identical object with `nodeID: "1"` passes AIF
structural import but fails native graph validation: the statement/assumption key
is not a ground atomic formula. The CLI and baseline form route already refused
that numeric-ID case; the repaired handler retains the refusal with an explicit
HTTP 400 status.

This is an existing boundary, not a change introduced by the repair. It matters
when planning AIFdb or external-graph ingestion. General AIF identifiers require
an explicit, tested mapping into native identifiers if they do not already meet
the native ground-formula requirement. The full-server evidence retains the
numeric case as a negative compatibility control and uses `p` for the positive
AIF control. It does not claim general AIFdb conformance.

## Exact proposed code changes

Only these four source files were changed or added by this reviewer:

| File under `third_party/carneades-4/src/web/` | Purpose |
| --- | --- |
| `evaluation.go` | Shared authored-graph validation and context-aware inference; form and raw HTTP handlers; typed error-to-HTTP mapping |
| `evaluation_test.go` | Three new permanent Go tests that exercise the real handlers, positive controls and cancellation plumbing |
| `web.go` | Wire the extracted handlers to the existing routes; remove their old inline implementations and unused imports |
| `profile.go` | Correct the comment about when the profile gate runs |

The status mapping is explicit:

| Typed outcome | HTTP status |
| --- | --- |
| `invalid input` | 400 Bad Request |
| `unsupported`, `inconsistent context`, `limit reached` | 422 Unprocessable Entity |
| `cancelled` | 408 Request Timeout |
| Internal/unclassified failure | 500 Internal Server Error |

This mapping does not change CLI exit codes. Raw-API failures are JSON strings,
encoded using the standard JSON encoder; form failures use the existing HTML
error template. Successful result formats and route names are preserved. The
raw API retains its existing OPTIONS/CORS behavior. The common helper stops
before labelling/export on validation or inference failure.

## Verification performed

### 1. Real CLI server, paired before/after observations

`http_probe.py` starts the compiled server on a fresh localhost port, waits for
its form route, sends the requests, records the full responses, and terminates
the server. Every HTTP request has a 12-second timeout; startup and termination
are bounded. Default and explicit-opt-in modes are separate server processes.
No public or third-party service receives these requests.

There are **36 requests per binary**: 15 form cases in each startup mode and six
raw-API cases in default mode. `http-before.json` and `http-after.json` include
status, headers, full body, input hash, and executable hash.

`summarize.py` independently compares the observed matrices with the stated
expected statuses. It verifies that:

- All 25 negative patched requests return the expected non-2xx status and export
  no result graph.
- All 11 positive patched requests return HTTP 200 and have the same parsed
  YAML/JSON graph as the baseline.
- The input hashes match between each before/after pair.

Parsed graph equality is deliberate: the YAML exporter may change map-key order
between processes. Byte-order differences are recorded separately and are not
misreported as semantic changes. `http-comparison.json` records every comparison
and the successful assertions.

### 2. Corresponding CLI controls

`cli_probe.py` runs seven matching cases against each binary: simple native
input, valid rule inference, inconsistent assumptions, unsupported guard,
wrong-predicate authored argument, wrong-instance authored argument, and numeric
AIF identifier. All 14 CLI runs return the expected exit codes. Failed CLI runs
have no stdout. The repair does not modify these CLI behaviors.

### 3. Permanent Go handler tests

`TestHTTPEvaluationRefusesFailedGraphs` uses real `httptest.NewServer` HTTP
requests for both handlers. It covers the two inference failures, both authored
argument failures, and an unknown YAML key. It requires the intended status,
typed outcome, and absence of result-graph fields. Raw failures must also be
valid JSON strings. A positive rule must derive and label `q(a)` as `in`.

`TestHTTPImportProfile` sends valid native YAML, restricted AIF and three legacy
fixtures with opt-in disabled and enabled. It also sends malformed legacy
fixtures and an unknown format. It demonstrates both admission and refusal,
including conversion after explicit opt-in.

`TestHTTPEvaluationUsesRequestCancellation` sends an already-cancelled request
context directly through each actual handler. Both return the cancellation
outcome with HTTP 408 and no result graph. This is a deterministic test of
request-context propagation; it is not a test of all network-disconnect timing
or of cancellation during output streaming.

The existing `TestWebImportProfile` also continues to pass. The package's four
top-level tests pass under **Go 1.24.13 and Go 1.27.4**, uncached. The Go 1.27.4
race-detector run also passes. `git diff --check` reports no whitespace errors.
Final logs and file hashes accompany this report. The parent reviewer owns the
complete repository gates and independent review of this new delta.

## Acceptance boundary and remaining integration work

The evidence supports accepting the narrow repair after a non-author reviews
the exact hashed source and re-runs the required gates. The new code fixes
admission and result integrity for the two CAES HTTP evaluation routes; it is
not a certification of the server as a general hosted service. This review does
not assess the Dung form handler, arbitrary request-size controls, service
authentication, multi-tenant capacity, external Graphviz execution, or atomic
failure handling after an output writer has already emitted bytes. Those
behaviors were not expanded by the import-gate task.

The new HTTP status code for limit/internal outcomes is defined in the common
mapping; this focused request matrix does not intentionally exhaust runtime
quotas or inject internal engine failures. Existing engine tests remain the
evidence for those engine outcome classifications.

The parent review should update the engine version and provenance, correct any
remaining “before upload is read” prose, document the non-2xx behavior and AIF ID
boundary, regenerate the reproducible source archive, and run the full required
gates. Tags or release claims should identify the resulting accepted commit,
not imply that this uncommitted repair already exists on public master.

## Reproduction and evidence inventory

The commands recorded in `verification.json` use the two installed Go binaries
and a temporary module file that replaces GoCHR with this repository's vendored
source. To reproduce elsewhere, make that replacement point to the local
checkout; the absolute paths in the captured module file are session metadata,
not portable project configuration.

For either compiled binary, run:

```sh
python3 http_probe.py --binary /path/to/carneades --repo /path/to/repo --out http-after.json
python3 cli_probe.py --binary /path/to/carneades --out cli-after.json
```

For the paired comparison, retain the baseline observations as
`http-before.json`, the patched observations as `http-after.json`, and run
`python3 summarize.py` in this evidence directory. It requires PyYAML only for
comparison of parsed output; the server probe itself uses Python's standard
library.

The evidence directory contains full HTTP and CLI JSON observations, concise
HTTP logs, server logs, both focused Go test logs, the race-detector log,
`http-comparison.json`, `findings.json`, `verification.json`, and
`reviewed-file-hashes.json`. Locally built binaries and the temporary module
directory can be omitted from a contributor archive; their identifying hashes
are preserved in the evidence. This report makes no claim that a locally built
binary is itself a release artifact.
