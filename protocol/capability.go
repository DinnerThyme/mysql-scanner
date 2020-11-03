package protocol

import (
	"fmt"
	"strings"
)

// https://dev.mysql.com/doc/internals/en/capability-flags.html
type CapabilityFlag uint32

func (r CapabilityFlag) Has(flag CapabilityFlag) bool {
	return r&flag != 0
}

func (r CapabilityFlag) String() string {
	var names []string
	for i := uint64(1); i <= uint64(1)<<31; i = i << 1 {
		if r.Has(CapabilityFlag(i)) {
			if name, ok := flags[CapabilityFlag(i)]; ok {
				names = append(names, fmt.Sprintf("\t- 0x%08x - %s", i, name))
			}

		}
	}
	return strings.Join(names, "\n")
}

const (
	clientLongPassword CapabilityFlag = 1 << iota
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConn
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)

var flags = map[CapabilityFlag]string{
	clientLongPassword:               "clientLongPassword",
	clientFoundRows:                  "clientFoundRows",
	clientLongFlag:                   "clientLongFlag",
	clientConnectWithDB:              "clientConnectWithDB",
	clientNoSchema:                   "clientNoSchema",
	clientCompress:                   "clientCompress",
	clientODBC:                       "clientODBC",
	clientLocalFiles:                 "clientLocalFiles",
	clientIgnoreSpace:                "clientIgnoreSpace",
	clientProtocol41:                 "clientProtocol41",
	clientInteractive:                "clientInteractive",
	clientSSL:                        "clientSSL",
	clientIgnoreSIGPIPE:              "clientIgnoreSIGPIPE",
	clientTransactions:               "clientTransactions",
	clientReserved:                   "clientReserved",
	clientSecureConn:                 "clientSecureConn",
	clientMultiStatements:            "clientMultiStatements",
	clientMultiResults:               "clientMultiResults",
	clientPSMultiResults:             "clientPSMultiResults",
	clientPluginAuth:                 "clientPluginAuth",
	clientConnectAttrs:               "clientConnectAttrs",
	clientPluginAuthLenEncClientData: "clientPluginAuthLenEncClientData",
	clientCanHandleExpiredPasswords:  "clientCanHandleExpiredPasswords",
	clientSessionTrack:               "clientSessionTrack",
	clientDeprecateEOF:               "clientDeprecateEOF",
}
