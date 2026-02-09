package md5hash_test

import (
	"log"
	md5 "monitorservice/pkg/encrypt/md5hash"
	"testing"
)

func TestMD5Hash(t *testing.T) {

	ss := "11234567890123hongkong3345678"
	s32 := md5.Hash32bit(ss)
	s16 := md5.Hash16bit(ss)

	log.Println(s32) //5d41402abc4b2a76b9719d911017c592
	log.Println(s16) //bc4b2a76b9719d91
}
