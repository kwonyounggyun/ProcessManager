package network

import (
	"net"
)

type Connection struct {
	conn     net.Conn
	read_buf []byte
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
	connection.id = counter()

	return connection
}

func (con *Connection) Connect(address string) {
	conn, err := net.Dial("tcp", address)
	con.conn = conn

	if err != nil {
		conn.Close()
		panic("Network initialize failed")
	}
}

func (con *Connection) Read() ([]byte, error) {
	read_bytes, err := con.conn.Read(con.read_buf)
	if err != nil {
		return nil, err
	}

	new := make([]byte, read_bytes)
	copy(new, con.read_buf)
	return new, nil
}

func (con *Connection) Write(msg []byte) error {
	_, err := con.conn.Write(msg)

	if err != nil {
		return err
	}

	return nil
}

func (con *Connection) Close() {
	if con.conn != nil {
		con.conn.Close()
		con.conn = nil
	}
}
