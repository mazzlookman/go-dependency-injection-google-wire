package google_wire

import "fmt"

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	conn := &Connection{file}
	return conn, func() {
		conn.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("Close Connection" + c.File.Name)
}
