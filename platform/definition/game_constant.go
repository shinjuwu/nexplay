package definition

const (
	/*
		game id sync with DB
			1. 此處定義遊戲 id ,此 id 必須唯一
			2. 每次新增遊戲 id 時,此id 必須與遊戲方同步
	*/

	GAME_ID_ALL             = 999  // 全部遊戲<自動過濾掉大廳>
	GAME_ID_LOBBY           = 0    // 大廳
	GAME_ID_FRIENDSROOM     = 2    // 好友房
	GAME_ID_BACCARAT        = 1001 // 百家樂
	GAME_ID_FANTAN          = 1002 // 番攤
	GAME_ID_COLORDISC       = 1003 // 色碟
	GAME_ID_PRAWNCRAB       = 1004 // 魚蝦蟹
	GAME_ID_HUNDREDSICBO    = 1005 // 百人骰寶
	GAME_ID_COCKFIGHT       = 1006 // 鬥雞
	GAME_ID_DOGRACING       = 1007 // 賽狗
	GAME_ID_ROCKET          = 1008 // 火箭
	GAME_ID_ANDARBAHAR      = 1009 // 安達巴哈
	GAME_ID_ROULETTE        = 1010 // 輪盤
	GAME_ID_BLACKJACK       = 2001 // 21點
	GAME_ID_SANGONG         = 2002 // 三公
	GAME_ID_BULLBULL        = 2003 // 牛牛
	GAME_ID_TEXAS           = 2004 // 德州撲克
	GAME_ID_RUMMY           = 2005 // 拉密
	GAME_ID_GOLDENFLOWER    = 2006 // 炸金花
	GAME_ID_POKDENG         = 2007 // 泰式博丁
	GAME_ID_CATTE           = 2008 // 越南Catte
	GAME_ID_CHINESEPOKER    = 2009 // 十三水
	GAME_ID_OKEY            = 2010 // 土耳其麻將
	GAME_ID_TEENPATTI       = 2011 // 印度炸金花
	GAME_ID_FRUITSLOT       = 3001 // 水果機
	GAME_ID_RCFISHING       = 3002 // 三國捕魚
	GAME_ID_PLINKO          = 3003 // 彈珠檯
	GAME_ID_HAPPYFISHING    = 3004 // 歡樂捕魚
	GAME_ID_FRUIT777SLOT    = 4001 // 水果777
	GAME_ID_MEGSHARKSLOT    = 4002 // 巨齒鯊
	GAME_ID_MIDASSLOT       = 4003 // 邁達斯之手
	GAME_ID_WILDGEMSLOT     = 4004 // 狂野寶石
	GAME_ID_JUMPHIGHSLOT    = 4005 // 跳高高
	GAME_ID_PYRTREASURESLOT = 4006 // 金字塔寶藏
	GAME_ID_FRIENDSTEXAS    = 5001 // 好友房德撲

	GAME_CODE_LOBBY           = "lobby"
	GAME_CODE_BACCARAT        = "baccarat"
	GAME_CODE_FANTAN          = "fantan"
	GAME_CODE_COLORDISC       = "colordisc"
	GAME_CODE_PRAWNCRAB       = "prawncrab"
	GAME_CODE_HUNDREDSICBO    = "hundredsicbo"
	GAME_CODE_COCKFIGHT       = "cockfight"
	GAME_CODE_DOGRACING       = "dogracing"
	GAME_CODE_ROCKET          = "rocket"
	GAME_CODE_ANDARBAHAR      = "andarbahar"
	GAME_CODE_ROULETTE        = "roulette"
	GAME_CODE_BLACKJACK       = "blackjack"
	GAME_CODE_SANGONG         = "sangong"
	GAME_CODE_BULLBULL        = "bullbull"
	GAME_CODE_TEXAS           = "texas"
	GAME_CODE_RUMMY           = "rummy"
	GAME_CODE_GOLDENFLOWER    = "goldenflower"
	GAME_CODE_POKDENG         = "pokdeng"
	GAME_CODE_CATTE           = "catte"
	GAME_CODE_CHINESEPOKER    = "chinesepoker"
	GAME_CODE_OKEY            = "okey"
	GAME_CODE_TEENPATTI       = "teenpatti"
	GAME_CODE_FRUITSLOT       = "fruitslot"
	GAME_CODE_RCFISHING       = "rcfishing"
	GAME_CODE_PLINKO          = "plinko"
	GAME_CODE_HAPPYFISHING    = "happyfishing"
	GAME_CODE_FRUIT777SLOT    = "fruit777slot"
	GAME_CODE_MEGSHARKSLOT    = "megsharkslot"
	GAME_CODE_MIDASSLOT       = "midasslot"
	GAME_CODE_WILDGEMSLOT     = "wildgemslot"
	GAME_CODE_JUMPHIGHSLOT    = "jumphighslot"
	GAME_CODE_PYRTREASURESLOT = "pyrtreasureslot"
	GAME_CODE_FRIENDSTEXAS    = "friendstexas"

	GAME_STATE_ALL      = -1 // 全部遊戲狀態
	GAME_STATE_ONLINE   = 1  // 遊戲正常啟動狀態
	GAME_STATE_MAINTAIN = 2  // 遊戲維護中狀態
	GAME_STATE_OFFLINE  = 0  // 遊戲關閉下架狀態

	GAME_TYPE_ALL             = 0
	GAME_TYPE_LOBBY           = 0 // 大廳
	GAME_TYPE_BAIREN          = 1 // 百人遊戲
	GAME_TYPE_CHIPAI          = 2 // 棋牌遊戲
	GAME_TYPE_ELECTRONIC_GAME = 3 // 電子遊戲
	GAME_TYPE_SLOT            = 4 // 老虎機
	GAME_TYPE_FRIENDSROOM     = 5 // 好友房
)

var (
	GAME_STATES_ONLINE_MAINTAIN_OFFLINE = []int16{GAME_STATE_ONLINE, GAME_STATE_MAINTAIN, GAME_STATE_OFFLINE}
	GAME_STATES_ONLINE_MAINTAIN         = []int16{GAME_STATE_ONLINE, GAME_STATE_MAINTAIN}
)
