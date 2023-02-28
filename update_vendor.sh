#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

function git_in() {
    # Tell git to act on a different directory
    # git_in <dir> <git args>
    git --git-dir="${1}"/.git --work-tree="${1}" "${@:2}" 
}

function git_should_update() {
    dep_dir=$1
    dep_version=$2

    if [ "$(git_in "${dep_dir}" rev-parse HEAD)" = "${dep_version}" ]; then
        return 1
    else
        return 0
    fi
}

function git_update() {
    dep_dir=$1
    dep_version=$2
    git_in "${dep_dir}" checkout -q "${dep_version}"
}

function git_download() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3

    git clone -q -- "${dep_url}" "${dep_dir}"
    git_in "${dep_dir}" checkout -q "${dep_version}"
}

function should_update() {
    case "${1}" in
    "git")
        git_should_update "${@:2}"
        ;;
    *)
        log_red "Vendor type ${1} is not supported"
        exit 1
        ;;
    esac
}

function update() {
    case "${1}" in
    "git")
        git_update "${@:2}"
        ;;
    *)
        log_red "Vendor type ${1} is not supported"
        exit 1
        ;;
    esac
}

function download() {
    case "${1}" in
    "git")
        git_download "${@:2}"
        ;;
    *)
        log_red "Vendor type ${1} is not supported"
        exit 1
        ;;
    esac
}

VENDOR_DIR="${REPO_ROOT}/.build/vendor"
VENDOR_FILE="${REPO_ROOT}/vendor.json"
VENDOR=($(cat "${VENDOR_FILE}" | jq -cM '.[]'))

for dep in "${VENDOR[@]}"; do
    dep_name="$(echo "${dep}" | jq -r '.name')"
    dep_type="$(echo "${dep}" | jq -r '.type')"
    dep_url="$(echo "${dep}" | jq -r '.url')"
    dep_version="$(echo "${dep}" | jq -r '.version')"

    dep_dir="${VENDOR_DIR}/${dep_name}"

    action_needed=""
    if [ -d "${dep_dir}" ]; then
        if should_update "${dep_type}" "${dep_dir}" "${dep_version}"; then
            action_needed="update"
        else
            action_needed="nothing"
        fi
    else
        action_needed="download"
    fi

    case "${action_needed}" in
    "update")
        log_blue "Updating ${dep_name} (${dep_type})"
        update "${dep_type}" "${dep_dir}" "${dep_version}"
        ;;
    "download")
        log_blue "Downloading ${dep_name} (${dep_type})"
        if [ ! -d "$(dirname "${dep_dir}")" ]; then
            mkdir -p "$(dirname "${dep_dir}")"
        fi

        download "${dep_type}" "${dep_dir}" "${dep_version}" "${dep_url}"
        ;;
    "nothing")
        log_yellow "Up to date! ${dep_name} (${dep_type})"
        ;;
    esac
done
