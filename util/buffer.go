package util

import (
	"bytes"
	"encoding/binary"
	"io"
)

var (
	// ErrIOEOF used for io.EOF.
	ErrIOEOF = io.EOF
)

// Buffer represents the buffer tuple.
type Buffer struct {
	pos  int
	seek int
	cap  int
	buf  []byte
}

func ReadBuffer(b []byte) *Buffer {
	return &Buffer{
		buf: b,
		pos: len(b),
	}
}

func (b *Buffer) Reset(data []byte) {
	b.buf = data
	b.pos = len(data)
	b.seek = 0
}

func (b *Buffer) Data() []byte {
	return b.buf[:b.pos]
}

func (b *Buffer) Length() int {
	return b.pos
}

func (b *Buffer) Seek() int {
	return b.seek
}

func (b *Buffer) AtEOF(bytes int) bool {
	return (b.seek + bytes) > b.pos
}

func (b *Buffer) ReadU8() (v uint8, err error) {
	if b.AtEOF(1) {
		err = ErrIOEOF
		return
	}

	v = b.buf[b.seek]
	b.seek++
	return
}

func (b *Buffer) ReadU16() (v uint16, err error) {
	if b.AtEOF(2) {
		err = ErrIOEOF
		return
	}

	v = binary.LittleEndian.Uint16(b.buf[b.seek : b.seek+2])
	b.seek += 2
	return
}

func (b *Buffer) ReadU32() (v uint32, err error) {
	if b.AtEOF(4) {
		err = ErrIOEOF
		return
	}
	v = binary.LittleEndian.Uint32(b.buf[b.seek : b.seek+4])
	b.seek += 4
	return
}

// ReadEOF used to read EOF.
func (b *Buffer) ReadEOF(n int) (err error) {
	return b.ReadZero(n)
}

// ReadZero used to read zero.
func (b *Buffer) ReadZero(n int) (err error) {
	if b.AtEOF(n) {
		err = ErrIOEOF
		return
	}
	b.seek += n
	return
}

// ReadStringNUL reads until the first NUL in the buffer
// returning a string containing the data up to and not including the NUL
func (b *Buffer) ReadStringNUL() (s string, err error) {
	var v []byte
	if v, err = b.readBytesWithToken(0x00); err != nil {
		return
	}
	s = string(v)
	return
}

func (b *Buffer) readBytesWithToken(token uint8) (v []byte, err error) {
	i := bytes.IndexByte(b.buf[b.seek:], token)
	end := b.seek + i + 1
	if i < 0 {
		b.seek = len(b.buf)
		err = ErrIOEOF
		return
	}
	v = b.buf[b.seek : end-1]
	b.seek = end
	return
}

// ReadBytes used to read bytes.
func (b *Buffer) ReadBytes(n int) (v []byte, err error) {
	if n == 0 {
		return nil, nil
	}

	if (b.seek + n) > b.pos {
		err = ErrIOEOF
		return
	}

	v = b.buf[b.seek:(b.seek + n)]
	b.seek += n
	return
}
