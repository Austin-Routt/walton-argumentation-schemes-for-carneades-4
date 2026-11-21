#!/usr/bin/env python3
"""Tests of tools/migrate_r2.py (PLAN.md 5.6).

    python3 tools/test_migrate_r2.py

Identifiers are renamed where they stand as identifiers, and natural language
is left alone: the values of language: and statements:, and meta blocks.
"""
import os
import contextlib
import io
import sys
import tempfile
import unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from migrate_r2 import Migrator, MigrationConflict, main  # noqa: E402

OLD = """meta:
  title: exeptional cases and the treat of posible waste
  note: >
    all_parts_propert(a,b) is mentioned here and stays
language:
  posible/1: "treat: %s is posible, said exeptional."
  all_parts_propert/2: "All the parts of %s have the property %s."
statements:
  "all_parts_propert(car,heavy)": "all_parts_propert in text stays"
  posible(x): posible stays here too
argument_schemes:
  - id: exeptional
    meta:
      title: exeptional
      source: "treat, posible"
    variables: [A]
    premises:
      - all_parts_propert(A,b)
    assumptions:
      - posible(A)
    conclusions:
      - treatment_required(A)
  - id: opposition-normative
    variables: [A]
    premises: [recieves_evaluation(A,good)]
    conclusions: [opposition-normative_2(A)]
arguments:
  a1:
    scheme: treat
    premises: [posible(x)]
    conclusion: not_fit_action(x)
assumptions:
  - all_parts_propert(car,heavy)
tests:
  in: [posible(x)]
"""


