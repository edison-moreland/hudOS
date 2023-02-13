#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BASE_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-barebone-20230203.img"
CHROOT_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-chroot.img"
CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if [ -d "${CHROOT_DIR}" ]
then
  if mount | grep "on ${CHROOT_DIR} type" > /dev/null
  then
    exit 0
  fi
fi

exit 1