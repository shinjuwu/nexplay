package aescbc_test

import (
	"encoding/base64"
	"log"
	"monitorservice/pkg/encrypt/aescbc"
	"testing"
)

func TestAescbc(t *testing.T) {

	paramData := "561h96516h51g6fh1g6fh"
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
