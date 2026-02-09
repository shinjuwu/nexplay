INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100305, '重置密碼', '/api/v1/user/resetpassword', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100305]'::jsonb, false)
WHERE "agent_id" = -1;