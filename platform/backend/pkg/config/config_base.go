package config

import (
	"log"
	"os"
	"path/filepath"
)

type config struct {
	Name             string          `yaml:"name" json:"name" usage:"Server’s node name - must be unique."`
	Config           []string        `yaml:"config" json:"config" usage:"The absolute file path to configuration YAML file."`
	ShutdownGraceSec int             `yaml:"shutdown_grace_sec" json:"shutdown_grace_sec" usage:"Maximum number of seconds to wait for the server to complete work before shutting down. Default is 0 seconds. If 0 the server will shut down immediately when it receives a termination signal."`
	Datadir          string          `yaml:"data_dir" json:"data_dir" usage:"An absolute path to a writeable folder where server will store its data."`
	App              *AppConfig      `yaml:"app" json:"app" usage:"app setting."`
	Jwt              *JwtConfig      `yaml:"jwt" json:"jwt" usage:"Jwt setting."`
	Logger           *LoggerConfig   `yaml:"logger" json:"logger" usage:"Logger levels and output."`
	Database         *DatabaseConfig `yaml:"database" json:"database" usage:"Database connection settings."`
	Redis            *RedisConfig    `yaml:"redis" json:"redis" usage:"Redis connection setting."`
	Captcha          *CaptchaConfig  `yaml:"captcha" json:"captcha" usage:"Captcha create and verify setting."`
	Backend          *BackendConfig  `yaml:"backend" json:"backend" usage:"Backend server setting."`
}

func NewConfig() *config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory.", err)
	}
	return &config{
		Name:             "defaultname",
		Datadir:          filepath.Join(cwd, "config"),
		ShutdownGraceSec: 0,
		App:              NewAppConfig(),
		Jwt:              NewJwtConfig(),
		Logger:           NewLoggerConfig(),
		Database:         NewDatabaseConfig(),
		Redis:            NewRedisConfig(),
		Captcha:          NewCaptchaConfig(),
		Backend:          NewBackendConfig(),
	}
}

func CheckConfigArgs(config *config) error {
	return nil
}

func (c *config) GetName() string {
	return c.Name
}

func (c *config) GetDataDir() string {
	return c.Datadir
}

func (c *config) GetShutdownGraceSec() int {
	return c.ShutdownGraceSec
}

func (c *config) GetApp() *AppConfig {
	return c.App
}

func (c *config) GetLogger() *LoggerConfig {
	return c.Logger
}

func (c *config) GetDatabase() *DatabaseConfig {
	return c.Database
}

func (c *config) GetRedis() *RedisConfig {
	return c.Redis
}

func (c *config) GetJwt() *JwtConfig {
	return c.Jwt
}

func (c *config) GetCaptcha() *CaptchaConfig {
	return c.Captcha
}

func (c *config) GetBackend() *BackendConfig {
	return c.Backend
}
