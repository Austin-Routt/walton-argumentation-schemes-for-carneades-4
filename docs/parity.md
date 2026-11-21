# Parity register

`tools/vv` is the one verification gate (PLAN.md, D7). The two other harnesses written during the reviews hold checks of their own. This register lists every one of them against the `tools/vv` check that covers it. A check marked **to port** is added to `tools/vv` before the evidence of that harness is set aside; when it is ported, its row names the new check.

The two harnesses are not in this repository. They are in the review archives:

* the second review's `tools/go/schemecheck` (Claude): branch `vv/engine-harness` of the review, delivered as `vv-engine-harness.patch`;
* the third review's `tools/enginecheck` and `tools/run_engine_audit.py` (ChatGPT): its review package, `repository/tools/`.

The third review's runner is not run beside `tools/vv`: it unpacks and audits the two archives, so once fixes land in `third_party/` it would keep testing the unmodified engine.

Status: **covered** (a `tools/vv` check tests the same thing), **partly** (the `tools/vv` check is narrower; the row says what is missing), **to port**, **dropped** (with the reason).

## Second review: `tools/go/schemecheck`

| Check | Status | Covered by |
|---|---|---|
| per scheme: `chr` (GoCHR accepts the derived rule) | covered | `chr-translation` |
| per scheme: `instantiate`, `infer`, `terminate`, `ground` | covered | `iso/fires`, `iso/infer-ok`, `iso/terminates`, `iso/ground` and the `full/` counterparts |
| per scheme: `label` (conclusions in) | covered | `iso/conclusions-in`, `full/conclusions-in` |
| per scheme: `exception` (each exception undercuts) | covered | `iso/exceptions-undercut`, `full/exceptions-undercut` |
| per scheme: `premise` (each premise needed) | covered | `iso/premises-needed`, `full/premises-needed` |
| per scheme: `assumption` (a withdrawn assumption) | to port | engine level only: `baseline-withdrawn-assumption-stays-withdrawn` |
| per scheme: `roundtrip` (export and import, parameters compared) | to port | engine level only: `export-keeps-compound-parameters`, `cli-reads-its-own-yaml` |
| per scheme: `negation-text` | covered | `iso/negation-texts` |
| static: `gochr-name` (built-in names as predicates) | covered | `static/no-reserved-functor`, `is-as-predicate-in-conclusion` |
| static: `shared-0ary` (statements without variables used by several schemes) | to port | none; belongs with PLAN.md 5.4 |
| static: `loop` (schemes that can re-trigger each other) | covered | `static/not-self-feeding`; the cycle report of `go run ./cmd/schemecheck` |
| probe: `global-state` | covered | `yaml-import-keeps-globals-clean` |
| probe: `argument-validation`, `argument-validation-panic` | covered | `validator-checks-arguments-against-schemes`, `validator-accepts-correct-argument` |
| probe: `argument-binding` | covered | `validator-enforces-shared-bindings` |
| probe: `inconsistent` | covered | `inconsistent-assumptions-detected`, `inconsistent-follows-current-assumptions` |
| probe: `infer-twice`, `infer-idempotent` | covered | `infer-twice-with-exception`, `infer-is-idempotent` |
| probe: `negated-goal` (a negated assumption as a goal) | partly | `negated-premises-match` tests a negated premise against a negated fact; a negated goal alone is not tested |
| probe: `assumption-semantics` | covered | `questioned-assumption-is-not-assumed`, `questioned-assumption-holds-when-supported` |
| probe: `export-import-fresh` | covered | `cli-reads-its-own-yaml` (a new process) |
| probe: `stress` (every premise of every scheme asserted at once) | to port | `no-scheme-is-starved` covers starvation, not the full load |
| probe: `proposed-fix` | dropped | measured the second review's prototype rules, which the plan drops (D8) |
| the six acceptance cases with expected labels (`cases/*.yml`) | to port | none: `tools/vv` derives its expectations per scheme; these state labels by hand for realistic graphs |

## Third review: `tools/enginecheck`

| Check | Status | Covered by |
|---|---|---|
| `TestBudgetExhaustionMustBeReported` | covered | `budget-exhaustion-is-reported` |
| `TestCompletionSignals` (limit, error and stdout together) | partly | `budget-exhaustion-is-reported`, `engine-keeps-stdout-clean`; the seven outcomes of D9 need one check each |
| `TestConcurrentImports`, `TestConcurrentSpellInitialization`, the race job | to port | outside the first profile (D3); port as known failures so the race stays visible |
| `TestDefaultAssumptionAndExplicitContraryIssue` | partly | `inconsistent-assumptions-detected`; the case of a default against an explicit contrary argument needs its own check |
| `TestImportedSchemeIsNotBasic`, `TestImportsAreIsolated` | covered | `yaml-import-keeps-globals-clean` |
| `TestInferenceRejectsInvalidRule`, `TestMalformedRuleMustBeRejected` | partly | `malformed-statement-is-an-error` covers a statement; GoCHR's own `AddRule` on malformed rule text is not tested directly |
| `TestMalformedGoalMustBeRejected` | to port | none |
| `TestInvalidExpectedLabelDoesNotPanic` | to port | none |
| `TestIssueExtensionKeepsBidirectionalMembership` | covered | `issue-extension-keeps-positions`, `labelling-returns-after-issue-gains-position` |
| `TestLinkedEvidenceAndExceptionLabels` | covered | `baseline-scheme-with-exception` |
| `TestManualValidArgumentDoesNotPanic` | covered | `validator-accepts-correct-argument` |
| `TestNegatedPremiseMustFire` | covered | `negated-premises-match` |
| `TestNoRuleFiringMustPreserveStore` | covered | GoCHR `TestNoRuleFiredKeepsGoals`, run by `tools/vv/engine-tests.sh` (PLAN.md 1.3) |
| `TestPositiveRule` | covered | `iso/fires` for every scheme |
| `TestRepeatedInferenceDoesNotDuplicateArguments`, `TestRepeatedInferenceWithException` | covered | `infer-is-idempotent`, `infer-twice-with-exception` |
| `TestRolesArePreserved` | covered | `scheme-roles-are-imported` |
| `TestRoundTripPreservesCompoundParameters`, `...ExpectedLabels`, `...Theory` | covered | `export-keeps-compound-parameters`, `export-keeps-expected-labels`, `yaml-export-can-be-reimported` |
| `TestSchemeParameterBindingsMustAgreeWithConclusion` | covered | `validator-enforces-shared-bindings` |
| `run_python_audit.py`: 20 schema probes of `tools/validate_schemes.py` | to port | none; PLAN.md 3.10 |
| `full_workflow_mutations.py`: two semantic mutations | to port | none; PLAN.md 3.10 |
| `sources.lock.json` with a hard failure on a hash mismatch | covered | `tools/vv/setup.sh` with `sources.lock` |
| per scheme: a second inference | covered | `iso/infer-twice`, `full/infer-twice` (ported) |
| per scheme: a withdrawn assumption | covered | `iso/assumptions-questionable`, `full/assumptions-questionable`: restated for the questioned state of D2 (ported) |

## Ported

Each port names its new check here and removes the row's **to port**.

| Check | Ported as | With |
|---|---|---|
| third review, per scheme: a second inference | `iso/infer-twice`, `full/infer-twice` (456 checks) | PLAN.md 2.5 |
| third review, per scheme: a withdrawn assumption | `iso/assumptions-questionable`, `full/assumptions-questionable` (456 checks) | PLAN.md 2.7 |
| second review, `TestNoRuleFiringMustPreserveStore` | GoCHR `TestNoRuleFiredKeepsGoals` | PLAN.md 1.3 |
