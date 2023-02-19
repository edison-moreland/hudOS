HUD_UI_VERSION = 0.0.1
HUD_UI_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-ui
HUD_UI_SITE_METHOD = local
HUD_UI_INSTALL_TARGET = YES

define HUD_UI_USERS
	hud -1 hud -1 * /opt/hud /bin/false seat hudOS
endef

define HUD_UI_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 644 $(HUD_UI_PKGDIR)/50-disable_getty.preset \
		$(TARGET_DIR)/etc/systemd/system-preset/50-disable_getty.preset;
endef

$(eval $(generic-package))