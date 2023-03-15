#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"
VENDOR_BUILDROOT_DIR="$(realpath "${SCRIPT_DIR}/../.build/vendor/buildroot")"

if [ -d "${BUILDROOT_DIR}" ]; then
	log_red "Buildroot already setup!"

	read -p "Do you want to overwrite it? [y/N]: " -n 1 -r
	echo
	if [[ ! $REPLY =~ ^[Yy]$ ]]; then
		exit 1
	fi

	log_blue "Removing old buildroot..."
	rm -r "${BUILDROOT_DIR}"
fi

if [ ! -d "${VENDOR_BUILDROOT_DIR}" ]; then
	"${SCRIPT_DIR}/../update_vendor.sh"
fi

ln -s "${VENDOR_BUILDROOT_DIR}" "${BUILDROOT_DIR}"

pushd "${BUILDROOT_DIR}"
make BR2_EXTERNAL="${SCRIPT_DIR}/br2_hudos" hudos_defconfig
popd
