#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname "$(realpath "$0")")"
CROSS_FILE="${SCRIPT_DIR}/meson_cross.ini"

STEP_CONFIG="$1"
SOURCE="$(echo "${STEP_CONFIG}" | jq -r '.source')"
PREFIX="$(echo "${STEP_CONFIG}" | jq -r '.prefix')"

PATH="${BUILDROOT_HOST_DIR}/bin:${PATH}"

BUILD_DIR="${CACHE_DIR}/${APP_NAME}/meson"

if [ ! -d "${BUILD_DIR}" ]; then
    mkdir -p "${BUILD_DIR}"
    envsubst < "${CROSS_FILE}" > "${BUILD_DIR}/cross.ini"

    meson setup \
        --cross-file="${BUILD_DIR}/cross.ini" \
        --prefix="${PREFIX}" \
        "${BUILD_DIR}" \
        "${SOURCE}"
fi

meson install -C "${BUILD_DIR}" --quiet