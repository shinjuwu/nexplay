package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/server/global"
	table_model "backend/server/table/model"
	"definition"
	"fmt"
	"sort"
	"strings"
	"time"

	"backend/pkg/encrypt"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemAgentApi struct {
	BasePath string
}

func NewSystemAgentApi(basePath string) api_cluster.IApiEach {
	return &SystemAgentApi{
		BasePath: basePath,
	}
}

func (p *SystemAgentApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemAgentApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/createagent",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateAgent))
	g.POST("/getagentlist", ginHandler.Handle(p.GetAgentList))
	g.POST("/getagentsecretkey", ginHandler.Handle(p.GetAgentSecretKey))
	g.POST("/getagentcoinsupplyinfo", ginHandler.Handle(p.GetAgentCoinSupplyInfo))
	g.POST("/setagentcoinsupplyinfo",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentCoinSupplyInfo))
	g.POST("/getagentgamelist", ginHandler.Handle(p.GetAgentGameList))
	g.POST("/setagentgamestate",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentGameState))
	g.POST("/getagentgameroomlist", ginHandler.Handle(p.GetAgentGameRoomList))
	g.POST("/setagentgameroomstate",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentGameRoomState))
	g.GET("/getagentpermissiontemplateinfo", ginHandler.Handle(p.GetAgentPermissionTemplateInfo))
	g.POST("/getagentpermissionlist", ginHandler.Handle(p.GetAgentPermissionList))
	g.POST("/getagentpermission", ginHandler.Handle(p.GetAgentPermission))
	g.POST("/createagentpermission",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateAgentPermission))
	g.POST("/setagentpermission",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentPermission))
	g.POST("/deleteagentpermission",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.DeleteAgentPermission))
	g.POST("/getagentipwhitelistlist", ginHandler.Handle(p.GetAgentIpWhitelistList))
	g.POST("/getagentipwhitelist", ginHandler.Handle(p.GetAgentIpWhitelist))
	g.POST("/setagentipwhitelist",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentIpWhitelist))
	g.POST("/getagentapiipwhitelist", ginHandler.Handle(p.GetAgentApiIpWhitelist))
	g.POST("/setagentapiipwhitelist",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentApiIpWhitelist))
	g.POST("/getagentwalletlist", ginHandler.Handle(p.GetAgentWalletList))
	g.POST("/setagentwallet", ginHandler.Handle(p.SetAgentWallet))
}

