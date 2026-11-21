# Critical response to the implementer's PR15 report

Reviewed public baseline: `e6d2d6f427c6ec0447cae514810c256903b15779`,
2026-11-21. This addendum preserves the original
`implementer-response.md` as historical evidence and supersedes the
conclusions below. It is not a claim that every statement in every earlier
report was correct.

## Findings that hold

The public GitHub evidence confirms PR15's merge and successful checks for
Go1.24.13, Go1.27.4 and their check-by-check comparison. A fresh full clone
of the exact public commit independently passes the engine gate and the
strict8,230-check suite on Go1.24.13. Final acceptance of the new candidate
uses separate clean copies and both pinned versions.

The criticism of the previous `fccec4f` candidate's practical capacity is
valid. Its low candidate-visit budgets rejected ordinary bounded inputs.
The implementer's ground-position optimization and higher budgets address
that problem. The binding-chain termination repair is justified, and
normalizing the authored exception comparison is appropriate. This review
retains those changes, subject to the repairs and limitations below.

The entry16 wording is source-consistent: the question concerns whether
vagueness creates a precision problem for the dialogue. The subsequent
independent static-burden review and positive/adverse cases justify its
acceptance. Entry33.3 should use the same likelihood-adequacy assessment as
entries33.1–33.2. That change is implemented and independently accepted;
saved positive or negative judgments about a *strong* probability threshold
are not automatically judgments about adequacy for a particular argument.

The old release pin `8c092b6` remains superseded. The Atkinson paper's17
critical questions and the five absent separate mappings are correctly
identified. Seven questions have direct treatments, five partial policies,
and five no separate mapping in the inherited form. That means **ten**, not
eight, lack full-source treatment. The earlier eight-item list omitted CQ8
and CQ9 despite explicitly admitting their aggregate side-effect policy lacks
the source's negative arguments and value comparison. Both reviewers now
agree with this correction. The historical narrower eight-item worklist is
preserved in the machine-readable decision table; it must not be described
as a full-coverage count. CQ10 remains positive support.

## Corrections to the response

### The defeated-default counterexample was real

The exact earlier probe gives two arguments and effective assumptions
`p`, `q`, `extra` on `96daf6c`, with downstream `r` and `s` accepted. On
`fccec4f` and `e6d2d6f`, it gives zero arguments, only `q` effectively
assumed, and `p` out. The supplied contrary is stipulated. Replacing that
stipulation with an ordinary argument supporting a contrary tests a different
priority rule. The implementer's claim of identical behavior on the exact
counterexample is false. This is supported by the executable probe and the
three recorded outputs in the review evidence, not an inference from the
contract text alone.

The response's table also contains nine rows, eight labelled reproduced/fixed
and one disputed. Its sentence “seven of the eight” does not describe that
table accurately. Counts of grouped counterexample families and counts of
individual defects must be named separately.

### The lookup optimization had an API regression

Ground lookup used canonical string keys against a map that the accepted API
could populate with valid noncanonical spacing. For example, matching
`accepts(a)` against key `accepts( a )` succeeded under the scan but failed
under direct lookup. The repair indexes parsed normalized formulas while
retaining graph nodes and checking typed matching. Separate identity and
numeric-type controls guard against merely replacing one key assumption
with another.

### Explicit double negation was not consistent across layers

With `given(X) -> q except ¬X`, facts `given(¬p)` and `p` must instantiate
the exception as `p`. The old graph builder instead retained an unsupported
`¬¬p` node and accepted q. The shared structural normalizer now cancels
pairs of explicit negations in constructed formulas, context identities,
authored matching and interchange. It preserves quoted values and types,
rejects ambiguous collisions, and preserves existing API node pointers on
successful publication. A further independent review found a premise-order
case in parameterless authored validation; the repair lets a later premise
bind a negated variable before comparison, with no invented inverse binding.

### Quotas remain finite limits

A quota of50,000,000 is a larger finite ceiling. It can still reject a valid
finite input that an unlimited evaluator might eventually finish. The
supported guarantee is explicit, atomic, diagnosable refusal when the
budget is reached. The implementation does not promise unbounded graph
capacity, a universal wall-clock deadline or a process-memory byte limit.
The “safety valves, not working limits” language is corrected in current
contracts and release guidance.

### Performance measurements need their actual fixtures and denominators

The lookup comparison test's600 insertions yield283 distinct statement keys.
The named1,200-input closure case contains1,800 graph nodes after generated
undercutters. Another3,000-positive-node fixture includes750 contraries,
giving3,750 actual statements. These are distinct measures. The original
mixed1,200-fact/2,199-node/14-issue benchmark file is not present in the
repository evidence examined, so its exact0.71-second figure is not
independently certified. New timings are reported only for their saved,
reproducible fixtures, machine and toolchain; they are not service guarantees.

### Source holds included executable defects

Some holds were conservative scope qualifications. Others exposed genuine
cross-agent or cross-case inference, omitted branches, unsupported reversal
of a conditional, or a stronger conclusion than the source permits. The
new context, part–whole, practical, credibility and explicit-variant changes
are substantive repairs. Approval totals cannot resolve these differences
by majority vote. Each accepted interpretation needs its own source and
discriminating example evidence.

### Source access is more specific than “the books are unavailable”

Relevant page-marked full text of Walton2006 and Walton2002 is now accessible
and was inspected, along with the author-hosted2003 witness article. This
changes several holds from missing access to a known formulation difference
or a dialogue-policy obligation. These are not facsimile-based checks of
every page of those books. The exact1995,1997, other cited editions, the
relevant2008 witness book passage,2015 passage and2019 correspondence remain
separate obligations where indicated in the current record register.

## Independence and decision rules

The implementer's additional review of eight changed source records is
useful corroboration and is retained. The earlier package was produced by
multiple cooperating review agents; the phrase “the first review alone”
does not establish which individual authored every change. The relevant
condition is clear per-change non-author review, not the product label on
the archive. New acceptance records name author/reviewer separation and
the exact cases and semantic fingerprints. These are AI-assisted technical
reviews, not claims of independent human or authorial endorsement.

The owner decisions are resolved operationally: add explicit part–whole and
other supported source alternatives; keep legacy AS1 outside the complete
source-approved set; scope full AS2 as a separate finite-model extension;
use the corrected accepted preview candidate; reserve1.0.0 for the unchanged
all-record stable gate. Publication requires actual authenticated write
access. Upstream correspondence remains prepared and unsent without explicit
authorization to contact others.

## Evidence and remaining work

The single review archive contains the baseline clean-run logs, public
GitHub API evidence, three-way default probe, lookup/negation controls,
performance scripts, source changes, independent approvals, final clean-run
evidence and exact candidate repository. The current closure register and
PLAN supersede old70/120 and50-hold summaries. Historical files retain their
original dates and numbers so the review history stays auditable.
