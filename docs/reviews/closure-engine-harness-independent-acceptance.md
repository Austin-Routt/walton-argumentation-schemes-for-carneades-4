# Independent engine and V&V harness acceptance

Reviewer: ChatGPT fresh non-author acceptance reviewer, 2026-11-21.

**Decision:** accept the reviewed code corrections for the declared contracts
below. This is code-review acceptance, supported by focused execution. Final
acceptance of the combined candidate still requires clean full gates on the
frozen commit under both pinned Go versions.

## Reviewed changes

I did not author these implementation changes. I reviewed their integrated
content in a separate clone, initially at
`d1da69e6d24d115763fc3e826ed6f7a708642ac6` and subsequently at
`22bfa8dc1a8921bb4e063745537f192ad6eab783` for the context-loss repair.

| Change | Original author-branch commit | Integrated commit reviewed |
|---|---|---|
| Preserve valid API statement keys in ground issue lookup | `ff110b70ba3bba64829b61c6c7bddaa05f4d9445` | `b7dc50bf1edd92c42722ac29ad1886cc058144c5` |
| Normalize explicit negation across graph identities and authored instances | `f66accc5de5ed5ae3a90d35186f39a56f706cace` | `15e3b1a1e3bc0eee5000e5c97aa9303bb8931a61` |
| Make bound-negation authored validation independent of premise order | `3a64f533cdce4a6d97daaa912d0dc8aa24b65662` | `60db373e2ad7839a93a79f382f395894fabacba2` |
| Check every declared case expectation across CLI round trips; select live library schemes | `2f5a0e686eabfb19384e649a9d49fe262909875b` | `d1da69e6d24d115763fc3e826ed6f7a708642ac6` |
| Preserve existing case context in every CLI round trip | Authored by the integration reviewer after my counterexample | `156e446df49730b546729339c28e567e3e01ce38` |

### Ground issue lookup

The index is built from parsed statement keys and retains the original
statement objects. A valid API key with noncanonical spacing therefore remains
discoverable. The additional typed match is necessary: a printed term key
alone is not sufficient to establish typed term equality. Open patterns and
enumerations retain the scan. The supplied differential controls compare
indexed and scanned behavior for aliases, lists, numeric types, duplicate
patterns, open patterns, enumeration and substituted explicit negation.

This accepts the correction to the optimized lookup. It does not assert that
every arbitrary issue topology has been formally verified. The documented
single-issue constraint and rejection of connections between distinct existing
issues still apply.

### Explicit negation and authored arguments

Paired explicit negations are cancelled structurally, including after
parameter substitution. The implementation does not evaluate arithmetic or
reinterpret quoted text. Graph normalization happens on the private inference
copy. Distinct objects that would collapse to one identity are rejected;
successful publication retains existing caller-visible objects. Assumptions,
defaults, questioning and expected labels use the same normalized identity.

The repaired example with `X = ¬p` correctly makes an exception `¬X` refer to
`p`. The generated argument, its exception, the authored form and its imported
form must all agree on that identity. Authored validation now obtains a shared
substitution from premises before judging a conclusion that depends on those
bindings. Where a negated variable still needs a binding, it tries a different
premise order. It preserves one-to-one premise matching and requires exception
witnesses to agree with the same substitution. This is not a general inverse
unification procedure for every logically equivalent formula.

I found no additional engine defect in this review of those changes. Focused
native tests passed with both Go 1.24.13 and Go 1.27.4. The logs are
`fresh-audit/engine-focused-go1.24.13.log` and
`fresh-audit/engine-focused-go1.27.4.log` in the review evidence. Full engine
build, vet, test and race gates remain part of final frozen-commit acceptance.

A subsequent boundary review found a separate public-API origin defect:
explicitly stipulating an existing questioned default did not promote it out
of default status. The reproduced finding, repair and independent acceptance
are recorded in `closure-assumption-origin-independent-acceptance.md`. That
addendum supersedes any inference that the focused review above established
the absence of further defects.

### Case expectations and actual library selection

The new observation layer checks declared `in`, `out`, `not_in`, required
arguments, forbidden argument instances, forbidden schemes and exact issue
position sets. Both first CLI output and the re-imported CLI output must meet
those expectations. The observer does not call inference again to repair the
result it is judging. A stable but incorrect result therefore fails.

Optional scheme selection retains the actual imported library records and
rebuilds a fresh scheme index. An omitted list retains all schemes; an explicit
empty list retains none. Unknown and duplicate selections are rejected. The
mutation control changes the selected library record and verifies that the
test observes the change. This avoids validating a copied fixture while the
actual library silently diverges.

## Additional independent finding: existing case context was lost

**FRESH-HARNESS-CONTEXT** is an additional harness defect found during this
fresh review. It is separate from the inference-engine defects discussed in
the earlier reviews.

At `d1da69e`, an inline case graph could contain `assumptions: [p]` and a rule
with premises `p, q` and conclusion `r`, while the case supplied the additional
assumption `q`. The API path appended the new assumption and derived `r`.
The CLI path, when no explicit scheme selection was present, appended a
second `assumptions` YAML key. The importer then retained only the latter
context, and `r` was absent. Existing `questioned` context had the same risk.

My independent probe reproduced the discrepancy. The strengthened result
oracle correctly reported the missing conclusion; it did not create the
underlying context-loss defect.

The integration reviewer repaired this in `156e446`: all case paths now
structurally combine the imported graph's context with the case's added facts
and questions, then export that **unevaluated** graph for the separate CLI
process. This preparatory import/export performs no inference. Dedicated raw
CLI parsing and refusal tests remain separate from this case-construction
helper.

I reviewed the repair and independently reran my unchanged original probe,
the new regression covering both existing assumptions and existing questions,
and the API/CLI mutations for every declared expectation class. They all pass.
The clone was clean again after the temporary probe was removed. Evidence:

- `fresh-audit/existing_case_context_probe.go` — unchanged independent probe.
- `fresh-audit/existing-case-context.log` — reproduced failure before repair.
- `fresh-audit/existing-case-context-fixed.log` — passing probe, context
  regression and adverse expectation controls after repair.

## Acceptance boundary

The reviewed corrections are accepted for the offline CLI and Go API profile
defined in `docs/contracts.md`. API callers still must perform the documented
structural validation and handle typed context/inference errors before using
labels. Finite solver quotas remain safety limits, with the separately stated
limits on input size, cancellation and resource guarantees.

This review grants no unavailable-source approval and promotes no inherited
source hold. The four new compendium-form decisions are recorded separately in
`closure-supplemental-independent-acceptance.md`. It also does not certify a
web service, unlimited inputs, all possible dialogue protocols or all future
theory combinations.

The final independent run must record the exact frozen commit, pinned Go
versions, clean status before and after, complete engine/race gate results,
strict V&V results, unique check identities and agreement between toolchains.
Until that evidence exists, the combined candidate's final gate is pending.
