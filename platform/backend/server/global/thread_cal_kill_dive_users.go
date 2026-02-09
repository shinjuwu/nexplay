package global

import (
	"backend/internal/notification"
	"backend/pkg/logger"
	"database/sql"
	"definition"
	"log"
	"time"
)

var (
	execKillDiveTimeForCheck = time.Duration(1) * time.Minute
)

func T_CalKillDiveState(logger *logger.RuntimeGoLogger, db *sql.DB) {
	c := time.NewTicker(execKillDiveTimeForCheck)
	for {
		_execKillDiveGameUsers(logger, db)
		<-c.C
	}
}

func _execKillDiveGameUsers(logger *logger.RuntimeGoLogger, db *sql.DB) {
	// 固定檢測時間
	now := time.Now().UnixNano()
	ftTimeNow := now / int64(time.Millisecond)

	logger.Printf("_execKillDiveGameUsers run, ftTimeNow: %v", ftTimeNow)

	needSendSetPlayerStatus := false
	userIds := make([]int, 0)
	states := make([]int, 0)

	f := func(key, value interface{}) bool {
		userId := key.(int)
		killDiveValMap := value.(map[string]float64)
		killDiveState := definition.GAMEUSERS_STATUS_KILLDIVE_CONFIGKILL
		// 記錄需要結束定額追殺狀態的用戶
		if killDiveValMap["kill_dive_value"] <= 0 {
			killDiveValMap["kill_dive_value"] = 0
			killDiveState = definition.GAMEUSERS_STATUS_KILLDIVE_NORMAL

			needSendSetPlayerStatus = true
			userIds = append(userIds, userId)
			states = append(states, killDiveState)
		}

		// 有變動才更新
		if killDiveValMap["need_update"] > 0 {
			// update game_users killdive state to db
			_, err := db.Exec(`UPDATE game_users SET kill_dive_state=$1, kill_dive_value=$2 WHERE "id"=$3;`,
				killDiveState, killDiveValMap["kill_dive_value"], userId)
			if err != nil {
				log.Printf("_execKillDiveGameUsers update failed, killDiveState= %v, killDiveMap= %v, userId= %v", killDiveState, killDiveValMap, userId)
			}
		}

		killDiveValMap["need_update"] = float64(0)
		GameUsersKillDiveValueList.Store(userId, killDiveValMap)
		return true
	}
	GameUsersKillDiveValueList.Range(f)

	// sent notify to gameserver
	if needSendSetPlayerStatus {
		apiResult, errCode, err := notification.SendSetPlayerStatus(userIds, states)
		if err != nil {
			logger.Printf("_execKillDiveGameUsers SendSetPlayerStatus failed, userIds: %v, apiResult= %v, errCode= %v, err= %v", userIds, apiResult, errCode, err)
		} else {
			// 刪除已結束定額追殺用戶
			f = func(key, value interface{}) bool {
				userId := key.(int)
				killDiveValMap := value.(map[string]float64)
				if killDiveValMap["kill_dive_value"] <= 0 {
					GameUsersKillDiveValueList.Delete(userId)
				}

				return true
			}
			GameUsersKillDiveValueList.Range(f)
		}
	}
}
