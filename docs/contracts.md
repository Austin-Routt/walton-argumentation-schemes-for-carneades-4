# Contracts

The supported behavior of the reviewed engine in `third_party/`. Each contract names executable checks. A green known-issues ledger records the tested cases; it is not proof that an untested path or every library interpretation is correct. Source approval is a separate gate in `PLAN.md`.

A contract changes only together with its checks, in one reviewed change.

## 1. The supported profile (D3)

**Entry points.** The `carneades` command (`check`, `eval`) and the Go API (`yaml.Import`, `ArgGraph.Infer`, `ArgGraph.GroundedLabelling`, `validation.Validate`), offline. The web server is outside the profile.

API clients evaluating authored graphs must first run `validation.Validate`, handle any problems, then call `CheckContext`/`InferContext` and handle their typed errors before labelling. CLI `check` and `eval` perform structural validation and context checking themselves. Calling `Infer` alone does not certify an authored argument's exception links.

**Concurrent use.** Calls from several goroutines are serialized, not run in parallel: an inference run, and the translation of a theory into rules (`caes.RuleStoreOf`), wait until no other run uses the rule engine, whose parser keeps package-level state; `yaml.Import` and `yaml.Export` wait for each other, since they share the importer's tables. Each call gives what it gives alone. A run that waits with a context that is done ends as cancelled (contract 3). GoCHR's spelling tables are filled once (`sync.Once`). A rule store returned by `RuleStoreOf` and used directly is outside this protection. Checked by the engines' `TestConcurrentInference`, `TestWaitingRunCanBeCancelled`, `TestConcurrentImports` and `TestConcurrentSpellInitialization`, run under the race detector by `tools/vv/engine-tests.sh`.

**Schemes.** A scheme is a propagation rule: its premises, together with no deletions and no guards, yield its conclusions, assumptions and exceptions as instantiated statements. A scheme with `deletions` or `guards` is outside the profile and is reported as **unsupported** (contract 3), not run. Every variable of a scheme occurs in a premise (range restriction).

**Terms.** A constant is a lower-case name, an integer or a quoted string; a capitalised name is a variable. A statement carries no arithmetic: an unquoted `1990-05-17` is invalid input, not the number 1968 (issue E-TERMS). A statement may be a variable of a scheme (a statement variable), as in `expert_opinion`.

**Facts.** The ground statements of a graph form a set (contract 2).

Checked by: `iso/*`, `full/*`, `cli/*`, `static/range-restricted`, `string-constants-survive`, `arithmetic-in-a-statement-is-invalid-input`, `deletions-and-guards-are-unsupported`, `yaml-import-keeps-globals-clean`.

## 2. One fact, one argument (D8)

**A fact** is a ground statement, identified by its canonical term: functor, arity and arguments compared structurally after normalization, so `p(a,b)` and `p(a, b)` are one fact and `p(a)` and `p(a,b)` are two. Adding a fact that is already present changes nothing. GoCHR keeps its multiset semantics for its own programs; the set semantics are switched on by Carneades for argument construction (PLAN.md D8).

Logical conjunction may reuse a fact: `person(alex)` can satisfy both `person(X)` and `person(Y)` with `X=Y=alex`. No implicit inequality is added. Carneades enables GoCHR's `ReuseConstraints` option; ordinary GoCHR programs keep distinct-head-constraint behavior by default. Matching streams candidate tuples, preserves propagation history and checks cancellation/work quotas during unsuccessful searches as well as successful applications.

**An argument** is identified by its scheme and its canonical parameter list. Two arguments for one conclusion from different schemes, or from one scheme with different parameters, are two arguments and both are kept. Running inference again on an unchanged graph adds nothing and leaves every label as it was.

**Termination.** Set semantics end the loops of rules that re-derive a fact. They do not bound rules that build ever larger terms, or matching work; the limits of contract 3 do.

Checked by: `rederivation-terminates`, `no-scheme-is-starved`, `infer-is-idempotent`, `infer-twice-with-exception`, `wrong-arity-fact-is-harmless`, `growing-term-is-bounded`; issues X-LOOP, E-IDEMP, E-MATCH, E-BUDGET.

## 3. How a run ends (D9)

A run of `check`, `eval` or `Infer` ends in exactly one of seven ways.

