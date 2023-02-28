#!/usr/bin/env bash
set -euo pipefail

STEPS_DIR="$(dirname $(realpath $0))"
BUILDROOT_HOST="${STEPS_DIR}/../../buildroot/.buildroot/output/host"
if [ ! -d "${BUILDROOT_HOST}" ]; then
	log_red "Buildroot hasn't been ran!"
	exit 1
fi

STEP_CONFIG="$1"
BUILD_FILE="$(echo "${STEP_CONFIG}" | jq -r '.build_file')"

# Make sure zig uses the pkg-config from buildroot
export PATH="${BUILDROOT_HOST}/bin:${PATH}"
zig build \
	--build-file "${BUILD_FILE}" \
    --cache-dir "${CACHE_DIR}/${APP_NAME}/zig-cache" \
    --global-cache-dir "${CACHE_DIR}/zig-global-cache" \
    --sysroot "${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot" \
    --search-prefix "${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot/usr" \
    -Dtarget="aarch64-linux-gnu"