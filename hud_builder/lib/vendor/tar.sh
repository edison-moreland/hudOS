#!/usr/bin/env bash
set -euo pipefail
source "${HB_LIB_DIR}/logging.sh"

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
        log_failure "SHA Mismatch!"
        log_info "Expected: ${dep_version}"
        log_info "     Got: ${downloaded_sha}"
        return 1
	fi

    mkdir -p "${dep_dir}"
    tar --strip-components=1 -xf "${tar_path}" -C "${dep_dir}"
}