| Outcome | Meaning | Exit status of the command |
|---|---|---|
| complete | the requested operation completed (`check` validates; `eval` also infers and labels) | 0 |
| invalid input | the file cannot be read, a statement cannot be parsed, or validation reports a problem | 1 |
| unsupported | the input uses something outside the profile of contract 1 | 2 |
| inconsistent context | the graph stipulates both a statement and its negation, or two positions of one issue (contract 4) | 3 |
| limit reached | a logical work/storage quota ran out while work remained | 4 |
| cancelled | the caller stopped the run | 5 |
| internal failure | a panic or a broken invariant of the engine | 6 |

**Through the API** `Infer` returns `nil` only for complete; otherwise an error whose kind is one of the six others, testable with `errors.Is`. On any error the graph is left as it was before the call: a partial graph is never returned as a result.

Inference constructs a private mutable graph copy and publishes it only on success. Failure and cancellation preserve the original topology, maps, slices and node contents; a successful publication retains the identities of pre-existing statement, argument and issue objects. The caller must not mutate the same graph concurrently. Theory metadata and weighting functions are treated as read-only during inference.

**Output.** The engine writes nothing to stdout; diagnostics go to stderr. On any outcome other than complete the command prints no evaluation.

**Limits.** Carneades sets these quotas for the supported inference path:

| Work or data | Limit |
|---|---:|
| Rule applications (`MAXRULEAPPS`) | 100,000 |
| Nodes in one term (`MAXTERMSIZE`) | 500 |
| Matching candidate visits, including failures (`MAXMATCHSTEPS`) | 50,000,000 |
| Inserted constraints (`MAXCONSTRAINTS`) | 100,000 |
| Cumulative inserted term nodes (`MAXSTORENODES`) | 2,000,000 |
| Issue candidate/pattern visits per materialization phase (`MAXISSUESTEPS`) | 50,000,000 |

The two visit budgets are finite safety limits. The earlier 1,000,000-visit quotas refused the reference four-way join over 120 facts and modest issue-closure examples. The larger limits admit those cases; elapsed time depends on the machine and theory. A larger finite evaluation can still be refused even when an unbounded evaluation would eventually finish. Issue construction indexes parsed statement keys when the remaining patterns are ground under the bindings of the first, preserving noncanonical API spacing and checking typed matching. Open patterns, statements with variables and enumerations retain the scan. An exact rule-application boundary is complete if no further firing remains. Duplicate facts consume no new insertion allowance. A caller's context cancels `ArgGraph.InferContext`; CLI evaluation connects interrupt/termination signals to that context. Cancellation is checked even when there are no rules or no matching rule, and during candidate search, issue construction and argument materialization. A cancelled/deadline context produces cancellation, rather than a successful partial evaluation.

These are logical work/storage quotas, not a hard byte-memory or end-to-end wall-clock guarantee. Parsing, copying an already supplied graph and grounded labelling are not fully preemptible by these solver checks. Term-node counts do not bound the byte length of a string. Applications exposed to untrusted oversized inputs must impose input-size limits and process resource controls; the present release does not claim that service profile.

Checked by: `cli-check-exit-status`, `cli-interrupt-is-cancelled`, `invalid-rule-is-reported`, `invalid-expected-label-is-rejected`, the reference cases `outcome/*`, the engines' `TestInferContextCancelled` and `TestExitStatus`, `cli-eval-signals-refusal`, `engine-keeps-stdout-clean`, `budget-exhaustion-is-reported`, `growing-term-is-bounded`, `malformed-statement-is-an-error`; issues E-EXIT, E-BUDGET, E-PARSE.

## 4. Negation and opposition (D10)

`¬S` is the explicit negation of `S`. It is encoded once, structurally, when a statement enters the rule store, under a functor no user statement can contain, so that it matches through statement variables and cannot collide with a user predicate such as `neg/1`. `¬¬S` is read as `S`. Language templates apply inside `¬`.

The same structural cancellation of paired explicit negations applies to imported/API statement identities, instantiated premises, conclusions and exceptions, and authored-instance validation after parameter binding. It leaves quoted text and other operators unchanged. For example, with `X=¬p`, an exception `¬X` is the statement `p`, so a stipulated `p` defeats that argument. Private API normalization preserves existing node identities on success. Two distinct declared statement objects that normalize to the same identity are refused rather than silently discarding their content; repeated assumptions alone remain set members. YAML rejects colliding normalized statement declarations. Questioned/default origins and expected labels use the same normalized identity.

Checked additionally by `TestNormalizeExplicitNegation`, `TestGeneratedNegationIdentityAndException`, `TestGeneratedNegationConclusionChains`, `TestAuthoredNegationIdentity`, `TestNormalizeAPIGraphPreservesIdentityAndAtomicity` and `TestNormalizedContextAndImport` in the required native-engine gate.

