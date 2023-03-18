LIBNOTIFY_VERSION = 69aff6e5fa2842e00b409c348bd73188548828b3
LIBNOTIFY_SITE = https://gitlab.gnome.org/GNOME/libnotify.git
LIBNOTIFY_SITE_METHOD = git
LIBNOTIFY_INSTALL_STAGING = YES
LIBNOTIFY_INSTALL_TARGET = YES
LIBNOTIFY_DEPENDENCIES = libgtk3 gdk-pixbuf
LIBNOTIFY_CONF_OPTS += -Dtests=false 
LIBNOTIFY_CONF_OPTS += -Dintrospection=disabled
LIBNOTIFY_CONF_OPTS += -Dman=false
LIBNOTIFY_CONF_OPTS += -Dgtk_doc=false
LIBNOTIFY_CONF_OPTS += -Ddocbook_docs=disabled

$(eval $(meson-package))