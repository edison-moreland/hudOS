#!/usr/bin/env bash
set -euo pipefail

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