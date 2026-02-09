package global

import (
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
)

func UdfCreateAgentPermission(db *sql.DB, agentId int, name, info string, accountType int, permission string,
	checkPermission bool) (agentPermission *table_model.AgentPermission, err error) {
	/*
		CREATE FUNCTION "public"."udf_create_agent_permission"("_agent_id" int4, "_name" varchar,
		    "_info" varchar, "_account_type" int2, "_permission" jsonb, "_check_permission" boolean)
		    RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 創建代理權限群組，並過濾多餘的權限
	// 2. 回傳 json 格式，包含創建的agentPermission的id及permission資料

	jsonResult := ""
	query := `SELECT "public"."udf_create_agent_permission"($1, $2, $3, $4, $5, $6)`

	err = db.QueryRow(query, agentId, name, info, accountType, permission,
		checkPermission).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	retId := result["id"].(string)
	retPermission := table_model.PermissionSlice{}
	utils.ToStruct([]byte(utils.ToJSON(result["permission"])), &retPermission)

	agentPermission = &table_model.AgentPermission{
		Id:          retId,
		AgentId:     agentId,
		AccountType: accountType,
		Name:        name,
		Info:        info,
		Permission:  retPermission,
	}

	return
}

func UdfUpdateAgentPermission(db *sql.DB, id string, agentId int, name, info string, accountType int,
	permission string, checkPermission, updateChildAgentPermission bool) (permissionSlice table_model.PermissionSlice, err error) {
	/*
		CREATE FUNCTION "public"."udf_update_agent_permission"("_id" uuid, "_agent_id" int4, "_name" varchar,
		    "_info" varchar, "_account_type" int2, "_permission" jsonb, "_check_permission" boolean,
			"_update_child_agent_permission" boolean)
		    RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 創建代理權限群組，並過濾多餘的權限
	// 2. 回傳 json 格式，包含創建的agentPermission的id及permission資料

	jsonResult := ""
	query := `SELECT "public"."udf_update_agent_permission"($1, $2, $3, $4, $5, $6, $7, $8)`

	err = db.QueryRow(query, id, agentId, name, info, accountType,
		permission, checkPermission, updateChildAgentPermission).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))
	utils.ToStruct([]byte(utils.ToJSON(result["permission"])), &permissionSlice)

	return
}
