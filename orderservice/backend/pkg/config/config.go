package config

import (
	"io/ioutil"
	"path/filepath"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config interface {
	GetName() string
	GetDataDir() string
	GetShutdownGraceSec() int
	GetApp() *AppConfig
	GetJwt() *JwtConfig
	GetLogger() *LoggerConfig
	GetDatabase() *DatabaseConfig
	GetRedis() *RedisConfig
	GetCaptcha() *CaptchaConfig
}

func ReadYamlFile(logger *zap.Logger) Config {
	c := NewConfig()

	var filename string

	filename, _ = filepath.Abs("./config.yml")

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Sugar().Fatalf("yamlFile.Get err  #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		logger.Fatal("Unmarshal: %v")
	}

	return c
}
