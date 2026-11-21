# Independent acceptance of explicit default promotion

Reviewer: ChatGPT fresh non-author acceptance reviewer, 2026-11-21.
Finding: **FRESH-DEFAULT-ORIGIN**.
Reviewed repair: `e6e4015eab2198d808e7b1c7aa817aff681c8fe9`.

**Decision:** accept this repair of the public-API origin contract and the
corresponding case-preparation paths. The complete combined-release gate
remains pending until the frozen commit passes both clean toolchain runs.

## Independent finding

While checking the boundary between defaults and stipulations, I found that
the public `AddAssumption` method could not promote an existing default.
Importing `defaults: [p]` and `questioned: [p]`, then explicitly adding `p`
through that method, left `p` out after another inference. Adding the equivalent
`¬¬p` failed too. The default marker remained, so the explicit stipulation was
still suppressed by questioning. A case's additional assumptions had the same
problem when its graph already contained that default.

The evidence is the original probe `fresh-audit/promote_default_probe.go` and
its failing run `fresh-audit/promote-default-before.log`. This finding is an
engine/API origin defect. It is distinct from FRESH-HARNESS-CONTEXT, the loss
of existing case context through duplicate YAML keys.

## Review of the author's repair

The integration reviewer authored the implementation; I did not modify it.
The public operation now normalizes the formula identity, removes all matching
default markers and appends the fact as a stipulation. Removing matching
aliases matters before the first inference: a raw API map can still contain a
spaced or doubly negated spelling. Unrelated defaults and the questioning
history remain intact.

A separate private append operation allows scheme instantiation to create a
default without passing through the public stipulating operation. This keeps
the distinction between the two origins. Repeated inference does not demote
an explicitly stipulated fact back to a default. Both ordinary case imports
and premise-withdrawal variants now use the stipulating API for supplied case
facts, instead of appending to the raw assumptions field.

API clients should use `AddAssumption` when explicitly accepting a fact,
especially when it previously appeared as a default. Clients that mutate the
public assumptions/default fields directly must maintain that origin invariant
themselves. The method is not a substitute for structural validation or
inconsistency checking before accepting an evaluation.

## Independent verification

I fetched the committed repair into the separate review clone and reran my
**unchanged** original probe under Go 1.24.13 and Go 1.27.4. The probe's SHA256
is `bca5a74ca51abf3c72917f527e05f101bce03e007d02a12aed11943c688ebb02`.
The before-fix failure and after-fix successes are retained separately.

Both toolchains also passed the author's native controls for all 18
combinations of evaluation state, marker spelling and added-fact spelling;
preservation of an unrelated default and existing statement pointer; a
scheme-generated default; and repeated inference after promotion. The harness
controls passed for canonical and doubly negated facts through the API, CLI
round trips and premise withdrawals. Existing context-preservation tests and
all seven adverse expectation classes remained green.

The runs used `GOTOOLCHAIN=local`, `GOFLAGS=-mod=mod`, the pinned toolchain at
the front of `PATH`, and `go test -count=1`. Logs:

- `fresh-audit/promote-default-fixed-go1.24.13.log`.
- `fresh-audit/promote-default-fixed-go1.27.4.log`.

The temporary probe was removed after its byte identity was checked against
the preserved original. The clone was clean before this documentation-only
acceptance change. No source record or source approval was changed. These
focused results resolve the reproduced origin defect; they do not replace
the required final engine/race and strict V&V gates on the frozen candidate.
