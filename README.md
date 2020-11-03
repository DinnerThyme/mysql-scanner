# MySql Scanner

CLI Tool that performs a handshake with a MySql Server and outputs the information
gained during the handshake process.

### Setup

To build the executable, run:

```shell script
make build
```

### Usage

To run the executable, run:

```shell script
mysql-scanner -h [hostname] -p [port]
```

CLI Arguments

```shell script
  -h string
    	network address to connect to (default "127.0.0.1")
  -p int
    	port to connect to (default 3306)
```

### Output

The application outputs some basic information about the MySql service that it connected with.
Specifically it displays:

- Protocol version (Always 10 in this implementation)
- Server version
- Connection Id
- Character Set
- Current Status
- Capability Flags
- Auth Data
- Auth Name

Example output:

```shell script
 MySql Details
 Protocol Version: 10
 Server Version: 8.0.22
 Connection ID: 84
 Character Set: unknown
 Status: 2
 Auth Name: caching_sha2_password
 Auth Data: [120 69 73 63 42 37 37 26]
 Capabilities:
 	- 0x00000001 - clientLongPassword
	- 0x00000002 - clientFoundRows
	- 0x00000004 - clientLongFlag
	- 0x00000008 - clientConnectWithDB
	- 0x00000010 - clientNoSchema
	- 0x00000020 - clientCompress
	- 0x00000040 - clientODBC
	- 0x00000080 - clientLocalFiles
	- 0x00000100 - clientIgnoreSpace
	- 0x00000200 - clientProtocol41
	- 0x00000400 - clientInteractive
	- 0x00000800 - clientSSL
	- 0x00001000 - clientIgnoreSIGPIPE
	- 0x00002000 - clientTransactions
	- 0x00004000 - clientReserved
	- 0x00008000 - clientSecureConn
	- 0x00010000 - clientMultiStatements
	- 0x00020000 - clientMultiResults
	- 0x00040000 - clientPSMultiResults
	- 0x00080000 - clientPluginAuth
	- 0x00100000 - clientConnectAttrs
	- 0x00200000 - clientPluginAuthLenEncClientData
	- 0x00400000 - clientCanHandleExpiredPasswords
	- 0x00800000 - clientSessionTrack
	- 0x01000000 - clientDeprecateEOF
```

### Testing

I did not use any particular testing framework, just the built in Golang testing tool along with a package to simplify assertions.

I have provided a Make command to simplify testing:

```shell script
make test
```
