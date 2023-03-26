#!/usr/bin/env bash
set -euo pipefail

APPBUILDER_DIR="$(dirname "$(realpath "$0")")"
REPO_DIR="$(realpath "${APPBUILDER_DIR}/..")"
source "${REPO_DIR}"/hud_builder/lib/logging.sh

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
    log_fatal "Manifest ${INPUT_MANIFEST} does not exist"
fi

STEP_DIR="${APPBUILDER_DIR}/steps"

export APP_NAME="$(jq -r '.app.name' "${INPUT_MANIFEST}")"
export APP_METADATA="$(jq -Mcr '.app' "${INPUT_MANIFEST}")"
export BUILD_DIR="${REPO_DIR}/.build"
export CACHE_DIR="${BUILD_DIR}/cache/apps"
export WORKSPACE_DIR="${BUILD_DIR}/app/${APP_NAME}/workspace"
export VENDOR_DIR="${BUILD_DIR}/vendor"
export PROTO_DIR="${BUILD_DIR}/proto"
export SOURCE_DIR="$(dirname "$(realpath "${INPUT_MANIFEST}")")"
export BUILDROOT_OUTPUT_DIR="${VENDOR_DIR}/buildroot/output"
export BUILDROOT_HOST_DIR="${BUILDROOT_OUTPUT_DIR}/host"
export BUILDROOT_SYSROOT_DIR="${BUILDROOT_HOST_DIR}/aarch64-buildroot-linux-gnu/sysroot"
export OUTPUT_BUNDLE

if [ ! -d "${BUILDROOT_OUTPUT_DIR}" ]; then
	log_fatal "Please run buildroot/build.sh!"
fi

# Workspace gets cleaned every time
if [ -d "${WORKSPACE_DIR}" ]; then
    rm -rf "${WORKSPACE_DIR}"
fi
mkdir -p "${WORKSPACE_DIR}"
mkdir -p "${CACHE_DIR}"

# Given the path to a .hud_app.json, generate a bundle.tar
log_section "Building ${APP_NAME}"

BUILD_STEPS=($(jq -cM '.build[]' "${INPUT_MANIFEST}"))
for step in "${BUILD_STEPS[@]}"; do
    step_subst="$(echo "${step}" | envsubst)"
    step_type="$(echo "${step_subst}" | jq -r '.step')"

    if [ "${DEBUG_STEP_CONFIG:-0}" == "1" ]; then
        echo "${step_subst}" | jq '.'
    else
        log_subsection "${step_type}"
    fi
    
    if [ -f "${SOURCE_DIR}/.build_steps/${step_type}.sh" ]; then
        step_file="${SOURCE_DIR}/.build_steps/${step_type}.sh"
    elif [ -f "${STEP_DIR}/${step_type}.sh" ]; then
        step_file="${STEP_DIR}/${step_type}.sh"
    else
        log_fatal "Step ${step_type} is not implemented"
    fi

    "${step_file}" "${step_subst}"
done