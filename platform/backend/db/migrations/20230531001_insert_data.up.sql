INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100301, '取得日結算報表列表', '/api/v1/record/getagentgameratiostatlist', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100302, '取得玩家遊玩紀錄列表', '/api/v1/record/getgameusersstathourlist', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100301, 100302]'::jsonb, false)
WHERE "agent_id" = -1;