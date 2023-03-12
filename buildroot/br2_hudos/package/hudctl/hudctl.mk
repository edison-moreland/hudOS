HUDCTL_VERSION = 0.0.1
HUDCTL_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/../../apps/hudctl
HUDCTL_SITE_METHOD = local
HUDCTL_INSTALL_TARGET = YES

define HUDCTL_BUILD_CMDS
	$(BR2_EXTERNAL_HUDOS_PATH)/../../app_builder/app_builder.sh -i $(@D)/.hud_app.json -o $(@D)/hudctl.tar
endef

define HUDCTL_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 644 $(@D)/hudctl.tar \
		$(TARGET_DIR)/opt/hud/bootstrap/hudctl.tar;
	$(INSTALL) -D -m 755 $(HUDCTL_PKGDIR)/bootstrap.sh \
		$(TARGET_DIR)/opt/hud/firstboot/00-hudctl-bootstrap.sh;
endef

$(eval $(generic-package))