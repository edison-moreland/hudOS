#!/usr/bin/env bash
set -euo pipefail

STEPS_DIR="$(dirname $(realpath $0))"
VENDOR_DIR="${STEPS_DIR}/../../.build/vendor"

STEP_CONFIG="$1"
TARGET="$(echo "${STEP_CONFIG}" | jq -r '.target')"

ln -s "${VENDOR_DIR}" "${TARGET}"