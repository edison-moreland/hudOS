#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_ROOT_SSH="deploy@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${REPO_ROOT}/buildroot/deploy_ed25519"

APP_BUNDLE="${REPO_ROOT}/bazel-bin/app_bundle.tar"
BUNDLE_DEPLOYER="${REPO_ROOT}/bundle_deployer.sh"

if [ ! -f "${PINEPHONE_SSH_KEY}" ]
then
  log_red "Deploy key not generated. Did you enable deploy user in buildroot?"
  exit 1
fi

log_blue "Building App Bundle"
bazel build '//:app_bundle'

log_blue "Deploying App Bundle"
scp -i "${PINEPHONE_SSH_KEY}" "${BUNDLE_DEPLOYER}" "${PINEPHONE_ROOT_SSH}":/home/deploy/bundle_deployer.sh
pv < "${APP_BUNDLE}" | ssh -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "/bin/bash /home/deploy/bundle_deployer.sh"

log_blue "Deployment Status"
ssh -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "systemctl -M hud@.host --user status"