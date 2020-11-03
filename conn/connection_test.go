package conn

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

// MockConn used to mock a net.Conn for testing purposes.
//type MockConn struct {
//	laddr  net.Addr
//	raddr  net.Addr
//	data   []byte
//	closed bool
//	read   int
//}

const greeting = "\x0A\x38\x2E\x30\x2E\x32\x32\x00\x3B\x00\x00\x00\x65\x05\x60\x11\x12\x0C\x0A\x17\x00\xFF\xFF\xFF\x02\x00\xFF\xC7\x15\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x69\x60\x35\x05\x4C\x51\x6D\x2F\x5C\x1D\x38\x7B\x00\x63\x61\x63\x68\x69\x6E\x67\x5F\x73\x68\x61\x32\x5F\x70\x61\x73\x73\x77\x6F\x72\x64\x00"

func TestConnect(t *testing.T) {
	addr := "127.0.0.200:12345"
	lstn, err := net.Listen("tcp", addr)
	assert.NoError(t, err)

	var svr net.Conn
	go func() {
		defer lstn.Close()
		svr, err = lstn.Accept()
		assert.NoError(t, err)

		_, err := svr.Write([]byte("\x01"))
		assert.NoError(t, err)
	}()

	_, err = Connect(addr, time.Duration(1) * time.Second)
	assert.NoError(t, err)
}

type MockConn struct {
	net.Conn
}

func (m *MockConn) Read(b []byte) (int, error) {
	// SeqId: 1, Length: 74
	n := copy(b, []byte{74, 0, 0, 1})
	n = copy(b[4:], greeting)
	return n, nil
}
