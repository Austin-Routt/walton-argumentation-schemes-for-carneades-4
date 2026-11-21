# Independent review of PR16 engine and interrupt findings

Reviewer: ChatGPT fresh non-author acceptance reviewer, 2026-11-21.
F2 implementation reviewed: `d86c1df987da708c556c84c857205fcee548f930`.
Baseline comparison: `2e7227352d1656d19d80a16de23c9d4ab27faa2f`.

**Decision:** accept the F2 implementation for the newly explicit per-anchor
issue-pattern contract. Accept the F4 workload replacement, with the timing
and work-count qualifications below. The final integrated candidate still
requires new clean full runs; the earlier 9,959-check matrix is pre-PR16
evidence and does not establish that these newly identified cases worked.

## F2: fresh bindings must not exclude alternative positions

For a non-enumeration issue scheme, the first pattern identifies an anchor
and binds its variables. Each remaining pattern independently describes
alternative positions under those anchor bindings. A new variable introduced
by a later pattern is local to that particular candidate match. It is not a
choice of one witness that constrains all competing statements or subsequent
alternative patterns.

Consequently, `[leads(X), trails(Y)]` must include every matching `trails`
position for an anchor. `[leads(X), trails(X,Y)]` must retain the anchor's
`X`. Repeating `Y` within one pattern still constrains the two occurrences to
the same value. Reusing a fresh variable's name in two alternative patterns
does not turn those alternatives into a conjunction with a common witness.
The new contract text states these distinctions explicitly.

The previous implementation retained the environment returned by a successful
candidate match. This accidentally selected one fresh-variable value and
omitted other valid positions. Sorting by itself would merely make that
arbitrary omission repeatable. The repair both sorts traversal and discards
each candidate's extension of the anchor environment. The term matcher uses
persistent binding-list extensions; discarding that return value leaves the
anchor's original bindings intact. Failed alternatives do not change them.

The grounded lookup path still applies the typed formula match. Enumeration
keeps its existing indexed-variable rule. The repair does not invent a general
multi-pattern join or change rule-engine conjunction matching. Existing
overlap, consistency and atomic-publication rules remain in effect.

I independently ran the exact same regression file against a separate
baseline clone and the repaired clone. All four position-set controls failed
on the baseline: fresh variables, preserved anchor bindings, fresh names
shared across alternative patterns, and repeated variables within one pattern.
The baseline's consistency control happened to pass in that particular run;
I do not describe it as a reproduced failure on every execution.

On the repair, all four controls and the consistency/refusal control pass on
both Go 1.24.13 and Go 1.27.4. They cover indexed and scan paths, multiple
insertion orders, exact position sets, repeated inference and every selected
stipulated alternative. An inconsistent stipulated pair is rejected with
`ErrInconsistentContext`, and no partial issue is published. Code inspection
confirms that inference still builds on a private graph and publishes only on
success; the new traversal and matching do not bypass that transaction.

The regression file is byte-identical in the two clones, SHA256
`0c4c2866df7ffb2031e90b8b974aee841793003d72774c832fca8a9ec4cc4d94`.
The baseline was augmented only with that test file; the repaired clone was
clean after execution. Logs are retained in `pr16-fresh-audit`:
`f2-before-go1.24.13.log`, `f2-fixed-go1.24.13.log` and
`f2-fixed-go1.27.4.log`.

## F4: cancellation fixture must not finish at a quota first

The old pair-producing workload is a valid limit-reaching computation. In my
independent run it ended at the 100,000 inserted-constraint quota with exit 4
after approximately 0.819 seconds. A test that assumes this workload always
remains active until a delayed interrupt can race a legitimate termination.
This is a fixture defect, not evidence that the engine misclassifies an
interrupt received while work is still active.

The replacement joins three `p` facts with a nonexistent `q(X,Y,Z)` fact.
With 300 `p` facts there are 27 million candidate triples, plus prefix and
matching work, and no successful argument. The public comment's phrase
"seconds of work on any machine" is not justified; duration depends on the
machine, engine and compiler. "27 million candidate visits" is also less
precise than "27 million candidate triples". Comments should describe current
quotas and measured runs, not a universal wall-clock guarantee.

I built separate CLIs from the reviewed F2 source with both pinned toolchains.
The new workload completed uninterrupted on Go 1.24.13 in approximately
6.560 seconds with exit 0, demonstrating completion within the current quotas
in that run. Interrupted after 0.3 seconds, it returned exit 5, no stdout and a
cancellation diagnostic on both Go versions, in approximately 0.303 and 0.302
seconds. Thus the replacement supplies a substantial measured interruption
window without weakening the required cancellation result.

These executions, fixture hashes, actual command paths, signal times, exits
and output sizes are recorded in `pr16-fresh-audit/f4-independent-runs.json`,
with the complete stdout/stderr and fixture files alongside it. This is
focused independent acceptance of F2 and the F4 workload; it grants no source
approval and does not replace the subsequent exact-head full matrix.
