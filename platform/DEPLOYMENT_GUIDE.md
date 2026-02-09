# Platform ETE 部署指南

## 目錄
- [系統需求](#系統需求)
- [快速開始](#快速開始)
- [Docker 部署](#docker-部署)
- [生產環境部署](#生產環境部署)
- [環境配置](#環境配置)
- [資料庫設置](#資料庫設置)
- [服務管理](#服務管理)
- [故障排除](#故障排除)

## 系統需求

### 硬體需求
- **CPU**: 2核心以上
- **記憶體**: 4GB以上
- **硬碟**: 20GB以上可用空間
- **網路**: 穩定的網際網路連線

### 軟體需求
- **作業系統**: Ubuntu 20.04+ / CentOS 8+ / RHEL 8+
- **Docker**: 20.10+
- **Docker Compose**: 1.29+
- **Git**: 2.25+
- **Make**: 4.2+

## 快速開始

### 1. 安裝必要軟體

#### Ubuntu/Debian
```bash
# 更新套件列表
sudo apt update

# 安裝基本工具
sudo apt install -y curl wget git make

# 安裝 Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# 安裝 Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/v2.12.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# 啟動 Docker 服務
sudo systemctl start docker
sudo systemctl enable docker

# 將當前用戶加入 docker 群組
sudo usermod -aG docker $USER
newgrp docker
```

#### CentOS/RHEL
```bash
# 安裝基本工具
sudo yum install -y curl wget git make

# 安裝 Docker
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io

# 安裝 Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/v2.12.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# 啟動 Docker 服務
sudo systemctl start docker
sudo systemctl enable docker

# 將當前用戶加入 docker 群組
sudo usermod -aG docker $USER
newgrp docker
```

### 2. 下載專案
```bash
# 克隆專案（假設有權限訪問）
git clone <repository-url> platform-ete
cd platform-ete

# 或者如果已有專案檔案，直接進入目錄
cd /path/to/platform-ete
```

## Docker 部署

### 1. 建立 Docker Volumes
```bash
# 建立持久化存儲卷
docker volume create dcc-pgsql-volumes
docker volume create dcc-redis-volumes
docker volume create dcc-pgadmin-volumes

# 驗證卷已建立
docker volume ls | grep dcc
```

### 2. 環境變數設置（可選）
```bash
# 建立環境變數檔案
cat > .env << EOF
PGADMIN_DEFAULT_EMAIL=admin@example.com
PGADMIN_DEFAULT_PASSWORD=admin123
PGADMIN_PORT=5050
EOF
```

### 3. 啟動服務
```bash
# 進入部署目錄
cd deployments

# 啟動所有服務
docker-compose up -d --build

# 檢查服務狀態
docker-compose ps
```

### 4. 驗證部署
```bash
# 檢查容器運行狀態
docker ps

# 檢查服務日誌
docker-compose logs backend
docker-compose logs pgsql
docker-compose logs redis

# 測試 API 連接
curl http://localhost:9986/api/v1/health
```

### 5. 服務訪問

| 服務 | 地址 | 認證資訊 |
|------|------|----------|
| Backend API | http://localhost:9986 | - |
| Swagger 文檔 | http://localhost:9986/swagger/index.html | - |
| PostgreSQL | localhost:5432 | postgres/123456 |
| Redis | localhost:6379 | 無密碼 |
| pgAdmin4 | http://localhost:5050 | admin/admin |

## 生產環境部署

### 1. 安裝 Go（用於本地編譯）
```bash
# 下載 Go 1.22.3
wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz

# 安裝
sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz

# 設置環境變數
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
source ~/.bashrc

# 驗證安裝
go version
```

### 2. 建立部署目錄
```bash
# 建立目標目錄
sudo mkdir -p /usr/local/bin/backend
sudo mkdir -p /usr/local/bin/backend/db/migrations

# 設置權限
sudo chown -R $USER:$USER /usr/local/bin/backend
```

### 3. 使用 Makefile 部署
```bash
# 確保在專案根目錄
cd /path/to/platform-ete

# 檢查 Makefile 配置
cat Makefile

# 執行完整部署
make all
```

### 4. 配置 systemd 服務
```bash
# 複製服務檔案
sudo cp deployments/systemctl/backend.service /etc/systemd/system/

# 檢查服務檔案內容
sudo cat /etc/systemd/system/backend.service

# 重新載入 systemd
sudo systemctl daemon-reload

# 啟用並啟動服務
sudo systemctl enable backend
sudo systemctl start backend

# 檢查服務狀態
sudo systemctl status backend
```

## 環境配置

### 配置檔案說明
專案支援多種環境配置，根據需要選擇：

| 配置檔案 | 用途 | 說明 |
|----------|------|------|
| config-dev.yml | 開發環境 | 本地開發使用 |
| config-ete-api.yml | ETE API環境 | ETE API 伺服器 |
| config-ete-client.yml | ETE 客戶端環境 | ETE 客戶端連接 |
| config-qa-api.yml | QA API環境 | 品質保證測試 |
| config-qa-backend.yml | QA 後端環境 | QA 後端服務 |
| config-docker.yml | Docker環境 | Docker 容器使用 |

### 修改配置
```bash
# 檢視當前配置
cat backend/config.yml

# 根據環境需要修改配置檔案
# 例如修改資料庫連接設定、Redis 設定等
```

## 資料庫設置

### 1. 初始化資料庫
```bash
# 進入部署目錄
cd deployments

# 執行資料庫初始化
sh dcc.docker.container.init.startup.sh
```

### 2. 手動資料庫遷移
```bash
# 僅執行資料庫遷移
sh dcc.docker.compose.pg.migration.sh

# 或使用 Docker Compose
docker-compose -f docker-compose-deploy.yml up migrate
```

### 3. 資料庫備份與還原
```bash
# 備份資料庫
docker exec dcc-pgsql pg_dump -U postgres dcc_game > backup_$(date +%Y%m%d_%H%M%S).sql

# 還原資料庫
docker exec -i dcc-pgsql psql -U postgres dcc_game < backup_file.sql
```

## 服務管理

### Docker 服務管理
```bash
# 啟動服務
docker-compose up -d

# 停止服務
docker-compose down

# 重啟特定服務
docker-compose restart backend

# 查看日誌
docker-compose logs -f backend

# 進入容器
docker exec -it dcc-backend sh
```

### systemd 服務管理
```bash
# 啟動服務
sudo systemctl start backend

# 停止服務
sudo systemctl stop backend

# 重啟服務
sudo systemctl restart backend

# 查看狀態
sudo systemctl status backend

# 查看日誌
sudo journalctl -u backend -f
```

### Makefile 指令
```bash
# 編譯專案
make build

# 完整部署
make all

# 啟動後端服務
make start_backend

# 停止後端服務
make stop_backend

# 重啟後端服務
make restart_backend

# 清理編譯檔案
make clean
```

## 故障排除

### 常見問題

#### 1. Docker 容器無法啟動
```bash
# 檢查 Docker 狀態
sudo systemctl status docker

# 檢查容器日誌
docker-compose logs

# 重新建立容器
docker-compose down
docker-compose up -d --build
```

#### 2. 資料庫連接失敗
```bash
# 檢查 PostgreSQL 容器
docker exec dcc-pgsql psql -U postgres -c "SELECT version();"

# 檢查網路連接
docker network ls
docker network inspect dcc_backend_networks
```

#### 3. 編譯失敗
```bash
# 檢查 Go 環境
go version
go env

# 清理 Go 模組快取
go clean -modcache

# 重新下載依賴
go mod download
```

#### 4. 權限問題
```bash
# 檢查檔案權限
ls -la /usr/local/bin/backend/

# 修復權限
sudo chown -R $USER:$USER /usr/local/bin/backend/
sudo chmod +x /usr/local/bin/backend/dcc_backend
```

### 日誌位置
- **Docker 日誌**: `docker-compose logs`
- **systemd 日誌**: `/var/log/syslog` 或 `journalctl`
- **應用程式日誌**: 根據配置檔案中的設定

### 監控指令
```bash
# 系統資源監控
top
htop
df -h
free -h

# Docker 資源使用
docker stats

# 網路連接檢查
netstat -tlnp | grep :9986
ss -tlnp | grep :9986
```

## 安全建議

1. **變更預設密碼**: 修改資料庫和管理界面的預設密碼
2. **防火牆設置**: 僅開放必要的端口
3. **SSL/TLS**: 在生產環境中啟用 HTTPS
4. **定期備份**: 建立自動化備份策略
5. **日誌監控**: 設置日誌監控和告警

## 效能調優

1. **資料庫調優**: 根據負載調整 PostgreSQL 配置
2. **Redis 記憶體**: 設置適當的 Redis 記憶體限制
3. **連接池**: 調整資料庫連接池大小
4. **負載均衡**: 在高負載環境中使用負載均衡器

---

## 聯絡支援

如有部署問題，請提供以下資訊：
- 作業系統版本
- Docker 版本
- 錯誤日誌
- 部署步驟

**版本**: 1.0.0  
**更新日期**: 2024-11-04