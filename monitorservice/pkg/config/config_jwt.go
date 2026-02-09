package config

// JwtConfig is configuration application connect information.
type JwtConfig struct {
	EncryptString string `yaml:"encrypt_string" json:"encrypt_string"`
	Issuer        string `yaml:"issuer" json:"issuer"`
	ExpireTimeMin int    `yaml:"expire_time_min" json:"expire_time_min"`
}

// NewJwtConfig creates a new JwtConfig struct.
func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		EncryptString: "#oHBda8jidas32R3G$Op",
		Issuer:        "monitorservice",
		ExpireTimeMin: 10080, // 7 days
	}
}
