package model

import (
	"backend/pkg/encrypt/aescbc"
	"definition"
	"encoding/base64"
	"net/url"
	"strconv"
)

const (
	SingleWalletRequest_Command_QueryMinusScore = 1  // 查詢平台可下分
	SingleWalletRequest_Command_AddScore        = 2  // 上分(平台扣款)
	SingleWalletRequest_Command_MinusScore      = 3  // 下分(平台加款)
	SingleWalletRequest_Command_CancelAddScore  = 11 // 取消上分(平台扣款)

	SingleWalletRequest_Command_CancelAddScore_resend = -999 // 取消上分(平台扣款) 重送使用旗標
)

func DispatchSWPlatformCode(platform string) string {
	ret := 0
	switch platform {
	case "dev":
		ret = 1
	case "qa":
		ret = 2
	case "ete":
		ret = 3
	case "pro":
		ret = 4
	default:
		ret = -1
	}

	if ret > 0 {
		return strconv.Itoa(ret)
	} else {
		return ""
	}
}

func CreateSingleWalletParamQueryWithCancel(command, cancelCommand int, username, point, orderId, gameId, roomId, currency, betId, wallerLedgerId string) (url.Values, url.Values) {

	return CreateSingleWalletParamQuery(command, username, point, orderId, gameId, roomId, currency, betId, wallerLedgerId),
		CreateSingleWalletParamQuery(cancelCommand, username, point, orderId, gameId, roomId, currency, betId, wallerLedgerId)
}

// 產生單一錢包網址參數
func CreateSingleWalletParamQuery(command int, username, point, orderId, gameId, roomId, currency, betId, wallerLedgerId string) url.Values {
	paramQuery := make(url.Values)
	paramQuery.Add("s", strconv.Itoa(command))
	paramQuery.Add("account", username)
	switch command {
	case SingleWalletRequest_Command_AddScore:
		fallthrough
	case SingleWalletRequest_Command_MinusScore:
		fallthrough
	case SingleWalletRequest_Command_CancelAddScore:
		paramQuery.Add("point", point)
		paramQuery.Add("orderid", orderId)
		paramQuery.Add("gameid", gameId)
		paramQuery.Add("roomid", roomId)
		paramQuery.Add("currency", currency)
		paramQuery.Add("wallerledgerid", wallerLedgerId) // 有上分才有下分
		if command == SingleWalletRequest_Command_MinusScore {
			paramQuery.Add("betid", betId)
		}
	}

	return paramQuery
}

// 參數加密字符串, paramQueryEncode: 欲加密的字串,
func CreateParamEncpy(paramQueryEncode, aesKey string) string {
	paramDataAesEncoding, err := aescbc.AesEncrypt([]byte(paramQueryEncode), []byte(aesKey))
	if err != nil {
		return ""
	}

	b64Encoding := base64.StdEncoding.EncodeToString([]byte(paramDataAesEncoding))

	return b64Encoding
}

func CreateSingleWalletCallbackURL(platform, agentStr, b64EncodingParam, timestampStr, md5hs32, connInfoScheme, connInfoDomain, connInfoPath string) url.URL {
	platfromCode := DispatchSWPlatformCode(platform)
	rawApiQuery := make(url.Values)
	rawApiQuery.Add("agent", agentStr)
	rawApiQuery.Add("param", b64EncodingParam)
	rawApiQuery.Add("timestamp", timestampStr)
	rawApiQuery.Add("key", md5hs32)
	if platfromCode != "" {
		rawApiQuery.Add("u", platfromCode)
	}

	return url.URL{
		Scheme:   connInfoScheme,
		Host:     connInfoDomain,
		Path:     connInfoPath,
		RawQuery: rawApiQuery.Encode()}
}

// 遊戲單一錢包(request from game)
// 取餘額: 只需帶入 UserId、Username
// WalletLedgerId: 只有下分要帶入，此碼在上分時取得
type SingleWalletRequest struct {
	Command        int     `json:"command"`          // 命令(1:取餘額；2:上分；3:下分)
	UserId         int     `json:"user_id"`          // 玩家id
	Username       string  `json:"username"`         // 玩家帳號
	Point          float64 `json:"point"`            // 目前點數
	GameId         int     `json:"game_id"`          // 遊戲id
	RoomId         int     `json:"room_id"`          // 房間id
	BetId          string  `json:"bet_id"`           // 局號
	WalletLedgerId string  `json:"wallet_ledger_id"` // 單一錢包上下分群組識別碼
}

func NewEmptySingleWalletRequest() *SingleWalletRequest {
	return &SingleWalletRequest{}
}

func (p *SingleWalletRequest) CheckParam() bool {
	if p.Command <= 0 || p.UserId <= 0 || p.Username == "" {
		return false
	}

	switch p.Command {
	case SingleWalletRequest_Command_QueryMinusScore:
		p.Point = 0
	case SingleWalletRequest_Command_AddScore:
		if p.Point <= 0 || p.GameId <= definition.GAME_ID_ALL || p.RoomId < definition.ROOM_TYPE_NEWBIE || p.WalletLedgerId == "" {
			return false
		}
	case SingleWalletRequest_Command_MinusScore:
		// tip:下分不需要在遊戲內
		if p.Point <= 0 || p.GameId <= definition.GAME_ID_ALL || p.RoomId < definition.ROOM_TYPE_NEWBIE || p.WalletLedgerId == "" {
			return false
		}
	default:
		return false
	}

	return true

}

// 遊戲單一錢包(request from game)
type SingleWalletResponse struct {
	Command        int     `json:"command"`          // 命令(1:取餘額；2:上分；3:下分)
	Money          float64 `json:"money"`            // 平台餘額
	Point          float64 `json:"point"`            // 目前點數
	WalletLedgerId string  `json:"wallet_ledger_id"` // 單一錢包上下分群組識別碼
}

func NewEmptySingleWalletResponse() *SingleWalletResponse {
	return &SingleWalletResponse{
		Command: -1,
		Point:   .0,
	}
}

// 單一錢包加密參數送外部平台
type SingleWalletParamsToPlatform struct {
	AgentId   string              `json:"agent_id"`  // 代理id
	Param     string              `json:"param"`     // 參數加密字符串(aes 加密)
	ParamMap  map[string][]string `json:"param_map"` // 參數解密後
	Key       string              `json:"key"`       // Md5 校驗字符串 Encrypt.MD5(agent+timestamp+MD5Key)
	TimeStamp string              `json:"timestamp"` // 時間戳(Unix 時間戳帶上毫秒),獲取當前時間（1488781836949）
}
