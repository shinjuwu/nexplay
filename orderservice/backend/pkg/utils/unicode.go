package utils

import (
	"unicode"
	"unicode/utf8"
)

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func WordLength(str string) int {
	return utf8.RuneCountInString(str)
}
