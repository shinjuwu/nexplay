package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemRecordApi struct {
	BasePath string
}

func NewSystemRecordApi(basePath string) api_cluster.IApiEach {
	return &SystemRecordApi{
		BasePath: basePath,
	}
}

func (p *SystemRecordApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemRecordApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/getuserplayloglist", ginHandler.Handle(p.GetUserPlayLogList))
	g.POST("/getbatchuserplayloglist", ginHandler.Handle(p.GetBatchUserPlayLogList))
	g.POST("/getplaylogcommon", ginHandler.Handle(p.GetPlayLogCommon))
	g.POST("/getwalletledgerlist", ginHandler.Handle(p.GetWalletLedgerList))
	g.POST("/confirmwalletledger",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.ConfirmWalletLedger))
	g.POST("/getagentwalletledgerlist", ginHandler.Handle(p.GetAgentWalletLedgerList))
	g.POST("/getbackendactionloglist", ginHandler.Handle(p.GetBackendActionLogList))
	g.POST("/getbackendloginloglist", ginHandler.Handle(p.GetBackendLoginLogList))
	g.POST("/getautoriskcontrolloglist", ginHandler.Handle(p.GetAutoRiskControlLogList))
	g.POST("/getagentgameratiostatlist", ginHandler.Handle(p.GetAgentGameRatioStatList))
	g.POST("/getusercreditloglist", ginHandler.Handle(p.GetUserCreditLogList))
	g.POST("/getgameusersstathourlist", ginHandler.Handle(p.GetGameUsersStatHourList))
	g.POST("/getfriendroomloglist", ginHandler.Handle(p.GetFriendRoomLogList))
}

/*
遊戲報表相關

輸贏報表
"指定不同時間、代理、遊戲、房間類型、用戶名、局號去篩選遊戲進行後的盈虧數據
（帳變時間、代理代碼、用戶名、遊戲、房間類型、初始金額、有效投注、押分、得分、局號）"

遊戲日誌結果
"指定局號的遊戲日誌
（局號、用戶名、遊戲、房間類型、遊戲詳情）"

上下分紀錄
"指定不同時間、代理、用戶名、帳變類型去篩選上下分的紀錄
（訂單號、帳變時間、代理代碼、用戶名、帳變類型、帳變前金額、帳變金額、帳變後金額）"
*/

