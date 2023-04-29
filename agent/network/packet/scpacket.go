package packet

import (
	"bytes"
	"encoding/binary"
)

const ReqeustExecuteID uint32 = 1
const ReqeustStopProcessID uint32 = 2
const ReqeustForceStopProcessID uint32 = 3

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
