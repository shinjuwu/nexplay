package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/internal/statistical"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"sort"
	"strings"
	"time"

	table_model "backend/server/table/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type SystemManageApi struct {
	BasePath string
}

/** create new api group instance
 * basePath: path of base group router
 */
func NewSystemManageApi(basePath string) api_cluster.IApiEach {
	return &SystemManageApi{
		BasePath: basePath,
	}
}

/** get base group path */
func (p *SystemManageApi) GetGroupPath() string {
	return p.BasePath
}

/** 註冊 api */
func (p *SystemManageApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/getmarqueelist", ginHandler.Handle(p.GetMarqueeList))
	g.POST("/getmarquee", ginHandler.Handle(p.GetMarquee))
	g.POST("/createmarquee",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateMarquee))
	g.POST("/deletemarquee",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.DeleteMarquee))
	g.POST("/updatemarquee",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.UpdateMarquee))
	g.GET("/getstatdata", ginHandler.Handle(p.GetStatData))
	g.GET("/getriskuserlist", ginHandler.Handle(p.GetRiskUserList))
	g.GET("/getgameleaderboards", ginHandler.Handle(p.GetGameLeaderboards))
	g.GET("/getintervaldaysbettordata", ginHandler.Handle(p.GetIntervalDaysBettorData))
	g.GET("/getintervaltotalscoredata", ginHandler.Handle(p.GetIntervalTotalScoreData))
	g.GET("/getintervaltotalbettorinfodata", ginHandler.Handle(p.GetIntervalTotalBettorInfoData))
	g.GET("/getintervalrealtimeuserdata", ginHandler.Handle(p.GetIntervalRealtimeUserData))
	g.POST("/getannouncementlist", ginHandler.Handle(p.GetAnnouncementList))
	g.POST("/getannouncement", ginHandler.Handle(p.GetAnnouncement))
	g.POST("/createannouncement",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateAnnouncement))
	g.POST("/deleteannouncement",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.DeleteAnnouncement))
	g.POST("/updateannouncement",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.UpdateAnnouncement))
	g.POST("/getmaintainpagesetting",
		ginHandler.Handle(p.GetMaintainPageSetting))
	g.POST("/setmaintainpagesetting",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetMaintainPageSetting))
}

// for swagger
type GetStatDataDataResponse struct {
	ActivePlayer       int `json:"active_player"`        // 活躍玩家(不重複)
	NumberBettors      int `json:"number_bettors"`       // 投注人數(不重複)
	NumberRegistrants  int `json:"number_registrants"`   // 註冊人數(不重複)
	OddNumber          int `json:"odd_number"`           // 注單數
	TotalBetting       int `json:"total_betting"`        // 總投注
	GameTax            int `json:"game_tax"`             // 遊戲抽水
	PlatformTotalScore int `json:"platform_total_score"` // 平台總輸贏分數
}

