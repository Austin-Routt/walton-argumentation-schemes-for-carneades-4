#!/usr/bin/env python3
"""Tests of tools/migrate_r2.py (PLAN.md 5.6).

    python3 tools/test_migrate_r2.py

Identifiers are renamed where they stand as identifiers, and natural language
is left alone: the values of language: and statements:, and meta blocks.
"""
import os
import sys
import unittest

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from migrate_r2 import Migrator  # noqa: E402

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
                  "[receives_evaluation(A,good)]", "scheme: threat", "[possible_goal(x)]", "conclusion: action_fits_quality(x)",
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

    def test_chains_and_merges(self):
        # posible -> can_be_realized -> possible_goal; part_propert ->
        # part_property -> has_property, which group_2's has/2 also became
        src = "premises: [part_propert(rose,red), has(g,q), posible(x), has_property(y,z)]\n"
        self.assertEqual(Migrator().text(src),
                         "premises: [has_property(rose,red), has_property(g,q), possible_goal(x), has_property(y,z)]\n")

    def test_by_hand(self):
        # a statement whose form changed for one record is listed, not rewritten
        src = "assumptions:\n  - carry_out_action(go)\n  - plausibly_true(rain)\n"
        m = Migrator()
        self.assertEqual(m.text(src), src)
        rows = {(r["record"], r["new"]) for r in m.needs_hand(src)}
        self.assertIn(("practical_inference", "should_do(P,A)"), rows)
        self.assertIn(("appeal_to_vanity_ad_populum", "ought_to_bring_about(X)"), rows)
        self.assertIn(("position_to_know_ad_populum", "A"), rows)


if __name__ == "__main__":
    unittest.main(verbosity=1)
