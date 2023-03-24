ALSA_UCM_CONF_VERSION = v1.2.8
ALSA_UCM_CONF_SITE = $(call github,alsa-project,alsa-ucm-conf,$(ALSA_UCM_CONF_VERSION))
ALSA_UCM_CONF_INSTALL_TARGET = YES
ALSA_UCM_CONF_DEPENDENCIES = alsa-lib alsa-utils

define ALSA_UCM_CONF_BUILD_CMDS
	$(INSTALL) -D -m 644 $(ALSA_UCM_CONF_PKGDIR)/ppp/PinePhonePro.conf \
		$(@D)/ucm2/Pine64/PinePhonePro/PinePhonePro.conf; \
	$(INSTALL) -D -m 644 $(ALSA_UCM_CONF_PKGDIR)/ppp/HiFi.conf \
		$(@D)/ucm2/Pine64/PinePhonePro/HiFi.conf; \
	$(INSTALL) -D -m 644 $(ALSA_UCM_CONF_PKGDIR)/ppp/VoiceCall.conf \
		$(@D)/ucm2/Pine64/PinePhonePro/VoiceCall.conf; \

    ln -s Pine64/PinePhonePro $(@D)/ucm2/PinePhonePro; \

    mkdir -p $(@D)/ucm2/conf.d/simple-card; \
    ln -sf ../../Pine64/PinePhonePro/PinePhonePro.conf \
        $(@D)/ucm2/conf.d/simple-card/PinePhonePro.conf; \

	cd $(@D) && \
	$(TAR) -cf alsa-ucm-conf.tar .
endef

define ALSA_UCM_CONF_INSTALL_TARGET_CMDS
	$(TAR) -xf $(@D)/alsa-ucm-conf.tar -C $(TARGET_DIR)/usr/share/alsa
endef

$(eval $(generic-package))