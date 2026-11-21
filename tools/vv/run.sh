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
echo "== tools/vv/setup.sh (engine sources)"
WALTON_REPO="$repo" "$here/setup.sh" || exit 1
echo "== go test (engine and library)"
# No package argument: in that mode `go test` shows the summary by issue even
# when the run is green.
(cd "$here" && WALTON_REPO="$repo" go test -count=1 "$@") || status=1
exit $status
