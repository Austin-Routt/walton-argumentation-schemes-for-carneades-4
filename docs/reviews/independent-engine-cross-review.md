# Independent cross-review of the engine repair

Reviewed the parent engine change `e819a63655acff2d25ff65dedb3c9c24fdff74cf` and the subsequent live context-validation work in `release-review/fixes`. This is a read-only cross-review; the source repair commit `bc0870880d2a150691d90cad856e52c32c090930` contains no engine edits. The observations below identify the specific implementation state exercised; the parent's final integration report is authoritative for later fixes.

## Streaming matching and logical reuse

The streaming matcher preserves the finite propagation closure for the tested supported profile. `candidateView` fixes a cutoff at the start of a rule's turn, so a body cannot add candidates to that same traversal. New constraints enter later rules or a subsequent round, as in the earlier materialized-tuple implementation. The earliest-new-head partition visits each ordered tuple in one branch; rule/tuple history prevents another firing. A head's bindings are persistent linked structures, so backtracking does not overwrite a sibling's bindings.

The explicit `ReuseConstraints` flag is the right separation. Carneades treats premises as logical propositions, for which P and P can use the same true fact. Ordinary CHR heads normally require distinct stored constraint identities. The flag is enabled for Carneades and has effect only for the pure Unique propagation solver. The generic CHR behavior and the pure solver's default retain distinctness.

Six independent manually specified closure controls passed, in `streaming_probe.jsonl`:

1. Two facts produce all four ordered pairs when logical reuse is enabled.
2. The default CHR reading produces only the two distinct pairs.
3. A three-node directed cycle reaches all nine ordered pairs under transitive closure.
4. A bound bare-variable head placed before its binding conditional correctly closes a proposition cycle.
5. A multi-head join backtracks correctly and permits one qualifying fact to fulfill both logical roles.
6. The same join without reuse excludes that repeated-identity case.

I found no new streaming correctness blocker in this scope. These finite controls complement the implementation review; they are not a universal proof of every CHR program or a certification of guards/deletions outside Carneades' supported profile.

## Resource and cancellation semantics

The new candidate-visit budget addresses unsuccessful joins as well as successful applications. Streaming also prevents constructing the full combinatorial list before checking the application limit. Constraints and term-node quotas count after duplicate suppression. These are substantial corrections to the previous practical resource bounds.

They remain logical quotas, not a process memory limit or an end-to-end deadline. Parsing the full goals list precedes store quotas. A huge atom/string can occupy one logical term node. Candidate list construction, issue matching, rendering and graph publication occur outside the match-step budget. Issue matching has quadratic scans and, in the inspected implementation, no context polling. The commit loop deliberately has a cancellation check before it begins and no midway rollback. Documentation should describe a pre-commit cancellation boundary and require input-size/process limits for untrusted or large inputs; it should not equate the propagation budget with total wall-clock or byte bounds.

The extra context check around acquiring the engine semaphore correctly removes the select race for an already-cancelled context. Returned errors before publication leave the graph untouched. Any error possible after publication would still need a transaction or advance validation; the new issue-pattern parse preflight is an appropriate way to eliminate the known late parse error.

## Context validation: independent findings and controls

The first issue preflight copied only effective assumptions and lost default origin. That would have treated an unmaterialized default/contrary pair as two stipulations. The parent independently corrected it while this review was running. The latest `context.go` copies raw assumptions, default metadata and questioned state onto a private graph. `preflight_probe.jsonl` confirms:

- Two stipulated positions are rejected with no statements/issues added.
- A default and stipulated contrary succeed; only the stipulation remains effective.
- Two contrary defaults succeed with both suspended.

Two additional bounded gaps were reproduced and sent to the parent:

**Malformed fact in an empty theory.** `Assumptions = ["p("]` with no argument schemes returns nil and materializes the malformed string. `ArithmeticProblem` ignores parse failure, and the no-rule path never invokes the GoCHR parser. A common validation gate should reject malformed accepted terms before mutation. See `preflight_probe.go/jsonl`.

**Connecting statement across multiple issue schemes.** With existing statement identities `p(a)`, `r(a)`, `q(a)`, assumptions `p(a)`, `q(a)`, and issue patterns `[p(X),r(X)]`, `[r(X),q(X)]`, the inspected per-scheme preflight sees no conflict. `Infer` returns nil; final materialization merges all three positions, after which `CheckContext` reports inconsistency. The existing but unassumed `r(a)` is the bridge omitted from the private probes. A shared private graph containing all existing statement identities and applying all patterns before checking consistency addresses that case. A bridge derived during inference requires validation on a prospective completed graph or an explicitly rejected overlap profile. See `overlap_probe.go/jsonl`.

These are precise acceptance boundaries. They should be corrected or declared outside the supported profile; a green single-issue case does not establish complete multi-issue atomicity.

## Consequence question advice

The supplied book's pp. 332–333 prints the shared critical questions for the consequence schemes, so a book-supported CQ1 question hook can be added without claiming to have inspected the separately cited originals. Source attribution and original-source verification must remain distinct.

CQ1 asks how strong the consequence likelihood is; CQ2 asks whether evidence supports that claimed strength; CQ3 asks about opposite consequences. The existing `outcome_likely/2` template says “strongly likely,” which imposes a high-likelihood reading not entailed by a general how-strong question. Low-probability severe consequences may still furnish a prudential reason. Therefore do not silently collapse causal uncertainty, evidential support and a decision threshold into one Boolean called likely.

For a qualitative preview, an explicit challengeable assumption that the stated likelihood assessment is adequate for this case can represent CQ1, with probability estimation and adequacy policy external. CQ2 should remain a challenge to evidence for the causal claim, and CQ3 a competing consequence argument. A calibrated quantitative claim needs parameters for the asserted degree/strength and an accepted policy; the compendium supplies no universal numerical cutoff. This distinction is a scope decision, not a reason to pretend the CQ text is inaccessible.

## Final transaction review: commit 4f2dd85

The parent resolved the three precise context defects found during this cross-review by staging inference on copied mutable topology, preserving default/question state, materializing the complete context before choosing CHR seeds, bounding and polling issue matching, and validating the final graph before publication. The copy preserves graph cycles through memoization; the commit remaps new links onto retained pre-existing statement, argument and issue objects.

I independently reran the original probes against commit `4f2dd85`:

- Malformed input in an empty theory now returns invalid input and adds no statement.
- A connecting pre-existing statement now makes context validation reject both accepted endpoints, with no issue committed.
- A defeated default no longer seeds a secondary default or its downstream conclusion; the graph contains no resulting arguments.
- One default versus a stipulation still resolves in favor of the stipulation; two defaults suspend together.

I also added two independent transaction controls in `transaction_probe.go`. A bridge produced only during rule inference now triggers an inconsistent-context error and leaves the original statement pointers, text, argument links, map alias and graph cardinalities unchanged. A successful multi-conclusion inference preserves existing Statement, Argument and Issue objects, extends the issue to a new position, and preserves canonical premise/conclusion pointers and both directions of the issue links. Both controls pass.

The earlier *_probe.jsonl files preserve the discovered failures; corresponding *_fixed.jsonl and `transaction_probe.jsonl` establish their resolution. I approve these inspected transaction/context repairs for the supported cases, subject to the parent's complete suite and issue-budget stress acceptance. The final claim is still bounded: cloned graph/input bytes and arbitrary string payloads are not bounded by logical term-node quotas, and the final commit has an intentional cancellation boundary.
