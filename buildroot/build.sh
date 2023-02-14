#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/../logging.sh

BUILDROOT_DIR="${SCRIPT_DIR}/.buildroot"

if [ ! -d "${BUILDROOT_DIR}" ]
then
  log_red "Buildroot not setup."
  exit 1
fi

pushd "${BUILDROOT_DIR}"
make
popd