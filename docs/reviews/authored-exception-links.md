# Independent runtime review: authored arguments bypassing exceptions

**Baseline:** repository `49ad425cbf7afeebcbd9a07ee479d60681ca33ea`; the unrelated Atkinson review commit does not change this engine behavior. **Finding:** confirmed. **Disposition:** structural validator repaired; authored graphs remain unchanged; new engine and fresh-process CLI regressions added.

## Reproduction and impact

A scheme has premise `p(X)`, exception `e(X)` and conclusion `q(X)`. The input contains an authored argument identifying that scheme, with premise `p(a)`, conclusion `q(a)`, and an undercutter pointing to an unrelated statement named `inert`. The graph stipulates both `p(a)` and `e(a)`.

The baseline validator only checked that the undercutter pointer was nonnil. It did not check whether the scheme's instantiated exception could support that undercutter. With explicit parameters `[a]`, inference recognizes the authored argument's scheme/parameter identity and does not construct the normal argument and its exception links. Without written parameters, a generated argument can coexist with the authored one, but the authored one still supplies unchallenged support. Both inputs incorrectly reached the same result:

| Input | `check` | `eval -t yaml` | Label of `q(a)` |
|---|---:|---:|---|
| Explicit `[a]` parameters | exit 0, no diagnostic | exit 0 | `in` |
| Parameters omitted | exit 0, no diagnostic | exit 0 | `in` |

This is an executable counterexample to the previous claim that a written-out argument naming a scheme necessarily carries its exceptions. It does not depend on the wording of a Walton scheme or the new Atkinson fixtures. The before-run inputs and exact outputs were retained with the review evidence.

## Repair and accepted structure

For each exception of a scheme, the validator now applies the same substitution used by the authored argument and requires a direct argument from that ground exception statement to the authored argument's actual undercutter. The supporting link must have exactly one premise and use the built-in linked weighing rule, either implicitly or explicitly. Both the link and its premise/conclusion objects must belong to the graph.

The validator rejects an inert undercutter, wrong-instance exception, link to a different undercutter, extra prerequisite, omitted exception, custom disabled weight or dangling link. It checks every exception independently. A complete generated graph already has this structure and continues to validate, including its exported and reimported form.

For arguments without explicit parameters, validation searches the possible premise/conclusion substitutions and requires the exception links to match one and the same substitution. It does not take the first symmetric premise assignment and falsely reject a valid later assignment; nor may different exceptions use incompatible assignments. The validator does not fill in parameters, add arguments or rewrite any authored graph.

This is a deliberately structural acceptance rule, now stated in contract 6. An arbitrary custom network might be logically equivalent to the direct links, but the validator does not prove that equivalence. Authors can supply the direct linked witnesses or request scheme instantiation. Exception arguments can still have their own undercutters: the engine's generated structure permits higher-order dispute. A valid link's presence does not guarantee that it wins every dispute.

## Regression evidence

The validation-package tests exercise seven shapes with both explicit and inferred parameters: inert undercutter, complete direct link, wrong instance, wrong target, extra prerequisite, missing second exception and complete two-exception structure. Further tests cover zero-weight custom links, links absent from the graph, backtracking with symmetric premises and incompatible exception bindings, positive/adverse complete authored arguments, and generated multi-conclusion structures with compound parameters and explicit negation surviving export/import.

Four named V&V checks execute the original counterexample through separate CLI processes: `check` and `eval`, each with and without explicit parameters. The corrected behavior is invalid input, exit 1, an exception-link diagnostic and empty stdout. Both an error status and absence of an evaluation are checked.

The focused validation run used Go 1.27.4. It passed all new engine tests and vet. The registered focused run passed **244 checks out of 244**: the four new CLI checks plus all 120 reference cases and their fresh-process round trips. There were zero known failures, unexpected failures or stale ledger entries. A direct rerun of both original reproducers independently confirmed exit 1 and zero stdout for both commands. The main reviewer owns the complete integration rerun, both toolchains and the combined named-check recount. This fix adds four registered checks beyond the separate Atkinson patch and does not alter the scheme library.

## Scope

The repaired layer is `validation.Validate` and the CLI's existing validation gate. This change does not make a bare direct call to `ArgGraph.Infer` run all external validator checks; that public-API contract remains subject to the main runtime review. A caller bypassing validation must not treat this patch as a new automatic input gate inside `Infer`.

The finding invalidates automatic certification of an older preview pin based solely on the previous green suite. The final release must refer to the commit containing the repair and must pass the integrated gate. The original reproducible passing count remains historical evidence about the checks that existed then, not evidence that this counterexample was already covered.
