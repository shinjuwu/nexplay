package system

import (
	"backend/api/v1/model"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"net"

	sq "github.com/Masterminds/squirrel"
)

func getAgentGameListSumInfo(db *sql.DB, sqAnd *sq.And) (int, int) {
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(agent_id)").
		From("view_agent_game").
		Where(sqAnd).
		ToSql()
	if err != nil {
		return 0, definition.ERROR_CODE_ERROR_DATABASE
	}

	var recordsTotal int
	err = db.QueryRow(query, args...).Scan(&recordsTotal)
	if err != nil && err != sql.ErrNoRows {
		return 0, definition.ERROR_CODE_ERROR_DATABASE
	}

	return recordsTotal, definition.ERROR_CODE_SUCCESS
}

func getAgentGameList(db *sql.DB, req *model.GetAgentGameListRequest, sqAnd *sq.And) ([]*model.GetAgentGameResponse, int) {
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "agent_name", "agent_level_code", "game_id", "game_code", "state").
		From("view_agent_game").
		Where(sqAnd).
		Limit(uint64(req.Length)).
		Offset(uint64(req.Start)).
		ToSql()
	if err != nil {
		return nil, definition.ERROR_CODE_ERROR_DATABASE
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, definition.ERROR_CODE_ERROR_DATABASE
	}

	defer rows.Close()

	var data = make([]*model.GetAgentGameResponse, 0)
	for rows.Next() {
		var temp model.GetAgentGameResponse
		if err := rows.Scan(&temp.AgentId, &temp.AgentName, &temp.AgentLevelCode,
			&temp.GameId, &temp.GameCode, &temp.State); err != nil {
			return nil, definition.ERROR_CODE_ERROR_DATABASE
		}
		data = append(data, &temp)
	}

	return data, definition.ERROR_CODE_SUCCESS
}

func toPermissionString(permissions []int) string {
	return utils.ToJSON(&table_model.PermissionSlice{
		List: permissions,
	})
}

func validIpAddress(ipAddress string) bool {
	checkIp := ipAddress

	if ipAddress[len(ipAddress)-1:] == "*" {
		checkIp = ipAddress[:len(ipAddress)-1] + "1"
	}

	ip := net.ParseIP(checkIp)

	return ip != nil && ip.To4() != nil
}
