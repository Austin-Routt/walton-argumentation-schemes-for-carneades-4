#!/usr/bin/env bash
# Check the engine sources this suite builds against.
#
#   tools/vv/setup.sh            check, and list how third_party/ differs
#
# The engine is kept as source in third_party/ at the repository root. It was
# imported byte for byte from the two archives at the repository root, which
# stay as the record of what upstream published. This script
#   1. refuses an archive whose SHA-256 or recorded commit is not the one in
#      sources.lock (whatever its archive comment says), and
#   2. lists every file of third_party/ that differs from its archive: these
#      are the engine fixes made in this repository, and nothing else should
#      appear. The two Windows executables of the GoCHR archive are not
#      imported and are not listed.
# It changes nothing.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo="${WALTON_REPO:-$(cd "$here/../.." && pwd)}"

lock="$repo/sources.lock"
field() { # engine key: read one value from sources.lock
  python3 -c 'import sys, yaml; print(yaml.safe_load(open(sys.argv[1]))[sys.argv[2]][sys.argv[3]])' "$lock" "$1" "$2"
}
want_c4="$(field carneades-4 commit)"; sha_c4="$(field carneades-4 sha256)"
want_chr="$(field GoCHR commit)";      sha_chr="$(field GoCHR sha256)"

check() { # file wanted-commit wanted-sha256
  local got_sha got_commit
  got_sha="$(sha256sum "$1" | cut -d' ' -f1)"
  got_commit="$(unzip -z "$1" | sed -n '2p')"
  [ "$got_sha" = "$3" ] || { echo "ERROR: $1 has sha256 $got_sha, expected $3" >&2; exit 1; }
  [ "$got_commit" = "$2" ] || { echo "ERROR: $1 is commit $got_commit, expected $2" >&2; exit 1; }
}
check "$repo/carneades-4.zip" "$want_c4" "$sha_c4"
check "$repo/GoCHR.zip" "$want_chr" "$sha_chr"
[ -f "$repo/third_party/carneades-4/go.mod" ] && [ -f "$repo/third_party/GoCHR/go.mod" ] \
  || { echo "ERROR: engine sources missing in $repo/third_party" >&2; exit 1; }

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
unzip -q "$repo/carneades-4.zip" -d "$tmp"
unzip -q "$repo/GoCHR.zip" -d "$tmp"
rm -f "$tmp/GoCHR-master/src/cmd/gochr/gochr.exe" "$tmp/GoCHR-master/src/engine/CHR/CHR.test.exe"
changed=0
for pair in "carneades-4-master:carneades-4" "GoCHR-master:GoCHR"; do
  from="$tmp/${pair%%:*}"; to="$repo/third_party/${pair##*:}"
  while IFS= read -r line; do
    echo "  third_party/${pair##*:}: $line"
    changed=$((changed + 1))
  done < <(diff -rq "$from" "$to" | sed -e "s|$from|archive|g" -e "s|$to|repository|g" || true)
done
echo "archives verified; $changed difference(s) between the archives and third_party/ (the engine fixes of this repository)"
