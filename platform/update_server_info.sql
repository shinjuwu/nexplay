-- ============================================
-- Server Info 更新腳本
-- 用途: 更新遊戲服務器IP地址配置
-- 日期: 2024-11-04
-- 說明: 將舊的IP地址更新為新的34.81.187.182
-- ============================================

-- 顯示更新前的配置
SELECT 
    code, 
    ip, 
    addresses, 
    is_enabled,
    update_time
FROM server_info 
WHERE code IN ('dev01', 'qa01', 'chat', 'api', 'maintain','monitor')
ORDER BY code;

-- 備份當前配置 (可選，建議執行)
-- CREATE TABLE server_info_backup_20241104 AS 
-- SELECT * FROM server_info WHERE code = 'dev';

-- 更新 dev01 配置 (主要的遊戲服務器配置)
UPDATE server_info 
SET 
    ip = '34.81.187.182',
    addresses = '{"notification": "http://34.81.187.182:9642/"}',
    update_time = now()
WHERE code = 'dev01';

-- 如果需要同時更新其他相關配置，取消下面的註釋
-- 更新 qa01 配置 (如果QA環境也需要指向同一個服務器)
-- UPDATE server_info 
-- SET 
--     ip = '34.81.187.182',
--     addresses = '{"notification": "http://34.81.187.182:9642/"}',
--     update_time = now()
-- WHERE code = 'qa01';

-- 確保配置已啟用
UPDATE server_info 
SET is_enabled = true 
WHERE code = 'dev01';

-- 顯示更新後的配置
SELECT 
    code, 
    ip, 
    addresses, 
    is_enabled,
    update_time
FROM server_info 
WHERE code IN ('dev01', 'qa01', 'chat', 'api', 'maintain','monitor')
ORDER BY code;

-- 驗證受影響的遊戲
SELECT 
    g.id,
    g.name,
    g.code,
    g.server_info_code,
    s.ip as server_ip,
    s.addresses,
    g.state
FROM game g
LEFT JOIN server_info s ON g.server_info_code = s.code
WHERE g.server_info_code = 'dev01'
ORDER BY g.id;

-- 腳本完成提示
SELECT 'Server info update completed successfully!' as status;