// @Tags 運營管理/數據總覽
// @Summary Get statistical data
// @Description 此接口用來取得當天資訊總覽的資料(UTC+0)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param date_type query string true "查詢時間類型"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=[]GetStatDataDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getstatdata [get]
func (p *SystemManageApi) GetStatData(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims != nil {

		/*
			targetLevelCode == "" 為搜尋全部，不管有沒有選擇全部
		*/

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"
		targetLevelCode := c.Request.URL.Query().Get("level_code")
		myLevelCode := claims.BaseClaims.LevelCode

		searchLevelCode := ""
		if targetLevelCode == "" {
			isSearchAll = true
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}

		dateType := c.Request.URL.Query().Get("date_type")
		// 時區計算
		tz := c.Request.URL.Query().Get("time_zone")
		timZoneMin := utils.StringToInt(tz, 0)

		tmpLastTimeMap := make(map[string]interface{})
		tmpThisTimeMap := make(map[string]interface{})

		var timeThisTimeStart, timeThisTimeEnd time.Time
		var timeLastTimeStart, timeLastTimeEnd time.Time

		calDay := 0
		calMonth := 0
		switch dateType {
		case "day":
			timeThisTimeStart, timeThisTimeEnd = utils.GetTimeUTCToday(timZoneMin)
			calDay = 1
		case "week":
			timeThisTimeStart, timeThisTimeEnd = utils.GetTimeUTCThisWeek(timZoneMin)
			calDay = 7
		case "month":
			timeThisTimeStart, timeThisTimeEnd = utils.GetTimeUTCThisMonth(timZoneMin)
			calMonth = 1
		default:
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		timeLastTimeStart = timeThisTimeStart.AddDate(0, -calMonth, -calDay)
		timeLastTimeEnd = timeLastTimeStart.AddDate(0, calMonth, calDay)

		// day example: 2023032016 - 2023032116
		timeStrThisTimeStart := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, timeThisTimeStart)
		timeStrThisTimeEnd := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, timeThisTimeStart.AddDate(0, calMonth, calDay))
		// day example: 2023031916 - 2023032016
		timeStrLastTimeStart := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, timeLastTimeStart)
		timeStrLastTimeEnd := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, timeLastTimeStart.AddDate(0, calMonth, calDay))

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		// 統計各種不重複人數
		// bettors 下注人數
		// active 活躍人數
		// register 註冊人數
		query := fmt.Sprintf(`SELECT 'bettors_yday' AS log_time_range, 
						COUNT(DISTINCT game_users_id)
						FROM game_users_stat_hour
						WHERE log_time >= $1 AND log_time < $2 AND %s
					UNION ALL
					SELECT 'bettors_today' AS log_time_range, 
						COUNT(DISTINCT game_users_id)
						FROM game_users_stat_hour
						WHERE log_time >= $3 AND log_time < $4 AND %s
					UNION ALL
					SELECT 'active_yday' AS log_time_range, COUNT(DISTINCT game_users_id) AS unique_users
					FROM (
						SELECT game_users_id
						FROM game_users_stat_hour
						WHERE log_time >=$5 AND log_time < $6 AND %s
						GROUP BY game_users_id
							UNION ALL
							SELECT game_user_id
						FROM user_login_log
						WHERE login_time >=$7 AND login_time < $8 AND %s
						GROUP BY game_user_id
					) AS active_yday
					UNION ALL
					SELECT 'active_today' AS log_time_range, COUNT(DISTINCT game_users_id) AS unique_users
					FROM (
						SELECT game_users_id
						FROM game_users_stat_hour
						WHERE log_time >= $9 AND log_time < $10 AND %s
						GROUP BY game_users_id
							UNION ALL
							SELECT game_user_id
						FROM user_login_log
						WHERE login_time >=$11 AND login_time < $12 AND %s
						GROUP BY game_user_id
					) AS active_today
					UNION ALL
					SELECT 'register_yday' AS log_time_range, 
						COUNT(DISTINCT id)
						FROM game_users
						WHERE create_time >=$13 AND create_time < $14 AND %s
					UNION ALL
					SELECT 'register_today' AS log_time_range, 
						COUNT(DISTINCT id)
						FROM game_users
						WHERE create_time >=$15 AND create_time < $16 AND %s
					;`,
			searchAllQueryPart, searchAllQueryPart, searchAllQueryPart, searchAllQueryPart, searchAllQueryPart,
			searchAllQueryPart, searchAllQueryPart, searchAllQueryPart)

		rows, err := c.DB().Query(query,
			timeStrLastTimeStart, timeStrLastTimeEnd, // bettors
			timeStrThisTimeStart, timeStrThisTimeEnd,
			timeStrLastTimeStart, timeStrLastTimeEnd, // active
			timeLastTimeStart, timeLastTimeEnd,
			timeStrThisTimeStart, timeStrThisTimeEnd,
			timeThisTimeStart, timeThisTimeEnd,
			timeLastTimeStart, timeLastTimeEnd, // register
			timeThisTimeStart, timeThisTimeEnd)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		for rows.Next() {
			var dbTag string
			var dbCount int
			if err := rows.Scan(&dbTag, &dbCount); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			if dbTag == "register_yday" {
				tmpLastTimeMap["number_registrants"] = dbCount
			} else if dbTag == "register_today" {
				tmpThisTimeMap["number_registrants"] = dbCount
			} else if dbTag == "active_yday" {
				tmpLastTimeMap["active_player"] = dbCount
			} else if dbTag == "active_today" {
				tmpThisTimeMap["active_player"] = dbCount
			} else if dbTag == "bettors_yday" {
				tmpLastTimeMap["number_bettors"] = dbCount
			} else if dbTag == "bettors_today" {
				tmpThisTimeMap["number_bettors"] = dbCount
			}
		}

		rows.Close()

		// 除了不重複人數之外的所有資訊從DB查詢
		/*組合語法查詢; 以下是範例
		SELECT 'yday' AS log_time_range,
		       COALESCE(SUM(play_count), 0) AS bu,
		       COALESCE(SUM(ya), 0) AS ya,
		       COALESCE(SUM(vaild_ya), 0) AS v_ya,
		       COALESCE(SUM(de), 0) AS de,
		       COALESCE(SUM(bonus), 0) AS bonus,
		       COALESCE(SUM(tax), 0) AS tax
		FROM (
		    SELECT *
		    FROM agent_game_ratio_stat_202310
		    WHERE log_time >= '2023093016' AND log_time < '2023100116' AND level_code LIKE '0001%'
		    UNION ALL
		    SELECT *
		    FROM agent_game_ratio_stat_202309
		    WHERE log_time >= '2023093016' AND log_time < '2023100116' AND level_code LIKE '0001%' ) as ydata
		UNION ALL
		SELECT 'tday' AS log_time_range,
		       COALESCE(SUM(play_count), 0) AS bu,
		       COALESCE(SUM(ya), 0) AS ya,
		       COALESCE(SUM(vaild_ya), 0) AS v_ya,
		       COALESCE(SUM(de), 0) AS de,
		       COALESCE(SUM(bonus), 0) AS bonus,
		       COALESCE(SUM(tax), 0) AS tax
		FROM (
		    SELECT *
		    FROM agent_game_ratio_stat_202310
		    WHERE log_time >= '2023100116' AND log_time < '2023100216' AND level_code LIKE '0001%'
		    UNION ALL
		    SELECT *
		    FROM agent_game_ratio_stat_202309
		    WHERE log_time >= '2023100116' AND log_time < '2023100216' AND level_code LIKE '0001%' ) as tdata
		*/

		months := []string{
			timeStrLastTimeStart[0:len("YYYYMM")],
			timeStrLastTimeEnd[0:len("YYYYMM")],
			timeStrThisTimeStart[0:len("YYYYMM")],
			timeStrThisTimeEnd[0:len("YYYYMM")],
		}
		// tablenames
		tablenames := []string{
			"agent_game_ratio_stat" + "_" + timeStrLastTimeStart[0:len("YYYYMM")],
			"agent_game_ratio_stat" + "_" + timeStrLastTimeEnd[0:len("YYYYMM")],
			"agent_game_ratio_stat" + "_" + timeStrThisTimeStart[0:len("YYYYMM")],
			"agent_game_ratio_stat" + "_" + timeStrThisTimeEnd[0:len("YYYYMM")],
		}

		queryCreateTable := `CALL "public"."usp_check_agent_game_ratio_stat"($1)`
		for i := 0; i < len(tablenames); i++ {
			// 沒table 就直接創建
			_, err := c.DB().Exec(queryCreateTable, months[i])
			if err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_EXEC)
				return
			}
		}

		tablenamesIdx := 0
		perTablenamesCount := 2
		queryCount := 2
		querys := make([]string, 0)
		// argsIdx := 1
		// perParamCount := 4
		for i := 0; i < queryCount; i++ {
			tmp := fmt.Sprintf(`SELECT *
				FROM %s
				WHERE log_time >= '%s' AND log_time < '%s' AND %s
				UNION ALL
				SELECT *
				FROM %s
				WHERE log_time >= '%s' AND log_time < '%s' AND %s `,
				tablenames[tablenamesIdx],
				timeStrLastTimeStart,
				timeStrLastTimeEnd,
				searchAllQueryPart,
				tablenames[tablenamesIdx+1],
				timeStrThisTimeStart,
				timeStrThisTimeEnd,
				searchAllQueryPart)
			tablenamesIdx = tablenamesIdx + perTablenamesCount
			querys = append(querys, tmp)
		}

		// trans []string to []any for fmt.Sprintf
		combQuery := make([]any, 0)
		for i := 0; i < len(querys); i++ {
			combQuery = append(combQuery, querys[i])
		}

		query = fmt.Sprintf(`SELECT 'yday' AS log_time_range, 
				COALESCE(SUM(play_count), 0) AS bu, 
				COALESCE(SUM(ya), 0) AS ya, 
				COALESCE(SUM(vaild_ya), 0) AS v_ya, 
				COALESCE(SUM(de), 0) AS de, 
				COALESCE(SUM(bonus), 0) AS bonus, 
				COALESCE(SUM(tax), 0) AS tax
		FROM (
			%s
			) as ydata
		WHERE log_time >= $1 AND log_time < $2
		UNION ALL
		SELECT 'tday' AS log_time_range, 
				COALESCE(SUM(play_count), 0) AS bu, 
				COALESCE(SUM(ya), 0) AS ya, 
				COALESCE(SUM(vaild_ya), 0) AS v_ya, 
				COALESCE(SUM(de), 0) AS de, 
				COALESCE(SUM(bonus), 0) AS bonus, 
				COALESCE(SUM(tax), 0) AS tax
		FROM (
			%s
			) as tdata
		WHERE log_time >= $3 AND log_time < $4
		;`, combQuery...)

		// 補上分鐘的位數
		rows, err = c.DB().Query(query,
			timeStrLastTimeStart+"00",
			timeStrLastTimeEnd+"00",
			timeStrThisTimeStart+"00",
			timeStrThisTimeEnd+"00")
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		for rows.Next() {
			var dbTag string
			var dbBetCount int
			var dbSumYa, dbSumVaildYa, dbSumDe, dbSumBonus, dbSumTax float64
			if err := rows.Scan(&dbTag, &dbBetCount, &dbSumYa, &dbSumVaildYa, &dbSumDe,
				&dbSumBonus, &dbSumTax); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			if dbTag == "yday" {
				tmpLastTimeMap["odd_number"] = dbBetCount
				tmpLastTimeMap["total_betting"] = dbSumYa
				tmpLastTimeMap["game_tax"] = dbSumTax
				tmpLastTimeMap["game_bonus"] = dbSumBonus
				tmpLastTimeMap["platform_total_win_score"] = dbSumYa
				tmpLastTimeMap["platform_total_lose_score"] = dbSumDe
				tmpLastTimeMap["platform_total_score"] = utils.DecimalSub(dbSumYa, dbSumDe)
			} else {
				tmpThisTimeMap["odd_number"] = dbBetCount
				tmpThisTimeMap["total_betting"] = dbSumYa
				tmpThisTimeMap["game_tax"] = dbSumTax
				tmpThisTimeMap["game_bonus"] = dbSumBonus
				tmpThisTimeMap["platform_total_win_score"] = dbSumYa
				tmpThisTimeMap["platform_total_lose_score"] = dbSumDe
				tmpThisTimeMap["platform_total_score"] = utils.DecimalSub(dbSumYa, dbSumDe)
			}
		}

		rows.Close()

		tmps := make(map[string]string, 0)
		tmps["stat_last_time"] = utils.ToJSON(tmpLastTimeMap)
		tmps["stat_this_time"] = utils.ToJSON(tmpThisTimeMap)

		c.Ok(tmps)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetRiskUserListData struct {
	UserId      int     `json:"user_id"`        // 玩家id
	Username    string  `json:"username"`       // 玩家帳號
	TotalYa     float64 `json:"total_ya"`       // 總投注
	TotaVaildYa float64 `json:"total_vaild_ya"` // 總有效投注
	TotalDe     float64 `json:"total_de"`       // 總派獎
	TotalTax    float64 `json:"total_tax"`      // 總抽水
	TotalBonus  float64 `json:"total_bonus"`    // 總紅利
}

