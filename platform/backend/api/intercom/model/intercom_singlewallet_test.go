package model_test

import (
	"backend/pkg/encrypt/aescbc"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"encoding/base64"
	"log"
	"net/url"
	"strconv"
	"testing"
	"time"
)

const (
	// 測試使用(之後要改成讀取代理自己的 KEY 作加解密)
	// DEFAULT_AGENT_ID  = "3"
	// DEFAULT_TIMESTAMP = "1234567890123"    // 13碼
	// DEFAULT_MD5_KEY   = "f7c637cd39679e99" // 16碼
	// DEFAULT_AES_KEY   = "ddxbst648uf7hdbc" // 16碼
	//5dfdb4ea61c520fd76847d764a510fa0

	DEFAULT_AGENT_ID  = "1026"
	DEFAULT_TIMESTAMP = "1234567890123"    // 13碼
	DEFAULT_MD5_KEY   = "0830f101232fdf06" // 16碼
	DEFAULT_AES_KEY   = "78133bb42ace6ac9" // 16碼
)

type ChannelHandleRequest struct {
	Agent     string              `json:"agent"`     // 代理編號（平台提供）
	TimeStamp string              `json:"timestamp"` // 時間戳(Unix 時間戳帶上毫秒),獲取當前時間（1488781836949）
	Param     string              `json:"param"`     // 參數加密字符串
	ParamMap  map[string][]string `json:"param_map"` // 參數解密後
	Key       string              `json:"key"`       // Md5 校驗字符串 Encrypt.MD5(agent+timestamp+MD5Key)
}

func TestSingleWalletRequest(t *testing.T) {

	// data init
	agent := DEFAULT_AGENT_ID
	timestamp := DEFAULT_TIMESTAMP
	key := md5.Hash32bit(agent + timestamp + DEFAULT_MD5_KEY)

	paramData := "s=1&account=kinco"

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

func TestSingleWalletParam(t *testing.T) {
	/*
	   s=0&account=111111&money=100&orderid=1000120170306143036949111111&ip=127.0.0.1&lineCode=text11&KindID=0

	   {"s":100,"m":"/channelHandle","d":{"code":0,"url":"https://h5.ky34.com/index.ht
	   ml?account=10001_111111&token=FBE54A7273EE4F15B363C3F98F32B19F&lang=zh-CN&KindI
	   D=0"}}
	*/
	aesKey := "78133bb42ace6ac9" //16bit
	md5Key := "0830f101232fdf06"
	agentId := 1026
	timestamp := time.Now().UnixMilli()
	log.Printf("時間戳(Unix 時間戳帶上毫秒) is %v", timestamp)

	paramQuery := make(url.Values)
	paramQuery.Add("s", "1")
	paramQuery.Add("account", "kinco")

	paramData := paramQuery.Encode()
	// paramData = "s=1&account=kinco"
	// paramData = "s=0&account=brvnd_brtest0168&money=0&orderid=100220230630133432604brvnd_brtest0168&kind=0"
	// paramData = "s=1&account=kinco"

	paramDataAesEncoding, err := aescbc.AesEncrypt([]byte(paramData), []byte(aesKey))
	if err != nil {
		t.Fatalf("EncryptAES err : %v", err)
	}

	b64Encoding := base64.StdEncoding.EncodeToString([]byte(paramDataAesEncoding))
	log.Printf("參數加密字符串 is %s", b64Encoding)

	ss := strconv.Itoa(agentId) + utils.Int64ToString(timestamp) + md5Key
	s32 := md5.Hash32bit(ss)
	log.Printf("Md5 校驗字符串 is %s", s32)
}
