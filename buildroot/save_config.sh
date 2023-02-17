#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"

DEFCONFIG_PATH="${SCRIPT_DIR}/br2_hudos/configs/hudos_defconfig"

log_blue "Saving configs to git repo"
pushd "${BUILDROOT_DIR}"
make savedefconfig BR2_DEFCONFIG="${DEFCONFIG_PATH}"
popd

# These settings should not be saved to git
log_blue "Filtering sensitive keys"
sed -i '/^BR2_PACKAGE_HUD_NETWORK_WIFI_/d' "${DEFCONFIG_PATH}"
sed -i '/^BR2_PACKAGE_HUD_DEPLOY_USER/d' "${DEFCONFIG_PATH}"
