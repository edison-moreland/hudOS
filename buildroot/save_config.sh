#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"

pushd "${BUILDROOT_DIR}"
make savedefconfig BR2_DEFCONFIG="${SCRIPT_DIR}/br2_hudos/configs/hudos_defconfig"
popd