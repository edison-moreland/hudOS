package protocol

import (
	"github.com/edison-moreland/nreal-hud/sdk/log"
	"math"
)

// From wayland/wayland-util.h
func fixedToFloat64(f int32) float64 {
	u_i := (1023+44)<<52 + (1 << 51) + int64(f)
	u_d := math.Float64frombits(uint64(u_i))
	return u_d - (3 << 43)
}

func fixedFromfloat64(d float64) int32 {
	u_d := d + (3 << (51 - 8))
	u_i := int64(math.Float64bits(u_d))
	return int32(u_i)
}

func roundup(v uint32) uint32 {
	return uint32((uint64(v) + (32 - 1)) / 32)
}

func UnmarshallInt32(offset int, data []byte) (int, int32) {
	val := int32(nativeEndian.Uint32(data[offset : offset+4]))

	log.Debug().
		Int("offset", offset).
		Int32("value", val).
		Msg("Unmarshalling Int32")

	return offset + 4, val
}

func UnmarshallUint32(offset int, data []byte) (int, uint32) {
	val := nativeEndian.Uint32(data[offset : offset+4])

	log.Debug().
		Int("offset", offset).
		Uint32("value", val).
		Msg("Unmarshalling Uint32")

	return offset + 4, val
}

func UnmarshallString(offset int, data []byte) (int, string) {
	strLength := nativeEndian.Uint32(data[offset : offset+4])
	val := string(data[offset+4 : (offset+4)+int(strLength)])

	log.Debug().
		Int("offset", offset).
		Uint32("length", strLength).
		Str("value", val).
		Msg("Unmarshalling String")

	return int(roundup(uint32(offset) + 4 + strLength)), val
}

func UnmarshallArray(offset int, data []byte) (int, []byte) {
	arrLength := nativeEndian.Uint32(data[offset : offset+4])

	val := make([]byte, arrLength)
	copy(val, data[offset+4:(offset+4)+int(arrLength)])

	log.Debug().
		Int("offset", offset).
		Uint32("length", arrLength).
		Hex("value", val).
		Msg("Unmarshalling Array")

	return int(roundup(uint32(offset) + 4 + arrLength)), val
}

func UnmarshallFd(offset int, data []byte) (int, uintptr) {
	log.Panic().Msg("UnmarshallFd not implemented")
	return offset, 0
}

func UnmarshallFixed(offset int, data []byte) (int, float64) {
	raw := nativeEndian.Uint32(data[offset : offset+4])
	val := fixedToFloat64(int32(raw))

	log.Debug().
		Int("offset", offset).
		Uint32("raw", raw).
		Float64("value", val).
		Msg("Unmarshalling Fixed")

	return offset + 4, fixedToFloat64(int32(val))
}
