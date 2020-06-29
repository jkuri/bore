package server

// Structure contains data of what address/port
// we should bind forwarded-tcpip connections
type bindInfo struct {
	Bound string
	Port  uint32
	Addr  string
}

// Information parsed from the authorized_keys file
type deviceInfo struct {
	LocalPorts  string
	RemotePorts string
	Comment     string
}

// RFC4254 7.2
type directTCPPayload struct {
	Addr       string // Connect to
	Port       uint32
	OriginAddr string
	OriginPort uint32
}

type forwardedTCPPayload struct {
	Addr       string // Connected to
	Port       uint32
	OriginAddr string
	OriginPort uint32
}

type tcpIPForwardPayload struct {
	Addr string
	Port uint32
}

type tcpIPForwardPayloadReply struct {
	Port uint32
}

type tcpIPForwardCancelPayload struct {
	Addr string
	Port uint32
}
