package model

import (
	"backend/server/global"
	"definition"
	"time"
)

type GetAgentIncomeRatioListRequest struct {
	AgentId   int `json:"agent_id"`   // 代理id (no search default: 0)
	StateType int `json:"state_type"` // 狀態類型 (-1:全部,0:停用,1:啟用)
}

func (p *GetAgentIncomeRatioListRequest) CheckParams() bool {
	if p.AgentId < definition.AGENT_ID_ALL ||
		p.StateType < definition.STATE_TYPE_ALL || p.StateType > definition.STATE_TYPE_ENABLED {
		return false
	}

	return true
}

type GetAgentIncomeRatioListResponse struct {
	AgentId    int       `json:"agent_id"` // 代理id (no search default: 0)
	AgentName  string    `json:"agent_name"`
	Ratio      float64   `json:"ratio"` // 總代理殺率
	State      bool      `json:"state"` // 狀態是否啟用
	Info       string    `json:"info"`  // 備註
	UpdateTime time.Time `json:"update_time"`
}

//***************************************************************************

type GetAgentIncomeRatioRequest struct {
	AgentId int `json:"agent_id"` // 代理id (no search default: 0)
}

func (p *GetAgentIncomeRatioRequest) CheckParams() bool {
	return p.AgentId > definition.AGENT_ID_ALL
}

type GetAgentIncomeRatioResponse struct {
	AgentId    int       `json:"agent_id"` // 代理id (no search default: 0)
	AgentName  string    `json:"agent_name"`
	Ratio      float64   `json:"ratio"` // 總代理殺率
	State      bool      `json:"state"` // 狀態是否啟用
	Info       string    `json:"info"`  // 備註
	UpdateTime time.Time `json:"update_time"`
}

//***************************************************************************

type SetAgentIncomeRatioRequest struct {
	AgentId int     `json:"agent_id"` // 代理id (no search default: 0)
	Ratio   float64 `json:"ratio"`    // 總代理殺率
	State   bool    `json:"state"`    // 狀態是否啟用
	Info    string  `json:"info"`     // 備註
}

func (p *SetAgentIncomeRatioRequest) CheckParams() bool {
	if p.AgentId <= definition.AGENT_ID_ALL ||
		p.Ratio < 0 {
		return false
	}
	return p.AgentId > definition.AGENT_ID_ALL
}

// type SetAgentIncomeRatioResponse struct {
// 	AgentId    int       `json:"agent_id"` // 代理id (no search default: 0)
// 	AgentName  string    `json:"agent_name"`
// 	Ratio      float64   `json:"ratio"`      // 總代理殺率
// 	StateType  bool      `json:"state_type"` // 狀態類型 (-1:全部,0:停用,1:啟用)
// 	Info       string    `json:"info"`       // 備註
// 	UpdateTime time.Time `json:"update_time"`
// }

// ***************************************************************************
type GetIncomeRatioListRequest struct {
	AgentId  int `json:"agent_id"`  // 代理id (no search default: 0)
	GameType int `json:"game_type"` // 遊戲類型 (no search default: -1)
	GameId   int `json:"game_id"`   // 遊戲id (no search default: 0)
	RoomType int `json:"room_type"` // 房間類型 (no search default: 0)
}

func (p *GetIncomeRatioListRequest) CheckParams() bool {
	if p.AgentId < definition.AGENT_ID_ALL ||
		p.GameType <= definition.GAME_TYPE_ALL || // can't select all
		p.GameId < definition.GAME_ID_ALL ||
		p.RoomType < definition.ROOM_TYPE_ALL {
		return false
	}

	return true
}

type GetIncomeRatioListResponse struct {
	Id             string    `json:"id"`               // 唯一碼(pKey,修改時傳入此碼)
	AgentId        int       `json:"agent_id"`         // 代理id
	AgnetName      string    `json:"agent_name"`       // 代理名稱
	GameType       int       `json:"game_type"`        // 遊戲類型
	GameId         int       `json:"game_id"`          // 遊戲id (no search default: 0)
	GameCode       string    `json:"game_code"`        // 遊戲代碼(唯一)
	RoomType       int       `json:"room_type"`        // 房間類型
	KillRatio      float64   `json:"kill_ratio"`       // 基礎殺率
	NewKillRatio   float64   `json:"new_kill_ratio"`   // 新手殺率
	ActiveNum      int       `json:"active_num"`       // 啟動人數
	LastUpdateTime time.Time `json:"last_update_time"` // 最後更新時間
	Info           string    `json:"info"`             // 備註
	IsParent       bool      `json:"is_parent"`        // 是否為上級設定
}

//***************************************************************************

