package config

// DatabaseConfig is configuration relevant to the Database storage.
type DatabaseConfig struct {
	DriverName   string   `yaml:"driver_name" json:"driver_name"`
	DefaultDbIdx int      `yaml:"default_db_idx" json:"default_db_idx"`
	ConnInfo     []string `yaml:"conn_info" json:"conn_info"`
}

// NewDatabaseConfig creates a new DatabaseConfig struct.
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DriverName:   "postgres",
		DefaultDbIdx: 0,
		ConnInfo:     []string{"host=127.0.0.1 port=5432 user=postgres dbname=postgres password= sslmode=disable"},
	}
}