// @Tags 報表管理/輸贏報表
// @Summary 取得個人遊戲紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetUserPlayLogListRequest true "取得個人遊戲紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetUserPlayLogResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getuserplayloglist [post]
func (p *SystemRecordApi) GetUserPlayLogList(c *ginweb.Context) {
	var req model.GetUserPlayLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.WinloseReportTimeRange,
		serverInfoSetting.WinloseReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	resp := make([]*model.GetUserPlayLogResponse, 0)

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_roomType := (req.RoomType > definition.ROOM_TYPE_ALL)
	checkExist_userName := (req.UserName != "")

	checkExist_betId := (req.BetId != "")
	checkExist_logNumber := (req.LogNumber != "")
	checkExist_singleWalletId := (req.SingleWalletId != "")
	checkExist_roomId := (req.RoomId != "")

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))

	tableGameId := definition.GAME_ID_ALL
	tableStartTime := today.Add(-time.Duration(serverInfoSetting.WinloseReportTimeBeforeDays) * 24 * time.Hour)
	tableEndTime := time.Now().UTC()

	sqAnd := sq.And{}
	if checkExist_betId {
		sqAnd = append(sqAnd, sq.Eq{"bet_id": req.BetId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else if checkExist_logNumber {
		// Notice: 局號規則 `${gameId}-${19長度數字}`
		strSplit := strings.Split(req.LogNumber, "-")
		if len(strSplit) != 2 || len(strSplit[1]) != 19 {
			c.Ok(resp)
			return
		}

		gameId := utils.ToInt(strSplit[0])
		game := global.GameCache.Get(gameId)
		if game == nil {
			c.Ok(resp)
			return
		}

		sqAnd = append(sqAnd, sq.Eq{"lognumber": req.LogNumber})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}

		tableGameId = game.Id
	} else if checkExist_singleWalletId {
		sqAnd = append(sqAnd, sq.Eq{"wallet_ledger_id": req.SingleWalletId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else if checkExist_roomId {
		sqAnd = append(sqAnd, sq.Eq{"room_id": req.RoomId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else {
		// 非自己或下層級的代理不能查詢
		if checkExist_agentId {
			agent := global.AgentCache.Get(req.AgentId)
			if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}

			sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
		} else {
			// 營運商額外搜尋有問題的訂單
			if userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN {
				sqAnd = append(sqAnd, sq.Or{
					sq.Like{"level_code": userAgent.LevelCode + "%"},
					sq.Eq{"level_code": ""},
				})
			} else {
				sqAnd = append(sqAnd, sq.Like{"level_code": userAgent.LevelCode + "%"})
			}
		}

		if checkExist_gameId {
			game := global.GameCache.Get(req.GameId)
			if game == nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}

			tableGameId = game.Id
		}

		if checkExist_roomType {
			sqAnd = append(sqAnd, sq.Eq{"room_type": req.RoomType})
		}

		if checkExist_userName {
			sqAnd = append(sqAnd, sq.Eq{"username": req.UserName})
		}

		tableStartTime = req.StartTime
		tableEndTime = req.EndTime
	}

	tableName := generateUserPlayLogListTableName(tableGameId, &tableStartTime, &tableEndTime)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("tmp.agent_id", "tmp.game_id", "tmp.room_type", "tmp.desk_id", "tmp.seat_id",
			"tmp.valid_score", "tmp.start_score", "tmp.end_score", "tmp.de_score", "tmp.ya_score",
			"tmp.bet_time", "tmp.username", "tmp.lognumber", "tmp.bet_id", "tmp.kill_type",
			"tmp.tax", "tmp.bonus", "tmp.jp_inject_water_rate", "tmp.jp_inject_water_score", "tmp.wallet_ledger_id",
			"tmp.kill_prob", "tmp.kill_level", "tmp.real_players", "tmp.room_id").
		From(tableName).
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	defer rows.Close()

	tmpAgents := make(map[int]*table_model.Agent)

	for rows.Next() {
		tmp := &model.GetUserPlayLogResponse{}
		if err = rows.Scan(&tmp.AgentId, &tmp.GameId, &tmp.RoomType, &tmp.DeskId, &tmp.SeatId,
			&tmp.ValidScore, &tmp.StartScore, &tmp.EndScore, &tmp.DeScore, &tmp.YaScore,
			&tmp.BetTime, &tmp.UserName, &tmp.LogNumber, &tmp.BetId, &tmp.KillType,
			&tmp.Tax, &tmp.Bonus, &tmp.JpInjectWaterRate, &tmp.JpInjectWaterScore, &tmp.WalletLedgerId,
			&tmp.KillProb, &tmp.KillLevel, &tmp.RealPlayers, &tmp.RoomId); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, tmp)
	}

	c.Ok(resp)
}

// @Tags 報表管理/輸贏報表
// @Summary 分批取得個人遊戲紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetBatchUserPlayLogListRequest true "取得個人遊戲紀錄列表參數"
// @Success 200 {object} response.Response{data=model.GetBatchUserPlayLogListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getbatchuserplayloglist [post]
func (p *SystemRecordApi) GetBatchUserPlayLogList(c *ginweb.Context) {
	var req model.GetBatchUserPlayLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.WinloseReportTimeRange,
		serverInfoSetting.WinloseReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	resp := &model.GetBatchUserPlayLogListResponse{}
	resp.Draw = req.Draw
	resp.Data = make([]interface{}, 0)

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_roomType := (req.RoomType > definition.ROOM_TYPE_ALL)
	checkExist_userName := (req.UserName != "")

	checkExist_betId := (req.BetId != "")
	checkExist_logNumber := (req.LogNumber != "")
	checkExist_singleWalletId := (req.SingleWalletId != "")
	checkExist_roomId := (req.RoomId != "")

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))

	tableGameId := definition.GAME_ID_ALL
	tableStartTime := today.Add(-time.Duration(serverInfoSetting.WinloseReportTimeBeforeDays) * 24 * time.Hour)
	tableEndTime := time.Now().UTC()

	// 搜尋順序
	// 單號 > 局號 > 單一錢包識別碼 > 房間編號 > other
	sqAnd := sq.And{}
	if checkExist_betId {
		sqAnd = append(sqAnd, sq.Eq{"bet_id": req.BetId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else if checkExist_logNumber {
		// Notice: 局號規則 `${gameId}-${19長度數字}`
		strSplit := strings.Split(req.LogNumber, "-")
		if len(strSplit) != 2 || len(strSplit[1]) != 19 {
			c.Ok(resp)
			return
		}

		gameId := utils.ToInt(strSplit[0])
		game := global.GameCache.Get(gameId)
		if game == nil {
			c.Ok(resp)
			return
		}

		sqAnd = append(sqAnd, sq.Eq{"lognumber": req.LogNumber})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}

		tableGameId = game.Id
	} else if checkExist_singleWalletId {
		sqAnd = append(sqAnd, sq.Eq{"wallet_ledger_id": req.SingleWalletId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else if checkExist_roomId {
		sqAnd = append(sqAnd, sq.Eq{"room_id": req.RoomId})
		if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else {
		// 非自己或下層級的代理不能查詢
		if checkExist_agentId {
			agent := global.AgentCache.Get(req.AgentId)
			if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}

			// 指定代理搜尋異常訂單必定找不到，所以可以直接回傳空的結果
			if req.BetslipStatus == 2 {
				c.Ok(resp)
				return
			}

			sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
		} else {
			// 營運商額外搜尋有問題的訂單
			if userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN {
				sqOr := sq.Or{}
				// 找異常的訂單
				if req.BetslipStatus != 1 {
					sqOr = append(sqOr, sq.Eq{"level_code": ""})
				}
				// 找正常的訂單
				if req.BetslipStatus != 2 {
					sqOr = append(sqOr, sq.Like{"level_code": userAgent.LevelCode + "%"})
				}
				sqAnd = append(sqAnd, sqOr)
			} else {
				sqAnd = append(sqAnd, sq.Like{"level_code": userAgent.LevelCode + "%"})
			}
		}

		if checkExist_gameId {
			game := global.GameCache.Get(req.GameId)
			if game == nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}

			tableGameId = game.Id
		}

		if checkExist_roomType {
			sqAnd = append(sqAnd, sq.Eq{"room_type": req.RoomType})
		}

		if checkExist_userName {
			sqAnd = append(sqAnd, sq.Eq{"username": req.UserName})
		}

		tableStartTime = req.StartTime
		tableEndTime = req.EndTime
	}

	tableName := generateUserPlayLogListTableName(tableGameId, &tableStartTime, &tableEndTime)

	// 限制只有第一次會進行加總
	if req.Draw == 0 {
		recordsTotal, totalKillGames, totalDiveGames, totalPlayerWinGames, totalValidScore,
			totalPlatformWinloseScore, totalJpInjectWaterScore, errorCode := getBatchUserPlayLogListSumInfo(c.DB(), tableName, &sqAnd)
		if errorCode != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(errorCode)
			return
		}

		resp.RecordsTotal = recordsTotal
		resp.TotalKillGames = totalKillGames
		resp.TotalDiveGames = totalDiveGames
		resp.TotalValidScore = totalValidScore
		resp.TotalPlatformWinloseScore = totalPlatformWinloseScore
		resp.TotalPlayerWinGames = totalPlayerWinGames
		resp.TotalJpInjectWaterScore = totalJpInjectWaterScore
	}

	data, errorCode := getBatchUserPlayLogList(c.DB(), &req, tableName, &sqAnd)
	if errorCode != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(errorCode)
		return
	}

	for i := 0; i < len(data); i++ {
		resp.Data = append(resp.Data, data[i])
	}

	c.Ok(resp)
}

// @Tags 運營管理/遊戲日誌解析
// @Summary 取得遊戲局記錄
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetPlayLogCommonRequest true "取得遊戲局記錄參數"
// @Success 200 {object} response.Response{data=model.GetPlayLogCommonResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getplaylogcommon [post]
func (p *SystemRecordApi) GetPlayLogCommon(c *ginweb.Context) {
	var req model.GetPlayLogCommonRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	retVal := new(model.GetPlayLogCommonResponse)

	// Notice: 局號規則 `${gameId}-${19長度數字}`
	strSplit := strings.Split(req.LogNumber, "-")
	if len(strSplit) != 2 || len(strSplit[1]) != 19 {
		c.Ok(nil)
		return
	}

	gameId := utils.ToInt(strSplit[0])
	game := global.GameCache.Get(gameId)
	if game == nil {
		c.Ok(nil)
		return
	}

	gameLog := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("lognumber", "game_id", "room_type", "playlog", "bet_time",
			"start_time", "end_time").
		From("play_log_common").
		Where(sq.Eq{"lognumber": req.LogNumber}).
		Limit(1)

	query, args, err := gameLog.ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	err = c.DB().QueryRow(query, args...).Scan(&retVal.LogNumber, &retVal.GameId, &retVal.RoomType, &retVal.PlayLogBytes, &retVal.BetTime,
		&retVal.StartTime, &retVal.EndTime)
	if err != nil && err != sql.ErrNoRows {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	} else if err == sql.ErrNoRows {
		c.Ok(nil)
		return
	}

	retVal.GameCode = game.Code

	json.Unmarshal(retVal.PlayLogBytes, &retVal.PlayLog)

	otherAgentUsers := 0
	tmpPlayerLogs := make([]map[string]interface{}, 0)
	playerLogs := retVal.PlayLog["playerlog"].([]interface{})

	for _, iplayerLog := range playerLogs {
		playerLog := iplayerLog.(map[string]interface{})
		playerIsRobot := playerLog["is_robot"].(float64)

		playerUserId := playerLog["user_id"].(float64)
		playerAgentId := global.AgentDataOfGameUserCache.Get(int(playerUserId))
		var playerAgent *table_model.Agent
		if playerAgentId != nil {
			playerAgent = global.AgentCache.Get(playerAgentId.AgentId)
		}

		// 非棋牌類要進行玩家身分檢查
		if game.Type != definition.GAME_TYPE_CHIPAI {
			// 機器人不顯示
			if playerIsRobot == 1 {
				continue
			}

			// 非管理者檢查玩家是否屬於自身玩家
			if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN &&
				(playerAgent == nil || !strings.HasPrefix(playerAgent.LevelCode, userAgent.LevelCode)) {
				continue
			}
		}

		// Notice: 機器人或是有問題的玩家 playerAgent 為 nil
		if playerAgent == nil || !strings.HasPrefix(playerAgent.LevelCode, userAgent.LevelCode) {
			playerUserName := playerLog["username"].(string)
			playerLog["username"] = "****" + playerUserName[len(playerUserName)-3:]
			otherAgentUsers++
		}

		// 刪除不需要提供的重要資訊
		delete(playerLog, "user_id")
		delete(playerLog, "is_robot")

		tmpPlayerLogs = append(tmpPlayerLogs, playerLog)
	}

	// 非開發商該局沒有屬於自代理底下的玩家則查無資訊
	if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN && otherAgentUsers == len(tmpPlayerLogs) {
		c.Ok(nil)
		return
	}

	retVal.PlayLog["playerlog"] = tmpPlayerLogs

	c.Ok(retVal)
}

// @Tags 報表管理/玩家分數紀錄
// @Summary 取得玩家分數紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetWalletLedgerListRequest true "取得玩家分數紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetWalletLedgerResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getwalletledgerlist [post]
func (p *SystemRecordApi) GetWalletLedgerList(c *ginweb.Context) {
	var req model.GetWalletLedgerListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(
		serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange*24,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// 必要條件檢查
	checkExist_id := (req.Id != "")
	checkExist_single_wallet_id := (req.SingleWalletId != "")
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_userName := (req.UserName != "")
	checkExist_kind := (req.Kind > definition.WALLET_LEDGER_KIND_ALL)

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))
	startTime := today.Add(-time.Duration(serverInfoSetting.CommonReportTimeBeforeDays) * 24 * time.Hour)
	endTime := time.Now().UTC()

	userClaims := c.GetUserClaims()
	meLevelCode := userClaims.LevelCode
	isAdmin := userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN

	sqAnd := sq.And{}
	// 輸入id時只用id進行查詢
	// 此要選擇輸入 single_wallet_id 時只用single_wallet_id進行查詢
	// id > single_wallet_id > other
	if checkExist_id {
		sqAnd = append(sqAnd, sq.Eq{"wl.id": req.Id})
		if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"a.level_code": meLevelCode + "%"})
		}
	} else if checkExist_single_wallet_id {
		sqAnd = append(sqAnd, sq.Eq{"wl.single_wallet_id": req.SingleWalletId})
		if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"a.level_code": meLevelCode + "%"})
		}
	} else {
		if checkExist_userName {
			sqAnd = append(sqAnd, sq.Eq{"wl.username": req.UserName})
		}

		if checkExist_agentId {
			sqAnd = append(sqAnd, sq.Eq{"wl.agent_id": req.AgentId})
		} else if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"a.level_code": meLevelCode + "%"})
		}

		if checkExist_kind {
			sqAnd = append(sqAnd, sq.Eq{"wl.kind": req.Kind})
		}

		startTime = req.StartTime
		endTime = req.EndTime
	}

	sqAnd = append(sqAnd, sq.GtOrEq{"wl.create_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"wl.create_time": endTime})

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("wl.id", "wl.agent_id", "a.name", "wl.username", "wl.create_time",
			"wl.kind", "wl.changeset", "wl.creator", "wl.status", "wl.error_code",
			"wl.info", "wl.request", "wl.single_wallet_id").
		From("wallet_ledger AS wl").
		InnerJoin("agent AS a ON wl.agent_id = a.id").
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	type WalletLedgerRequestParamMap struct {
		Money []string `json:"money"`
	}
	type WalletLedgerRequest struct {
		ParamMap interface{} `json:"param_map"`
	}

	resp := make([]*model.GetWalletLedgerResponse, 0)
	for rows.Next() {
		var requestBytes []byte

		var temp model.GetWalletLedgerResponse
		if err := rows.Scan(&temp.Id, &temp.AgentId, &temp.AgentName, &temp.UserName,
			&temp.CreateTime, &temp.Kind, &temp.ChangeSet, &temp.Creator, &temp.Status,
			&temp.ErrorCode, &temp.Info, &requestBytes, &temp.SingleWalletId); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		// API 上下分不提供Info資訊
		if temp.Kind >= definition.WALLET_LEDGER_KIND_API_UP &&
			temp.Kind <= definition.WALLET_LEDGER_KIND_API_DOWN {
			temp.Info = ""
		}

		var reqeust WalletLedgerRequest
		json.Unmarshal(requestBytes, &reqeust)

		var paramMap WalletLedgerRequestParamMap
		json.Unmarshal([]byte(utils.ToJSON(reqeust.ParamMap)), &paramMap)
		if len(paramMap.Money) > 0 {
			if coinAmount, err := strconv.ParseFloat(paramMap.Money[0], 64); err == nil {
				temp.CoinAmount = coinAmount
			}
		}

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 報表管理/玩家分數紀錄
// @Summary 更新玩家分數紀錄狀態
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ConfirmWalletLedgerRequest true "更新玩家分數紀錄狀態參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/confirmwalletledger [post]
func (p *SystemRecordApi) ConfirmWalletLedger(c *ginweb.Context) {
	var req model.ConfirmWalletLedgerRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("wallet_ledger").
		SetMap(sq.Eq{
			"status":      definition.WALLET_LEDGER_STATUS_SUCCESS,
			"update_time": time.Now().UTC(),
		}).
		Where(sq.And{
			sq.Eq{"id": req.Id},
			sq.Eq{"agent_id": c.GetUserClaims().BaseClaims.ID},
		}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Id)
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 報表管理/代理分數紀錄
// @Summary 取得代理分數紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentWalletLedgerListRequest true "取得代理分數紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentWalletLedgerResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getagentwalletledgerlist [post]
func (p *SystemRecordApi) GetAgentWalletLedgerList(c *ginweb.Context) {
	req := &model.GetAgentWalletLedgerListRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange*24,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))
	isAdmin := userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN

	// 必要條件檢查
	checkExist_id := (req.Id != "")
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))
	startTime := today.Add(-time.Duration(serverInfoSetting.CommonReportTimeBeforeDays) * 24 * time.Hour)
	endTime := time.Now().UTC()

	sqAnd := sq.And{}
	if checkExist_id {
		sqAnd = append(sqAnd, sq.Eq{"awl.id": req.Id})
		if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"a.level_code": userAgent.LevelCode + "%"})
		}
	} else {
		if checkExist_agentId {
			reqAgent := global.AgentCache.Get(int(req.AgentId))
			if reqAgent == nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			} else if !strings.HasPrefix(reqAgent.LevelCode, userAgent.LevelCode) {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			}

			sqAnd = append(sqAnd, sq.Eq{"awl.agent_id": req.AgentId})
		} else if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"a.level_code": userAgent.LevelCode + "%"})
		}

		startTime = req.StartTime
		endTime = req.EndTime
	}
	sqAnd = append(sqAnd, sq.GtOrEq{"awl.update_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"awl.update_time": endTime})

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("awl.id", "a.name", "awl.update_time", "awl.kind", "awl.changeset", "awl.creator", "awl.info").
		From("agent_wallet_ledger AS awl").
		InnerJoin("agent AS a ON awl.agent_id = a.id").
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetAgentWalletLedgerResponse, 0)
	for rows.Next() {
		var temp model.GetAgentWalletLedgerResponse

		if err := rows.Scan(&temp.Id, &temp.AgentName, &temp.UpdateTime, &temp.Kind,
			&temp.ChangeSet, &temp.Creator, &temp.Info); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 運營管理/後台操作紀錄
// @Summary 取得後台操作紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetBackendActionLogListRequest true "取得後台操作紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetBackendActionLogResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getbackendactionloglist [post]
func (p *SystemRecordApi) GetBackendActionLogList(c *ginweb.Context) {
	req := &model.GetBackendActionLogListRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	reqAgent := global.AgentCache.Get(int(req.AgentId))
	if reqAgent == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	} else if !strings.HasPrefix(reqAgent.LevelCode, userAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	userAgentAdmin := global.AdminUserCache.Get(userAgent.Id, userAgent.AdminUsername)
	userAgentAdminAgentPermission := global.AgentPermissionCache.Get(userAgentAdmin.PermissionId)

	wantFeatureCodes := len(req.FeatureCodes) > 0
	req.FeatureCodes = utils.ArrayIntersection(req.FeatureCodes, userAgentAdminAgentPermission.Permission.List)
	if wantFeatureCodes && len(req.FeatureCodes) == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	checkExist_actionType := (req.ActionType > definition.ACTION_LOG_TYPE_ALL)
	checkExist_featureCodes := (len(req.FeatureCodes) > 0)

	sqAnd := sq.And{}
	if checkExist_actionType {
		sqAnd = append(sqAnd, sq.Eq{"action_type": req.ActionType})
	}
	if checkExist_featureCodes {
		if len(req.FeatureCodes) == 1 {
			sqAnd = append(sqAnd, sq.Eq{"feature_code": req.FeatureCodes[0]})
		} else {
			sqAnd = append(sqAnd, sq.Eq{"feature_code": req.FeatureCodes})
		}
	}
	sqAnd = append(sqAnd, sq.Eq{"agent_id": reqAgent.Id})
	sqAnd = append(sqAnd, sq.Eq{"error_code": definition.ERROR_CODE_SUCCESS})
	sqAnd = append(sqAnd, sq.GtOrEq{"create_time": req.StartTime})
	sqAnd = append(sqAnd, sq.Lt{"create_time": req.EndTime})

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "create_time", "username", "action_type", "action_log", "feature_code").
		From("admin_user_backend_action_log").
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetBackendActionLogResponse, 0)
	for rows.Next() {
		var temp model.GetBackendActionLogResponse

		if err := rows.Scan(&temp.Id, &temp.CreateTime, &temp.Username, &temp.ActionType,
			&temp.ActionLog, &temp.FeatureCode); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 系統管理/後台登入紀錄
// @Summary 取得後台登入紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetBackendLoginLogListRequest true "取得後台登入紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetBackendLoginLogListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getbackendloginloglist [post]
func (p *SystemRecordApi) GetBackendLoginLogList(c *ginweb.Context) {
	req := &model.GetBackendLoginLogListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.EarningReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	reqAgent := global.AgentCache.Get(int(req.AgentId))
	if reqAgent == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	} else if !strings.HasPrefix(reqAgent.LevelCode, userAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
	sqAnd = append(sqAnd, sq.GtOrEq{"login_time": req.StartTime})
	sqAnd = append(sqAnd, sq.Lt{"login_time": req.EndTime})

	if req.UserName != "" {
		sqAnd = append(sqAnd, sq.Eq{"username": req.UserName})
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("login_time", "agent_id", "username", "ip", "error_code").
		From("backend_login_log").
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	resp := make([]*model.GetBackendLoginLogListResponse, 0)
	for rows.Next() {
		var temp model.GetBackendLoginLogListResponse

		if err := rows.Scan(&temp.LoginTime, &temp.AgentId, &temp.UserName, &temp.Ip, &temp.ErrorCode); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 風控管理/自動風控紀錄
// @Summary 取得自動風控列表(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAutoRiskControlLogListRequest true "取得自動風控列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAutoRiskControlLogResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getautoriskcontrolloglist [post]
func (p *SystemRecordApi) GetAutoRiskControlLogList(c *ginweb.Context) {
	req := &model.GetAutoRiskControlLogListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// 只有開發商（營運商）可使用
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.GtOrEq{"l.create_time": req.StartTime})
	sqAnd = append(sqAnd, sq.Lt{"l.create_time": req.EndTime})

	if req.AgentId > definition.AGENT_ID_ALL {
		sqAnd = append(sqAnd, sq.Eq{"l.agent_id": req.AgentId})
	}
	if req.UserName != "" {
		sqAnd = append(sqAnd, sq.Eq{"l.username": req.UserName})
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("l.agent_id", "a.name", "l.user_id", "l.username", "l.risk_code", "l.create_time").
		From("auto_risk_control_log AS l").
		InnerJoin(`agent AS a ON a."id" = l.agent_id`).
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_QUERY_COMBITION)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	resp := make([]*model.GetAutoRiskControlLogResponse, 0)
	for rows.Next() {
		var temp model.GetAutoRiskControlLogResponse

		if err := rows.Scan(&temp.AgentId, &temp.AgentName, &temp.UserId, &temp.UserName, &temp.RiskCode, &temp.CreateTime); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 報表管理/日結算報表
// @Summary 取得日結算報表列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentGameRatioStatListRequest true "取得日結算報表列表參數"
// @Success 200 {object} response.Response{data=model.GetAgentGameRatioStatListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getagentgameratiostatlist [post]
func (p *SystemRecordApi) GetAgentGameRatioStatList(c *ginweb.Context) {
	var req model.GetAgentGameRatioStatListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeRange, serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	resp := &model.GetAgentGameRatioStatListResponse{}
	resp.Draw = req.Draw
	resp.Data = make([]interface{}, 0)

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_roomType := (req.RoomType > definition.ROOM_TYPE_ALL)

	args := make([]interface{}, 0)
	queryAnd := make([]string, 0)
	// 非自己或下層級的代理不能查詢
	if checkExist_agentId {
		agent := global.AgentCache.Get(req.AgentId)
		if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		queryAnd = append(queryAnd, "agent_id = $%d")
		args = append(args, req.AgentId)
	} else {
		queryAnd = append(queryAnd, "level_code LIKE $%d")
		args = append(args, userAgent.LevelCode+"%")
	}

	if checkExist_gameId {
		game := global.GameCache.Get(req.GameId)
		if game == nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
		queryAnd = append(queryAnd, "game_id = $%d")
		args = append(args, req.GameId)
	}

	if checkExist_roomType {
		queryAnd = append(queryAnd, "room_type = $%d")
		args = append(args, req.RoomType)
	}

	queryAnd = append(queryAnd, "log_time >= $%d")
	args = append(args, utils.TransUnsignedTimeUTCFormat("15min", req.StartTime))
	queryAnd = append(queryAnd, "log_time <= $%d")
	args = append(args, utils.TransUnsignedTimeUTCFormat("15min", req.EndTime))

	// 產生欲查詢 tablename 並檢查 table 是否存在
	// 取開始到結束時間所有的格式化時間字串
	months := utils.GetTimeIntervalList("month", req.StartTime, req.EndTime)
	checkTablenames := make([]string, 0)
	for i := 0; i < len(months); i++ {
		checkTablenames = append(checkTablenames, "agent_game_ratio_stat"+"_"+months[i])
	}

	// 因為資料表是自動產生，故在此檢查資料表是否存在
	tablenames := make([]string, 0)
	for i := 0; i < len(checkTablenames); i++ {
		var tableExist bool
		query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`

		err := c.DB().QueryRow(query, checkTablenames[i]).Scan(&tableExist)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if tableExist {
			tablenames = append(tablenames, checkTablenames[i])
		}
	}

	// 沒有資料表，直接送空資料
	if len(tablenames) == 0 {
		c.Ok(nil)
		return
	}
	sort.Strings(tablenames)

	// 限制只有第一次會進行加總
	if req.Draw == 0 {
		// args := make([]interface{}, 0)
		// queryAnd := make([]string, 0)
		recordsTotal, totalPlatformWinloseScore, errorCode := getAgentGameRatioStatListSumInfo(c.DB(), tablenames, queryAnd, args)
		if errorCode != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(errorCode)
			return
		}

		resp.RecordsTotal = recordsTotal
		resp.TotalPlatformWinloseScore = totalPlatformWinloseScore
	}

	data, errorCode := getAgentGameRatioStatList(c.DB(), &req, tablenames, queryAnd, args)
	if errorCode != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(errorCode)
		return
	}

	resp.Data = append(resp.Data, data...)

	c.Ok(resp)
}

// @Tags 報表管理/玩家遊玩紀錄
// @Summary 取得玩家遊玩紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUsersStatHourListRequest true "取得玩家遊玩紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetGameUsersStatHourResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getgameusersstathourlist [post]
func (p *SystemRecordApi) GetGameUsersStatHourList(c *ginweb.Context) {
	var req model.GetGameUsersStatHourListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeRange, serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))
	agent := global.AgentCache.Get(req.AgentId)
	if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	req.StartTime = utils.TruncateToHour(req.StartTime)
	req.EndTime = utils.TruncateToHour(req.EndTime)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"a.name", "s.agent_id", "gu.original_username", "s.game_users_id", "s.game_id",
			"SUM(s.play_count)", "SUM(s.de)", "SUM(s.ya)", "SUM(s.tax)", "SUM(s.bonus)",
		).
		From("game_users_stat_hour AS s").
		InnerJoin("agent AS a ON s.agent_id = a.id").
		InnerJoin("game_users AS gu ON s.game_users_id = gu.id").
		Where(sq.And{
			sq.Eq{"s.agent_id": req.AgentId},
			sq.Eq{"gu.original_username": req.Username},
			sq.GtOrEq{"s.log_time": utils.TransUnsignedTimeUTCFormat("hour", req.StartTime)},
			sq.Lt{"s.log_time": utils.TransUnsignedTimeUTCFormat("hour", req.EndTime)},
		}).
		GroupBy("s.agent_id", "s.level_code", "a.name", "gu.original_username", "s.game_users_id", "s.game_id").
		OrderBy("s.level_code asc").
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		return
	}

	defer rows.Close()

	resp := make([]*model.GetGameUsersStatHourResponse, 0)
	for rows.Next() {
		tmp := &model.GetGameUsersStatHourResponse{}
		if err = rows.Scan(&tmp.AgentName, &tmp.AgentId, &tmp.Username, &tmp.UserId, &tmp.GameId,
			&tmp.PlayCount, &tmp.DeScore, &tmp.YaScore, &tmp.Tax, &tmp.Bonus); err != nil {
			return
		}

		resp = append(resp, tmp)
	}

	c.Ok(resp)
}

// GetFriendRoomLogList
// @Tags 報表管理/好友房建房紀錄
// @Summary 取得好友房建房紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetFriendRoomLogListRequest true "取得好友房建房紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetFriendRoomLogListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getfriendroomloglist [post]
func (p *SystemRecordApi) GetFriendRoomLogList(c *ginweb.Context) {
	var req model.GetFriendRoomLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_roomId := (req.RoomId != "")
	checkExist_username := (req.Username != "")

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.GtOrEq{"create_time": req.StartTime})
	sqAnd = append(sqAnd, sq.Lt{"create_time": req.EndTime})

	// 非自己或下層級的代理不能查詢
	if checkExist_agentId {
		agent := global.AgentCache.Get(req.AgentId)
		if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
		sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
	} else if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		sqAnd = append(sqAnd, sq.Like{"level_code": userAgent.LevelCode + "%"})
	}

	if checkExist_gameId {
		game := global.GameCache.Get(req.GameId)
		if game == nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
		sqAnd = append(sqAnd, sq.Eq{"game_id": req.GameId})
	}

	if checkExist_roomId {
		sqAnd = append(sqAnd, sq.Eq{"room_id": req.RoomId})
	}

	if checkExist_username {
		sqAnd = append(sqAnd, sq.Eq{"username": req.Username})
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"id", "agent_id", "game_id", "room_id", "username",
			"tax", "taxpercent", "create_time", "end_time", "detail",
		).
		From("friend_room_log").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetFriendRoomLogListResponse, 0)
	tmpAgents := make(map[int]*table_model.Agent)
	for rows.Next() {
		tmp := &model.GetFriendRoomLogListResponse{}
		if err = rows.Scan(
			&tmp.Id, &tmp.AgentId, &tmp.GameId, &tmp.RoomId, &tmp.Username,
			&tmp.Tax, &tmp.Taxpercent, &tmp.CreateTime, &tmp.EndTime, &tmp.Detail,
		); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, tmp)
	}

	c.Ok(resp)

}

// @Tags 報表管理/玩家帳變紀錄
// @Summary 取得玩家帳變紀錄列表(only 管理可用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetUserCreditLogListRequest true "取得個人遊戲紀錄列表參數"
// @Success 200 {object} response.Response{data=model.GetUserCreditLogListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/record/getusercreditloglist [post]
func (p *SystemRecordApi) GetUserCreditLogList(c *ginweb.Context) {

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	} else {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}
	}

	var req model.GetUserCreditLogListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.WinloseReportTimeRange,
		serverInfoSetting.WinloseReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_userName := (req.UserName != "")

	// 必需指定代理跟用戶查詢
	if !checkExist_agentId || !checkExist_userName {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	tableGameId := definition.GAME_ID_ALL

	agentName := ""
	sqAnd := sq.And{}
	// 非自己或下層級的代理不能查詢
	if checkExist_agentId {
		agent := global.AgentCache.Get(req.AgentId)
		if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
		agentName = agent.Name

		sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
	}

	if checkExist_gameId {
		game := global.GameCache.Get(req.GameId)
		if game == nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		tableGameId = game.Id
	}

	if checkExist_userName {
		sqAnd = append(sqAnd, sq.Eq{"username": req.UserName})
	}

	startTime := req.StartTime
	endTime := req.EndTime

	tableName := generateUserPlayLogListTableName(tableGameId, &startTime, &endTime)

	gameDatas, errorCode := getUserPlayLogList(c.DB(), tableName, &sqAnd)
	if errorCode != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(errorCode)
		return
	}
	resp := make([]*model.GetUserCreditLogListResponse, 0)

	// 轉換遊戲資料
	checkGameDatasDuplicate := make(map[string]bool)
	for _, gameData := range gameDatas {
		tmp := new(model.GetUserCreditLogListResponse)
		tmp.TranGameData(gameData)

		isAdd := true
		if checkGameDatasDuplicate[tmp.CreditId] {
			// 21點特例判斷，同局號注單不記錄
			if tmp.GameId == definition.GAME_ID_BLACKJACK {
				isAdd = false
			}
		} else {
			checkGameDatasDuplicate[tmp.CreditId] = true
		}

		if isAdd {
			resp = append(resp, tmp)
		}
	}

	// 轉換帳變資料
	sqAnd = append(sqAnd, sq.GtOrEq{"create_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"create_time": endTime})

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", fmt.Sprintf(`'%s' as "name"`, agentName), "username", "create_time",
			"kind", "changeset", "creator", "status", "error_code",
			"request", "single_wallet_id").
		From("wallet_ledger").
		Where(sqAnd).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	type WalletLedgerRequestParamMap struct {
		Money []string `json:"money"`
	}
	type WalletLedgerRequest struct {
		ParamMap interface{} `json:"param_map"`
	}

	// walletDate := make([]*model.GetWalletLedgerResponse, 0)
	for rows.Next() {
		var requestBytes []byte

		var temp model.GetWalletLedgerResponse
		if err := rows.Scan(&temp.Id, &temp.AgentId, &temp.AgentName, &temp.UserName,
			&temp.CreateTime, &temp.Kind, &temp.ChangeSet, &temp.Creator, &temp.Status,
			&temp.ErrorCode, &requestBytes, &temp.SingleWalletId); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		var reqeust WalletLedgerRequest
		json.Unmarshal(requestBytes, &reqeust)

		var paramMap WalletLedgerRequestParamMap
		json.Unmarshal([]byte(utils.ToJSON(reqeust.ParamMap)), &paramMap)
		if len(paramMap.Money) > 0 {
			if coinAmount, err := strconv.ParseFloat(paramMap.Money[0], 64); err == nil {
				temp.CoinAmount = coinAmount
			}
		}

		tmp := new(model.GetUserCreditLogListResponse)
		tmp.TranWalletData(&temp)

		resp = append(resp, tmp)
	}

	rows.Close()

	// 轉換JP資料
	sqAnd = sqAnd[:len(sqAnd)-2]
	sqAnd = append(sqAnd, sq.GtOrEq{"winning_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"winning_time": endTime})

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"bet_id", "lognumber", "token_id", "agent_id", "username",
			"jp_bet", "token_create_time", "prize_score", "prize_item", "winning_time",
			"show_pool", "real_pool", "is_robot",
		).
		From("jackpot_log").
		Where(sqAnd).
		ToSql()

	rows, err = c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	for rows.Next() {
		var temp model.GetJackpotListResponse

		if err := rows.Scan(&temp.BetId, &temp.Lognumber, &temp.TokenId, &temp.AgentId, &temp.Username,
			&temp.JpBet, &temp.TokenCreateTime, &temp.PrizeScore, &temp.PrizeItem, &temp.WinningTime,
			&temp.ShowPool, &temp.RealPool, &temp.IsRobot); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		temp.AgentName = agentName
		tmp := new(model.GetUserCreditLogListResponse)
		tmp.TranJPData(&temp)

		resp = append(resp, tmp)
	}
	rows.Close()

	c.Ok(resp)
}
