package protocol

import (
	"fmt"
	"math"
	"mysql-scanner/util"
)

const errMalformedPacket = "malformed packet: %v"

/**
1              [0a] protocol version
string[NUL]    server version
4              connection id
string[8]      auth-plugin-data-part-1
1              [00] filler
2              capability flags (lower 2 bytes)
  if more data in the packet:
1              character set
2              status flags
2              capability flags (upper 2 bytes)
  if capabilities & CLIENT_PLUGIN_AUTH {
1              length of auth-plugin-data
  } else {
1              [00]
  }
string[10]     reserved (all [00])
  if capabilities & CLIENT_SECURE_CONNECTION {
string[$len]   auth-plugin-data-part-2 ($len=MAX(13, length of auth-plugin-data - 8))
  if capabilities & CLIENT_PLUGIN_AUTH {
string[NUL]    auth-plugin name
  }
*/
type Greeting struct {
	ProtoVersion uint8
	SvrVersion   string
	ConnId       uint32
	AuthData     []byte
	Filler       byte
	Capability   CapabilityFlag
	Status       uint16
	CharacterSet uint8
	AuthName     string
}

// UnPack used to unpack the greeting packet.
func (g *Greeting) UnPack(payload []byte) error {
	var err error
	buf := util.ReadBuffer(payload)

	// 1: [0a] protocol version
	if g.ProtoVersion, err = buf.ReadU8(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting protocol-version failed")
	}

	// string[NUL]: server version
	if g.SvrVersion, err = buf.ReadStringNUL(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting server-version failed")
	}

	// 4: connection id
	if g.ConnId, err = buf.ReadU32(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting connection-id failed")
	}

	// string[8]: auth-plugin-data-part-1
	var authData1 []byte
	if authData1, err = buf.ReadBytes(8); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting auth-plugin-data-part-1 failed")
	}
	g.AuthData = authData1

	// 1: [00] filler
	if err = buf.ReadZero(1); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting filler failed")
	}

	// 2: capability flags (lower 2 bytes)
	var capLower uint16
	if capLower, err = buf.ReadU16(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting capability-flags failed")
	}

	// 1: character set
	if g.CharacterSet, err = buf.ReadU8(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting charset failed")
	}

	// 2: status flags
	if g.Status, err = buf.ReadU16(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting status-flags failed")
	}

	// 2: capability flags (upper 2 bytes)
	var capUpper uint16
	if capUpper, err = buf.ReadU16(); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting capability-flags-upper failed")
	}
	g.Capability = (CapabilityFlag(capUpper) << 16) | (CapabilityFlag(capLower))

	// 1: length of auth-plugin-data-part-1
	var authLen byte
	if g.Capability.Has(clientPluginAuth) {
		if authLen, err = buf.ReadU8(); err != nil {
			return fmt.Errorf(errMalformedPacket, "extracting greeting auth-plugin-data length failed")
		}
	} else {
		if err = buf.ReadZero(1); err != nil {
			return fmt.Errorf(errMalformedPacket, "extracting greeting zero failed")
		}
	}

	// string[10]: reserved (all [00])
	if err = buf.ReadZero(10); err != nil {
		return fmt.Errorf(errMalformedPacket, "extracting greeting reserved failed")
	}

	//// string[$len]: auth-plugin-data-part-2 ($len=MAX(13, length of auth-plugin-data - 8))
	if g.Capability.Has(clientSecureConn) {
		toRead := int(math.Max(13, float64(authLen)-8))
		var authData2 []byte
		if authData2, err = buf.ReadBytes(toRead); err != nil {
			return fmt.Errorf(errMalformedPacket, "extracting greeting salt2 failed")
		}
		copy(g.AuthData[8:], authData2[:toRead-1])
	}

	// string[NUL]    auth-plugin name
	if g.Capability.Has(clientPluginAuth) {
		if g.AuthName, err = buf.ReadStringNUL(); err != nil {
			return fmt.Errorf(errMalformedPacket, "extracting greeting auth-plugin-name failed")
		}
	}
	return nil
}

func (g *Greeting) String() string {
	out := "MySql Details\n"
	out += fmt.Sprintf(" Protocol Version: %d\n", g.ProtoVersion)
	out += fmt.Sprintf(" Server Version: %s\n", g.SvrVersion)
	out += fmt.Sprintf(" Connection ID: %d\n", g.ConnId)
	out += fmt.Sprintf(" Character Set: %s\n", CharacterSetMap[g.CharacterSet])
	out += fmt.Sprintf(" Status: %d\n", g.Status)
	out += fmt.Sprintf(" Auth Name: %s\n", g.AuthName)
	out += fmt.Sprintf(" Auth Data: %d\n", g.AuthData)
	out += fmt.Sprintf(" Capabilities:\n %s", g.Capability)
	return out
}
