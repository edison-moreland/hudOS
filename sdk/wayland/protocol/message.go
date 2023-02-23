package protocol

import (
	"errors"
)

type Message struct {
	MessageHeader
	Data []byte
}

func (m *Message) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8+len(m.Data))
	headerData, err := m.MessageHeader.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(data[:8], headerData)
	copy(data[8:], m.Data)

	return data, nil
}

type MessageHeader struct {
	ObjectID      uint32
	MessageLength uint16
	Opcode        uint16
}

func (m *MessageHeader) UnmarshallBinary(data []byte) error {
	if len(data) < 8 {
		return errors.New("not enough Data")
	}

	m.ObjectID = nativeEndian.Uint32(data[0:4])
	m.MessageLength = nativeEndian.Uint16(data[4:6])
	m.Opcode = nativeEndian.Uint16(data[6:8])

	return nil
}

func (m *MessageHeader) MarshalBinary() ([]byte, error) {
	data := make([]byte, 8)
	nativeEndian.PutUint32(data[0:4], m.ObjectID)
	nativeEndian.PutUint16(data[4:6], m.MessageLength)
	nativeEndian.PutUint16(data[6:8], m.Opcode)

	return data, nil
}
