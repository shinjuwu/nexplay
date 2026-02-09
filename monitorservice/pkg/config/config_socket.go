package config

// LoggerConfig is configuration relevant to logging levels and output.
type SocketConfig struct {
	ServerKey    string `yaml:"server_key" json:"server_key" usage:"Server key to use to establish a connection to the server."`
	Port         int    `yaml:"port" json:"port" usage:"The port for accepting connections from the client for the given interface(s), address(es), and protocol(s). Default 7350."`
	Address      string `yaml:"address" json:"address" usage:"The IP address of the interface to listen for client traffic on. Default listen on all available addresses/interfaces."`
	SingleSocket bool   `yaml:"single_socket" json:"single_socket" usage:"Only allow one socket per user. Older sessions are disconnected. Default false."`
}

// NewLoggerConfig creates a new LoggerConfig struct.
func NewLSocketConfig() *SocketConfig {
	return &SocketConfig{
		ServerKey:    "defaultkey",
		Port:         17782,
		Address:      "",
		SingleSocket: false,
	}
}
