package conn

import (
	"bufio"
	"io"
	"net"
)

const bufferSize = 128

type Stream struct {
	pktMaxSize int
	header     []byte
	reader     *bufio.Reader
	writer     *bufio.Writer
}

// NewStream creates a new stream.
func NewStream(conn net.Conn, pktMaxSize int) *Stream {
	return &Stream{
		pktMaxSize: pktMaxSize,
		header:     []byte{0, 0, 0, 0},
		reader:     bufio.NewReaderSize(conn, bufferSize),
		writer:     bufio.NewWriterSize(conn, bufferSize),
	}
}

// Read reads the next packet from the reader
// The returned pkt.Data is only guaranteed to be valid until the next read
func (s *Stream) Read() (*Packet, error) {
	// Header.
	if _, err := io.ReadFull(s.reader, s.header); err != nil {
		return nil, err
	}

	// Length.
	pkt := &Packet{}
	pkt.SequenceID = s.header[3]
	length := int(uint32(s.header[0]) | uint32(s.header[1])<<8 | uint32(s.header[2])<<16)
	if length == 0 {
		return pkt, nil
	}

	// Data.
	data := make([]byte, length)
	if _, err := io.ReadFull(s.reader, data); err != nil {
		return nil, err
	}
	pkt.Data = data

	// Single packet.
	if length < s.pktMaxSize {
		return pkt, nil
	}

	// There is more than one packet, read them all.
	next, err := s.Read()
	if err != nil {
		return nil, err
	}
	pkt.SequenceID = next.SequenceID
	pkt.Data = append(pkt.Data, next.Data...)
	return pkt, nil
}