class Migrate(unittest.TestCase):
    def setUp(self):
        self.new = Migrator().text(OLD)

    def test_identifiers(self):
        for s in ["- id: exceptional", "- all_parts_property(A,b)", "- possible_goal(A)", "- id: opposition_normative",
                  "[receives_evaluation(A,good)]", "scheme: threat", "[possible_goal(x)]", "conclusion: not_fit_action(x)",
                  "  - all_parts_property(car,heavy)", "in: [possible_goal(x)]"]:
            self.assertIn(s, self.new, s)

    def test_keys_but_not_text(self):
        self.assertIn('  possible_goal/1: "treat: %s is posible, said exeptional."', self.new)
        self.assertIn('  all_parts_property/2: "All the parts of %s have the property %s."', self.new)
        self.assertIn('  "all_parts_property(car,heavy)": "all_parts_propert in text stays"', self.new)
        self.assertIn('  possible_goal(x): posible stays here too', self.new)

    def test_meta_untouched(self):
        self.assertIn("  title: exeptional cases and the treat of posible waste", self.new)
        self.assertIn("    all_parts_propert(a,b) is mentioned here and stays", self.new)
        self.assertIn("      title: exeptional\n", self.new)
        self.assertIn('      source: "treat, posible"', self.new)

    def test_whole_words_only(self):
        self.assertIn("- treatment_required(A)", self.new)
        self.assertIn("[opposition-normative_2(A)]", self.new)

    def test_idempotent(self):
        self.assertEqual(Migrator().text(self.new), self.new)

    def test_ethotic_strength_change_is_instance_specific_and_manual(self):
        src = 'arguments:\n  a1:\n    scheme: ethotic2\n    conclusion: "¬claim"\n'
        mig = Migrator()
        self.assertEqual(mig.text(src), src)
        self.assertTrue(any(r['record'] == 'ethotic2' for r in mig.needs_hand(src)))
        unrelated = 'assumptions: ["¬claim"]\n'
        self.assertEqual(mig.text(unrelated), unrelated)
        self.assertFalse(any(r['record'] == 'ethotic2' for r in mig.needs_hand(unrelated)))

    def test_chains_and_merges(self):
        # posible -> can_be_realized -> possible_goal; part_propert ->
        # part_property -> has_property, which group_2's has/2 also became
        src = "premises: [part_propert(rose,red), has(g,q), posible(x), has_property(y,z)]\n"
        self.assertEqual(Migrator().text(src),
                         "premises: [has_property(rose,red), has_property(g,q), possible_goal(x), has_property(y,z)]\n")
        shared = "premises: [same_properties(tweety,bird,flies), person_judgment_value(ann,honest)]\n"
        self.assertEqual(Migrator().text(shared),
                         "premises: [has_properties(tweety,bird,flies), person_value(ann,honest)]\n")

    def test_perception_memory_conclusion_identity(self):
        # Old saved perception keys/terms and the old memory misspelling
        # converge, while human-authored text remains the user's text.
        src = ('language:\n  reasonable_case/1: "reasonable_case in prose stays."\n'
               'statements:\n  reasonable_case(door_closed(case1)): "reasonable_case in prose stays."\n'
               'assumptions: [reasonable_case(door_closed(case1)), reasonable_belive(door_closed(case2))]\n')
        want = ('language:\n  reasonable_believe/1: "reasonable_case in prose stays."\n'
                'statements:\n  reasonable_believe(door_closed(case1)): "reasonable_case in prose stays."\n'
                'assumptions: [reasonable_believe(door_closed(case1)), reasonable_believe(door_closed(case2))]\n')
        migrated = Migrator().text(src)
        self.assertEqual(migrated, want)
        self.assertEqual(Migrator().text(migrated), migrated)

    def test_many_to_one_key_collisions_do_not_discard_text(self):
        for section, old, new in [
            ("language", "reasonable_case/1", "reasonable_believe/1"),
            ("statements", "reasonable_case(p)", "reasonable_believe(p)"),
        ]:
            with self.subTest(section=section):
                src = f'{section}:\n  {old}: "first text"\n  {new}: "different text"\n'
                with self.assertRaisesRegex(MigrationConflict, "duplicate migrated key"):
                    Migrator().text(src)
                with tempfile.TemporaryDirectory() as tmp:
                    path = os.path.join(tmp, "graph.yml")
                    with open(path, "w", encoding="utf-8") as file:
                        file.write(src)
                    stdout, stderr = io.StringIO(), io.StringIO()
                    with contextlib.redirect_stdout(stdout), contextlib.redirect_stderr(stderr):
                        self.assertEqual(main(["migrate_r2.py", path]), 2)
                    with open(path, encoding="utf-8") as file:
                        self.assertEqual(file.read(), src)
                    self.assertEqual(stdout.getvalue(), "")
                    self.assertIn("reconcile the original entries by hand", stderr.getvalue())

    def test_likelihood_threshold_change_requires_review(self):
        # A denial of strong likelihood need not deny adequacy for a modest
        # defeasible claim. Do not silently equate the old and new predicates.
        src = "questioned: [outcome_likely(tax,recession)]\nassumptions: [¬outcome_likely(tax,recession)]\n"
        mig = Migrator()
        self.assertEqual(mig.text(src), src)
        rows = {(r["record"], r["new"]) for r in mig.needs_hand(src)}
        self.assertIn(("negative_consequences_reasoning", "consequence_likelihood_adequate(A,B)"), rows)

    def test_by_hand(self):
        # a statement whose form changed for one record is listed, not rewritten
        src = "assumptions:\n  - carry_out_action(go)\n  - plausibly_true(rain)\n"
        m = Migrator()
        self.assertEqual(m.text(src), src)
        rows = {(r["record"], r["new"]) for r in m.needs_hand(src)}
        self.assertIn(("practical_inference", "should_do(P,A)"), rows)
        self.assertIn(("appeal_to_vanity_ad_populum", "ought_to_bring_about(X)"), rows)
        self.assertIn(("position_to_know_ad_populum", "A"), rows)

    def test_new_premise_and_role_changes_require_evidence(self):
        src = "assumptions: [is_a(whale,mammal), successful_diagram(c1), based_on_false_belief(p), action_character(ann)]\n"
        mig = Migrator()
        # Roles and coextension cannot be invented by a spelling migration.
        self.assertEqual(mig.text(src), src)
        rows = {(r["record"], r["new"]) for r in mig.needs_hand(src)}
        self.assertIn(("rule_physical_world_1", "referentially_coextensive(W,D)"), rows)
        self.assertIn(("abductive_forward", "diagram_conclusion(G1,C1)"), rows)
        self.assertIn(("memory", "based_on_false_belief(P,M)"), rows)
        self.assertIn(("abductive_action_to_character", "action_character(A,X,Q)"), rows)

    def test_general_partition_witnesses_are_not_manufactured(self):
        src = "assumptions: ['whole_of(team,ann,bob,cy)', 'part_whole(team,strong,ann,bob,cy)']\n"
        mig = Migrator()
        self.assertEqual(mig.text(src), src)
        rows = {(r["record"], r["new"]) for r in mig.needs_hand(src)}
        self.assertIn(("whole_into_parts_selected", "whole_partition(X,S)"), rows)
        self.assertIn(("whole_into_parts_selected", "some_part_sufficient_for_whole(X,S,A)"), rows)
        self.assertIn(("whole_into_parts", "selected_part_sufficient_for_whole(X,X1,A)"), rows)


if __name__ == "__main__":
    unittest.main(verbosity=1)
