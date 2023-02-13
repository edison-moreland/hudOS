#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BASE_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-barebone-20230203.img"
CHROOT_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-chroot.img"
CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if "${SCRIPT_DIR}"/is-up.sh
then
  log_red "Chroot is currently up! Running down.sh"
  "${SCRIPT_DIR}"/down.sh
fi

if [ -f "${CHROOT_ARCH_IMAGE}" ]
then
  log_blue "Destroying image"
  rm "${CHROOT_ARCH_IMAGE}"
fi

if [ -d "${CHROOT_DIR}" ]
then
  log_blue "Destroying directory"
  rm -r "${CHROOT_DIR}"
fi