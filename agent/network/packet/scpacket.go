package packet

import (
	"bytes"
	"encoding/binary"
)

const ReqeustExecuteID uint32 = 1
const ReqeustStopProcessID uint32 = 2
const ReqeustForceStopProcessID uint32 = 3

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

type ReqeustExecute struct {
	Path string
	Args string
}

func (p *ReqeustExecute) Serialize() []byte {
	packet := make([]byte, 0, 4096)
	buf := bytes.NewBuffer(packet)

	writeString(buf, p.Path)
	writeString(buf, p.Args)

	return packet[:buf.Len()]
}

func (p *ReqeustExecute) Unserialize(data []byte) {
	buf := bytes.NewReader(data)

	p.Path = readString(buf)
	p.Args = readString(buf)
}

type ReqeustStopProcess struct {
	PID int
}

func (p *ReqeustStopProcess) Serialize() []byte {
	packet := make([]byte, 0, 4096)
	buf := bytes.NewBuffer(packet)

	binary.Write(buf, binary.BigEndian, &p.PID)

	return packet[:buf.Len()]
}

func (p *ReqeustStopProcess) Unserialize(data []byte) {
	buf := bytes.NewReader(data)

	binary.Read(buf, binary.BigEndian, &p.PID)
}

type ReqeustForceStopProcess struct {
	PID int
}

func (p *ReqeustForceStopProcess) Serialize() []byte {
	packet := make([]byte, 0, 4096)
	buf := bytes.NewBuffer(packet)

	binary.Write(buf, binary.BigEndian, &p.PID)

	return packet[:buf.Len()]
}

func (p *ReqeustForceStopProcess) Unserialize(data []byte) {
	buf := bytes.NewReader(data)

	binary.Read(buf, binary.BigEndian, &p.PID)
}
