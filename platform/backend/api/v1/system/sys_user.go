package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/internal/statistical"
	"backend/pkg/encrypt"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemUserApi struct {
	BasePath string
}

func NewSystemUserApi(basePath string) api_cluster.IApiEach {
	return &SystemUserApi{
		BasePath: basePath,
	}
}

func (p *SystemUserApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemUserApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.GET("/ping", ginHandler.Handle(p.Ping))
	g.POST("/getalivetokenlist", ginHandler.Handle(p.GetAliveTokenList))
	g.POST("/blacktoken", ginHandler.Handle(p.BlackToken))
	g.POST("/createadminuser",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateAdminUser))
	g.POST("/getadminusers", ginHandler.Handle(p.GetAdminUsers))
	g.POST("/getadminuserinfo", ginHandler.Handle(p.GetAdminUserInfo))
	g.POST("/updateadminuserinfo",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.UpdateAdminUserInfo))
	g.POST("/getgameusers", ginHandler.Handle(p.GetGameUsers))
	g.POST("/getgameuserinfo", ginHandler.Handle(p.GetGameUserInfo))
	g.POST("/updategameuserinfo",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.UpdateGameUserInfo))
	g.POST("/getgameuserwalletlist", ginHandler.Handle(p.GetGameUserWalletList))
	g.POST("/setgameuserwallet", ginHandler.Handle(p.SetGameUserWallet))
	g.POST("/setpersonalinfo",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetPersonalInfo))
	g.POST("/setpersonalpassword",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetPersonalPassword))
	g.POST("/resetpassword",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.ResetPassword))
	g.POST("/getgameuserbalance", ginHandler.Handle(p.GetGameUserBalance))
	g.POST("/getgameuserplaycountdata", ginHandler.Handle(p.GetGameUserPlayCountData))
}

// @Tags User
// @Summary ping
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/ping [get]
func (p *SystemUserApi) Ping(c *ginweb.Context) {
	c.Ok("pong")
}

// @Tags 代理帳號管理/後台帳號
// @Summary 取得目前已產生的有效 token list(已登入後台帳號列表)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.AliveTokenListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getalivetokenlist [post]
func (p *SystemUserApi) GetAliveTokenList(c *ginweb.Context) {

	tokenList := c.GetJwt().TokenCacheList.GetAll()

	var temp model.AliveTokenListResponse
	temp.Response(&tokenList)
	c.Ok(temp)
}

// @Tags 代理帳號管理/後台帳號
// @Summary 將用戶 token 列入黑名單(主動使已登入後台帳號強制登出)
// @Produce  application/json
// @Security BearerAuth
// @Param token formData string true "token"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/blacktoken [post]
func (p *SystemUserApi) BlackToken(c *ginweb.Context) {
	token := c.PostForm("token")

	if token == "" {
		c.Fail("parse request error")
		return
	}
	claims, err := c.GetJwt().ParseToken(token)
	if err != nil || claims == nil {
		c.Fail("parse parameter error")
		return
	}
	c.GetJwt().AddBlackToken(token, claims)
	c.Ok("")
}

// @Tags 代理帳號管理/後台帳號
// @Summary 依照查詢者角色權限列出自身權限下的後台帳號列表
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.GetAdminUsersData,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getadminusers [post]
func (p *SystemUserApi) GetAdminUsers(c *ginweb.Context) {

	claims := c.GetUserClaims()
	if claims != nil {

		myAgentId := int(claims.BaseClaims.ID)
		myIsAdded := claims.BaseClaims.IsAdded

		query := `SELECT au.username, au.nickname, au.is_enabled, au.create_time, au.login_time, au.is_added, ap."name"
						  FROM admin_user au, agent_permission ap
						  WHERE au.agent_id = $1 AND au."role" = ap."id"`
		args := make([]interface{}, 0)
		args = append(args, myAgentId)

		// 如果是分身帳號
		if myIsAdded {
			query = query + ` AND "is_added" = $2;`
			args = append(args, myIsAdded)
		}

		at := global.AgentCache.Get(myAgentId)
		topUsername := ""
		if at != nil {
			topUsername = at.AdminUsername
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST)
			return
		}

		rows, err := c.DB().Query(query, args...)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		defer rows.Close()

		data := make([]*model.GetAdminUsersData, 0)
		for rows.Next() {
			var temp model.GetAdminUsersData
			isAdded := false
			if err := rows.Scan(&temp.Username, &temp.Nickname, &temp.IsEnabled, &temp.CreateTime, &temp.LoginTime,
				&isAdded, &temp.RoleName); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}

			if isAdded {
				// 分身
				temp.TopUsername = topUsername
			} else {
				// 自己
				temp.TopUsername = "-"
			}

			data = append(data, &temp)
		}

		c.Ok(data)
		return

	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}
}

