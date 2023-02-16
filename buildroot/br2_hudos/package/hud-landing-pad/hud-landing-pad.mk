HUD_LANDING_PAD_VERSION = 0.0.1
HUD_LANDING_PAD_SITE = $(BR2_EXTERNAL_HUDOS_PATH)/package/hud-landing-pad
HUD_LANDING_PAD_SITE_METHOD = local

define HUD_LANDING_PAD_USERS
	hud -1 hud -1 * /opt/hud - - hudOS
	deploy -1 deploy -1 =changechangethis /home/deploy /bin/bash wheel,sudo hudOS Deploy
endef

$(eval $(generic-package))