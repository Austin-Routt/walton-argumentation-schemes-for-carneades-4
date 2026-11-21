# Finite dialogue snapshots

`tools/dialogue_snapshot.py` implements an explicit question, accepted-answer,
and reopening protocol around an already constructed Carneades graph. It
produces a fresh graph for evaluation. The canonical scheme records keep their
ordinary premises, assumptions, and supported exceptions.

The source basis is the supplied *Argumentation Schemes* (2008), printed p.96:
an appropriate unanswered critical question can temporarily defeat a practical
inference, and an adequate answer can remove that challenge. Position-to-know
reasoning has an analogous burden discussion in Walton, *Legal Argumentation
and Evidence* (2002), pp.46–48; access and the limits of that source comparison
are recorded in [the closure adjudication](reviews/closure-source-adjudication.md).
The adapter is an executable, bounded application protocol for those passages.
It does not infer that every question in every scheme has this burden policy.

## Supported question inventories

| Scheme | Critical questions |
|---|---:|
| `practical_inference` | 1–5 |
| `practical_necessary_condition` | 1–5 |
| `practical_sufficient_condition` | 1–5 |
| `value_based_practical_reasoning` | 1–7 |
| `two_person_practical` | 1–5 |
| `position_to_know` | 1–3 |

The numbers refer to the source-ledger inventory of each named record. A
dialogue operator identifies an appropriate question and explicitly adjudicates
whether supplied evidence answers it. The program validates those declarations;
it does not interpret natural-language questions, establish the truth of an
observation, or decide the adequacy of an answer from an arbitrary atom's name.

## Inputs and exact targeting

First construct and evaluate the case using the desired live library records:

```sh
python3 tools/graph_from_library.py practical_inference > case.yml
# Supply the case's actual statements and assumptions in case.yml.
carneades eval -t yaml case.yml > base.yml
```

Keep `base.yml` unchanged for this dialogue. A target names a constructed
argument ID **and** its actual scheme **and** its complete ordered parameter
list. Copy these from the exported `arguments` mapping. Argument IDs are not
stable across independently reconstructed graphs, so a journal is bound to
this specific base, whose canonical content hash is included in each result.
The target must have premises, a known conclusion, and one unambiguous schema
with the corresponding parameter arity. An unexpanded inference request is
insufficient.

For `practical_inference`, the parameter order is action, goal, agent. For
example, if the actual exported argument ID is `a42`, a journal can be:

```yaml
schema: walton-dialogue-snapshot/v1
dialogue: treatment_review_1
targets:
  exercise_argument:
    argument: a42
    scheme: practical_inference
    parameters: [exercise, health, ann]
events:
  - {id: question_1, type: ask, target: exercise_argument, cq: 1}
  - id: answer_1
    type: answer
    question: question_1
    evidence: [answer_evidence_1]
  - id: reopen_1
    type: reopen
    question: question_1
    reason: A new objection challenges the accepted adequacy assessment.
evidence:
  answer_evidence_1:
    dialogue: treatment_review_1
    target: exercise_argument
    question: question_1
    answer: answer_1
    statement: accepted_goal_review(ann,health)
    accepted: true
    accepted_by: Designated case reviewer
```

The example IDs and evidence are illustrative inputs. The argument must
actually exist, and `accepted_goal_review(ann,health)` must already be an
assumption or an `in` statement in the supplied evaluated base. The operator's
acceptance declaration must be justified by the application; naming a fact
`accepted_goal_review` does not prove its adequacy.

Ground term normalization removes whitespace outside quoted literals and
preserves all content and escapes inside them. `"New York"` and `"NewYork"`
therefore remain different. Single-quoted atoms are also preserved. Equivalent
escape spellings are not recanonicalized; copy the engine's exported spelling.
Free variables and invalid or unterminated terms are rejected.

## Producing the states

```sh
python3 tools/dialogue_snapshot.py base.yml journal.yml --through initial > initial.yml
python3 tools/dialogue_snapshot.py base.yml journal.yml --through question_1 > questioned.yml
python3 tools/dialogue_snapshot.py base.yml journal.yml --through answer_1 > answered.yml
python3 tools/dialogue_snapshot.py base.yml journal.yml --through reopen_1 > reopened.yml
carneades eval -t yaml questioned.yml > questioned-evaluated.yml
```

Omit `--through` to apply the whole journal. Every journal entry is validated,
including entries after a requested prefix. An invalid later answer therefore
also prevents rendering an earlier prefix of that invalid journal. The command
returns exit status 2 and no partial graph on invalid input.

Each open question adds an assumed marker and an ordinary supported `linked`
argument for the target's existing undercutter. If the target has no
undercutter, a fresh one is created. All original arguments and support for an
existing undercutter remain present. The target is identified independently of
its conclusion: another argument for that conclusion survives the challenge.
Answering one question leaves any other unanswered question active.

