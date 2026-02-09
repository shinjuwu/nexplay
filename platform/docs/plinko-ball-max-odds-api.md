# Plinko 球倍率设定功能 - Game Server API 技術文檔

## 📋 概述

平台後台新增了兩個 API 端點來管理 Plinko 遊戲的球倍率上限設定。Game Server 需要實現相應的端點來配合前端的倍率控制功能。

---

## 🔗 需要實作的 Game Server API

### 1. **取得 Plinko 球倍率上限**

**端點**: `GET /getplinkoballmaxodds`  
**用途**: 查詢指定代理商的 Plinko 遊戲球倍率上限設定

#### 請求參數
```
agentName=代理商名稱
```

#### 請求範例
```http
GET /getplinkoballmaxodds?agentName=test_agent
```

#### 回應格式
```json
{
  "code": 0,
  "message": "success", 
  "data": "{\"max_odds\": 10.0}"
}
```

#### 回應說明
- `code`: 狀態碼 (0=成功, 非0=錯誤)
- `data`: JSON 字串，包含 `max_odds` 欄位 (float64)
- `max_odds`: 當前設定的最大倍率值

---

### 2. **設定 Plinko 球倍率上限**

**端點**: `POST /setplinkoballmaxodds`  
**用途**: 設定 Plinko 遊戲的球倍率上限

#### 請求格式
```json
{
  "agent_name": "test_agent",
  "max_odds": 15.0
}
```

#### 請求參數說明
- `agent_name` (string): 代理商名稱，必填
- `max_odds` (float64): 要設定的最大倍率值，必須大於 0

