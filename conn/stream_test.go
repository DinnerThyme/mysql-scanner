package conn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStream(t *testing.T) {
	strm := NewStream(&MockConn{}, maxPacketSize)
	assert.NotNil(t, strm)
}

func TestStream_Read(t *testing.T) {
	strm := NewStream(&MockConn{}, maxPacketSize)
	pkt, err := strm.Read()
	assert.NoError(t, err)
	assert.Equal(t,  uint8(1), pkt.SequenceID)
	assert.Equal(t,  74, len(pkt.Data))
}
