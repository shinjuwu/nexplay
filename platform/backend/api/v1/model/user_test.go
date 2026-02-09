package model_test

import (
	"backend/api/v1/model"
	"testing"
)

func TestGetGameUserNewbieResponse(t *testing.T) {
	/*
	   type GameUserPlayCountData struct {
	   	GameCode  string `json:"game_code"`  // 遊戲代碼
	   	RoomId    int    `json:"room_id"`    // 房間ID
	   	PlayCount int    `json:"play_count"` // 遊戲局數加總
	   }
	*/
	tmp := new(model.GetGameUserPlayCountDataResponse)
	tmp.UserId = 123
	tmp.TotalNewbieLimit = 100
	jsonStr := `[{"game_code":"t1","room_id":100,"play_count":10},{"game_code":"t2","room_id":200,"play_count":20}]`

	if succ := tmp.DataConvert(jsonStr); !succ {
		t.Error(tmp)
	}
	t.Log(tmp)
}
