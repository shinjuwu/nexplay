package definition

const (
	// 上下分類型 上分:奇數 下分:偶數 (DB方便加總處理)
	WALLET_LEDGER_KIND_ALL                = iota // 全部上下分類型
	WALLET_LEDGER_KIND_API_UP                    // api上分
	WALLET_LEDGER_KIND_API_DOWN                  // api下分
	WALLET_LEDGER_KIND_BACKEND_UP                // 後台上分
	WALLET_LEDGER_KIND_BACKEND_DOWN              // 後台下分
	WALLET_LEDGER_KIND_SINGLE_WALLET_UP          // 單一錢包上分
	WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN        // 單一錢包下分
)

const (
	// 上下分狀態
	WALLET_LEDGER_STATUS_SUCCESS            = iota + 1 // 訂單完成
	WALLET_LEDGER_STATUS_CREATED                       // 訂單建立
	WALLET_LEDGER_STATUS_AGENT_DEDUCTED                // 代理已扣款
	WALLET_LEDGER_STATUS_GAME_SERVER_FAILED            // 遊戲伺服器訂單處理失敗
	WALLET_LEDGER_STATUS_CANCEL                        // 訂單取消
	WALLET_LEDGER_STATUS_CANCEL_FAILED                 // 訂單取消失敗
)
