#!/usr/bin/env bash
# One command for a developer or a CI job: the two Python checks of the
# repository, then the engine-level suite.
set -uo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo="${WALTON_REPO:-$(cd "$here/../.." && pwd)}"
status=0
echo "== tools/validate_schemes.py"
python3 "$repo/tools/validate_schemes.py" "$repo" | tail -1 || status=1
echo "== tools/build_category_files.py --check"
python3 "$repo/tools/build_category_files.py" "$repo" --check || status=1
echo "== tools/test_validate_schemes.py (the validator's probes, mutations and instances)"
python3 "$repo/tools/test_validate_schemes.py" "$repo" 2>&1 | tail -3
[ "${PIPESTATUS[0]}" -eq 0 ] || status=1
echo "== tools/test_migrate_r2.py (the migration tool of release R2)"
python3 "$repo/tools/test_migrate_r2.py" 2>&1 | tail -1
[ "${PIPESTATUS[0]}" -eq 0 ] || status=1
echo "== tools/test_graph_from_library.py (graphs started from the library)"
python3 "$repo/tools/test_graph_from_library.py" 2>&1 | tail -1
[ "${PIPESTATUS[0]}" -eq 0 ] || status=1
echo "== tools/check_source_approvals.py (reviewed source decisions)"
python3 "$repo/tools/check_source_approvals.py" --repo "$repo" || status=1
echo "== tools/test_check_source_approvals.py (approval integrity mutations)"
python3 "$repo/tools/test_check_source_approvals.py" 2>&1 | tail -3
[ "${PIPESTATUS[0]}" -eq 0 ] || status=1
echo "== tools/vv/setup.sh (engine sources)"
WALTON_REPO="$repo" "$here/setup.sh" || exit 1
echo "== go vet, and the unit tests of the suite itself"
(cd "$here" && go vet ./... && go test -count=1 ./internal/...) || status=1
echo "== go test (engine and library)"
# No package argument: in that mode `go test` shows the summary by issue even
# when the run is green.
(cd "$here" && WALTON_REPO="$repo" go test -count=1 "$@") || status=1
exit $status
