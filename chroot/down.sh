#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BASE_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-barebone-20230203.img"
CHROOT_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-chroot.img"
CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if ! "${SCRIPT_DIR}"/is-up.sh
then
  log_red "Chroot is not up!"
  exit 1
fi

LOOP_DEV=$(sudo losetup --associated "${CHROOT_ARCH_IMAGE}" | cut -d':' -f1)

log_blue "Destroying loop device"
sudo umount -d "${CHROOT_DIR}"
sudo losetup -d "${LOOP_DEV}"