package definition

const (
	// 遊戲罐頭語定義
	CannedType_Default = 0 // 預設
	CannedType_Custome = 1 // 自定義

	// 罐頭語預設編號 1~10
	CannedDefault_Start = 1
	CannedDefault_End   = 10
	// 罐頭語自定義編號 11~20
	CannedCustome_Start = 11
	CannedCustome_End   = 20

	// 罐頭語代理開關
	CannedStatus_Close = 0
	CannedStatus_Opne  = 1

	// 館頭語情緒定義
	CannedEmojiType_hp = 1 // 開心
	CannedEmojiType_ex = 2 // 興奮
	CannedEmojiType_cl = 3 // 平靜
	CannedEmojiType_dp = 4 // 沮喪
	CannedEmojiType_ag = 5 // 激動

)

var (
	CANNED_LANG_TYPE_MAP = map[string]bool{
		LANG_TYPE_CHS: true,
		LANG_TYPE_CHT: false,
		LANG_TYPE_EN:  true,
		LANG_TYPE_VI:  true,
		LANG_TYPE_TH:  false,
		LANG_TYPE_PT:  true,
		LANG_TYPE_TR:  true,
	}
)
