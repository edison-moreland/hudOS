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
VENDOR_DIR="$6"
CACHE_DIR="$7"

export PATH="${BUILDROOT_HOST}/bin:${PATH}"
# Take a page out of bazels book, copy everything into the workspace, link the vendor directory into the same folder
# Magic
 
cp -ar "${APP_DIR}"/* "${RULE_WORKSPACE}/"
ln -s "${VENDOR_DIR}" "${RULE_WORKSPACE}/vendor"

zig build \
	--build-file "${RULE_WORKSPACE}/build.zig" \
    --cache-dir "${CACHE_DIR}/${APP_NAME}" \
    --sysroot "${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot" \
    --search-prefix "${BUILDROOT_HOST}/aarch64-buildroot-linux-gnu/sysroot/usr" \
    -Dtarget="aarch64-linux-gnu" \

cat <<EOF | envsubst >"${RULE_WORKSPACE}/hud-${APP_NAME}.service"
[Unit]
Description=${APP_NAME}, a HUD Application
ConditionPathExists=/opt/hud/run/wayland-0
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
EOF

# TODO: Make a script to do this \/
BUNDLE_CONFIG=$(jq -n -c \
	--arg unit "hud-${APP_NAME}.service" \
	--arg binary "zig-out/bin/${APP_NAME}" \
	--arg script "050-${APP_NAME}.sh" \
	'{"units": [$unit], "binaries": [$binary], "scripts": [$script]}' |
	base64 -w 0)

"${RULES_DIR}"/raw.sh \
	"${APP_NAME}" \
	"${RULE_WORKSPACE}" \
	"${BUNDLE_CONFIG}" \
	"${RULE_WORKSPACE}" \
	"${RULE_BUNDLE_OUT}"