// @Tags 代理帳號管理/後台帳號
// @Summary 創建後台帳號(只能創建自己的後台帳號)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.AdminUserCreateRequest true "代理編號, 用戶名, 密碼"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/createadminuser [post]
func (p *SystemUserApi) CreateAdminUser(c *ginweb.Context) {
	/*
		Tip: 創建新帳號時只輸入必要欄位，ex 帳號、密碼......
		其它個人化設定，一律讀取預設值
	*/

	var l model.AdminUserCreateRequest
	_ = c.ShouldBindJSON(&l)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	myAgentId := -1 // default  最高權限
	myAccountType := 0
	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	myAgentId = int(claims.BaseClaims.ID)
	myAccountType = claims.AccountType

	// 取得自已登入的後台帳號所屬代理的 SecretKey 加密後台設定密碼
	myAgentObj := global.AgentCache.Get(myAgentId)
	if myAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
		return
	}
	secretKey := myAgentObj.SecretKey

	// 將密碼做加密
	decryptPwd, err := encrypt.EncryptSaltToken(l.Password, secretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	checkDataExist := 0
	// 檢查權限群組是否存在
	query := `SELECT count(*)
	FROM agent_permission 
	WHERE "id" = $1;`

	err = c.DB().QueryRow(query, l.Role).Scan(&checkDataExist)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	} else if checkDataExist == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION_ROLE_NOT_EXIST)
		return
	}

	query = `SELECT count(*)
	FROM admin_user
	WHERE "username" = $1;`
	err = c.DB().QueryRow(query, l.Username).Scan(&checkDataExist)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	} else if checkDataExist > 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_EXIST)
		return
	}

	adminUser, err := global.UdfCreateNewAdminUser(c.DB(), myAgentId, l.Username, decryptPwd, l.Nickname, myAccountType,
		definition.ACCOUNT_READONLY_OFF, l.Role, l.Info)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	global.AdminUserCache.Add(adminUser)

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(adminUser.Username)
	c.Set("action_log", actionLog)

	c.Ok("Create successed.")
}

// @Tags 代理帳號管理/後台帳號
// @Summary 指定查詢某後台帳號狀態
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAdminUserInfoRequest true "用戶id"
// @Success 200 {object} response.Response{data=model.GetAdminUserInfoResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getadminuserinfo [post]
func (p *SystemUserApi) GetAdminUserInfo(c *ginweb.Context) {

	var l model.GetAdminUserInfoRequest
	_ = c.ShouldBindJSON(&l)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	myAgentId := int(claims.BaseClaims.ID)

	acoountType := 0
	resp := model.GetAdminUserInfoResponse{}

	query := `SELECT "role", "info", "is_enabled", "account_type"
				  FROM admin_user 
				  WHERE "username" = $1 AND "agent_id" = $2;`

	err := c.DB().QueryRow(query, l.Username, myAgentId).Scan(&resp.Role, &resp.Info, &resp.IsEnabled, &acoountType)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	c.Ok(resp)
}

