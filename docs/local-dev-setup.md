# 本地前端開發環境搭建指南

本文件說明如何在本機搭建 NextPlay (DCC) 前端開發環境，包含所需的後端服務。

> **快速開始**：若不想逐步操作，可直接使用一鍵啟動腳本，詳見 [一鍵啟動腳本](#quick-start)。

---

## 目錄

1. [前置需求](#1-前置需求)
2. [架構總覽](#2-架構總覽)
3. [Step 1: 建立 Docker Volumes](#3-step-1-建立-docker-volumes)
4. [Step 2: 啟動基礎設施](#4-step-2-啟動基礎設施)
5. [Step 3: 啟動 Mock Game Server (選用)](#5-step-3-啟動-mock-game-server-選用)
6. [Step 4: 啟動其他微服務 (選用)](#6-step-4-啟動其他微服務-選用)
7. [Step 5: 安裝前端依賴 & 啟動開發伺服器](#7-step-5-安裝前端依賴--啟動開發伺服器)
8. [存取服務](#8-存取服務)
9. [環境設定說明](#9-環境設定說明)
10. [資料庫管理](#10-資料庫管理)
11. [常見問題排除](#11-常見問題排除)
12. [服務埠號速查表](#12-服務埠號速查表)

---

## 1. 前置需求

請確認本機已安裝以下工具：

| 工具 | 最低版本 | 說明 |
|------|---------|------|
| **Docker Desktop** | 4.x+ | 執行後端服務及資料庫 |
| **Node.js** | 16.x+ | 前端建置環境 |
| **npm** | 8.x+ | 套件管理 (隨 Node.js 安裝) |
| **Git** | 2.x+ | 版本控制與 submodule |

> **Windows 使用者**：確認 Docker Desktop 已啟用且正在運行，建議使用 WSL 2 backend。

Clone 專案後，記得初始化 git submodule：

```bash
git clone <repository-url> nexplay
cd nexplay
git submodule update --init --recursive
```

---

## 2. 架構總覽

本地開發時，前端透過 Vite dev server 的 proxy 機制將 API 請求轉發到後端服務，整體架構如下：

```
┌─────────────────────┐
│   瀏覽器 (Browser)    │
│  http://localhost:9985│
└─────────┬───────────┘
          │
┌─────────▼───────────┐
│  Vite Dev Server     │
│  (port 9985)         │
│  ┌────────────────┐  │
│  │  /api/* proxy  │──┼──────────┐
│  └────────────────┘  │          │
└──────────────────────┘          │
                                  │
          ┌───────────────────────▼──────────────────────┐
          │           Docker Network (172.99.0.0/16)       │
          │                                                │
          │  ┌──────────────┐  ┌──────────┐  ┌──────────┐ │
          │  │  Platform     │  │PostgreSQL│  │  Redis   │ │
          │  │  Backend      │  │  :5432   │  │  :6379   │ │
          │  │  :9986        │  │ 172.99.  │  │ 172.99.  │ │
          │  │  172.99.0.10  │  │  0.12    │  │  0.11    │ │
          │  └──────────────┘  └──────────┘  └──────────┘ │
          │                                                │
          │  ┌──────────────┐  ┌──────────┐  ┌──────────┐ │
          │  │ OrderService  │  │ChatService│ │  pgAdmin │ │
          │  │  :9988 (選用) │  │ :8896    │  │  :5050   │ │
          │  │  172.99.0.99  │  │(選用)    │  │ 172.99.  │ │
          │  └──────────────┘  │172.99.0.21│  │  0.14    │ │
          │                    └──────────┘  └──────────┘ │
          └────────────────────────────────────────────────┘
```

**核心流程**：瀏覽器 → Vite Dev Server (localhost:9985) → proxy `/api/*` → Platform Backend (localhost:9986) → PostgreSQL / Redis

---

## 3. Step 1: 建立 Docker Volumes

首次設定時，需手動建立 `docker-compose.yml` 中宣告為 `external` 的 Docker volumes：

```bash
docker volume create dcc-pgsql-volumes
docker volume create dcc-redis-volumes
docker volume create dcc-pgadmin-volumes
```

> 這些 volumes 用於持久化資料庫和 Redis 資料，僅需執行一次。

驗證：

```bash
docker volume ls | grep dcc
```

應看到三個 volume 已建立。

---

## 4. Step 2: 啟動基礎設施

進入 `platform/deployments/` 目錄啟動主要後端服務：

```bash
cd platform/deployments
docker-compose up -d
```

此指令會啟動以下容器：

| 容器名稱 | 映像檔 | 說明 |
|----------|--------|------|
| `dcc-backend` | `dcc/backend:1.0.0` | Platform 主後端 API |
| `dcc-pgsql` | `postgres:14.5` | PostgreSQL 資料庫 |
| `dcc-redis` | `redis:7` | Redis 快取 |
| `dcc-pgadmin` | `dpage/pgadmin4` | 資料庫管理介面 |

> **首次啟動**：`dcc-backend` 映像檔需要建置，可能需要數分鐘。PostgreSQL 啟動時會自動執行 `pgsql-dump/` 目錄下的初始化 SQL。

驗證容器狀態：

```bash
docker-compose ps
```

確認所有容器狀態為 `Up`。若 `dcc-backend` 異常，檢查日誌：

```bash
docker logs dcc-backend
```

---

## 5. Step 3: 啟動 Mock Game Server (選用)

若需要測試與遊戲伺服器的互動，可啟動 Mock Game Server：

```bash
# 回到專案根目錄
cd <project-root>
docker-compose -f docker-compose.dev.yml up -d
```

> **注意**：此指令需要在 `dcc_backend_networks` 網路已建立之後執行 (即 Step 2 完成後)，因為 `docker-compose.dev.yml` 宣告該網路為 `external`。

| 容器名稱 | 說明 | 位址 |
|----------|------|------|
| `mock-gameserver` | Python 模擬遊戲伺服器 | 172.99.0.50:9642 |

---

## 6. Step 4: 啟動其他微服務 (選用)

根據開發需求，可選擇性啟動以下服務。這些服務各有獨立的 `docker-compose.yml`，且共用同一個 Docker 網路。

### ChatService

```bash
cd chatservice/deployments
docker-compose up -d
```

| 容器名稱 | 說明 | 位址 |
|----------|------|------|
| `chatservice` | WebSocket 聊天服務 | 172.99.0.21:8896 |

> 首次啟動需建置映像檔。

### OrderService

```bash
cd orderservice/deployments
docker-compose up -d
```

| 容器名稱 | 說明 | 位址 |
|----------|------|------|
| `dcc-orderservice` | 訂單/遊戲紀錄服務 | 172.99.0.99:9988 |

> OrderService 依賴 Platform 的 PostgreSQL 和 Redis，請確認 Step 2 已完成。

---

## 7. Step 5: 安裝前端依賴 & 啟動開發伺服器

```bash
cd dcc_front/frontend
npm install
```

啟動開發伺服器 (擇一)：

```bash
# 管理後台 (Manager)
npm run manager-dev

# 代理端後台 (Agent)
npm run agent-dev
```

啟動後瀏覽器會自動開啟，預設位址為 `http://localhost:9985`。

---

## 8. 存取服務

開發環境啟動後，可透過以下 URL 存取各服務：

| 服務 | URL | 說明 |
|------|-----|------|
| **Manager 管理後台** | http://localhost:9985 | Vite dev server (manager-dev) |
| **Agent 代理端後台** | http://localhost:9985 | Vite dev server (agent-dev) |
| **Platform API** | http://localhost:9986 | 後端 REST API |
| **pgAdmin** | http://localhost:5050 | 資料庫管理介面 |
| **PostgreSQL** | localhost:5432 | 可用 DBeaver 等工具直連 |
| **Redis** | localhost:6379 | 可用 RedisInsight 等工具直連 |

> Manager 和 Agent 共用同一個 port (9985)，但不能同時運行，需擇一啟動。

---

## 9. 環境設定說明

### 環境變數檔案

環境變數檔案位於 `dcc_front/frontend/env/`：

| 檔案 | 用途 |
|------|------|
| `.env` | **預設/本地開發** — 定義 proxy 設定 |
| `.env.dev` | dev 環境建置 |
| `.env.qa` | QA 環境建置 |
| `.env.qc` | QC 環境建置 |

本地開發時使用 `.env` 中的設定：

```properties
VITE_SERVER_CLI_PORT=9985              # Vite dev server 埠號
VITE_SERVER_PROXY_API_PATHNAME=/api    # 需要 proxy 的 API 路徑前綴
VITE_SERVER_PROXY_CLI_ORIGIN=http://127.0.0.1  # 後端 host
VITE_SERVER_PROXY_API_PORT=9986        # 後端 API 埠號
```

### Vite Proxy 機制

Vite 配置檔 (`vite.config.js`) 中設定了 proxy，將前端對 `/api/*` 的請求轉發到後端：

```
瀏覽器請求: http://localhost:9985/api/v1/login
       ↓ Vite proxy
後端接收: http://127.0.0.1:9986/api/v1/login
```

Proxy 同時會在請求中注入 `X-Dcc-Header`，用以區分 Manager 和 Agent：

- **Manager**: `X-Dcc-Header: M7NJpSXxh`
- **Agent**: `X-Dcc-Header: aikDg57I1`

這些值分別定義在 `vite.config.manager.js` 和 `vite.config.agent.js` 中。

---

## 10. 資料庫管理

### pgAdmin 登入

| 欄位 | 值 |
|------|---|
| URL | http://localhost:5050 |
| Email | `pgadmin4@pgadmin.org` |
| Password | `admin` |

### 連線 PostgreSQL

在 pgAdmin 中新增 Server 連線：

| 欄位 | 值 |
|------|---|
| Host | `dcc-pgsql` (容器內) 或 `localhost` (本機) |
| Port | `5432` |
| Username | `postgres` |
| Password | `123456` |
| Database | `dcc_game` |

### 資料庫初始化

PostgreSQL 首次啟動時，會自動執行 `platform/deployments/pgsql-dump/` 目錄下的 SQL 檔案。Platform backend 啟動時也會執行 DB migration (`load_db_migration: true`)，SQL 遷移檔案位於 `platform/backend/db/migrations/`。

初始化 SQL 腳本參考 (位於 `platform/backend/db/init/`)：
- `init_database.sql` — 建立資料庫
- `init_table.sql` — 建立資料表
- `init_insert_data.sql` — 插入初始資料

---

## 11. 常見問題排除

### Agent 登入失敗 — IP 白名單

**現象**：Agent 端登入時回傳 IP 不在白名單內。

**原因**：`agent` 資料表的 `ip_whitelist` 欄位限制了可登入的 IP。

**解法**：在資料庫中更新該 agent 的 IP 白名單，加入 Docker gateway IP：

```sql
UPDATE agent SET ip_whitelist = '172.99.0.*,127.0.0.1' WHERE ...;
```

### ChatService 連線失敗

**現象**：前端無法連線 WebSocket 聊天服務。

**解法**：確認 `server_info` 資料表中 `code='chat'` 的 domain 設定正確：

```sql
UPDATE server_info SET domain = '127.0.0.1:8896' WHERE code = 'chat';
```

### Game Server 連線失敗

**現象**：遊戲相關功能無法正常運作。

**解法**：確認 `server_info` 資料表中 `code='dev01'` 的 notification URL 指向 mock-gameserver：

```sql
UPDATE server_info SET notification = 'http://172.99.0.50:9642/...' WHERE code = 'dev01';
```

### Port 衝突

**現象**：`docker-compose up` 失敗，提示 port 已被佔用。

**解法**：

```bash
# 查看佔用 port 的程序 (Windows)
netstat -ano | findstr :5432

# 查看佔用 port 的程序 (macOS/Linux)
lsof -i :5432
```

關閉佔用程序後重新啟動，或修改 `docker-compose.yml` 中的 port 映射。

### dcc-backend 啟動失敗

**現象**：`dcc-backend` 容器反覆重啟或立即退出。

**排查步驟**：

1. 檢查日誌：`docker logs dcc-backend`
2. 確認 PostgreSQL 已完全啟動 (首次初始化需要時間)
3. 確認 Redis 正在運行：`docker logs dcc-redis`
4. 嘗試重啟：`docker-compose restart backend`

### Docker Network 衝突

**現象**：啟動 `docker-compose.dev.yml` 或其他服務的 compose 時，提示網路已存在或子網衝突。

**解法**：確保 Step 2 的 `platform/deployments/docker-compose.yml` 先啟動，因為它會建立 `dcc_backend_networks`。其他 compose 檔案宣告此網路為 `external`，需依賴已存在的網路。

### GameKillDiveInfoReset

若 `storage` 資料表中 `GameKillDiveInfoReset` flag 為 `false`，platform 會從 DB 載入資料 (可能為空)。如需從 game server 重新取得，將此 flag 設為 `true`。

---

## 12. 服務埠號速查表

| 服務 | 埠號 | Docker IP | 容器名稱 |
|------|------|-----------|----------|
| Platform Backend | 9986 | 172.99.0.10 | dcc-backend |
| PostgreSQL | 5432 | 172.99.0.12 | dcc-pgsql |
| Redis | 6379 | 172.99.0.11 | dcc-redis |
| pgAdmin | 5050 | 172.99.0.14 | dcc-pgadmin |
| ChatService | 8896 | 172.99.0.21 | chatservice |
| OrderService | 9988 | 172.99.0.99 | dcc-orderservice |
| Mock Game Server | 9642 | 172.99.0.50 | mock-gameserver |
| Vite Dev Server | 9985 | — | — (本機) |

---

<a id="quick-start"></a>

## 一鍵啟動腳本 (Windows)

專案提供了 Windows batch 腳本，可一鍵完成所有啟動步驟。

<a id="prerequisites"></a>

### 前置配置 (首次必做)

在執行腳本之前，請確認以下項目皆已完成：

#### 1. 安裝必要工具

確認 Docker Desktop、Node.js (16+)、npm、Git 已安裝。

```bash
docker --version
node --version
npm --version
git --version
```

#### 2. 初始化 Git Submodule

Platform 後端的建置依賴 `definition` submodule。**若未初始化，Docker build 會失敗。**

```bash
cd <project-root>
git submodule update --init --recursive
```

> 此 submodule 來源為內部 GitLab (`gitlab.int.dayongtek.com`)，請確認你的 SSH key 或 Git 憑證已設定，且有該倉庫的存取權限。

驗證：確認 `platform/definition/` 目錄下有 `.go` 檔案。

#### 3. 準備資料庫初始化 SQL

PostgreSQL 首次啟動時會執行 `platform/deployments/pgsql-dump/` 目錄下的 SQL 檔案。請確認該目錄內包含所需的初始化 SQL dump。

```bash
# 確認目錄內有 SQL 檔案
ls platform/deployments/pgsql-dump/
```

> 若目錄為空，請向團隊成員取得最新的 SQL dump 檔案並放入此目錄。Platform backend 啟動時會自動執行 DB migration，但 migration 可能需要基礎資料存在才能正確運行。

#### 4. 確認 Docker Desktop 已啟動

腳本會自動檢查 Docker 狀態，但請在執行前確認 Docker Desktop 已開啟且狀態正常 (系統匣圖示為綠色)。

#### 前置配置檢查清單

| 項目 | 檢查方式 | 狀態 |
|------|---------|------|
| Docker Desktop 已安裝且運行 | `docker info` 無錯誤 | ☐ |
| Node.js 16+ 已安裝 | `node --version` | ☐ |
| Git submodule 已初始化 | `platform/definition/` 有 `.go` 檔案 | ☐ |
| pgsql-dump 已備妥 | `platform/deployments/pgsql-dump/` 有 SQL 檔案 | ☐ |
| GitLab 存取權限 | 可 `git pull` submodule | ☐ |

---

### 啟動

```bash
# 預設：啟動基礎設施 + Manager 前端
scripts\dev-start.bat

# 啟動基礎設施 + Agent 前端
scripts\dev-start.bat --agent

# 啟動所有服務 (含 mock-gameserver, chatservice, orderservice) + Manager 前端
scripts\dev-start.bat --all

# 只啟動後端服務，不啟動前端
scripts\dev-start.bat --no-front

# 自由組合
scripts\dev-start.bat --mock --chat
```

腳本會自動完成以下步驟：
1. 前置檢查 (Docker、Git submodule、Node.js)
2. 建立 Docker Volumes (若不存在)
3. 啟動基礎設施 (Platform + PostgreSQL + Redis + pgAdmin)
4. 啟動選用服務 (依參數決定)
5. 等待 Backend 就緒
6. 安裝前端依賴 (首次) 並啟動 Vite dev server

> 腳本會自動檢查 Docker、Git submodule、Node.js 是否就緒，若不符合會提示錯誤訊息。但 **pgsql-dump SQL 檔案**和 **GitLab 存取權限**需事先手動準備，請參考上方[前置配置](#prerequisites)。

### 停止

```bash
# 停止所有 Docker 服務
scripts\dev-stop.bat
```
