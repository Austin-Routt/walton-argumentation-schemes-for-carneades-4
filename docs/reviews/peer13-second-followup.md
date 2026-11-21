# PR13 follow-up: perception, memory and shared wording

Prepared from independent comparison of remote PR13
`96daf6c0545f7cb84fe5e30cfbffdde5d6f6c0d1` with the integrated candidate
`a708add511f52fb3603d92ba368875ef3d3f2c42`. This is a bounded amendment to the
entries 28–60 review, not a wholesale adoption of PR13's 90 source approvals.

## Accepted correction

The supplied compendium's perception and memory schemes, printed pp. 345–346,
both conclude that belief in the same proposition or case state is reasonable.
The earlier independent review correctly repaired memory's recaller scope but
missed the unnecessary split between `reasonable_case/1` and
`reasonable_believe/1`. That omission is now acknowledged and corrected.

`perception` now uses memory's existing `reasonable_believe/1`. The particular
case belongs in the ground proposition: `door_closed(case1)` is distinct from
`door_closed(case2)`. Neither scheme is allowed to make two unrelated events
identical merely because their description uses the same property word.

The five new cases in `tools/vv/cases/independent_peer13.yml` require two actual
arguments for one proposition, surviving support when either one is defeated,
loss of support when both are defeated, and failure of another case's memory
to rescue a defeated perception. The separate observer/context weakness of
perception's `not_indicator(I)` remains in its source disposition. Perception
stays `transcription-checked`; memory's approved formula is unchanged.

The migration table and a historical saved-graph fixture exercise the rename
through the native CLI. If a saved graph contains both old and canonical keyed
entries, renaming can collide. The tool now rejects duplicate migrated
`language:`/`statements:` keys before emitting or overwriting the file; it does
not silently choose between text, metadata or labels. This safeguard applies
to the same risk in earlier many-to-one predicate migrations.

## Other wording decisions

The imported original `third_party/carneades-4/examples/AGs/YAML/walton.yml`
at import commit `c7d0eaee5cd95ba86c0c20dfd64efb8aa3c3ff71` defines
`biased/1` with the generic template "%s is biased." The current library had
added "Arguer". Restoring the imported wording honors D5 and remains suitable
for the arguer in the book's entry 46; no predicate or logical rule changes.

The split between `morally_bad/1` (47, pp. 338–339) and `bad_character/1`
(42, p. 336) had not been documented separately. The explicit decision is to
retain it: moral badness is a more specific judgment than the unqualified
character description. The source does not establish automatic equivalence.
Entry 47's third premise requires the moral judgment rather than deriving it
from bias alone. An application may justify an explicit bridge, but the
generic library should not silently supply one.

## Approval and verification boundary

No stage is promoted by these changes. The changed `biased/1` rendering affects
semantic fingerprints of records that reference it; the perception rename
also changes its fingerprint. The registry must be refreshed only after an
independent integrator accepts this bounded change. The implementation commit
deliberately leaves the old fingerprints in place, so the gate detects the
pending acceptance instead of treating changed bytes as automatically approved.

The complete 51-record PR13 comparison is supplied separately in
`peer13-comparison-second.md` and its JSON companion. Its distinction between
baseline approval, repaired-source approval and remaining source/context
qualifications remains in force.

## Focused measured results

On Go 1.24.13, the five new source cases and two saved-graph migration
fixtures produce **12 checks, all passing**, including CLI round trips. The
source and migration-table audit produces **651 checks, all passing**. The
Python migration tests pass **10 of 10**, the category files regenerate cleanly,
and the validator reports **0 errors and 11 warnings**. No engine files were
changed and no additional engine unit-test run is claimed for this amendment.

Running the five new cases against `a708add`'s unchanged library with the same
engine produces **3 failing checks out of 10**: the split perception conclusion
persists, and the canonical belief conclusion loses support when memory is
defeated even though perception remains available (also after CLI round trip).
These failures demonstrate that the new expectations discriminate the defect.

The unchanged approval registry correctly reports only two stale approvals,
`bias` and `credible_source`, because their shared template changed. The other
changed fingerprints belong to already-held `perception` and
`witness_testimony`. `memory`'s own scheme and semantic fingerprint do not
change. Full-suite inventory regeneration and independent acceptance belong
to integration, not this focused implementation commit.
