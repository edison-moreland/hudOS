#!/bin/bash
set -euo pipefail

# Parsing of arguments defined in json:
# {
#     "args": [
#         "arg-one",
#         "arg-two"
#     ],
#     "flags": {
#         "flag": {},
#         "short-flag": {"short": "s"},
#         "required-flag": {"short": "r", "required": true}
#     }
# }

# Generated use:
# <arg-one> <arg-two> [--flag <flag>] [(-s/--short-flag) <short-flag>] (-r/--required-flag) <required-flag>

# Generated variables:
# ARG_ARG_ONE="<arg_one>"
# ARG_ARG_TWO="<arg_one>"
# ARG_FLAG_PROVIDED="true"
# ARG_FLAG="<flag>"
# ARG_SHORT_FLAG_PROVIDED="true"
# ARG_SHORT_FLAG="<short-flag>"
# ARG_REQUIRED_FLAG="<required-flag>"

function parse_args() {
    # TODO
    # It would be nice to grab the json blob from the callers frontmatter
    :
}