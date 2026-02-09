package system

import (
	"backend/pkg/utils"
	"context"
	"database/sql"
	"definition"
)

func getConnInfo(db *sql.DB, context context.Context, code string) (map[string]interface{}, int) {
	var address []byte
	isEnabled := false
	query := "SELECT addresses, is_enabled FROM server_info WHERE code = $1 "
	err := db.QueryRowContext(context, query, code).Scan(&address, &isEnabled)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, definition.ERROR_CODE_ERROR_PERMISSION
		} else {
			return nil, definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS
		}
	} else if !isEnabled {
		return nil, definition.ERROR_CODE_ERROR_FEATURE_DISABLED
	}

	connInfo := utils.ToMap(address)

	return connInfo, definition.ERROR_CODE_SUCCESS
}

func getChatConnInfo(db *sql.DB, context context.Context) (map[string]interface{}, int) {
	return getConnInfo(db, context, "chat")
}

func getMatainConnInfo(db *sql.DB, context context.Context) (map[string]interface{}, int) {
	return getConnInfo(db, context, "maintain")
}
