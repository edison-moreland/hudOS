#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_ROOT_SSH="deploy@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${REPO_ROOT}/buildroot/deploy_ed25519"

TARGET="${2:-}"

APP_BUNDLE="${REPO_ROOT}/.build/bundle.tar.gz"

if [ ! -f "${PINEPHONE_SSH_KEY}" ]; then
	log_red "Deploy key not generated. Did you enable deploy user in buildroot?"
	exit 1
fi

log_blue "Building App Bundle"
buildargs='-e bluetooth'
if [ "${TARGET}" != "" ]; then
	buildargs+=" -i ${TARGET}"
fi

"${REPO_ROOT}"/build.sh ${buildargs}

log_blue "Deploying App Bundle"
pv <"${APP_BUNDLE}" | ssh \
	-o StrictHostKeychecking=no \
	-o UserKnownHostsFile=/dev/null \
	-i "${PINEPHONE_SSH_KEY}" \
	"${PINEPHONE_ROOT_SSH}" \
	"sudo hudctl deploy"

# log_blue "Deployment Status"
#ssh -i "${PINEPHONE_SSH_KEY}" "${PINEPHONE_ROOT_SSH}" "systemctl -M hud@.host --user status"