// @Tags 運營管理/數據總覽
// @Summary Get risk user list
// @Description 此接口用來取得今日風險玩家清單(前100名)(UTC+0)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=[]GetRiskUserListData,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getriskuserlist [get]
func (p *SystemManageApi) GetRiskUserList(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims != nil {
		myLevelCode := claims.BaseClaims.LevelCode

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

		// 時區計算
		tz := c.Request.URL.Query().Get("time_zone")
		offset := utils.StringToInt(tz, 0)
		startTime, endTime := utils.GetTimeUTCToday(offset)

		startStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, startTime)
		endStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, endTime)

		targetLevelCode := c.Request.URL.Query().Get("level_code")

		searchLevelCode := ""
		if targetLevelCode == "" {
			isSearchAll = true
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		query := ""
		var args []interface{}
		// 找出自己以下的全部代理資料
		// if isAdmin && isSearchAll {
		query = fmt.Sprintf(`SELECT gu.original_username, gush.game_users_id, SUM(gush.ya) AS ya_sum, SUM(gush.vaild_ya) AS vaild_ya_sum, SUM(gush.de) AS de_sum, SUM(gush.tax) AS tax_sum, SUM(gush.bonus) AS bonus_sum
				FROM game_users_stat_hour gush, game_users gu  
				WHERE gush.game_users_id = gu.id AND gush.log_time >= $1 AND gush.log_time < $2  AND gu.%s
				GROUP BY gu.original_username, gush.game_users_id
				HAVING SUM(gush.de) - SUM(gush.ya) > 0
				ORDER BY SUM(gush.de) - SUM(gush.ya) DESC
				LIMIT 100
				;`, searchAllQueryPart)
		args = append(args, startStr, endStr)

		rows, err := c.DB().Query(query, args...)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		defer rows.Close()

		records := make([]*GetRiskUserListData, 0)
		for rows.Next() {
			temp := &GetRiskUserListData{}

			if err := rows.Scan(&temp.Username, &temp.UserId, &temp.TotalYa, &temp.TotaVaildYa, &temp.TotalDe, &temp.TotalTax, &temp.TotalBonus); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			records = append(records, temp)
		}

		c.Ok(records)
		return

	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetGameLeaderboardsData struct {
	GameId       int     `json:"game_id"`        // 遊戲Id
	GameName     string  `json:"game_name"`      // 遊戲名稱
	TotalYa      float64 `json:"total_ya"`       // 總投注
	TotalValidYa float64 `json:"total_valid_ya"` // 總有效投注
	TotalDe      float64 `json:"total_de"`       // 總派獎
	TotalTax     float64 `json:"total_tax"`      // 總抽水
	TotalBonus   float64 `json:"total_bonus"`    // 總紅利
}

// @Tags 運營管理/數據總覽
// @Summary Get game leaderboards
// @Description 此接口用來取得今日遊戲輸贏排行榜(UTC+0)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=[]GetGameLeaderboardsData,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getgameleaderboards [get]
func (p *SystemManageApi) GetGameLeaderboards(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims != nil {
		lbd := make([]*GetGameLeaderboardsData, 0)
		myLevelCode := claims.BaseClaims.LevelCode

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

		// 時區計算
		tz := c.Request.URL.Query().Get("time_zone")
		offset := utils.StringToInt(tz, 0)
		startTime, endTime := utils.GetTimeUTCToday(offset)

		targetLevelCode := c.Request.URL.Query().Get("level_code")

		searchLevelCode := ""
		if targetLevelCode == "" {
			// isSearchAll = true
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		query := ""
		// 找出自己以下的全部代理資料
		gGameDatas := global.GameCache.GetAll()
		gameNames := make([]string, 0)
		for _, game := range gGameDatas {
			if game.Type != definition.GAME_TYPE_LOBBY {
				gameNames = append(gameNames, game.Code)
			}
		}

		mvTablename := generateMVCalGameStatHourSQL(
			fmt.Sprintf("bet_time >= '%s' AND bet_time < '%s' AND %s", startTime.Format(time.RFC3339), endTime.Format(time.RFC3339), searchAllQueryPart),
			gameNames)

		query = fmt.Sprintf(`SELECT tmps.log_time, tmps.game_id, SUM(tmps.sum_ya) as sum_ya, SUM(tmps.sum_valid_ya) as sum_valid_ya, SUM(tmps.sum_de) as sum_de, SUM(tmps.sum_bonus) as sum_bonus, SUM(tmps.sum_tax) as sum_tax
				FROM (%s) tmps
				GROUP BY tmps.log_time ,tmps.game_id
				ORDER BY tmps.log_time DESC;`, mvTablename)

		rows, err := c.DB().Query(query)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		defer rows.Close()

		records := make(map[int]*table_model.RtRealtimeGameHourData, 0)
		recordTime := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_DAY, time.Now().Add(time.Duration(-offset)*time.Minute))
		for rows.Next() {
			temp := &table_model.RtRealtimeGameHourData{}

			var lt time.Time
			if err := rows.Scan(&lt, &temp.GameId, &temp.YaScore, &temp.ValidYaScore, &temp.DeScore,
				&temp.Bonus, &temp.Tax); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			temp.LogTime = recordTime

			record, ok := records[temp.GameId]
			if ok {
				record.YaScore = utils.DecimalAdd(record.YaScore, temp.YaScore)
				record.ValidYaScore = utils.DecimalAdd(record.ValidYaScore, temp.ValidYaScore)
				record.DeScore = utils.DecimalAdd(record.DeScore, temp.DeScore)
				record.Bonus = utils.DecimalAdd(record.Bonus, temp.Bonus)
				record.Tax = utils.DecimalAdd(record.Tax, temp.Tax)
			} else {
				records[temp.GameId] = temp
			}
		}

		// 依照目前有設定開放的遊戲，來設定遊戲即時資料
		for _, gGameData := range gGameDatas {
			// 已關閉遊戲不統計 & 大廳不統計
			if gGameData.State > 0 && gGameData.Type != definition.GAME_TYPE_LOBBY {
				val := &GetGameLeaderboardsData{}
				record, ok := records[gGameData.Id]
				if ok {
					val.GameId = gGameData.Id
					val.GameName = gGameData.Name
					val.TotalYa = record.YaScore
					val.TotalValidYa = record.ValidYaScore
					val.TotalDe = record.DeScore
					val.TotalTax = record.Tax
					val.TotalBonus = record.Bonus
				} else {
					val.GameId = gGameData.Id
					val.GameName = gGameData.Name
					val.TotalYa = 0
					val.TotalValidYa = 0
					val.TotalDe = 0
					val.TotalTax = 0
					val.TotalBonus = 0
				}

				lbd = append(lbd, val)
			}
		}

		tmp := utils.ToJSON(lbd)

		c.Ok(tmp)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetIntervalDaysBettorDataResponse struct {
	Logtime       string `json:"log_time"`       // 日期(YYYYMMDD)
	ActivePlayer  int    `json:"active_player"`  // 活躍玩家(不重複)
	NumberBettors int    `json:"number_bettors"` // 投注人數(不重複)
}

// @Tags 運營管理/數據總覽
// @Summary Get game winning leaderboard
// @Description 此接口用來取得最近一段時間(預設30日)的活躍玩家數&日投注人數(UTC+0)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=[]GetIntervalDaysBettorDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getintervaldaysbettordata [get]
func (p *SystemManageApi) GetIntervalDaysBettorData(c *ginweb.Context) {

	claims := c.GetUserClaims()
	if claims != nil {
		idbdr := make(map[string]*GetIntervalDaysBettorDataResponse, 0)
		myLevelCode := claims.BaseClaims.LevelCode

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

		tz := c.Request.URL.Query().Get("time_zone")
		offset := utils.StringToInt(tz, 0)

		transHour := offset / 60
		transMin := offset % 60

		targetLevelCode := c.Request.URL.Query().Get("level_code")

		searchLevelCode := ""
		if targetLevelCode == "" {
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}
		nowTime := time.Now().UTC()
		startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()-30, transHour, transMin, 0, 0, nowTime.UTC().Location())
		endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), transHour, transMin, 0, 0, nowTime.UTC().Location())

		startTimeStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, startTime)
		endTimeStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, endTime)

		// dateList := utils.GetTimeIntervalList("day", startTime, endTime)
		// 產生實際時區的時間，之後依照時間歸納資料
		dateList := utils.GetTimeIntervalList(statistical.SAVE_PERIOD_DAY, startTime.Add(time.Duration(-offset)*time.Minute), endTime.Add(time.Duration(-offset)*time.Minute))

		if len(dateList) > 1 {
			dateList = dateList[:len(dateList)-1]
		}

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		// 不重複日投注人數計算 start
		caseQuery := "WHEN log_time >= '%s' AND log_time < '%s' THEN '%s'"
		caseQuerys := ""

		for _, timeStr := range dateList {
			targetTime, _ := utils.GetUnsignedTimeUTCFromStr(timeStr, statistical.SAVE_PERIOD_DAY)
			targetTime = targetTime.Add(time.Minute * time.Duration(offset))
			targetStartTimeStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, targetTime)
			targetEndTimeStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, targetTime.AddDate(0, 0, 1))

			caseQuerys += fmt.Sprintf(caseQuery+" ", targetStartTimeStr, targetEndTimeStr, timeStr)
		}
		fromQueryBettors := fmt.Sprintf(`FROM
		game_users_stat_hour
		WHERE
		log_time >= $1 AND log_time < $2 AND %s`, searchAllQueryPart)

		finalQuery := fmt.Sprintf(`SELECT
			CASE
			%s
			END as day_interval,
			COUNT(DISTINCT game_users_id) AS user_count
			%s
			GROUP BY
			CASE
			%s
			END;`, caseQuerys, fromQueryBettors, caseQuerys)

		rows, err := c.DB().Query(finalQuery, startTimeStr, endTimeStr)
		if err != nil && err != sql.ErrNoRows {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		for rows.Next() {
			temp := &GetIntervalDaysBettorDataResponse{}
			if err := rows.Scan(&temp.Logtime, &temp.NumberBettors); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			idbdr[temp.Logtime] = temp
		}

		rows.Close()
		// 不重複下注人數計算 end

		fromQueryActives := fmt.Sprintf(`FROM (
			SELECT game_users_id, log_time
			FROM game_users_stat_hour
			WHERE log_time >=$1 AND log_time < $2 AND %s
			GROUP BY game_users_id, log_time
			UNION ALL
			SELECT game_user_id, to_char(login_time, 'YYYYMMDDhh24') AS log_time
			FROM user_login_log
			WHERE login_time >= to_timestamp($3, 'YYYYMMDDhh24') AND login_time < to_timestamp($4, 'YYYYMMDDhh24') AND %s
			GROUP BY game_user_id, login_time
			) as ttt`, searchAllQueryPart, searchAllQueryPart)

		finalQuery = fmt.Sprintf(`SELECT
			CASE
			%s
			END as day_interval,
			COUNT(DISTINCT game_users_id) AS user_count
			%s
			GROUP BY
			CASE
			%s
			END;`, caseQuerys, fromQueryActives, caseQuerys)

		rows, err = c.DB().Query(finalQuery, startTimeStr, endTimeStr, startTimeStr, endTimeStr)
		if err != nil && err != sql.ErrNoRows {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		for rows.Next() {
			temp := &GetIntervalDaysBettorDataResponse{}
			if err := rows.Scan(&temp.Logtime, &temp.ActivePlayer); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			idbd, ok := idbdr[temp.Logtime]
			if ok {
				idbd.ActivePlayer = temp.ActivePlayer
			} else {
				idbdr[temp.Logtime] = temp
			}
		}

		rows.Close()

		// 沒有資料的日期要補0
		for _, v := range dateList {
			if _, ok := idbdr[v]; !ok {
				tmp := &GetIntervalDaysBettorDataResponse{
					Logtime:       v,
					ActivePlayer:  0,
					NumberBettors: 0,
				}

				idbdr[v] = tmp
			}
		}

		c.Ok(idbdr)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetIntervalTodayTotalScoreDataResponse struct {
	GameId            int     `json:"game_id"`              // 遊戲id
	Gamename          string  `json:"game_name"`            // 遊戲名稱
	Logtime           string  `json:"log_time"`             // 時段
	TotlaDeScore      float64 `json:"total_de_score"`       // 總德分
	TotlaYaScore      float64 `json:"total_ya_score"`       // 總壓分
	TotlaValidYaScore float64 `json:"total_valid_ya_score"` // 總有效壓分
	TotalTaxScore     float64 `json:"total_tax_score"`      // 總抽水
	TotalBonus        float64 `json:"total_bonus"`          // 總紅利
}

// @Tags 運營管理/數據總覽
// @Summary Get today's winners and losers
// @Description 此接口用來取得今日各時段輸贏(昨日&今日)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=map[][]GetIntervalTodayTotalScoreDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getintervaltotalscoredata [get]
func (p *SystemManageApi) GetIntervalTotalScoreData(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims != nil {
		rtmp := make([]*GetIntervalTodayTotalScoreDataResponse, 0)
		myLevelCode := claims.BaseClaims.LevelCode

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

		tz := c.Request.URL.Query().Get("time_zone")
		timZoneMin := utils.StringToInt(tz, 0)

		transHour := timZoneMin / 60
		transMin := timZoneMin % 60

		targetLevelCode := c.Request.URL.Query().Get("level_code")

		searchLevelCode := ""
		if targetLevelCode == "" {
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		query := ""
		calHour := 24 - transHour

		gGameDatas := global.GameCache.GetAll()
		gameNames := make([]string, 0)
		for _, game := range gGameDatas {
			if game.Type != definition.GAME_TYPE_LOBBY {
				gameNames = append(gameNames, game.Code)
			}
		}

		mvTablename := generateMVCalGameStatHourSQL(
			fmt.Sprintf("bet_time >= CURRENT_DATE - '%d hours'::interval AND %s", calHour, searchAllQueryPart),
			gameNames)

		query = fmt.Sprintf(`SELECT tmps.log_time, tmps.game_id, SUM(tmps.sum_ya) as sum_ya, SUM(tmps.sum_valid_ya) as sum_valid_ya, SUM(tmps.sum_de) as sum_de, SUM(tmps.sum_tax) as sum_tax, SUM(tmps.sum_bonus) as sum_bonus
				FROM (%s) tmps
				GROUP BY tmps.log_time, tmps.game_id;`, mvTablename)
		rows, err := c.DB().Query(query)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		defer rows.Close()

		// 時區計算
		nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
		nowTime := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, nowTimeFromTimeZone.UTC().Location())
		startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()-1, transHour, transMin, 0, 0, nowTime.UTC().Location())
		endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, transHour, transMin, 0, 0, nowTime.UTC().Location())
		dateList := utils.GetTimeIntervalList(statistical.SAVE_PERIOD_HOUR, startTime, endTime)

		for rows.Next() {
			temp := &GetIntervalTodayTotalScoreDataResponse{}

			if err := rows.Scan(&temp.Logtime, &temp.GameId, &temp.TotlaYaScore, &temp.TotlaValidYaScore, &temp.TotlaDeScore, &temp.TotalTaxScore, &temp.TotalBonus); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			rtmp = append(rtmp, temp)
		}

		gResult := make(map[int]interface{}, 0)

		for _, gc := range gGameDatas {
			if gc.Type == definition.GAME_TYPE_LOBBY {
				continue
			}

			tmp := make([]*GetIntervalTodayTotalScoreDataResponse, 0)

			dateMap := make(map[string]bool, 0)
			for _, logtime := range dateList {
				for _, rt := range rtmp {
					t, _ := time.Parse("2006-01-02T15:04:05Z", rt.Logtime)
					rtLogtime := utils.GetUnsignedTimeUTC(t, statistical.SAVE_PERIOD_HOUR)
					if rtLogtime == logtime && rt.GameId == gc.Id {
						rt.Gamename = gc.Name
						rt.Logtime = rtLogtime
						tmp = append(tmp, rt)
						dateMap[logtime] = true
					}
				}
			}

			// 沒有資料的日期要補0
			for _, logtime := range dateList {
				if !dateMap[logtime] {
					gg := &GetIntervalTodayTotalScoreDataResponse{
						GameId:   gc.Id,
						Gamename: gc.Name,
						Logtime:  logtime,
					}
					tmp = append(tmp, gg)
					dateMap[logtime] = true
				}
			}

			gResult[gc.Id] = tmp
		}

		c.Ok(gResult)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetIntervalTotalBettorInfoDataResponse struct {
	Logtime       string `json:"log_time"`        // 時段
	GameId        int    `json:"game_id"`         // 遊戲id
	Gamename      string `json:"game_name"`       // 遊戲名稱
	TotalBettor   int    `json:"total_bettor"`    // 總投注人數
	TotalBetCount int    `json:"total_bet_count"` // 總注單數
}

// @Tags 運營管理/數據總覽
// @Summary Get number of bettors for each time period and game today
// @Description 此接口用來取得今日各時段、各遊戲投注人數(昨日&今日)
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=map[][]GetIntervalTotalBettorInfoDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getintervaltotalbettorinfodata [get]
func (p *SystemManageApi) GetIntervalTotalBettorInfoData(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims != nil {
		rtmp := make([]*GetIntervalTotalBettorInfoDataResponse, 0)
		myLevelCode := claims.BaseClaims.LevelCode

		isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

		tz := c.Request.URL.Query().Get("time_zone")
		timZoneMin := utils.StringToInt(tz, 0)

		transHour := timZoneMin / 60
		transMin := timZoneMin % 60

		targetLevelCode := c.Request.URL.Query().Get("level_code")

		searchLevelCode := ""
		if targetLevelCode == "" {
			// isSearchAll = true
			searchLevelCode = myLevelCode
		} else {
			// 檢查是否為自己以下的權限
			isOk := global.CheckTargetLevelCodeIsPassing(myLevelCode, targetLevelCode)
			if !isOk {
				c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
				return
			} else {
				searchLevelCode = targetLevelCode
			}
		}

		searchAllQueryPart := ""
		if isSearchAll {
			searchAllQueryPart = fmt.Sprintf("level_code LIKE '%s'", searchLevelCode+"%")
		} else {
			searchAllQueryPart = fmt.Sprintf("level_code = '%s'", searchLevelCode)
		}

		query := ""
		calHour := 24 - transHour
		gGameDatas := global.GameCache.GetAll()
		gameNames := make([]string, 0)
		for _, game := range gGameDatas {
			if game.Type != definition.GAME_TYPE_LOBBY {
				gameNames = append(gameNames, game.Code)
			}
		}

		mvTablename := generateMVCalGameStatHourSQL(
			fmt.Sprintf("bet_time >= CURRENT_DATE - '%d hours'::interval AND %s", calHour, searchAllQueryPart),
			gameNames)

		query = fmt.Sprintf(`SELECT tmps.log_time, tmps.game_id, SUM(tmps.bet_user) as sum_bet_user, SUM(tmps.bet_count) as sum_bet_count
				FROM (%s) tmps
				GROUP BY tmps.log_time, tmps.game_id;`, mvTablename)
		rows, err := c.DB().Query(query)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		defer rows.Close()

		// 時區計算
		nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
		nowTime := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, nowTimeFromTimeZone.UTC().Location())
		startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()-1, transHour, transMin, 0, 0, nowTime.UTC().Location())
		endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+1, transHour, transMin, 0, 0, nowTime.UTC().Location())
		dateList := utils.GetTimeIntervalList(statistical.SAVE_PERIOD_HOUR, startTime, endTime)

		for rows.Next() {
			temp := &GetIntervalTotalBettorInfoDataResponse{}

			if err := rows.Scan(&temp.Logtime, &temp.GameId, &temp.TotalBettor, &temp.TotalBetCount); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			rtmp = append(rtmp, temp)
		}

		gResult := make(map[int]interface{}, 0)

		for _, gc := range gGameDatas {
			if gc.Type == definition.GAME_TYPE_LOBBY {
				continue
			}

			tmp := make([]*GetIntervalTotalBettorInfoDataResponse, 0)

			dateMap := make(map[string]bool, 0)
			for _, logtime := range dateList {
				for _, rt := range rtmp {
					t, _ := time.Parse("2006-01-02T15:04:05Z", rt.Logtime)
					rtLogtime := utils.GetUnsignedTimeUTC(t, statistical.SAVE_PERIOD_HOUR)
					if rtLogtime == logtime && rt.GameId == gc.Id {
						rt.Gamename = gc.Name
						rt.Logtime = rtLogtime
						tmp = append(tmp, rt)
						dateMap[logtime] = true
					}
				}
			}

			// 沒有資料的日期要補0
			for _, logtime := range dateList {
				if !dateMap[logtime] {
					gg := &GetIntervalTotalBettorInfoDataResponse{
						GameId:   gc.Id,
						Gamename: gc.Name,
						Logtime:  logtime,
					}
					tmp = append(tmp, gg)
					dateMap[logtime] = true
				}
			}

			gResult[gc.Id] = tmp
		}

		c.Ok(gResult)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

type GetIntervalRealtimeUserDataResponse struct {
	Logtime                  string `json:"log_time"`                     // 時段
	TotalOnlineGameUserCount int    `json:"total_online_game_user_count"` // 總人數
}

// @Tags 運營管理/數據總覽
// @Summary Get today and yesterday online people info
// @Description 此接口用來取得今日昨日的各時段線上人數
// @Produce  application/json
// @Security BearerAuth
// @Param level_code query string false "查詢目標層級碼"
// @Param time_zone query int true "時間區域(UTC+X, X時間, 單位分鐘, ex: UTC+8 要傳入 -60*8=-480)"
// @Param is_search_all query int true "查詢是否包含自身以下層級資料(1:是, other:否)"
// @Success 200 {object} response.Response{data=[]GetIntervalRealtimeUserDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getintervalrealtimeuserdata [get]
func (p *SystemManageApi) GetIntervalRealtimeUserData(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	searchLevelCode := c.Request.URL.Query().Get("level_code")
	timZoneMin := utils.StringToInt(c.Request.URL.Query().Get("time_zone"), 0)
	isSearchAll := c.Request.URL.Query().Get("is_search_all") == "1"

	if searchLevelCode == "" {
		searchLevelCode = claims.BaseClaims.LevelCode
	}

	// 檢查是否為自己以下的權限
	if isOk := global.CheckTargetLevelCodeIsPassing(claims.BaseClaims.LevelCode, searchLevelCode); !isOk {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	nowTimeFromTimeZone := time.Now().UTC().Add(time.Duration(-timZoneMin) * time.Minute)
	todayFromTimeZone := time.Date(nowTimeFromTimeZone.Year(), nowTimeFromTimeZone.Month(), nowTimeFromTimeZone.Day(), 0, 0, 0, 0, time.UTC)
	today := todayFromTimeZone.Add(time.Duration(timZoneMin) * time.Minute)
	startTime := today.Add(-time.Hour * 24)
	endTime := today.Add(time.Hour * 24)

	sqAnd := sq.And{
		sq.GtOrEq{"log_time": utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_MINUTE, startTime)},
		sq.Lt{"log_time": utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_MINUTE, endTime)},
	}
	if isSearchAll {
		sqAnd = append(sqAnd, sq.Like{"level_code": searchLevelCode + "%"})
	} else {
		sqAnd = append(sqAnd, sq.Eq{"level_code": searchLevelCode})
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("log_time", "SUM(online_game_user_count)").
		From("agent_game_users_stat_min").
		Where(sqAnd).
		GroupBy("log_time").
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}
	defer rows.Close()

	irudr := make(map[string]struct{})
	resp := make([]*GetIntervalRealtimeUserDataResponse, 0)
	for rows.Next() {
		var temp GetIntervalRealtimeUserDataResponse
		// newline every scan 5 column
		if err := rows.Scan(&temp.Logtime, &temp.TotalOnlineGameUserCount); err != nil {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARSE_ROWS_FAILED)
			return
		}

		resp = append(resp, &temp)
		irudr[temp.Logtime] = struct{}{}
	}

	// 產生實際時區的時間，之後依照時間歸納資料
	dateList := utils.GetTimeIntervalList(statistical.SAVE_PERIOD_MINUTE, startTime, endTime)
	for _, logtime := range dateList {
		if _, find := irudr[logtime]; !find {
			resp = append(resp, &GetIntervalRealtimeUserDataResponse{Logtime: logtime})
		}
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Logtime < resp[j].Logtime
	})

	c.Ok(resp)
}

