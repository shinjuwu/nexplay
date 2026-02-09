package main

import (
	"bytes"
	"definition"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		createFrontendConstant(os.Args[1])
	}
	if len(os.Args) > 2 && os.Args[2] != "" {
		createSystemFeatureCodeConstant(os.Args[2])
	}
}

func createFrontendConstant(path string) {
	// fs.FileMode check https://pkg.go.dev/io/fs#FileMode
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, fs.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	var buffer bytes.Buffer

	buffer.WriteString("export default {\n")

	createConstant(&buffer, "TableDefaultLength", definition.TABLE_DEFAULT_LENGTH)
	createConstant(&buffer, "TableDefaultLengthMenu", definition.TABLE_DEFAULT_LENGTH_MENU)

	createEnum(&buffer, "AccountStatus", map[string]int{
		"Disable": definition.ACCOUNT_STATUS_DISABLE,
		"Open":    definition.ACCOUNT_STATUS_OPEN,
	})
	createEnum(&buffer, "AccountType", map[string]int{
		"Admin":   definition.ACCOUNT_TYPE_ADMIN,
		"General": definition.ACCOUNT_TYPE_GENERAL,
		"Nornam":  definition.ACCOUNT_TYPE_NORMAL,
	})
	createEnum(&buffer, "Agent", map[string]int{
		"All": definition.AGENT_ID_ALL,
	})
	createEnum(&buffer, "AgentCooperation", map[string]int{
		"BuyPoint": definition.AGENT_COOPERATION_BUY_POINT,
		"Trust":    definition.AGENT_COOPERATION_TRUST,
	})
	createEnum(&buffer, "ErrorCode", map[string]int{
		"Success":                         definition.ERROR_CODE_SUCCESS,
		"Fail":                            definition.ERROR_CODE_FAIL,
		"ErrorRequestData":                definition.ERROR_CODE_ERROR_REQUEST_DATA,
		"ErrorDatabase":                   definition.ERROR_CODE_ERROR_DATABASE,
		"ErrorNotification":               definition.ERROR_CODE_ERROR_NOTIFICATION,
		"ErrorJWT":                        definition.ERROR_CODE_ERROR_JWT,
		"ErrorPermission":                 definition.ERROR_CODE_ERROR_PERMISSION,
		"ErrorCaptcha":                    definition.ERROR_CODE_ERROR_CAPTCHA,
		"ErrorAccount":                    definition.ERROR_CODE_ERROR_ACCOUNT,
		"ErrorAccountDisabled":            definition.ERROR_CODE_ERROR_ACCOUNT_DISABLED,
		"ErrorAccountExist":               definition.ERROR_CODE_ERROR_ACCOUNT_EXIST,
		"ErrorLocalCacheNotExist":         definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST,
		"ErrorTimeRange":                  definition.ERROR_CODE_ERROR_TIME_RANGE,
		"ErrorTimeBeforeDays":             definition.ERROR_CODE_ERROR_TIME_BEFORE_DAYS,
		"ErrorTopAgentSetting":            definition.ERROR_CODE_ERROR_TOP_AGENT_SETTING,
		"ErrorRedis":                      definition.ERROR_CODE_ERROR_REDIS,
		"ErrorNoDataBeAffected":           definition.ERROR_CODE_ERROR_NO_DATA_BE_AFFECTED,
		"ErrorAccountRoleUsing":           definition.ERROR_CODE_ERROR_ACCOUNT_ROLE_USING,
		"ErrorRoleNameExist":              definition.ERROR_CODE_ERROR_ROLE_NAME_EXIST,
		"ErrorPermissionRoleNotExist":     definition.ERROR_CODE_ERROR_PERMISSION_ROLE_NOT_EXIST,
		"ErrorAgentIpWhitelistCount":      definition.ERROR_CODE_ERROR_AGENT_IP_WHITELIST_COUNT,
		"ErrorAgentNameExist":             definition.ERROR_CODE_ERROR_AGENT_NAME_EXIST,
		"ErrorAgentWalletAmountNotEnough": definition.ERROR_CODE_ERROR_AGENT_WALLET_AMOUNT_NOT_ENOUGH,
		"ErrorDatabaseNoRows":             definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS,
		"ErrorFeatureDisabled":            definition.ERROR_CODE_ERROR_FEATURE_DISABLED,
		"ErrorChannelId":                  definition.ERROR_CODE_ERROR_CHANNEL_ID,
		"ErrorChatSendMessageFailed":      definition.ERROR_CODE_ERROR_CHAT_SEND_MESSAGE_FAILED,
		"ErrorApiServerRequestFailed":     definition.ERROR_CODE_ERROR_API_SERVER_REQUEST_FAILED,
		"ErrorAccountNotExist":            definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST,
		"ErrorQueryCombition":             definition.ERROR_CODE_ERROR_QUERY_COMBITION,
		"ErrorDatabaseCommit":             definition.ERROR_CODE_ERROR_DATABASE_COMMIT,
		"ErrorDatabaseExec":               definition.ERROR_CODE_ERROR_DATABASE_EXEC,
		"ErrorGameServerNotInMaintenance": definition.ERROR_CODE_ERROR_GAME_SERVER_NOT_IN_MAINTENANCE,
		"ErrorIpFormat":                   definition.ERROR_CODE_ERROR_IP_FORMAT,
		"ErrorGameUsersNotExist":          definition.ERROR_CODE_ERROR_GAME_USERS_NOT_EXIST,
		"ErrorGameUsersBeBlock":           definition.ERROR_CODE_ERROR_GAME_USERS_BE_BLOCK,
		"ErrorAnalyzeFailed":              definition.ERROR_CODE_ERROR_ANALYZE_FAILED,
		"ErrorCaptchaExoired":             definition.ERROR_CODE_ERROR_CAPTCHA_EXPIRED,
		"ErrorPasswordFormat":             definition.ERROR_CODE_ERROR_PASSWORD_FORMAT,
		"ErrorPassword":                   definition.ERROR_CODE_ERROR_PASSWORD,
		"ErrorConfirmPassword":            definition.ERROR_CODE_ERROR_CONFIRM_PASSWORD,
		"ErrorPasswordSame":               definition.ERROR_CODE_ERROR_PASSWORD_SAME,
		"ErrorMarqueeOverdue":             definition.ERROR_CODE_ERROR_MARQUEE_OVERDUE,
		"ErrorGameServerInMaintenance":    definition.ERROR_CODE_ERROR_GAME_SERVER_IN_MAINTENANCE,
		"ErrorAdminUserTypeIllegal":       definition.ERROR_CODE_ERROR_ADMIN_USER_TYPE_ILLEGAL,
		"ErrorAnnouncementNotExist":       definition.ERROR_CODE_ERROR_ANNOUNCEMENT_NOT_EXIST,
		"ErrorAgentIpWhitelist":           definition.ERROR_CODE_ERROR_AGENT_IP_WHITELIST,
		"ErrorCurrencyNotSupported":       definition.ERROR_CODE_ERROR_CURRENCY_NOT_SUPPORTED,
		"ErrorKillRatioParentOpen":        definition.ERROR_CODE_ERROR_KILLRATION_PARENT_OPEN,
		"ErrorWalletUrlParseFailed":       definition.ERROR_CODE_ERROR_WALLET_URL_PARSE_FAILED,
		"ErrorNoChangeInData":             definition.ERROR_CODE_ERROR_NO_CHANGE_IN_DATA,
		"ErrorIncomeRatioData":            definition.ERROR_CODE_ERROR_INCOME_RATIO_DATA,
		"ErrorStringLimit":                definition.ERROR_CODE_ERROR_STRING_LIMIT,
		"ErrorLocal":                      definition.ERROR_CODE_ERROR_LOCAL,
	})

	createEnum(&buffer, "Game", map[string]int{
		"Lobby":           definition.GAME_ID_LOBBY,
		"Friendsroom":     definition.GAME_ID_FRIENDSROOM,
		"All":             definition.GAME_ID_ALL,
		"Baccarat":        definition.GAME_ID_BACCARAT,
		"Fantan":          definition.GAME_ID_FANTAN,
		"Colordisc":       definition.GAME_ID_COLORDISC,
		"Prawncrab":       definition.GAME_ID_PRAWNCRAB,
		"Hundredsicbo":    definition.GAME_ID_HUNDREDSICBO,
		"Cockfight":       definition.GAME_ID_COCKFIGHT,
		"Dogracing":       definition.GAME_ID_DOGRACING,
		"Rocket":          definition.GAME_ID_ROCKET,
		"Andarbahar":      definition.GAME_ID_ANDARBAHAR,
		"Roulette":        definition.GAME_ID_ROULETTE,
		"Blackjack":       definition.GAME_ID_BLACKJACK,
		"Sangong":         definition.GAME_ID_SANGONG,
		"Bullbull":        definition.GAME_ID_BULLBULL,
		"Texas":           definition.GAME_ID_TEXAS,
		"Rummy":           definition.GAME_ID_RUMMY,
		"Goldenflower":    definition.GAME_ID_GOLDENFLOWER,
		"Pokdeng":         definition.GAME_ID_POKDENG,
		"Catte":           definition.GAME_ID_CATTE,
		"Chinesepoker":    definition.GAME_ID_CHINESEPOKER,
		"Okey":            definition.GAME_ID_OKEY,
		"Teenpatti":       definition.GAME_ID_TEENPATTI,
		"Fruitslot":       definition.GAME_ID_FRUITSLOT,
		"Rcfishing":       definition.GAME_ID_RCFISHING,
		"Plinko":          definition.GAME_ID_PLINKO,
		"Happyfishing":    definition.GAME_ID_HAPPYFISHING,
		"Fruit777slot":    definition.GAME_ID_FRUIT777SLOT,
		"Megsharkslot":    definition.GAME_ID_MEGSHARKSLOT,
		"Midasslot":       definition.GAME_ID_MIDASSLOT,
		"Wildgemslot":     definition.GAME_ID_WILDGEMSLOT,
		"Jumphighslot":    definition.GAME_ID_JUMPHIGHSLOT,
		"Pyrtreasureslot": definition.GAME_ID_PYRTREASURESLOT,
		"Friendstexas":    definition.GAME_ID_FRIENDSTEXAS,
	})
	createEnum(&buffer, "GameType", map[string]int{
		"All":            definition.GAME_TYPE_ALL,
		"Lobby":          definition.GAME_TYPE_LOBBY,
		"BaiRen":         definition.GAME_TYPE_BAIREN,
		"ChiPai":         definition.GAME_TYPE_CHIPAI,
		"ElectronicGame": definition.GAME_TYPE_ELECTRONIC_GAME,
		"Slot":           definition.GAME_TYPE_SLOT,
		"FriendsRoom":    definition.GAME_TYPE_FRIENDSROOM,
	})
	createEnum(&buffer, "GameState", map[string]int{
		"All":      definition.GAME_STATE_ALL,
		"Online":   definition.GAME_STATE_ONLINE,
		"Offline":  definition.GAME_STATE_OFFLINE,
		"Maintain": definition.GAME_STATE_MAINTAIN,
	})
	createEnum(&buffer, "RoomType", map[string]int{
		"All":            definition.ROOM_TYPE_ALL,
		"Newbie":         definition.ROOM_TYPE_NEWBIE,
		"Common":         definition.ROOM_TYPE_COMMON,
		"Master":         definition.ROOM_TYPE_MASTER,
		"GrandMaster":    definition.ROOM_TYPE_GRAND_MASTER,
		"Beginner":       definition.ROOM_TYPE_BEGINNER,
		"SmallSuccess":   definition.ROOM_TYPE_SMALL_SUCCESS,
		"GreateFacility": definition.ROOM_TYPE_GREATE_FACILITY,
		"Perfection":     definition.ROOM_TYPE_PERFECTION,
	})
	createEnum(&buffer, "TableSortDirection", map[string]int{
		"Asc":  definition.TABLE_SORT_DIRECTION_ASC,
		"Desc": definition.TABLE_SORT_DIRECTION_DESC,
	})
	createEnum(&buffer, "WalletLedgerKind", map[string]int{
		"All":              definition.WALLET_LEDGER_KIND_ALL,
		"ApiUp":            definition.WALLET_LEDGER_KIND_API_UP,
		"ApiDown":          definition.WALLET_LEDGER_KIND_API_DOWN,
		"BackendUp":        definition.WALLET_LEDGER_KIND_BACKEND_UP,
		"BackendDown":      definition.WALLET_LEDGER_KIND_BACKEND_DOWN,
		"SingleWalletUp":   definition.WALLET_LEDGER_KIND_SINGLE_WALLET_UP,
		"SingleWalletDown": definition.WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN,
	})

	createEnum(&buffer, "MarqueeType", map[string]int{
		"All":    definition.MARQUEE_TYPE_NONE,
		"System": definition.MARQUEE_TYPE_SYSTEM,
		"Event":  definition.MARQUEE_TYPE_EVENT,
	})

	createConstant(&buffer, "LangType", map[string]string{
		"All": definition.LANG_TYPE_NONE,
		"CHS": definition.LANG_TYPE_CHS,
		"CHT": definition.LANG_TYPE_CHT,
		"EN":  definition.LANG_TYPE_EN,
		"VI":  definition.LANG_TYPE_VI,
		"TH":  definition.LANG_TYPE_TH,
		"PT":  definition.LANG_TYPE_PT,
		"TR":  definition.LANG_TYPE_TR,
	})

	createEnum(&buffer, "ActionLogType", map[string]int{
		"All":    definition.ACTION_LOG_TYPE_ALL,
		"Create": definition.ACTION_LOG_TYPE_CREATE,
		"Update": definition.ACTION_LOG_TYPE_UPDATE,
		"Delete": definition.ACTION_LOG_TYPE_DELETE,
	})

	createEnum(&buffer, "AnnounceType", map[string]int{
		"All":    definition.ANNOUNCEMENT_TYPE_NONE,
		"System": definition.ANNOUNCEMENT_TYPE_SYSTEM,
		"Event":  definition.ANNOUNCEMENT_TYPE_EVENT,
	})

	createConstant(&buffer, "CurrencyType", map[string]string{
		"CNY": definition.CUR_CNY,
		"VND": definition.CUR_VND,
		"THB": definition.CUR_THB,
		"MYR": definition.CUR_MYR,
		"PHP": definition.CUR_PHP,
		"INR": definition.CUR_INR,
		"BRL": definition.CUR_BRL,
		"TRY": definition.CUR_TRY,
	})

	createConstant(&buffer, "AutoRiskCode", map[string]int{
		"ApiRequestLimit":        definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_API_REQUEST,
		"CoinInAndOutRatioLimit": definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_COIN_IN_AND_COIN_OUT_DIFF,
		"WinRateLimit":           definition.AUTO_RISK_CONTROL_RISK_CODE_GAME_USER_WIN_RATE,
	})

	createEnum(&buffer, "KillDive", map[string]int{
		"Normal":     definition.GAMEUSERS_STATUS_KILLDIVE_NORMAL,
		"ConfigKill": definition.GAMEUSERS_STATUS_KILLDIVE_CONFIGKILL,
		"BlackKill":  definition.GAMEUSERS_STATUS_KILLDIVE_BLACKKILL,
	})

	createEnum(&buffer, "AgentWallet", map[string]int{
		"Transfer": definition.AGENT_WALLET_TRANSFER,
		"Single":   definition.AGENT_WALLET_SINGLE,
	})

	createEnum(&buffer, "AgentLobbySwitch", map[string]int{
		"Normal":      definition.LOBBY_NORMAL_SWITCH,
		"FriendsRoom": definition.LOBBY_FRIENDSROOM_SWITCH,
	})

	createConstant(&buffer, "Device", map[string]string{
		"All":     definition.DEVICE_ID_ALL,
		"PC":      definition.DEVICE_ID_PC,
		"Mac":     definition.DEVICE_ID_Mac,
		"IOS":     definition.DEVICE_ID_IOS,
		"Android": definition.DEVICE_ID_ANDROID,
		"Other":   definition.DEVICE_ID_OTHER,
	})
	createConstant(&buffer, "Browser", map[string]string{
		"All": definition.BROWSER_ID_ALL,
	})
	createConstant(&buffer, "Country", map[string]string{
		"All": definition.COUNTRY_ID_ALL,
	})

	createEnum(&buffer, "CannedLanguageType", map[string]int{
		"Default": definition.CannedType_Default,
		"Custome": definition.CannedType_Custome,
	})

	createEnum(&buffer, "CannedStatus", map[string]int{
		"Close": definition.CannedStatus_Close,
		"Open":  definition.CannedStatus_Opne,
	})

	createEnum(&buffer, "CannedEmojiType", map[string]int{
		"Happy":      definition.CannedEmojiType_hp,
		"Excited":    definition.CannedEmojiType_ex,
		"Calm":       definition.CannedEmojiType_cl,
		"Depression": definition.CannedEmojiType_dp,
		"Agitated":   definition.CannedEmojiType_ag,
	})

	createConstant(&buffer, "CannedLangTypeSetting", definition.CANNED_LANG_TYPE_MAP)

	buffer.WriteString("}")

	err := os.WriteFile(path, buffer.Bytes(), 0)
	if err != nil {
		panic(err)
	}
}

