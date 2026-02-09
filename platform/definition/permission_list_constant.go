package definition

// test
const (
	PERMISSION_LIST_FEATURE_CODE_EXAMPLE_HEALTH = iota + 100001 // 檢查伺服器是否活著
)

// 遊戲SERVER串接使用
const (
	PERMISSION_LIST_FEATURE_CODE_INTERCOM_GET_LOGIN_TOKEN     = iota + 100100 // client取得驗證token
	PERMISSION_LIST_FEATURE_CODE_INTERCOM_LOGIN_GAME                          // 驗證token並登入遊戲
	PERMISSION_LIST_FEATURE_CODE_INTERCOM_LOGOUT_GAME                         // 遊戲用戶登出
	PERMISSION_LIST_FEATURE_CODE_INTERCOM_CREATE_GAME_RECORD                  // 創建每局遊戲紀錄
	PERMISSION_LIST_FEATURE_CODE_INTERCOM_GET_MARQUEE_SETTING                 // 此接口供遊戲伺服器取得跑馬燈設定列表
)

// 後台使用
const (
	PERMISSION_LIST_FEATURE_CODE_AGENT_CREATE_AGENT                               = iota + 100200 // 創建代理帳號
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_LIST                                             // 取得代理底下所有代理資料
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_SECRET_KEY                                       // 秘鑰資訊顯示
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_COIN_SUPPLY_INFO                                 // 取得指定代理補分相關資料設定
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_COIN_SUPPLY_INFO                                 // 修改指定代理補分相關資料設定
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_GAME_LIST                                        // 取得代理遊戲列表
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_GAME_STATE                                       // 設置代理遊戲狀態
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_GAME_ROOM_LIST                                   // 取得代理遊戲房間列表
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_GAME_ROOM_STATE                                  // 設置代理遊戲房間狀態
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION_TEMPLATE_INFO                         // 取得代理權限群組權限樣板
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION_LIST                                  // 取得代理權限群組列表
	PERMISSION_LIST_FEATURE_CODE_AGENT_CREATE_AGENT_PERMISSION                                    // 創建代理權限群組
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_PERMISSION                                       // 修改代理權限群組
	PERMISSION_LIST_FEATURE_CODE_AGENT_DELETE_AGENT_PERMISSION                                    // 刪除代理權限群組
	PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_LIST                                               // 取得遊戲列表
	PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_STATE                                              // 修改遊戲狀態
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_RELOAD_GLOBAL_DATA                                        // 此接口用來重新載入本地端資料
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_CHECK_GAME_ROOM_SETTING                                   // 此接口用來檢查並重設 game_room setting
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_LIST                                            // 取得全部代理商list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_ALL_GAME_LIST                                         // 取得全部遊戲list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_GAME_LIST                                             // 取得上線及維護中的遊戲list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_ROOM_TYPE_LIST                                        // 取得房間類型list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_PERMISSION_LIST                                 // 取得權限群組層級list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_LOGIN_LOGIN                                                      // 用戶登錄
	PERMISSION_LIST_FEATURE_CODE_LOGIN_CAPTCHA                                                    // 取得驗證碼
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MARQUEE_LIST                                          // 取得目前跑馬燈設定列表：開發者、總代理、子代理，皆可查看跑馬燈資訊全內容
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MARQUEE                                               // 指定取得某筆跑馬燈設定
	PERMISSION_LIST_FEATURE_CODE_MANAGE_CREATE_MARQUEE                                            // 添加跑馬燈功能（開發者）：管理後台才可添加【活動】類型跑馬燈
	PERMISSION_LIST_FEATURE_CODE_MANAGE_UPDATE_MARQUEE                                            // 編輯跑馬燈功能（開發者）：管理後台才可編輯【活動】類型跑馬燈
	PERMISSION_LIST_FEATURE_CODE_MANAGE_DELETE_MARQUEE                                            // 刪除跑馬燈功能（開發者）：管理後台才可刪除【活動】類型跑馬燈
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_USER_PLAY_LOG_LIST                                    // 取得個人遊戲紀錄列表
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_PLAY_LOG_COMMON                                       // 取得遊戲局記錄
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_WALLET_LEDGER_LIST                                    // 取得帳變資料列表
	PERMISSION_LIST_FEATURE_CODE_USER_PING                                                        // ping
	PERMISSION_LIST_FEATURE_CODE_USER_GET_ALIVE_TOKEN_LIST                                        // 取得目前已產生的有效 token list
	PERMISSION_LIST_FEATURE_CODE_USER_BLACK_TOKEN                                                 // 將用戶 token 列入黑名單(主動使登入token 失效)
	PERMISSION_LIST_FEATURE_CODE_USER_CREATE_ADMIN_USER                                           // 創建後台帳號(只能創建自己的後台帳號)
	PERMISSION_LIST_FEATURE_CODE_USER_GET_ADMIN_USERS                                             // 依照查詢者角色權限列出自身權限下的子帳號列表
	PERMISSION_LIST_FEATURE_CODE_USER_GET_ADMIN_USER_INFO                                         // 指定查詢某後台帳號狀態
	PERMISSION_LIST_FEATURE_CODE_USER_UPDATE_ADMIN_USER_INFO                                      // 指定設定某後台帳號狀態
	PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USERS                                              // 依照查詢者角色權限列出遊戲會員帳號清單
	PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_INFO                                          // 指定查詢某遊戲會員帳號信息
	PERMISSION_LIST_FEATURE_CODE_USER_UPDATE_GAME_USER_INFO                                       // 指定修改某遊戲會員帳號信息
	PERMISSION_LIST_FEATURE_CODE_CAL_CAL_PERFORMANCE_REPORT                                       // 此接口用來重新計算業績報表
	PERMISSION_LIST_FEATURE_CODE_CAL_GET_JOB_SHEDULER_LIST                                        // 此接口用來取得當前 job 的資訊清單
	PERMISSION_LIST_FEATURE_CODE_CAL_GET_PERFORMANCE_REPORT                                       // 取得指定時間區段代理時間區段的統計資料
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_IP_WHITELIST_LIST                                // 取得代理後台IP資訊列表
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_IP_WHITELIST                                     // 取得代理後台IP資訊
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_IP_WHITELIST                                     // 設置代理後台IP資訊
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_PERMISSION                                       // 取得代理權限群組
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_WALLET_LIST                                      // 取得代理錢包餘額列表
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_WALLET                                           // 設置代理錢包餘額
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_AGENT_WALLET_LEDGER_LIST                              // 取得代理分數紀錄列表
	PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_WALLET_LIST                                   // 取得玩家錢包餘額列表
	PERMISSION_LIST_FEATURE_CODE_USER_SET_GAME_USER_WALLET                                        // 設置玩家錢包餘額
	PERMISSION_LIST_FEATURE_CODE_RECORD_CONFIRM_WALLET_LEDGER                                     // 更新玩家分數紀錄狀態
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_STAT_DATA                                             // 此接口用來取得當天資訊總覽的資料
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_RISK_USER_LIST                                        // 此接口用來取得今日風險玩家清單(前100名)
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_GAME_LEADERBOARDS                                     // 此接口用來取得今日遊戲輸贏排行榜
	PERMISSION_LIST_FEATURE_CODE_NOTIFY_GET_CHAT_SERVICE_CONN_INFO                                // 此接口用來取得chat service 連線資訊
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_INCOME_RATIO_LIST                                // 此接口用來取得當前殺數設定列表(只有管理可以使用)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_INCOME_RATIO                                     // 此接口用來取得指定id殺數設定資料(只有管理可以使用)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_INCOME_RATIO                                     // 此接口用來設定指定id殺數設定資料(只有管理可以使用)
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_DAYS_BETTOR_DATA                             // 此接口用來取得最近一段時間(預設30日)的活躍玩家數&日投注人數
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_TOTAL_SCORE_DATA                             // 此接口用來取得今日各時段輸贏(昨日&今日)
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_TOTAL_BETTOR_INFO_DATA                       // 此接口用來取得今日各時段(昨日&今日)
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_BACKEND_ACTION_LOG                                    // 取得後台操作紀錄列表
	PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_SERVER_STATE                                       // 取得遊戲server狀態
	PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_SERVER_STATE                                       // 設置遊戲server狀態
	PERMISSION_LIST_FEATURE_CODE_GAME_NOTIFY_GAME_SERVER                                          // 創建更新遊戲相關設定(遊戲server維護中才可以使用)
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_ANNOUNCEMENT_LIST                                     // 取得目前後台公告設定列表：開發者、總代理、子代理，皆可查看跑馬燈資訊全內容
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_ANNOUNCEMENT                                          // 指定取得某筆後台公告設定
	PERMISSION_LIST_FEATURE_CODE_MANAGE_CREATE_ANNOUNCEMENT                                       // 添加後台公告功能（開發者）
	PERMISSION_LIST_FEATURE_CODE_MANAGE_UPDATE_ANNOUNCEMENT                                       // 編輯後台公告功能（開發者）
	PERMISSION_LIST_FEATURE_CODE_MANAGE_DELETE_ANNOUNCEMENT                                       // 刪除後台公告功能（開發者）
	PERMISSION_LIST_FEATURE_CODE_CAL_GET_PERFORMANCE_REPORT_LIST                                  // 取得指定時間區段代理總和的資料
	PERMISSION_LIST_FEATURE_CODE_USER_SET_PERSONAL_INFO                                           // 修改個人資訊
	PERMISSION_LIST_FEATURE_CODE_USER_SET_PERSONAL_PASSWORD                                       // 修改個人密碼
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO_LIST                          // 取得當前總代理風控設定資料列表（只有管理）
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO                               // 取得指定ID總代理風控設定資料（只有管理）
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AGENT_INCOME_RATIO                               // 設定指定ID總代理風控設定資料（只有管理）
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_CUSTOM_TAG_SETTING_LIST                    // 取得玩家標示設定資料(只有總代理、子代理可以使用)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AGENT_CUSTOM_TAG_SETTING_LIST                    // 設定玩家標示設定資料(只有總代理、子代理可以使用)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USERS_CUSTOM_TAG_LIST                       // 取得玩家標示資料列表(開發商（營運商）可看到全部代理商資訊；總代理、子代理可看到自身&下級)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USERS_CUSTOM_TAG                            // 取得標示玩家資料(開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_USERS_CUSTOM_TAG                            // 設定標示玩家資料(開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AGENT_INCOME_RATIO_AND_GAME_DATA                 // 此接口用來取得代理設定機率&遊戲輸贏結果
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_CUSTOM_TAG_SETTING_LIST                         // 此接口用來取得玩家標示設定下拉選單資料list(供下拉選單使用,Drop Down Menu)
	PERMISSION_LIST_FEATURE_CODE_SYSTEM_GET_SERVER_SETTING                                        // 此接口用來取得server相關設定的參數
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_BACKEND_LOGIN_LOG_LIST                                // 取得後台登入紀錄列表
	PERMISSION_LIST_FEATURE_CODE_GAME_GET_GAME_ICON_LIST                                          // 取得遊戲icon list
	PERMISSION_LIST_FEATURE_CODE_GAME_SET_GAME_ICON_LIST                                          // 設置遊戲遊戲icon list
	PERMISSION_LIST_FEATURE_CODE_AGENT_GET_AGENT_API_IP_WHITELIST                                 // 取得代理API IP資訊
	PERMISSION_LIST_FEATURE_CODE_AGENT_SET_AGENT_API_IP_WHITELIST                                 // 設置代理API IP資訊
	PERMISSION_LIST_FEATURE_CODE_SYSTEM_GET_EXCHANGE_DATA_LIST                                    // 取得匯率設定
	PERMISSION_LIST_FEATURE_CODE_SYSTEM_SET_EXCHANGE_DATA_LIST                                    // 設定匯率設定
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AUTO_RISKCONTROL_SETTING                         // 取得自動風控設定
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_AUTO_RISKCONTROL_SETTING                         // 設定自動風控設定
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_USER_RISKCONTROL_TAG                        // 取得指定玩家處置設定
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_USER_RISKCONTROL_TAG                        // 設定指定玩家處置設定
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_AUTO_RISKCONTROL_LOG_LIST                        // 取得自動風控紀錄
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_AGENT_GAME_RATIO_STAT_LIST                            // 取得日結算報表列表
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_GAME_USERS_STAT_HOUR_LIST                             // 取得玩家遊玩紀錄列表
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_MAINTAIN_PAGE_SETTING                                 // 取得維護頁設定(只有開發商(營用商)可以使用)
	PERMISSION_LIST_FEATURE_CODE_MANAGE_SET_MAINTAIN_PAGE_SETTING                                 // 設定維護頁設定(只有開發商(營用商)可以使用)
	PERMISSION_LIST_FEATURE_CODE_USER_RESET_PASSWORD                                              // 重置密碼
	PERMISSION_LIST_FEATURE_CODE_MANAGE_GET_INTERVAL_REAL_TIME_USER_DATA                          // 此接口用來取得今日昨日的各時段線上人數
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_AGENT_JACKPOT_LIST                                   // 取得總代理jackpot設定列表(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_SET_AGENT_JACKPOT                                        // 設定總代理jackpot(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_SETTING                                      // 取得平台jackpot設定(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_SET_JACKPOT_SETTING                                      // 設定平台jackpot設定(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_TOKEN_LIST                                   // 取得jackpot代幣紀錄列表(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_CREATE_JACKPOT_TOKEN                                     // 建立jackpot代幣(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LIST                                         // 取得jackpot紀錄列表
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_POOL_DATA                                    // 取得jackpot獎池資訊(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_NOTIFY_GAME_SERVER_AGENT_JACKPOT_INFO                    // server同步jackpot資訊(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LEADER_BOARD                                 // 取得jackpot玩家貢獻度(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_AGENT_JACKPOT                                        // 取得指定總代理jackpot設定(只有開發商(營運商)可使用)
	PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_BALANCE                                       // 取得玩家目前餘額
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_GAME_SETTING                                     // 此接口用來取得遊戲基礎設定(只有開發商（營運商）可使用)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_GAME_SETTING                                     // 此接口用來設定遊戲基礎設定(只有開發商（營運商）可使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_AGENT_ADMIN_USER_PERMISSION_LIST                      // 此接口用來取得代理父帳號權限list(供下拉選單使用,Drop Down Menu)
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_GET_REAL_TIME_GAME_RATIO                             // 此接口用來取得遊戲即時殺率資訊(只有開發商（營運商）可使用)
	PERMISSION_LIST_FEATURE_CODE_USER_GET_GAME_USER_PLAY_COUNT_DATA                               // 此接口用來取得玩家目前遊戲局數狀態
	PERMISSION_LIST_FEATURE_CODE_RISKCONTROL_SET_INCOME_RATIOS                                    // 此接口用來批次設定殺數設定資料(只有管理可以使用)
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_FRIEND_ROOM_LOG_LIST                                  // 取得好友房建房紀錄列表
	PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_DEVICE_USAGE_RATIO                                    // 此接口用來取得今日用戶登入裝置使用比例
	PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_SOURCE_LOC_RANKING                                    // 此接口用來取得今日用戶登入來源位置排行
	PERMISSION_LIST_FEATURE_CODE_USER_LOGIN_DATA                                                  // 此接口用來取得玩家登入資訊列表
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_USER_LOGIN_LOG_COUNTRY_SHORT_LIST                     // 此接口用來取得用戶登入國家資料list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_GLOBAL_GET_USER_LOGIN_LOG_BROWSER_LIST                           // 此接口用來取得用戶登入瀏覽器資料list(供下拉選單使用)
	PERMISSION_LIST_FEATURE_CODE_RECORD_GET_USER_CREDIT_LOG_LIST                                  // 取得玩家帳變紀錄列表(only 管理可用)
	PERMISSION_LIST_FEATURE_CODE_GAME_GET_CANNED_LIST                                             // 取得遊戲 canned list
	PERMISSION_LIST_FEATURE_CODE_GAME_SET_CANNED_LIST                                             // 設置遊戲 canned list
)

// 對外串接
const (
	PERMISSION_LIST_FEATURE_CODE_CHANNEL_HANDLE = iota + 200101
	PERMISSION_LIST_FEATURE_CODE_GET_RECORD_HANDLE
)
