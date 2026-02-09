package database

import (
	"backend/pkg/config"
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrNoChange   = migrate.ErrNoChange
	ErrNilVersion = migrate.ErrNoChange
)

type SMigrate struct {
	sourceURL   string
	databaseURL string
	m           *migrate.Migrate
}

func NewSMigration(config config.Config) (*SMigrate, error) {
	//"host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable"    # for local
	connParams := strings.Split(config.GetDatabase().ConnInfo[0], " ")
	connParseParamsSlice := make(map[string]string, 0)
	for _, v := range connParams {
		tmp := strings.Split(v, "=")
		connParseParamsSlice[tmp[0]] = tmp[1]
	}

	driverName := config.GetDatabase().DriverName
	dbName := connParseParamsSlice["dbname"]
	username := connParseParamsSlice["user"]
	password := connParseParamsSlice["password"]
	hostname := connParseParamsSlice["host"]
	port := connParseParamsSlice["port"]

	// check databse be created
	// "postgres://username:password@hostname:port/?sslmode=disable"
	createDatabaseURL := fmt.Sprintf(
		"%s://%s:%s@%s:%s/?sslmode=disable",
		driverName,
		username,
		password,
		hostname,
		port)

	db, err := sql.Open(driverName, createDatabaseURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT datname FROM pg_catalog.pg_database WHERE datname = $1;", dbName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database already exists!")
	} else {
		createQuery := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		_, err = db.Exec(createQuery)
		if err != nil {
			fmt.Println("Failed to create database:", err.Error())
			return nil, err
		}
		fmt.Println("Database created successfully!")
	}

	sourceURL := config.GetDatabase().MigrateSourceUrl
	//"postgres://postgres:123456@localhost:5432/example?sslmode=disable")
	databaseURL := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		driverName,
		username,
		password,
		hostname,
		port,
		dbName)

	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		return nil, err
	}

	res := &SMigrate{
		sourceURL:   sourceURL,
		databaseURL: databaseURL,
		m:           m,
	}

	return res, err
}

func (p *SMigrate) SourceURL() string {
	return p.sourceURL
}

func (p *SMigrate) DatabaseURL() string {
	return p.databaseURL
}

// Down looks at the currently active migration version
// and will migrate all the way down (applying all down migrations).
// return version, is dirty, error is sql select return
func (p *SMigrate) Up() (uint, bool, error) {
	err := p.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return 0, false, err
	}
	ver, dirty, err := p.m.Version()
	return ver, dirty, err
}

// Up looks at the currently active migration version
// and will migrate all the way up (applying all up migrations).
// return version, is dirty, error is sql select return
func (p *SMigrate) Down() (uint, bool, error) {
	err := p.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return 0, false, err
	}
	ver, dirty, err := p.m.Version()
	return ver, dirty, err
}

// get migration veriosn and state of migrate.
// return version, is dirty, error is sql select return
func (p *SMigrate) Version() (uint, bool, error) {
	ver, dirty, err := p.m.Version()
	if err == migrate.ErrNilVersion {
		return ver, dirty, nil
	} else {
		return ver, dirty, err
	}
}

// 設定 migration 狀態為 true
// return error is sql select return
func (p *SMigrate) Force(currentlyVersion int) error {
	return p.m.Force(currentlyVersion)
}

// The currently active migration version. It will migrate up to n.
// return version, is dirty, error is sql select return
func (p *SMigrate) UpTo(currentlyVersion int) (uint, bool, error) {
	err := p.m.Steps(currentlyVersion)
	if err != nil {
		return 0, false, err
	}
	ver, dirty, err := p.m.Version()
	if err != nil {
		return 0, false, err
	}
	if dirty {
		err = p.m.Force(int(ver))
		if err != nil {
			return 0, false, err
		}
	}
	err = p.m.Migrate(ver)
	if err != nil && err != migrate.ErrNoChange {
		return 0, false, err
	}
	ver, dirty, err = p.m.Version()
	if err != nil {
		return 0, false, err
	}
	return ver, dirty, err
}

// The currently active migration version. It will migrate down to 0.
// return version, is dirty, error is sql select return
func (p *SMigrate) DownTo(currentlyVersion int) (uint, bool, error) {
	err := p.m.Steps(-currentlyVersion)
	if err != nil {
		return 0, false, err
	}
	ver, dirty, err := p.m.Version()
	if err != nil {
		return 0, false, err
	}
	if dirty {
		err = p.m.Force(int(ver))
		if err != nil {
			return 0, false, err
		}
	}
	err = p.m.Migrate(ver)
	if err != nil && err != migrate.ErrNoChange {
		return 0, false, err
	}
	ver, dirty, err = p.m.Version()
	if err != nil {
		return 0, false, err
	}

	return ver, dirty, err
}
