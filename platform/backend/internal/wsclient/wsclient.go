package wsclient

import (
	"backend/internal/wsclient/client"
	"backend/internal/wsclient/envelope"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type ChatConnInfo struct {
	Scheme       string `json:"scheme"`         // "http" or "https"
	Host         string `json:"host"`           // "127.0.0.1"
	Port         string `json:"port"`           // "8986"
	Domain       string `json:"domain"`         // "127.0.0.1:8986"
	ApiKey       string `json:"api_key"`        // "defaultkey"
	ApiLoginPath string `json:"api_Login_path"` // "/chatservice.api/v1/login"
	WsConnPath   string `json:"ws_conn_path"`   // "/chatservice.ws"
	Platform     string `json:"platform"`       // 平台碼
	Username     string `json:"username"`       // 用戶帳號
	AgentId      string `json:"agent_id"`       // 代理id
}

type ChatServiceManager struct {
	ConnInfo *ChatConnInfo
	WsClient *client.WebSocketClient
	Retries  int // 重試次數
}

func NewChatServiceManager() *ChatServiceManager {
	return &ChatServiceManager{
		ConnInfo: new(ChatConnInfo),
		WsClient: new(client.WebSocketClient),
	}
}

func (p *ChatServiceManager) RegisterEvent(id string, fn func(string)) error {
	return nil
}

func (p *ChatServiceManager) ParseConnInfo(connInfo string) error {
	return json.Unmarshal([]byte(connInfo), &p.ConnInfo)
}

func (p *ChatServiceManager) Login(username string, agentId int) string {
	token := p.getLoginToken(username, agentId)
	if token != "" {
		p.loginWithWebsocketClient(token)
	}

	return token
}

/*
get login token by username and agent id, the server side will be check account exist, if not,
server will create login token and usr account, return token after check.
*/
func (p *ChatServiceManager) getLoginToken(username string, agentId int) string {

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("username", username)
	rawApiQuery.Add("agent_id", fmt.Sprintf("%d", agentId))

	u := url.URL{Scheme: p.ConnInfo.Scheme, Host: p.ConnInfo.Domain, Path: p.ConnInfo.ApiLoginPath, RawQuery: rawApiQuery.Encode()}

	body, err := utils.GetBasicAuthAPI(u.String(), p.ConnInfo.ApiKey, "")
	if err != nil {
		log.Println(err)
	}

	bodyMap := make(map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(body), &bodyMap)

	data := bodyMap["data"].(map[string]interface{})
	if utils.ToInt(bodyMap["code"], -1) != 0 {
		fmt.Println("error api reponse")
		fmt.Println(bodyMap["code"])
		return ""
	}
	token := utils.ToString(data["token"], "")
	if token == "" {
		fmt.Println("token is empty")
		return ""
	}

	return token
}

/*
use login token to create websocket and connect server.
*/
func (p *ChatServiceManager) loginWithWebsocketClient(token string) {
	// use token to login ws
	rawQuery := make(url.Values)
	rawQuery.Add("token", token)

	wsclient, err := client.NewWebSocketClient(p.ConnInfo.Domain, p.ConnInfo.WsConnPath, rawQuery.Encode())
	if err != nil {
		log.Printf("create websocket connect fialed, err: %v", err)
	}

	p.WsClient = wsclient
}

// 跑馬燈發話
func (p *ChatServiceManager) SendMarquee(msg string) error {

	m := envelope.NewMessage(envelope.CTWC_MARQUEE, msg, 0)

	return p.WsClient.Write(m.ToJson())
}