// @Tags 代理帳號管理/代理帳號
// @Summary 創建代理帳號
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.CreateAgentRequest true "添加代理必要資訊"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/createagent [post]
func (p *SystemAgentApi) CreateAgent(c *ginweb.Context) {
	req := model.NewCreateAgentRequest()
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if !req.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查錢包類型與錢包設定URL
	// 只有單一錢包要檢查
	if req.WalletType == definition.AGENT_WALLET_SINGLE {
		if success, _ := req.WalletConnInfo.ParseUrl(req.WalletUrl); !success {
			c.OkWithCode(definition.ERROR_CODE_ERROR_WALLET_URL_PARSE_FAILED)
			return
		}
	} else {
		req.WalletConnInfo = new(table_model.WalletConnInfo)
	}

	// 創建後臺帳號時一律轉小寫
	req.Account = strings.ToLower(req.Account)

	myAgentId := -1 // default  最高權限
	myAccountType := 0
	levelCode := ""
	creator := ""
	cooperation := req.Cooperation
	currency := req.Currency
	lobbySwitchInfo := req.LobbySwitchInfo
	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	// 子代理以下不可再創
	if claims.AccountType >= definition.ACCOUNT_TYPE_NORMAL {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 檢查代理名稱有無重複
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("agent").
		Where(sq.And{
			sq.Eq{"name": req.Nickname},
		}).
		Suffix(")").
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var isExists bool
	err = c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_AGENT_NAME_EXIST)
		return
	}

	// 檢查後台帳號有無重複
	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("admin_user").
		Where(sq.And{
			sq.Eq{"username": req.Account},
		}).
		Suffix(")").
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	err = c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_EXIST)
		return
	}

	myAgentId = int(claims.BaseClaims.ID)
	myAccountType = claims.AccountType
	// 上層代理 level code + 上層代理id 轉16進位碼(字母小寫)
	// levelCode = claims.LevelCode + fmt.Sprintf("%04x", topAgentId)
	levelCode = claims.LevelCode

	// 只有開發者可以設定合作模式(代理結帳類型, 1: 買分, 2: 信用)
	// 不是開發者，合作模式一律繼承上級設定
	if myAccountType != definition.ACCOUNT_TYPE_ADMIN {
		cooperation = claims.Cooperation
		currency = claims.Currency

		myAgent := global.AgentCache.Get(myAgentId)
		lobbySwitchInfo = myAgent.LobbySwitchInfo
	}

	// 檢查幣種
	exchangeData := global.ExchangeDataCache.Get(currency)
	if exchangeData == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CURRENCY_NOT_SUPPORTED)
		return
	}

	creator = c.GetUsername()
	/*
		產生必要資訊
	*/
	timeNowStr := time.Now().String()
	// 8碼  代理編號
	agentCode := md5.Hash8bit(req.Nickname + timeNowStr)
	// 16碼 md5 key
	md5Key := md5.Hash16bit(agentCode + timeNowStr)
	// 16碼 aes key
	aesKey := md5.Hash16bit(md5Key + timeNowStr)
	// 32碼 secret key
	secretKey := md5.Hash32bit(aesKey + timeNowStr)

	// 將密碼做加密
	decryptPwd, err := encrypt.EncryptSaltToken(req.Password, secretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	if global.CheckAdminUserIsExist(c.DB(), req.Account) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_EXIST)
		return
	}

	// 建立ip白名單
	now := utils.GetUnixTimeNowUTC()
	ipWhitelist := make([]*table_model.AgentIPWhitelistObj, 0)
	for _, ipAddress := range strings.Split(req.IpWhitelist, ";") {
		if !validIpAddress(ipAddress) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_IP_FORMAT)
			return
		}

		ipWhitelist = append(ipWhitelist, &table_model.AgentIPWhitelistObj{
			CreateTime: now,
			IPAddress:  ipAddress,
			Creator:    creator,
		})
	}

	// 開發者代理只能用SQL語法創建, 所以 isTopAgent = false
	isTopAgent := false
	agent, adminUser, agentGames, agentGameRooms, err := global.UdfCreateNewAgent(c.DB(),
		req.Nickname, agentCode, levelCode, req.Info, secretKey,
		aesKey, md5Key, currency, ipWhitelist, creator,
		req.Commission, cooperation, myAgentId, isTopAgent, req.WalletType,
		req.WalletConnInfo, lobbySwitchInfo, req.Account, decryptPwd, req.Nickname, req.Role,
		req.Info, myAccountType+1, definition.ACCOUNT_READONLY_OFF)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	global.AgentCache.Add(agent)
	global.AdminUserCache.Add(adminUser)

	if err := global.AgentGameRatioCache.CreateNewAgentToCacheFromJson(agent.Id); err != nil {
		c.Logger().Printf("CreateNewAgentToCacheFromJson() has error: %v", err)
	}

	if err := global.AgentCustomTagInfoCache.CreateNewAgentToCacheFromJson(agent.Id, agent.LevelCode); err != nil {
		c.Logger().Printf("CreateNewAgentToCacheFromJson() has error: %v", err)
	}

	if err := global.AgentGameIconListCache.CreateNewAgentToCacheFromJson(myAgentId, agent.Id, agent.LevelCode, false, true); err != nil {
		c.Logger().Printf("CreateNewAgentToCacheFromJson() has error: %v", err)
	}

	gameInfos := make([]map[string]interface{}, 0)
	for i := 0; i < len(agentGames); i++ {
		global.AgentGameCache.Add(agentGames[i])

		game := global.GameCache.Get(agentGames[i].GameId)
		gameInfos = append(
			gameInfos,
			getGameInfo(
				agentGames[i].AgentId,
				agentGames[i].GameId,
				game.Code,
				agentGames[i].State,
			),
		)
	}

	lobbyInfos := make([]map[string]interface{}, 0)
	killDiveInfo := make([]map[string]interface{}, 0)

	for i := 0; i < len(agentGameRooms); i++ {
		global.AgentGameRoomCache.Add(agentGameRooms[i])

		gameRoom := global.GameRoomCache.Get(agentGameRooms[i].GameRoomId)
		if gameRoom == nil {
			c.Logger().Error("CreateAgent agentGameRoom not find game room, agentId: %d, gameRoomId: %d", agentGameRooms[i].AgentId, agentGameRooms[i].GameRoomId)
			continue
		}

		game := global.GameCache.Get(gameRoom.GameId)
		if game == nil {
			c.Logger().Error("CreateAgent gameRoom not find game, agentId: %d, gameId: %d", agentGameRooms[i].AgentId, gameRoom.GameId)
			continue
		}

		agentId := agentGameRooms[i].AgentId
		gameId := game.Id

		lobbyInfos = append(
			lobbyInfos,
			getLobbyInfo(
				agentId,
				gameId,
				gameRoom.Id,
				agentGameRooms[i].State,
			),
		)

		// 百人類及好友房遊戲不用做殺放處理
		if game.Type == definition.GAME_TYPE_BAIREN ||
			game.Type == definition.GAME_TYPE_FRIENDSROOM {
			continue
		}

		gameType := game.Type
		roomType := gameRoom.RoomType

		agentGameRatioKey := global.AgentGameRatioCache.GetKey(agentId, gameId, gameType, roomType)
		agentGameRatio, ok := global.AgentGameRatioCache.SelectOne(agentGameRatioKey)
		if !ok {
			c.Logger().Error("CreateAgent agentGameRatio not find, agentId: %d, gameId: %d, gameType: %d, roomType: %d, key: %s", agentId, gameId, gameType, roomType, agentGameRatioKey)
			continue
		}

		roomId, convErr := getRoomId(agentGameRatio.GameId, agentGameRatio.RoomType)
		if convErr != nil {
			c.Logger().Error("CreateAgent agentGameRatio getRoomId err, agentId: %d, gameId: %d, roomType: %d, err: %v", agentGameRatio.AgentId, agentGameRatio.GameId, agentGameRatio.RoomType, convErr)
			continue
		}

		killDiveInfo = append(
			killDiveInfo,
			getKillDiveInfo(
				agentGameRatio.AgentId,
				agentGameRatio.GameId,
				roomId,
				agentGameRatio.ActiveNum,
				agentGameRatio.KillRatio,
				agentGameRatio.NewKillRatio,
			),
		)
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(agent.Name)
	c.Set("action_log", actionLog)

	// 通知game server新建立的代理遊戲設定
	if code, _ := notification.SendSetGameInfo(gameInfos); code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Info("create agent notification.SendSetGameInfo fail, gameInfos=%v, resp code=%d", gameInfos, code)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}
	if code, _ := notification.SendSetLobbyInfo(lobbyInfos); code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Info("create agent notification.SendSetLobbyInfo fail, lobbyInfos=%v, resp code=%d", lobbyInfos, code)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}
	if apiResult, errCode, err := notification.SendJapotInfo([]int{agent.Id}, []int64{agent.JackpotStartTime.UnixMilli()}, []int64{agent.JackpotEndTime.UnixMilli()}); errCode != definition.ERROR_CODE_SUCCESS {
		c.Logger().Info("create agent notification.SendJapotInfo fail, agent id=%d, start=%v, end=%v, resp=%s, resp code=%d, err=%v", agent.Id, agent.JackpotStartTime, agent.JackpotEndTime, apiResult, errCode, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}
	// TO DO: 追蹤Bug問題加入的Log，問題解決後要移除
	c.Logger().Info("[代理殺放追蹤]新增代理設置殺放，殺放內容: %v", killDiveInfo)
	if apiResult, errCode, err := notification.SendSetKillDives(killDiveInfo); errCode != definition.ERROR_CODE_SUCCESS {
		c.Logger().Info("create agent notification.SendSetKillDives fail, agent id=%d, resp=%s, resp code=%d, err=%v", agent.Id, apiResult, errCode, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	c.Ok("")
}

// @Tags 代理帳號管理/代理帳號
// @Summary 取得代理底下所有代理資料
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.GetAgentSecretKeyRequest true "指定代理id"
// @Success 200 {object} response.Response{data=[]model.GetAgentListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentlist [post]
func (p *SystemAgentApi) GetAgentList(c *ginweb.Context) {
	req := model.NewGetAgentListRequest()

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

	myAgentId := int(claims.BaseClaims.ID)
	myAgentObj := global.AgentCache.Get(myAgentId)
	if myAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	if req.Id == 0 {
		req.Id = myAgentId
	}
	targetAgentObj := global.AgentCache.Get(req.Id)
	if targetAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查權限
	isOk := global.CheckTargetLevelCodeIsPassing(myAgentObj.LevelCode, targetAgentObj.LevelCode)
	if !isOk {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	res := model.NewGetAgentListResponseSlice()

	agents := global.AgentCache.GetAll()
	for _, v := range agents {
		// 目標權限比較小或相等
		myLen := len(targetAgentObj.LevelCode)
		targetLen := len(v.LevelCode)
		if targetLen >= myLen {
			if v.LevelCode[0:myLen] == targetAgentObj.LevelCode {
				tmp := model.NewGetAgentListResponse()
				if v.TopAgentId == -1 {
					tmp.TopAgentName = v.Name
				} else {
					at := global.AgentCache.Get(v.TopAgentId)
					if at != nil {
						tmp.TopAgentName = at.Name
					}
				}

				vAdminUser := global.AdminUserCache.Get(v.Id, v.AdminUsername)
				if vAdminUser != nil {
					vAgentPermission := global.AgentPermissionCache.Get(vAdminUser.PermissionId)
					if vAgentPermission != nil {
						tmp.RoleName = vAgentPermission.Name
					}
				}

				tmp.TransVal(v)
				res = append(res, tmp)
			}
		}
	}

	// 查詢代理錢包餘額
	query := `select agent_id, amount from agent_wallet where 1=1;`
	rows, err := c.DB().Query(query)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS)
		return
	}
	defer rows.Close()

	// map["agentId"]amount
	var data = make(map[int]float64, 0)
	for rows.Next() {
		var agentId int
		var amount float64
		if err := rows.Scan(&agentId, &amount); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}
		data[agentId] = amount
	}
	for _, v := range res {
		amo, ok := data[v.Id]
		if ok {
			v.Amount = amo
		}
	}

	c.Ok(res)
}

