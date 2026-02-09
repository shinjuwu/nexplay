package table

import (
	"backend/internal/squirrel"
	table_model "backend/server/table/model"
	"database/sql"
)

var (
	table_cluster = make(map[string]*squirrel.SquirrelModel)
)

const (
	default_connoection  = "dcc_game"
	TABLENAME_ADMIN_USER = "admin_user"
)

func CreateTableCluster(db *sql.DB) {

	table_cluster := make(map[string]*squirrel.SquirrelModel)
	table_cluster[TABLENAME_ADMIN_USER] = squirrel.NewSquirrelModel(TABLENAME_ADMIN_USER, table_model.AdminUser{})

}

func GetTableCluster(table string) *squirrel.SquirrelModel {
	return table_cluster[table]
}
