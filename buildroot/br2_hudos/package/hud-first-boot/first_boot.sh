#!/bin/bash
set -euo pipefail

echo "RUNNING FIRST BOOT HOOKS!"

for first_boot_script in /opt/hud/firstboot/*.sh; do
    "${first_boot_script}"
done

rm /usr/sbin/first_boot.sh
rm /usr/lib/systemd/system/first_boot.service
rm /etc/systemd/system-preset/50-first_boot.preset