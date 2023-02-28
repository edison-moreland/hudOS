#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
ENABLE_UNITS=($(echo "${STEP_CONFIG}" | jq -cM '.enable_units[]'))
SETCAP="$(echo "${STEP_CONFIG}" | jq -r '.setcap // ""')"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"

cat <<EOF > "${OUT}"
#!/bin/bash

EOF

for unit in "${ENABLE_UNITS[@]}"; do
cat <<EOF | envsubst >> "${OUT}"
systemctl enable ${unit}
EOF
done

if [[ "${SETCAP}" != "" ]]; then
cat <<EOF | envsubst >> "${OUT}"
setcap '${SETCAP}' /opt/hud/bin/${APP_NAME}
EOF
fi