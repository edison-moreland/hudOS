#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(dirname $(realpath $0))"
source "${REPO_ROOT}"/logging.sh

function tar_archive_path() {
    dep_dir=$1
    dep_url=$2

    # For tar dependecies, the archive is stored next to the extracted folder
    # The .tar.* extension is kept on the file so the gnu tar can guess the compression method
    ext="$(basename "${dep_url}" | grep -oE '\.tar\.[^./]+$')"
    printf '%s%s' "${dep_dir}" "${ext}"
}

function tar_should_update() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3
    
    tar_path="$(tar_archive_path "${dep_dir}" "${dep_url}")"

    existing_sha="$(shasum -a 256 "${tar_path}" | cut -d' ' -f1)"
    if [[ "${existing_sha}" == "${dep_version}" ]]; then
        return 1
    else
        return 0
	fi 
}

function tar_update() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3

    tar_path="$(tar_archive_path "${dep_dir}" "${dep_url}")"

    rm -f "${tar_path}"
    rm -rf "${dep_dir}"
    tar_download "${dep_dir}" "${dep_version}" "${dep_url}" 
}

function tar_download() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3

    tar_path="$(tar_archive_path "${dep_dir}" "${dep_url}")"

    wget --show-progress --no-verbose -O "${tar_path}" "${dep_url}" 
	
    downloaded_sha="$(shasum -a 256 "${tar_path}" | cut -d' ' -f1)"
    if [[ "${downloaded_sha}" != "${dep_version}" ]]; then
        log_red "SHA Mismatch!"
        log_red "Expected: ${dep_version}"
        log_red "     Got: ${downloaded_sha}"
        return 1
	fi

    mkdir -p "${dep_dir}"
    tar --strip-components=1 -xf "${tar_path}" -C "${dep_dir}"
}

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
    "tar")
        tar_should_update "${@:2}"
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
    "tar")
        tar_update "${@:2}"
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
    "tar")
        tar_download "${@:2}"
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
    dep_patches=($(echo "${dep}" | jq -r '.patches // [] | .[]'))

    dep_dir="${VENDOR_DIR}/${dep_name}"

    action_needed=""
    if [ -d "${dep_dir}" ]; then
        if should_update "${dep_type}" "${dep_dir}" "${dep_version}" "${dep_url}"; then
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
        update "${dep_type}" "${dep_dir}" "${dep_version}" "${dep_url}"

        if (( ${#dep_patches[@]} != 0 )); then
            log_yellow "Updating a dependency with patches may not work properly!"
        fi

        ;;
    "download")
        log_blue "Downloading ${dep_name} (${dep_type})"
        if [ ! -d "$(dirname "${dep_dir}")" ]; then
            mkdir -p "$(dirname "${dep_dir}")"
        fi

        download "${dep_type}" "${dep_dir}" "${dep_version}" "${dep_url}"

        if (( ${#dep_patches[@]} != 0 )); then
            log_blue "Applying patches..."
            for dep_patch in "${dep_patches[@]}"; do
                log_blue "- $(basename "${dep_patch}")"
                patch -p1 -d "${dep_dir}" < "${REPO_ROOT}/${dep_patch}"
            done
        fi
        ;;
    "nothing")
        log_yellow "Up to date! ${dep_name} (${dep_type})"
        ;;
    esac
done
