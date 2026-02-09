package global

import (
	"backend/api/game/model"
	"database/sql"
	"definition"
)

func UpdateGameUserRiskControlStatusByIndex(db *sql.DB, userId, riskControlStatusIndex int, riskControlStatus string) (errCode int) {
	riskControlStatusFmt := ""
	switch riskControlStatusIndex {
	case definition.RISK_CONTROL_LOGIN_IDX:
		riskControlStatusFmt = "$2 || substring(risk_control_status,2,3)"
	case definition.RISK_CONTROL_BET_IDX:
		riskControlStatusFmt = "substring(risk_control_status,1,1) || $2 || substring(risk_control_status,3,2)"
	case definition.RISK_CONTROL_COIN_IN_IDX:
		riskControlStatusFmt = "substring(risk_control_status,1,2) || $2 || substring(risk_control_status,4,1)"
	case definition.RISK_CONTROL_COIN_OUT_IDX:
		riskControlStatusFmt = "substring(risk_control_status,1,3) || $2"
	}

	if riskControlStatusFmt == "" {
		return model.Response_Local_Error
	}

	query := "UPDATE game_users SET risk_control_status = " + riskControlStatusFmt + " WHERE id = $1"

	_, err := db.Exec(query, userId, riskControlStatus)
	if err != nil {
		return model.Response_Exec_Error
	}

	return model.Response_Success
}