// @Tags 代理帳號管理/後台帳號
// @Summary 指定設定某後台帳號狀態
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.UpdateAdminUserInfoRequest true "用戶id"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/updateadminuserinfo [post]
func (p *SystemUserApi) UpdateAdminUserInfo(c *ginweb.Context) {

	var l model.UpdateAdminUserInfoRequest
	_ = c.ShouldBindJSON(&l)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	myAgentId := int(claims.BaseClaims.ID)
	myAccountType := int(claims.BaseClaims.AccountType)

	// TODO: 檢查設定的角色資料是否存在
	checkDataExist := 0
	// 檢查權限群組是否存在
	query := `SELECT count(*)
		FROM agent_permission 
		WHERE "id" = $1 AND "agent_id" = $2 AND "account_type" = $3;`

	err := c.DB().QueryRow(query, l.PermissionId, myAgentId, myAccountType).Scan(&checkDataExist)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	} else if checkDataExist == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION_ROLE_NOT_EXIST)
		return
	}

	updateTime, err := global.UdfUpdateAdminUser(c.DB(), myAgentId, l.Username, l.PermissionId, l.Info, l.IsEnabled)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	adminUser := global.AdminUserCache.Get(myAgentId, l.Username)
	// admin user變更權限群組，將舊token加入黑名單重新登入取的新權限
	// 如果帳號被封停也要把舊TOKEN加入黑名單
	if adminUser.PermissionId != l.PermissionId || l.IsEnabled == 0 {
		c.Jwt.AddBlackTokenByUsername(adminUser.Username)
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(adminUser.Username)
	if adminUser.PermissionId != l.PermissionId {
		beforeAgentPermission := global.AgentPermissionCache.Get(adminUser.PermissionId)
		afterAgentPermission := global.AgentPermissionCache.Get(l.PermissionId)
		actionLog["agent_permission"] = createBancendActionLogDetail(beforeAgentPermission.Name, afterAgentPermission.Name)
	}
	if adminUser.Info != l.Info {
		actionLog["info"] = createBancendActionLogDetail(adminUser.Info, l.Info)
	}
	if adminUser.IsEnabled != l.IsEnabled {
		actionLog["number_status"] = createBancendActionLogDetail(adminUser.IsEnabled, l.IsEnabled)
	}
	c.Set("action_log", actionLog)

	// 更新cache資料
	adminUser.PermissionId = l.PermissionId
	adminUser.Info = l.Info
	adminUser.IsEnabled = l.IsEnabled
	adminUser.UpdateTime = updateTime

	c.Ok("")
}

// @Tags 運營管理/玩家帳號相關
// @Summary 依照查詢者角色權限列出遊戲會員帳號清單
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.GetGameUsersData,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getgameusers [post]
func (p *SystemUserApi) GetGameUsers(c *ginweb.Context) {

	claims := c.GetUserClaims()
	if claims != nil {

		myLevelCode := claims.BaseClaims.LevelCode

		query := `SELECT gu."id", gu.original_username, gu.is_enabled, gu.agent_id, aa."name",
				gu.create_time, gu.sum_coin_in, gu.sum_coin_out, gu.is_risk, gu.kill_dive_state,
				gu.kill_dive_value, gu.custom_status, gu.last_login_time, gu.risk_control_status, aa.wallet_type
				FROM game_users gu, agent aa
				WHERE gu.agent_id = aa."id"  AND gu.level_code LIKE $1 ORDER BY gu."id" ASC;`

		rows, err := c.DB().Query(query, myLevelCode+"%")
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		defer rows.Close()

		inGameUserData, err := c.Redis().LoadHAllValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER)
		if err != nil {
			c.Result(definition.ERROR_CODE_ERROR_REDIS, "", nil)
			return
		}

		// inGameUserData

		resp := make([]*model.GetGameUsersData, 0)
		for rows.Next() {
			var temp model.GetGameUsersData
			if err := rows.Scan(&temp.Id, &temp.Username, &temp.IsEnabled, &temp.AgentId, &temp.AgentName,
				&temp.CreateTime, &temp.CoinIn, &temp.CoinOut, &temp.HighRisk, &temp.KillDiveState,
				&temp.KillDiveValue, &temp.TagList, &temp.LastLoginTime, &temp.RiskControlTagList, &temp.WalletType); err != nil {
				c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
				return
			}

			if _, ok := inGameUserData[strconv.Itoa(temp.Id)]; ok {
				temp.IsOnline = true
			}

			agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(temp.AgentId)
			if !ok {
				c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
				return
			}

			temp.CustomTagInfo = agentGameRatioCache.CustomTagInfo

			resp = append(resp, &temp)
		}
		c.Ok(resp)
		return

	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}
}

