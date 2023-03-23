#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"
ON_GLASSES="$(echo "${STEP_CONFIG}" | jq -r '.on_glasses // false')"
SVC_USER="$(echo "${STEP_CONFIG}" | jq -r '.user // "hud"')"


cat <<EOF | envsubst >"${OUT}"
[Unit]
Description=${APP_NAME}, a HUD Application
AssertPathExists=/opt/hud/run/wayland-0
EOF

if [[ "${ON_GLASSES}" == "true" ]]; then
cat <<EOF | envsubst >>"${OUT}"
PartOf=hud-glasses.target
EOF
else
cat <<EOF | envsubst >>"${OUT}"
PartOf=hud-apps.target
EOF
fi


cat <<EOF | envsubst >>"${OUT}"

[Service]
User=${SVC_USER}
Group=${SVC_USER}
WorkingDirectory=/opt/hud

Slice=hud.slice
Type=simple
Environment=XDG_RUNTIME_DIR=/opt/hud/run
ExecStart=/usr/bin/${APP_NAME}
Restart=always

EOF

if [[ "${ON_GLASSES}" == "true" ]]; then
cat <<EOF | envsubst >>"${OUT}"
[Install]
WantedBy=hud-glasses.target
EOF
else
cat <<EOF | envsubst >>"${OUT}"
[Install]
WantedBy=hud-apps.target
EOF
fi

