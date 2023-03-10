#!/bin/bash
set -euo pipefail

function query_blob() {
	JSON_BLOB="$1"
    JSON_QUERY="$2"

    echo "${JSON_BLOB}" | jq -Mcr "${JSON_QUERY}"
}

function query_file() {
   	JSON_FILE="$1"
    JSON_QUERY="$2" 
    
    jq -Mcr "${JSON_QUERY}" "${JSON_FILE}"
}

function query_file_array() {
    JSON_FILE="$1"
    INTO_VARIABLE="$2"
    JSON_QUERY="$3"

    mapfile -t "${INTO_VARIABLE}" < <(query_file "${JSON_FILE}" "${JSON_QUERY}")
}

function json_set_difference() {
    JSON_ARRAY_A="${1}"
    JSON_ARRAY_B="${2}"

    jq -nMcr \
       --argjson setA "${JSON_ARRAY_A}" \
       --argjson setB "${JSON_ARRAY_B}" \
       '$setA - $setB'
}