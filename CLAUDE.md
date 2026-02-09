# CLAUDE.md

本文件為 Claude Code (claude.ai/code) 在此倉庫中操作時提供指引。

## 專案概述

NextPlay (DCC) 是一個遊戲平台，由多個 Go 微服務和 Vue 3 前端管理後台組成。此倉庫為 monorepo 架構，各服務獨立運作並共用一個 definition 常數模組。

## 倉庫結構

```
nextplay/
├── platform/          # 主要後端服務 (Go 1.22.3, module: backend)
│   ├── backend/       # 應用程式碼
│   └── definition/    # 共用常數庫 (git submodule)
├── chatservice/       # WebSocket 聊天服務 (Go 1.22.3, module: chatservice)
├── orderservice/      # 訂單/遊戲紀錄服務 (Go 1.18, module: backend)
│   └── backend/       # 應用程式碼
├── monitorservice/    # 監控服務 (Go 1.22.3, module: monitorservice)
│   ├── cmd/baseserver/  # 主程式入口 + config.yml + Dockerfile
│   ├── cmd/baseclient/  # 測試用 Console client
│   ├── internal/        # API 框架、模組系統
│   ├── pkg/             # 共用套件
│   ├── migrate/         # SQL 遷移腳本
│   └── deployments/     # Docker、systemctl、nginx 部署設定
├── dcc_monitor/       # 監控前端 (Vue 3 + TypeScript + Vite)
│   └── frontend/      # 獨立 SPA，與 dcc_front 分離
└── dcc_front/         # 前端
    └── frontend/      # Vue 3 + Vite 多應用管理後台
```

## 建置與執行指令

### 前端 (dcc_front/frontend/)

```bash
npm install

# 開發伺服器
npm run manager-dev    # 管理後台開發伺服器
npm run agent-dev      # 代理端開發伺服器

# 建置 (依環境區分)
npm run build-manager-dev / build-agent-dev
npm run build-manager-qa  / build-agent-qa
npm run build-manager-qc  / build-agent-qc

# 預覽建置產出
npm run preview-manager / npm run preview-agent

# 程式碼檢查與格式化
npm run lint           # ESLint 自動修正
npm run format         # Prettier 格式化
```

### Go 後端服務

```bash
# 建置任一服務
cd <service>/backend   # (chatservice 則直接 cd chatservice)
go build -o <binary_name> .

# 執行測試
go test ./...
go test ./pkg/jwt -v           # 單一套件測試
go test ./internal/... -cover  # 含覆蓋率

# 產生 Swagger 文件 (platform)
cd platform/backend
swag init

# Platform Makefile 目標
cd platform
make all      # swag + 建置 + 部署 + 重啟
make build    # 僅編譯
make swag     # 產生 Swagger 文件
```

### Docker 建置

每個服務都有 Dockerfile，使用多階段建置 (golang -> alpine)。建置參數 `web_mode` 用於選擇設定檔：

```bash
docker-compose -f deployments/docker-compose.yml up
```

## 架構說明

### Go 服務共同模式

所有 Go 服務遵循相同架構：

- **Web 框架**：Gin，搭配 CORS、日誌記錄、例外處理中介層
- **資料庫**：PostgreSQL，透過 `lib/pq` + `squirrel` 查詢建構器
- **資料庫遷移**：`golang-migrate` (orderservice/platform) 或自訂遷移工具 (chatservice)
- **快取**：Redis (go-redis v7/v9)
- **認證**：JWT (`golang-jwt/jwt/v4`) + 驗證碼
- **日誌**：Zap + lumberjack 日誌輪替
- **設定檔**：依環境區分的 YAML 檔案 (`config.yml`、`config-dev.yml`、`config-docker.yml` 等)
- **排程任務**：`robfig/cron` 背景排程

**路由模式** (ApiCluster)：
```
gin.Engine → Middleware → ApiCluster.RouterGroupRegister()
  → IApiEach.RegisterApiRouter() 依功能群組註冊
```

**標準套件目錄結構**：
- `api/` — 依功能分組的 API 處理器，含 `controller/` + `model/` 子目錄
- `server/router/` — 路由初始化與中介層設定
- `server/global/` — 全域狀態 (DB、Redis、設定檔參照)
- `pkg/` — 可重用套件 (cache, config, database, encrypt, jwt, logger, redis, utils)
- `internal/` — 服務專用內部套件
- `db/migrations/` — SQL 遷移檔案

**設定載入順序**：YAML → Logger → Database → Redis → JWT → Captcha → JobScheduler → Router → Gin server

### 共用 Definition 模組

`platform/definition/` 是一個 git submodule，包含所有後端服務共用的常數。錯誤碼範圍：
- 0：成功；0-99：後台；600-699：內部溝通；800-899：chatservice；1001-1999：外部 API

### 前端架構

**技術棧**：Vue 3 (Composition API, `<script setup>`) + Vite + Tailwind CSS + Pinia

**多應用建置**：兩個獨立 SPA 共用同一個基礎層：
- **Manager** (`src/manager/`) — 管理後台 (header: `M7NJpSXxh`)
- **Agent** (`src/agent/`) — 代理端 (header: `aikDg57I1`)
- **Shared** (`src/base/`) — 共用元件、API 模組、Store、Composables、i18n

