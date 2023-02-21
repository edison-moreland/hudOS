#!/bin/bash

# Deploys an app_bundle.tar on the pinephone
# Use:
#   ssh root@pinephone "bundle_deployer.sh" < app_bundle.tar

# App bundle structure
# bin/
#    ...app binaries
# services/
#    ...systemd unit files
# post_deploy/
#    ...post_deploy scripts
# configs/
#    ...app configs

WORKSPACE=$(mktemp -d)
tar xf - -C "${WORKSPACE}" <&0

function clean_workspace {
	rm -r "${WORKSPACE}"
}
trap clean_workspace EXIT

HUD_USER="hud"
HUD_PREFIX="/opt/hud"
HUD_BIN_DIR="${HUD_PREFIX}/bin"
HUD_UNIT_DIR="${HUD_PREFIX}/systemd/system"
HUD_CONFIG_DIR="${HUD_PREFIX}/.config"

function sync_folder {
	SOURCE=$1
	DESTINATION=$2

	if [ ! -d "${DESTINATION}" ]; then
		mkdir -p "${DESTINATION}"
	fi
	rsync -ap --delete "${SOURCE}" "${DESTINATION}"
}

if [ -d "${WORKSPACE}/bin" ]; then
	echo "Syncing binaries"
	sync_folder "${WORKSPACE}/bin/" "${HUD_BIN_DIR}"
	chmod +x "${WORKSPACE}"/bin/*
else
	echo "No binaries to sync..."
fi

if [ -d "${WORKSPACE}/services" ]; then
	echo "Syncing systemd units"
	sync_folder "${WORKSPACE}/services/" "${HUD_UNIT_DIR}"
	systemctl daemon-reload
else
	echo "No systemd units to sync..."
fi

if [ -d "${WORKSPACE}/configs" ]; then
	echo "Syncing configs"
	sync_folder "${WORKSPACE}/configs/" "${HUD_CONFIG_DIR}"
else
	echo "No configs to sync..."
fi

chown -R ${HUD_USER}:${HUD_USER} ${HUD_PREFIX}

if [ -d "${WORKSPACE}/post_deploy" ]; then
	echo "Running post deploy scripts"
	for post_deploy_script in "${WORKSPACE}"/post_deploy/*.sh; do
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