// @Tags 運營管理/跑馬燈設定
// @Summary Get marquee list
// @Description 取得目前跑馬燈設定列表：開發者、總代理、子代理，皆可查看跑馬燈資訊全內容
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.GetMarqueeResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getmarqueelist [post]
func (p *SystemManageApi) GetMarqueeList(c *ginweb.Context) {

	/*
		開發者、總代理、子代理，皆可查看跑馬燈資訊全內容，故不檢查權限
	*/

	// column: 11
	query := `SELECT "id", "lang", "type", "order", "freq", "is_enabled", "is_open", "content", "start_time", "end_time", "create_time", "update_time"
			FROM marquee 
			WHERE "is_open" = true;`

	rows, err := c.DB().Query(query)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	defer rows.Close()

	resp := make([]*model.GetMarqueeResponse, 0)

	for rows.Next() {
		temp := model.NewEmptyGetMarqueeResponse()
		// newline every scan 5 column
		if err := rows.Scan(&temp.Id, &temp.Lang, &temp.Type, &temp.Order, &temp.Freq,
			&temp.IsEnabled, &temp.IsOpen, &temp.Content, &temp.StartTime, &temp.EndTime,
			&temp.CreateTime, &temp.UpdateTime); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		resp = append(resp, temp)
	}

	c.Ok(resp)
}

