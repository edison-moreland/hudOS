#!/bin/bash
set -euo pipefail

ANSI_RED="\033[0;31m"
ANSI_GREEN="\033[0;32m"
ANSI_YELLOW="\033[0;33m"
ANSI_BLUE="\033[0;94m"
ANSI_BOLD="\033[1m"
ANSI_ITALIC="\033[3m"
ANSI_RESET="\033[0m"

function log() {
	printf "$@\n" >&2
}

function log_blue() {
	log "${ANSI_BLUE}${ANSI_ITALIC}$@${ANSI_RESET}"
}

function log_red() {
	log "${ANSI_RED}${ANSI_BOLD}$@${ANSI_RESET}"
}

function log_green() {
	log "${ANSI_GREEN}$@${ANSI_RESET}"
}

function log_yellow() {
	log "${ANSI_YELLOW}$@${ANSI_RESET}"
}
