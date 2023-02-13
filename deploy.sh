#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_ROOT_SSH="root@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${REPO_ROOT}/.deploy/id_ed25519"

APP_BUNDLE="${REPO_ROOT}/bazel-bin/app_bundle.tar"
BUNDLE_DEPLOYER="${REPO_ROOT}/bundle_deployer.sh"

log_blue "Building App Bundle"
bazel build '//:app_bundle'

log_blue "Deploying App Bundle"
scp -i "${PINEPHONE_SSH_KEY}" "${BUNDLE_DEPLOYER}" "${PINEPHONE_ROOT_SSH}":/root/bundle_deployer.sh
pv < "${APP_BUNDLE}" | ssh -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "/bin/bash /root/bundle_deployer.sh"

log_blue "Deployment Status"
ssh -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "systemctl -M hud@.host --user status"