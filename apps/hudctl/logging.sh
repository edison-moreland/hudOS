#!/bin/bash
set -euo pipefail

ANSI_RED='\033[0;31m'
ANSI_GREEN='\033[0;32m'
ANSI_YELLOW='\033[0;33m'
ANSI_BLUE='\033[0;34m'
ANSI_PURPLE='\033[0;35m'
ANSI_CYAN='\033[0;36m'

# ANSI_RED="\033[0;31m"
# ANSI_GREEN="\033[0;32m"
# ANSI_YELLOW="\033[0;33m"
# ANSI_BLUE="\033[0;94m"
ANSI_BOLD="\033[1m"
ANSI_ITALIC="\033[3m"
ANSI_RESET="\033[0m"

function log() {
	printf "$@\n" >&2
}

function log_section() {
	log "${ANSI_BLUE}$@${ANSI_RESET}"	
}

function log_subsection() {
	log "${ANSI_CYAN}- ${ANSI_ITALIC}$@${ANSI_RESET}"	
}

function log_action() {
	printf "${ANSI_PURPLE}${ANSI_ITALIC}%10s %s ${ANSI_RESET}\n" "${1:0:10}" "${@:2}" >&2
}

function log_warning() {
	log "${ANSI_YELLOW}${ANSI_BOLD}$@${ANSI_RESET}"
}

function log_success() {
	log "${ANSI_GREEN}SUCCESS! $@${ANSI_RESET}"
}

function log_failure() {
	log "${ANSI_RED}FAILURE! $@${ANSI_RESET}"
}

function log_fatal() {
	log_failure "$@"
	exit 1
}

function log_info() {
	log "$@"
}