#!/usr/bin/env bash
set -euo pipefail

BUILDROOT_DIR="$(dirname $(realpath $0))"
REPO_ROOT="${BUILDROOT_DIR}/.."
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_DEPLOY_SSH="deploy@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${REPO_ROOT}/buildroot/deploy_ed25519"

if [ ! -f "${PINEPHONE_SSH_KEY}" ]; then
	log_red "Deploy key not generated. Did you enable deploy user in buildroot?"
	exit 1
fi

ssh \
    -o StrictHostKeychecking=no \
    -o UserKnownHostsFile=/dev/null \
    -i "${PINEPHONE_SSH_KEY}" \
    "${PINEPHONE_DEPLOY_SSH}"