type GetIncomeRatioRequest struct {
	AgentId  int `json:"agent_id"`  // 代理id (no search default: 0)
	GameType int `json:"game_type"` // 遊戲類型 (no search default: -1)
	GameId   int `json:"game_id"`   // 遊戲id (no search default: 0)
	RoomType int `json:"room_type"` // 房間類型 (no search default: 0)
}

func (p *GetIncomeRatioRequest) CheckParams() bool {
	// all condition can't select all
	if p.GameType == definition.GAME_TYPE_BAIREN {
		// if gameType == BAIREN, don't need check agentId
		if p.GameType <= definition.GAME_TYPE_ALL ||
			p.GameId <= definition.GAME_ID_ALL ||
			p.RoomType <= definition.ROOM_TYPE_ALL {
			return false
		}
	} else {
		if p.AgentId <= definition.AGENT_ID_ALL ||
			p.GameType <= definition.GAME_TYPE_ALL ||
			p.GameId <= definition.GAME_ID_ALL ||
			p.RoomType <= definition.ROOM_TYPE_ALL {
			return false
		}
	}

	return true
}

type GetIncomeRatioResponse struct {
	Id             string    `json:"id"`               // 唯一碼(pKey,修改時傳入此碼)
	AgentId        int       `json:"agent_id"`         // 代理id
	AgnetName      string    `json:"agent_name"`       // 代理名稱
	GameType       int       `json:"game_type"`        // 遊戲類型
	GameId         int       `json:"game_id"`          // 遊戲id (no search default: 0)
	GameCode       string    `json:"game_code"`        // 遊戲代碼(唯一)
	RoomType       int       `json:"room_type"`        // 房間類型
	KillRatio      float64   `json:"kill_ratio"`       // 基礎殺率
	NewKillRatio   float64   `json:"new_kill_ratio"`   // 新手殺率
	ActiveNum      int       `json:"active_num"`       // 啟動人數
	LastUpdateTime time.Time `json:"last_update_time"` // 最後更新時間
	Info           string    `json:"info"`             // 備註
	IsParent       bool      `json:"is_parent"`        // 是否為上級設定
}

//***************************************************************************

type SetIncomeRatioRequest struct {
	Id           string  `json:"id"`             // 唯一碼(pKey,修改時傳入此碼)
	KillRatio    float64 `json:"kill_ratio"`     // 基礎殺率
	NewKillRatio float64 `json:"new_kill_ratio"` // 新手殺率
	ActiveNum    int     `json:"active_num"`     // 啟動人數
	Info         string  `json:"info"`           // 備註
}

func (p *SetIncomeRatioRequest) CheckParams() bool {
	return true
}

type SetIncomeRatioResponse struct {
	AgentId        int       `json:"agent_id"`         // 代理id
	GameId         int       `json:"game_id"`          // 遊戲id
	GameType       int       `json:"game_type"`        // 遊戲類型
	RoomType       int       `json:"room_type"`        // 房間類型
	KillRatio      float64   `json:"kill_ratio"`       // 基礎殺率
	DiveRatio      float64   `json:"dive_ratio"`       // 基礎放水率
	UpRatioLimit   float64   `json:"up_ratio_limit"`   // 上水線
	DownRatioLimit float64   `json:"down_ratio_limit"` // 下水線
	LastUpdateTime time.Time `json:"last_update_time"` // 最後更新時間
	Info           string    `json:"info"`             // 備註
}

//***************************************************************************

type SetIncomeRatiosRequest struct {
	AgentIdList  []int   `json:"agent_id_list"`  // 代理 id list
	GameIdList   []int   `json:"game_id_list"`   // 遊戲 id list
	RoomTypeList []int   `json:"room_type_list"` // 房間類型list
	KillRatio    float64 `json:"kill_ratio"`     // 基礎殺率
	NewKillRatio float64 `json:"new_kill_ratio"` // 新手殺率
	ActiveNum    int     `json:"active_num"`     // 啟動人數
	Info         string  `json:"info"`           // 備註
}

type SetIncomeRatiosResponse struct {
	ErrorList []map[string]int `json:"error_list"` // 錯誤列表會提供 agent_id, game_id, room_type
}

//***************************************************************************

type GetAgentIncomeRatioAndGameRequest struct {
	AgentId   int       `json:"agent_id"`   // 代理id
	GameId    int       `json:"game_id"`    // 遊戲id
	RoomType  int       `json:"room_type"`  // 房間類型
	StartTime time.Time `json:"start_time"` // 開始時間(UTC+0)，請前端直接依照時區送入換算後的時間
	EndTime   time.Time `json:"end_time"`   // 結束時間(UTC+0)，請前端直接依照時區送入換算後的時間
}

