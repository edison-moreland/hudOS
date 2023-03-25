#!/usr/bin/env bash
set -euo pipefail
source "${HB_LIB_DIR}/logging.sh"

function zip_should_update() {
    dep_dir=$1
    dep_version=$2

    zip_path="${dep_dir}.zip"

    existing_sha="$(shasum -a 256 "${zip_path}" | cut -d' ' -f1)"
    if [[ "${existing_sha}" == "${dep_version}" ]]; then
        return 1
    else
        return 0
	fi 
}

function zip_update() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3

    zip_path="${dep_dir}.zip"

    rm -f "${zip_path}"
    rm -rf "${dep_dir}"
    zip_download "${dep_dir}" "${dep_version}" "${dep_url}" 
}

function zip_download() {
    dep_dir=$1
    dep_version=$2
    dep_url=$3

    zip_path="${dep_dir}.zip"
    
    wget --show-progress --no-verbose -O "${zip_path}" "${dep_url}"

    downloaded_sha="$(shasum -a 256 "${zip_path}" | cut -d' ' -f1)"
    if [[ "${downloaded_sha}" != "${dep_version}" ]]; then
        log_failure "SHA Mismatch!"
        log_info "Expected: ${dep_version}"
        log_info "     Got: ${downloaded_sha}"
        return 1
	fi

    mkdir -p "${dep_dir}"

    unzip "${zip_path}" -d "${dep_dir}"
}