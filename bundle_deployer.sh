#!/bin/bash
set -euo pipefail

BUNDLE=$(mktemp -d)
tar -xzf - -C "${BUNDLE}" <&0
function clean_bundle {
	rm -r "${BUNDLE}"
}
trap clean_bundle EXIT

# Temporary until app link locations get changed
# export PATH="/opt/hud/bin:${PATH}"

if [ -f "${BUNDLE}/weston.tar" ]; then
	# Special case, weston always gets installed first
	# It would be nice if apps could have dependecies on each other
	hudctl install "${BUNDLE}/weston.tar"
	rm "${BUNDLE}/weston.tar"
fi

for app_bundle in "${BUNDLE}"/*.tar; do
	hudctl install "${app_bundle}"
done

systemctl restart hud.target 
systemctl restart hud-apps.target