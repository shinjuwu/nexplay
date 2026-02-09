INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
    (100250, '取得代理錢包餘額列表', '/api/v1/agent/getagentwalletlist', true, '後台使用', true),
    (100251, '設置代理錢包餘額', '/api/v1/agent/setagentwallet', true, '後台使用', true);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100250, 100251]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" IN (1, 2);
