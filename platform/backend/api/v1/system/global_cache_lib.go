package system

import (
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"definition"
	"strings"

	"golang.org/x/exp/slices"
)

func GetAgentList(agentId int) []*table_model.Agent {
	agent := global.AgentCache.Get(agentId)
	if agent == nil {
		return nil
	}

	agents := make([]*table_model.Agent, 0)
	for _, targetAgent := range global.AgentCache.GetAll() {
		if strings.HasPrefix(targetAgent.LevelCode, agent.LevelCode) {
			agents = append(agents, targetAgent)
		}
	}
	return agents
}

func GetAgentInfoList(agentId int) []*map[string]interface{} {
	agents := GetAgentList(agentId)
	if agents == nil {
		return nil
	}

	agentInfo := make([]*map[string]interface{}, 0)
	for _, agent := range agents {
		if agent != nil {
			agentInfo = append(agentInfo, &map[string]interface{}{
				"id":          agent.Id,
				"name":        agent.Name,
				"level_code":  agent.LevelCode,
				"cooperation": agent.Cooperation,
			})
		}
	}
	return agentInfo
}

func GetGameList(accountType, agentId int, states []int16) []*table_model.Game {
	agent := global.AgentCache.Get(agentId)
	if agent == nil {
		return nil
	}

	gameCache := global.GameCache.GetAll()

	games := make([]*table_model.Game, 0)

	for _, game := range gameCache {
		// 大廳類僅為設定用不用顯示
		if game.Type == definition.GAME_TYPE_LOBBY {
			continue
		}

		if slices.IndexFunc(states, func(state int16) bool { return game.State == state }) < 0 {
			continue
		}

		// 總代理、子代理要檢查 agent game
		if accountType != definition.ACCOUNT_TYPE_ADMIN {
			agentGame := global.AgentGameCache.Get(agentId, game.Id)
			if agentGame != nil {
				if slices.IndexFunc(states, func(state int16) bool { return agentGame.State == state }) < 0 {
					continue
				}
			} else {
				continue
			}
		}

		games = append(games, game)
	}

	return games
}

func GetGameIdList(accountType, agentId int, states []int16) []int {
	games := GetGameList(accountType, agentId, states)
	if games == nil {
		return nil
	}

	gameIds := make([]int, 0)
	for _, game := range games {
		gameIds = append(gameIds, game.Id)
	}
	return gameIds
}

func GetRoomTypeList(accountType, agentId int) []int {
	games := GetGameList(accountType, agentId, definition.GAME_STATES_ONLINE_MAINTAIN)
	if games == nil {
		return nil
	}

	roomTypeMap := make(map[int]interface{}, 0)
	for _, game := range games {
		gameRooms := global.GameRoomCache.GetGameRooms(game.Id)

		for _, gameRoom := range gameRooms {
			// 總代理、子代理要檢查 agent game room
			if accountType != definition.ACCOUNT_TYPE_ADMIN {
				agentGameRoom := global.AgentGameRoomCache.Get(agentId, gameRoom.Id)

				// 總代理、子代理下架的遊戲房間跳過
				if agentGameRoom.State == definition.GAME_STATE_OFFLINE {
					continue
				}
			}

			roomTypeMap[gameRoom.RoomType] = nil
		}
	}

	roomTypes := make([]int, 0)
	for roomType := range roomTypeMap {
		roomTypes = append(roomTypes, roomType)
	}
	return roomTypes
}

func GetTemplatePermissions(agenId int) map[int][]int {
	agent := global.AgentCache.Get(agenId)
	if agent == nil {
		return nil
	}

	adminUser := global.AdminUserCache.Get(agent.Id, agent.AdminUsername)
	agentPermission := global.AgentPermissionCache.Get(adminUser.PermissionId)

	permissionMap := make(map[int][]int)
	for _, template := range global.AgentPermissionCache.GetTemplates(agentPermission.AccountType) {
		// 開發商取得總代的template不用過濾權限(權限全開)
		// 其餘(開發商取得開發商、總代取得總代、總代取的子代、子代取得子代)要過去權限
		if adminUser.AccountType == definition.ACCOUNT_TYPE_ADMIN &&
			template.AccountType == definition.ACCOUNT_TYPE_GENERAL {
			permissionMap[template.AccountType] = template.Permission.List
		} else {
			permissionMap[template.AccountType] = utils.ArrayIntersection(template.Permission.List, agentPermission.Permission.List)
		}
	}
	return permissionMap
}

func GetAgentPermissionList(agentId int, accountType int, agentPermissionId string) []*map[string]interface{} {
	userAgentPermission := global.AgentPermissionCache.Get(agentPermissionId)
	if userAgentPermission == nil {
		return nil
	}

	agentPermissions := make([]*map[string]interface{}, 0)
	for _, agentPermission := range global.AgentPermissionCache.GetByAgentAccountType(agentId, accountType) {
		agentPermissions = append(agentPermissions, &map[string]interface{}{
			"id":   agentPermission.Id,
			"name": agentPermission.Name,
		})
	}

	return agentPermissions
}
