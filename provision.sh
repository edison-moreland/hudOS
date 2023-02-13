#!/usr/bin/env bash
set -euo pipefail

# Note: This script should be safe to run multiple times, but it will redo a lot of work
# Use: `./provision.sh '<pinephone_ip>

SCRIPT_DIR="$(dirname $(realpath $0))"
source "${SCRIPT_DIR}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_ALARM_SSH="alarm@${PINEPHONE_HOST}"
PINEPHONE_ROOT_SSH="root@${PINEPHONE_HOST}"

log_blue "Syncing provision scripts to pinephone"
PROVISION_PREFIX="/home/alarm"
PROVISION_SCRIPTS="${PROVISION_PREFIX}/provision"
scp -r "${SCRIPT_DIR}/provision" "${PINEPHONE_ALARM_SSH}":"${PROVISION_PREFIX}"

# Generate an ssh key that can be used to deploy the hud
# Right now this gets put under root
DEPLOY_SSH_KEY="${SCRIPT_DIR}/.deploy/id_ed25519"
DEPLOY_SSH_KEY_PUB="${DEPLOY_SSH_KEY}.pub"
if [ ! -d "${SCRIPT_DIR}/.deploy" ]
then
  log_blue "Generating SSH keys"
  mkdir "${SCRIPT_DIR}/.deploy"
  ssh-keygen -t ed25519 -C "PinePhone deploy key" -f "${DEPLOY_SSH_KEY}"
fi

log_blue "Installing SSH key"
# shellcheck disable=SC2029
ssh "${PINEPHONE_ALARM_SSH}" "${PROVISION_SCRIPTS}/install_ssh_key.sh $(base64 < "${DEPLOY_SSH_KEY_PUB}")"

log_blue "Installing Weston"
ssh -i "${DEPLOY_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "${PROVISION_SCRIPTS}/install_weston.sh"

log_blue "Creating hud user"
ssh -i "${DEPLOY_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "${PROVISION_SCRIPTS}/create_hud_user.sh"

log_blue "Updating OS"
ssh -i "${DEPLOY_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "${PROVISION_SCRIPTS}/update.sh"

log_blue "Cleaning up + rebooting"
ssh -i "${DEPLOY_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "rm -r ${PROVISION_SCRIPTS}; reboot now"

# If things worked, pinephone should log in as hud automatically
