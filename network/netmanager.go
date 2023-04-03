package network

import (
	"net"
	"sync"
)

type NetManager struct {
	connections map[int]*Connection
	mu          *sync.Mutex
	ch          chan bool
}

func CreateManager() *NetManager {
	manager := new(NetManager)
	manager.mu = &sync.Mutex{}
	manager.connections = make(map[int]*Connection)

	return manager
}

func (m *NetManager) Listen(port int) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("Listen failed")
	}

	go func() {
		<-m.ch
		listener.Close()
	}()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				break
			}

			connection := Create(4096)
			connection.conn = conn

			m.mu.Lock()
			m.connections[connection.id] = connection
			m.mu.Unlock()
		}
	}()
}
