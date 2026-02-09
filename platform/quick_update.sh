#!/bin/bash

# 快速更新腳本 - 直接執行SQL命令
echo "正在更新server_info配置..."

# 更新dev01配置
docker exec dcc-pgsql psql -U postgres -d dcc_game -c "
UPDATE server_info 
SET 
    ip = '34.81.187.182',
    addresses = '{\"notification\": \"http://34.81.187.182:9642/\"}',
    update_time = now()
WHERE code = 'dev01';
"

# 確保配置已啟用
docker exec dcc-pgsql psql -U postgres -d dcc_game -c "
UPDATE server_info 
SET is_enabled = true 
WHERE code = 'dev01';
"

# 顯示更新結果
echo "更新後的配置:"
docker exec dcc-pgsql psql -U postgres -d dcc_game -c "
SELECT code, ip, addresses, is_enabled 
FROM server_info 
WHERE code = 'dev01';
"

echo "配置更新完成! 請重啟後端服務:"
echo "cd deployments && docker-compose restart backend"