func (p *GetAgentIncomeRatioAndGameRequest) CheckParams() bool {
	if p.AgentId < definition.AGENT_ID_ALL ||
		p.GameId < definition.GAME_ID_ALL ||
		p.RoomType < definition.ROOM_TYPE_ALL ||
		p.StartTime.After(p.EndTime) {
		return false
	}

	return true
}

type GetAgentIncomeRatioAndGameResponse struct {
	LogTime      string  `json:"log_time"`       // 時間字串(YYYYMMDD)
	Id           string  `json:"id"`             // 唯一碼
	LevelCode    string  `json:"level_code"`     // 層級碼
	AgentId      int     `json:"agent_id"`       // 代理id
	AgnetName    string  `json:"agent_name"`     // 代理名稱
	GameType     int     `json:"game_type"`      // 遊戲類型
	GameId       int     `json:"game_id"`        // 遊戲id (no search default: 0)
	GameCode     string  `json:"game_code"`      // 遊戲代碼(唯一)
	RoomType     int     `json:"room_type"`      // 房間類型
	BetCount     int     `json:"bet_count"`      // 下注次數
	KillRatio    float64 `json:"kill_ratio"`     // 基礎殺率
	DeScore      float64 `json:"de_score"`       // 累積玩家總贏分(float64)
	YaScore      float64 `json:"ya_score"`       // 累積玩家總投注(float64)
	VaildYaScore float64 `json:"vaild_ya_score"` // 累積玩家總有效投注(float64)
	Tax          float64 `json:"tax"`            // 累積玩家總遊戲抽水(float64)
	Bonus        float64 `json:"bonus"`          // 累積玩家總紅利(float64)
}

//***************************************************************************

type GetPlayerIncomeRatioAndGameRequest struct {
	AgentId  int    `json:"agent_id"` // 代理id
	Username string `json:"username"` // 玩家username
}

func (p *GetPlayerIncomeRatioAndGameRequest) CheckParams() bool {
	return (p.AgentId > -1)
}

type GetPlayerIncomeRatioAndGameResponse struct {
	AgentId      int     `json:"agent_id"`       // 代理id
	AgnetName    string  `json:"agent_name"`     // 代理名稱
	GameUsersId  int     `json:"game_users_id"`  // 玩家ID
	GameUsername string  `json:"game_username"`  // 玩家username
	RiskTag      string  `json:"risk_tag"`       // 風險玩家旗標
	DeScore      float64 `json:"de_score"`       // 總得遊戲分(float64)
	YaScore      float64 `json:"ya_score"`       // 總壓遊戲分(float64)
	VaildYaScore float64 `json:"vaild_ya_score"` // 總有效壓遊戲分(float64)
	BetResult    float64 `json:"bet_result"`     // 輸贏結果(float64)
}

//***************************************************************************

type GetAgentCustomTagSettingResponse struct {
	AgentId       int                           `json:"agent_id"`        // 代理id
	CustomTagInfo map[int]*global.CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
	UpdateTime    time.Time                     `json:"update_time"`     // 更新時間
}

//***************************************************************************

type SetAgentCustomTagSettingListRequest struct {
	AgentId       int                           `json:"agent_id"`        // 代理編號
	CustomTagInfo map[int]*global.CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
}

func (p *SetAgentCustomTagSettingListRequest) CheckParams() bool {
	if p.AgentId <= 0 || len(p.CustomTagInfo) == 0 {
		return false
	}
	return true
}

// type SetAgentCustomTagSettingResponse struct {
// 	AgentId        int       `json:"agent_id"`         // 代理id
// 	GameId         int       `json:"game_id"`          // 遊戲id
// 	GameType       int       `json:"game_type"`        // 遊戲類型
// 	RoomType       int       `json:"room_type"`        // 房間類型
// 	KillRatio      float64   `json:"kill_ratio"`       // 基礎殺率
// 	DiveRatio      float64   `json:"dive_ratio"`       // 基礎放水率
// 	UpRatioLimit   float64   `json:"up_ratio_limit"`   // 上水線
// 	DownRatioLimit float64   `json:"down_ratio_limit"` // 下水線
// 	LastUpdateTime time.Time `json:"last_update_time"` // 最後更新時間
// 	Info           string    `json:"info"`             // 備註
// }

//***************************************************************************

