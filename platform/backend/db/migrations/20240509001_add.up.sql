INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100331,'取得玩家帳變紀錄列表(only 管理可用)','/api/v1/record/getusercreditloglist','t','後台使用','t',0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100331]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;