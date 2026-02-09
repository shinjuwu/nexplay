INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100323,'此接口用來取得玩家目前遊戲局數狀態','/api/v1/user/getgameusernewbie','t','後台使用','t',0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100323]'::jsonb, false)
WHERE "permission"->'list' @> '100240' AND "permission"->'list' @> '100241' AND "permission"->'list' @> '100318';