package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/pkg/database"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"time"
)

/*
	3.2.7 查询游戏注单
		此接口用来获取游戏对局注单
*/
func CheckGameRecord(logger ginweb.ILogger, idb database.IDatabase, rdb redis.IRedisCliect, data model.ChannelHandleRequest) (bool, *model.ChannelHandleResponse, error) {

	type CheckGameRecordResponseData struct {
		Code  int         `json:"code"`  // 錯誤碼
		Count int         `json:"count"` // 資料筆數
		List  interface{} `json:"list"`  // 資料
		Start int64       `json:"start"` // 查詢開始時間(timestamp: 10碼)
		End   int64       `json:"end"`   // 查詢結束時間(timestamp: 10碼)
	}

	returnData := CheckGameRecordResponseData{
		Start: time.Now().Unix(),
		End:   time.Now().Unix(),
	}
	returnTmp := model.CreateChannelHandleResponse(model.GetRecordHandle_CheckGameRecord, &returnData)

	startTimeStamp := utils.ToInt64(data.ParamMap["start_time"][0], 0)
	endTimeStamp := utils.ToInt64(data.ParamMap["end_time"][0], 0)

	if startTimeStamp <= 0 || endTimeStamp <= 0 {
		returnData.End = time.Now().Unix()
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	startTime := time.Unix(startTimeStamp, 0).UTC()
	endTime := time.Unix(endTimeStamp, 0).UTC()

	// 計算 startTime 和 endTime 之間的時間間隔
	duration := endTime.Sub(startTime)

	// 檢查時間間隔是否超過設定時間間隔
	if duration > TIMEDRUATION_INTERVALS_MIN {
		returnData.End = time.Now().Unix()
		returnData.Code = model.Response_TimeIntervalSetting_Error
		return true, returnTmp, nil
	}

	// 查詢DB資料，並轉成注單格式送出
	agentId := utils.StringToInt(data.Agent, 0)
	// tip: 因為每個代理幣別都是固定的，所以只要送出查詢代理的幣別資料即可
	ret, err := getGameData(logger, idb.GetDB(global.DB_IDX_ORDER), rdb, agentId, data.Currency, startTime, endTime)
	if err != nil {
		returnData.End = time.Now().Unix()
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, nil
	}
	returnData.Count = len(ret)
	returnData.List = ret
	returnData.End = time.Now().Unix()
	returnTmp.D = returnData
	return true, returnTmp, nil
}