// @Tags 運營管理/玩家帳號相關
// @Summary 指定查詢某遊戲會員帳號信息
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetGameUserInfoRequest true "用戶id"
// @Success 200 {object} response.Response{data=[]model.GetGameUserInfoResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getgameuserinfo [post]
func (p *SystemUserApi) GetGameUserInfo(c *ginweb.Context) {

	var l model.GetGameUserInfoRequest
	_ = c.ShouldBindJSON(&l)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {

		myLevelCode := claims.BaseClaims.LevelCode

		gameUserInfo := ""
		gameUserIsEnabled := false
		// 檢查被查詢的遊戲帳號 id 與查詢者的 level code 是否合法
		query := `SELECT "info", "is_enabled"
				  FROM game_users 
				  WHERE "id" = $1 AND "level_code" LIKE $2;`

		err := c.DB().QueryRow(query, l.Id, myLevelCode+"%").Scan(&gameUserInfo, &gameUserIsEnabled)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		query = `SELECT COALESCE(SUM(de),0), COALESCE(SUM(ya),0), COALESCE(SUM(vaild_ya),0), COALESCE(SUM(tax),0), COALESCE(SUM(bonus),0) FROM game_users_stat WHERE game_users_id = $1;`

		var sumDe, sumYa, sumValidYa, sumTax, sumBonus float64
		err = c.DB().QueryRow(query, l.Id).Scan(&sumDe, &sumYa, &sumValidYa, &sumTax, &sumBonus)
		if err != nil {
			if err != sql.ErrNoRows {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}
		}

		// 時區計算
		// 找出當月時間
		nowTime := time.Now().UTC()
		transEndMonth := nowTime.Month()

		timZoneMin := l.TimeZone

		transHour := timZoneMin / 60
		transMin := timZoneMin % 60

		beforeTime := time.Date(nowTime.Year(), transEndMonth, 1, transHour, transMin, 0, 0, nowTime.UTC().Location())
		beforeTimeStr := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_HOUR, beforeTime)

		query = `SELECT COALESCE(SUM(de),0) as de, COALESCE(SUM(ya),0) as ya, COALESCE(SUM(vaild_ya),0) as v_ya, COALESCE(SUM(tax),0) as tax, COALESCE(SUM(bonus),0) as bonus
				FROM game_users_stat_hour 
				WHERE game_users_id = $1 AND log_time >= $2
				;`

		var sumDeP, sumYaP, sumValidYaP, sumTaxP, sumBonusP float64
		err = c.DB().QueryRow(query, l.Id, beforeTimeStr).Scan(&sumDeP, &sumYaP, &sumValidYaP, &sumTaxP, &sumBonusP)
		if err != nil {
			if err != sql.ErrNoRows {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}
		}

		// 总有效投注、总盈利、当月有效投注、当月盈利
		// 玩家視角
		resp := model.NewEmptyGetGameUserInfoResponse()
		resp.ValidBetSum = sumValidYa
		resp.ProfitSum = utils.DecimalSub(sumDe, sumYa)
		resp.TaxSum = sumTax
		resp.BonusSum = sumBonus
		resp.ValidBet = sumValidYaP                    // 當月有效投注
		resp.Profit = utils.DecimalSub(sumDeP, sumYaP) // 當月營利
		resp.Tax = sumTaxP                             // 當月抽水
		resp.Bonus = sumBonusP                         // 當月紅利
		resp.Info = gameUserInfo
		resp.IsEnabled = gameUserIsEnabled

		c.Ok(resp)
		return

	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}
}