// @Tags 運營管理/跑馬燈設定
// @Summary Get marquee
// @Description 指定取得某筆跑馬燈設定
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.GetMarqueeRequest true "id"
// @Success 200 {object} response.Response{data=model.GetMarqueeResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getmarquee [post]
func (p *SystemManageApi) GetMarquee(c *ginweb.Context) {

	/*
		開發者、總代理、子代理，皆可查看跑馬燈資訊全內容，故不檢查權限
	*/

	req := model.NewEmptyGetMarqueeRequest()
	_ = c.ShouldBindJSON(&req)

	// 檢查參數是否合法
	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	id := req.Id

	temp := model.NewEmptyGetMarqueeResponse()
	// column: 11
	query := `SELECT "id", "lang", "type", "order", "freq", "is_enabled", "is_open", "content", "start_time", "end_time", "create_time", "update_time"
			FROM marquee 
			WHERE "is_open" = true AND "id" = $1;`

	err := c.DB().QueryRow(query, id).Scan(&temp.Id, &temp.Lang, &temp.Type, &temp.Order, &temp.Freq,
		&temp.IsEnabled, &temp.IsOpen, &temp.Content, &temp.StartTime, &temp.EndTime,
		&temp.CreateTime, &temp.UpdateTime)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	c.Ok(temp)
}

