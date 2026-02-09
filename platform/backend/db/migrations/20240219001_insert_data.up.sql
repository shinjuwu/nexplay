INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100322,'此接口用來取得遊戲即時殺率資訊(只有開發商（營運商）可使用)','/api/v1/riskcontrol/getrealtimegameratio','t','後台使用','t',0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100322]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;