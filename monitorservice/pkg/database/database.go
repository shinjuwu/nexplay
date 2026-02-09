package database

import (
	"database/sql"
	"monitorservice/pkg/config"
	"sync"

	"go.uber.org/zap"
)

type IDatabase interface {
	DBSet() *sync.Map
	Version() string
	GetDB(int) *sql.DB
	GetDefaultDB() *sql.DB
	CreateDatabaseObject(int, string, string) error
}

const (
	DRIVER_POSTGRESQL = "postgres"
	DRIVER_MYSQL      = "mysql"
	DRIVER_SQLITE     = "sqlite"
)

func DBConnectDispatcher(multiLogger *zap.Logger, config config.Config) IDatabase {

	switch config.GetDatabase().DriverName {
	case DRIVER_POSTGRESQL:
		return NewPostgreDB(multiLogger, config)
	case DRIVER_MYSQL:
		// TODO:
	case DRIVER_SQLITE:
		// TODO:
	default:
		multiLogger.Debug("Unknow database type", zap.String("driver_name", config.GetDatabase().DriverName))
	}
	return nil
}
