# Independent acceptance of implicit linked-argument observations

Reviewer: ChatGPT fresh non-author acceptance reviewer, 2026-11-21.
Finding: **FRESH-NULL-SCHEME-OBSERVATION**.
Reviewed repair: `df64bf46d64f1394488b9f614a2cfa1007050c58`.

**Decision:** accept the observer repair. Final combined acceptance remains
subject to the frozen commit's complete clean runs on both Go versions.

Generated exception arguments can have a nil `Scheme` pointer and use the
engine's default linked interpretation. The CLI observer already treated
these as `linked`, but the API observer skipped them. My independent example
used a rule from `p` to `q`, with exception `e`, stipulated `p`, and deliberately
forbade any `linked` argument. Before the fix, the API observation incorrectly
passed; the CLI observation correctly reported the generated exception link.

This was an API-oracle false pass, not an incorrect inference by the engine.
The combined API-plus-CLI case test still failed because its CLI component
caught the forbidden argument. The finding therefore does not show that such
a complete paired case silently passed. It corrects a real discrepancy between
the two observation paths.

The integration reviewer authored the repair. The API observer now records a
non-nil argument with a nil scheme as `linked`, using the same argument and
parameter observation routine as named schemes and the CLI path. I reviewed
the committed change, then ran my unchanged original probe and the author's
positive and forbidden-linked controls independently on Go 1.24.13 and
Go 1.27.4. Both commands passed. The original probe now observes the expected
violation in both API and CLI outputs; the positive control confirms that a
valid expectation of `linked()` succeeds.

Evidence retained separately:

- `fresh-audit/generated_link_observation_probe.go` — unchanged probe.
- `fresh-audit/generated-link-observation-before.log` — API false pass before
  repair, while the CLI correctly detected the forbidden linked argument.
- `fresh-audit/generated-link-observation-fixed-go1.24.13.log` and
  `fresh-audit/generated-link-observation-fixed-go1.27.4.log` — passing
  independent probe and positive/negative controls after repair.

The runs used the pinned toolchains, `GOTOOLCHAIN=local`, `GOFLAGS=-mod=mod`
and `go test -count=1`. I removed the temporary probe after checking it was
byte-identical to the preserved original. No implementation or source record
was changed by this reviewer. Further acceptance work is the required full
clean matrix, not an unbounded search for a proof that no defect can exist.
