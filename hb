#!/usr/bin/env bash
set -euo pipefail

export HB_REPOSITORY_DIR="$(dirname "$(realpath "$0")")"
export HB_COMMANDS_DIR="${HB_REPOSITORY_DIR}/hud_builder"
export HB_LIB_DIR="${HB_COMMANDS_DIR}/lib"

export HB_BUILD_DIR="${HB_REPOSITORY_DIR}/.build"

export PATH="${HB_COMMANDS_DIR}:${PATH}"

source "${HB_LIB_DIR}/logging.sh"

cmd="${1:-help}"
hb_command="hb-$cmd"
if ! which "${hb_command}" > /dev/null; then
    log_fatal "Unknown command $cmd"
fi

{ 
    if ! flock -n 9; then
        log_fatal "Could not obtain lock"
    fi

    "${hb_command}" "${@:2}"
} 9>"${HB_BUILD_DIR}/.hblock"