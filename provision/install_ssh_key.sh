#!/bin/bash
set -euo pipefail

PUBLIC_KEY_B64="${1}"

authorizedkeys_file=$(mktemp)
echo "${PUBLIC_KEY_B64}" | base64 -d > "${authorizedkeys_file}"
chmod 644 "${authorizedkeys_file}"

if [ ! -d "$HOME/.ssh" ]
then
  mkdir -m 700 "$HOME/.ssh"
fi
cp --preserve=mode "${authorizedkeys_file}" "$HOME/.ssh/authorized_keys"

echo '123456' | sudo -S mv "${authorizedkeys_file}" "/root/.ssh/authorized_keys"
echo '123456' | sudo -S chown root:root "/root/.ssh/authorized_keys"