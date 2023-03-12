HUD_FIRST_BOOT_VERSION = 0.0.1
HUD_FIRST_BOOT_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-first-boot
HUD_FIRST_BOOT_SITE_METHOD = local
HUD_FIRST_BOOT_INSTALL_TARGET = YES

define HUD_FIRST_BOOT_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 744 $(HUD_FIRST_BOOT_PKGDIR)/first_boot.sh \
		$(TARGET_DIR)/usr/sbin/first_boot.sh; \
	$(INSTALL) -D -m 644 $(HUD_FIRST_BOOT_PKGDIR)/first_boot.service \
		$(TARGET_DIR)/usr/lib/systemd/system/first_boot.service; \
	$(INSTALL) -D -m 644 $(HUD_FIRST_BOOT_PKGDIR)/50-first_boot.preset \
		$(TARGET_DIR)/etc/systemd/system-preset/50-first_boot.preset;
endef

$(eval $(generic-package))