#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

if [[ $EUID -eq 0 ]]
then
  log_red "This script should NOT be ran as root!"
  exit 1
fi

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"

if [ ! -d "${BUILDROOT_DIR}" ]
then
  log_red "Buildroot not setup."
  exit 1
fi

# make BR2_EXTERNAL=/path/to/foo menuconfig

log_blue "Running buildroot"
log_yellow "THIS MAY TAKE A LONG TIME!"
pushd "${BUILDROOT_DIR}"
make "$@"
popd

BUILDROOT_IMAGE_OUT="${BUILDROOT_DIR}/output/images/hudOS.img"
if [ -f "${BUILDROOT_IMAGE_OUT}"  ]
then
  log_blue "Buildroot produced a new image!"
  if [ -f "${SCRIPT_DIR}/hudOS.img" ]
  then
    rm "${SCRIPT_DIR}/hudOS.img"
  fi
  mv "${BUILDROOT_IMAGE_OUT}" "${SCRIPT_DIR}/hudOS.img"
fi

if [ "${BUILDROOT_DIR}/.config" -nt "${SCRIPT_DIR}/br2_hudos/configs/hudos_defconfig" ]
then
  read -p "Save config changes to hudos_defconfig? [Y/n]: " -n 1 -r
  echo
  if [[ ! $REPLY =~ ^[Nn]$ ]]
  then
    "${SCRIPT_DIR}"/save_config.sh
  fi
fi

# TODO: Pull this from images dir
BUILDROOT_DEPLOY_KEY_OUT="${BUILDROOT_DIR}/output/build/hud-landing-pad-0.0.1/deploy_ed25519"
if [ "${BUILDROOT_DEPLOY_KEY_OUT}" -nt "${SCRIPT_DIR}/deploy_ed25519" ]
then
  log_blue "Deploy key updated"
  cp "${BUILDROOT_DEPLOY_KEY_OUT}" "${SCRIPT_DIR}"
  cp "${BUILDROOT_DEPLOY_KEY_OUT}.pub" "${SCRIPT_DIR}"
fi