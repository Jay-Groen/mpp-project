#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TMP_DIR="$ROOT_DIR/.tmp"
COVERPROFILE="$TMP_DIR/coverage.out"
COVER_TXT="$TMP_DIR/coverage.txt"
COVER_HTML="$TMP_DIR/coverage.html"

mkdir -p "$TMP_DIR"

# Run tests with coverage for all packages
go test ./... -count=1 -covermode=atomic -coverprofile="$COVERPROFILE" >/dev/null

# Human-readable summary
go tool cover -func="$COVERPROFILE" | tee "$COVER_TXT" >/dev/null

if [[ "${1:-}" == "--html" ]]; then
  go tool cover -html="$COVERPROFILE" -o "$COVER_HTML"
  echo "HTML report written to: $COVER_HTML"
fi

# Print the total coverage line as proof
tail -n 1 "$COVER_TXT"
