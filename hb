#!/usr/bin/env bash
set -euo pipefail

export HB_REPOSITORY_DIR="$(dirname "$(realpath "$0")")"
export HB_COMMANDS_DIR="${HB_REPOSITORY_DIR}/hud_builder"
export HB_LIB_DIR="${HB_COMMANDS_DIR}/lib"
export HB_APP_BUILDER="${HB_REPOSITORY_DIR}/app_builder/app_builder.sh"
export HB_APPS_DIR="${HB_REPOSITORY_DIR}/apps"
export HB_BUILD_DIR="${HB_REPOSITORY_DIR}/.build"
export HB_FINAL_BUNDLE="${HB_BUILD_DIR}/bundle.tar"
export HB_BUILD_CACHE_DIR="${HB_BUILD_DIR}/cache"
export HB_BUILD_APP_DIR="${HB_BUILD_DIR}/app"
export HB_VENDOR_FILE="${HB_REPOSITORY_DIR}/vendor.json"
export HB_VENDOR_DIR="${HB_BUILD_DIR}/vendor"
export HB_OUTPUT_DIR="${HB_REPOSITORY_DIR}/output"

export PATH="${HB_COMMANDS_DIR}:${PATH}"

source "${HB_LIB_DIR}/logging.sh"

if ! which jq >/dev/null; then
	log_fatal "Please install jq!"
fi

cmd="${1:-help}"
hb_command="hb-$cmd"
if ! which "${hb_command}" > /dev/null; then
    log_fatal "Unknown command $cmd"
fi

if [[ "${HB_NO_LOCK:-}" == "buildroot" ]]; then
    # This is only to be used with buildroot hooks
    # Don't do something stupid
    "${hb_command}" "${@:2}"
    exit $?
fi

{ 
    if ! flock -n 9; then
        log_fatal "Could not obtain lock"
    fi

    "${hb_command}" "${@:2}"
} 9>"${HB_BUILD_DIR}/.hblock"