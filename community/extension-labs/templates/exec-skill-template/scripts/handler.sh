#!/usr/bin/env bash
set -euo pipefail

payload=$(cat)
logger "[loqa-exec] received payload: $payload"

cat <<JSON
{"reply":"Hello from exec skill"}
JSON
