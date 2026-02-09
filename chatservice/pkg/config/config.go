package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config interface {
	GetName() string
	GetDataDir() string
	GetShutdownGraceSec() int
	GetApp() *AppConfig
	GetLogger() *LoggerConfig
	GetSocket() *SocketConfig
	GetDatabase() *DatabaseConfig
}

func ReadYamlFile(logger *zap.Logger) Config {
	c := NewConfig()

	args := os.Args[1:]

	filename := ""
	if len(args) > 1 { // for systemctl
		if args[0] == "--config" {
			filename = args[1]
			yamlFile, err := ioutil.ReadFile(filename)
			if err != nil {
				logger.Fatal("yamlFile.Get err: ", zap.Error(err))
			}

			err = yaml.Unmarshal(yamlFile, &c)
			if err != nil {
				logger.Fatal("Unmarshal: ", zap.Error(err))
			}
		} else {
			logger.Fatal("yamlFile.Get err #")
		}
	} else {
		filename, _ = filepath.Abs("./config.yml")

		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			logger.Fatal("yamlFile.Get err: ", zap.Error(err))
		}

		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			logger.Fatal("Unmarshal: ", zap.Error(err))
		}
	}

	return c
}
