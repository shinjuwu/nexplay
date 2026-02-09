INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
    (100245, '依照指定報表類型取得指定時間區段的資料', '/api/v1/cal/getperformancereport', true, '後台使用', true),
    (100246, '取得代理後台IP資訊列表', '/api/v1/agent/getagentipwhitelistlist', true, '後台使用', true),
    (100247, '取得代理後台IP資訊', '/api/v1/agent/getagentipwhitelist', true, '後台使用', true),
    (100248, '設置代理後台IP資訊', '/api/v1/agent/setagentipwhitelist', true, '後台使用', true);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100243,100244]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100245,100246,100247,100248]'::jsonb, false)
WHERE "agent_id" = -1;

UPDATE "public"."agent"
SET "ip_whitelist" = '[{"info":"","creator":"admin","ip_address":"127.0.0.1","create_time":1666592735}]';
