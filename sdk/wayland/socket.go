package wayland

import (
	"errors"
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"github.com/edison-moreland/nreal-hud/sdk/xdg"
	"net"
	"os"
	"path"
	"strconv"
)

var (
	ErrInvalidFileDescriptor = errors.New("file descriptor is invalid")
	ErrNoSocketFound         = errors.New("wayland socket could not be found after exhausting all options")
)

func OpenSocket() (net.Conn, error) {
	if wsocket, ok := os.LookupEnv("WAYLAND_SOCKET"); ok {
		conn, err := openExistingSocket(wsocket)
		if err != nil {
			log.Warn().
				Err(err).
				Str("WAYLAND_SOCKET", wsocket).
				Msg("Could not open WAYLAND_SOCKET! Moving on...")
		} else {
			return conn, nil
		}
	}

	xdgRuntimeDir := xdg.RuntimeDir()

	if wdisplay, ok := os.LookupEnv("WAYLAND_DISPLAY"); ok {
		conn, err := openSocketPath(xdgRuntimeDir, wdisplay)
		if err != nil {
			log.Warn().
				Err(err).
				Str("WAYLAND_DISPLAY", wdisplay).
				Str("XDG_RUNTIME_DIR", xdgRuntimeDir).
				Msg("Could not open WAYLAND_DISPLAY! Moving on...")
		} else {
			return conn, nil
		}
	}

	conn, err := openSocketPath(xdgRuntimeDir, "wayland-0")
	if err != nil {
		log.Warn().
			Err(err).
			Str("XDG_RUNTIME_DIR", xdgRuntimeDir).
			Msg("Could not open wayland-0! Giving up.")
	} else {
		return conn, nil
	}

	return nil, ErrNoSocketFound
}

func openExistingSocket(socketFd string) (net.Conn, error) {
	fd, err := strconv.ParseUint(socketFd, 0, 64)
	if err != nil {
		return nil, err
	}

	socket := os.NewFile(uintptr(fd), "wayland-0")
	if socket == nil {
		return nil, ErrInvalidFileDescriptor
	}

	conn, err := net.FileConn(socket)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func openSocketPath(base, socket string) (net.Conn, error) {
	return net.Dial("unix", path.Join(base, socket))
}
