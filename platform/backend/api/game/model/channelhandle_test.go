package model_test

import (
	"backend/pkg/encrypt/aescbc"
	"backend/pkg/encrypt/md5hash"
	"net/url"
	"testing"
)

const (
	// 測試使用(之後要改成讀取代理自己的 KEY 作加解密)
	DEFAULT_MD5_KEY = "hongkong3345678"  // 15碼
	DEFAULT_AES_KEY = "1234567890123456" // 16碼
)

type ChannelHandleRequest struct {
	Agent     string              `json:"agent"`     // 代理編號（平台提供）
	TimeStamp string              `json:"timestamp"` // 時間戳(Unix 時間戳帶上毫秒),獲取當前時間（1488781836949）
	Param     string              `json:"param"`     // 參數加密字符串
	ParamMap  map[string][]string `json:"param_map"` // 參數解密後
	Key       string              `json:"key"`       // Md5 校驗字符串 Encrypt.MD5(agent+timestamp+MD5Key)
}

func TestChannelHandleRequest(t *testing.T) {

	// data init
	agent := "1"
	timestamp := "1234567890123" // 13碼
	key := md5hash.Hash32bit(agent + timestamp + DEFAULT_MD5_KEY)

	paramData := "s=6&startTime=1488781836949&endTime=1488781836949"

	// url param encrypt by aes
	paramDataAesEncoding, err := aescbc.AesEncrypt([]byte(paramData), []byte(DEFAULT_AES_KEY))
	if err != nil {
		t.Fatalf("EncryptAES err : %v", err)
	}

	param := string(paramDataAesEncoding)

	t.Logf("原始數據: %v", paramData)
	t.Logf("aes 加密後數據: %v", param)

	temp := &ChannelHandleRequest{
		Agent:     agent,
		TimeStamp: timestamp,
		Param:     param,
		ParamMap:  nil,
		Key:       key,
	}

	// aes decrypt
	decryptString, err := aescbc.AesDecrypt([]byte(temp.Param), []byte(DEFAULT_AES_KEY))
	if err != nil {
		t.Fatalf("ParamToJSON err : %v", err)
	}

	temp.ParamMap, err = url.ParseQuery(string(decryptString))
	if err != nil {
		t.Fatalf("ParamToJSON err : %v", err)
	}

	t.Logf("aes 解密後數據: %v", decryptString)
	t.Logf("url parse 後結構數據: %v", temp.ParamMap)

}
