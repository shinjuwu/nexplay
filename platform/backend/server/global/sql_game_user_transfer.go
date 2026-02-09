package global

import (
	"backend/api/game/model"
	"backend/pkg/utils"
	"database/sql"
)

func UdfGameUserStartCoinIn(db *sql.DB, id string, userId int, usename string, agentId int, agentLevelCode string,
	kind, status int, info, creator, request string,
	addAgentWalletAmount, addAgentSumCoinOut float64, walletLedgerId string) (code int, err error) {
	/*
		CREATE FUNCTION "public"."udf_game_user_start_coin_in"("_id" varchar,
		  "_user_id" int4, "_username" varchar, "_agent_id" int4,
		  "_agent_level_code" varchar, "_kind" int2, "_status" int2,
		  "_info" varchar, "_creator" varchar, "_request" jsonb,
		  "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric)
		  RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 代理錢包扣款及統計
	// 2. 創建上下分紀錄
	// 3. 回傳 json 格式，code:結果(0:成功，其他:錯誤)

	jsonResult := ""
	query := `SELECT "public"."udf_game_user_start_coin_in" ($1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10,
		$11, $12, $13)`

	err = db.QueryRow(query, id, userId, usename, agentId, agentLevelCode,
		kind, status, info, creator, request,
		addAgentWalletAmount, addAgentSumCoinOut, walletLedgerId).Scan(&jsonResult)
	if err != nil {
		code = model.Response_QueryRow_Error
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	code = TranfromUdfGameUserStartCoinInCodeToModelErrorCode(int(result["code"].(float64)))

	return
}

func TranfromUdfGameUserStartCoinInCodeToModelErrorCode(code int) int {
	switch code {
	case 0:
		return model.Response_Success
	case 1:
		return model.Response_ParseParam_Error
	case 2:
		return model.Response_AgentWalletAmountNotEnough_Error
	case 3:
		return model.Response_OrderExist_Error
	default:
		return model.Response_Local_Error
	}
}

func UdfGameUserFinishCoinIn(db *sql.DB, id, changeset string, status, errorCode, agentId int,
	addAgentWalletAmount, addAgentSumCoinOut float64, userId int, addUserSumCoinIn float64) (code int, err error) {
	/*
		CREATE FUNCTION "public"."udf_game_user_finish_coin_in"("_id" varchar,
		  "_changeset" jsonb, "_status" int2, "_error_code" int2, "_agent_id" int4,
		  "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric,
		  "_user_id" int4, "_add_user_sum_coin_in" numeric)
	*/
	// 此 sql function 內含
	// 1. 失敗代理錢包及資訊更新，成功玩家資訊更新
	// 2. 更新上下分紀錄
	// 3. 回傳 json 格式，code:結果(0:成功，其他:錯誤)

	jsonResult := ""
	query := `SELECT "public"."udf_game_user_finish_coin_in" ($1, $2, $3, $4, $5,
		$6, $7, $8, $9)`

	err = db.QueryRow(query, id, changeset, status, errorCode, agentId,
		addAgentWalletAmount, addAgentSumCoinOut, userId, addUserSumCoinIn).Scan(&jsonResult)
	if err != nil {
		code = model.Response_QueryRow_Error
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	code = TranfromUdfGameUserFinishCoinInCodeToModelErrorCode(int(result["code"].(float64)))

	return
}

func UdfGameUserFinishCoinInCancel(db *sql.DB, id, changeset string, status, errorCode, agentId int,
	addAgentWalletAmount, addAgentSumCoinOut float64, userId int, addUserSumCoinIn float64) (code int, err error) {
	/*
		CREATE FUNCTION "public"."udf_game_user_finish_coin_in"("_id" varchar,
		  "_changeset" jsonb, "_status" int2, "_error_code" int2, "_agent_id" int4,
		  "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric,
		  "_user_id" int4, "_add_user_sum_coin_in" numeric)
	*/
	// 此 sql function 內含
	// 1. 失敗代理錢包及資訊更新，成功玩家資訊更新
	// 2. 更新上下分紀錄
	// 3. 回傳 json 格式，code:結果(0:成功，其他:錯誤)

	jsonResult := ""
	query := `SELECT "public"."udf_game_user_finish_coin_in_cancel" ($1, $2, $3, $4, $5,
		$6, $7, $8, $9)`

	err = db.QueryRow(query, id, changeset, status, errorCode, agentId,
		addAgentWalletAmount, addAgentSumCoinOut, userId, addUserSumCoinIn).Scan(&jsonResult)
	if err != nil {
		code = model.Response_QueryRow_Error
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	code = TranfromUdfGameUserFinishCoinInCodeToModelErrorCode(int(result["code"].(float64)))

	return
}

func TranfromUdfGameUserFinishCoinInCodeToModelErrorCode(code int) int {
	switch code {
	case 0:
		return model.Response_Success
	case 1:
		return model.Response_ParseParam_Error
	default:
		return model.Response_Local_Error
	}
}

func UdfGameUserStartCoinOut(db *sql.DB, id string, userId int, usename string, agentId int, agentLevelCode string,
	kind, status int, info, creator, request, walletLedgerId string) (code int, err error) {
	/*
		CREATE FUNCTION "public"."udf_game_user_start_coin_out"("_id" varchar,
		  "_user_id" int4, "_username" varchar, "_agent_id" int4,
		  "_agent_level_code" varchar, "_kind" int2, "_status" int2,
		  "_info" varchar, "_creator" varchar, "_request" jsonb)
		  RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 創建上下分紀錄
	// 2. 回傳 json 格式，code:結果(0:成功，其他:錯誤)

	jsonResult := ""
	query := `SELECT "public"."udf_game_user_start_coin_out" ($1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10, $11)`

	err = db.QueryRow(query, id, userId, usename, agentId, agentLevelCode,
		kind, status, info, creator, request,
		walletLedgerId).Scan(&jsonResult)
	if err != nil {
		code = model.Response_QueryRow_Error
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	code = TranfromUdfGameUserStartCoinOutCodeToModelErrorCode(int(result["code"].(float64)))

	return
}

func TranfromUdfGameUserStartCoinOutCodeToModelErrorCode(code int) int {
	switch code {
	case 0:
		return model.Response_Success
	case 1:
		return model.Response_OrderExist_Error
	default:
		return model.Response_Local_Error
	}
}

func UdfGameUserFinishCoinOut(db *sql.DB, id, changeset string, status, errorCode, agentId int,
	addAgentWalletAmount, addAgentSumCoinIn float64, userId int, addUserSumCoinOut float64) (code int, err error) {
	/*
		CREATE FUNCTION "public"."udf_game_user_finish_coin_out"("_id" varchar,
		  "_changeset" jsonb, "_status" int2, "_error_code" int2, "_agent_id" int4,
		  "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_in" numeric,
		  "_user_id" int4, "_add_user_sum_coin_out" numeric)
	*/
	// 此 sql function 內含
	// 1. 失敗代理錢包及資訊更新，成功玩家資訊更新
	// 2. 更新上下分紀錄
	// 3. 回傳 json 格式，code:結果(0:成功，其他:錯誤)

	jsonResult := ""
	query := `SELECT "public"."udf_game_user_finish_coin_out" ($1, $2, $3, $4, $5,
		$6, $7, $8, $9)`

	err = db.QueryRow(query, id, changeset, status, errorCode, agentId,
		addAgentWalletAmount, addAgentSumCoinIn, userId, addUserSumCoinOut).Scan(&jsonResult)
	if err != nil {
		code = model.Response_QueryRow_Error
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	code = TranfromUdfGameUserFinishCoinOutCodeToModelErrorCode(int(result["code"].(float64)))

	return
}

func TranfromUdfGameUserFinishCoinOutCodeToModelErrorCode(code int) int {
	switch code {
	case 0:
		return model.Response_Success
	case 1:
		return model.Response_ParseParam_Error
	default:
		return model.Response_Local_Error
	}
}
