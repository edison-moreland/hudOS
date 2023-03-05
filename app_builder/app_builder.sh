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
export VENDOR_DIR="${BUILD_DIR}/vendor"
export SOURCE_DIR="$(dirname "$(realpath "${INPUT_MANIFEST}")")"
export BUILDROOT_OUTPUT_DIR="${REPO_DIR}/buildroot/.buildroot/output"
export BUILDROOT_HOST_DIR="${BUILDROOT_OUTPUT_DIR}/host"
export BUILDROOT_SYSROOT_DIR="${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot"
export OUTPUT_BUNDLE

if [ ! -d "${BUILDROOT_OUTPUT_DIR}" ]; then
	log_red "Please run buildroot/build.sh!"
	exit 1
fi

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

    if [ "${DEBUG_STEP_CONFIG:-0}" == "1" ]; then
        echo "${step_subst}" | jq '.'
    else
        log_blue " - ${step_type}"
    fi

    if [ ! -f "${STEP_DIR}/${step_type}.sh" ]; then
        log_red "Step ${step_type} is not implemented"
        exit 1
    fi

    "${STEP_DIR}/${step_type}.sh" "${step_subst}"
done