// @Tags 代理帳號管理/代理帳號
// @Summary 秘鑰資訊顯示
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.GetAgentSecretKeyRequest true "指定代理id(0:找自己)"
// @Success 200 {object} response.Response{data=model.GetAgentSecretKeyResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentsecretkey [post]
func (p *SystemAgentApi) GetAgentSecretKey(c *ginweb.Context) {

	req := model.NewGetAgentSecretKeyRequest()

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

	myAgentId := int(claims.BaseClaims.ID)
	myAgentObj := global.AgentCache.Get(myAgentId)
	if req.Id == 0 {
		req.Id = myAgentId
	}
	targetAgentObj := global.AgentCache.Get(req.Id)
	if targetAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查權限
	isOk := global.CheckTargetLevelCodeIsPassing(myAgentObj.LevelCode, targetAgentObj.LevelCode)
	if !isOk {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	res := model.NewGetAgentSecretKeyResponse()
	res.AesKey = targetAgentObj.AesKey
	res.Md5Key = targetAgentObj.Md5Key

	c.Ok(res)
}

// @Tags 代理帳號管理/代理帳號
// @Summary 取得指定代理補分相關資料設定
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.GetAgentCoinSupplyInfoRequest true "指定代理id"
// @Success 200 {object} response.Response{data=model.GetAgentCoinSupplyInfoResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentcoinsupplyinfo [post]
func (p *SystemAgentApi) GetAgentCoinSupplyInfo(c *ginweb.Context) {
	req := model.NewGetAgentCoinSupplyInfoRequest()

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

	myAgentId := int(claims.BaseClaims.ID)
	if req.Id == 0 {
		req.Id = myAgentId
	}
	myAgentObj := global.AgentCache.Get(myAgentId)
	if myAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	targetAgentObj := global.AgentCache.Get(req.Id)
	if targetAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查權限
	isOk := global.CheckTargetLevelCodeIsPassing(myAgentObj.LevelCode, targetAgentObj.LevelCode)
	if !isOk {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	adminUser := global.AdminUserCache.Get(targetAgentObj.Id, targetAgentObj.AdminUsername)
	if adminUser == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	agentPermission := global.AgentPermissionCache.Get(adminUser.PermissionId)
	if agentPermission == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	res := model.NewGetAgentCoinSupplyInfoResponse()
	res.Name = targetAgentObj.Name
	res.Commission = targetAgentObj.Commission
	res.CoinLimit = targetAgentObj.CoinLimit
	res.CoinUse = targetAgentObj.CoinUse
	res.Cooperation = targetAgentObj.Cooperation
	res.Info = targetAgentObj.Info
	res.TopAgentId = targetAgentObj.TopAgentId
	res.Role = adminUser.PermissionId
	res.RoleName = agentPermission.Name
	res.WalletType = targetAgentObj.WalletType
	res.WalletConnInfo = targetAgentObj.WalletConnInfo.GetUrlPath()
	res.LobbySwitchInfo = targetAgentObj.LobbySwitchInfo

	c.Ok(res)
}

// @Tags 代理帳號管理/代理帳號
// @Summary 修改指定代理補分相關資料設定
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.SetAgentCoinSupplyInfoRequest true "指定代理id"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentcoinsupplyinfo [post]
func (p *SystemAgentApi) SetAgentCoinSupplyInfo(c *ginweb.Context) {
	req := model.NewSetAgentCoinSupplyInfoRequest()

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

	myAgentId := int(claims.BaseClaims.ID)
	myAgentObj := global.AgentCache.Get(myAgentId)

	targetAgentObj := global.AgentCache.Get(req.Id)
	if targetAgentObj == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查錢包類型，單一錢包必須要有api url
	wTmp := new(table_model.WalletConnInfo)
	if targetAgentObj.WalletType == definition.AGENT_WALLET_SINGLE {
		if req.WalletConnInfo == "" {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		if success, _ := wTmp.ParseUrl(req.WalletConnInfo); !success {
			c.OkWithCode(definition.ERROR_CODE_ERROR_WALLET_URL_PARSE_FAILED)
			return
		}
	}

	// 檢查權限
	isOk := global.CheckTargetLevelCodeIsPassing(myAgentObj.LevelCode, targetAgentObj.LevelCode)
	if !isOk {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 不能設定自己
	isMy := myAgentObj.Id == targetAgentObj.Id
	if isMy {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 總代理、子代理不能設定大廳開關
	if len(myAgentObj.LevelCode) > 4 {
		req.LobbySwitchInfo = targetAgentObj.LobbySwitchInfo
	}

	checkLobbyInfoChanged := (targetAgentObj.LobbySwitchInfo != req.LobbySwitchInfo)
	if checkLobbyInfoChanged {
		gameServerInfoStorage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO)
		if !ok {
			c.Logger().Error("NotifyGameServer() STORAGE_KEY_GAMESERVERINFO is empty")
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}
		resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
		cs, ok := resp["state"].(float64)
		if !ok {
			c.Logger().Error("getGameServerState code:%d", definition.ERROR_CODE_ERROR_DATABASE)
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
		state := int(cs)

		if state != definition.GAME_STATE_MAINTAIN {
			c.OkWithCode(definition.ERROR_CODE_ERROR_GAME_SERVER_NOT_IN_MAINTENANCE)
			return
		}
	}

	// 檢查名稱有無重複
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("agent").
		Where(sq.And{
			sq.NotEq{"id": req.Id},
			sq.Eq{"name": req.Name},
		}).
		Suffix(")").
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var isExists bool
	err = c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_AGENT_NAME_EXIST)
		return
	}

	targetAdminUserObj := global.AdminUserCache.Get(targetAgentObj.Id, targetAgentObj.AdminUsername)
	checkRoleChanged := (targetAdminUserObj.PermissionId != req.Role)

	agentUpdateTime, adminUserUpdateTime, err := global.UdfUpdateAgent(c.DB(), targetAgentObj.Id, req.Name, req.Info, req.Commission,
		targetAdminUserObj.Username, req.Role, targetAdminUserObj.Info, targetAdminUserObj.IsEnabled, checkRoleChanged,
		utils.ToJSON(wTmp), req.LobbySwitchInfo)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Name)
	if targetAgentObj.Name != req.Name {
		actionLog["name"] = createBancendActionLogDetail(targetAgentObj.Name, req.Name)
	}
	if targetAgentObj.Commission != req.Commission {
		actionLog["commission"] = createBancendActionLogDetail(targetAgentObj.Commission, req.Commission)
	}
	if targetAgentObj.Info != req.Info {
		actionLog["info"] = createBancendActionLogDetail(targetAgentObj.Info, req.Info)
	}
	if targetAdminUserObj.PermissionId != req.Role {
		beforeAgentPermission := global.AgentPermissionCache.Get(targetAdminUserObj.PermissionId)
		afterAgentPermission := global.AgentPermissionCache.Get(req.Role)
		actionLog["agent_permission"] = createBancendActionLogDetail(beforeAgentPermission.Name, afterAgentPermission.Name)
	}
	if targetAgentObj.WalletType == definition.AGENT_WALLET_SINGLE {
		if targetAgentObj.WalletConnInfo.ApiKey != wTmp.ApiKey ||
			targetAgentObj.WalletConnInfo.Domain != wTmp.Domain ||
			targetAgentObj.WalletConnInfo.Path != wTmp.Path ||
			targetAgentObj.WalletConnInfo.Scheme != wTmp.Scheme {
			actionLog["wallet_conninfo"] = createBancendActionLogDetail(targetAgentObj.WalletConnInfo, wTmp)
		}
	}
	if targetAgentObj.LobbySwitchInfo != req.LobbySwitchInfo {
		actionLog["lobby_switch_info"] = createBancendActionLogDetail(targetAgentObj.LobbySwitchInfo, req.LobbySwitchInfo)
	}
	c.Set("action_log", actionLog)

	// 修改cache資料，agent是從redis取出，所以更新玩從新add進行更新
	targetAgentObj.Name = req.Name
	targetAgentObj.Commission = req.Commission
	// targetAgentObj.Cooperation = req.Cooperation
	// targetAgentObj.CoinSupplySetting = req.CoinSupplySetting
	targetAgentObj.Info = req.Info
	targetAgentObj.UpdateTime = agentUpdateTime
	targetAgentObj.WalletConnInfo = wTmp
	targetAgentObj.WalletConnInfoBytes = []byte(utils.ToJSON(wTmp))
	global.AgentCache.Add(targetAgentObj)

	targetAdminUserObj.PermissionId = req.Role
	targetAdminUserObj.UpdateTime = adminUserUpdateTime

	// 權限群組變更則要將該代理及下層級的的所有後台帳號加入黑名單強迫重新登入及修正權限cache
	if checkRoleChanged {
		agents := make([]*table_model.Agent, 0)
		agents = append(agents, targetAgentObj)

		childAgents := global.AgentCache.GetChildAgents(targetAgentObj.Id)
		if childAgents != nil {
			agents = append(agents, childAgents...)
		}

		targetAgentPermission := global.AgentPermissionCache.Get(req.Role)
		for _, agent := range agents {
			for _, agentPermission := range global.AgentPermissionCache.GetByAgentId(agent.Id) {
				if agentPermission == targetAgentPermission {
					continue
				}

				agentPermission.Permission.List = utils.ArrayIntersection(agentPermission.Permission.List, targetAgentPermission.Permission.List)
			}

			for _, adminUser := range global.AdminUserCache.GetAgentAdminUsers(agent.Id) {
				c.Jwt.AddBlackTokenByUsername(adminUser.Username)
			}
		}
	}

	if checkLobbyInfoChanged {
		agents := make([]*table_model.Agent, 0)
		agents = append(agents, targetAgentObj)

		childAgents := global.AgentCache.GetChildAgents(targetAgentObj.Id)
		if childAgents != nil {
			agents = append(agents, childAgents...)
		}

		for _, agent := range agents {
			agent.LobbySwitchInfo = req.LobbySwitchInfo
		}

		global.AgentCache.Adds(agents)
	}

	c.Ok("")
}

// @Tags 遊戲設置/遊戲管理
// @Summary 取得代理遊戲列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentGameListRequest true "取得代理遊戲列表參數"
// @Success 200 {object} response.Response{data=model.DataTablesResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentgamelist [post]
func (p *SystemAgentApi) GetAgentGameList(c *ginweb.Context) {
	var req model.GetAgentGameListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userCliams := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userCliams.BaseClaims.ID))

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_gameId := (req.GameId > definition.GAME_ID_ALL)
	checkExist_state := (req.State > definition.GAME_STATE_ALL)

	sqAnd := sq.And{}
	// 非自己或下層級不能查詢
	if checkExist_agentId {
		agent := global.AgentCache.Get(req.AgentId)
		if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		sqAnd = append(sqAnd, sq.Eq{"agent_id": agent.Id})
	} else {
		sqAnd = append(sqAnd, sq.Like{"agent_level_code": userAgent.LevelCode + "%"})
	}

	// 下架的遊戲不能查
	if checkExist_gameId {
		game := global.GameCache.Get(req.GameId)
		if game == nil || game.State == definition.GAME_STATE_OFFLINE {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		sqAnd = append(sqAnd, sq.Eq{"game_id": req.GameId})
	} else {
		sqAnd = append(sqAnd, sq.Eq{"game_state": definition.GAME_STATES_ONLINE_MAINTAIN})
	}

	// 非系統商只能顯示線上及維護中的資料
	if checkExist_state {
		if userCliams.AccountType != definition.ACCOUNT_TYPE_ADMIN && req.State == definition.GAME_STATE_OFFLINE {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		sqAnd = append(sqAnd, sq.Eq{"state": req.State})
	}

	resp := &model.DataTablesResponse{}
	resp.Draw = req.Draw
	resp.Data = make([]interface{}, 0)

	// 限制只有第一次會進行加總
	if req.Draw == 0 {
		recordsTotal, errorCode := getAgentGameListSumInfo(c.DB(), &sqAnd)
		if errorCode != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(errorCode)
			return
		}

		resp.RecordsTotal = recordsTotal
	}

	data, errorCode := getAgentGameList(c.DB(), &req, &sqAnd)
	if errorCode != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(errorCode)
		return
	}

	for i := 0; i < len(data); i++ {
		resp.Data = append(resp.Data, data[i])
	}

	c.Ok(resp)
}

