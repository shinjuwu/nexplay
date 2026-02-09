package utils

import "regexp"

var (
	HttpOrHttpsUrl *regexp.Regexp = regexp.MustCompile(`^(?:(http|https):\/\/)?((?:[\w-]+\.)+[a-z0-9]+)((?:\/[^/?#]*)+)?(\?[^#]+)?(#.+)?$`)

	EnglishAndNumber4To16          *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9]){4,16}$`)
	EnglishAndNumber8To16          *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9]){8,16}$`)
	LowercaseEnglishAndNumber4To16 *regexp.Regexp = regexp.MustCompile(`^(?:[a-z0-9]){4,16}$`)
)
