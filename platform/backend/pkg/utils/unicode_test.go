package utils_test

import (
	"backend/pkg/utils"
	"log"
	"testing"
)

func TestUnicode(t *testing.T) {

	testStr := []string{"123 我我我",
		"123我我我",
		"aaaa123我我我",
		"我我我aaaa123我我我",
		"我我我aaaa123",
		"我我我aaaa我我我123我我我",
		"123",
		"123ssssss",
		"sssss123ssssss",
	}

	for _, k := range testStr {
		result := utils.IsChinese(k)
		log.Printf("test string is %v, result is %v", k, result)
	}
	// utils.IsChinese(thirdAccount)
}