// @Tags 遊戲設置/遊戲管理
// @Summary 設置代理遊戲狀態
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetAgentGameStateRequest true "設置代理遊戲狀態必要資訊"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentgamestate [post]
func (p *SystemAgentApi) SetAgentGameState(c *ginweb.Context) {
	var req model.SetAgentGameStateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userCliams := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userCliams.BaseClaims.ID))

	// 非系統商不能設定下架
	if userCliams.AccountType != definition.ACCOUNT_TYPE_ADMIN &&
		req.State == definition.GAME_STATE_OFFLINE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查設定，並過濾掉上級有設定的
	// 下架的遊戲不能設定、非下層級或自己不能設定
	sort.SliceStable(req.List, func(i, j int) bool {
		return req.List[i].AgentId < req.List[j].AgentId
	})

	gameAgentLevelCodes := make(map[int][]string)
	reqAgentGames := make([]*model.SetAgentGameStateGameRequest, 0)

	for _, item := range req.List {
		itemAgent := global.AgentCache.Get(item.AgentId)
		itemGame := global.GameCache.Get(item.GameId)

		if itemGame == nil || itemGame.State == definition.GAME_STATE_OFFLINE ||
			itemAgent == nil || !strings.HasPrefix(itemAgent.LevelCode, userAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		if _, findGame := gameAgentLevelCodes[itemGame.Id]; !findGame {
			gameAgentLevelCodes[itemGame.Id] = make([]string, 0)
		}

		// 上級有設定就過濾掉
		findAgent := false
		for _, levelCode := range gameAgentLevelCodes[itemGame.Id] {
			if strings.HasPrefix(itemAgent.LevelCode, levelCode) {
				findAgent = true
				break
			}
		}

		if !findAgent {
			gameAgentLevelCodes[itemGame.Id] = append(gameAgentLevelCodes[itemGame.Id], itemAgent.LevelCode)
			reqAgentGames = append(reqAgentGames, item)
		}
	}

	gameInfos := make([]map[string]interface{}, 0)
	actionLogGames := make([]map[string]interface{}, 0)

	var sqlSb strings.Builder
	sqlTopAgentFmt := `CALL "public"."procedure_update_agent_game_state_by_top_agent"(%d,'%s',%d,%d);`
	sqlAgentFmt := `CALL "public"."procedure_update_agent_game_state_by_agent"(%d,%d,%d);`

	for _, reqAgentGame := range reqAgentGames {
		agent := global.AgentCache.Get(reqAgentGame.AgentId)
		game := global.GameCache.Get(reqAgentGame.GameId)
		agentGame := global.AgentGameCache.Get(reqAgentGame.AgentId, reqAgentGame.GameId)

		if agentGame.State == req.State {
			continue
		}

		// 上級代理非系統代理需要檢查上級代理的設定，系統代理不會有設定
		topAgentGame := global.AgentGameCache.Get(agent.TopAgentId, game.Id)
		if topAgentGame != nil {
			if topAgentGame.State == definition.GAME_STATE_MAINTAIN &&
				req.State == definition.GAME_STATE_ONLINE {
				c.OkWithCode(definition.ERROR_CODE_ERROR_TOP_AGENT_SETTING)
				return
			} else if topAgentGame.State == definition.GAME_STATE_OFFLINE &&
				req.State != definition.GAME_STATE_OFFLINE {
				c.OkWithCode(definition.ERROR_CODE_ERROR_TOP_AGENT_SETTING)
				return
			}
		}

		// gameInfo加入自己
		gameInfos = append(gameInfos, getGameInfo(agent.Id, game.Id, game.Code, req.State))

		// gameInfo加入子代
		childAgents := global.AgentCache.GetChildAgents(agent.Id)
		if len(childAgents) > 0 {
			for _, childAgent := range childAgents {
				childAgentGame := global.AgentGameCache.Get(childAgent.Id, reqAgentGame.GameId)

				if childAgentGame.State == req.State {
					continue
				}

				gameInfos = append(gameInfos, getGameInfo(childAgent.Id, game.Id, game.Code, req.State))
			}

			sqlSb.WriteString(fmt.Sprintf(sqlTopAgentFmt, agent.Id, agent.LevelCode, game.Id, req.State))
		} else {
			sqlSb.WriteString(fmt.Sprintf(sqlAgentFmt, agent.Id, game.Id, req.State))
		}

		actionLogGames = append(actionLogGames, map[string]interface{}{
			"agent_name": agent.Name,
			"game_id":    game.Id,
			"state":      createBancendActionLogDetail(agentGame.State, req.State),
		})
	}

	if len(gameInfos) > 0 {
		_, err := c.DB().Exec(sqlSb.String())
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if code, _ := notification.SendSetGameInfo(gameInfos); code != definition.ERROR_CODE_SUCCESS {
			c.Logger().Info("notification.SendSetGameInfo fail, gameInfos=%v, resp code=%d", gameInfos, code)

			// 通知失敗要重設DB資料
			sqlSb.Reset()
			for _, gameInfo := range gameInfos {
				agentId := gameInfo["AgentId"].(int)
				gameId := gameInfo["GameId"].(int)
				agentGame := global.AgentGameCache.Get(agentId, gameId)

				sqlSb.WriteString(fmt.Sprintf(sqlAgentFmt, agentId, gameId, agentGame.State))
			}

			_, err := c.DB().Exec(sqlSb.String())
			if err != nil {
				c.Logger().Warn("notification.SendSetGameInfo fail and db rollback fail, req=%v, sql=%s", req, sqlSb.String())
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}

			// 重設DB資料完成才傳通知送失敗
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	if len(actionLogGames) > 0 {
		actionLog["agent_games"] = actionLogGames
	}
	c.Set("action_log", actionLog)

	// 全部成功才更新cache
	for _, gameInfo := range gameInfos {
		agentId := gameInfo["AgentId"].(int)
		gameId := gameInfo["GameId"].(int)

		agentGame := global.AgentGameCache.Get(agentId, gameId)
		agentGame.State = req.State
	}

	c.Ok(nil)
}

// @Tags 遊戲設置/遊戲管理
// @Summary 取得代理遊戲房間列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentGameRoomListRequest true "取得代理遊戲房間列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentGameRoomResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentgameroomlist [post]
func (p *SystemAgentApi) GetAgentGameRoomList(c *ginweb.Context) {
	var req model.GetAgentGameRoomListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userCliams := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userCliams.BaseClaims.ID))

	// 非自己或下層級不能查詢
	agent := global.AgentCache.Get(req.AgentId)
	if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 下架的遊戲不能查詢
	game := global.GameCache.Get(req.GameId)
	if game == nil || game.State == definition.GAME_STATE_OFFLINE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	sqAnd := sq.And{
		sq.Eq{"agent_id": req.AgentId},
		sq.Eq{"game_id": req.GameId},
	}

	// 非開發商不能查詢被關閉的設定
	if userCliams.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		sqAnd = append(sqAnd, sq.Eq{"state": definition.GAME_STATES_ONLINE_MAINTAIN})
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "agent_name", "game_id", "game_room_id", "room_type", "state").
		From("view_agent_game_room").
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

	defer rows.Close()

	data := make([]*model.GetAgentGameRoomResponse, 0)
	for rows.Next() {
		temp := model.GetAgentGameRoomResponse{}
		if err := rows.Scan(&temp.AgentId, &temp.AgentName, &temp.GameId,
			&temp.GameRoomId, &temp.RoomType, &temp.State); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
		data = append(data, &temp)
	}

	c.Ok(data)
}

