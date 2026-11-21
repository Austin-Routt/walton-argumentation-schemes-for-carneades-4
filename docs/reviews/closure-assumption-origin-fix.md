# Explicit stipulations must replace default origin

Date: 2026-11-21. Finding: **FRESH-DEFAULT-ORIGIN**. The fresh independent
reviewer supplied the counterexample; the integration reviewer authored this
repair. Independent acceptance of the repair and the final clean gates are
recorded separately.

Import a graph with statement `p`, `defaults: [p]`, and `questioned: [p]`.
After inference, p correctly remains out. Calling the public
`AddAssumption("p")` must make p an explicit stipulation. Previously it appended
the fact but retained `Defaults[p]`, so another inference still treated it as
an unanswered default and left it out. The same failure occurred when a case
fixture appended new stipulated facts to a graph that already contained a
default for the same formula.

The public API now normalizes the added identity and removes every matching
default marker, including noncanonical aliases present before a first
inference. It preserves unrelated defaults and the history that the statement
was questioned. A private append operation lets a scheme introduce a default
without accidentally stipulating it through the public API. Case preparation
and premise-withdrawal variants now add supplied facts through the public
stipulating operation. Repeated scheme instantiation cannot demote that
stipulation back to a default.

The focused native test covers 18 combinations of pre-evaluation state,
default-marker spelling and added-fact spelling, plus a generated-default
control and repeated inference after promotion. It checks continued suspension
of an unrelated default and existing statement-pointer preservation. The
existing question/default tests still pass. The harness tests cover canonical
and doubly negated supplied facts through the API, both CLI round trips and
the isolated premise-withdrawal path. These are native/unit controls, reported
separately from the registered V&V check inventory.

Author runs passed on Go1.24.13 with `GOTOOLCHAIN=local` and
`GOFLAGS=-mod=mod`; full logs are `assumption-audit/native.log` and
`assumption-audit/harness.log` in the review archive. The fresh reviewer's
unchanged failing probe is preserved in `fresh-audit/promote_default_probe.go`
with its before/after execution evidence. This defect is separate from the
earlier harness-only loss of existing assumptions through duplicate YAML keys.
