package process

import "strings"

type manager struct {
	processes map[int]*Process
}

var Manager manager

func (m *manager) ExecuteProcess(path, args string) {
	s_args := strings.Split(args, " ")
	p := CreateProcess(path, s_args...)
	p.Run()
	m.processes[p.PID] = p
}

func (m *manager) StopProcess(id int) {
	p := m.processes[id]
	p.Stop()
}

func (m *manager) ForceStopProcess(id int) {
	p := m.processes[id]
	p.ForceStop()
}
