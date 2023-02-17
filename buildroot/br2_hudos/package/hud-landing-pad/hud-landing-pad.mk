HUD_LANDING_PAD_VERSION = 0.0.1
HUD_LANDING_PAD_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-landing-pad
HUD_LANDING_PAD_SITE_METHOD = local
PINE_FIRMWARE_INSTALL_TARGET = YES


ifeq ($(BR2_PACKAGE_HUD_LANDING_PAD_DEPLOY_USER),y)
HUD_LANDING_PAD_DEPLOY_USER = deploy -1 deploy -1 - /home/deploy /bin/bash wheel,sudo hudOS Deploy

define HUD_LANDING_PAD_BUILD_DEPLOY_USER
	rm -f $(@D)/deploy_ed25519 $(@D)/deploy_ed25519.pub; \
	ssh-keygen -t ed25519 -N "" -C "hudOS Development Key" -f $(@D)/deploy_ed25519
endef

define HUD_LANDING_PAD_SETUP_DEPLOY_USER
	$(INSTALL) -D -m 644 $(@D)/deploy_ed25519.pub \
		$(TARGET_DIR)/home/deploy/.ssh/authorized_keys; \

	$(INSTALL) -D -m 440 $(HUD_LANDING_PAD_PKGDIR)/deploy_sudoers \
		$(TARGET_DIR)/etc/sudoers.d/deploy;
endef

endif

define HUD_LANDING_PAD_SETUP_WIFI
	if [ "$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_SSID)" == "" ]; then \
		echo "Please use 'build.sh nconfig' to set wifi settings under External Settings"; \
		exit 1; \
	fi; \

	mkdir -p $(TARGET_DIR)/etc/wpa_supplicant/; \

	cat $(HUD_LANDING_PAD_PKGDIR)/wlan0.conf | \
	sed s/WPA_SUPPLICANT_SSID/$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_SSID)/ | \
	sed s/WPA_SUPPLICANT_PSK/$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_PASSWORD)/ > \
		$(TARGET_DIR)/etc/wpa_supplicant/wpa_supplicant-wlan0.conf; \

	$(INSTALL) -D -m 644 $(HUD_LANDING_PAD_PKGDIR)/20-wifi.preset \
		$(TARGET_DIR)/usr/lib/systemd/system-preset/20-wifi.preset; \
	$(INSTALL) -D -m 644 $(HUD_LANDING_PAD_PKGDIR)/wlan0.network \
		$(TARGET_DIR)/etc/systemd/network/wlan0.network;
endef

define HUD_LANDING_PAD_SETUP_AUTOLOGIN
	mkdir -p $(TARGET_DIR)/etc/systemd/system/getty@tty1.service.d/; \
	$(INSTALL) -D -m 644 $(HUD_LANDING_PAD_PKGDIR)/autologin.conf \
		$(TARGET_DIR)/etc/systemd/system/getty@tty1.service.d/autologin.conf;
endef

define HUD_LANDING_PAD_BUILD_CMDS
	$(HUD_LANDING_PAD_BUILD_DEPLOY_USER)
endef

define HUD_LANDING_PAD_USERS
	hud -1 hud -1 * /opt/hud /bin/bash - hudOS
	$(HUD_LANDING_PAD_DEPLOY_USER)
endef

define HUD_LANDING_PAD_INSTALL_TARGET_CMDS
	$(HUD_LANDING_PAD_SETUP_WIFI)
	$(HUD_LANDING_PAD_SETUP_AUTOLOGIN)
	$(HUD_LANDING_PAD_SETUP_DEPLOY_USER)
endef

$(eval $(generic-package))