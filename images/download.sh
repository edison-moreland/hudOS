#!/bin/bash
set -euo pipefail

IMAGES_DIR="$(dirname $(realpath $0))"
source "${IMAGES_DIR}"/../logging.sh

ARCH_IMAGE="${IMAGES_DIR}/archlinux-pinephone-barebone-20230203.img"
ARCH_IMAGE_SHA="236ece93a51eed4a2ca8f3b12d91da424875e7db465bfe5e2cf99c7cb02dd89d"
ARCH_DOWNLOAD_URL="https://github.com/dreemurrs-embedded/Pine64-Arch/releases/download/20230203/archlinux-pinephone-barebone-20230203.img.xz"
ARCH_DOWNLOAD_SHA="b74ea1cc3632e827fa01eb25cfc04819ab468ca0323cd922dd3082da61a08ef0"

TOWBOOT_IMAGE="${IMAGES_DIR}/mmcboot.installer.img"
TOWBOOT_IMAGE_SHA="d3666e31001dc87021206ab253005bb9a7850ecae06097bfcf8ae5499e635c04"
TOWBOOT_DOWNLOAD_URL="https://github.com/Tow-Boot/Tow-Boot/releases/download/release-2021.10-005/pine64-pinephoneA64-2021.10-005.tar.xz"
TOWBOOT_DOWNLOAD_SHA="b9c004362a11a8baa73b5ed7b708291e3c9340c0d5e35ea573dc6f679a6c663e"


function checksha {
  TARGET_FILE=$1
  EXPECTED_SHA=$2

  if [ ! -f "${TARGET_FILE}" ]
  then
    return 1
  fi

  if [[ $(shasum -a 256 "${TARGET_FILE}") != "${EXPECTED_SHA}" ]]
  then
    return 0
  else
    return 1
  fi
}


ARCH_DID_UPDATE="false"
if checksha "${ARCH_IMAGE}" "${ARCH_IMAGE_SHA}"
then
  log_blue "Arch image is up to date!"
else
  log_green "Updating Arch image"
  ARCH_DID_UPDATE="true"

  arch_download="$(mktemp)"
  function arch_cleanup {
    log "Cleaning temporary files..."
    rm -f "${arch_download}"
  }
  trap arch_cleanup EXIT

  wget "${ARCH_DOWNLOAD_URL}" --show-progress -O "${arch_download}"

  if ! checksha "${arch_download}" "${ARCH_DOWNLOAD_SHA}"
  then
    log_red "Downloaded image doesn't match expected sha!"
    exit 1
  fi

  if [ -f "${ARCH_IMAGE}" ]
  then
    log "Deleting old image..."
    rm "${ARCH_IMAGE}"
  fi

  log "Extracting new image..."
  xz -dc "${arch_download}" | pv > "${ARCH_IMAGE}"
fi

if checksha "${TOWBOOT_IMAGE}" "${TOWBOOT_IMAGE_SHA}"
then
  log_blue "Tow-Boot image is up to date!"
else
  log_green "Updating Tow-Boot image"

  towboot_download="$(mktemp)"
  function towboot_cleanup {
    log "Cleaning temporary files..."
    rm -f "${towboot_download}"
  }
  trap towboot_cleanup EXIT

  wget "${TOWBOOT_DOWNLOAD_URL}" --show-progress -O "${towboot_download}"

  if ! checksha "${towboot_download}" "${TOWBOOT_DOWNLOAD_SHA}"
  then
    log_red "Downloaded image doesn't match expected sha!"
    exit 1
  fi

  if [ -f "${TOWBOOT_IMAGE}" ]
  then
    log "Deleting old image..."
    rm "${TOWBOOT_IMAGE}"
  fi

  log "Extracting new image..."
  image_name=$(tar -tf "${towboot_download}" | grep "mmcboot.installer.img")
  tar -xf "${towboot_download}" "${image_name}" -O | pv > "${TOWBOOT_IMAGE}"
fi

if [[ "${ARCH_DID_UPDATE}" == "true" ]]
then
  log_yellow "Arch image was updated, don't forget to rebuild chroot!"
fi