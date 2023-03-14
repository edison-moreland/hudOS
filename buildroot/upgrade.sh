#!/usr/bin/env bash
set -euo pipefail

BUILDROOT_DIR="$(dirname $(realpath $0))"
REPO_ROOT="${BUILDROOT_DIR}/.."
source "${REPO_ROOT}"/logging.sh

PINEPHONE_HOST=$1
PINEPHONE_DEPLOY_SSH="deploy@${PINEPHONE_HOST}"
PINEPHONE_SSH_KEY="${BUILDROOT_DIR}/deploy_ed25519"

BUILDROOT_ROOT_IMAGE="${BUILDROOT_DIR}/.buildroot/output/images/rootfs.ext4"

if [ ! -f "${BUILDROOT_ROOT_IMAGE}" ]; then
    log_fatal "No root image. Did you run buildroot?"
fi

gzip -c "${BUILDROOT_ROOT_IMAGE}" | pv | ssh \
	-o StrictHostKeychecking=no \
	-o UserKnownHostsFile=/dev/null \
	-i "${PINEPHONE_SSH_KEY}" \
	"${PINEPHONE_DEPLOY_SSH}" \
	"sudo hudctl root upgrade"