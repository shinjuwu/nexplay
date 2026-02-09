package aescbc_test

import (
	"backend/pkg/encrypt/aescbc"
	"encoding/base64"
	"log"
	"testing"
)

func TestAescbc(t *testing.T) {
	/*
	   s=0&account=111111&money=100&orderid=1000120170306143036949111111&ip=127.0.0.1&lineCode=text11&KindID=0

	   {"s":100,"m":"/channelHandle","d":{"code":0,"url":"https://h5.ky34.com/index.ht
	   ml?account=10001_111111&token=FBE54A7273EE4F15B363C3F98F32B19F&lang=zh-CN&KindI
	   D=0"}}
	*/

	paramData := "s=0&account=kenny&money=100&orderid=1001&ip=172.30.0.152&linecode=1&kind=0"
	paramData = "s=2&account=test4545&money=100&orderid=1000120170306143036949111111"
	paramData = "s=3&account=test4545&money=100&orderid=1000120170306143036949111112"
	// paramData = "s=1&account=dcc_00000001"
	aesKey := "1234567890123456" //16bit
	paramDataAesEncoding, err := aescbc.AesEncrypt([]byte(paramData), []byte(aesKey))
	if err != nil {
		t.Fatalf("EncryptAES err : %v", err)
	}

	b64Encoding := base64.StdEncoding.EncodeToString([]byte(paramDataAesEncoding))
	log.Printf("original data is %s", paramData)
	log.Printf("AesEncrypt data is %s", paramDataAesEncoding)
	log.Printf("AesEncrypt & base64 encoding data is %s", b64Encoding)

	b64Decoding, _ := base64.StdEncoding.DecodeString(b64Encoding)

	log.Printf("AesEncrypt & base64 decode data is %s", b64Decoding)

	paramDataAesDecoding, err := aescbc.AesDecrypt([]byte(b64Decoding), []byte(aesKey))
	if err != nil {
		t.Fatalf("ParamToJSON err : %v", err)
	}

	log.Printf("AesDecrypt data is %s", string(paramDataAesDecoding))

}
