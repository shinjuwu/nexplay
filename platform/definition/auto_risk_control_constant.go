package definition

const (
	AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_API_REQUEST               = iota + 1 // 玩家每秒請求api次數過多
	AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_COIN_IN_AND_COIN_OUT_DIFF            // 玩家上下分的差距過高
	AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_WIN_RATE                             // 玩家勝率過高
	AUTO_RISK_CONTROL_RISK_CODE_COUNT                                          // 風險代碼數量
)

const (
	RISK_CONTROL_STATUS_ENABLED  = "1" // 風控啟用
	RISK_CONTROL_STATUS_DISABLED = "0" // 風控禁用
)

const (
	RISK_CONTROL_LOGIN_IDX    = iota // 登入index
	RISK_CONTROL_BET_IDX             // 下注index
	RISK_CONTROL_COIN_IN_IDX         // 上分index
	RISK_CONTROL_COIN_OUT_IDX        // 下分index
)
