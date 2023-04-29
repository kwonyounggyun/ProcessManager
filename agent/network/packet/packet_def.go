package packet

import (
	"bytes"
	"encoding/binary"
)

type PacketData interface {
	Serialize() []byte
	Unserialize(data []byte)
}

func MakePacket(packet_id uint32, data PacketData) []byte {
	s_data := data.Serialize()

	header := Header{ID: packet_id, SIZE: uint32(len(s_data))}
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, &header)

	return append(buf.Bytes(), s_data...)
}

func writeString(buf *bytes.Buffer, str string) {
	len := uint16(len(str))
	binary.Write(buf, binary.BigEndian, len)
	buf.WriteString(str)
}

func readString(buf *bytes.Reader) string {
	var len uint16
	binary.Read(buf, binary.BigEndian, &len)
	str_byte := make([]byte, len)
	buf.Read(str_byte)

	return string(str_byte)
}

// ===================================================================================
type Header struct {
	ID   uint32
	SIZE uint32
}