// @Tags 運營管理/玩家帳號相關
// @Summary 指定修改某遊戲會員帳號信息
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.UpdateGameUserInfoRequest true "用戶id, "
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/updategameuserinfo [post]
func (p *SystemUserApi) UpdateGameUserInfo(c *ginweb.Context) {

	var l model.UpdateGameUserInfoRequest
	_ = c.ShouldBindJSON(&l)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode

	dbUsername := ""
	dbInfo := ""
	dbIsEnabled := false
	dbAgentName := ""

	query := `SELECT "gu"."original_username", "gu"."info", "gu"."is_enabled", "a"."name"
	            FROM "public"."game_users" AS "gu"
				INNER JOIN "public"."agent" AS "a" ON "gu"."agent_id" = "a"."id"
				WHERE "gu"."id" = $1 AND "gu"."level_code" LIKE $2;`
	err := c.DB().QueryRow(query, l.Id, myLevelCode+"%").Scan(&dbUsername, &dbInfo, &dbIsEnabled, &dbAgentName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST)
			return
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
	}

	query = `UPDATE "public"."game_users"
				SET "info" = $1, "is_enabled" = $2
				WHERE "id" = $3 AND level_code LIKE $4;`
	_, err = c.DB().Exec(query, l.Info, l.IsEnabled, l.Id, myLevelCode+"%")
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// TODO: if disable game user account, kick user if in game.
	if dbIsEnabled != l.IsEnabled {
		// 遊戲那邊封鎖是傳 true ,所以 IsEnabled 要加!
		if msg, code, err := notification.SendSetBlockPlayer([]int{l.Id}, !l.IsEnabled); err != nil {
			c.Logger().Printf("UpdateGameUserInfo() SendSetPlayerBlock has error, msg: %s, code: %d, err: %v", msg, code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}

		// 如果是帳號封停
		if !l.IsEnabled {
			// delete user login token
			reqUserId := utils.IntToString(l.Id)
			serverInfoCode, err := c.Redis().LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
			if err == nil && serverInfoCode != "" {
				_ = c.Redis().DeleteHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
				_ = c.Redis().DeleteHValue(
					global.REDIS_IDX_LOGIN_INFO,
					global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, serverInfoCode),
					reqUserId)
			}
		}
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["username"] = dbUsername
	actionLog["agent_name"] = dbAgentName
	if dbInfo != l.Info {
		actionLog["info"] = createBancendActionLogDetail(dbInfo, l.Info)
	}
	if dbIsEnabled != l.IsEnabled {
		actionLog["bool_status"] = createBancendActionLogDetail(dbIsEnabled, l.IsEnabled)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 運營管理/後台玩家上下分
// @Summary 取得玩家錢包餘額列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUserWalletListRequest true "取得玩家錢包餘額列表參數"
// @Success 200 {object} response.Response{data=model.GetGameUserWalletListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getgameuserwalletlist [post]
func (p *SystemUserApi) GetGameUserWalletList(c *ginweb.Context) {
	req := &model.GetGameUserWalletListRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()

	// 必要條件檢查
	checkExist_username := (req.Username != "")

	sqAnd := sq.And{}
	if checkExist_username {
		sqAnd = append(sqAnd, sq.Eq{"gu.original_username": req.Username})
	}

	sqAnd = append(sqAnd, sq.Eq{"gu.agent_id": userClaims.BaseClaims.ID})

	resp := &model.GetGameUserWalletListResponse{}
	resp.Draw = req.Draw
	resp.Data = make([]interface{}, 0)

	// 限制只有第一次會進行加總
	if req.Draw == 0 {
		recordsTotal, code := getGameUserWalletListSumInfo(c.DB(), &sqAnd)
		if code != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(code)
			return
		}

		resp.RecordsTotal = recordsTotal
	}

	data, code := getGameUserWalletList(c.DB(), req, &sqAnd)
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	if len(data) > 0 {
		users := make([]int, 0)
		for i := 0; i < len(data); i++ {
			users = append(users, data[i].UserId)
			resp.Data = append(resp.Data, data[i])
		}

		gameUserGolds, code, _ := notification.GetGameServerGolds(users)
		if code != definition.ERROR_CODE_SUCCESS {
			c.Logger().Info("GetGameUserWalletList notification.GetGameServerGolds fail, users=%v, resp code=%d", users, code)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return

		}
		for _, iData := range resp.Data {
			data := iData.(*model.GetGameUserWalletResponse)
			data.Gold = gameUserGolds[data.UserId]["gold"].(float64)
			data.LockGold = gameUserGolds[data.UserId]["lockgold"].(float64)
		}
	}

	agentBalance, code := getAgentWalletAmount(c.DB(), int(userClaims.BaseClaims.ID))
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}
	resp.AgentBalance = agentBalance

	c.Ok(resp)
}

