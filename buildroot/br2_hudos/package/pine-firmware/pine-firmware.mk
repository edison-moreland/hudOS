# Copied almost entirely from buildroot's linux-firmware package

PINE_FIRMWARE_VERSION = 5c4c2b89f30a42f5ffabb5b5bcbc799d8ac9f66f
PINE_FIRMWARE_SITE = "https://megous.com/git/linux-firmware"
PINE_FIRMWARE_SITE_METHOD = git
PINE_FIRMWARE_INSTALL_IMAGES = YES
PINE_FIRMWARE_INSTALL_TARGET = YES


ifeq ($(BR2_PACKAGE_PINE_FIRMWARE_ANX7688),y)
PINE_FIRMWARE_FILES += anx7688-fw.bin
endif

ifeq ($(BR2_PACKAGE_PINE_FIRMWARE_OV5640),y)
PINE_FIRMWARE_FILES += ov5640_af.bin
endif

ifeq ($(BR2_PACKAGE_PINE_FIRMWARE_RTL8723CS),y)
PINE_FIRMWARE_FILES += \
	rtl_bt/rtl8723cs_xx_config-pinephone.bin \
	rtl_bt/rtl8723cs_xx_config.bin \
	rtl_bt/rtl8723cs_xx_fw.bin
endif


ifneq ($(PINE_FIRMWARE_FILES),)

define PINE_FIRMWARE_BUILD_CMDS
	cd $(@D) && \
	$(TAR) cf br-firmware.tar $(sort $(PINE_FIRMWARE_FILES))
endef

define PINE_FIRMWARE_INSTALL_FW
	mkdir -p $(1)
	$(TAR) xf $(@D)/br-firmware.tar -C $(1)
endef

endif  # PINE_FIRMWARE_FILES

define PINE_FIRMWARE_INSTALL_TARGET_CMDS
	echo "Installing pine-firmware to target"
	$(call PINE_FIRMWARE_INSTALL_FW, $(TARGET_DIR)/lib/firmware)
endef

define PINE_FIRMWARE_INSTALL_IMAGES_CMDS
	echo "Installing pine-firmware to images"
	$(call PINE_FIRMWARE_INSTALL_FW, $(BINARIES_DIR))
endef

$(eval $(generic-package))
