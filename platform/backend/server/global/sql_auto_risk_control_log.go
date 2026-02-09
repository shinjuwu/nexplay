package global

import (
	"backend/api/game/model"
	"database/sql"
)

func InsertAutoRiskControlLog(db *sql.DB, agentId, userId, riskCode int, username, levelCode string) (errCode int) {
	query := `INSERT INTO "public"."auto_risk_control_log" ("agent_id", "user_id", "username", "level_code", "risk_code")
	VALUES ($1,$2,$3,$4,$5);`

	_, err := db.Exec(query, agentId, userId, username, levelCode, riskCode)
	if err != nil {
		return model.Response_Exec_Error
	}

	return model.Response_Success
}
