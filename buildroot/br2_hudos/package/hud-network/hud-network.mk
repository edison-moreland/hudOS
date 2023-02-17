HUD_NETWORK_VERSION = 0.0.1
HUD_NETWORK_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-network
HUD_NETWORK_SITE_METHOD = local
HUD_NETWORK_INSTALL_TARGET = YES


define HUD_NETWORK_INSTALL_TARGET_CMDS
	if [ "$(BR2_PACKAGE_HUD_NETWORK_WIFI_SSID)" == "" ]; then \
		echo "Please use 'build.sh nconfig' to set wifi settings under External Settings"; \
		exit 1; \
	fi; \

	mkdir -p $(TARGET_DIR)/etc/wpa_supplicant/; \

	cat $(HUD_NETWORK_PKGDIR)/wlan0.conf | \
	sed s/WPA_SUPPLICANT_SSID/$(BR2_PACKAGE_HUD_NETWORK_WIFI_SSID)/ | \
	sed s/WPA_SUPPLICANT_PSK/$(BR2_PACKAGE_HUD_NETWORK_WIFI_PSK)/ > \
		$(TARGET_DIR)/etc/wpa_supplicant/wpa_supplicant-wlan0.conf; \

	$(INSTALL) -D -m 644 $(HUD_NETWORK_PKGDIR)/20-wlan0.preset \
		$(TARGET_DIR)/usr/lib/systemd/system-preset/20-wlan0.preset; \
	$(INSTALL) -D -m 644 $(HUD_NETWORK_PKGDIR)/wlan0.network \
		$(TARGET_DIR)/etc/systemd/network/wlan0.network;
endef

$(eval $(generic-package))