**關鍵模式**：
- API 模組位於 `src/base/api/sys*.js` — Axios 搭配 `Dcc-Token` header，base URL `/api/v1`
- Pinia stores 位於 `src/base/store/` — `userStore` (認證/權限)、`breadcrumbStore`、`chatServiceStore`
- 路由守衛檢查 localStorage 中的 token，若不存在則導向 `/Login`
- i18n：中文 (zh-cn.json)，錯誤碼翻譯透過 `errorCode__${code}` 鍵值
- 遊戲日誌解析器位於 `src/base/common/gameLog/`，支援 30+ 種遊戲類型

**程式碼風格**：無分號、單引號、120 字元寬度、尾隨逗號 (es5)、CRLF 換行

### MonitorService 監控服務

**功能**：遊戲平台的即時監控與資料分析系統

- **服務狀態監控** — 監控 platform、orderservice、chatservice 等服務的線上狀態
- **玩家錢包監控** — 追蹤玩家上下分 (coin in/out) 交易
- **異常輸贏監控** — 偵測異常的遊戲輸贏模式
- **RTP 監控** — 追蹤各遊戲的 Return-to-Player 統計 (日/週/月)
- **資料收集** — 接收外部平台推送的遊戲資料

**API 路徑結構**：
- `/monitor/login.api/v1/*` — 登入與驗證碼
- `/monitor/account.api/v1/*` — 帳號管理
- `/monitor/monitor.api/v1/*` — 監控資料查詢 (需 JWT)
- `/monitor/admin.api/v1/*` — 管理員操作 (註冊、封鎖、修改)
- `/monitor/collector.api/v1/*` — 外部平台資料推送入口

**資料庫表** (PostgreSQL `monitor` DB)：
- `users` — 監控系統使用者帳號 (admin/dev/qa/ete/pro 角色)
- `service_info` — 外部服務設定與健康檢查端點
- `wallet_ledger` — 玩家上下分交易紀錄
- `user_play_log` — 玩家遊戲輸贏紀錄
- `game` — 遊戲列表 (19 種遊戲)
- `{platform}_game_ratio_stat_{YYYYMM}` — 動態月份 RTP 統計表

**部署方式**：
```bash
# Docker (含 PostgreSQL + Redis)
cd monitorservice/deployments
docker-compose up

# Nginx 反向代理 (生產環境)
# 監聽 9527 → 代理到 localhost:17782/monitor
# 靜態前端 → /var/www/dcc/rtp_monitor
```

### dcc_monitor 監控前端

**技術棧**：Vue 3 + TypeScript + Vite + Tailwind CSS + Pinia（獨立於 dcc_front）

**頁面路由**：
- `/login` — 登入頁
- `/` — 首頁 (Layout)
- `/personal-info` — 個人資訊
- `/member-management` — 會員管理 (管理員)
- `/system-monitor/:platform` — 系統監控主頁面

**主要功能 (SystemMonitor.vue)**：
- 交易紀錄監控 (coin in/out)
- 異常遊戲紀錄
- RTP 監控 (日/週/月，依遊戲類型篩選)
- 每 60 秒自動刷新

## 服務埠號

| 服務 | 埠號 |
|------|------|
| platform | 9986 |
| chatservice | 8896 |
| orderservice | 9988 |
| monitorservice (HTTP API) | 17782 |
| monitorservice (WebSocket) | 17783 |
| monitorservice (Nginx 對外) | 9527 |
| mock-gameserver | 9642 |

## 環境切換

- Go 服務：設定 `web_mode` (debug/docker/qa/ete) 以載入對應的 `config-{mode}.yml`
- 前端：環境檔案位於 `dcc_front/frontend/env/` (`.env`、`.env.dev`、`.env.qa`、`.env.qc`)

## 本地開發環境

### Docker 服務啟動

```bash
# 1. 基礎設施 + 主要後端服務
cd platform/deployments
docker-compose up -d        # PostgreSQL, Redis, pgAdmin, platform, orderservice, chatservice

# 2. Mock Game Server (替代 stone/GameHub)
cd nextplay
docker-compose -f docker-compose.dev.yml up -d   # mock-gameserver (172.99.0.50:9642)

# 3. MonitorService (可選)
cd monitorservice/deployments
docker-compose up -d        # monitorservice + 專屬 PostgreSQL + Redis
```

### 前端開發

```bash
cd dcc_front/frontend
npm install
node node_modules/vite/bin/vite.js --host 127.0.0.1 --config vite.config.manager.js  # 管理後台
node node_modules/vite/bin/vite.js --host 127.0.0.1 --config vite.config.agent.js    # 代理後台
```

### Docker 網路

所有服務使用 `dcc_backend_networks` (172.99.0.0/16)：
- dcc-backend (platform): 172.99.0.10
- dcc-pgsql: 172.99.0.12
- dcc-redis: 172.99.0.13
- mock-gameserver: 172.99.0.50
- monitorservice: 172.99.0.21

### 已知注意事項

- 登入 IP 白名單：agent 表的 `ip_whitelist` 需包含 Docker gateway IP (`172.99.0.*`)
- chatservice 連線：`server_info` 表中 `code='chat'` 的 domain 需設為 `127.0.0.1:8896`
- game server：`server_info` 表中 `code='dev01'` 的 notification 需指向 mock-gameserver
- `GameKillDiveInfoReset`：若 storage 表中此 flag 為 false，platform 會從 DB 載入（可能為空），設為 true 則從 game server 重新取得
