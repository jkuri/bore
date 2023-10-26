package server

// Structure contains data of what address/port
// we should bind forwarded-tcpip connections
type bindInfo struct {
	Bound string
	Port  uint32
	Addr  string
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

type idRequestPayload struct {
	ID string
}

type clientResponse struct {
	id     string
	port   uint32
	domain string
}
