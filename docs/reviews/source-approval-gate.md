# Source approval integrity gate

`tools/check_source_approvals.py` detects missing or stale recorded approvals. It does not perform a source review, authenticate a reviewer, or create an approval. The independent reviewer must first make a documented decision; the registry records that decision and the exact representation it approved.

## Commands

```sh
# Check recorded approvals and require them for every promoted row.
python3 tools/check_source_approvals.py --repo /path/to/repository

# The positional repository argument is equivalent.
python3 tools/check_source_approvals.py /path/to/repository

# Additionally require every library record to have reached example-passed.
python3 tools/check_source_approvals.py --repo /path/to/repository --require-stable

# Print candidate fingerprints as JSON. This grants no approval.
python3 tools/check_source_approvals.py --repo /path/to/repository --fingerprints

# Run the synthetic controls and mutations.
python3 tools/test_check_source_approvals.py
```

With no repository argument, the script uses the repository containing itself. `--fingerprints` is an inventory mode, mutually exclusive with `--require-stable`; it succeeds without an approval registry and does not write one. Exit status is 0 for the requested successful check/inventory, 1 for failed approval or stable-stage requirements, and 2 for malformed inputs or command usage.

The default mode requires a registry entry for each row at `interpretation-approved`, `questions-sorted` or `example-passed`. It also checks every existing approval, including entries whose ledger rows have an earlier stage. An empty or absent registry is valid when no row is promoted. Unknown or duplicate scheme, ledger or approval IDs are rejected, and the source ledger must cover exactly the nonempty master library.

Stable mode treats **every record in `walton.yml` as released**. There is no implicit exemption or optional exclusion list. A deliberately smaller stable distribution must contain its actual library and corresponding ledger in the tree being checked. This gate checks the source stages; the complete V&V gate must also pass.

## Registry format

The path is `docs/reviews/independent-approvals.json`. The following is a format example, not an actual approval:

```json
{
  "schema": "walton-source-approvals/v1",
  "approvals": [
    {
      "record": "RECORD_ID",
      "fingerprint": "sha256:REVIEWED_FINGERPRINT",
      "evidence": ["A report/source locator identifying the independent comparison and its scope"],
      "examples": ["RECORD_ID/book"]
    }
  ]
}
```

`evidence` and `examples` must be nonempty lists of nonblank strings. Evidence identifies the material and reasoning supporting the approval; the tool does not establish that the cited reasoning is true or that a person actually performed it. Example IDs must exist in `tools/vv/cases/*.yml`. Existence is checked; execution and success belong to the V&V run. Additional registry fields may record reviewer, date, scope or limitations, without pretending to provide cryptographic identity authentication.

## Fingerprint formula

The value is `sha256:` followed by SHA-256 over the UTF-8 encoding of this JSON object, serialized with sorted mapping keys, `ensure_ascii=False`, compact separators `(',', ':')` and `allow_nan=False`:

```text
{
  format: "walton-source-semantics/v1",
  scheme: the complete raw argument_schemes record,
  language: the statically referenced templates,
  issue_schemes: the complete master library issue_schemes object,
  source_contract: {
    origin, upstream, entry, cites, claim, source_questions,
    questions: all question fields except each question's note
  }
}
```

Missing selected scalar ledger fields become JSON `null`; missing questions become an empty list; missing issue schemes become an empty object. The complete raw scheme includes metadata, roles, weights and all other fields. Mapping order and YAML comments do not affect the hash. Lists and statement strings retain their order and content. This is a conservative representation fingerprint, not a proof that two differently written schemes are logically equivalent.

Ledger stage and explanatory notes do not affect the fingerprint. An alteration to a question's number, kind, mapping or any other non-note field does. Changing a raw scheme metadata note does affect the fingerprint, because the whole raw scheme is included.

The existing validator's `parse_statement` selects the outer predicate's template from premises, assumptions, exceptions, conclusions, deletions and guards. It also identifies the underlying predicate of a negated statement. This matches the current renderer: nested term arguments are rendered as terms, rather than by applying their own language templates. No unrelated language template is included. A missing referenced template or an unsupported statement syntax fails closed.

**Statement-variable limitation:** a scheme can conclude or consume a whole statement through a variable. The concrete predicate and its template are supplied by a graph at runtime, so they cannot be determined from the library record. Those runtime templates are outside this per-record fingerprint. Global issue schemes are included because changing opposition can change an approved argument's treatment. Engine code, case contents/results and arbitrary graph-specific extensions remain outside this fingerprint and require their separate verification gates.

## Mutation evidence and integration

The test module uses only temporary synthetic records and synthetic approvals. Fourteen tests cover an unchanged approved record; changed conclusion and metadata; missing approvals at every promoted stage; changed CQ scope, kind and source contract; relevant language and global opposition changes; irrelevant-template and ledger-note stability; mapping-key order; unknown and duplicate IDs; missing evidence; nonexistent examples; stale approvals on unpromoted rows; absent/empty registries; both repository argument forms; inventory without approval creation; and stable-mode refusal of an incomplete library.

This change adds the checker, its tests and this documentation. It does not edit source stages, create the approval registry, or alter `PLAN.md` or `tools/vv/run.sh`. The integrating reviewer generates the real registry from final reviewed dispositions and wires the default check and mutation tests into CI. Stable-release execution uses `--require-stable` in addition to the full runtime/library test gate.
