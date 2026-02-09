package api_module

import (
	"encoding/json"
	"fmt"
	"monitorservice/cmd/baseserver/api_module/data"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/cmd/baseserver/api_module/module_lib"
	"monitorservice/internal/api"
	"monitorservice/pkg/utils"
)

// @Tags collector.api
// @Summary 同步平台遊戲列表
// @accept application/json
// @Produce  application/json
// @param data body model.CollectorGameListRequest true "紀錄平台RTP資料輸贏資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/collector.api/v1/collectorgamelist [post]
func CollectorGameList(c api.IContext, params string) (string, error) {

	var req model.CollectorGameListRequest
	req.Data = make([]*model.GameList, 0)

	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	// update db
	_, err := c.DB().Exec(
		`INSERT INTO "public"."game"("id", "list")
			SELECT 
				$1, $2
			ON CONFLICT ON CONSTRAINT "game_pkey" DO 
			UPDATE SET 	
				list = EXCLUDED.list;`,
		req.ID, utils.ToJSON(req.Data))
	if err != nil {
		return "", fmt.Errorf("CollectorGameList() failed, err is : %v", err)
	}

	// update local memory
	var temp model.Game
	temp.ID = req.ID
	temp.GameList = req.Data
	temp.List = utils.ToJSON(temp.GameList)

	data.Game.Add(&temp)

	return "", nil
}

// @Tags collector.api
// @Summary 玩家上下分資料收集
// @accept application/json
// @Produce  application/json
// @param data body model.CollectorCoinInOutRequest true "紀錄玩家上下分資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/collector.api/v1/collectorcoininout [post]
func CollectorCoinInOut(c api.IContext, params string) (string, error) {

	var req model.CollectorCoinInOutRequest
	req.Data = make([]*model.WalletLedger, 0)

	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	sqlStr := `INSERT INTO
	"public"."wallet_ledger" ("platform", "id", "agent_id", "agent_name", "user_id", 
	"username", "kind", "status", "changeset", "create_time") VALUES `

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //10

	idx := 1
	vals := []interface{}{}

	for _, v := range req.Data {
		if v != nil {
			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4,
				idx+5, idx+6, idx+7, idx+8, idx+9)
			idx += 10

			vals = append(vals, req.ID, v.ID, v.AgentID, v.AgentName, v.UserID,
				v.Username, v.Kind, v.Status, v.ChangeSet, v.CreateTime)
		}
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := c.DB().Prepare(sqlStr)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return "", err
		}
	}

	return "", nil
}

// @Tags collector.api
// @Summary 異常輸贏資料收集
// @accept application/json
// @Produce  application/json
// @param data body model.CollectorAbnormalWinAndLoseRequest true "紀錄玩家輸贏資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/collector.api/v1/collectorabnormalwinandlose [post]
func CollectorAbnormalWinAndLose(c api.IContext, params string) (string, error) {

	var req model.CollectorAbnormalWinAndLoseRequest
	req.Data = make([]*model.UserPlayLog, 0)

	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	sqlStr := `INSERT INTO
	"public"."user_play_log" ("platform", "lognumber", "agent_id", "agent_name", "user_id", 
	"username", "game_id", "game_name", "room_type", "room_name",
	"de_score", "bonus", "bet_id", "bet_time", "create_time") VALUES `

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //15

	idx := 1
	vals := []interface{}{}

	for _, v := range req.Data {
		if v != nil {
			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4,
				idx+5, idx+6, idx+7, idx+8, idx+9,
				idx+10, idx+11, idx+12, idx+13, idx+14)
			idx += 15

			vals = append(vals, req.ID, v.LogNumber, v.AgentID, v.AgentName, v.UserID,
				v.Username, v.GameID, v.GameName, v.RoomType, v.RoomName,
				v.DeScore, v.Bonus, v.BetID, v.BetTime, v.CreateTime)
		}
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := c.DB().Prepare(sqlStr)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return "", err
		}
	}

	return "", nil
}

// @Tags collector.api
// @Summary RTP 統計資料收集
// @accept application/json
// @Produce  application/json
// @param data body model.CollectorRTPStatRequest true "紀錄平台RTP資料輸贏資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/collector.api/v1/collectorrtpstat [post]
func CollectorRTPStat(c api.IContext, params string) (string, error) {

	var req model.CollectorRTPStatRequest
	req.Data = make([]*model.CollectorRTPStat, 0)

	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	for _, rtpData := range req.Data {
		module_lib.UspInsertGameRatioStat(c.DB(), req.ID,
			rtpData.GameId, rtpData.GameType, rtpData.RoomType, rtpData.PlayCount, rtpData.Ya,
			rtpData.De, rtpData.VaildYa, rtpData.Tax, rtpData.Bonus, rtpData.UpdateTime)
	}

	return "", nil
}
