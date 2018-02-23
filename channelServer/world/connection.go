package world

import (
	"net"

	"github.com/Hucaru/Valhalla/common/connection"
	"github.com/Hucaru/gopacket"
)

type Connection struct {
	conn      *connection.ServerConnection
	worldID   byte
	channelID byte
}

func newConnection(conn net.Conn) *Connection {
	serverConn := &Connection{conn: connection.NewServerConnection(conn)}
	return serverConn
}

func (c *Connection) Write(p gopacket.Packet) error {
	return c.conn.Write(p)
}

func (c *Connection) Read(p gopacket.Packet) error {
	return c.conn.Read(p)
}

func (c *Connection) String() string {
	return c.conn.String()
}

func (c *Connection) SetWorldID(val byte) {
	c.worldID = val
}

func (c *Connection) GetWorldID() byte {
	return c.worldID
}
func (c *Connection) SetchannelID(val byte) {
	c.channelID = val
}

func (c *Connection) GetchannelID() byte {
	return c.channelID
}

func (c *Connection) Close() {
	connected <- false
	c.conn.Close()
}