package dbus

import "github.com/godbus/dbus"

type SystemDbus dbus.Conn

func NewSystemDbus() (*SystemDbus, error) {
	systemConn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	return (*SystemDbus)(systemConn), nil
}