func getAgentWalletAmount(db *sql.DB, agentId int) (amount float64, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("amount").
		From("agent_wallet").
		Where(sq.Eq{"agent_id": agentId}).
		ToSql()
	if err != nil {
		return
	}

	err = db.QueryRow(query, args...).Scan(&amount)
	if err != nil {
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return

}

func getGameUserWalletListSumInfo(db *sql.DB, sqAnd *sq.And) (recordsTotal int, code int) {
	code = definition.ERROR_CODE_ERROR_DATABASE

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(gu.id)").
		From("game_users AS gu").
		Where(sqAnd).
		ToSql()
	if err != nil {
		return
	}

	err = db.QueryRow(query, args...).Scan(&recordsTotal)
	if err != nil {
		return
	}

	code = definition.ERROR_CODE_SUCCESS

	return
}

func getGameUserWalletList(db *sql.DB, req *model.GetGameUserWalletListRequest, sqAnd *sq.And) (data []*model.GetGameUserWalletResponse, code int) {
	code = definition.ERROR_CODE_ERROR_DATABASE

	sqBuilder := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("gu.id", "gu.original_username", "a.name", "gu.is_enabled", "gu.create_time").
		From("game_users AS gu").
		InnerJoin("agent AS a ON gu.agent_id = a.id").
		Where(sqAnd)
	sqBuilder = getGameUserWalletListEvaluateOrderBy(sqBuilder, req.SortColumn, req.SortDirection)

	query, args, err := sqBuilder.
		Limit(uint64(req.Length)).Offset(uint64(req.Start)).
		ToSql()
	if err != nil {
		return
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}

	defer rows.Close()

	data = make([]*model.GetGameUserWalletResponse, 0)
	for rows.Next() {
		var temp model.GetGameUserWalletResponse
		if err := rows.Scan(&temp.UserId, &temp.Username, &temp.AgentName, &temp.IsEnabled,
			&temp.CreateTime); err != nil {
			return
		}
		data = append(data, &temp)
	}

	code = definition.ERROR_CODE_SUCCESS

	return
}

func getGameUserWalletListEvaluateOrderBy(sqBuilder sq.SelectBuilder, sortColumn string, sortDirection int) sq.SelectBuilder {
	if sortColumn == "username" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("gu.original_username asc")
		} else {
			return sqBuilder.OrderBy("gu.original_username desc")
		}
	} else if sortColumn == "agentname" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("a.name asc")
		} else {
			return sqBuilder.OrderBy("a.name desc")
		}
	} else if sortColumn == "isenabled" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("gu.is_enabled asc")
		} else {
			return sqBuilder.OrderBy("gu.is_enabled desc")
		}
	} else if sortColumn == "createtime" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("gu.create_time asc")
		} else {
			return sqBuilder.OrderBy("gu.create_time desc")
		}
	} else {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("gu.id asc")
		} else {
			return sqBuilder.OrderBy("gu.id desc")
		}
	}
}

// @Tags 運營管理/後台玩家上下分
// @Summary 設置玩家錢包餘額
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetGameUserWalletRequest true "設置玩家錢包餘額參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/setgameuserwallet [post]
func (p *SystemUserApi) SetGameUserWallet(c *ginweb.Context) {
	req := &model.SetGameUserWalletRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	// 開發商不可使用此功能
	if userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 取得基本資料
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("gu.original_username", "a.cooperation", "a.level_code", "a.aes_key", "a.md5_key").
		From("game_users AS gu").
		InnerJoin("agent AS a ON gu.agent_id = a.id").
		Where(sq.And{
			sq.Eq{"gu.id": req.UserId},
			sq.Eq{"gu.agent_id": userClaims.BaseClaims.ID},
		}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var username string
	var agentCooperation int
	var agentLevelCode string
	var agentAesKey string
	var agentMd5Key string
	err = c.DB().QueryRow(query, args...).Scan(&username, &agentCooperation, &agentLevelCode, &agentAesKey, &agentMd5Key)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 取得api server資料
	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("addresses", "is_enabled").
		From("server_info").
		Where(sq.Eq{"code": "api"}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var address []byte
	var isEnabled bool
	err = c.DB().QueryRowContext(c.Request.Context(), query, args...).Scan(&address, &isEnabled)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	} else if !isEnabled {
		c.OkWithCode(definition.ERROR_CODE_ERROR_FEATURE_DISABLED)
		return
	}
	connInfo := utils.ToMap(address)

	// 參數預設上分
	s := 2
	account := username
	money := req.ChangeScore
	kind := definition.WALLET_LEDGER_KIND_BACKEND_UP

	// 參數調整成下分
	if req.ChangeScore < 0 {
		s = 3
		money = -req.ChangeScore
		kind = definition.WALLET_LEDGER_KIND_BACKEND_DOWN
	}

	timeNow := time.Now()

	orderId := fmt.Sprintf("%d%d%02d%02d%02d%02d%02d%06d%s", userClaims.BaseClaims.ID, timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.UnixMicro()%int64(time.Microsecond), username)
	info := req.Info
	creator := c.GetUsername()

	param := fmt.Sprintf("s=%d&account=%s&money=%v&orderid=%s&beinfo=%s&becreator=%s&bekind=%d", s, account, money, orderId, info, creator, kind)

	resp, err := sendApiServer(connInfo, int64(userClaims.BaseClaims.ID), param, agentAesKey, agentMd5Key)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_API_SERVER_REQUEST_FAILED)
		return
	}

	respData := resp.Data.(map[string]interface{})
	respDataD := respData["d"].(map[string]interface{})

	c.OkWithCode(int(respDataD["code"].(float64)))
}

