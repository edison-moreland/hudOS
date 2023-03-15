#!/bin/bash
set -euo pipefail
source "${HB_LIB_DIR}/json.sh"

NAME_PARTS="${HB_LIB_DIR}/name-generator.json"

json_query_file_array "${NAME_PARTS}" NAME_EMOTIONS '.emotions[]'
json_query_file_array "${NAME_PARTS}" NAME_OBJECTS '.objects[]'

random_element() { 
    arr=("${!1}")
    printf '%s' ${arr[$((RANDOM % ${#arr[@]}))]}
}

random_name() {
    printf '%s-%s' \
           "$(random_element "NAME_EMOTIONS[@]")" \
           "$(random_element "NAME_OBJECTS[@]")"
}