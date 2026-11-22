# Independent cross-review of the final AATS adapter amendments

Date: 2026-11-22 UTC. Base repository:
`98a2335d2c9909cf597ae85a9f359f20d092ae6c`.
Implementation author: the separate `final_atkinson` reviewer.
Independent code reviewer: ChatGPT `final_sources`.

## Decision

**Accept the final adapter amendments and their regression evidence within
the declared finite, goal-labelled AATS interpretation.** I did not author
the adapter implementation, its Python regression changes, or its actual
engine Go tests. I read those changes, constructed additional adverse
inputs, reported three input-boundary problems to the author, and
independently reran the repaired results. This is model-assisted technical
independence, not authenticated human certification.

The acceptance covers the exact file hashes below. It does not turn the
project's action-comparison policy into a uniquely derived algorithm from
the paper, a general VAF evaluator, or an exhaustive verification of every
possible model.

## Reviewed behavior

**Input shape and duplicate keys.** The adapter now uses the existing
duplicate-aware safe YAML loader. Duplicate literal keys are rejected
with their line numbers; intentional YAML merges with explicit overrides
remain usable by both supplied model files. Mapping/list/atom checks
prevent a scalar proposition string from being treated as its characters
and prevent malformed transition/effect objects from being silently
interpreted as valid model components. Parser and filesystem errors are
reported as `ModelError`. These checks are finite-model input checks,
not a claim of arbitrary hostile-input resource limits.

**Effects belong to the state pair.** The emitted predicates carry source
and result state, not a joint-action identifier. Therefore two transitions
from the same state to the same result must carry the same set of
goal-labelled value effects. The new validation enforces that invariant
and independently rejects duplicate effect entries. Ordering of the effect
list is immaterial. This prevents another action from acquiring an
unlisted promotion merely because its transition reaches the same state.
Equal state-pair effects remain valid, and the actual engine test expects
both actions' positive and negative arguments explicitly.

**Equal proposition sets do not identify states.** The paper defines an
interpretation function, not an injective function. Distinct state IDs
may therefore have the same currently true propositions. The adapter
retains those IDs. A proposition-list description matching more than one
state now requires an explicit ID rather than choosing the first state
silently. The actual-engine fixture contains equal-interpretation state
IDs, and the Python controls inspect both ID lookup and ambiguous-list
refusal.

**Unnamed joint tuples.** A well-formed tuple containing one declared
action per declared agent can belong to the Cartesian product even when
the input assigns it no name or transition. In that case the problem is
the missing claimed transition, CQ2, rather than an invalid tuple, CQ13.
The new code preserves that distinction. My independent positive/control
probes produced no questions for the valid named claim, CQ13 for a
well-formed unknown joint name, and exactly CQ2 for the valid unnamed
tuple used in the regression.

**Fact-based oracle.** The new Python argument reconstruction joins the
emitted facts and does not read the adapter's predicted argument list.
That removes a circular dependency from the previous differential
comparison. The Go boundary tests also read only the emitted facts and
use separately authored expectations against the actual current library
and vendored engine. The assertions cover exact positive/negative
argument construction and the equal-weight opposition behavior, not
merely successful parsing.

## Additional counterexamples found and repaired during this review

1. **Unknown transition joint caused a raw exception.** Changing a
   farmer transition's `by` field to `unknown_joint` was detected by the
   validator, but the final option-existence check subsequently indexed
   the invalid joint and raised `KeyError`. The author replaced that
   cascading lookup with a guarded check. The Model API now raises
   `ModelError`; the CLI exits 2, emits a useful error and no traceback,
   and emits no result graph.
2. **Misspelled audience silently selected every action.** The CLI
   previously treated `--audience definitely_unknown` as a value order,
   filtered its undeclared value away and reported all four farmer
   actions as equally good. The author now validates named and ad-hoc
   orders against the declared model value universe and rejects unknown,
   duplicate, empty or empty-component orders. Both API and CLI reject
   the counterexample. Intentional valid subsets remain valid, and a
   declared value held only by another modeled agent can still produce
   the explicitly intended tie when no relevant value of the deciding
   agent remains. This is distinct from silently accepting a typo.
3. **Malformed claim shapes leaked implementation exceptions.** A list
   instead of a claim, a list instead of a joint name/tuple, and nested
   unhashable proposition entries previously raised `AttributeError` or
   `TypeError`. The author added explicit field and type guards. Both
   incomplete malformed inputs and complete claims with those malformed
   values now produce `ModelError`. Well-formed unknown propositions or
   joint names retain their appropriate critical-question diagnostics.

The author made these repairs; this reviewer made no changes to the
three implementation/test files listed below. The final adverse results
are retained in `aats-cross-probes.json`.

## Independent execution

| Check | Result |
|---|---|
| `python tools/test_aats.py` | 71 tests pass; the suite reports 2 models, 17 situations and 32 recorded decision checks |
| Fresh unknown-joint Model/CLI controls | Typed refusal; CLI exit 2; no traceback or successful result output |
| Fresh unknown-audience CLI control | Typed refusal; CLI exit 2; no choice output |
| Fresh malformed-claim controls | Typed refusal for all reported shapes |
| Fresh valid/named/unnamed joint controls | Valid claim: `[]`; unknown name: `[13]`; valid unnamed tuple: `[2]` |
| Valid subset and other-agent value controls | Accepted and retain the declared comparison behavior |
| Actual-engine adapter regressions, Go 1.24.13 | Both tests pass |
| Actual-engine adapter regressions, Go 1.27.4 | Both tests pass |

The actual-engine command was `go test -count=1 -run '^TestAATS' -v
./internal/vvlib` from `tools/vv`, using each pinned compiler with
`GOWORK=off`, `GOTOOLCHAIN=local` and `GOFLAGS=-buildvcs=false`.
`aats-cross-execution.json` and the two accompanying Go logs contain
the exact commands and outputs. The full project gate and clean-copy
cross-version comparison remain the integration reviewer's separate
acceptance evidence.

## Exact accepted files

| File | SHA-256 |
|---|---|
| `tools/aats.py` | `d9f316f44cbaa02b79119555f26dc6ef077d530e4a0c80d9f01cc5c5c987379f` |
| `tools/test_aats.py` | `d4acc0190f9f04acfc2ad16781fe0354f28ecb5ab2aac339641443a75b7fcaf9` |
| `tools/vv/internal/vvlib/aats_adapter_test.go` | `fccc7e1a7c94cb4fb8c804a8970fe985b2cef2ec4ea3d13611715594434f616c` |

No concrete defect from this cross-review remains unresolved in those
files. Further empirical adequacy and capacity evaluation belongs to the
planned validation/benchmark program; it is not represented by this
bounded implementation acceptance.