// @Tags 遊戲設置/遊戲管理
// @Summary 設置代理遊戲房間狀態
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetAgentGameRoomStateRequest true "設置代理遊戲房狀態必要資訊"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentgameroomstate [post]
func (p *SystemAgentApi) SetAgentGameRoomState(c *ginweb.Context) {
	var req model.SetAgentGameRoomStateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userCliams := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userCliams.BaseClaims.ID))

	// 非自己或下層級不能查詢
	agent := global.AgentCache.Get(req.AgentId)
	if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 下架的遊戲不能查詢
	game := global.GameCache.Get(req.GameId)
	if game == nil || game.State == definition.GAME_STATE_OFFLINE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 遊戲房間要能對上遊戲
	for _, item := range req.List {
		itemGameRoom := global.GameRoomCache.Get(item.GameRoomId)
		if itemGameRoom == nil || game.Id != itemGameRoom.GameId {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
	}

	gameRooms := global.GameRoomCache.GetGameRoomMaps(game.Id)
	lobbyInfos := make([]map[string]interface{}, 0)
	actionLogGameRooms := make([]map[string]interface{}, 0)

	var sqlSb strings.Builder
	sqlTopAgentFmt := `CALL "public"."procedure_update_agent_game_room_state_by_top_agent"(%d,'%s',%d,%d);`
	sqlAgentFmt := `CALL "public"."procedure_update_agent_game_room_state_by_agent"(%d,%d,%d);`

	for _, reqAgentGameRoom := range req.List {
		gameRoom := gameRooms[reqAgentGameRoom.GameRoomId]
		agentGameRoom := global.AgentGameRoomCache.Get(req.AgentId, gameRoom.Id)

		if agentGameRoom.State == reqAgentGameRoom.State {
			continue
		}

		// 上級代理非系統代理需要檢查上級代理的設定，系統代理不會有設定
		topAgentGameRoom := global.AgentGameRoomCache.Get(agent.TopAgentId, gameRoom.Id)
		if topAgentGameRoom != nil {
			if topAgentGameRoom.State == definition.GAME_STATE_MAINTAIN &&
				reqAgentGameRoom.State == definition.GAME_STATE_ONLINE {
				c.OkWithCode(definition.ERROR_CODE_ERROR_TOP_AGENT_SETTING)
				return
			} else if topAgentGameRoom.State == definition.GAME_STATE_OFFLINE &&
				reqAgentGameRoom.State != definition.GAME_STATE_OFFLINE {
				c.OkWithCode(definition.ERROR_CODE_ERROR_TOP_AGENT_SETTING)
				return
			}
		}

		// lobbyInfo加入自己
		lobbyInfos = append(lobbyInfos, getLobbyInfo(
			agent.Id,
			game.Id,
			gameRoom.Id,
			reqAgentGameRoom.State,
		))

		// lobbyInfo加入子代
		childAgents := global.AgentCache.GetChildAgents(agent.Id)
		if len(childAgents) > 0 {
			for _, childAgent := range global.AgentCache.GetChildAgents(agent.Id) {
				childAgentGameRoom := global.AgentGameRoomCache.Get(childAgent.Id, gameRoom.Id)

				if childAgentGameRoom.State == reqAgentGameRoom.State {
					continue
				}

				lobbyInfos = append(lobbyInfos, getLobbyInfo(
					childAgent.Id,
					game.Id,
					gameRoom.Id,
					reqAgentGameRoom.State,
				))
			}

			sqlSb.WriteString(fmt.Sprintf(sqlTopAgentFmt, agent.Id, agent.LevelCode, gameRoom.Id, reqAgentGameRoom.State))
		} else {
			sqlSb.WriteString(fmt.Sprintf(sqlAgentFmt, agent.Id, gameRoom.Id, reqAgentGameRoom.State))
		}

		actionLogGameRooms = append(actionLogGameRooms, map[string]interface{}{
			"room_type": gameRoom.RoomType,
			"state":     createBancendActionLogDetail(agentGameRoom.State, reqAgentGameRoom.State),
		})
	}

	if len(lobbyInfos) > 0 {
		_, err := c.DB().Exec(sqlSb.String())
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if code, _ := notification.SendSetLobbyInfo(lobbyInfos); code != definition.ERROR_CODE_SUCCESS {
			c.Logger().Info("notification.SendSetLobbyInfo fail, lobbyInfos=%v, resp code=%d", lobbyInfos, code)

			// 通知失敗要重設DB資料
			sqlSb.Reset()
			for _, lobbyInfo := range lobbyInfos {
				agentId := lobbyInfo["AgentId"].(int)
				gameRoomId := lobbyInfo["TableId"].(int)
				agentGameRoom := global.AgentGameRoomCache.Get(agentId, gameRoomId)

				sqlSb.WriteString(fmt.Sprintf(sqlAgentFmt, agentId, gameRoomId, agentGameRoom.State))
			}

			_, err := c.DB().Exec(sqlSb.String())
			if err != nil {
				c.Logger().Warn("notification.SendSetLobbyInfo fail and db rollback fail, req=%v, sql=%s", req, sqlSb.String())
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}

			// 重設DB資料完成才傳通知送失敗
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["agent_name"] = agent.Name
	actionLog["game_id"] = game.Id
	if len(actionLogGameRooms) > 0 {
		actionLog["agent_game_rooms"] = actionLogGameRooms
	}
	c.Set("action_log", actionLog)

	// 全部成功才更新cache
	for _, lobbyInfo := range lobbyInfos {
		agentId := lobbyInfo["AgentId"].(int)
		gameRoomId := lobbyInfo["TableId"].(int)
		state := lobbyInfo["Status"].(int16)

		agentGameRoom := global.AgentGameRoomCache.Get(agentId, gameRoomId)
		agentGameRoom.State = state
	}

	c.Ok(nil)
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 取得代理權限群組權限樣板
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=map[int][]int,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentpermissiontemplateinfo [get]
func (p *SystemAgentApi) GetAgentPermissionTemplateInfo(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	c.Ok(GetTemplatePermissions(int(userClaims.BaseClaims.ID)))
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 取得代理權限群組列表
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentPermissionListRequest true "取得代理權限群組列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentPermissionListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentpermissionlist [post]
func (p *SystemAgentApi) GetAgentPermissionList(c *ginweb.Context) {
	req := &model.GetAgentPermissionListRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userClaims := c.GetUserClaims()

	sqAnd := sq.And{
		sq.Eq{"agent_id": userClaims.BaseClaims.ID},
	}

	checkExist_name := req.Name != ""
	if checkExist_name {
		sqAnd = append(sqAnd, sq.Eq{"name": req.Name})
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", "name", "info", "account_type").
		From("agent_permission").
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

	defer rows.Close()

	resp := make([]*model.GetAgentPermissionListResponse, 0)
	for rows.Next() {
		temp := model.GetAgentPermissionListResponse{}
		if err = rows.Scan(&temp.Id, &temp.AgentId, &temp.Name, &temp.Info,
			&temp.AccountType); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		// Notice: 目前不做任何限制，只要賦予權限就可以取得所有資訊
		// // 自己權限群組不顯示
		// if temp.Id == userClaims.PermissionId {
		// 	continue
		// }

		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 取得代理權限群組
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentPermissionRequest true "取得代理權限群組參數"
// @Success 200 {object} response.Response{data=model.GetAgentPermissionResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentpermission [post]
func (p *SystemAgentApi) GetAgentPermission(c *ginweb.Context) {
	req := &model.GetAgentPermissionRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()

	sqAnd := sq.And{
		sq.Eq{"id": req.Id},
		sq.Eq{"agent_id": userClaims.BaseClaims.ID},
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", "permission", "name", "info", "account_type").
		From("agent_permission").
		Where(sqAnd).
		ToSql()

	resp := &model.GetAgentPermissionResponse{}

	err := c.DB().QueryRow(query, args...).Scan(&resp.Id, &resp.AgentId, &resp.PermissionBytes,
		&resp.Name, &resp.Info, &resp.AccountType)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var permissionSlice table_model.PermissionSlice
	utils.ToStruct(resp.PermissionBytes, &permissionSlice)
	resp.Permissions = permissionSlice.List

	c.Ok(resp)
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 創建代理權限群組
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.CreateAgentPermissionRequest true "創建代理權限群組參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/createagentpermission [post]
func (p *SystemAgentApi) CreateAgentPermission(c *ginweb.Context) {
	req := &model.CreateAgentPermissionRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userClaims := c.GetUserClaims()
	// 層級只能設置自己跟下一層級
	if (req.AccountType != userClaims.AccountType &&
		req.AccountType != userClaims.AccountType+1) ||
		req.AccountType > definition.ACCOUNT_TYPE_NORMAL {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("agent_permission").
		Where(sq.And{
			sq.Eq{"agent_id": userClaims.BaseClaims.ID},
			sq.Eq{"name": req.Name},
		}).
		Suffix(")").
		ToSql()

	var isExists bool
	err := c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 設置名稱不能與其他的權限群組相同
	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ROLE_NAME_EXIST)
		return
	}

	var agentPermission *table_model.AgentPermission
	if agentPermission, err = global.UdfCreateAgentPermission(c.DB(), int(userClaims.BaseClaims.ID), req.Name, req.Info, req.AccountType,
		toPermissionString(req.Permissions), userClaims.BaseClaims.AccountType >= definition.ACCOUNT_TYPE_GENERAL); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// cache新增agentPermission
	global.AgentPermissionCache.Add(agentPermission)

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Name)
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 修改代理權限群組
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.SetAgentPermissionRequest true "修改代理權限群組參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentpermission [post]
func (p *SystemAgentApi) SetAgentPermission(c *ginweb.Context) {
	req := &model.SetAgentPermissionRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userClaims := c.GetUserClaims()

	// 層級只能設置自己跟下一層級
	if (req.AccountType != userClaims.AccountType &&
		req.AccountType != userClaims.AccountType+1) ||
		req.AccountType > definition.ACCOUNT_TYPE_NORMAL {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "permission", "account_type", "name", "info").
		From("agent_permission").
		Where(sq.Eq{"id": req.Id}).
		ToSql()

	var permissionAgentId uint
	var permissionAccountType int
	var permissionBytes []byte
	var permissionName string
	var permissionInfo string
	err := c.DB().QueryRow(query, args...).Scan(&permissionAgentId, &permissionBytes, &permissionAccountType, &permissionName, &permissionInfo)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 不能設置不是自己代理的權限群組
	if permissionAgentId != userClaims.BaseClaims.ID {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	var permission table_model.PermissionSlice
	utils.ToStruct(permissionBytes, &permission)

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("agent_permission").
		Where(sq.And{
			sq.NotEq{"id": req.Id},
			sq.Eq{"agent_id": userClaims.BaseClaims.ID},
			sq.Eq{"name": req.Name},
		}).
		Suffix(")").
		ToSql()

	var isExists bool
	err = c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 設置名稱不能與其他的權限群組相同
	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ROLE_NAME_EXIST)
		return
	}

	// 變更層級要檢查有沒有目前使用者使用中
	if permissionAccountType != req.AccountType {
		query, args, _ = sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("1").
			Prefix("SELECT EXISTS (").
			From("admin_user").
			Where(sq.Eq{"role": req.Id}).
			Suffix(")").
			ToSql()

		err = c.DB().QueryRow(query, args...).Scan(&isExists)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		// 有其他帳號設置在該群組下不能變更層級
		if isExists {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_ROLE_USING)
			return
		}
	}

	// 比對建立操作紀錄
	permissionMap := make(map[int]int)
	for _, permission := range permission.List {
		permissionMap[permission] = 1
	}
	for _, permission := range req.Permissions {
		if _, find := permissionMap[permission]; find {
			permissionMap[permission] = 0
		} else {
			permissionMap[permission] = 2
		}
	}

	beforePermissions := make([]int, 0)
	afterPermissions := make([]int, 0)
	for permission, status := range permissionMap {
		if status == 1 {
			beforePermissions = append(beforePermissions, permission)
		} else if status == 2 {
			afterPermissions = append(afterPermissions, permission)
		}
	}

	isNameChanged := permissionName != req.Name
	isInfoChanged := permissionInfo != req.Info
	isAccountTypeChanged := permissionAccountType != req.AccountType
	isPermissionChanged := len(beforePermissions) > 0 || len(afterPermissions) > 0

	if isNameChanged || isInfoChanged || isAccountTypeChanged || isPermissionChanged {
		// 有更新權限時，修改DB前先踢除使用者
		if !isAccountTypeChanged && isPermissionChanged {
			// 修改自己的權限只要剔除該群組的使用者
			if req.AccountType == userClaims.AccountType {
				for _, adminUser := range global.AdminUserCache.GetAgentAdminUsers(int(userClaims.BaseClaims.ID)) {
					if adminUser.PermissionId != req.Id {
						continue
					}
					c.Jwt.AddBlackTokenByUsername(adminUser.Username)
				}
			} else {
				// 修改下層級的權限要剔除所有下層級的使用者
				for _, agent := range global.AgentCache.GetChildAgents(int(userClaims.BaseClaims.ID)) {
					for _, adminUser := range global.AdminUserCache.GetAgentAdminUsers(agent.Id) {
						c.Jwt.AddBlackTokenByUsername(adminUser.Username)
					}
				}
			}
		}

		var permissionSlice table_model.PermissionSlice
		if permissionSlice, err = global.UdfUpdateAgentPermission(c.DB(), req.Id, int(userClaims.BaseClaims.ID), req.Name, req.Info, req.AccountType,
			toPermissionString(req.Permissions), userClaims.BaseClaims.AccountType > definition.ACCOUNT_TYPE_ADMIN,
			userClaims.BaseClaims.AccountType < req.AccountType); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		// cache更新agentPermission
		agentPermission := global.AgentPermissionCache.Get(req.Id)
		agentPermission.Name = req.Name
		agentPermission.Info = req.Info
		agentPermission.AccountType = req.AccountType
		agentPermission.Permission = permissionSlice

		// cache更新下層級的agentPermission
		if !isAccountTypeChanged && isPermissionChanged && req.AccountType > userClaims.AccountType {
			for _, agent := range global.AgentCache.GetChildAgents(int(userClaims.BaseClaims.ID)) {
				for _, agentPermission := range global.AgentPermissionCache.GetByAgentId(agent.Id) {
					agentPermission.Permission.List = utils.ArrayIntersection(agentPermission.Permission.List, permissionSlice.List)
				}
			}
		}
	}

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(req.Name)
	if isNameChanged {
		actionLog["name"] = createBancendActionLogDetail(permissionName, req.Name)
	}
	if isInfoChanged {
		actionLog["info"] = createBancendActionLogDetail(permissionInfo, req.Info)
	}
	if isAccountTypeChanged {
		actionLog["level"] = createBancendActionLogDetail(permissionAccountType, req.AccountType)
	}
	if isPermissionChanged {
		actionLog["permissions"] = createBancendActionLogDetail(beforePermissions, afterPermissions)
	}
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

// @Tags 代理帳號管理/權限群組管理
// @Summary 刪除代理權限群組
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.DeleteAgentPermissionRequest true "刪除代理權限群組參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/deleteagentpermission [post]
func (p *SystemAgentApi) DeleteAgentPermission(c *ginweb.Context) {
	req := &model.DeleteAgentPermissionRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 不能刪除自己的權限群組
	userClaims := c.GetUserClaims()
	if req.Id == userClaims.PermissionId {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_ROLE_USING)
		return
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "permission", "name").
		From("agent_permission").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var permissionAgentId uint
	var permissionBytes []byte
	var permissionName string
	err = c.DB().QueryRow(query, args...).Scan(&permissionAgentId, &permissionBytes, &permissionName)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 不能刪除不是自己代理的權限群組
	if permissionAgentId != userClaims.BaseClaims.ID {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_ROLE_USING)
		return
	}

	var permission table_model.PermissionSlice
	utils.ToStruct(permissionBytes, &permission)

	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("admin_user").
		Where(sq.And{
			sq.Eq{"role": req.Id},
		}).
		Suffix(")").
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	var isExists bool
	err = c.DB().QueryRow(query, args...).Scan(&isExists)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 有其他帳號設置在該權限群組下不能刪除
	if isExists {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_ROLE_USING)
		return
	}

	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Delete("agent_permission").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	err = c.DB().QueryRow(query, args...).Err()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// cache刪除agentPermission
	global.AgentPermissionCache.Remove(global.AgentPermissionCache.Get((req.Id)))

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(permissionName)
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

// @Tags 系統管理/後台IP白名單
// @Summary 取得代理後台IP資訊列表
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentIpWhitelistListRequest true "取得代理後台IP資訊列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentIpWhitelistListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentipwhitelistlist [post]
func (p *SystemAgentApi) GetAgentIpWhitelistList(c *ginweb.Context) {
	req := &model.GetAgentIpWhitelistListRequest{}
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

	myAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))

	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	if checkExist_agentId {
		reqAgent := global.AgentCache.Get(req.AgentId)
		if reqAgent == nil || !strings.HasPrefix(reqAgent.LevelCode, myAgent.LevelCode) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
	}

	resp := make([]*model.GetAgentIpWhitelistListResponse, 0)
	for _, agent := range global.AgentCache.GetAll() {
		if !strings.HasPrefix(agent.LevelCode, myAgent.LevelCode) {
			continue
		}

		if checkExist_agentId && agent.Id != req.AgentId {
			continue
		}

		resp = append(resp, &model.GetAgentIpWhitelistListResponse{
			Id:                agent.Id,
			Name:              agent.Name,
			LevelCode:         agent.LevelCode,
			IpAddressCount:    len(agent.IPWhitelist),
			ApiIpAddressCount: len(agent.ApiIPWhitelist),
			AdminUsername:     agent.AdminUsername,
		})
	}

	c.Ok(resp)
}

// @Tags 系統管理/後台IP白名單
// @Summary 取得代理後台IP資訊
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentIpWhitelistRequest true "取得代理後台IP資訊參數"
// @Success 200 {object} response.Response{data=[]model.AgentIPWhitelistObj,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentipwhitelist [post]
func (p *SystemAgentApi) GetAgentIpWhitelist(c *ginweb.Context) {
	req := &model.GetAgentIpWhitelistRequest{}
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

	myAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))
	reqAgent := global.AgentCache.Get(req.AgentId)
	if reqAgent == nil || !strings.HasPrefix(reqAgent.LevelCode, myAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	c.Ok(reqAgent.IPWhitelist)
}

// @Tags 系統管理/後台IP白名單
// @Summary 設置代理後台IP資訊
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.SetAgentIpWhitelistRequest true "設置代理後台IP資訊參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentipwhitelist [post]
func (p *SystemAgentApi) SetAgentIpWhitelist(c *ginweb.Context) {
	req := &model.SetAgentIpWhitelistRequest{}
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

	myAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))
	reqAgent := global.AgentCache.Get(req.AgentId)
	if reqAgent == nil || !strings.HasPrefix(reqAgent.LevelCode, myAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	creator := c.GetUsername()

	now := utils.GetUnixTimeNowUTC()
	newIpWhitelist := make([]*table_model.AgentIPWhitelistObj, 0)
	for i := 0; i < len(req.IpWhitelist); i++ {
		// 空白ip跳過不設置
		if req.IpWhitelist[i].IPAddress == "" {
			continue
		}

		if !validIpAddress(req.IpWhitelist[i].IPAddress) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_IP_FORMAT)
			return
		}

		// 空白時間為新增的ip位置
		if req.IpWhitelist[i].CreateTime == 0 {
			req.IpWhitelist[i].CreateTime = now
			req.IpWhitelist[i].Creator = creator
		}

		newIpWhitelist = append(newIpWhitelist, &req.IpWhitelist[i])
	}

	if len(newIpWhitelist) <= 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("agent").
		Set("ip_whitelist", utils.ToJSON(&newIpWhitelist)).
		Where(sq.Eq{"id": req.AgentId}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 修改DB資料
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	actionLog := createBackendActionLogWithTitle(reqAgent.Name)
	// 操作紀錄
	actionLog["ip"] = createBancendActionLogDetail(reqAgent.IPWhitelist, req.IpWhitelist)
	c.Set("action_log", actionLog)

	// 修改cache資料
	reqAgent.IPWhitelist = newIpWhitelist
	global.AgentCache.Add(reqAgent)

	c.Ok("")
}