func createSystemFeatureCodeConstant(path string) {
	// fs.FileMode check https://pkg.go.dev/io/fs#FileMode
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, fs.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	var buffer bytes.Buffer

	buffer.WriteString("export default {\n")

	createConstant(&buffer, "AgentCreateAgent", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_CREATE_AGENT)
	createConstant(&buffer, "AgentGetAgentList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_LIST)
	createConstant(&buffer, "AgentGetAgentSecretKey", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_SECRET_KEY)
	createConstant(&buffer, "AgentGetAgentCoinSupplyInfo", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_COIN_SUPPLY_INFO)
	createConstant(&buffer, "AgentSetAgentCoinSupplyInfo", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_COIN_SUPPLY_INFO)
	createConstant(&buffer, "AgentGetAgentGameList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_GAME_LIST)
	createConstant(&buffer, "AgentSetAgentGameState", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_GAME_STATE)
	createConstant(&buffer, "AgentGetAgentGameRoomList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_GAME_ROOM_LIST)
	createConstant(&buffer, "AgentSetAgentGameRoomState", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_GAME_ROOM_STATE)
	createConstant(&buffer, "AgentGetAgentPermissionTemplateInfo", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION_TEMPLATE_INFO)
	createConstant(&buffer, "AgentGetAgentPermissionList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION_LIST)
	createConstant(&buffer, "AgentCreateAgentPermission", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_CREATE_AGENT_PERMISSION)
	createConstant(&buffer, "AgentSetAgentPermission", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_PERMISSION)
	createConstant(&buffer, "AgentDeleteAgentPermission", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_DELETE_AGENT_PERMISSION)
	createConstant(&buffer, "GameGetGameList", definition.PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_LIST)
	createConstant(&buffer, "GameSetGameState", definition.PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_STATE)
	createConstant(&buffer, "GlobalGetAgentList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_LIST)
	createConstant(&buffer, "GlobalGetAllGameList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_ALL_GAME_LIST)
	createConstant(&buffer, "GlobalGetGameList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_GAME_LIST)
	createConstant(&buffer, "GlobalGetRoomTypeList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_ROOM_TYPE_LIST)
	createConstant(&buffer, "GlobalGetAgentPermissionList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_PERMISSION_LIST)
	createConstant(&buffer, "ManageGetMarqueeList", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MARQUEE_LIST)
	createConstant(&buffer, "ManageGetMarquee", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MARQUEE)
	createConstant(&buffer, "ManageCreateMarquee", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_CREATE_MARQUEE)
	createConstant(&buffer, "ManageUpdateMarquee", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_UPDATE_MARQUEE)
	createConstant(&buffer, "ManageDeleteMarquee", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_DELETE_MARQUEE)
	createConstant(&buffer, "RecordGetUserPlayLogList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_USER_PLAY_LOG_LIST)
	createConstant(&buffer, "RecordGetPlayLogCommon", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_PLAY_LOG_COMMON)
	createConstant(&buffer, "RecordGetWalletLedgerList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_WALLET_LEDGER_LIST)
	createConstant(&buffer, "UserPing", definition.PERMISSION_LIST_FEATURE_CODE_USER_PING)
	createConstant(&buffer, "UserCreateAdminUser", definition.PERMISSION_LIST_FEATURE_CODE_USER_CREATE_ADMIN_USER)
	createConstant(&buffer, "UserGetAdminUsers", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_ADMIN_USERS)
	createConstant(&buffer, "UserGetAdminUserInfo", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_ADMIN_USER_INFO)
	createConstant(&buffer, "UserUpdateAdminUserInfo", definition.PERMISSION_LIST_FEATURE_CODE_USER_UPDATE_ADMIN_USER_INFO)
	createConstant(&buffer, "UserGetGameUsers", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USERS)
	createConstant(&buffer, "UserGetGameUserInfo", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_INFO)
	createConstant(&buffer, "UserUpdateGameUserInfo", definition.PERMISSION_LIST_FEATURE_CODE_USER_UPDATE_GAME_USER_INFO)
	createConstant(&buffer, "CalCalPerformanceReport", definition.PERMISSION_LIST_FEATURE_CODE_CAL_CAL_PERFORMANCE_REPORT)
	createConstant(&buffer, "CalGetJobShedulerList", definition.PERMISSION_LIST_FEATURE_CODE_CAL_GET_JOB_SHEDULER_LIST)
	createConstant(&buffer, "CalGetPerformanceReport", definition.PERMISSION_LIST_FEATURE_CODE_CAL_GET_PERFORMANCE_REPORT)
	createConstant(&buffer, "AgentGetAgentIpWhitelistList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_IP_WHITELIST_LIST)
	createConstant(&buffer, "AgentGetAgentIpWhitelist", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_IP_WHITELIST)
	createConstant(&buffer, "AgentSetAgentIpWhitelist", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_IP_WHITELIST)
	createConstant(&buffer, "AgentGetAgentPermission", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION)
	createConstant(&buffer, "AgentGetAgentWalletList", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_WALLET_LIST)
	createConstant(&buffer, "AgentSetAgentWallet", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_WALLET)
	createConstant(&buffer, "RecordGetAgentWalletLedgerList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_AGENT_WALLET_LEDGER_LIST)
	createConstant(&buffer, "UserGetGameUserWalletList", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_WALLET_LIST)
	createConstant(&buffer, "UserSetGameUserWallet", definition.PERMISSION_LIST_FEATURE_CODE_USER_SET_GAME_USER_WALLET)
	createConstant(&buffer, "RecordConfirmWalletLedger", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_CONFIRM_WALLET_LEDGER)
	createConstant(&buffer, "ManageGetStatData", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_STAT_DATA)
	createConstant(&buffer, "ManageGetRiskUserList", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_RISK_USER_LIST)
	createConstant(&buffer, "ManageGetGameLeaderboards", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_GAME_LEADERBOARDS)
	createConstant(&buffer, "NotifyGetChatServiceConnInfo", definition.PERMISSION_LIST_FEATURE_CODE_NOTIFY_GET_CHAT_SERVICE_CONN_INFO)
	createConstant(&buffer, "RiskControlGetIncomeRatioList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_INCOME_RATIO_LIST)
	createConstant(&buffer, "RiskControlGetIncomeRatio", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_INCOME_RATIO)
	createConstant(&buffer, "RiskControlSetIncomeRatio", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_INCOME_RATIO)
	createConstant(&buffer, "ManageGetIntervalDaysBettorData", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_DAYS_BETTOR_DATA)
	createConstant(&buffer, "ManageGetIntervalTotalScoreData", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_TOTAL_SCORE_DATA)
	createConstant(&buffer, "ManageGetIntervalTotalBettorInfoData", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_TOTAL_BETTOR_INFO_DATA)
	createConstant(&buffer, "RecordGetBackendActionLog", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_BACKEND_ACTION_LOG)
	createConstant(&buffer, "GameGetGameServerState", definition.PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_SERVER_STATE)
	createConstant(&buffer, "GameSetGameServerState", definition.PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_SERVER_STATE)
	createConstant(&buffer, "GameNotifyGameServer", definition.PERMISSION_LIST_FEATURE_CODE_GAME_NOTIFY_GAME_SERVER)
	createConstant(&buffer, "ManageGetAnnouncementList", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_ANNOUNCEMENT_LIST)
	createConstant(&buffer, "ManageGetAnnouncement", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_ANNOUNCEMENT)
	createConstant(&buffer, "ManageCreateAnnouncement", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_CREATE_ANNOUNCEMENT)
	createConstant(&buffer, "ManageUpdateAnnouncement", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_UPDATE_ANNOUNCEMENT)
	createConstant(&buffer, "ManageDeleteAnnouncement", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_DELETE_ANNOUNCEMENT)
	createConstant(&buffer, "CalGetPerformanceReportList", definition.PERMISSION_LIST_FEATURE_CODE_CAL_GET_PERFORMANCE_REPORT_LIST)
	createConstant(&buffer, "UserSetPersonalInfo", definition.PERMISSION_LIST_FEATURE_CODE_USER_SET_PERSONAL_INFO)
	createConstant(&buffer, "UserSetPersonalPassword", definition.PERMISSION_LIST_FEATURE_CODE_USER_SET_PERSONAL_PASSWORD)
	createConstant(&buffer, "RiskControlGetAgentIncomeRatioList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO_LIST)
	createConstant(&buffer, "RiskControlGetAgentIncomeRatio", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO)
	createConstant(&buffer, "RiskControlSetAgentIncomeRatio", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AGENT_INCOME_RATIO)
	createConstant(&buffer, "RiskControlGetAgentCustomTagList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_CUSTOM_TAG_SETTING_LIST)
	createConstant(&buffer, "RiskControlSetAgentCustomTagList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AGENT_CUSTOM_TAG_SETTING_LIST)
	createConstant(&buffer, "RiskControlGetGameUsersCustomTagList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USERS_CUSTOM_TAG_LIST)
	createConstant(&buffer, "RiskControlGetGameUsersCustomTag", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USERS_CUSTOM_TAG)
	createConstant(&buffer, "RiskControlSetGameUsersCustomTag", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_USERS_CUSTOM_TAG)
	createConstant(&buffer, "RiskControlGetAgentIncomeRatioAndGameData", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO_AND_GAME_DATA)
	createConstant(&buffer, "GlobalGetAgentCustomTagSettingList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_CUSTOM_TAG_SETTING_LIST)
	createConstant(&buffer, "RecordGetBackendLoginLogList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_BACKEND_LOGIN_LOG_LIST)
	createConstant(&buffer, "GameGetGameIconList", definition.PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_ICON_LIST)
	createConstant(&buffer, "GameSetGameIconList", definition.PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_ICON_LIST)
	createConstant(&buffer, "AgentGetAgentApiIpWhitelist", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_API_IP_WHITELIST)
	createConstant(&buffer, "AgentSetAgentApiIpWhitelist", definition.PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_API_IP_WHITELIST)
	createConstant(&buffer, "SystemGetExchangeDataList", definition.PERMISSION_LIST_FEATURE_CODE_SYSTEM_GET_EXCHANGE_DATA_LIST)
	createConstant(&buffer, "SystemSetExchangeDataList", definition.PERMISSION_LIST_FEATURE_CODE_SYSTEM_SET_EXCHANGE_DATA_LIST)
	createConstant(&buffer, "RiskControlGetAutoRiskControlSetting", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AUTO_RISKCONTROL_SETTING)
	createConstant(&buffer, "RiskControlSetAutoRiskControlSetting", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AUTO_RISKCONTROL_SETTING)
	createConstant(&buffer, "RiskControlGetGameUserRiskControlTag", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USER_RISKCONTROL_TAG)
	createConstant(&buffer, "RiskControlSetGameUserRiskControlTag", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_USER_RISKCONTROL_TAG)
	createConstant(&buffer, "RiskControlGetAutoRiskControlLogList", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AUTO_RISKCONTROL_LOG_LIST)
	createConstant(&buffer, "RecordGetAgentGameRatioStatList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_AGENT_GAME_RATIO_STAT_LIST)
	createConstant(&buffer, "RecordGetGameUsersStatHourList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_GAME_USERS_STAT_HOUR_LIST)
	createConstant(&buffer, "ManageGetMaintainPageSetting", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MAINTAIN_PAGE_SETTING)
	createConstant(&buffer, "ManageSetMaintainPageSetting", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_SET_MAINTAIN_PAGE_SETTING)
	createConstant(&buffer, "UserResetPassword", definition.PERMISSION_LIST_FEATURE_CODE_USER_RESET_PASSWORD)
	createConstant(&buffer, "ManageGetIntervalRealTimeUserData", definition.PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_REAL_TIME_USER_DATA)
	createConstant(&buffer, "JackpotGetAgentJackpotList", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_AGENT_JACKPOT_LIST)
	createConstant(&buffer, "JackpotSetAgentJackpot", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_SET_AGENT_JACKPOT)
	createConstant(&buffer, "JackpotGetJackpotSetting", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_SETTING)
	createConstant(&buffer, "JackpotSetJackpotSetting", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_SET_JACKPOT_SETTING)
	createConstant(&buffer, "JackpotGetJackpotTokenList", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_TOKEN_LIST)
	createConstant(&buffer, "JackpotCreateJackpotToken", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_CREATE_JACKPOT_TOKEN)
	createConstant(&buffer, "JackpotGetJackpotList", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LIST)
	createConstant(&buffer, "JackpotGetJackpotPoolData", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_POOL_DATA)
	createConstant(&buffer, "JackpotNotifyGameServerAgentJackpotInfo", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_NOTIFY_GAME_SERVER_AGENT_JACKPOT_INFO)
	createConstant(&buffer, "JackpotGetJackpotLeaderBoard", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LEADER_BOARD)
	createConstant(&buffer, "JackpotGetAgentJackpot", definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_AGENT_JACKPOT)
	createConstant(&buffer, "UserGetGameUserBalance", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_BALANCE)
	createConstant(&buffer, "RiskControlGetGameSetting", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_SETTING)
	createConstant(&buffer, "RiskControlSetGameSetting", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_SETTING)
	createConstant(&buffer, "GlobalGetAgentAdminUserPermissionList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_ADMIN_USER_PERMISSION_LIST)
	createConstant(&buffer, "RiskControlGetRealTimeGameRatio", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_REAL_TIME_GAME_RATIO)
	createConstant(&buffer, "UserGetGameUserPlayCountData", definition.PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_PLAY_COUNT_DATA)
	createConstant(&buffer, "RiskControlSetIncomeRatios", definition.PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_INCOME_RATIOS)
	createConstant(&buffer, "RecordGetFriendRoomLogList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_FRIEND_ROOM_LOG_LIST)
	createConstant(&buffer, "UserLoginDeviceUsageRatio", definition.PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_DEVICE_USAGE_RATIO)
	createConstant(&buffer, "UserLoginSourceLocRanking", definition.PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_SOURCE_LOC_RANKING)
	createConstant(&buffer, "UserLoginData", definition.PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_DATA)
	createConstant(&buffer, "GlobalGetUserLoginLogCountryShortList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_USER_LOGIN_LOG_COUNTRY_SHORT_LIST)
	createConstant(&buffer, "GlobalGetUserLoginLogBrowserList", definition.PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_USER_LOGIN_LOG_BROWSER_LIST)
	createConstant(&buffer, "RecordGetUserCreditLogList", definition.PERMISSION_LIST_FEATURE_CODE_RECORD_GET_USER_CREDIT_LOG_LIST)
	createConstant(&buffer, "GameGetCannedList", definition.PERMISSION_LIST_FEATURE_CODE_GAME_GET_CANNED_LIST)
	createConstant(&buffer, "GameSetCannedList", definition.PERMISSION_LIST_FEATURE_CODE_GAME_SET_CANNED_LIST)
	buffer.WriteString("}")

	err := os.WriteFile(path, buffer.Bytes(), 0)
	if err != nil {
		panic(err)
	}
}

func createConstant(buffer *bytes.Buffer, name string, value interface{}) {
	v, _ := json.Marshal(value)
	buffer.WriteString(fmt.Sprintf("\t%s: %s,\n", name, string(v)))
}

func createEnum(buffer *bytes.Buffer, name string, values map[string]int) {
	buffer.WriteString(fmt.Sprintf("\t%s: {\n", name))

	proprtyNames := make([]string, 0, len(values))
	for proprtyName := range values {
		proprtyNames = append(proprtyNames, proprtyName)
	}

	sort.SliceStable(proprtyNames, func(i, j int) bool {
		return values[proprtyNames[i]] < values[proprtyNames[j]]
	})

	for _, proprtyName := range proprtyNames {
		v, _ := json.Marshal(values[proprtyName])
		buffer.WriteString(fmt.Sprintf("\t\t%s: %s,\n", proprtyName, string(v)))
	}

	buffer.WriteString("\t},\n")
}
