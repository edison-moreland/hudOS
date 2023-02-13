#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

CHROOT_DIR="${SCRIPT_DIR}/.chroot"

if ! "${SCRIPT_DIR}"/is-up.sh
then
  log_red "Chroot is not up!"
  read -p "Do you want to start it now? [Y/n]: " -n 1 -r
  echo
  if [[ $REPLY =~ ^[Nn]$ ]]
  then
    exit 1
  fi

  "${SCRIPT_DIR}"/up.sh
fi

sudo arch-chroot "${CHROOT_DIR}"
