#!/bin/bash
set -euo pipefail

BUNDLE=$(mktemp -d)
tar -xf - -C "${BUNDLE}" <&0

function clean_bundle {
	rm -r "${BUNDLE}"
}
trap clean_bundle EXIT

HUD_USER="hud"
HUD_PREFIX="/opt/hud"
HUD_APPS_DIR="${HUD_PREFIX}/apps"
HUD_BIN_DIR="/usr/bin"
HUD_DATA_DIR="/usr/share"

HUDCTL_APP_DIR="${HUD_APPS_DIR}/hudctl"

function mkdir_hud() {
    if [ ! -d "${1}" ]; then
	    mkdir -p "${1}"
        chown -R "${HUD_USER}:${HUD_USER}" "${1}"
    fi
}

mkdir_hud "${HUDCTL_APP_DIR}"
mkdir_hud "${HUD_APPS_DIR}"
mkdir -p "${HUD_BIN_DIR}"
mkdir -p "${HUD_DATA_DIR}"

rsync -ap --chown="${HUD_USER}:${HUD_USER}" --delete "${BUNDLE}/" "${HUDCTL_APP_DIR}"
chmod 755 "${HUDCTL_APP_DIR}"

HUD_APPS_CATALOG="${HUD_APPS_DIR}/catalog.json"
jq -nM \
    --arg app_name 'hudctl' \
    --arg app_manifest "${HUDCTL_APP_DIR}/manifest.json" \
    '[{ "name": $app_name, "manifest": $app_manifest, "enabled": false }]' \
    > "${HUD_APPS_CATALOG}"
chown "${HUD_USER}:${HUD_USER}" "${HUD_APPS_CATALOG}"
chmod 644 "${HUD_APPS_CATALOG}"

# Manually link the minimun amount needed for hudctl to link itself
ln -sf "${HUDCTL_APP_DIR}/binaries/hudctl" "${HUD_BIN_DIR}/hudctl"
ln -sf "${HUDCTL_APP_DIR}/binaries/hudctl-enable" "${HUD_BIN_DIR}/hudctl-enable"
ln -sf "${HUDCTL_APP_DIR}/binaries/hudctl-catalog" "${HUD_BIN_DIR}/hudctl-catalog"
ln -sf "${HUDCTL_APP_DIR}/binaries/hudctl-appenv" "${HUD_BIN_DIR}/hudctl-appenv"
ln -sf "${HUDCTL_APP_DIR}/data/logging.sh" "${HUD_DATA_DIR}/logging.sh"
ln -sf "${HUDCTL_APP_DIR}/data/json.sh" "${HUD_DATA_DIR}/json.sh"

chown -hR "${HUD_USER}:${HUD_USER}" "${HUD_PREFIX}"

PATH="${HUD_BIN_DIR}:${PATH}" hudctl enable hudctl