#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
SETCAP="$(echo "${STEP_CONFIG}" | jq -r '.setcap // ""')"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"

cat <<EOF > "${OUT}"
#!/bin/bash

EOF

if [[ "${SETCAP}" != "" ]]; then
cat <<EOF | envsubst >> "${OUT}"
setcap '${SETCAP}' /opt/hud/apps/${APP_NAME}/binaries/${APP_NAME}
EOF
fi