// @Tags 個人資訊/修改個人資訊
// @Summary 修改個人資訊
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetPersonalInfoRequest true "修改個人資訊參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/setpersonalinfo [post]
func (p *SystemUserApi) SetPersonalInfo(c *ginweb.Context) {
	req := &model.SetPersonalInfoRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myName := c.GetUsername()
	myNickname := c.GetNickname()
	myAgentId := claims.BaseClaims.ID

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("admin_user").
		Set("nickname", req.Nickname).
		Where(sq.And{
			sq.Eq{"agent_id": myAgentId},
			sq.Eq{"username": myName},
		}).
		ToSql()

	_, err := c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(myName)
	if myNickname != req.Nickname {
		actionLog["nickname"] = createBancendActionLogDetail(myNickname, req.Nickname)
	}
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

// @Tags 個人資訊/修改個人密碼
// @Summary 修改個人密碼
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetPersonalPasswordRequest true "修改個人密碼參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/setpersonalpassword [post]
func (p *SystemUserApi) SetPersonalPassword(c *ginweb.Context) {
	req := &model.SetPersonalPasswordRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myName := c.GetUsername()
	myAgentId := claims.BaseClaims.ID
	myAgentObj := global.AgentCache.Get(int(myAgentId))
	if myAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("password").
		From("admin_user").
		Where(sq.And{
			sq.Eq{"agent_id": myAgentId},
			sq.Eq{"username": myName},
		}).
		ToSql()

	myOldPassword := ""
	err := c.DB().QueryRow(query, args...).Scan(&myOldPassword)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	secretKey := myAgentObj.SecretKey

	// DB密碼為加密後的密碼，所以需要解密來確認是否相同
	decryptOldPassword, err := encrypt.DecryptSaltToken(myOldPassword, secretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	if req.OldPassword != decryptOldPassword {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PASSWORD)
		return
	}

	encryptNewPassword, err := encrypt.EncryptSaltToken(req.NewPassword, secretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("admin_user").
		Set("password", encryptNewPassword).
		Where(sq.And{
			sq.Eq{"agent_id": myAgentId},
			sq.Eq{"username": myName},
		}).
		ToSql()

	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["username"] = myName
	c.Set("action_log", actionLog)

	// 密碼修改成功後要將用戶登出，強迫使用新密碼重新登入
	c.Jwt.AddBlackTokenByUsername(myName)

	c.Ok(nil)
}

// @Tags 帳號/重置密碼
// @Summary 重置密碼
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ResetPasswordRequest true "重置密碼參數"
// @Success 200 {object} response.Response{data=model.ResetPasswordResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/resetpassword [post]
func (p *SystemUserApi) ResetPassword(c *ginweb.Context) {
	req := &model.ResetPasswordRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("au.agent_id", "a.top_agent_id", "a.secret_key").
		From("admin_user AS au").
		InnerJoin(`agent AS a ON a."id" = au.agent_id`).
		Where(sq.Eq{"au.username": req.Username}).
		ToSql()

	var userAgentId, userAgentTopAgentId int
	var userAgentSecretKey string
	err := c.DB().QueryRow(query, args...).Scan(&userAgentId, &userAgentTopAgentId, &userAgentSecretKey)
	if err != nil {
		errCode := definition.ERROR_CODE_ERROR_DATABASE
		if err == sql.ErrNoRows {
			errCode = definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST
		}
		c.OkWithCode(errCode)
		return
	}

	myAgentId := int(claims.BaseClaims.ID)
	if userAgentId != myAgentId && userAgentTopAgentId != myAgentId {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	timeNowStr := time.Now().String()
	newPassword := md5.Hash10bit(req.Username + timeNowStr)
	// DB密碼為加密後的密碼
	encryptNewPassword, err := encrypt.EncryptSaltToken(newPassword, userAgentSecretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("admin_user").
		Set("password", encryptNewPassword).
		Where(sq.And{
			sq.Eq{"username": req.Username},
		}).
		ToSql()

	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["username"] = req.Username
	c.Set("action_log", actionLog)

	// 密碼修改成功後要將用戶登出，強迫使用新密碼重新登入
	c.Jwt.AddBlackTokenByUsername(req.Username)

	c.Ok(&model.ResetPasswordResponse{
		Username:    req.Username,
		NewPassword: newPassword,
	})
}

// @Tags 運營管理/玩家帳號相關
// @Summary 取得玩家目前餘額
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUserBalanceRequest true "取得玩家目前餘額參數"
// @Success 200 {object} response.Response{data=model.GetGameUserBalanceResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getgameuserbalance [post]
func (p *SystemUserApi) GetGameUserBalance(c *ginweb.Context) {
	req := &model.GetGameUserBalanceRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil || claims.WalletType == definition.AGENT_WALLET_SINGLE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("level_code").
		From("game_users").
		Where(sq.Eq{"id": req.UserId}).
		ToSql()

	userAgentLevelCode := ""
	err := c.DB().QueryRow(query, args...).Scan(&userAgentLevelCode)
	if err != nil {
		code := definition.ERROR_CODE_ERROR_DATABASE
		if err == sql.ErrNoRows {
			code = definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST
		}

		c.OkWithCode(code)
		return
	}

	myAgentLevelCode := claims.LevelCode
	if !strings.HasPrefix(userAgentLevelCode, myAgentLevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	gold, _, err := notification.GetGameServerGold(int64(req.UserId))
	if err != nil {
		c.Logger().Info("GetGameUserBalance notification GetGameServerGold failed, userId=%d, err=%v", req.UserId, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	resp := new(model.GetGameUserBalanceResponse)
	resp.UserId = req.UserId
	resp.WalletBalance = gold

	c.Ok(resp)
}

// @Tags 運營管理/玩家帳號相關
// @Summary 此接口用來取得玩家目前遊戲局數狀態
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUserPlayCountDataRequest true "取得玩家目前遊戲局數狀態"
// @Success 200 {object} response.Response{data=model.GetGameUserPlayCountDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/user/getgameuserplaycountdata [post]
func (p *SystemUserApi) GetGameUserPlayCountData(c *ginweb.Context) {
	req := &model.GetGameUserPlayCountDataRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("level_code").
		From("game_users").
		Where(sq.Eq{"id": req.UserId}).
		ToSql()

	userAgentLevelCode := ""
	err := c.DB().QueryRow(query, args...).Scan(&userAgentLevelCode)
	if err != nil {
		code := definition.ERROR_CODE_ERROR_DATABASE
		if err == sql.ErrNoRows {
			code = definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST
		}

		c.OkWithCode(code)
		return
	}

	// 管理後台、代理後台皆可使用
	myAgentLevelCode := claims.LevelCode
	if !strings.HasPrefix(userAgentLevelCode, myAgentLevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	playCountData, limit, apiCode, err := notification.GetGameServerUserIsNewbie(int64(req.UserId))
	code := notification.TransformApiCodeToModelErrorCode(apiCode)
	if err != nil {
		c.Logger().Info("GetGameUserPlayCountData notification GetGameServerUserIsNewbie failed, userId=%d, apiCode=%v, err=%v", req.UserId, apiCode, err)
		c.OkWithCode(code)
		return
	}

	resp := new(model.GetGameUserPlayCountDataResponse)
	resp.UserId = req.UserId
	resp.TotalNewbieLimit = limit
	if succ := resp.DataConvert(playCountData); !succ {
		c.Logger().Info("GetGameUserPlayCountData DataConvert failed, userId=%d, playCountData=%v", req.UserId, playCountData)
	}

	c.Ok(resp)
}
