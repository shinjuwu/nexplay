package definition

const (
	/*
		調整【輸贏報表】的【风控】欄位，需依照下列分類：
			0. 一般、1. 基礎追殺、2. 放水、3. 控牌追殺、4. 單間追殺、5. 限制倍率、6. 禁開分數
	*/
	KILLDIVE_KILL_TYPE_NORMAL           = 0
	KILLDIVE_KILL_TYPE_BASIC            = 1
	KILLDIVE_KILL_TYPE_RELEASE          = 2
	KILLDIVE_KILL_TYPE_CARD_CONTROL     = 3
	KILLDIVE_KILL_TYPE_SINGLE_ROOM      = 4
	KILLDIVE_KILL_TYPE_LIMIT_BET        = 5
	KILLDIVE_KILL_TYPE_FORBIDDEN_POINTS = 6

	/*
		【輸贏報表】新增【设定机率】欄位作為觸發時的顯示，下列舉例：
			0. 一般 = －、1. 基礎追殺 = 10%、2. 放水 = 3%、3. 控牌追殺 = 50%、4. 單間追殺 = 20%、5. 限制倍率 = x10、6. 禁開分數 = 100,000
		若該局為一般無觸發風控，則顯示【－】即可
	*/
	KILLDIVE_KILL_PROB_NONE             = "-"
	KILLDIVE_KILL_PROB_NORMAL           = 0
	KILLDIVE_KILL_PROB_BASIC            = 1
	KILLDIVE_KILL_PROB_RELEASE          = 2
	KILLDIVE_KILL_PROB_CARD_CONTROL     = 3
	KILLDIVE_KILL_PROB_SINGLE_ROOM      = 4
	KILLDIVE_KILL_PROB_LIMIT_BET        = 5
	KILLDIVE_KILL_PROB_FORBIDDEN_POINTS = 6

	/*
		新增【杀放层级】欄位作為判斷觸發時的功能數值來源，下列舉例：
			1. 總代理風控設定、2. 遊戲風控設定、3. 遊戲基礎設定、4. 預設（GameSetting）
	*/
	KILLDIVE_KILL_LEVEL_AGENT_RISK_SETTING      = 1
	KILLDIVE_KILL_LEVEL_GAME_RISK_RISK_SETTING  = 2
	KILLDIVE_KILL_LEVEL_GAME_BASIC_RISK_SETTING = 3
	KILLDIVE_KILL_LEVEL_DEFAULT_RISK_SETTING    = 4
)
