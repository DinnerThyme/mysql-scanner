package conn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPackets(t *testing.T) {
	pkt := NewPackets(&MockConn{})
	assert.NotNil(t, pkt)
}

//func TestPackets_Next(t *testing.T) {
//	pkt := NewPackets(&MockConn{})
//	bytes, err := pkt.Next()
//	assert.NoError(t, err)
//	assert.Equal(t, []uint8(greeting), bytes)
//}