The library supplies the opposition of `S` and `¬S` as an issue scheme, so that arguments for both are weighed in one issue under its proof standard. A graph that stipulates both `S` and `¬S` is an **inconsistent context** (contract 3): it is reported and not resolved by picking a side. `ArgGraph.CheckContext()` examines raw assumptions and complete issue closure even without predeclared statement nodes or a rule firing. It reports malformed input, limits and inconsistency distinctly. `Inconsistent()` is the Boolean inconsistency query; callers needing all failures use `CheckContext()`.

Overlapping issue schemes may connect accepted positions through an existing or derived third statement. Inference checks before seeding and again after construction. Connecting two distinct pre-existing issues is rejected as invalid input rather than arbitrarily merging their proof standards. These checks operate on a private copy.

For a non-enumeration issue scheme, the first pattern selects an anchor statement and its bindings. Each remaining pattern independently identifies alternative positions under those anchor bindings. Variables first introduced in a later pattern are local to that candidate match; their values do not carry over to competing statements or another alternative pattern. Thus `[leads(X), trails(Y)]` includes every matching `trails` position for an anchor, while `[leads(X), trails(X,Y)]` retains the anchor's `X`. Repeating `Y` within one pattern still requires equal values. Reusing a fresh `Y` name in two later patterns does not turn the issue scheme into a conjunction requiring both patterns to have a common witness. This is per-anchor pattern matching, not an unrestricted multi-pattern join. Anchors and candidates are visited in sorted key order; sorting alone must not select one arbitrary fresh-variable value and omit the others. Enumeration schemes retain their existing indexed-variable rule.

Checked by: `negated-premises-match`, `negation-through-statement-variable`, `baseline-user-predicate-neg-is-ordinary`, `baseline-negation-issue`, `inconsistent-assumptions-detected`, `inconsistent-follows-current-assumptions`, `iso/negation-texts`; issues X-NEG, E-NEGTEXT, E-INCONS, E-ISSUE.

## 5. Critical questions and their states (D2)

The book's three kinds are kept (pp. 33, 373–375). An ordinary **premise** has to be supported. An **assumption** holds until it is called into question, and then has to be supported like a premise. An **exception** counts against the argument only once it is shown.

A statement of a graph can have three origins, kept apart:

* **stipulated**: listed under `assumptions:` in the graph file, or added through the API;
* **default**: listed under a scheme's `assumptions:` and instantiated by an argument;
* **supported**: concluded by an argument.

A graph can list **questioned** statements: `questioned:` in YAML, and through the API. A statement is assumed when it is stipulated, or when it is a default that is not questioned. Questioning suspends a default only: it asserts no negation and labels nothing out by itself. A questioned default can still be supported by an argument. A stipulated statement stays assumed until the stipulation is withdrawn; questioning does not withdraw it. Removing a statement from `questioned:` restores the default.

| State of an assumption | Assumed | The argument it belongs to |
|---|---|---|
| undisputed | yes | applicable if its premises hold |
| questioned, no support | no | not applicable |
| questioned, supported | through its support | applicable if the support holds |
| question withdrawn | yes | applicable if its premises hold |

| State of an exception | The argument it belongs to |
|---|---|
| not shown | applicable |
| shown (stipulated or supported) | undercut |

A contrary argument does not question an assumption by itself; it is weighed against it in an issue (contract 4). A contrary that is *assumed* is different: a default is not assumed while its negation, or another position of its issue, is stipulated, and two defaults that contradict each other are both suspended. No side is picked between defaults, as none is between stipulations (contract 3); a stipulation outranks a default. Without this, a default and its stipulated negation were both labelled in. The questioned list and the origin of each assumption survive a second `Infer`, export and import.

Issue closure and this priority are resolved before choosing the facts that seed rule inference. Otherwise an already defeated default could generate another default and incorrectly accept a downstream conclusion. The independent regression gate covers this ordering as well as direct default-versus-stipulation labels.

**Representation.** `ArgGraph.Questioned` and `questioned:` hold the questioned statements; `ArgGraph.Defaults` records which assumptions a scheme instantiation added; export writes them under `questioned:` and `defaults:`, apart from the stipulated `assumptions:`. `ArgGraph.Assumed()` lists what is assumed now.

