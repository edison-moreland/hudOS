#!/bin/bash
set -euo pipefail

STEPS_DIR="$(dirname $(realpath $0))"
BUILDROOT_HOST="${STEPS_DIR}/../../buildroot/.buildroot/output/host"
if [ ! -d "${BUILDROOT_HOST}" ]; then
	log_red "Buildroot hasn't been ran!"
	exit 1
fi

STEP_CONFIG="$1"
TARGET="$(echo "${STEP_CONFIG}" | jq -r '.target')"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"

export GOOS=linux
export GOARCH=arm64
export GOARM=7 # Pinephone processor is Armv8-A, but go only supports up to v7
export CGO_ENABLED=1
export PATH="${PATH}:${BUILDROOT_HOST}/bin"
#export CGO_CFLAGS="--sysroot ${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot"
#export CGO_LDFLAGS="--sysroot ${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot"
export AR=aarch64-none-linux-gnu-ar
export CC=aarch64-none-linux-gnu-gcc
export CXX=aarch64-none-linux-gnu-g++
export FC=aarch64-none-linux-gnu-gfortran
go build -o "${OUT}" "${TARGET}"