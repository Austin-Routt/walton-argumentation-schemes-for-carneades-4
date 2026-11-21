# Contracts

What the engine in `third_party/` and the scheme library promise once the plan's phases 1 to 4 are done (PLAN.md, item 0.5). Each contract names the checks of `tools/vv` that test it. Where today's engine falls short, the check is in `tools/vv/known_issues.yml` under the issue named here; a contract is met when its issues have left the ledger.

A contract changes only together with its checks, in one reviewed change.

## 1. The supported profile (D3)

**Entry points.** The `carneades` command (`check`, `eval`) and the Go API (`yaml.Import`, `ArgGraph.Infer`, `ArgGraph.GroundedLabelling`, `validation.Validate`), offline, one graph at a time in one process. Concurrent use and the web server are outside the profile until they have their own tests (issue E-GLOBALS covers what sequential use needs).

**Schemes.** A scheme is a propagation rule: its premises, together with no deletions and no guards, yield its conclusions, assumptions and exceptions as instantiated statements. A scheme with `deletions` or `guards` is outside the profile and is reported as **unsupported** (contract 3), not run. Every variable of a scheme occurs in a premise (range restriction).

**Terms.** A constant is a lower-case name, an integer or a quoted string; a capitalised name is a variable. A statement carries no arithmetic: an unquoted `1990-05-17` is invalid input, not the number 1968 (issue E-TERMS). A statement may be a variable of a scheme (a statement variable), as in `expert_opinion`.

**Facts.** The ground statements of a graph form a set (contract 2).

Checked by: `iso/*`, `full/*`, `cli/*`, `static/range-restricted`, `string-constants-survive`, `arithmetic-in-a-statement-is-invalid-input`, `deletions-and-guards-are-unsupported`, `yaml-import-keeps-globals-clean`.

## 2. One fact, one argument (D8)

**A fact** is a ground statement, identified by its canonical term: functor, arity and arguments compared structurally after normalization, so `p(a,b)` and `p(a, b)` are one fact and `p(a)` and `p(a,b)` are two. Adding a fact that is already present changes nothing. GoCHR keeps its multiset semantics for its own programs; the set semantics are switched on by Carneades for argument construction (PLAN.md 2.2).

**An argument** is identified by its scheme and its canonical parameter list. Two arguments for one conclusion from different schemes, or from one scheme with different parameters, are two arguments and both are kept. Running inference again on an unchanged graph adds nothing and leaves every label as it was.

**Termination.** Set semantics end the loops of rules that re-derive a fact. They do not bound rules that build ever larger terms, or matching work; the limits of contract 3 do.

Checked by: `rederivation-terminates`, `no-scheme-is-starved`, `infer-is-idempotent`, `infer-twice-with-exception`, `wrong-arity-fact-is-harmless`, `growing-term-is-bounded`; issues X-LOOP, E-IDEMP, E-MATCH, E-BUDGET.

## 3. How a run ends (D9)

A run of `check`, `eval` or `Infer` ends in exactly one of seven ways.

| Outcome | Meaning | Exit status of the command |
|---|---|---|
| complete | the graph was read, validated, inferred and labelled | 0 |
| invalid input | the file cannot be read, a statement cannot be parsed, or validation reports a problem | 1 |
| unsupported | the input uses something outside the profile of contract 1 | 2 |
| inconsistent context | the graph stipulates both a statement and its negation, or two positions of one issue (contract 4) | 3 |
| limit reached | a limit of rule applications, term size or time ran out while work remained | 4 |
| cancelled | the caller stopped the run | 5 |
| internal failure | a panic or a broken invariant of the engine | 6 |

**Through the API** `Infer` returns `nil` only for complete; otherwise an error whose kind is one of the six others, testable with `errors.Is`. On any error the graph is left as it was before the call: a partial graph is never returned as a result.

**Output.** The engine writes nothing to stdout; diagnostics go to stderr. On any outcome other than complete the command prints no evaluation.

