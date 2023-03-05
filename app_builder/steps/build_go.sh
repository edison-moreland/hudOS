#!/bin/bash
set -euo pipefail

STEP_CONFIG="$1"
TARGET="$(echo "${STEP_CONFIG}" | jq -r '.target')"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"

GO="${VENDOR_DIR}/go/go/bin/go"

export GOOS=linux
export GOARCH=arm64
export GOARM=7 # Pinephone processor is Armv8-A, but go only supports up to v7
export CGO_ENABLED=1
export PATH="${PATH}:${BUILDROOT_HOST_DIR}/bin"
#export CGO_CFLAGS="--sysroot ${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot"
#export CGO_LDFLAGS="--sysroot ${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot"
export AR=aarch64-none-linux-gnu-ar
export CC=aarch64-none-linux-gnu-gcc
export CXX=aarch64-none-linux-gnu-g++
export FC=aarch64-none-linux-gnu-gfortran
$GO build -o "${OUT}" "${TARGET}"