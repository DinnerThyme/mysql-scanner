package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufferGetters(t *testing.T) {
	data := []byte{152, 21}
	reader := ReadBuffer(data)

	assert.Equal(t, []byte{152, 21}, reader.Data())
	assert.Equal(t, 2, reader.Length())
	assert.Equal(t, 0, reader.Seek())
}

func TestBufferRead8(t *testing.T) {
	data := []byte{152, 21}
	reader := ReadBuffer(data)
	value, _ := reader.ReadU8()
	assert.Equal(t, uint8(152), value)

	data = []byte{}
	reader = ReadBuffer(data)
	value, err := reader.ReadU8()
	assert.Error(t, err)
}

func TestBufferRead16(t *testing.T) {
	data := []byte{152, 21, 83, 1}
	reader := ReadBuffer(data)
	value, _ := reader.ReadU16()
	assert.Equal(t, uint16(5528), value)

	data = []byte{}
	reader = ReadBuffer(data)
	value, err := reader.ReadU16()
	assert.Error(t, err)
}

func TestBufferRead32(t *testing.T) {
	data := []byte{152, 21, 83, 1, 97, 98, 99, 0, 0}
	reader := ReadBuffer(data)
	value, _ := reader.ReadU32()
	assert.Equal(t, uint32(22222232), value)

	data = []byte{}
	reader = ReadBuffer(data)
	value, err := reader.ReadU32()
	assert.Error(t, err)
}
