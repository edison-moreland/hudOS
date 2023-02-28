#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
BINARIES=($(echo "${STEP_CONFIG}" | jq -r '.binaries // [] | .[]'))
UNITS=($(echo "${STEP_CONFIG}" | jq -r '.units // [] | .[]'))
CONFIGS=($(echo "${STEP_CONFIG}" | jq -r '.configs // [] | .[]'))
SCRIPTS=($(echo "${STEP_CONFIG}" | jq -r '.scripts // [] | .[]'))


STAGING_BUNDLE="$(mktemp)"
function remove_staging_bundle {
	rm "${STAGING_BUNDLE}"
}
trap remove_staging_bundle EXIT

function add_to_bundle {
	DIR_IN_BUNDLE="${1}"
	FILE="${2}"

	path_transform="s,^,${DIR_IN_BUNDLE}/,"

	tar -rf "${STAGING_BUNDLE}" -C "$(dirname "${FILE}")" --transform "${path_transform}" "$(basename "${FILE}")"
}

for binary in "${BINARIES[@]}"; do
	add_to_bundle 'bin' "${binary}"
done

for unit in "${UNITS[@]}"; do
	add_to_bundle 'services' "${unit}"
done

for config in "${CONFIGS[@]}"; do
	add_to_bundle 'configs' "${config}"
done

for script in "${SCRIPTS[@]}"; do
	add_to_bundle 'post_deploy' "${script}"
done

cp "${STAGING_BUNDLE}" "${OUTPUT_BUNDLE}"