// @Tags 系統管理/後台IP白名單
// @Summary 取得代理後台IP資訊
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentIpWhitelistRequest true "取得代理API IP資訊"
// @Success 200 {object} response.Response{data=[]model.AgentIPWhitelistObj,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentapiipwhitelist [post]
func (p *SystemAgentApi) GetAgentApiIpWhitelist(c *ginweb.Context) {
	req := &model.GetAgentIpWhitelistRequest{}
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

	myAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))
	reqAgent := global.AgentCache.Get(req.AgentId)
	if reqAgent == nil || !strings.HasPrefix(reqAgent.LevelCode, myAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	c.Ok(reqAgent.ApiIPWhitelist)
}

// @Tags 系統管理/後台IP白名單
// @Summary 設置代理後台IP資訊
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.SetAgentIpWhitelistRequest true "設置代理API IP資訊"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentapiipwhitelist [post]
func (p *SystemAgentApi) SetAgentApiIpWhitelist(c *ginweb.Context) {
	req := &model.SetAgentIpWhitelistRequest{}
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

	myAgent := global.AgentCache.Get(int(claims.BaseClaims.ID))
	reqAgent := global.AgentCache.Get(req.AgentId)
	if reqAgent == nil || !strings.HasPrefix(reqAgent.LevelCode, myAgent.LevelCode) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	creator := c.GetUsername()

	now := utils.GetUnixTimeNowUTC()
	newIpWhitelist := make([]*table_model.AgentIPWhitelistObj, 0)
	for i := 0; i < len(req.IpWhitelist); i++ {
		// 空白ip跳過不設置
		if req.IpWhitelist[i].IPAddress == "" {
			continue
		}

		if !validIpAddress(req.IpWhitelist[i].IPAddress) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_IP_FORMAT)
			return
		}

		// 空白時間為新增的ip位置
		if req.IpWhitelist[i].CreateTime == 0 {
			req.IpWhitelist[i].CreateTime = now
			req.IpWhitelist[i].Creator = creator
		}

		newIpWhitelist = append(newIpWhitelist, &req.IpWhitelist[i])
	}

	if len(newIpWhitelist) <= 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("agent").
		Set("api_ip_whitelist", utils.ToJSON(&newIpWhitelist)).
		Where(sq.Eq{"id": req.AgentId}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 修改DB資料
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	actionLog := createBackendActionLogWithTitle(reqAgent.Name)
	// 操作紀錄
	actionLog["api_ip"] = createBancendActionLogDetail(reqAgent.ApiIPWhitelist, req.IpWhitelist)
	c.Set("action_log", actionLog)

	// 修改cache資料
	reqAgent.ApiIPWhitelist = newIpWhitelist
	global.AgentCache.Add(reqAgent)

	c.Ok("")
}

