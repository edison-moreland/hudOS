#!/bin/bash
set -euo pipefail

# https://nerdstuff.org/posts/2020/2020-003_simplest_way_to_create_an_arm_chroot/
# https://man.archlinux.org/man/extra/arch-install-scripts/arch-chroot.8.en

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BASE_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-barebone-20230203.img"
CHROOT_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-chroot.img"
CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if "${SCRIPT_DIR}"/is-up.sh
then
  log_red "Chroot is already up!"
  exit 1
fi

FIRST_SETUP="false"
LOOP_DEV=$(sudo losetup --find)
if [ -f "${CHROOT_ARCH_IMAGE}" ]
then
  log_blue "Skipping copy"
else
  log_blue "Making copy of arch image"
  FIRST_SETUP="true"
  cp "${BASE_ARCH_IMAGE}" "${CHROOT_ARCH_IMAGE}"
fi

if [ ! -d "${CHROOT_DIR}" ]
then
  mkdir "${CHROOT_DIR}"
fi

log_blue "Mounting image through ${LOOP_DEV}"
sudo losetup -P "${LOOP_DEV}" "${CHROOT_ARCH_IMAGE}"
sudo mount "${LOOP_DEV}p2" "${CHROOT_DIR}"

if [[ "${FIRST_SETUP}" == "true" ]]
then
  log_blue "Running first time setup"
  sudo arch-chroot "${CHROOT_DIR}" /bin/bash << EOF
pacman-key --init
pacman-key --populate archlinuxarm

pacman -Syu --noconfirm
pacman -S --needed --noconfirm base-devel
EOF
else
  log_blue "Skipping first time setup"
fi