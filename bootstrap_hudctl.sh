#!/usr/bin/env bash
set -euo pipefail

# This script will go away once hudctl is added to buildroot

REPO_ROOT="$(dirname "$(realpath "$0")")"
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_ROOT_SSH="deploy@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${REPO_ROOT}/buildroot/deploy_ed25519"

HUDCTL_BUNDLE="${REPO_ROOT}/.build/app/hudctl.tar"
HUDCTL_BUNDLE_DEPLOYER="${REPO_ROOT}/bootstrap_hudctl_deployer.sh"

log_blue "Building hudctl bundle"
"${REPO_ROOT}"/build.sh -i hudctl

log_blue "Deploying hudctl... "
scp -o StrictHostKeychecking=no -o UserKnownHostsFile=/dev/null -i "${PINEPHONE_SSH_KEY}" "${HUDCTL_BUNDLE_DEPLOYER}" "${PINEPHONE_ROOT_SSH}":/home/deploy/bootstrap_hudctl_deployer.sh
pv <"${HUDCTL_BUNDLE}" | ssh -o StrictHostKeychecking=no -o UserKnownHostsFile=/dev/null -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "/bin/sudo /bin/bash /home/deploy/bootstrap_hudctl_deployer.sh"
