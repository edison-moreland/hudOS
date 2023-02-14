#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

BUILDROOT_DIR="${REPO_ROOT}/buildroot"
BUILDROOT_URL="https://buildroot.org/downloads/buildroot-2022.11.tar.xz"


if [ -d "${BUILDROOT_DIR}" ]
then
  log_red "Buildroot already setup!"

  read -p "Do you want to overwrite it? [y/N]: " -n 1 -r
  echo    # (optional) move to a new line
  if [[ ! $REPLY =~ ^[Yy]$ ]]
  then
      exit 1
  fi

  log_blue "Removing old buildroot..."
  rm -r "${BUILDROOT_DIR}"
fi

mkdir "${BUILDROOT_DIR}"

curl "${BUILDROOT_URL}" | tar -xJf - -C "${BUILDROOT_DIR}" --strip-components=1