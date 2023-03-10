HUDOS_NETWORK_VERSION = 0.0.1
HUDOS_NETWORK_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hudOS-network
HUDOS_NETWORK_SITE_METHOD = local
HUDOS_NETWORK_INSTALL_TARGET = YES


define HUDOS_NETWORK_INSTALL_TARGET_CMDS
	if [ "$(BR2_PACKAGE_HUDOS_NETWORK_WIFI_SSID)" == "" ]; then \
		echo "Please use 'build.sh nconfig' to set wifi settings under External Settings"; \
		exit 1; \
	fi; \

	mkdir -p $(TARGET_DIR)/etc/wpa_supplicant/; \

	cat $(HUDOS_NETWORK_PKGDIR)/wlan0.conf | \
	sed s/WPA_SUPPLICANT_SSID/$(BR2_PACKAGE_HUDOS_NETWORK_WIFI_SSID)/ | \
	sed s/WPA_SUPPLICANT_PSK/$(BR2_PACKAGE_HUDOS_NETWORK_WIFI_PSK)/ > \
		$(TARGET_DIR)/etc/wpa_supplicant/wpa_supplicant-wlan0.conf; \

	$(INSTALL) -D -m 644 $(HUDOS_NETWORK_PKGDIR)/50-wpa_supplicant.preset \
		$(TARGET_DIR)/etc/systemd/system-preset/50-wpa_supplicant.preset; \
	$(INSTALL) -D -m 644 $(HUDOS_NETWORK_PKGDIR)/wlan0.network \
		$(TARGET_DIR)/etc/systemd/network/wlan0.network;
endef

$(eval $(generic-package))