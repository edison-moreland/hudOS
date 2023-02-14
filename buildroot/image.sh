#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

REALUSER="${USER}"
if [[ $EUID -eq 0 ]]; then
  if [[ -n "$SUDO_USER" ]]; then
    REALUSER="${SUDO_USER}"
  fi
else
  log_red "This script must be run as root."
  exit 1
fi

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"
BUILDROOT_OUTPUT_DIR="${BUILDROOT_DIR}/output/images"

if [ ! -d "${BUILDROOT_DIR}" ]
then
  log_red "Buildroot has not been setup!"
  exit 1
fi

if [ ! -d "${BUILDROOT_OUTPUT_DIR}" ]
then
  log_red "Buildroot has not been ran!"
  exit 1
fi

OUTPUT_IMAGE="${SCRIPT_DIR}/hudOS.img"
OUTPUT_IMAGE_MB="3072"

if [ -f "${OUTPUT_IMAGE}" ]
then
  log_red "Overwriting old image!"
  rm "${OUTPUT_IMAGE}"
fi

log_blue "Reseving space for new image"
dd if=/dev/zero of="${OUTPUT_IMAGE}" bs=1M count=0 seek="${OUTPUT_IMAGE_MB}"
LODEV=$(losetup --find)
losetup "${LODEV}" "${OUTPUT_IMAGE}"
function clean_lodev {
  losetup -d "${LODEV}"
}
trap clean_lodev EXIT

log_blue "Partitioning..."
parted -s "${LODEV}" mklabel msdos
parted -s "${LODEV}" unit s mkpart primary 8192 491519
parted -s "${LODEV}" unit s mkpart primary 524288 100%

log_blue "Formatting boot partition"
mkfs.vfat -n BOOT "${LODEV}p1"

log_blue "Copying rootfs partition to image"
pv < "${BUILDROOT_OUTPUT_DIR}/rootfs.ext4" > "${LODEV}p2"

MNT_ROOT="$(mktemp -d)"
mount "${LODEV}p2" "${MNT_ROOT}"
mkdir "${MNT_ROOT}/boot"
mount "${LODEV}p1" "${MNT_ROOT}/boot"
function clean_mnt {
  umount "${MNT_ROOT}/boot"
  rm -r "${MNT_ROOT:?}/boot"
  umount "${MNT_ROOT}"
}
trap clean_mnt EXIT


ls -la "${MNT_ROOT}"

# Fix fstab
# Copy boot files
#https://xnux.eu/howtos/install-arch-linux-arm.html
#https://github.com/megous/linux
#https://marcocetica.com/posts/buildroot-tutorial/