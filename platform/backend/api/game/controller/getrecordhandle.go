package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/pkg/redis"
	"database/sql"
)

/*
	3.2.7 查询游戏注单
		此接口用来获取游戏对局注单
*/
func CheckGameRecord(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest) (bool, *model.ChannelHandleResponse, error) {

	return true,
		&model.ChannelHandleResponse{
			M: model.GetRecordHandle_Path,
			S: 0,
			D: `{"s":106,"m":"/getRecordHandle","d":{"list":{"GameID":["062036007452964330-255"],
		"Accounts":["test"],"ServerID":[3602],"KindID":[620],"TableID":[1],"ChairID":[3],
		"UserCount":[2],"CardValue":["0709292a0000000000000000252b0000211104281d181a"],
		"CellScore":[800],"AllBet":[800],"Profit":[760],"Revenue":[40],"GameStartTime":["2017-04-21 14:41:25"],
		"GameEndTime":[" 2017-04-21 14:45:25"],"ChannelID":[10001]},"count":1,"code":0,"start":1490590800,"end":1490598000}}`,
		}, nil
}
