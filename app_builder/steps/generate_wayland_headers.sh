#!/usr/bin/env bash
set -euo pipefail

STEP_CONFIG="$1"
PROTOCOLS=($(echo "${STEP_CONFIG}" | jq -r '.protocols // [] | .[]'))
DESTINATION="$(echo "${STEP_CONFIG}" | jq -r '.destination')"

WAYLAND_SCANNER="${BUILDROOT_HOST_DIR}/bin/wayland-scanner"

PROTOCOL_ROOT="${BUILDROOT_SYSROOT_DIR}/usr/share/wayland-protocols"

if [ ! -d "${DESTINATION}" ]; then
    mkdir -p "${DESTINATION}"
fi

c_files=""
for protocol in "${PROTOCOLS[@]}"; do
    protocol_name="$(basename "${protocol}" | cut -d'.' -f1)"


    if [[ ${protocol} =~ ^/ ]]; then
        protocol_path="${BUILDROOT_SYSROOT_DIR}${protocol}"
    else
        protocol_path="${PROTOCOL_ROOT}/${protocol}"
    fi

    if [ ! -f "${protocol_path}" ]; then
        echo "Path doesn't exist ${protocol_path}"
        exit 1
    fi

    ${WAYLAND_SCANNER} client-header "${protocol_path}" "${DESTINATION}/${protocol_name}-client-protocol.h"
    ${WAYLAND_SCANNER} private-code "${protocol_path}" "${DESTINATION}/${protocol_name}-protocol.c"

    c_files+="'${protocol_name}-protocol.c',"
done

cat <<EOF > "${DESTINATION}/meson.build"
srcs_libprotocols = [
${c_files}
]
deps_libprotocols = [dep_wayland_client]

lib_libprotocols = static_library(
	'protocols',
	srcs_libprotocols,
	dependencies: deps_libprotocols,
	pic: true,
	install: false
)
dep_libprotocols = declare_dependency(
	link_with: lib_libprotocols,
	dependencies: deps_libprotocols
)
EOF