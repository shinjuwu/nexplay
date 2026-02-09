package global

func CheckTargetLevelCodeIsPassing(my string, target string) (isOk bool) {

	isOk = false
	myLen := len(my)
	targetLen := len(target)
	// level code 長度越短，權限越大
	// 被查詢目標權限大於查詢者
	// level code 格式不對不給查
	if myLen > targetLen || myLen%4 != 0 || targetLen%4 != 0 {
		return
	}
	// 開發者帳號 4碼
	// 如果檢查的是自己
	if my == target {
		isOk = true
		return
	}

	// 查詢目標如果同是開發者，就不給查
	// 平級只能查自己，其他不給查
	if targetLen > 4 {
		if my == target[:myLen] {
			isOk = true
			return
		}
	}

	return
}
