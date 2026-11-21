# Registered inventory review

The public baseline contains8,230 named checks. The closure candidate contains9,959:1,747 additions and18 replaced source-language/migration identities. Every emitted result is required exactly once. The development inventory-update run passed9,959/9,959; it is not clean release acceptance. Final acceptance runs without update/skip flags on both toolchains.

All18 removed identities are metadata checks for nine obsolete language names and nine superseded migration destinations. No engine behavioral check or source-case check was removed. The complete added/removed lists, replacement identities and grouping are in `development/inventory-diff.json` in the review archive. Scoped participant arities, explicit sufficiency and likelihood-adequacy changes require manual migration where new information is absent; the replacement tests retain that requirement.

New checks cover all23 records in master/category API/CLI paths, source placement and examples, the new discriminating cases and migration rules. Native engine tests, harness oracle mutations and thirteen actual-engine dialogue test groups are separate gates; they are not inflated into the registered9,959 total. `known_issues.yml` remains empty and its minimum rises to9,959. The inventory and minimum are tightened, not weakened.

Six inherited unused templates and five dependency-cycle groups account for all eleven validator warnings. Warnings are disclosed structural observations, not known failing executable checks.
