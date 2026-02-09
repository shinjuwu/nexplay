INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
    (100252, '取得代理分數紀錄列表', '/api/v1/record/getagentwalletledgerlist', true, '後台使用', true);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100252]'::jsonb, false)
WHERE "agent_id" = -1;
