package system

import (
	"backend/internal/ginweb/response"
	"backend/pkg/encrypt/aescbc"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"encoding/base64"
	"net/url"
	"strconv"
	"time"
)

func sendApiServer(connInfo map[string]interface{}, agentId int64, params, aesKey, md5Key string) (*response.Response, error) {
	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")

	agent := strconv.FormatInt(agentId, 10)
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	key := md5.Hash32bit(agent + timestamp + md5Key)

	paramBytes, _ := aescbc.AesEncrypt([]byte(params), []byte(aesKey))
	param := base64.StdEncoding.EncodeToString(paramBytes)

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("agent", agent)
	rawApiQuery.Add("timestamp", timestamp)
	rawApiQuery.Add("param", param)
	rawApiQuery.Add("key", key)

	u := url.URL{Scheme: scheme, Host: domain, Path: path}

	resultStr, err := utils.GetAPI(u.String(), rawApiQuery.Encode(), "")
	if err != nil {
		return nil, err
	}

	result := &response.Response{}
	utils.ToStruct([]byte(resultStr), result)

	return result, nil
}
