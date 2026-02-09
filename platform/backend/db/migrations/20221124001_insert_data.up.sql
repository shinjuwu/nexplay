INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
    (100253, '取得玩家錢包餘額列表', '/api/v1/user/getgameuserwalletlist', true, '後台使用', true),
    (100254, '設置玩家錢包餘額', '/api/v1/user/setgameuserwallet', true, '後台使用', true),
    (100255, '更新玩家分數紀錄狀態', '/api/v1/record/confirmwalletledger', true, '後台使用', true);


UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100253, 100254, 100255]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" IN (1, 2);