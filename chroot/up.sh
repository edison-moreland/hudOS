#!/bin/bash
set -euo pipefail

# https://nerdstuff.org/posts/2020/2020-003_simplest_way_to_create_an_arm_chroot/
# https://man.archlinux.org/man/extra/arch-install-scripts/arch-chroot.8.en

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BASE_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-barebone-20230203.img"
CHROOT_ARCH_IMAGE="${SCRIPT_DIR}/../images/archlinux-pinephone-chroot.img"
CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if [ ! -f "${BASE_ARCH_IMAGE}" ]
then
  log_red "No Arch image found!"
  read -p "Do you want to download images now? [Y/n]: " -n 1 -r
  echo
  if [[ $REPLY =~ ^[Nn]$ ]]
  then
    exit 1
  fi

  "${SCRIPT_DIR}"/../images/download.sh
fi

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
  log_blue "Making copy of Arch image"
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
  sudo cp "${SCRIPT_DIR}/setup.sh" "${CHROOT_DIR}/root/setup.sh"
  sudo arch-chroot "${CHROOT_DIR}" /root/setup.sh
  sudo mount --bind "${SCRIPT_DIR}/../" "${CHROOT_DIR}/opt/build/"
else
  log_blue "Skipping first time setup"
fi