#!/bin/bash
set -euo pipefail

BUNDLE=$(mktemp -d)
tar -xzf - -C "${BUNDLE}" <&0
function clean_bundle {
	rm -r "${BUNDLE}"
}
trap clean_bundle EXIT

# Temporary until app link locations get changed
export PATH="/opt/hud/bin:${PATH}"

for app_bundle in "${BUNDLE}"/*.tar; do
	hudctl install "${app_bundle}"
done


hudctl catalog