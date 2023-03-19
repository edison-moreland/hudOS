#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
OUT="$(echo "${STEP_CONFIG}" | jq -r '.out')"

cat <<EOF | envsubst >"${OUT}"
[Unit]
Description=${APP_NAME}, a HUD Application
AssertPathExists=/opt/hud/run/wayland-0
PartOf=hud-apps.target

[Service]
User=hud
Group=hud
WorkingDirectory=/opt/hud

Slice=hud.slice
Type=simple
Environment=XDG_RUNTIME_DIR=/opt/hud/run
ExecStart=/usr/bin/${APP_NAME}
Restart=always

[Install]
WantedBy=hud-apps.target
EOF