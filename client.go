package mysql_scanner

import (
	"fmt"
	"mysql-scanner/conn"
	"mysql-scanner/protocol"
	"time"
)

type Client struct {
	addr string
	port int
}

func (c *Client) Probe(hostname string, port int) (*protocol.Greeting, error) {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	timeout := time.Duration(5) * time.Second
	cnn, err := conn.Connect(addr, timeout)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			cnn.Cleanup()
		}
	}()

	if err := cnn.Read(); err != nil {
		return nil, err
	}

	return cnn.Greeting(), nil
}
