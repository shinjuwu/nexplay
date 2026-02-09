package data

import (
	"database/sql"
	"encoding/json"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/internal/api"
	"monitorservice/pkg/utils"
	"sync"

	sq "github.com/Masterminds/squirrel"
)

var (
	ServiceInfo *ServiceInfoCache
	DBConnInfo  *DBConnInfoCache
	DBSet       *sync.Map
	Game        *GameCache
)

func InitData(logger api.ILogger, db *sql.DB) error {
	if err := InitServiceInfo(logger, db); err != nil {
		return err
	}

	// if err := InitDBConnInfo(logger, db); err != nil {
	// 	return err
	// }

	// if err := InitDBSet(logger, db); err != nil {
	// 	return err
	// }

	if err := InitGame(logger, db); err != nil {
		return err
	}

	return nil
}

// read service info from db
func InitServiceInfo(logger api.ILogger, db *sql.DB) error {

	if ServiceInfo == nil {
		ServiceInfo = NewServiceInfoCache()
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "name", "sub_name", "api_urls", "info",
			"is_enabled", "create_time", "update_time", "disable_time").
		From("service_info").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.ServiceInfo

		if err := rows.Scan(&temp.ID, &temp.Name, &temp.SubName, &temp.APIURLsBytes, &temp.Info,
			&temp.IsEnabled, &temp.CreateTime, &temp.UpdateTime, &temp.DisableTime); err != nil {
			return err
		}

		utils.ToStruct(temp.APIURLsBytes, &temp.APIURLs)

		ServiceInfo.Add(&temp)
	}

	return nil
}

// read db conn info from db
func InitDBConnInfo(logger api.ILogger, db *sql.DB) error {

	if DBConnInfo == nil {
		DBConnInfo = NewDBConnInfoCache()
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "addresses", "info", "is_enabled", "create_time",
			"update_time", "disable_time").
		From("db_conn_info").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.DBConnInfo

		if err := rows.Scan(&temp.ID, &temp.AddressesBytes, &temp.Info, &temp.IsEnabled, &temp.CreateTime,
			&temp.UpdateTime, &temp.DisableTime); err != nil {
			return err
		}

		utils.ToStruct(temp.AddressesBytes, &temp.Addresses)

		DBConnInfo.Add(&temp)
	}

	return nil
}

//  create DB 物件
func InitDBSet(logger api.ILogger, db *sql.DB) error {

	if DBSet == nil {
		DBSet = new(sync.Map)
	}

	dbConnInfos := DBConnInfo.GetAll()

	for _, dbConnInfo := range dbConnInfos {
		db, err := sql.Open(dbConnInfo.Addresses.Driver, dbConnInfo.Addresses.GetConnInfoString())
		if err != nil {
			return err
		}
		DBSet.Store(dbConnInfo.ID, db)
	}

	return nil
}

// read game from db
func InitGame(logger api.ILogger, db *sql.DB) error {

	if Game == nil {
		Game = NewGameCache()
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "list").
		From("game").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp model.Game

		if err := rows.Scan(&temp.ID, &temp.List); err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(temp.List), &temp.GameList); err != nil {
			return err
		}

		Game.Add(&temp)
	}

	return nil
}
