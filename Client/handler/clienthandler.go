package handler

import (
	"ProcessManager/agent/network/packet"
	"ProcessManager/agent/process"
)

var Handle = map[uint32](func([]byte) bool){
	packet.ReqeustExecuteID:          HandleRequestExecute,
	packet.ReqeustStopProcessID:      HandleRequestStopProcess,
	packet.ReqeustForceStopProcessID: HandleRequestForceStopProcess,
}

func HandleRequestExecute(msg []byte) bool {
	var packetMsg packet.ReqeustExecute
	packetMsg.Unserialize(msg)

	process.Manager.ExecuteProcess(packetMsg.Path, packetMsg.Args)

	return true
}

func HandleRequestStopProcess(msg []byte) bool {
	var packetMsg packet.ReqeustStopProcess
	packetMsg.Unserialize(msg)

	process.Manager.StopProcess(packetMsg.PID)
	return true
}

func HandleRequestForceStopProcess(msg []byte) bool {
	var packetMsg packet.ReqeustForceStopProcess
	packetMsg.Unserialize(msg)

	process.Manager.ForceStopProcess(packetMsg.PID)
	return true
}
