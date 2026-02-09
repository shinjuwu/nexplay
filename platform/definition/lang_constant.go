package definition

const (
	LANG_TYPE_NONE = "All"   // 所有語系
	LANG_TYPE_CHT  = "zh-tw" // 繁體中文
	LANG_TYPE_CHS  = "zh-cn" // 簡體中文
	LANG_TYPE_EN   = "en-us" // 英語 - 美國
	LANG_TYPE_VI   = "vi-vn" // 越南文
	LANG_TYPE_TH   = "th-th" // 泰文
	LANG_TYPE_PT   = "pt-br" // 巴西文
	LANG_TYPE_TR   = "tr-tr" // 土耳其
)

var (
	LANG_TYPE_MAP = map[string]bool{
		LANG_TYPE_CHS: true,
		LANG_TYPE_CHT: true,
		LANG_TYPE_EN:  true,
		LANG_TYPE_VI:  true,
		LANG_TYPE_TH:  true,
		LANG_TYPE_PT:  true,
		LANG_TYPE_TR:  true,
	}
)