**API updates.** Use `ArgGraph.AddAssumption(id)` to explicitly stipulate a
fact, including a fact previously introduced as a default. The operation
normalizes its identity and clears every equivalent default-origin marker;
question history and unrelated defaults are retained. A later scheme
instantiation cannot demote the stipulation back to a default. Direct writes
to the public `Assumptions` and `Defaults` fields require the caller to maintain
that origin invariant; appending a fact while retaining its default marker
does not express promotion. Native tests cover promotion before and after
inference, equivalent spellings, preserved node identity and repeated runs.

**Dialogue policy.** A request for an answer is distinct from proving an
exception. For the six explicitly supported practical/position-to-know
families, `tools/dialogue_snapshot.py` implements the separate finite
question/accepted-answer/reopening protocol in `docs/dialogue-snapshots.md`.
It targets an individual constructed argument, retains independent support
and original defeaters, and requires fresh evaluation of each snapshot.
Ignorance-family burden-holder/probandum allocations are checked separately
from the issue's proof threshold. The application supplies the appropriateness
of questions, adequacy of answers and relevance of the ultimate probandum;
the engine does not authenticate those judgments or infer them from atom names.

Checked by: `default-yields-to-stipulated-negation`, `default-yields-to-stipulated-position`, `contrary-defaults-suspend-each-other`, `questioned-assumption-is-not-assumed`, `questioned-assumption-holds-when-supported`, `questioned-stipulation-stays-assumed`, `questioned-state-survives-export`, `baseline-undisputed-assumption-holds`, `baseline-withdrawn-assumption-stays-withdrawn`, `withdrawn-assumption-survives-second-infer`, `iso/exceptions-undercut`, `iso/assumptions-questionable`; issue X-ASSUME.

## 6. Arguments in an input file that name a scheme (D12)

**A written-out argument** (premises, conclusion and a `scheme:`) is validated against its scheme with one substitution across its parameters, premises, assumptions, exceptions and conclusion. If it does not carry the scheme's critical questions, that is reported. Inference does not add a second argument with the same identity (contract 2). A written-out argument is never rewritten.

**A request to instantiate** (a `scheme:` and `parameters:`, nothing else) is carried out, giving the argument the scheme's premises, conclusion and questions, or is refused as **unsupported** until that is implemented. It is never accepted silently as an empty argument.

Implemented: for every instantiated exception, the validator requires a direct, single-premise linked argument concluding the authored argument's actual undercutter. An arbitrary nonnil undercutter, a link for another instance or another target, an extra prerequisite, a missing exception, or a custom weighing function cannot stand in for that structure. The accepted link has either no scheme or the built-in `linked` scheme, as produced by instantiation and preserved by export/import. Links and their statements must belong to the graph. A link may itself be attacked, as generated exception arguments may; validation checks the structure without deciding the outcome of such a dispute.

When parameters are omitted, premises, conclusion and exception links must all agree under one substitution, including cases with several possible premise matchings. Validation does not write inferred parameters or missing links into the authored graph. Equivalent custom networks without direct linked witnesses are outside this structural acceptance rule and are reported for the author to restate. The importer carries out a request (an argument with `scheme:` and `parameters:` only) by instantiating the scheme.

Checked by: `validator-checks-arguments-against-schemes`, `validator-enforces-shared-bindings`, `validator-accepts-correct-argument`, `explicit-argument-without-questions-is-reported`, `argument-from-scheme-and-parameters`; issues E-ARGVAL, E-EXPLICIT.

Also checked by `TestAuthoredExceptionLinksCLI` (both `check` and `eval`, with and without parameters) and the validation package's exception-link tests, including complete authored controls, ambiguous bindings and exported generated graphs.

## 7. Names the validator rejects (D11)

As predicate names, at any arity: `is`, `true`, `false`, and the internal functor of negation (contract 4). In a conclusion GoCHR executes `is` as an assignment; `true` and `false` are booleans to the engine's validator. Other operator names (`in`, `or`, `div`, `mod`) work as predicates in premises and conclusions and stay allowed.

Checked by: `static/no-reserved-functor`, `reserved-predicate-names-are-rejected`, `engine-validator`; issues X-IS, L-TRUE.

## Interchange (Phase 4)

A graph exported and read back in a new process has the same theory, language, schemes (with their weights and roles), weighing functions, parameters, expected labels (under `tests:`), questioned statements, defaults and stipulations, and recomputes the same labels. Checked by: `yaml-export-can-be-reimported`, `export-keeps-compound-parameters`, `export-keeps-expected-labels`, `cli-reads-its-own-yaml`, `cli-output-stays-valid`, `scheme-roles-are-imported`; issues E-EXPORT, E-ROLES.
