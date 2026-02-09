package controller_test

import (
	"backend/pkg/encrypt/aescbc"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"encoding/base64"
	"log"
	"strconv"
	"testing"
	"time"
)

// 參數加解密測式
func TestEncryptChannelhandleParam(t *testing.T) {
	/*
	   s=0&account=111111&money=100&orderid=1000120170306143036949111111&ip=127.0.0.1&lineCode=text11&KindID=0

	   {"s":100,"m":"/channelHandle","d":{"code":0,"url":"https://h5.ky34.com/index.ht
	   ml?account=10001_111111&token=FBE54A7273EE4F15B363C3F98F32B19F&lang=zh-CN&KindI
	   D=0"}}
	*/
	aesKey := "ddxbst648uf7hdbc" //16bit
	md5Key := "f7c637cd39679e99"
	agentId := 3
	timestamp := time.Now().UnixMilli()
	log.Printf("時間戳(Unix 時間戳帶上毫秒) is %v", timestamp)

	paramData := "s=0&account=kinco&money=0&orderid=1000120170306143076949222244&ip=172.30.0.152&linecode=1&kind=1001"
	// paramData = "s=1&account=kenny&money=100&orderid=1000120170306143036949111111"
	// paramData = "s=2&account=kenny&money=100&orderid=1000121171306143036949111111"
	// paramData = "s=3&account=kinco&money=100&orderid=100012017030614303694911555"
	// paramData = "s=1&account=dcc_00000001"
	paramData = "s=6&start_time=1661495400&end_time=1661496000"

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
