#!/bin/bash

# ============================================
# Server Info 更新腳本
# 用途: 執行數據庫更新以修改遊戲服務器IP配置
# 日期: 2024-11-04
# ============================================

# 腳本配置
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SQL_FILE="${SCRIPT_DIR}/update_server_info.sql"
BACKUP_FILE="${SCRIPT_DIR}/server_info_backup_$(date +%Y%m%d_%H%M%S).sql"

# 顏色定義
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日誌函數
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 檢查Docker容器是否運行
check_postgres_container() {
    log_info "檢查PostgreSQL容器狀態..."
    if ! docker ps | grep -q "dcc-pgsql"; then
        log_error "PostgreSQL容器未運行，請先啟動Docker服務"
        log_info "執行: cd deployments && docker-compose up -d"
        exit 1
    fi
    log_success "PostgreSQL容器正在運行"
}

# 創建數據庫備份
create_backup() {
    log_info "創建server_info表備份..."
    docker exec dcc-pgsql pg_dump -U postgres -d dcc_game -t server_info > "${BACKUP_FILE}" 2>/dev/null
    if [ $? -eq 0 ]; then
        log_success "備份已創建: ${BACKUP_FILE}"
    else
        log_warning "備份創建失敗，但繼續執行更新"
    fi
}

# 執行SQL更新
execute_update() {
    log_info "執行數據庫更新..."
    if [ ! -f "${SQL_FILE}" ]; then
        log_error "SQL文件不存在: ${SQL_FILE}"
        exit 1
    fi
    
    # 執行SQL文件
    docker exec -i dcc-pgsql psql -U postgres -d dcc_game < "${SQL_FILE}"
    
    if [ $? -eq 0 ]; then
        log_success "數據庫更新完成"
    else
        log_error "數據庫更新失敗"
        exit 1
    fi
}

# 驗證更新結果
verify_update() {
    log_info "驗證更新結果..."
    
    # 檢查dev01配置是否更新成功
    RESULT=$(docker exec dcc-pgsql psql -U postgres -d dcc_game -t -c "SELECT ip FROM server_info WHERE code = 'dev01';")
    
    if echo "$RESULT" | grep -q "34.81.187.182"; then
        log_success "IP地址更新成功: 34.81.187.182"
    else
        log_error "IP地址更新失敗"
        exit 1
    fi
}

# 重啟後端服務
restart_backend() {
    log_info "重啟後端服務以應用新配置..."
    cd "${SCRIPT_DIR}/deployments" 2>/dev/null || cd "${SCRIPT_DIR}"
    
    if [ -f "docker-compose.yml" ]; then
        docker-compose restart backend
        if [ $? -eq 0 ]; then
            log_success "後端服務重啟完成"
        else
            log_warning "後端服務重啟失敗，請手動重啟"
        fi
    else
        log_warning "找不到docker-compose.yml，請手動重啟後端服務"
    fi
}

# 主函數
main() {
    echo "========================================"
    echo "Server Info 更新腳本"
    echo "將遊戲服務器IP更新為: 34.81.187.182"
    echo "========================================"
    
    # 確認執行
    read -p "是否確認執行更新? (y/N): " confirm
    if [[ ! $confirm =~ ^[Yy]$ ]]; then
        log_info "操作已取消"
        exit 0
    fi
    
    # 執行步驟
    check_postgres_container
    create_backup
    execute_update
    verify_update
    
    # 詢問是否重啟服務
    read -p "是否重啟後端服務以應用新配置? (y/N): " restart_confirm
    if [[ $restart_confirm =~ ^[Yy]$ ]]; then
        restart_backend
    fi
    
    echo "========================================"
    log_success "所有操作完成!"
    echo "新的遊戲服務器地址: http://34.81.187.182:9642/"
    echo "備份文件位置: ${BACKUP_FILE}"
    echo "========================================"
}

# 執行主函數
main