package xdg

import (
	"github.com/rs/zerolog/log"
	"os"
	"path"
)

// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

func ConfigHome() string {
	return getXDGPath("XDG_CONFIG_HOME", "$HOME/.config/")
}

func CacheHome() string {
	return getXDGPath("XDG_CACHE_HOME", "$HOME/.cache/")
}

func DataHome() string {
	return getXDGPath("XDG_DATA_HOME", "$HOME/.local/share/")
}

func StateHome() string {
	return getXDGPath("XDG_STATE_HOME", "$HOME/.local/state")
}

func RuntimeDir() string {
	// RuntimeDir is the only one that doesn't have a default
	return getXDGPath("XDG_RUNTIME_DIR", "")
}

func getXDGPath(name, defaultPath string) string {
	if xdgPath, ok := os.LookupEnv(name); ok {
		if path.IsAbs(xdgPath) {
			return xdgPath
		}
		log.Warn().
			Str(name, xdgPath).
			Msg("Path is not absolute, ignoring.")
	} else {
		log.Warn().
			Str("xdg_dir", name).
			Msg("Env var is not set, moving on.")
	}

	if defaultPath == "" {
		log.Panic().
			Str("xdg_dir", name).
			Msg("Could not find dir.")
	}

	xdgPath := os.ExpandEnv(defaultPath)
	if _, err := os.Stat(xdgPath); os.IsNotExist(err) {
		log.Panic().
			Str("xdg_dir", name).
			Str("default", xdgPath).
			Msg("Default path does not exist!")

	}

	return xdgPath
}
