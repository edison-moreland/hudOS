package dbus

import (
	"errors"
	"fmt"

	"github.com/godbus/dbus"
)

var ErrInvalidType = errors.New("invalid type")

type DBusInterfaceProxy struct {
	dbus.BusObject
	interfaceName string
}

func NewDBusInterfaceProxy(conn *dbus.Conn, service string, node dbus.ObjectPath, interfaceName string) *DBusInterfaceProxy {
	return &DBusInterfaceProxy{
		BusObject:     conn.Object(service, node),
		interfaceName: interfaceName,
	}
}

func (d *DBusInterfaceProxy) Key(key string) string {
	return fmt.Sprintf("%s.%s", d.interfaceName, key)
}

type DBusType interface {
	byte | bool | int | uint |
		int16 | uint16 |
		int32 | uint32 |
		int64 | uint64 |
		float64 | string |
		dbus.ObjectPath |
		dbus.Signature |
		dbus.UnixFDIndex
}

func GetInterfaceProperty[T DBusType](d *DBusInterfaceProxy, key string) (T, error) {
	var zero T

	variant, err := d.GetProperty(key)
	if err != nil {
		return zero, err
	}

	return variant.Value().(T), nil
}

func (d *DBusInterfaceProxy) GetProperty(key string) (dbus.Variant, error) {
	return d.BusObject.GetProperty(d.Key(key))
}

func (d *DBusInterfaceProxy) GetPropertyByte(key string) (byte, error) {
	return GetInterfaceProperty[byte](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyBool(key string) (bool, error) {
	return GetInterfaceProperty[bool](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyInt(key string) (int, error) {
	return GetInterfaceProperty[int](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyUint(key string) (uint, error) {
	return GetInterfaceProperty[uint](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyInt16(key string) (int16, error) {
	return GetInterfaceProperty[int16](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyUint16(key string) (uint16, error) {
	return GetInterfaceProperty[uint16](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyInt32(key string) (int32, error) {
	return GetInterfaceProperty[int32](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyUint32(key string) (uint32, error) {
	return GetInterfaceProperty[uint32](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyInt64(key string) (int64, error) {
	return GetInterfaceProperty[int64](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyUint64(key string) (uint64, error) {
	return GetInterfaceProperty[uint64](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyFloat64(key string) (float64, error) {
	return GetInterfaceProperty[float64](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyString(key string) (string, error) {
	return GetInterfaceProperty[string](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyObjectPath(key string) (dbus.ObjectPath, error) {
	return GetInterfaceProperty[dbus.ObjectPath](d, key)
}

func (d *DBusInterfaceProxy) GetPropertySignature(key string) (dbus.Signature, error) {
	return GetInterfaceProperty[dbus.Signature](d, key)
}

func (d *DBusInterfaceProxy) GetPropertyUnixFDIndex(key string) (dbus.UnixFDIndex, error) {
	return GetInterfaceProperty[dbus.UnixFDIndex](d, key)
}
