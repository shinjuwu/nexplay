package config

// DatabaseConfig is configuration relevant to the Database storage.
type DatabaseConfig struct {
	DriverName       string   `yaml:"driver_name" json:"driver_name"`
	DefaultDbIdx     int      `yaml:"default_db_idx" json:"default_db_idx"`
	ConnInfo         []string `yaml:"conn_info" json:"conn_info"`
	MigrateSourceUrl string   `yaml:"migrate_source_url" json:"migrate_source_url"`
}

// NewDatabaseConfig creates a new DatabaseConfig struct.
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DriverName:       "postgres",
		DefaultDbIdx:     0,
		ConnInfo:         []string{"host=34.87.61.82 port=5432 user=postgres dbname=dcc_order password=postgres01 sslmode=disable"},
		MigrateSourceUrl: "file://db/migrations",
	}
}
