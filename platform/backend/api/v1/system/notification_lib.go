package system

import (
	"fmt"
	"strconv"
)

func getGameList(gameId int, gameCode string, state int16) map[string]interface{} {
	return map[string]interface{}{
		"GameId":   gameId,
		"GameCode": gameCode,
		"Status":   state,
	}
}

func getGameInfo(agentId, gameId int, gameCode string, state int16) map[string]interface{} {
	return map[string]interface{}{
		"AgentId":  agentId,
		"GameId":   gameId,
		"GameCode": gameCode,
		"Status":   state,
	}
}

func getLobbyInfo(agentId, gameId, tableId int, state int16) map[string]interface{} {
	return map[string]interface{}{
		"AgentId": agentId,
		"GameId":  gameId,
		"TableId": tableId,
		"Status":  state,
	}
}

func getKillDiveInfo(agentId, gameId, roomId, activeNum int, killRate, newKillRate float64) map[string]interface{} {
	return map[string]interface{}{
		"AgentId":     agentId,
		"GameId":      gameId,
		"RoomId":      roomId,
		"Killrate":    killRate,
		"Newkillrate": newKillRate,
		"Activenum":   activeNum,
	}
}

func getRoomId(gameId, roomType int) (int, error) {
	return strconv.Atoi(fmt.Sprintf("%d%d", gameId, roomType))
}
