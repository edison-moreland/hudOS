#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
BINARIES=($(echo "${STEP_CONFIG}" | jq -r '.binaries // [] | .[]'))
UNITS=($(echo "${STEP_CONFIG}" | jq -r '.units // [] | .[]'))
CONFIGS=($(echo "${STEP_CONFIG}" | jq -r '.configs // [] | .[]'))
SCRIPTS=($(echo "${STEP_CONFIG}" | jq -r '.scripts // [] | .[]'))
DATAS=($(echo "${STEP_CONFIG}" | jq -r '.data // [] | .[]'))

STAGING_BUNDLE="$(mktemp)"
function remove_staging_bundle {
	rm "${STAGING_BUNDLE}"
}
trap remove_staging_bundle EXIT

MANIFEST="$(jq -ncM --arg app_name "${APP_NAME}" '{"name":$app_name}')"
function add_to_bundle {
	DIR_IN_BUNDLE="${1}"
	MODE="${2}"
	FILE="${3}"

	if [ ! -f "${FILE}" ]; then
		echo "File does not exist: ${FILE}"
		exit 1
	fi
	echo "   - ${DIR_IN_BUNDLE}/$(basename "${FILE}")"

	MANIFEST="$(
		echo "${MANIFEST}" | \
		jq -Mc \
		--arg type "${DIR_IN_BUNDLE}" \
		--arg file "$(basename "${FILE}")" \
		'.[$type] = (.[$type] + [$file])' \
	)"

	path_transform="s,^,${DIR_IN_BUNDLE}/,"

	tar -rf "${STAGING_BUNDLE}" \
	    -C "$(dirname "${FILE}")" \
		--mode "${MODE}" \
		--transform "${path_transform}" \
		"$(basename "${FILE}")"
}

for binary in "${BINARIES[@]}"; do
	add_to_bundle 'binaries' '755' "${binary}"
done

for unit in "${UNITS[@]}"; do
	add_to_bundle 'units' '644' "${unit}"
done

for config in "${CONFIGS[@]}"; do
	add_to_bundle 'configs' '640' "${config}"
done

for script in "${SCRIPTS[@]}"; do
	add_to_bundle 'scripts' '755' "${script}"
done

for data in "${DATAS[@]}"; do
	add_to_bundle 'data' '644' "${data}"
done

manifest_file="$(mktemp)"
echo "${MANIFEST}" | jq -M '.' > "${manifest_file}"
tar -rf "${STAGING_BUNDLE}" \
	-C "$(dirname "${manifest_file}")" \
	--mode "644" \
	--transform 's:.*:manifest.json:' \
	"$(basename "${manifest_file}")"

# tar -rf "${manifest_file}" \
# 	--mode "644" \
# 	--transform 's:.*:manifest.json:' \
# 	"${manifest_file}"
rm "${manifest_file}"

cp "${STAGING_BUNDLE}" "${OUTPUT_BUNDLE}"
