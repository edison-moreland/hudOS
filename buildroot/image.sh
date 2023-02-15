#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

REALUSER="${USER}"
if [[ $EUID -eq 0 ]]
then
  if [[ -n "$SUDO_USER" ]]
  then
    REALUSER="${SUDO_USER}"
  fi
else
  log_red "This script must be run as root."
  exit 1
fi

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"
BUILDROOT_OUTPUT_DIR="${BUILDROOT_DIR}/output"

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
OUTPUT_IMAGE_MB="300"

if [ -f "${OUTPUT_IMAGE}" ]
then
  log_red "Overwriting old image!"
  rm "${OUTPUT_IMAGE}"
fi

log_blue "Reseving space for new image"
dd if=/dev/zero of="${OUTPUT_IMAGE}" bs=1M count=0 seek="${OUTPUT_IMAGE_MB}"
sync
LODEV=$(losetup --find)
losetup "${LODEV}" "${OUTPUT_IMAGE}"
function clean_lodev {
  log_yellow "Cleaning lodev"
  losetup -d "${LODEV}"
}
trap clean_lodev EXIT

log_blue "Partitioning..."
parted -s "${LODEV}" mklabel gpt
parted -s "${LODEV}" unit s mkpart primary 8192 25%
parted -s "${LODEV}" unit s mkpart primary 25% 100%
sync

log_blue "Formatting boot partition"
mkfs.vfat -n BOOT "${LODEV}p1"

log_blue "Copying rootfs partition to image"
pv < "${BUILDROOT_OUTPUT_DIR}/images/rootfs.ext4" > "${LODEV}p2"
sync

log_blue "Mounting boot partition"
MNT_BOOT="$(mktemp -d)"
mount "${LODEV}p1" "${MNT_BOOT}"
function clean_boot {
  log_yellow "Unmounting boot partition"
  umount "${MNT_BOOT}"
  rm -r "${MNT_BOOT}"
  clean_lodev
}
trap clean_boot EXIT

log_blue "Copying images"
KERNEL_IMAGE="${BUILDROOT_OUTPUT_DIR}/images/Image.gz"
INITRAMFS="${BUILDROOT_OUTPUT_DIR}/images/rootfs.cpio"
DEVICE_TREE="${BUILDROOT_OUTPUT_DIR}/images/sun50i-a64-pinephone-1.2.dtb"

cp "${KERNEL_IMAGE}" "${MNT_BOOT}/"
cp "${INITRAMFS}" "${MNT_BOOT}/initramfs-linux.img"
mkdir -p "${MNT_BOOT}/dtbs/allwinner/"
cp "${DEVICE_TREE}" "${MNT_BOOT}/dtbs/allwinner/"

log_blue "Generating boot script"
"${BUILDROOT_OUTPUT_DIR}"/host/bin/mkimage -A arm -O linux -T script -C none -n "U-Boot boot script" -d "${SCRIPT_DIR}/boot.txt" "${MNT_BOOT}/boot.scr"

log_blue "Finalizing!"
chown "${REALUSER}":"${REALUSER}" "${OUTPUT_IMAGE}"

# Fix fstab
# Copy boot files
#https://xnux.eu/howtos/install-arch-linux-arm.html
#https://github.com/megous/linux
#https://marcocetica.com/posts/buildroot-tutorial/