An accepted answer removes only its question's added support in the fresh
snapshot. It does not assert the target's conclusion and does not cancel an
original supported exception. Reopening makes that question active again.
The adapter rejects a shared undercutter, because supporting it would affect
another argument. It also rejects a target undercutter that belongs to a
competing issue; this bounded adapter does not implement a priority policy
that could guarantee suspension against such competing positions.

Every state is rebuilt from the unmodified base. Passing a prior adapter
snapshot back as the base is rejected, preventing stale question attacks from
accumulating. Old statement labels, calculated values, and expected test
outputs are removed. The output's metadata says `requires_fresh_evaluation:
true`; evaluate it with Carneades before interpreting its conclusion labels.

## Answer integrity and the conservative dependence check

An answer has one or more nonduplicate evidence references. Every referenced
record must match this exact dialogue, target, question, and answer event;
contain an existing accepted base statement; and name an accepting operator.
Foreign, reused, missing, unaccepted, or unused evidence declarations are
rejected. Answering a closed question, reopening an open question, duplicate
event IDs, duplicate target declarations, duplicate YAML keys, unknown fields,
out-of-range question numbers, and references to future questions are also
rejected. A previously answered question is reopened explicitly instead of
creating a second declaration of the same target and CQ number.

An `in` label alone does not establish independent answer evidence. The target
conclusion and its undercutter cannot answer their own challenge. For derived
evidence, the adapter traces argument dependencies and issue competition from
those two statements. It rejects an answer if any such path reaches it. This
also rejects evidence derived indirectly from the challenged conclusion.
Other explicit base assumptions are independent stipulations under this
protocol; their factual adequacy is the operator's responsibility.

This check is deliberately conservative. It may refuse a statement that has
both a dependent proof and a separate valid independent proof. It does not
search for a minimal independent proof or recompute answer acceptability in
every possible dialogue state. Applications can supply a separately identified
independent observation with justified acceptance, or use a richer proof-aware
dialogue controller. They must not relabel a dependent conclusion merely to
evade the requirement.

## Ignorance: allocating a burden separately from proving a claim

Every constructed `ignorance`, `negative_normal_expectations`,
`negative_practical`, or `negative_practical_true` argument in the base needs
one explicit `burdens` record, even if it is not a question target. For example:

```yaml
burdens:
  - argument: a7
    scheme: ignorance
    parameters: [alarm]
    dialogue: treatment_review_1
    probandum: ¬alarm
    burden_holder: proponent
    issue: alarm_issue
    standard: PE
```

The argument and parameter binding must match the actual constructed graph.
The probandum must exist and belong to the named issue; the proof standard must
be `PE`, `CCE`, or `BRD` and match that issue. A nonempty burden holder and the
current dialogue are required. Missing or duplicate allocations and foreign
references are errors. These records are copied into the snapshot metadata.

The ultimate probandum may belong to a broader dialogue issue rather than be
identical to this argument's conclusion. Its relevance to the target argument,
and whether the designated party actually fulfills that broader burden, are
explicit application judgments. The checker verifies the allocation's target
identity, declared issue membership, and proof standard. It does not establish
logical pertinence or successful fulfillment merely because those references
are valid.

This closes a representation/validation gap: a proof standard is an evidence
threshold, not an identity for the burden holder or the ultimate claim at
issue. The adapter checks that an application supplies consistent metadata; it
does not decide who should have the burden, select a legal standard, infer
search completeness, or turn ignorance into proof. Existing assumption
challenges and the source-branch contract still govern the actual arguments.

For the two practical presumption choices, applications that choose the
bounded binary protocol declare its two positive choice statements explicitly
in one issue, as described in [the source branch contract](source-branch-contracts.md).
No automatic equivalence to the negation of the opposite choice is introduced.

## Scope and verification

This is a finite snapshot adapter with externally accepted questions and
answers. It does not implement a complete autonomous dialogue system, protocol
legality beyond the listed events, probability estimation, answer generation,
or the Atkinson–Bench-Capon AS2 transition/value model. Canonical static graphs
can still be used directly with their declared static interpretation.

`tools/test_dialogue_snapshot.py` tests the actual local engine, not a mock.
It runs every listed CQ in all six supported families, all four transition
states, repeated export/import/evaluation, independent support for the same
conclusion, surviving original undercutters, multiple simultaneous questions,
quoted identities, indirect self-support, independent derived evidence,
malformed journals, the command-line interface, and ignorance allocations.
`tools/vv/run.sh` includes it in the ordinary and strict CI gate. These Python
test groups are reported separately from the Go harness's registered check
count. Source approval and combined final acceptance remain independent gates.
