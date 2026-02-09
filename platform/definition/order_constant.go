package definition

const (
	ORDER_TYPE_AGENT_WALLET_LEDGER = iota + 1 // 代理上下分紀錄
	ORDER_TYPE_WALLET_LEDGER                  // 玩家上下分紀錄
	ORDER_TYPE_PLAY_LOG_COMMON                // 遊戲局紀錄
	ORDER_TYPE_USER_PLAY_LOG                  // 玩家遊戲紀錄
	ORDER_TYPE_JACKPOT_TOKEN_LOG              // jackpot代幣紀錄
	ORDER_TYPE_JACKPOT_LOG                    // jackpot紀錄
	ORDER_TYPE_SINGLE_WALLET_LOG              // 單一錢包紀錄
	ORDER_TYPE_FRIEND_ROOM_LOG                // 好友房紀錄
)
