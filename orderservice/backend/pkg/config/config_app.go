package config

// AppConfig is configuration application connect information.
type AppConfig struct {
	Env                string `yaml:"env" json:"env"`
	Addr               int    `yaml:"addr" json:"addr"`
	AccManageHeader    string `yaml:"acc_manage_header" json:"acc_manage_header"`
	AccNormalHeader    string `yaml:"acc_normal_header" json:"acc_normal_header"`
	CheckHeader        bool   `yaml:"check_header" json:"check_header"`
	UseMultipoint      bool   `yaml:"use_multipoint" json:"use_multipoint"`
	LoadFront          bool   `yaml:"load_front" json:"load_front"`
	LoadSwagger        bool   `yaml:"load_swagger" json:"load_swagger"`
	LoadDBMigration    bool   `yaml:"load_db_migration" json:"load_db_migration"`
	LoadApiChannel     bool   `yaml:"load_api_channel" json:"load_api_channel"`
	LoadBackendChannel bool   `yaml:"load_backend_channel" json:"load_backend_channel"`
	LoadJobScheduler   bool   `yaml:"load_job_scheduler" json:"load_job_scheduler"`
}

// NewAppConfig creates a new AppConfig struct.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Env:                "debug",
		Addr:               8080,
		AccManageHeader:    "8887",
		AccNormalHeader:    "8888",
		CheckHeader:        false,
		UseMultipoint:      false,
		LoadFront:          false,
		LoadSwagger:        false,
		LoadDBMigration:    false,
		LoadApiChannel:     false,
		LoadBackendChannel: false,
		LoadJobScheduler:   false,
	}
}
