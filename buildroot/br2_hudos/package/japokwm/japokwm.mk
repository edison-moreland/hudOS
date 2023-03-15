JAPOKWM_VERSION = 0.4.2
JAPOKWM_SITE = $(call github,werererer,japokwm,v$(JAPOKWM_VERSION))
JAPOKWM_INSTALL_TARGET = YES
JAPOKWM_DEPENDENCIES = libnotify libxcb json-c wlroots lua libuv wayland wayland-protocols mpfr
JAPOKWM_CONF_OPTS += -Dzsh-completions=false
JAPOKWM_CONF_OPTS += -Dfish-completions=false

$(eval $(meson-package))