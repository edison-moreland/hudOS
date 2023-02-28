#!/usr/bin/env bash
set -euo pipefail

APPBUILDER_DIR="$(dirname "$(realpath "$0")")"
REPO_DIR="$(realpath "${APPBUILDER_DIR}/..")"
source "${REPO_DIR}"/logging.sh

INPUT_MANIFEST=""
OUTPUT_BUNDLE=""
while getopts i:o: opt; do
	case $opt in
	i)
		INPUT_MANIFEST=$OPTARG
		;;

	o)
		OUTPUT_BUNDLE=$OPTARG
		;;

	\?) ;; #TODO
	:) ;;
	esac
done
shift $((OPTIND - 1))

if [ ! -f "${INPUT_MANIFEST}" ]; then
    log_red "Manifest ${INPUT_MANIFEST} does not exist"
    exit 1
fi

STEP_DIR="${APPBUILDER_DIR}/steps"

export APP_NAME="$(jq -r '.app.name' "${INPUT_MANIFEST}")"
export BUILD_DIR="${REPO_DIR}/.build"
export CACHE_DIR="${BUILD_DIR}/cache"
export WORKSPACE_DIR="${BUILD_DIR}/app/${APP_NAME}/workspace"
export SOURCE_DIR="$(dirname "$(realpath "${INPUT_MANIFEST}")")"
export OUTPUT_BUNDLE

# Workspace gets cleaned every time
if [ -d "${WORKSPACE_DIR}" ]; then
    rm -rf "${WORKSPACE_DIR}"
fi
mkdir -p "${WORKSPACE_DIR}"
mkdir -p "${CACHE_DIR}"

# Given the path to a .hud_app.json, generate a bundle.tar
log_blue "Building ${APP_NAME}"

BUILD_STEPS=($(jq -cM '.build[]' "${INPUT_MANIFEST}"))
for step in "${BUILD_STEPS[@]}"; do
    step_subst="$(echo "${step}" | envsubst)"
    step_type="$(echo "${step_subst}" | jq -r '.step')"

    if [ ! -f "${STEP_DIR}/${step_type}.sh" ]; then
        log_red "Step ${step_type} is not implemented"
        exit 1
    fi

    log_blue " - ${step_type}"
    "${STEP_DIR}/${step_type}.sh" "${step_subst}"
done