package dbus

// https://github.com/svenwltr/i3-statusbar/blob/master/upower/upower.go

import (
	"time"

	"github.com/godbus/dbus"
)

const (
	upowerService           = "org.freedesktop.UPower"
	upowerDeviceInterface   = "org.freedesktop.UPower.Device"
	upowerDisplayDeviceNode = "/org/freedesktop/UPower/devices/DisplayDevice"
)

type UPowerDaemon dbus.Conn

func (s *SystemDbus) UPowerDaemon() *UPowerDaemon {
	return (*UPowerDaemon)(s)
}

type UPowerDevice DBusInterfaceProxy

func (u *UPowerDaemon) DisplayDevice() *UPowerDevice {
	return u.Device(upowerDisplayDeviceNode)
}

func (u *UPowerDaemon) Device(deviceNode dbus.ObjectPath) *UPowerDevice {
	d := NewDBusInterfaceProxy((*dbus.Conn)(u), upowerService, deviceNode, upowerDeviceInterface)

	return (*UPowerDevice)(d)
}

func (d *UPowerDevice) GetPercentage() (int, error) {
	f, err := (*DBusInterfaceProxy)(d).GetPropertyFloat64("Percentage")
	if err != nil {
		return 0, err
	}

	return int(f), nil

}

func (d *UPowerDevice) GetTimeToFull() (time.Duration, error) {
	i, err := (*DBusInterfaceProxy)(d).GetPropertyInt64("TimeToFull")
	if err != nil {
		return 0, err
	}

	return time.Duration(int(i) * int(time.Second)), nil

}

func (d *UPowerDevice) GetTimeToEmpty() (time.Duration, error) {
	i, err := (*DBusInterfaceProxy)(d).GetPropertyInt64("TimeToEmpty")
	if err != nil {
		return 0, err
	}

	return time.Duration(int(i) * int(time.Second)), nil

}

type UPowerDeviceState uint32

const (
	UNKNOWN           UPowerDeviceState = iota
	CHARGING          UPowerDeviceState = iota
	DISCHARGING       UPowerDeviceState = iota
	EMPTY             UPowerDeviceState = iota
	FULLY_CHARGED     UPowerDeviceState = iota
	PENDING_CHARGE    UPowerDeviceState = iota
	PENDING_DISCHARGE UPowerDeviceState = iota
)

func (d *UPowerDevice) GetState() (UPowerDeviceState, error) {
	i, err := (*DBusInterfaceProxy)(d).GetPropertyUint32("State")
	if err != nil {
		return 0, err
	}

	return UPowerDeviceState(i), nil

}
