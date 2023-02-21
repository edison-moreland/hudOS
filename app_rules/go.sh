#!/usr/bin/env bash
set -euo pipefail

RULES_DIR="$(dirname $(realpath $0))"
REPO_ROOT="${RULES_DIR}/.."
source "${REPO_ROOT}"/logging.sh

BUILDROOT_HOST="${REPO_ROOT}/buildroot/.buildroot/output/host"
if [ ! -d "${BUILDROOT_HOST}" ]; then
	log_red "Buildroot hasn't been ran!"
	exit 1
fi

APP_NAME=$1
APP_DIR=$2
RULE_CONFIG=$(echo "$3" | base64 -d)
RULE_WORKSPACE=$4
RULE_BUNDLE_OUT=$5

SETCAP=$(echo "${RULE_CONFIG}" | jq -r '.setcap // ""')

export GOOS=linux
export GOARCH=arm64
export GOARM=7 # Pinephone processor is Armv8-A, but go only supports up to v7
export CGO_ENABLED=1
export PATH="${PATH}:${BUILDROOT_HOST}/bin"
#export CGO_CFLAGS="--sysroot ${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot"
#export CGO_LDFLAGS="--sysroot ${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot"
export AR=aarch64-none-linux-gnu-ar
export CC=aarch64-none-linux-gnu-gcc
export CXX=aarch64-none-linux-gnu-g++
export FC=aarch64-none-linux-gnu-gfortran
go build -o "${RULE_WORKSPACE}/${APP_NAME}" "${APP_DIR}"

cat <<EOF | envsubst >"${RULE_WORKSPACE}/hud-${APP_NAME}.service"
[Unit]
Description=${APP_NAME}, a HUD Application
PartOf=hud-apps.target

[Service]
User=hud
Group=hud
WorkingDirectory=/opt/hud

Slice=hud.slice
Type=simple
Environment=XDG_RUNTIME_DIR=/opt/hud/run XCOMPOSEFILE=/dev/null
ExecStart=/opt/hud/bin/${APP_NAME}
Restart=always

[Install]
WantedBy=hud-apps.target
EOF

cat <<EOF | envsubst >"${RULE_WORKSPACE}/050-${APP_NAME}.sh"
#!/bin/bash

systemctl enable hud-${APP_NAME}.service

if [[ "${SETCAP}" != "" ]]
then
  setcap '${SETCAP}' /opt/hud/bin/${APP_NAME}
fi
EOF

# TODO: Make a script to do this \/
BUNDLE_CONFIG=$(jq -n -c \
	--arg unit "hud-${APP_NAME}.service" \
	--arg binary "${APP_NAME}" \
	--arg script "050-${APP_NAME}.sh" \
	'{"units": [$unit], "binaries": [$binary], "scripts": [$script]}' |
	base64 -w 0)

"${RULES_DIR}"/raw.sh \
	"${APP_NAME}" \
	"${RULE_WORKSPACE}" \
	"${BUNDLE_CONFIG}" \
	"${RULE_WORKSPACE}" \
	"${RULE_BUNDLE_OUT}"
