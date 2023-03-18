HUDOS_DEPLOY_USER_VERSION = 0.0.1
HUDOS_DEPLOY_USER_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hudOS-deploy-user
HUDOS_DEPLOY_USER_SITE_METHOD = local
HUDOS_DEPLOY_USER_INSTALL_TARGET = YES

define HUDOS_DEPLOY_USER_USERS
	deploy -1 deploy -1 - /home/deploy /bin/bash wheel,sudo hudOS Deploy
endef

# brhook will copy the default devices keys into the given directory
define HUDOS_DEPLOY_USER_BUILD_CMDS
	HB_NO_LOCK="buildroot" \
	$(BR2_EXTERNAL_HUDOS_PATH)/../hb devices brhook $(@D)
endef

define HUDOS_DEPLOY_USER_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 644 $(@D)/ed25519.pub \
		$(TARGET_DIR)/home/deploy/.ssh/authorized_keys; \

	$(INSTALL) -D -m 440 $(HUDOS_DEPLOY_USER_PKGDIR)/deploy_sudoers \
		$(TARGET_DIR)/etc/sudoers.d/deploy;
endef

$(eval $(generic-package))