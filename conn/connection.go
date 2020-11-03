package conn

import (
	"mysql-scanner/protocol"
	"net"
	"time"
)

type Conn interface {
	Read() error
	Closed() bool
	Cleanup()
	NextPacket() ([]byte, error)
	Greeting() *protocol.Greeting
}

type conn struct {
	netConn  net.Conn
	greeting *protocol.Greeting
	packets  *Packets
}

func (c *conn) Greeting() *protocol.Greeting {
	return c.greeting
}

func Connect(address string, timeout time.Duration) (Conn, error) {
	var err error
	c := &conn{greeting:&protocol.Greeting{}}
	if c.netConn, err = net.DialTimeout("tcp", address, timeout); err != nil {
		return nil, err
	}

	// Set KeepAlive to True and period to 180s.
	if tcpConn, ok := c.netConn.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(time.Second * 180)
		c.netConn = tcpConn
	}

	c.netConn.SetReadDeadline(time.Now().Add(timeout))
	defer c.netConn.SetReadDeadline(time.Time{})

	c.packets = NewPackets(c.netConn)

	return c, nil
}

func (c *conn) NextPacket() ([]byte, error) {
	return c.packets.Next()
}

func (c *conn) Cleanup() {
	if c.netConn != nil {
		c.netConn.Close()
		c.netConn = nil
	}
}

func (c *conn) Closed() bool {
	return c.netConn == nil
}

func (c *conn) Read() error {
	var err error
	var data []byte

	if data, err = c.packets.Next(); err != nil {
		return err
	}

	if err = c.greeting.UnPack(data); err != nil {
		return err
	}

	return nil
}
