package client

// Config holds configuration data.
type Config struct {
	RemoteServer string
	RemotePort   int
	LocalServer  string
	LocalPort    int
	BindPort     int
	KeepAlive    bool
}