#### 回應格式
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "agent_name": "test_agent",
    "max_odds": 15.0,
    "updated_at": "2025-11-07T16:30:00Z"
  }
}
```

#### 回應說明
- `code`: 狀態碼 (0=成功, 非0=錯誤)
- `data`: 包含設定後的結果
- `agent_name`: 代理商名稱
- `max_odds`: 設定後的最大倍率值
- `updated_at`: 更新時間 (可選)

---

## 🎮 功能邏輯說明

### 球倍率控制機制
```
遊戲邏輯流程：
1. 玩家投球時，系統隨機產生倍率
2. 隨機倍率範圍：1 倍 ~ max_odds (上限值)
3. 上限值由平台後台動態設定
4. 預設建議值：10 倍
```

### 資料存儲建議
```go
// 建議的資料結構
type PlinkoBallConfig struct {
    AgentName string    `json:"agent_name"` // 代理商名稱
    MaxOdds   float64   `json:"max_odds"`   // 最大倍率
    UpdatedAt time.Time `json:"updated_at"` // 更新時間
}
```

### 資料庫設計建議
```sql
-- 建議的表結構
CREATE TABLE plinko_ball_config (
    agent_name VARCHAR(255) PRIMARY KEY,
    max_odds DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

## 🔧 平台後台調用流程

### 調用時機
1. **查詢**: 前端頁面載入時調用
2. **設定**: 管理員修改倍率上限時調用

### 調用範例

#### 平台發送查詢請求
```http
GET http://game-server:9642/getplinkoballmaxodds?agentName=test_agent
Authorization: Bearer xxx
```

#### 平台發送設定請求  
```http
POST http://game-server:9642/setplinkoballmaxodds
Content-Type: application/json
Authorization: Bearer xxx

{
  "agent_name": "test_agent",
  "max_odds": 20.0
}
```

---

## ⚠️ 重要注意事項

### 1. **資料驗證**
- `agent_name` 不可為空，必須是有效的代理商名稱
- `max_odds` 必須大於 0
- 建議設定合理範圍 (例如：1-100)
- 需要處理無效參數的錯誤情況

### 2. **錯誤處理**
```json
// 參數錯誤回應格式
{
  "code": 1001,
  "message": "agent_name cannot be empty",
  "data": null
}

// 倍率錯誤回應格式
{
  "code": 1002,
  "message": "max_odds must be greater than 0",
  "data": null
}

// 代理商不存在錯誤
{
  "code": 1003,
  "message": "agent not found",
  "data": null
}
```

### 3. **初始值**
- 如果沒有設定過，建議回傳預設值 `1.0`
- 確保遊戲可以正常運行

### 4. **即時生效**
- 設定後立即生效，影響下一次投球
- 無需重啟遊戲服務

---

## 🧪 測試案例

### 測試案例 1: 查詢預設值
```bash
curl -X GET "http://localhost:9642/getplinkoballmaxodds?agentName=test_agent"
# 預期回應: max_odds = 1.0 (預設值)
```

### 測試案例 2: 設定正常值
```bash
curl -X POST http://localhost:9642/setplinkoballmaxodds \
  -H "Content-Type: application/json" \
  -d '{"agent_name": "test_agent", "max_odds": 10.0}'
# 預期回應: code = 0
```

### 測試案例 3: 設定無效值
```bash
curl -X POST http://localhost:9642/setplinkoballmaxodds \
  -H "Content-Type: application/json" \
  -d '{"agent_name": "test_agent", "max_odds": -5.0}'
# 預期回應: code != 0, 錯誤訊息
```

### 測試案例 4: 測試空代理名稱
```bash
curl -X POST http://localhost:9642/setplinkoballmaxodds \
  -H "Content-Type: application/json" \
  -d '{"agent_name": "", "max_odds": 10.0}'
# 預期回應: code != 0, 代理商名稱不可為空
```

### 測試案例 5: 驗證設定生效
```bash
# 1. 設定代理商 test_agent 的倍率上限為 15
curl -X POST http://localhost:9642/setplinkoballmaxodds \
  -H "Content-Type: application/json" \
  -d '{"agent_name": "test_agent", "max_odds": 15.0}'

# 2. 查詢確認
curl -X GET "http://localhost:9642/getplinkoballmaxodds?agentName=test_agent"
# 預期回應: max_odds = 15.0

# 3. 測試不同代理商有不同設定
curl -X POST http://localhost:9642/setplinkoballmaxodds \
  -H "Content-Type: application/json" \
  -d '{"agent_name": "test_agent2", "max_odds": 20.0}'

curl -X GET "http://localhost:9642/getplinkoballmaxodds?agentName=test_agent2"
# 預期回應: max_odds = 20.0

# 4. 實際遊戲測試
# 多次投球，確認產生的倍率都在對應代理商的設定範圍內
```

---

## 📊 平台後台 API 端點

為配合此功能，平台後台已實現以下 API：

### 平台查詢 API
```http
GET /api/v1/game/getplinkoballmaxodds?agentName=代理商名稱&gameId=3003
```

### 平台設定 API
```http
POST /api/v1/game/setplinkoballmaxodds
Content-Type: application/json

{
  "agentName": "代理商名稱",
  "gameId": 3003,
  "maxOdds": 10.0
}
```

---

## 🔄 資料流程圖

```
前端管理介面
    ↓ 查詢/設定
平台後台 API
    ↓ HTTP 請求
Game Server API
    ↓ 存儲/讀取
遊戲配置資料
    ↓ 應用設定
Plinko 遊戲邏輯
    ↓ 隨機倍率
玩家遊戲結果
```

---

## 📞 聯絡資訊

如有技術問題或需要澄清需求，請聯絡：
- **平台後台負責人**: [你的聯絡方式]
- **API 版本**: v1
- **文檔版本**: 1.0
- **最後更新**: 2025-11-07

---

## 📝 變更日誌

| 版本 | 日期 | 變更內容 | 作者 |
|------|------|----------|------|
| 1.0  | 2025-11-07 | 初始版本，新增 Plinko 球倍率設定 API | Platform Team |

---

## 🔗 相關文檔

- [平台後台 API 文檔](./platform-api-docs.md)
- [Game Server 整合指南](./game-server-integration.md)
- [錯誤碼對照表](./error-codes.md)