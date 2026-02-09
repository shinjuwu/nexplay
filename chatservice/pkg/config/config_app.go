package config

// AppConfig is configuration application connect information.
type AppConfig struct {
	Env           string `yaml:"env" json:"env"`
	Addr          int    `yaml:"addr" json:"addr"`
	UseMultipoint bool   `yaml:"use_multipoint" json:"use_multipoint"`
	LoadFront     bool   `yaml:"load_front" json:"load_front"`
	LoadSwagger   bool   `yaml:"load_swagger" json:"load_swagger"`
}

// NewAppConfig creates a new AppConfig struct.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Env:           "debug",
		Addr:          8080,
		UseMultipoint: false,
		LoadFront:     false,
		LoadSwagger:   false,
	}
}
