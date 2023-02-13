#!/bin/bash
set -euo pipefail

# This script is ran inside the chroot to get it ready for development
pacman-key --init
pacman-key --populate archlinuxarm

#full system update might be too heavy. Maybe once build is more stable
#pacman -Syu --noconfirm
#pacman -S --needed --noconfirm base-devel

# Install bazelisk
BAZELISK_URL="https://github.com/bazelbuild/bazelisk/releases/download/v1.16.0/bazelisk-linux-arm64"
BAZELISK_SHA="8cc337c69b6e6a71aac86c7e3514c29b2f18d384c0ed991fd3ea2ac7762584ee"

bazelisk_download="$(mktemp)"
function bazelisk_cleanup {
  if [ -f "${bazelisk_download}" ]
  then
    echo "Cleaning temporary files..."
#    rm -f "${bazelisk_download}"
  fi
}
trap bazelisk_cleanup EXIT

curl -L "${BAZELISK_URL}" > "${bazelisk_download}"

bazelisk_download_sha=$(sha256sum "${bazelisk_download}" | cut -d' ' -f1)
if [[ "${bazelisk_download_sha}" != "${BAZELISK_SHA}" ]]
then
  echo "SHA mismatch!"
  echo "bazelisk-linux-arm64 sha256sum '${bazelisk_download_sha}' != '${BAZELISK_SHA}'"
  exit 1
fi

mv "${bazelisk_download}" /usr/bin/bazel
chmod +x "/usr/bin/bazel"

# Setup build user and build area
useradd -U -m -d /home/build -r build
mkdir -p /opt/build
chown build:build /opt/build