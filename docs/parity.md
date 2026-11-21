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
| per scheme: `assumption` (a withdrawn assumption) | covered | `iso/assumptions-questionable`, `full/assumptions-questionable`, restated for the questioned state of D2 |
| per scheme: `roundtrip` (export and import, parameters compared) | covered | `iso/roundtrip` (ported): every argument with its scheme, parameters, premises and undercutter, and every label, after export and import with clean globals; `roundtrip` of every reference case in a second process |
| per scheme: `negation-text` | covered | `iso/negation-texts` |
| static: `gochr-name` (built-in names as predicates) | covered | `static/no-reserved-functor`, `is-as-predicate-in-conclusion` |
| static: `shared-0ary` (statements without variables used by several schemes) | covered | `static/shared-0ary` in `TestSource`: rejects unreviewed cross-scheme nullary predicates, including use beneath negation. The reviewed exception is `do_obligator/0`, with its exact universal-warrant wording and the exact positive-premise consumer set `regulative_rule`, `regulative_rule_prohibition` (PLAN.md 5.4; CHANGELOG.md 3.2 and 5.4). A changed consumer set or wording fails. Mutation tests cover unexpected sharing, changed consumers, polarity and burden; controls keep local nullaries, statement variables and other arities distinct. |
| static: `loop` (schemes that can re-trigger each other) | covered | `static/not-self-feeding`; the cycle report of `go run ./cmd/schemecheck` |
| probe: `global-state` | covered | `yaml-import-keeps-globals-clean` |
| probe: `argument-validation`, `argument-validation-panic` | covered | `validator-checks-arguments-against-schemes`, `validator-accepts-correct-argument` |
| probe: `argument-binding` | covered | `validator-enforces-shared-bindings` |
| probe: `inconsistent` | covered | `inconsistent-assumptions-detected`, `inconsistent-follows-current-assumptions` |
| probe: `infer-twice`, `infer-idempotent` | covered | `infer-twice-with-exception`, `infer-is-idempotent` |
| probe: `negated-goal` (a negated assumption as a goal) | covered | `negated-premises-match`; `negated-assumption-alone` (ported) |
| probe: `assumption-semantics` | covered | `questioned-assumption-is-not-assumed`, `questioned-assumption-holds-when-supported` |
| probe: `export-import-fresh` | covered | `cli-reads-its-own-yaml` (a new process) |
| probe: `stress` (every premise of every scheme asserted at once) | covered | `stress` of each scheme file (ported): the run completes, every scheme constructs its argument, every statement is ground; a scheme whose premise negates one already asserted gets constants of its own |
| probe: `proposed-fix` | dropped | measured the second review's prototype rules, which the plan drops (D8) |
| the six acceptance cases with expected labels (`cases/*.yml`) | covered | `tools/vv/cases/acceptance.yml` (ported), with the book entry of each |

## Third review: `tools/enginecheck`

| Check | Status | Covered by |
|---|---|---|
| `TestBudgetExhaustionMustBeReported` | covered | `budget-exhaustion-is-reported` |
| `TestCompletionSignals` (limit, error and stdout together) | covered | `budget-exhaustion-is-reported`, `engine-keeps-stdout-clean`; one check for each outcome of D9: complete (every case), invalid input, unsupported, inconsistent context and limit (`outcome/*` reference cases), cancelled (`cli-interrupt-is-cancelled`, `TestInferContextCancelled`), internal failure (`TestExitStatus`) |
| `TestConcurrentImports`, `TestConcurrentSpellInitialization`, the race job | covered | the engines' `TestConcurrentImports`, `TestConcurrentSpellInitialization`, `TestConcurrentInference` and `TestWaitingRunCanBeCancelled`, under `-race` in `tools/vv/engine-tests.sh` (ported). Without the fix both review tests race; calls are now serialized (contract 1) |
| `TestDefaultAssumptionAndExplicitContraryIssue` | covered | `default-yields-to-stipulated-negation`, `default-yields-to-stipulated-position`, `contrary-defaults-suspend-each-other` (ported; they found that a default and its stipulated negation were both in, contract 5) |
| `TestImportedSchemeIsNotBasic`, `TestImportsAreIsolated` | covered | `yaml-import-keeps-globals-clean` |
| `TestInferenceRejectsInvalidRule`, `TestMalformedRuleMustBeRejected` | covered | `invalid-rule-is-reported` (ported); GoCHR `TestAddRuleRefusesMalformedText`, which now also refuses a number or a string as a rule body (the review's `numeric_body`) |
| `TestMalformedGoalMustBeRejected` | covered | GoCHR `TestInferRefusesMalformedGoal`, run by `tools/vv/engine-tests.sh`; `malformed-statement-is-an-error` |
| `TestInvalidExpectedLabelDoesNotPanic` | covered | `invalid-expected-label-is-rejected` (ported) |
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
| `run_python_audit.py`: 20 schema probes of `tools/validate_schemes.py` | covered | `tools/test_validate_schemes.py`, `SchemaProbes` (ported; the 17 gaps closed) |
| `full_workflow_mutations.py`: two semantic mutations | covered | `tools/test_validate_schemes.py`, `Mutations` (ported, with three more) |
| `sources.lock.json` with a hard failure on a hash mismatch | covered | `tools/vv/setup.sh` with `sources.lock` |
| per scheme: a second inference | covered | `iso/infer-twice`, `full/infer-twice` (ported) |
| per scheme: a withdrawn assumption | covered | `iso/assumptions-questionable`, `full/assumptions-questionable`: restated for the questioned state of D2 (ported) |

## Ported

Each port names its new check here and removes the row's **to port**.

| Check | Ported as | With |
|---|---|---|
| third review, per scheme: a second inference | `iso/infer-twice`, `full/infer-twice` (456 checks) | PLAN.md 2.5 |
| third review, per scheme: a withdrawn assumption | `iso/assumptions-questionable`, `full/assumptions-questionable` (456 checks) | PLAN.md 2.7 |
| third review, `run_python_audit.py` | `tools/test_validate_schemes.py` | PLAN.md 3.10 |
| third review, `full_workflow_mutations.py` | `tools/test_validate_schemes.py` | PLAN.md 3.10 |
| second review, `TestNoRuleFiringMustPreserveStore` | GoCHR `TestNoRuleFiredKeepsGoals` | PLAN.md 1.3 |
| second review, static `shared-0ary` | `static/shared-0ary`, `TestSharedNullaryMutations` and positive controls | PLAN.md 5.4; reviewed shared universal warrant retained |
