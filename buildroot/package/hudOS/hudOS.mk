HUDOS_VERSION = 0.0.1
HUDOS_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hudOS
HUDOS_SITE_METHOD = local
HUDOS_INSTALL_TARGET = YES

define HUDOS_USERS
	hud -1 hud -1 * /opt/hud /bin/false seat hudOS
endef

define HUDOS_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 644 $(HUDOS_PKGDIR)/50-disable_getty.preset \
		$(TARGET_DIR)/etc/systemd/system-preset/50-disable_getty.preset; \
	$(INSTALL) -D -m 644 $(HUDOS_PKGDIR)/fw_env.config \
		$(TARGET_DIR)/etc/fw_env.config; \
	mkdir -p $(TARGET_DIR)/opt/hud/systemd/system/; \
	mkdir -p $(TARGET_DIR)/usr/share/icons/;
endef

# These could probably find a better place to live
# This is so setcap work for the bluetooth daemon
define HUDOS_LINUX_CONFIG_FIXUPS
	$(call KCONFIG_ENABLE_OPT,CONFIG_EXT4_FS_SECURITY)
	$(call KCONFIG_ENABLE_OPT,CONFIG_EXT4_FS_XATTR)
endef

$(eval $(generic-package))