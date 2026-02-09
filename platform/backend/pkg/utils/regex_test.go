package utils_test

import (
	"backend/pkg/utils"
	"log"
	"testing"
)

func TestRegex(t *testing.T) {
	links := []string{
		"http://yahoo.com/index.html",
		"http://123456.com/index.html",
		"http://yes123.com/index",
		"https://yahoo.com/123456",
		"https://123456.1.123/index.html",
		"https://yes123.com/index.html",
	}

	for i := 0; i < len(links); i++ {
		result := utils.HttpOrHttpsUrl.MatchString(links[i])
		log.Printf("link: %s matching result is %v", links[i], result)
	}

	type RegexTest struct {
		str string // 測試字串
		r1  bool   // EnglishAndNumber4To16 result
		r2  bool   // EnglishAndNumber8To16 result
		r3  bool   // LowercaseEnglishAndNumber4To16 result
	}

	regexTestTable := []*RegexTest{
		{str: "12345678", r1: true, r2: true, r3: true},
		{str: "1234567890123456", r1: true, r2: true, r3: true},
		{str: "12345678901234567", r1: false, r2: false, r3: false},
		{str: "1a2A3z4Z56789", r1: true, r2: true, r3: false},
		{str: "1a2A3z4Z56789@", r1: false, r2: false, r3: false},
		{str: "aaa大俠好帥aaa", r1: false, r2: false, r3: false},
		{str: "", r1: false, r2: false, r3: false},
		{str: "  a  a  ", r1: false, r2: false, r3: false},
		{str: "123", r1: false, r2: false, r3: false},
		{str: "1234", r1: true, r2: false, r3: true},
		{str: "123@", r1: false, r2: false, r3: false},
		{str: "12345678901234567890", r1: false, r2: false, r3: false},
	}

	for _, regexTest := range regexTestTable {
		r1 := utils.EnglishAndNumber4To16.MatchString(regexTest.str)
		r2 := utils.EnglishAndNumber8To16.MatchString(regexTest.str)
		r3 := utils.LowercaseEnglishAndNumber4To16.MatchString(regexTest.str)

		log.Printf("case: %s, r1 expected %v actual %v, ", regexTest.str, regexTest.r1, r1)
		log.Printf("case: %s, r2 expected %v actual %v, ", regexTest.str, regexTest.r2, r2)
		log.Printf("case: %s, r3 expected %v actual %v, ", regexTest.str, regexTest.r3, r3)
	}
}
