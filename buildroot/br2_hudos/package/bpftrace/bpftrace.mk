BPFTRACE_VERSION = v0.17.0
BPFTRACE_SITE = https://github.com/iovisor/bpftrace.git
BPFTRACE_SITE_METHOD = git
BPFTRACE_INSTALL_STAGING = YES
BPFTRACE_INSTALL_TARGET = YES
BPFTRACE_CONF_OPTS = -DUSE_SYSTEM_BPF_BCC=ON -DSTATIC_LINKING=OFF
BPFTRACE_DEPENDENCIES = bcc binutils cereal libbpf libpcap
BPFTRACE_SUPPORTS_IN_SOURCE_BUILD = NO

define BPFTRACE_INSTALL_LIB_STAGING
	for file in $$(find $(@D)/buildroot-build/ -type f -name "*.so"); \
	do $(INSTALL) -D -m 755 "$${file}" $(STAGING_DIR)/usr/lib/; \
	done;
endef
BPFTRACE_POST_INSTALL_STAGING_HOOKS += BPFTRACE_INSTALL_LIB_STAGING

define BPFTRACE_INSTALL_LIB_TARGET
	for file in $$(find $(@D)/buildroot-build/ -type f -name "*.so"); \
	do $(INSTALL) -D -m 755 "$${file}" $(TARGET_DIR)/usr/lib/; \
	done;
endef
BPFTRACE_POST_INSTALL_TARGET_HOOKS += BPFTRACE_INSTALL_LIB_TARGET

$(eval $(cmake-package))