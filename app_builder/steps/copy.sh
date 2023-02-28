#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
SOURCE="$(echo "${STEP_CONFIG}" | jq -r '.source')"
DESTINATION="$(echo "${STEP_CONFIG}" | jq -r '.destination')"

mkdir -p "$(dirname "${DESTINATION}")"

if [ -d "${SOURCE}" ]; then
    cp -ar "${SOURCE}"/* "${DESTINATION}"
else
    cp -a "${SOURCE}" "${DESTINATION}"
fi