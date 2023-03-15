#!/bin/bash
set -euo pipefail

# Parsing for the custom frontmatter format used for documentation
# Example:
# ```bash
# #!/bin/bash
# #-<key>: <value>
# ...
# ```

function get_frontmatter() {
    # Returns the frontmatter as a json blob
    FILE_PATH="$1"

    # This awk will give us every line that starts with #-, and it will split the fields
    # In:
    #   #!/usr/bin/env bash
    #   #-Help: Get/set ip address for a device
    #   #-Use: <device-name> [<ip-address>]
    # Out: ('-' is a placeholder for unit seperator '\x1f')
    #   help-Get/set ip address for a device
    #   use-<device-name> [<ip-address>]
    mapfile -t raw_frontmatter < <(awk -F': ' '/^#-/ {gsub(/^#-/, "", $1); print tolower($1) "\x1F" $2} ' < "${FILE_PATH}" )

    frontmatter="{}"
    for line in "${raw_frontmatter[@]}"; do
        key="$(echo "${line}" | cut -d$'\x1f' -f1)"
        val="$(echo "${line}" | cut -d$'\x1f' -f2)"

        frontmatter="$(
            echo "${frontmatter}" | \
            jq -M \
               --arg k "${key}" \
               --arg v "${val}" \
               '.[$k] = $v'
        )"
    done

    echo "${frontmatter}" | jq -M '.'
}