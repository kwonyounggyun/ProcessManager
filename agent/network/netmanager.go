package network

import (
	"fmt"
	"net"
	"sync"
)

type NetManager struct {
	connections map[string]*Connection
	wg          *sync.WaitGroup
	filter      map[string]string

	exit_event   chan bool
	accept_event chan *Connection
	close_event  chan *Connection
}

func CreateManager() *NetManager {
	manager := new(NetManager)
	manager.connections = make(map[string]*Connection)
	manager.wg = &sync.WaitGroup{}
	manager.filter = make(map[string]string)

	manager.exit_event = make(chan bool)
	manager.accept_event = make(chan *Connection)
	manager.close_event = make(chan *Connection)

	return manager
}

func (m *NetManager) Listen(port int) {
	con_str := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", con_str)

	if err != nil {
		panic("Listen failed")
	}

	m.wg.Add(2)
	//event process loop
	go func() {
		defer m.wg.Done()
		for {
			select {
			case accept_con := <-m.accept_event:
				{
					m.connections[accept_con.name] = accept_con

					//create read loop
					m.wg.Add(1)
					go m.clientReadLoop(accept_con)
					break
				}
			case close_con := <-m.close_event:
				{
					find := m.connections[close_con.name]
					if find != nil {
						close_con.conn.Close()
						delete(m.connections, close_con.name)
					}
					break
				}
			case <-m.exit_event:
				{
					listener.Close()
					for _, con := range m.connections {
						con.Close()
					}
					return
				}
			}
		}
	}()

	//listen process loop
	go func() {
		defer m.wg.Done()
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

			m.accept_event <- connection
		}
	}()
}

func (m *NetManager) clientReadLoop(con *Connection) {
	defer m.wg.Done()
	for {
		read, err := con.Read()
		if err != nil {
			m.close_event <- con
			break
		}

		fmt.Print(string(read))
	}
}

func (m *NetManager) Stop() {
	m.exit_event <- false
	m.wg.Wait()
}

func (m *NetManager) SendPacket(name string, data []byte) {
	con := m.connections[name]
	if con == nil {
		return
	}

	con.Write(data)
}
