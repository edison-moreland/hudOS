#!/bin/bash
set -euo pipefail

pacman -Sy --needed --noconfirm weston seatd rsync
systemctl enable seatd
systemctl start seatd