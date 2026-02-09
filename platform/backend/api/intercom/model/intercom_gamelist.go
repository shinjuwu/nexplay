package model

/*
遊戲方資料表結構 - 遊戲列表
*/
type GSGameList struct {
	Id       int    `json:"id,omitempty"`        //流水號
	GameId   int    `json:"game_id,omitempty"`   //遊戲id(唯一)
	GameCode string `json:"game_code,omitempty"` //遊戲代碼(唯一)
	GameName string `json:"game_name,omitempty"` //名稱
	Status   bool   `json:"status,omitempty"`    //目前開放狀態
}

/*
遊戲方資料表結構 - 房間設定
*/
type GSRoomList struct {
	Id       int    `json:"id,omitempty"`        //流水號
	AgentId  int    `json:"agent_id,omitempty"`  //代理編號
	GameId   int    `json:"game_id,omitempty"`   //遊戲id(唯一, ex: 1001))
	TableId  int    `json:"table_id,omitempty"`  //桌子id(唯一, 格式: 遊戲id+RoomType+桌碼, ex: 100101)
	GameCode string `json:"game_code,omitempty"` //遊戲代碼(唯一)
	GameType int    `json:"game_type,omitempty"` //遊戲類型
	GameName string `json:"game_name,omitempty"` //遊戲名稱
	RoomType int    `json:"room_name,omitempty"` //房間類型
	Language string `json:"language,omitempty"`  //支援語系(json)
	Status   bool   `json:"status,omitempty"`    //目前開放狀態
}

type GameListRequest struct {
	// AgentId int `json:"agent_id,omitempty"` //代理編號
	GameId int `json:"game_id,omitempty"` //遊戲id(唯一)
}

//?
type OrderSetting struct {
	Order   int `json:"order"`   // 順序數字
	Arrange int `json:"arrange"` // 排列順序(1:大->小, 2:小->大)
}

type GameListResponse struct {
	OrderList OrderSetting `json:"order_setting"` //遊戲icon排列順序(json)
	GameList  GSGameList   `json:"gamelist"`
	RoomList  GSRoomList   `json:"roomlist"`
}

type AgentGameListRequest struct {
	AgentId int `json:"agent_id,omitempty"` //代理編號
	GameId  int `json:"game_id,omitempty"`  //遊戲id(唯一)
}
