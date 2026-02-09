package database

import (
	"backend/pkg/config"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type PostgreDB struct {
	// db object map [dbIdx, *sql.DB]
	dbSet        *sync.Map
	version      string
	defaultDbIdx int
}

func NewPostgreDB(multiLogger *zap.Logger, config config.Config) IDatabase {

	dbc := &PostgreDB{
		dbSet:        new(sync.Map),
		version:      "",
		defaultDbIdx: config.GetDatabase().DefaultDbIdx,
	}

	deiverName := config.GetDatabase().DriverName

	multiLogger.Debug("Database connect info", zap.String("driver_name", deiverName),
		zap.Strings("conn_info", config.GetDatabase().ConnInfo))

	for idx, connInfo := range config.GetDatabase().ConnInfo {
		dbc.CreateDatabaseObject(multiLogger, idx, deiverName, connInfo)
	}

	// defaultDbIdx := config.GetDatabase().DefaultDbIdx
	// use dafault db index from config
	if db := dbc.GetDefaultDB(); db == nil {
		multiLogger.Fatal("Get database failed error", zap.Int("index", 0))
	} else {
		err := db.QueryRow("select version()").Scan(&dbc.version)
		if err != nil {
			multiLogger.Fatal("QueryRow database version error", zap.Error(err))
		}
	}

	return dbc
}

func (p *PostgreDB) DBSet() *sync.Map { return p.dbSet }
func (p *PostgreDB) Version() string  { return p.version }

func (p *PostgreDB) GetDB(idx int) *sql.DB {
	val, ok := p.dbSet.Load(idx)
	if ok {
		db, ok := val.(*sql.DB)
		if ok {
			return db
		}
	}

	return nil
}

func (p *PostgreDB) GetDefaultDB() *sql.DB {
	return p.GetDB(p.defaultDbIdx)
}

// 創建一個 db 物件
func (p *PostgreDB) CreateDatabaseObject(logger *zap.Logger, dbIdx int, driverName, connInfo string) {

	p.CheckDatabaseExist(logger, dbIdx, driverName, connInfo)

	db, err := sql.Open(driverName, connInfo)
	if err != nil {
		logger.Fatal("Database create error", zap.Error(err))
	} else {
		p.dbSet.Store(dbIdx, db)
	}
}

/* 檢查指定資料庫是否存在, 如果不存在就創建
 */
func (p *PostgreDB) CheckDatabaseExist(logger *zap.Logger, dbIdx int, driverName, connInfo string) {

	// parse connect db connInfo and check database is exist, if not exist create it.
	connParams := strings.Split(connInfo, " ")

	log.Println(connParams)

	var dbname string
	connParseParamsSlice := make([]string, 0)
	for _, v := range connParams {
		tmp := strings.Split(v, "=")
		if tmp[0] != "dbname" {
			connParseParamsSlice = append(connParseParamsSlice, v)
		} else {
			dbname = tmp[1]
		}
	}

	afterConnInfo := strings.Join(connParseParamsSlice, " ")

	// "host=127.0.0.1 port=5432 user=postgres dbname=dcc_chat password=123456 sslmode=disable"
	db, err := sql.Open(driverName, afterConnInfo)
	if err != nil {
		logger.Fatal("Failed to open database", zap.Error(err))
	}

	var dBExists bool
	if err = db.QueryRow("SELECT EXISTS (SELECT 1 from pg_database WHERE datname = $1)", dbname).Scan(&dBExists); err != nil {
		// var pgErr *pgconn.PgError
		if err != nil {
			dBExists = false
		} else {
			db.Close()
			logger.Fatal("Failed to check if db exists", zap.String("db", dbname), zap.Error(err))
		}
	}

	if !dBExists {
		// Database does not exist, create it
		logger.Info("Creating new database", zap.String("name", dbname))

		if _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %q", dbname)); err != nil {
			db.Close()
			logger.Fatal("Failed to create database", zap.Error(err))
		}
		db.Close()

		db, err = sql.Open(driverName, connInfo)
		if err != nil {
			db.Close()
			logger.Fatal("Failed to open database", zap.Error(err))
		}
	}

	// Get database version
	var dbVersion string
	if err = db.QueryRow("SELECT version()").Scan(&dbVersion); err != nil {
		db.Close()
		logger.Fatal("Error querying database version", zap.Error(err))
	}

	log.Println("Database information", zap.String("version", dbVersion))

	if err = db.Ping(); err != nil {
		db.Close()
		logger.Fatal("Error pinging database", zap.Error(err))
	}

	db.Close()
}
