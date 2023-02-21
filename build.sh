#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

if ! which jq >/dev/null; then
	log_red "Please install jq!"
	exit 1

fi

EXCLUDE_APPS=""
while getopts "e:" opt; do
	case $opt in
	e)
		EXCLUDE_APPS=$OPTARG
		;;
	\?) ;;
	:) ;;
	esac
done
shift $((OPTIND - 1))

APPS_DIR="${REPO_ROOT}/apps"
RULES_DIR="${REPO_ROOT}/app_rules"
BUILD_OUTPUT="${REPO_ROOT}/.build"
APP_BUNDLE="${BUILD_OUTPUT}/bundle.tar"

if [ -d "${BUILD_OUTPUT}" ]; then
	rm -r "${BUILD_OUTPUT}"
fi
mkdir "${BUILD_OUTPUT}"

for app_manifest in "${APPS_DIR}"/**/.hud_app.json; do
	app_dir="$(dirname "${app_manifest}")"
	app_name="$(jq -r '.name' "${app_manifest}")"
	app_type="$(jq -r '.type' "${app_manifest}")"
	build_rule="${RULES_DIR}/${app_type}.sh"
	rule_config="$(jq -c --arg type "${app_type}" '.[$type]' "${app_manifest}" | base64 -w 0)"
	# This is a workspace that the build rule can use however it wants
	rule_workspace="${BUILD_OUTPUT}/app/${app_name}/build"
	# This is where the build rule is expected to put the final app bundle
	rule_bundle_out="${BUILD_OUTPUT}/app/${app_name}/bundle.tar"

	if [[ "${EXCLUDE_APPS}" =~ ${app_name} ]]; then
		log_yellow "Excluding $app_name ($app_type)"
		continue
	fi
	log_blue "Building $app_name ($app_type)"

	if [ ! -f "${build_rule}" ]; then
		log_red "No implementation for $app_type rule"
		exit 1
	fi

	mkdir -p "${rule_workspace}"

	if "${build_rule}" "${app_name}" "${app_dir}" "${rule_config}" "${rule_workspace}" "${rule_bundle_out}"; then
		if [ ! -f "${rule_bundle_out}" ]; then
			log_yellow "Rule did NOT produce an app bundle!"
		else
			log_green "Success!"
		fi
	else
		log_red "Failure :("
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

cp "${TEMP_BUNDLE}" "${APP_BUNDLE}"
