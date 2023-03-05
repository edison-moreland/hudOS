#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
SOURCE="$(echo "${STEP_CONFIG}" | jq -r '.source')"
TARGET="$(echo "${STEP_CONFIG}" | jq -r '.target')"

ln -s "${SOURCE}" "${TARGET}"