type GetGameUsersCustomTagListRequest struct {
	AgentId      int       `json:"agent_id"`       // 代理編號
	TagIdx       int       `json:"tag_idx"`        // 自定義標籤索引
	CalTranInOut float64   `json:"cal_tran_inout"` // 累積轉入轉出
	RTP          float64   `json:"rtp"`            // RTP
	WinPercent   float64   `json:"win_percent"`    // 勝率
	StartTime    time.Time `json:"start_time"`     // 開始時間(UTC+0))
	EndTime      time.Time `json:"end_time"`       // 結束時間(UTC+0))
}

func (p *GetGameUsersCustomTagListRequest) CheckParams() bool {
	if p.AgentId < definition.AGENT_ID_ALL ||
		p.TagIdx < 0 ||
		p.RTP < 0 ||
		p.CalTranInOut < 0 ||
		p.WinPercent < 0 ||
		p.StartTime.After(p.EndTime) {
		return false
	}
	return true
}

type GetGameUsersCustomTagListResponse struct {
	CustomTagInfo map[int]*global.CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
	DataList      []GameUsersCustomTagList      `json:"data_list"`
}

type GameUsersCustomTagList struct {
	AgentId       int     `json:"agent_id"` // 代理id
	AgentName     string  `json:"agent_name"`
	GameUserId    int     `json:"game_users_id"`   // 玩家id
	GameUsersName string  `json:"game_users_name"` // 玩家帳號
	IsEnabled     bool    `json:"is_enabled"`      // 是否為封停
	HighRisk      bool    `json:"high_risk"`       // 是否為高風險
	KillDiveState int     `json:"kill_dive_state"` // 殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)
	KillDiveValue float64 `json:"kill_dive_value"` // 定點額度(只有在殺放狀態設定為定點時有效)
	TagList       string  `json:"tag_list"`        // 玩家自訂義標示(8bit)
	Ya            float64 `json:"ya"`              // 玩家總壓分
	ValidYa       float64 `json:"valid_ya"`        // 玩家總有效壓分
	De            float64 `json:"de"`              // 玩家總得分
	Tax           float64 `json:"tax"`             // 玩家總抽水
	PlayCount     int     `json:"play_count"`      // 總遊玩次數
	WinCount      int     `json:"win_count"`       // 總贏次數
	BgiWinCount   int     `json:"big_win_count"`   // 大獎次數
	Bonus         float64 `json:"bonus"`           // 紅利
}

//***************************************************************************

type GetGameUsersCustomTagRequest struct {
	GameUserId    int    `json:"game_users_id"`   // 玩家id
	GameUsersName string `json:"game_users_name"` // 玩家帳號
}

func (p *GetGameUsersCustomTagRequest) CheckParams() bool {
	if p.GameUserId < 0 ||
		p.GameUsersName == "" {
		return false
	}
	return true
}

type GetGameUsersCustomTagResponse struct {
	GameUsersId   int                           `json:"game_users_id"`   // 玩家id
	GameUsersname string                        `json:"game_users_name"` // 玩家帳號
	HighRisk      bool                          `json:"high_risk"`       // 是否為高風險
	KillDiveState int                           `json:"kill_dive_state"` // 殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)
	KillDiveValue float64                       `json:"kill_dive_value"` // 定點額度(只有在殺放狀態設定為定點時有效)
	TagList       string                        `json:"tag_list"`        // 玩家自訂義標示(8bit)
	CustomTagInfo map[int]*global.CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
}

//***************************************************************************

type SetGameUsersCustomTagRequest struct {
	GameUsersId   int                           `json:"game_users_id"`   // 玩家id
	GameUsername  string                        `json:"game_username"`   // 玩家帳號
	HighRisk      bool                          `json:"high_risk"`       // 是否為高風險
	KillDiveState int                           `json:"kill_dive_state"` // 殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)
	KillDiveValue float64                       `json:"kill_dive_value"` // 定點額度(只有在殺放狀態設定為定點時有效)
	TagList       string                        `json:"tag_list"`        // 玩家自訂義標示(8bit)
	CustomTagInfo map[int]*global.CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊,對應TagList內每個位置的字串(array length:8)
}

func (p *SetGameUsersCustomTagRequest) CheckParams() bool {
	if p.GameUsersId <= 0 ||
		p.GameUsername == "" ||
		p.KillDiveState < definition.GAMEUSERS_STATUS_KILLDIVE_NORMAL ||
		p.KillDiveState > definition.GAMEUSERS_STATUS_KILLDIVE_BLACKKILL ||
		p.KillDiveValue < 0 ||
		(p.KillDiveState == definition.GAMEUSERS_STATUS_KILLDIVE_CONFIGKILL && p.KillDiveValue <= 0) ||
		len(p.TagList) != 8 ||
		len(p.CustomTagInfo) != 8 {
		return false
	}
	return true
}

// type SetGameUsersCustomTagResponse struct {

// }
