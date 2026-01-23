#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BASELINE_FILE="$ROOT_DIR/.coverage_baseline"
MIN="${COVERAGE_MIN:-0}"

# Generate fresh coverage output (also acts as proof artifact in CI logs)
TOTAL_LINE="$(bash "$ROOT_DIR/scripts/coverage.sh")"
echo "Coverage summary: $TOTAL_LINE"

# Extract numeric percent (e.g. 12.3 from 'total: (statements) 12.3%')
CURRENT="$(echo "$TOTAL_LINE" | awk '{print $3}' | tr -d '%')"

# Compare against minimum
awk -v cur="$CURRENT" -v min="$MIN" 'BEGIN {
  if (cur+0 < min+0) {
    printf("FAIL: coverage %.2f%% is below minimum %.2f%%\n", cur, min);
    exit 1;
  } else {
    printf("OK: coverage %.2f%% meets minimum %.2f%%\n", cur, min);
  }
}'

# Compare against baseline if file exists
if [[ -f "$BASELINE_FILE" ]]; then
  BASE="$(cat "$BASELINE_FILE" | tr -d ' %\n\r\t')"
  if [[ -n "$BASE" ]]; then
    awk -v cur="$CURRENT" -v base="$BASE" 'BEGIN {
      if (cur+0 <= base+0) {
        printf("FAIL: coverage %.2f%% did not increase vs baseline %.2f%%\n", cur, base);
        exit 1;
      } else {
        printf("OK: coverage %.2f%% increased vs baseline %.2f%%\n", cur, base);
      }
    }'
  fi
else
  echo "NOTE: No .coverage_baseline found; skipping 'increase vs baseline' check."
fi
