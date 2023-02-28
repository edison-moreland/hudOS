#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

if ! which jq >/dev/null; then
	log_red "Please install jq!"
	exit 1

fi

EXCLUDE_APPS=""
INCLUDE_APPS=""
while getopts i:e: opt; do
	case $opt in
	e)
		EXCLUDE_APPS=$OPTARG
		;;

	i)
		INCLUDE_APPS=$OPTARG
		;;

	\?) ;;
	:) ;;
	esac
done
shift $((OPTIND - 1))

APP_BUILDER="${REPO_ROOT}/app_builder/app_builder.sh"
APPS_DIR="${REPO_ROOT}/apps"
BUILD_OUTPUT="${REPO_ROOT}/.build"
FINAL_BUNDLE="${BUILD_OUTPUT}/bundle.tar"

log_blue "Updating vendor"
"${REPO_ROOT}"/update_vendor.sh

log_blue "Building apps"
for app_manifest in "${APPS_DIR}"/**/.hud_app.json; do
	app_name="$(jq -r '.app.name' "${app_manifest}")"
	bundle_out="${BUILD_OUTPUT}/app/${app_name}/bundle.tar"

	if [[ "${INCLUDE_APPS}" != "" ]]; then
		if ! [[ "${INCLUDE_APPS}" =~ ${app_name} ]]; then
			log_yellow "Not Included $app_name"
			continue
		fi
	else
		if [[ "${EXCLUDE_APPS}" =~ ${app_name} ]]; then
			log_yellow "Excluding $app_name"
			continue
		fi
	fi

	if "${APP_BUILDER}" -i "${app_manifest}" -o "${bundle_out}"; then
		if [ ! -f "${bundle_out}" ]; then
			log_yellow "Build did NOT produce an app bundle! (${app_name})"
		else
			log_green "Success! (${app_name})"
		fi
	else
		log_red "Failure! (${app_name})"
		exit 1
	fi
done

log_blue "Building final bundle"
TEMP_BUNDLE="$(mktemp)"
function remove_temp_bundle {
	rm "${TEMP_BUNDLE}"
}
trap remove_temp_bundle EXIT

for app_bundle in "${BUILD_OUTPUT}"/app/**/bundle.tar; do
	tar --concatenate --file="${TEMP_BUNDLE}" "${app_bundle}"
done

cp "${TEMP_BUNDLE}" "${FINAL_BUNDLE}"
