#!/usr/bin/env bash
# Run the engines' own unit tests on the sources in third_party/.
#
#   tools/vv/engine-tests.sh
#
# GoCHR's tests run as shipped. Carneades' tests run against the GoCHR of
# third_party/ (a replace directive in a copy of its go.mod; the module's own
# files are not touched). Three exclusions, each tracked under E-BUILD in
# known_issues.yml (PLAN.md 1.7):
#   -vet=off        go vet fails on caes.go as shipped
#   dung/...        two packages of the dung module no longer compile
#   TestCouchDB     needs a CouchDB server on localhost
# VV_GOMOD_EXTRA, as for the suite, names a file of extra go.mod lines
# (replace directives for an offline machine).
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo="${WALTON_REPO:-$(cd "$here/../.." && pwd)}"
tp="$repo/third_party"

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

echo "-- GoCHR"
(cd "$tp/GoCHR" && go test -vet=off -count=1 ./...)

echo "-- carneades-4 (against third_party/GoCHR)"
cp "$tp/carneades-4/go.mod" "$tmp/go.mod"
cp "$tp/carneades-4/go.sum" "$tmp/go.sum"
go mod edit -replace "github.com/hfried/GoCHR=$tp/GoCHR" "$tmp/go.mod"
if [ -n "${VV_GOMOD_EXTRA:-}" ]; then cat "$VV_GOMOD_EXTRA" >> "$tmp/go.mod"; fi
(cd "$tp/carneades-4" && go test -mod=mod -modfile="$tmp/go.mod" -vet=off -count=1 \
  ./src/engine/caes/... ./src/engine/terms/... ./src/engine/validation/... ./src/test/ \
  -skip '^TestCouchDB$')
