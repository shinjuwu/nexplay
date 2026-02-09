INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
    (100249, '取得代理權限群組', '/api/v1/agent/getagentpermission', true, '後台使用', true);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100249]'::jsonb, false)
WHERE "agent_id" = -1;