# Dialogue adapter implementation and verification record

Date: 2026-11-21. Author: ChatGPT closure scope implementation author.
**No independent source approval is granted by this implementation record.**

The earlier all-50 adjudication identified a concrete practical-reasoning
burden gap and a separate ignorance metadata gap. The implementation is now
present in `tools/dialogue_snapshot.py`; its full interface, source basis, and
bounded contract are in `docs/dialogue-snapshots.md`. It leaves the canonical
scheme formulas unchanged.

The adapter adds supported per-argument question attacks, uses exact scheme
and ordered parameter binding, preserves original undercutter evidence, and
rebuilds each state from one immutable evaluated base. An answer requires
explicit external adequacy acceptance and accepted independent evidence.
Ignorance-family arguments require separate checked dialogue, probandum,
burden-holder, issue, and standard metadata. These are supplied allocations,
not machine-discovered facts.

Independent review during implementation identified two material risks.
The final ground-term normalizer preserves quoted spaces and escaped content;
it cannot silently match `"New York"` with `"NewYork"`. The earlier supported
grammar rejected quoted terms rather than admitting the proposed collision,
but that also excluded legitimate exported parameters, so quote support and
identity controls were still required. The final answer check also rejects
indirect dependency on the challenged conclusion or undercutter, including
paths through competing issues. A frozen `in` label is insufficient.

The dependence rule is intentionally conservative; an alternative independent
proof does not override a detected dependent path. Explicit assumptions other
than the challenged claim/undercutter are treated as independently stipulated
evidence, whose adequacy is asserted by the designated operator. Shared or
issue-competing target undercutters are refused because this limited adapter
cannot guarantee its exact-target suspension policy for those graphs.

## Focused evidence

The command below passed with the local Go1.24.13 toolchain,
`GOTOOLCHAIN=local`, and `GOFLAGS=-mod=mod`:

```sh
python -m unittest -v tools/test_dialogue_snapshot.py
```

Result: **12 Python test groups passed**, using actual Carneades evaluation.
They cover all **30 listed critical questions across six families**, four
snapshot states, another argument for the same conclusion, surviving original
defeaters, multiple open questions, quoted terms, indirect answer dependency,
independent derived answers, 22 malformed-journal/base variations plus other
reference controls, ignorance metadata, duplicate YAML keys, and actual Python
CLI output/error behavior. Practical transition states are also evaluated a
second time after export/import. These are not 12 new Go registered checks.

Evidence retained by the integration reviewer:
`scope-audit/dialogue-tests-final.log`, SHA-256
`79d0a522f5d2664bd0ebb23535495f7dec441e772a8e8e49afacc62d1d2ef50d`.
Earlier failed attempts and the reviewed intermediate test run remain in the
same evidence directory; they are not substituted for the final run.

The strict runner now invokes this test after engine-source setup, so the
combined toolchain gate includes it. The final integrated clean-copy run is
owned by the independent integration reviewer and is not claimed by this
focused author run.

This change also corrects this reviewer's earlier Atkinson wording: the
partition of seven direct, five partial, and five absent mappings leaves
**ten questions without full-source coverage**. The historical eight-item
worklist omitted the partial CQ8–9 treatments. Qualified AS1 remains outside
the source-approved set; full AS2 remains the separately specified extension.
