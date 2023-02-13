#!/bin/bash
set -euo pipefail

HUD_USER="hud"
HUD_HOME="/opt/hud"

if ! $(id ${HUD_USER} > /dev/null 2>&1)
then
  useradd -U -d "${HUD_HOME}" -r "${HUD_USER}"
fi

mkdir -p "/etc/systemd/system/getty@tty1.service.d/"
cat << EOF > "/etc/systemd/system/getty@tty1.service.d/autologin.conf"
[Service]
ExecStart=
ExecStart=-/sbin/agetty -o '-p -f -- \\u' --noclear --autologin ${HUD_USER} %I \$TERM
Environment=XDG_SESSION_TYPE=wayland
EOF