// @Tags 運營管理/跑馬燈設定
// @Summary create new marquee
// @Description 添加跑馬燈功能（開發者）：管理後台才可添加【活動】類型跑馬燈
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.CreateMarqueeRequest true "id, 創建內容相關參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/createmarquee [post]
func (p *SystemManageApi) CreateMarquee(c *ginweb.Context) {
	req := model.NewEmptyCreateMarqueeRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 跑馬燈新增時不能逾期
	if req.EndTime.Before(time.Now()) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_MARQUEE_OVERDUE)
		return
	}

	// 只能新增活動類型跑馬燈
	// if req.Type != definition.MARQUEE_TYPE_EVENT {
	// 	c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
	// 	return
	// }

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode

	// 最高權限才能新增
	if len(myLevelCode) != 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	isEnabled := true
	isOpen := true

	// 先傳送給game server，傳送成功再新增db，新增db失敗讓使用再重新操作即可
	notifyData := make(map[string]interface{}, 0)

	notifyData["method"] = "create"
	notifyData["id"] = id
	notifyData["lang"] = req.Lang
	notifyData["type"] = req.Type
	notifyData["order"] = req.Order
	notifyData["freq"] = req.Freq
	notifyData["content"] = req.Content
	notifyData["start_time"] = req.StartTime
	notifyData["end_time"] = req.EndTime
	notifyData["is_enabled"] = isEnabled
	notifyData["is_open"] = isOpen

	code, err := notification.SendMarquee(notifyData)
	if err != nil || code != 0 {
		c.Logger().Info("create marquee notification.SendMarquee fail, notifyData=%v, resp code=%d, err=%v", notifyData, code, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("marquee").
		Columns("id", "lang", "\"type\"", "\"order\"", "freq",
			"is_enabled", "is_open", "\"content\"", "start_time", "end_time").
		Values(id, req.Lang, req.Type, req.Order, req.Freq,
			isEnabled, isOpen, req.Content, req.StartTime, req.EndTime).
		ToSql()

	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Content)
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 運營管理/跑馬燈設定
// @Summary update new marquee
// @Description 編輯跑馬燈功能（開發者）：管理後台才可編輯【活動】類型跑馬燈
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.UpdateMarqueeRequest true "id, 更新內容相關參數"
// @Success 200 {object} response.Response{data=model.GetMarqueeResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/updatemarquee [post]
func (p *SystemManageApi) UpdateMarquee(c *ginweb.Context) {
	req := model.NewEmptyUpdateMarqueeRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 跑馬燈修改時不能逾期
	if req.EndTime.Before(time.Now()) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_MARQUEE_OVERDUE)
		return
	}

	// 只能修改活動類型跑馬燈
	// if req.Type != definition.MARQUEE_TYPE_EVENT {
	// 	c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
	// 	return
	// }

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode

	// 最高權限才能新增
	if len(myLevelCode) != 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 先查詢確定有資料再傳送修改資訊給game server
	marqueeLang := ""
	marqueeType := definition.MARQUEE_TYPE_NONE
	marqueeOrder := definition.MARQUEE_ORDER_MIN - 1
	marqueeFreq := -1
	marqueeIsEnabled := false
	marqueeContent := ""
	marqueeStartTime := time.Now().UTC()
	marqueeEndTime := time.Now().UTC()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("lang", "\"type\"", "\"order\"", "freq", "is_enabled", "\"content\"", "start_time", "end_time").
		From("marquee").
		Where(sq.Eq{"id": req.Id}).
		ToSql()

	err := c.DB().QueryRow(query, args...).Scan(&marqueeLang, &marqueeType, &marqueeOrder,
		&marqueeFreq, &marqueeIsEnabled, &marqueeContent, &marqueeStartTime, &marqueeEndTime)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 全部流程由後台這邊控制
	// 如果is_enabled=false，送method="delete"，之後如果改設定is_enabled=true，走下面的流程即可
	// 如果is_enabled=true，送method="update"，如果game server回應資料不存在，重新送method="create"進行建立資料
	// 先傳送給game server，傳送成功再修改db，修改db失敗讓使用再重新操作即可
	notifyData := make(map[string]interface{}, 0)
	notifyData["id"] = req.Id

	if req.IsEnabled {
		notifyData["method"] = "update"
		notifyData["lang"] = req.Lang
		notifyData["type"] = req.Type
		notifyData["order"] = req.Order
		notifyData["freq"] = req.Freq
		notifyData["is_enabled"] = req.IsEnabled
		notifyData["content"] = req.Content
		notifyData["start_time"] = req.StartTime
		notifyData["end_time"] = req.EndTime

		code, err := notification.SendMarquee(notifyData)
		if code != 0 && code != notification.API_ERROR_MARQUEE_DO_NOT_EXIST {
			c.Logger().Info("update marquee notification.SendMarquee fail, notifyData=%v, resp code=%d, err=%v", notifyData, code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}

		if code == notification.API_ERROR_MARQUEE_DO_NOT_EXIST {
			notifyData["method"] = "create"
			code, err = notification.SendMarquee(notifyData)

			if err != nil || code != 0 {
				c.Logger().Info("update recreate marquee notification.SendMarquee fail, notifyData=%v, resp code=%d, err=%v", notifyData, code, err)
				c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
				return
			}
		}
	} else {
		notifyData["method"] = "delete"
		notifyData["is_open"] = false

		code, err := notification.SendMarquee(notifyData)
		if code != 0 && code != notification.API_ERROR_MARQUEE_DO_NOT_EXIST {
			c.Logger().Info("update but delete marquee notification.SendMarquee fail, notifyData=%v, resp code=%d, err=%v", notifyData, code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("marquee").
		Set("lang", req.Lang).
		Set("\"type\"", req.Type).
		Set("\"order\"", req.Order).
		Set("freq", req.Freq).
		Set("is_enabled", req.IsEnabled).
		Set("\"content\"", req.Content).
		Set("start_time", req.StartTime).
		Set("end_time", req.EndTime).
		Set("update_time", time.Now().UTC()).
		Where(sq.Eq{"id": req.Id}).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Content)
	if marqueeLang != req.Lang {
		actionLog["lang"] = createBancendActionLogDetail(marqueeLang, req.Lang)
	}
	if marqueeType != req.Type {
		actionLog["marquee_type"] = createBancendActionLogDetail(marqueeType, req.Type)
	}
	if marqueeOrder != req.Order {
		actionLog["order"] = createBancendActionLogDetail(marqueeOrder, req.Order)
	}
	if marqueeFreq != req.Freq {
		actionLog["freq"] = createBancendActionLogDetail(marqueeFreq, req.Freq)
	}
	if marqueeIsEnabled != req.IsEnabled {
		actionLog["bool_status"] = createBancendActionLogDetail(marqueeIsEnabled, req.IsEnabled)
	}
	if !marqueeStartTime.Equal(req.StartTime) {
		actionLog["start_time"] = createBancendActionLogDetail(marqueeStartTime, req.StartTime)
	}
	if !marqueeEndTime.Equal(req.EndTime) {
		actionLog["end_time"] = createBancendActionLogDetail(marqueeEndTime, req.EndTime)
	}
	if marqueeContent != req.Content {
		actionLog["content"] = createBancendActionLogDetail(marqueeContent, req.Content)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 運營管理/跑馬燈設定
// @Summary delete marquee
// @Description 刪除跑馬燈功能（開發者）：管理後台才可刪除【活動】類型跑馬燈
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.DeleteMarqueeRequest true "id"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/deletemarquee [post]
func (p *SystemManageApi) DeleteMarquee(c *ginweb.Context) {
	req := model.NewEmptyDeleteMarqueeRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode

	// 最高權限才能新增
	if len(myLevelCode) != 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 先傳送給game server，傳送成功再修改db，修改db失敗讓使用再重新操作即可
	// isEnabled := true
	isOpen := false

	notifyData := make(map[string]interface{}, 0)

	notifyData["method"] = "delete"
	notifyData["id"] = req.Id
	notifyData["is_open"] = isOpen

	// game server不存在表示已經過期，可以直接將本地資料刪除(關閉)
	code, err := notification.SendMarquee(notifyData)
	if code != 0 && code != notification.API_ERROR_MARQUEE_DO_NOT_EXIST {
		c.Logger().Info("delete marquee notification.SendMarquee fail, notifyData=%v, resp code=%d, err=%v", notifyData, code, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("marquee").
		Set("is_open", isOpen).
		Set("update_time", time.Now().UTC()).
		Where(sq.Eq{"id": req.Id}).
		Suffix("RETURNING \"content\"").
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	content := ""
	if rows.Next() {
		rows.Scan(&content)
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(content)
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 運營管理/後台公告
// @Summary Get announcement list
// @Description 取得目前後台公告列表：開發者、總代理、子代理，皆可查看跑馬燈資訊全內容
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.GetAnnouncementResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getannouncementlist [post]
func (p *SystemManageApi) GetAnnouncementList(c *ginweb.Context) {

	/*
		開發者、總代理、子代理，皆可查看公告資訊全內容，故不檢查權限
	*/

	// column: 6
	query := `SELECT "id", "type", "subject", "content", "create_time", "update_time"
			FROM announcement
			WHERE "is_open" = true;`

	rows, err := c.DB().Query(query)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	defer rows.Close()

	resp := make([]*model.GetAnnouncementResponse, 0)

	for rows.Next() {
		temp := model.NewEmptyGetAnnouncementResponse()
		// newline every scan 5 column
		if err := rows.Scan(&temp.Id, &temp.Type, &temp.Subject, &temp.Content, &temp.CreateTime,
			&temp.UpdateTime); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		resp = append(resp, temp)
	}

	c.Ok(resp)
}

// @Tags 運營管理/後台公告
// @Summary Get announcement
// @Description 指定取得某筆後台公告設定
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.GetAnnouncementRequest true "id"
// @Success 200 {object} response.Response{data=model.GetAnnouncementResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getannouncement [post]
func (p *SystemManageApi) GetAnnouncement(c *ginweb.Context) {
	/*
		開發者、總代理、子代理，皆可查看公告資訊全內容，故不檢查權限
	*/

	req := model.NewEmptyGetAnnouncementRequest()
	_ = c.ShouldBindJSON(&req)

	// 檢查參數是否合法
	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	id := req.Id

	temp := model.NewEmptyGetAnnouncementResponse()
	// column: 6
	query := `SELECT "id", "type", "subject", "content", "create_time", "update_time"
			FROM announcement
			WHERE "is_open" = true AND "id" = $1;`

	err := c.DB().QueryRow(query, id).Scan(&temp.Id, &temp.Type, &temp.Subject, &temp.Content, &temp.CreateTime,
		&temp.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ANNOUNCEMENT_NOT_EXIST)
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		}
		return
	}
	c.Ok(temp)
}

// @Tags 運營管理/後台公告
// @Summary create new announcement
// @Description 添加後台公告功能（開發者）
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.CreateAnnouncementRequest true "id, 創建內容相關參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/createannouncement [post]
func (p *SystemManageApi) CreateAnnouncement(c *ginweb.Context) {
	req := model.NewEmptyCreateAnnouncementRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 最高權限才能新增
	if claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	isOpen := true

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("announcement").
		Columns("\"type\"", "\"subject\"", "\"content\"", "is_open").
		Values(req.Type, req.Subject, req.Content, isOpen).
		Suffix("RETURNING id").
		ToSql()

	id := ""
	err := c.DB().QueryRow(query, args...).Scan(&id)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Subject)
	c.Set("action_log", actionLog)

	// 送公告通知
	connInfo, code := getChatConnInfo(c.DB(), c.Request.Context())
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	channelId, ok := connInfo["channel"].(string)
	if !ok {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHANNEL_ID)
		return
	}

	message := utils.ToJSON(map[string]interface{}{
		"id":     id,
		"method": "create",
	})

	success := notification.SendNotifyToFrontend(channelId, message, notification.ChatNotification_broadcast, connInfo, notification.API_CHAT_MESSAGE_SUBJECT_ANNOUNCEMENT)
	if !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHAT_SEND_MESSAGE_FAILED)
		return
	}

	c.Ok("")
}

// @Tags 運營管理/後台公告
// @Summary update new announcement
// @Description 編輯後台公告功能（開發者）
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.UpdateAnnouncementRequest true "id, 更新內容相關參數"
// @Success 200 {object} response.Response{data=model.GetMarqueeResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/updateannouncement [post]
func (p *SystemManageApi) UpdateAnnouncement(c *ginweb.Context) {
	req := model.NewEmptyUpdateAnnouncementRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 最高權限才能新增
	if claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	announcementType := definition.MARQUEE_TYPE_NONE
	announcementSubject := ""
	announcementContent := ""

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("\"type\"", "\"subject\"", "\"content\"").
		From("announcement").
		Where(sq.Eq{"id": req.Id}).
		ToSql()

	err := c.DB().QueryRow(query, args...).Scan(&announcementType, &announcementSubject, &announcementContent)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("announcement").
		Set("\"type\"", req.Type).
		Set("\"subject\"", req.Subject).
		Set("\"content\"", req.Content).
		Set("update_time", time.Now().UTC()).
		Where(sq.Eq{"id": req.Id}).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Subject)
	if announcementType != req.Type {
		actionLog["announcement_type"] = createBancendActionLogDetail(announcementType, req.Type)
	}
	if announcementSubject != req.Subject {
		actionLog["subject"] = createBancendActionLogDetail(announcementSubject, req.Subject)
	}
	if announcementContent != req.Content {
		actionLog["content"] = createBancendActionLogDetail(announcementContent, req.Content)
	}
	c.Set("action_log", actionLog)

	// 送公告通知
	connInfo, code := getChatConnInfo(c.DB(), c.Request.Context())
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	channelId, ok := connInfo["channel"].(string)
	if !ok {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHANNEL_ID)
		return
	}

	message := utils.ToJSON(map[string]interface{}{
		"id":     req.Id,
		"method": "update",
	})

	success := notification.SendNotifyToFrontend(channelId, message, notification.ChatNotification_broadcast, connInfo, notification.API_CHAT_MESSAGE_SUBJECT_ANNOUNCEMENT)
	if !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHAT_SEND_MESSAGE_FAILED)
		return
	}

	c.Ok("")
}

