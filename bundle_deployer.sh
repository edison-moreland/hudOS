#!/bin/bash
set -euo pipefail

# Deploys an app_bundle.tar on the pinephone
# Use:
#   ssh root@pinephone "bundle_deployer.sh" < app_bundle.tar

# App bundle structure
# binaries/
#    ...app binaries
# units/
#    ...systemd unit files
# scripts/
#    ...post_deploy scripts
# configs/
#    ...app configs
# data/
#    ...app resources/static data

BUNDLE=$(mktemp -d)
tar xf - -C "${BUNDLE}" <&0

function clean_bundle {
	rm -r "${BUNDLE}"
}
trap clean_bundle EXIT

HUD_USER="hud"
HUD_PREFIX="/opt/hud"
HUD_BIN_DIR="${HUD_PREFIX}/bin"
HUD_UNIT_DIR="${HUD_PREFIX}/systemd/system"
HUD_CONFIG_DIR="${HUD_PREFIX}/.config"
HUD_DATA_DIR="${HUD_PREFIX}/.local/share"

function sync_bundle_dir {
	BUNDLE_DIR=$1
	DESTINATION_DIR=$2
	
	full_bundle_dir="${BUNDLE}/${BUNDLE_DIR}/"
	if [ -d "${full_bundle_dir}" ]; then
		echo "Syncing ${BUNDLE_DIR}"
		
		if [ ! -d "${DESTINATION_DIR}" ]; then
			mkdir -p "${DESTINATION_DIR}"
		fi
		rsync -ap --delete "${full_bundle_dir}" "${DESTINATION_DIR}"
		
		return 0
	fi

	echo "No ${BUNDLE_DIR} to sync..."
	return 1
}

if sync_bundle_dir "binaries" "${HUD_BIN_DIR}"; then
	chmod +x "${HUD_BIN_DIR}"/*
fi

if sync_bundle_dir "units" "${HUD_UNIT_DIR}"; then
	systemctl daemon-reload
fi

if sync_bundle_dir "configs" "${HUD_CONFIG_DIR}"; then
	: # Intentionally left blank
fi

if sync_bundle_dir "data" "${HUD_DATA_DIR}"; then 
	: # Intentionally left blank
fi

chown -R ${HUD_USER}:${HUD_USER} ${HUD_PREFIX}

bundle_scripts="${BUNDLE}/scripts"
if [ -d "${bundle_scripts}" ]; then
	echo "Running post deploy scripts"
	for post_deploy_script in "${bundle_scripts}"/*.sh; do
		echo "- $(basename "${post_deploy_script}")"
		if bash "${post_deploy_script}"; then
			echo "Success!"
		else
			echo "Failure :("
			exit 1
		fi
	done
else
	echo "No post deploy scripts to run..."
fi

echo "Restarting UI"
systemctl restart hud.target
systemctl restart hud-apps.target
