ENVSUBST_VERSION = v1.4.2
ENVSUBST_SITE = $(call github,a8m,envsubst,$(ENVSUBST_VERSION))
ENVSUBST_EXTRA_DOWNLOADS = https://github.com/a8m/envsubst/releases/download/$(ENVSUBST_VERSION)/envsubst-Linux-arm64
ENVSUBST_INSTALL_TARGET = YES

define ENVSUBST_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 755 $(ENVSUBST_DL_DIR)/envsubst-Linux-arm64 \
		$(TARGET_DIR)/usr/bin/envsubst
endef

$(eval $(generic-package))