// @Tags 運營管理/後台公告
// @Summary delete announcement
// @Description 刪除後台公告功能（開發者）
// @accept application/json
// @Produce application/json
// @Security BearerAuth
// @Param data body model.DeleteAnnouncementRequest true "id"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/deleteannouncement [post]
func (p *SystemManageApi) DeleteAnnouncement(c *ginweb.Context) {
	req := model.NewEmptyDeleteAnnouncementRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 最高權限才能新增
	if claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	isOpen := false

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("announcement").
		Set("is_open", isOpen).
		Set("update_time", time.Now().UTC()).
		Where(sq.Eq{"id": req.Id}).
		Suffix("RETURNING \"subject\"").
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	subject := ""
	if rows.Next() {
		rows.Scan(&subject)
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(subject)
	c.Set("action_log", actionLog)

	c.Ok("")
}

type MaintainPageSetting struct {
	StartTime string `json:"start_time"` // 開始時間 YYYY-MM-DD HH:mm
	EndTime   string `json:"end_time"`   // 結束時間 YYYY-MM-DD HH:mm
	Timezone  string `json:"timezone"`   // 時區
}

type SendNotifyToGetMaintainPageSettingResp struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

// @Tags 運營管理/維護頁設定
// @Summary 取得維護頁設定(只有開發商(營用商)可以使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=MaintainPageSetting,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/getmaintainpagesetting [post]
func (p *SystemManageApi) GetMaintainPageSetting(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	connInfo, code := getMatainConnInfo(c.DB(), c.Request.Context())
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	isSuccess, result := notification.SendNotifyToGetMaintainPageSetting(connInfo, notification.MaintainNotification_action_read)
	if !isSuccess {
		c.Logger().Info("notification.SendNotifyToGetMaintainPageSetting failed, result =%s", result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	var resp SendNotifyToGetMaintainPageSettingResp
	utils.ToStruct([]byte(result), &resp)

	if resp.Code != 0 && resp.Message != "No Data" {
		c.Logger().Info("notification.SendNotifyToGetMaintainPageSetting failed, result =%s", result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	ret := &MaintainPageSetting{}
	if resp.Code == 0 && len(resp.Data) > 0 {
		ret.StartTime = resp.Data[0]["starttime"].(string)
		ret.EndTime = resp.Data[0]["endtime"].(string)
		ret.Timezone = resp.Data[0]["timezone"].(string)
	}

	c.Ok(ret)
}

// @Tags 運營管理/維護頁設定
// @Summary 設定維護頁設定(只有開發商(營用商)可以使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body MaintainPageSetting true "維護頁設資料"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/manage/setmaintainpagesetting [post]
func (p *SystemManageApi) SetMaintainPageSetting(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	var req MaintainPageSetting
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.Timezone == "" || utils.WordLength(req.Timezone) > 20 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	checkTimeFormat := func(timeStr string) bool {
		sp := strings.Split(timeStr, " ")
		if len(sp) != 2 {
			return false
		}

		dateSp := strings.Split(sp[0], "-")
		if len(dateSp) != 3 {
			return false
		}

		timeSp := strings.Split(sp[1], ":")
		if len(timeSp) != 2 {
			return false
		}

		year := utils.ToInt(dateSp[0])
		month := time.Month(utils.ToInt(dateSp[1]))
		day := utils.ToInt(dateSp[2])
		hour := utils.ToInt(timeSp[0])
		min := utils.ToInt(timeSp[1])

		time := time.Date(year, month, day, hour, min, 0, 0, time.Local)
		result := fmt.Sprintf("%d-%02d-%02d %02d:%02d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute())
		fmt.Println(result, timeStr)

		return result == timeStr
	}

	if !checkTimeFormat(req.StartTime) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !checkTimeFormat(req.EndTime) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	connInfo, code := getMatainConnInfo(c.DB(), c.Request.Context())
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	isSuccess, result := notification.SendNotifyToGetMaintainPageSetting(connInfo, notification.MaintainNotification_action_read)
	if !isSuccess {
		c.Logger().Info("notification.SendNotifyToGetMaintainPageSetting failed, result =%s", result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	var resp SendNotifyToGetMaintainPageSettingResp
	utils.ToStruct([]byte(result), &resp)

	if resp.Code != 0 && resp.Message != "No Data" {
		c.Logger().Info("notification.SendNotifyToGetMaintainPageSetting failed, result =%s", result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	backup := &MaintainPageSetting{}
	if resp.Code == 0 && len(resp.Data) > 0 {
		backup.StartTime = resp.Data[0]["starttime"].(string)
		backup.EndTime = resp.Data[0]["endtime"].(string)
		backup.Timezone = resp.Data[0]["timezone"].(string)
	}

	isSuccess, result = notification.SendNotifyToModifyMaintainPageSetting(
		connInfo,
		notification.MaintainNotification_action_write,
		req.Timezone,
		req.StartTime,
		req.EndTime)
	if !isSuccess {
		c.Logger().Info("notification.SendNotifyToModifyMaintainPageSetting failed, req = %v, result = %s", req, result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	utils.ToStruct([]byte(result), &resp)
	if resp.Code != 0 {
		c.Logger().Info("notification.SendNotifyToModifyMaintainPageSetting failed, req = %v, result = %s", req, result)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	if backup.StartTime != req.StartTime {
		actionLog["start_time"] = createBancendActionLogDetail(backup.StartTime, req.StartTime)
	}
	if backup.EndTime != req.EndTime {
		actionLog["end_time"] = createBancendActionLogDetail(backup.EndTime, req.EndTime)
	}
	if backup.Timezone != req.Timezone {
		actionLog["timezone"] = createBancendActionLogDetail(backup.Timezone, req.Timezone)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}