**Limits.** The rule-application limit stays 100,000 by default (`caes.MAXRULEAPPS`). A run counts as interrupted only if a rule could still fire when the limit was reached. A limit on the size of derived terms, 500 nodes (`caes.MAXTERMSIZE`), ends runs whose terms grow without end; the library derives none larger than 30. A caller's context ends a run from outside (`ArgGraph.InferContext`). Implemented: GoCHR's `RuleStore.MaxTermSize`, `Done` and `Stopped`, with `ErrLimit` and `ErrCancelled`.

Checked by: `cli-check-exit-status`, `cli-eval-signals-refusal`, `engine-keeps-stdout-clean`, `budget-exhaustion-is-reported`, `growing-term-is-bounded`, `malformed-statement-is-an-error`; issues E-EXIT, E-BUDGET, E-PARSE.

## 4. Negation and opposition (D10)

`¬S` is the explicit negation of `S`. It is encoded once, structurally, when a statement enters the rule store, under a functor no user statement can contain, so that it matches through statement variables and cannot collide with a user predicate such as `neg/1`. `¬¬S` is read as `S`. Language templates apply inside `¬`.

The library supplies the opposition of `S` and `¬S` as an issue scheme, so that arguments for both are weighed in one issue under its proof standard. A graph that stipulates both `S` and `¬S` is an **inconsistent context** (contract 3): it is reported and not resolved by picking a side. `ArgGraph.Inconsistent()` answers from the current assumptions.

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

A contrary argument does not question an assumption by itself; it is weighed against it in an issue (contract 4). The questioned list and the origin of each assumption survive a second `Infer`, export and import.

**Representation.** `ArgGraph.Questioned` and `questioned:` hold the questioned statements; `ArgGraph.Defaults` records which assumptions a scheme instantiation added; export writes them under `questioned:` and `defaults:`, apart from the stipulated `assumptions:`. `ArgGraph.Assumed()` lists what is assumed now.

Checked by: `questioned-assumption-is-not-assumed`, `questioned-assumption-holds-when-supported`, `questioned-stipulation-stays-assumed`, `questioned-state-survives-export`, `baseline-undisputed-assumption-holds`, `baseline-withdrawn-assumption-stays-withdrawn`, `withdrawn-assumption-survives-second-infer`, `iso/exceptions-undercut`, `iso/assumptions-questionable`; issue X-ASSUME.

## 6. Arguments in an input file that name a scheme (D12)

**A written-out argument** (premises, conclusion and a `scheme:`) is validated against its scheme with one substitution across its parameters, premises, assumptions, exceptions and conclusion. If it does not carry the scheme's critical questions, that is reported. Inference does not add a second argument with the same identity (contract 2). A written-out argument is never rewritten.

**A request to instantiate** (a `scheme:` and `parameters:`, nothing else) is carried out, giving the argument the scheme's premises, conclusion and questions, or is refused as **unsupported** until that is implemented. It is never accepted silently as an empty argument.

Checked by: `validator-checks-arguments-against-schemes`, `validator-enforces-shared-bindings`, `validator-accepts-correct-argument`, `explicit-argument-gets-critical-questions`, `argument-from-scheme-and-parameters`; issues E-ARGVAL, E-EXPLICIT.

## 7. Names the validator rejects (D11)

As predicate names, at any arity: `is`, `true`, `false`, and the internal functor of negation (contract 4). In a conclusion GoCHR executes `is` as an assignment; `true` and `false` are booleans to the engine's validator. Other operator names (`in`, `or`, `div`, `mod`) work as predicates in premises and conclusions and stay allowed.

Checked by: `static/no-reserved-functor`, `is-as-predicate-in-conclusion`; issues X-IS, L-TRUE.

## Interchange (Phase 4)

A graph exported and read back in a new process has the same theory, language, schemes, parameters, expected labels, questioned statements and stipulations, and recomputes the same labels. Checked by: `yaml-export-can-be-reimported`, `export-keeps-compound-parameters`, `export-keeps-expected-labels`, `cli-reads-its-own-yaml`, `cli-output-stays-valid`, `scheme-roles-are-imported`; issues E-EXPORT, E-ROLES.
