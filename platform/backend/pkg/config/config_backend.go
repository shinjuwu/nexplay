package config

// BackendConfig is configuration relevant to the backend server.
type BackendConfig struct {
	Address string `yaml:"address" json:"address"`
}

// NewBackendConfig creates a new BackendConfig struct.
func NewBackendConfig() *BackendConfig {
	return &BackendConfig{
		Address: "",
	}
}
