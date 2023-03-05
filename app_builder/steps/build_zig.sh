#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
BUILD_FILE="$(echo "${STEP_CONFIG}" | jq -r '.build_file')"

# Make sure zig uses the pkg-config from buildroot
export PATH="${BUILDROOT_HOST_DIR}/bin:${PATH}"
zig build \
	--build-file "${BUILD_FILE}" \
    --cache-dir "${CACHE_DIR}/${APP_NAME}/zig-cache" \
    --global-cache-dir "${CACHE_DIR}/zig-global-cache" \
    --sysroot "${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot" \
    --search-prefix "${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot/usr" \
    -Dtarget="aarch64-linux-gnu"