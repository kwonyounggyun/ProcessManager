package network

import (
	"fmt"
	"net"
	"sync"
)

type NetManager struct {
	connections map[string]*Connection
	mu          *sync.Mutex
	ch          chan bool
	wg          *sync.WaitGroup
	filter      map[string]string
}

func CreateManager() *NetManager {
	manager := new(NetManager)
	manager.mu = &sync.Mutex{}
	manager.connections = make(map[string]*Connection)
	manager.ch = make(chan bool)
	manager.wg = &sync.WaitGroup{}
	manager.filter = make(map[string]string)

	return manager
}

func (m *NetManager) Listen(port int) {
	con_str := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", con_str)

	if err != nil {
		panic("Listen failed")
	}

	m.wg.Add(2)
	go func() {
		<-m.ch
		listener.Close()

		m.wg.Done()
	}()

	go func() {
		defer func() {
			m.mu.Lock()
			for _, con := range m.connections {
				con.Close()
				delete(m.connections, con.name)
			}
			m.mu.Unlock()
		}()

		for {
			conn, err := listener.Accept()
			if err != nil {
				break
			}

			remote_addr := conn.RemoteAddr().String()

			name := m.filter[remote_addr]
			if len(name) == 0 {
				conn.Close()
				continue
			}

			connection := Create(name)
			connection.conn = conn

			m.mu.Lock()
			m.connections[connection.name] = connection
			m.mu.Unlock()

			// var data mypack.ReqeustExecute
			// data.Path = "calc"
			// data.Args = "nono"

			// packet := mypack.MakePacket(mypack.ReqeustExecuteID, &data)

			// connection.Write(packet)
			m.wg.Add(1)
			go m.clientReadLoop(connection)
		}

		m.wg.Done()
	}()
}

func (m *NetManager) clientReadLoop(con *Connection) {
	for {
		read, err := con.Read()
		if err != nil {
			m.mu.Lock()

			con.Close()
			delete(m.connections, con.name)
			m.mu.Unlock()

			m.wg.Done()
			break
		}

		fmt.Print(string(read))
	}
}

func (m *NetManager) Stop() {
	m.ch <- false
	m.wg.Wait()
}

func (m *NetManager) SendPacket(name string, data []byte) {
	con := m.connections[name]
	if con == nil {
		return
	}

	con.Write(data)
}