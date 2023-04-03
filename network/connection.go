package network

import (
	"net"
)

type Connection struct {
	conn     net.Conn
	read_buf []byte
	ch       chan bool
	id       int
}

func static_counter() (f func() int) {
	var i int = 0
	f = func() int {
		i++
		return i
	}
	return
}

var counter func() int = static_counter()

func Create(byte_size int) *Connection {
	connection := new(Connection)
	connection.read_buf = make([]byte, byte_size)
	connection.ch = make(chan bool, 1)
	connection.id = counter()

	return connection
}

func (con *Connection) Connect(address string) {
	conn, err := net.Dial("tcp", ":8080")
	con.conn = conn

	if err != nil {
		conn.Close()
		panic("Network initialize failed")
	}
}

func (con *Connection) Read() []byte {
	read_bytes, err := con.conn.Read(con.read_buf)
	if err != nil || read_bytes == 0 {
		return nil
	}

	var new []byte
	copy(new, con.read_buf)
	return new
}

func (con *Connection) Write(msg []byte) {
	con.conn.Write(msg)
}
