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

if [ ! -f "${SCRIPT_DIR}/hudOS.img" ]
then
  ln -s -t "${SCRIPT_DIR}/" "${BUILDROOT_DIR}/output/images/hudOS.img"
fi