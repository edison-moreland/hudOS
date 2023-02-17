HUD_DEPLOY_USER_VERSION = 0.0.1
HUD_DEPLOY_USER_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-deploy-user
HUD_DEPLOY_USER_SITE_METHOD = local
HUD_DEPLOY_USER_INSTALL_TARGET = YES

define HUD_DEPLOY_USER_USERS
	deploy -1 deploy -1 - /home/deploy /bin/bash wheel,sudo hudOS Deploy
endef

define HUD_DEPLOY_USER_BUILD_CMDS
	rm -f $(@D)/deploy_ed25519 $(@D)/deploy_ed25519.pub; \
	ssh-keygen -t ed25519 -N "" -C "hudOS Development Key" -f $(@D)/deploy_ed25519
endef

define HUD_DEPLOY_USER_INSTALL_TARGET_CMDS
	$(INSTALL) -D -m 644 $(@D)/deploy_ed25519.pub \
		$(TARGET_DIR)/home/deploy/.ssh/authorized_keys; \

	$(INSTALL) -D -m 440 $(HUD_DEPLOY_USER_PKGDIR)/deploy_sudoers \
		$(TARGET_DIR)/etc/sudoers.d/deploy;
endef

$(eval $(generic-package))