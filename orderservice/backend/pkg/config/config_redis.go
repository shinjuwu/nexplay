package config

// DatabaseConfig is configuration relevant to the Database storage.
type RedisConfig struct {
	Address  string `yaml:"address" json:"address"`
	Password string `yaml:"password" json:"password"`
	DbIndex  []int  `yaml:"db_index" json:"db_index"`
}

// NewDatabaseConfig creates a new DatabaseConfig struct.
func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Address:  "127.0.0.1:6379",
		Password: "",
		DbIndex:  []int{14},
	}
}
