#!/bin/bash
set -euo pipefail

images_dir="$1"
target_dir="${images_dir}/../target"

# Remove the old image to make sure the build fails if this script fails
rm -f "${images_dir}/initramfs-linux.img"

staging="$(mktemp -d)"
trap "rm -rf ${staging}" EXIT

pushd "${target_dir}"
find lib/firmware | cpio --pass-through --make-directories "${staging}"
# find lib/{firmware,systemd} | cpio --pass-through --make-directories "${staging}"
# find bin/{cat,echo,[,[[,sh,mkdir,mount,printf,systemd*} | cpio --pass-through --make-directories "${staging}"
# find sbin/switch_root | cpio --pass-through --make-directories "${staging}"
popd

(cd "${staging}"; find . | cpio -o -H newc | gzip ) > "${images_dir}/initramfs-linux.img"