// @Tags 運營管理/後台代理上下分
// @Summary 取得代理錢包餘額列表
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentWalletListRequest true "取得代理錢包餘額列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentWalletListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/getagentwalletlist [post]
func (p *SystemAgentApi) GetAgentWalletList(c *ginweb.Context) {
	req := &model.GetAgentWalletListRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))
	if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN &&
		userAgent.Cooperation != definition.AGENT_COOPERATION_BUY_POINT {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 必要條件檢查
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)

	sqOr := sq.Or{}
	if checkExist_agentId {
		reqAgent := global.AgentCache.Get(req.AgentId)
		if reqAgent == nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		} else if reqAgent.TopAgentId != userAgent.Id {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		sqOr = append(sqOr, sq.Eq{"aw.agent_id": []int{userAgent.Id, reqAgent.Id}})
	} else {
		sqOr = append(sqOr, sq.Eq{"aw.agent_id": userAgent.Id})
		sqOr = append(sqOr, sq.Eq{"a.top_agent_id": userAgent.Id})
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"a.cooperation": definition.AGENT_COOPERATION_BUY_POINT})
	sqAnd = append(sqAnd, sqOr)

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("au.username", "aw.agent_id", "a.name", "a.level_code", "a.commission", "aw.amount", "a.is_enabled", "a.create_time").
		From("agent_wallet AS aw").
		InnerJoin("agent AS a ON aw.agent_id = a.id").
		InnerJoin("admin_user AS au ON aw.agent_id = au.agent_id AND au.is_added = false").
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

	defer rows.Close()

	resp := make([]*model.GetAgentWalletListResponse, 0)
	for rows.Next() {
		var temp model.GetAgentWalletListResponse
		if err := rows.Scan(&temp.AdminUserUsername, &temp.Id, &temp.Name, &temp.LevelCode,
			&temp.Commission, &temp.Balance, &temp.IsEnabled, &temp.CreateTime); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 運營管理/後台代理上下分
// @Summary 設置代理錢包餘額
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetAgentWalletRequest true "設置代理錢包餘額參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/agent/setagentwallet [post]
func (p *SystemAgentApi) SetAgentWallet(c *ginweb.Context) {
	req := &model.SetAgentWalletRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	reqAgent := global.AgentCache.Get(req.AgentId)
	if reqAgent == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	userClaims := c.GetUserClaims()
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))
	// 開發商(營運商)及買分的代理才可以使用此功能
	if userClaims.AccountType != definition.ACCOUNT_TYPE_ADMIN &&
		userAgent.Cooperation != definition.AGENT_COOPERATION_BUY_POINT {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 只能修改自己下一級的代理
	if reqAgent.TopAgentId != userAgent.Id {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	var fromAgentId, toAgentId, kind int
	if req.ChangeScore > 0 {
		fromAgentId = userAgent.Id
		toAgentId = reqAgent.Id
		kind = definition.WALLET_LEDGER_KIND_BACKEND_UP
	} else {
		fromAgentId = reqAgent.Id
		toAgentId = userAgent.Id
		kind = definition.WALLET_LEDGER_KIND_BACKEND_DOWN
	}

	creator := c.GetUsername()
	saltFormat := "%d%d%d%v%s_%d"

	var query string
	var resultJson string
	var err error
	if userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN {
		salt := fmt.Sprintf(saltFormat, definition.ORDER_TYPE_AGENT_WALLET_LEDGER, req.AgentId, kind, req.ChangeScore, creator, time.Now().UnixMilli()%int64(time.Millisecond))
		orderId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_AGENT_WALLET_LEDGER, salt, time.Now())

		query = `SELECT "public"."udf_backend_update_agent_wallect" ($1, $2, $3, $4, $5, $6)`
		err = c.DB().QueryRow(query, orderId, reqAgent.Id, req.ChangeScore, req.Info, kind, creator).Scan(&resultJson)
	} else {
		fromSalt := fmt.Sprintf(saltFormat, definition.ORDER_TYPE_AGENT_WALLET_LEDGER, fromAgentId, definition.WALLET_LEDGER_KIND_BACKEND_DOWN, req.ChangeScore, creator, time.Now().UnixMilli()%int64(time.Millisecond))
		fromOrderId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_AGENT_WALLET_LEDGER, fromSalt, time.Now())

		toSalt := fmt.Sprintf(saltFormat, definition.ORDER_TYPE_AGENT_WALLET_LEDGER, toAgentId, definition.WALLET_LEDGER_KIND_BACKEND_UP, -req.ChangeScore, creator, time.Now().UnixMilli()%int64(time.Millisecond))
		toOrderId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_AGENT_WALLET_LEDGER, toSalt, time.Now())

		if req.ChangeScore < 0 {
			req.ChangeScore = -req.ChangeScore
		}

		query := `SELECT "public"."udf_backend_update_agent_wallects" ($1, $2, $3, $4, $5, $6, $7)`
		err = c.DB().QueryRow(query, fromOrderId, fromAgentId, toOrderId, toAgentId, req.ChangeScore, req.Info, creator).Scan(&resultJson)
	}

	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	result := utils.ToMap([]byte(resultJson))
	code := int(result["code"].(float64))

	errorCode := definition.ERROR_CODE_SUCCESS
	switch code {
	case 1:
		errorCode = definition.ERROR_CODE_ERROR_REQUEST_DATA
	case 2:
		errorCode = definition.ERROR_CODE_ERROR_AGENT_WALLET_AMOUNT_NOT_ENOUGH
	}

	// 餘額不足才會失敗
	if errorCode != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(errorCode)
		return
	}

	c.Ok("")
}
