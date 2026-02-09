package model

/*
	遊戲相關接口
	KY 接口參考

	3.2.1 登录游戏
		此接口用以验证游戏账号，如果账号不存在则创建游戏账号并为账号上分。
	*3.2.2 查询可下分
		此接口用来查询玩家的可下分余额
	*3.2.3 上分
		此接口用来为账号上分
	*3.2.4 下分
		此接口用来为账号下分
	3.2.5 查询订单(單筆)
		此接口用来查询玩家上下分的订单信息，通过 status 状态来判断上下分是否成功。
	*3.2.6 查询玩家在线状态
		此接口用来查询玩家是否在线
	3.2.7 查询游戏注单
		此接口用来获取游戏对局注单
	*3.2.8 查询玩家总分
		此接口用来查询玩家的游戏内总分、玩家可下分余额、玩家在线状态
	*3.2.9 踢玩家下线
		此接口用以将在线的玩家强制离线
	3.2.10 查询代理余额(非必要)
		此接口用以查询代理余额
*/

const (
	// request code
	ChannelHandle_GameLogin             = iota // 登入遊戲
	ChannelHandle_CheckCoinOutLimit            // 查詢可下分
	ChannelHandle_CoinIn                       // 上分
	ChannelHandle_CoinOut                      // 下分
	ChannelHandle_CheckUserOrder               // 查詢上下分訂單
	ChannelHandle_CheckUserOnlineStatus        // 查詢玩家在線狀態
	GetRecordHandle_CheckGameRecord            // 查詢遊戲注單
	ChannelHandle_CheckUserCoin                // 查詢玩家總分
	ChannelHandle_KickUser                     // 踢玩家下線
	ChannelHandle_CheckAgentCoin               // 查詢代理余額
	ChannelHandle_Max
)

const (

	// response code = request code + ChannelHandle_ResponseBase
	ChannelHandle_ResponseBase = 100

	// game server side error code = api return code + GameServerAPI_ErrorCodeBase
	// GameServerAPI_ErrorCodeBase = 1000

	ChannelHandle_Path   = "/channel/channelHandle"
	GetRecordHandle_Path = "/record/getRecordHandle"
)

const (
	// TO DO: move to definition
	// response error code
	Response_Success                          = 0    // 成功
	Response_NoRowResultSet_Error             = 1001 // 查詢不到資料
	Response_QueryRow_Error                   = 1002 // 查詢失敗(1)
	Response_Exec_Error                       = 1003 // 查詢失敗(2)
	Response_Commit_Error                     = 1004 // 查詢失敗(3)
	Response_ParseParam_Error                 = 1005 // 參數錯誤
	Response_AgentWalletAmountNotEnough_Error = 1006 // 代理餘額不足
	Response_OrderExist_Error                 = 1007 // 訂單編號已存在
	Response_AccountExist_Error               = 1008 // 帳號已存在
	Response_GameServeDepositFailed_Error     = 1009 // 遊戲上/下分失敗
	Response_GameServer_Error                 = 1010 // 遊戲發生未知錯誤
	Response_AccountBlock_Error               = 1011 // 帳號已被封停
	Response_CoinInOutValueFailed_Error       = 1012 // 上/下分數值不合法
	Response_GameServerBlockPlayer_Error      = 1013 // 遊戲踢除用戶失敗
	Response_PlayerNotOnline_Error            = 1014 // 用戶必須在線上
	Response_GameServerGold_Error             = 1015 // 遊戲查詢用戶餘額失敗
	Response_ParseJsonFailed_Error            = 1016 // 解析 json 失敗
	Response_TypeTransFailed_Error            = 1017 // 參數轉換失敗
	Response_GameUserNotExist_Error           = 1018 // 用戶不存在
	Response_TooManyRequests_Error            = 1019 // 請求太頻繁
	Response_RiskControlLogin_Error           = 1020 // 用戶禁止登入
	Response_RiskControlCoinIn_Error          = 1021 // 用戶禁止上分
	Response_RiskControlCoinOut_Error         = 1022 // 用戶禁止下分
	Response_TimeIntervalSetting_Error        = 1023 // 超過設定間隔時間限制

	Response_Local_Error = 1999 // 發生未知錯誤
)

const (
	// 測試使用(之後要改成讀取代理自己的 KEY 作加解密)
	DEFAULT_MD5_KEY = "hongkong3345678"  // 15碼
	DEFAULT_AES_KEY = "1234567890123456" // 16碼
)
