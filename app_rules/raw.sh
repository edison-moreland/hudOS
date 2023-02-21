#!/usr/bin/env bash
set -euo pipefail

RULES_DIR="$(dirname $(realpath $0))"
REPO_ROOT="${RULES_DIR}/.."
source "${REPO_ROOT}"/logging.sh

APP_NAME=$1
APP_DIR=$2
RULE_CONFIG=$(echo "$3" | base64 -d)
RULE_WORKSPACE=$4
RULE_BUNDLE_OUT=$5

STAGING_BUNDLE="$(mktemp)"
function remove_staging_bundle {
	rm "${STAGING_BUNDLE}"
}
trap remove_staging_bundle EXIT

log "Bundling $APP_NAME"
# todo(edison) fix shell check warning below
binaries=($(echo "${RULE_CONFIG}" | jq -r '.binaries // [] | .[]'))
units=($(echo "${RULE_CONFIG}" | jq -r '.units // [] | .[]'))
configs=($(echo "${RULE_CONFIG}" | jq -r '.configs // [] | .[]'))
scripts=($(echo "${RULE_CONFIG}" | jq -r '.scripts // [] | .[]'))

function add_to_bundle {
	DIR_IN_BUNDLE="${1}"
	FILE="${2}"

	path_transform="s,^,${DIR_IN_BUNDLE}/,"

	tar -rf "${STAGING_BUNDLE}" -C "$(dirname "${FILE}")" --transform "${path_transform}" "$(basename "${FILE}")"
}

for binary in "${binaries[@]}"; do
	binary="${APP_DIR}/${binary}"
	add_to_bundle 'bin' "${binary}"
done

for unit in "${units[@]}"; do
	unit="${APP_DIR}/${unit}"
	add_to_bundle 'services' "${unit}"
done

for config in "${configs[@]}"; do
	config="${APP_DIR}/${config}"
	add_to_bundle 'configs' "${config}"
done

for script in "${scripts[@]}"; do
	script="${APP_DIR}/${script}"
	add_to_bundle 'post_deploy' "${script}"
done

cp "${STAGING_BUNDLE}" "${RULE_BUNDLE_OUT}"
