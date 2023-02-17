HUD_LANDING_PAD_VERSION = 0.0.1
HUD_LANDING_PAD_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-landing-pad
HUD_LANDING_PAD_SITE_METHOD = local
PINE_FIRMWARE_INSTALL_TARGET = YES

define HUD_LANDING_PAD_USERS
	hud -1 hud -1 * /opt/hud - - hudOS
	deploy -1 deploy -1 =changechangethis /home/deploy /bin/bash wheel,sudo hudOS Deploy
endef

# TODO
# Add depends for wpa_supplicant, and systemd-networkd
# Created symlink /etc/systemd/system/multi-user.target.wants/wpa_supplicant@wlan0.service → /usr/lib/systemd/system/wpa_supplicant@.service

define HUD_LANDING_PAD_INSTALL_TARGET_CMDS
	if [ "$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_SSID)" == "" ]; then \
		echo "Please use 'build.sh nconfig' to set wifi settings under External Settings"; \
		exit 1; \
	fi; \

	mkdir -p $(TARGET_DIR)/etc/wpa_supplicant/; \

	cat $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-landing-pad/wlan0.conf | \
	sed s/WPA_SUPPLICANT_SSID/$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_SSID)/ | \
	sed s/WPA_SUPPLICANT_PSK/$(BR2_PACKAGE_HUD_LANDING_PAD_WIFI_PASSWORD)/ > \
		$(TARGET_DIR)/etc/wpa_supplicant/wpa_supplicant-wlan0.conf; \

	$(INSTALL) -D -m 644 $(HUD_LANDING_PAD_PKGDIR)/20-wifi.preset \
		$(TARGET_DIR)/usr/lib/systemd/system-preset/20-wifi.preset;
endef

$(eval $(generic-package))