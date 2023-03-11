#!/bin/bash
set -euo pipefail

function json_query_blob() {
	JSON_BLOB="$1"
    JSON_QUERY="$2"

    echo "${JSON_BLOB}" | jq -Mcr "${JSON_QUERY}"
}

function json_query_file() {
   	JSON_FILE="$1"
    JSON_QUERY="$2" 
    
    jq -Mcr "${JSON_QUERY}" "${JSON_FILE}"
}

function json_query_file_array() {
    JSON_FILE="$1"
    INTO_VARIABLE="$2"
    JSON_QUERY="$3"

    mapfile -t "${INTO_VARIABLE}" < <(json_query_file "${JSON_FILE}" "${JSON_QUERY}")
}

function json_array_difference() {
    JSON_ARRAY_A="${1}"
    JSON_ARRAY_B="${2}"

    jq -nMcr \
       --argjson setA "${JSON_ARRAY_A}" \
       --argjson setB "${JSON_ARRAY_B}" \
       '$setA - $setB'
}

function json_array_into() {
    # Read a json array into a bash array
    INTO_VARIABLE="$1"
    JSON_ARRAY="$2"

    mapfile -t "${INTO_VARIABLE}" < <(json_query_blob "${JSON_ARRAY}" '. // [] | .[]')
}