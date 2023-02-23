package protocol

import (
	"encoding/binary"
	"errors"
)

//go:generate go run ./scanner/ --package protocol --out ./wayland.go --protocol ../../../buildroot/.buildroot/output/build/wayland-1.21.0/protocol/wayland.xml

var (
	ErrInvalidOp = errors.New("opcode was not recognized by object")
)

var nativeEndian = binary.LittleEndian
