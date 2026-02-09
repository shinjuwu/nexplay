INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100318, '取得玩家目前餘額', '/api/v1/user/getgameuserbalance', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100318]'::jsonb, false)
WHERE "permission"->'list' @> '100240' AND "permission"->'list' @> '100241';