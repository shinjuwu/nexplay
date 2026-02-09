package global

import (
	"backend/api/game/model"
	"database/sql"
)

//
func CreateWalletChangeset(beforeCoin, addCoin, AfterCoin, toCoin float64, currency string) (changeset map[string]interface{}) {

	changeset = make(map[string]interface{}, 0)
	changeset["before_coin"] = beforeCoin
	changeset["add_coin"] = addCoin
	changeset["after_coin"] = AfterCoin
	changeset["currency"] = currency
	changeset["to_coin"] = toCoin

	return
}

//
func CreateWalletInfo(resaon string) (info map[string]interface{}) {

	info = make(map[string]interface{}, 0)
	info["reason"] = resaon

	return
}

func InsertIntoWalletLedgeRecord(db *sql.DB, orderId, username, levelCode, changeset, info string, userId, agentId, kind int) (errCode int) {
	query := `INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "changeset", "info", "agent_id", "kind", "level_code") 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8);`

	tx, _ := db.Begin()
	_, err := tx.Exec(query, orderId, userId, username, changeset, info, agentId, kind, levelCode)
	if err != nil {
		tx.Rollback()
		errCode = model.Response_Exec_Error
	} else {
		err = tx.Commit()
		if err != nil {
			errCode = model.Response_Commit_Error
		}
	}

	return
}
