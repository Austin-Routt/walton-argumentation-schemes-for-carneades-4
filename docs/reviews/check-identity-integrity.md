# Verification inventory repair

Baseline: `3a1ea97809f6cc7207d7bcd1ed431990d4cc328b`.

Repair commit: `45f571e8627c33f3cdd9d1dd010d6b8435d4965b`, created in the separate full clone `release-review/gate-fix`.

## Finding and effect

The previous complete Go 1.27.4 result file contained 8,206 result objects but only 8,201 distinct `(file, scheme, check)` identities. Five pairs of historical migration forms produced the same current predicate. Their check names included only migration kind and current form, so each pair collided. The Go manifest gate silently deduplicated result identities; the Python comparator silently overwrote matching identities. The comparator also coerced arbitrary `pass` values to booleans, so the JSON string `"false"` became true. These are harness defects, independently of whether the underlying scheme executions passed.

The five affected record families are `practical_necessary_condition`, `practical_sufficient_condition`, `retroductive`, `abductive_forward`, and `abductive_action_to_character`. Their common target forms are respectively `necessary_action(P,X,S,A)`, `sufficient_action(P,X,S,A)`, `reference_class(A,E,Q)`, `successful_diagram(G,G1)`, and `action_fits_quality(A,X,Q)`.

All 204 migration rows have distinct original `(record, kind, old, trigger, new)` tuples. Migration names now use:

```
migration/<kind> old="<original old>" trigger="<original trigger>" new="<original new>"
```

The original values are captured before predicate rename normalization. This matters because historical spelling corrections can otherwise collapse two different old forms even when `old` is added to the name. File and record remain separate parts of the full result identity. Reordering rows leaves identities unchanged; a truly repeated mapping is rejected, not assigned an arbitrary row number.

## Enforcement

`ResultKeys` returns the complete sorted inventory or an error for duplicate keys, regardless of their pass values or details. `TestMain` checks every selected run, including runs that omit the full ledger test. `TestLedgerIsUsed` uses the same validator before considering manifest updates.

`ParseExpectedChecks` rejects repeated manifest entries. Manifest update mode validates the existing manifest before writing, so it cannot quietly repair or erase duplicated entries. An absent or empty manifest can still be initialized explicitly. The existing minimum-check and missing/added-check gates remain responsible for complete suite coverage.

`compare_results.py` now requires each record to be a JSON object with string `file`, `scheme`, and `check` fields and an actual JSON boolean `pass`. Duplicate result identities are rejected even when both outcomes agree. Bad JSON, malformed fields, duplicates and unreadable inputs exit 2 with a diagnostic; valid disagreement exits 1; valid agreement exits 0. Matching failures still constitute agreement, not a successful strict release gate.

## Verification performed

- Go 1.27.4: all `internal/...` harness tests passed; `go vet ./...` passed.
- Five added Go test functions cover duplicate result identities with matching/conflicting outcomes, preservation of distinct full identities, duplicate manifest entries including CRLF, valid manifests, and migration identity changes/reordering/true duplicates.
- The focused strict `TestSource` run passed 649 of 649 checks, with 649 distinct identities. Its actual result file is `source-results.jsonl`.
- Eleven Python comparator tests passed, with subcases for malformed booleans, missing/invalid key fields, non-object JSON records, duplicate records on either side, line-number reporting, valid reordering, matching false outcomes, outcome/key disagreements, unreadable files, and malformed-JSON CLI exit status.
- The repaired comparator rejected the actual historical Go 1.24.13 log with exit 2 at line 7888, identifying a duplicate of line 7769.
- An isolated compiled harness probe confirmed actual execution paths: a minimal valid inventory passed; duplicate emitted results failed a focused run; duplicate manifest entries failed ordinary and update modes. The manifest remained byte-for-byte unchanged after both rejection paths. Outputs are in `actual-gate-mutations.json`.
- An inventory-only analysis replaced the 205 historical migration results (readability check plus 204 rows) with the 205 freshly executed results. All non-identity fields and outcomes matched. The resulting inventory contained exactly 8,206 rows and 8,206 unique keys. This establishes the intended identity count; it is not a new complete runtime execution. Details are in `inventory-analysis.json`.

The temporary probe test was compiled solely to exercise the actual harness guards, then removed. It injected two identical passing `Result` objects into `collect`; a temporary working directory held a minimal zero-minimum ledger and either an empty manifest or two identical `f|s|c` lines. The committed tests and release sources contain no probe bypass.

## Integration boundary

The patch changes nine files under `tools/vv`: the source migration names, ledger helpers and tests, main suite guard, comparator and tests, runner, and README. It changes no engine or library semantics, source stages, approvals, expected manifest, or minimum count. The integrating reviewer must regenerate `expected_checks.txt` once, then rerun the complete suite under both toolchains and compare those fresh, duplicate-free logs. Do not use the historical duplicate-containing reports as final cross-toolchain evidence.
