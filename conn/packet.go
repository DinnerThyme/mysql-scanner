package conn

import "net"

const maxPacketSize = (1<<24 - 1) // (16MB - 1）

//https://dev.mysql.com/doc/internals/en/mysql-packet.html
type Packet struct {
	SequenceID byte
	Data       []byte
}

type Packets struct {
	seq    uint8
	stream *Stream
}

func NewPackets(c net.Conn) *Packets {
	return &Packets{
		stream: NewStream(c, maxPacketSize),
	}
}

func (p *Packets) Next() ([]byte, error) {
	pkt, err := p.stream.Read()
	if err != nil {
		return nil, err
	}

	p.seq++
	return pkt.Data, nil
}
