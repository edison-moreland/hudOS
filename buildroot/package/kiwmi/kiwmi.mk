KIWMI_VERSION = 17814972abe6a8811a586fa87c99a2b16a86075f
KIWMI_SITE = https://github.com/buffet/kiwmi.git
KIWMI_SITE_METHOD = git
KIWMI_INSTALL_TARGET = YES
KIWMI_DEPENDENCIES = wlroots lua pixman

$(eval $(meson-package))