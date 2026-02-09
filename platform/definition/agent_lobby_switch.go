package definition

const (
	LOBBY_NORMAL_SWITCH      = 1 << 0 // 一般大廳
	LOBBY_FRIENDSROOM_SWITCH = 1 << 1 // 好友房大廳

	LOBBY_COUNT = 2
)

var (
	GameIdToLobbySwitch = map[int]int{
		GAME_ID_LOBBY:       LOBBY_NORMAL_SWITCH,
		GAME_ID_FRIENDSROOM: LOBBY_FRIENDSROOM_SWITCH,
	}
)
