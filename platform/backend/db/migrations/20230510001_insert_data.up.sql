INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100296, '此接口用來取得自動風控設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/getautoriskcontrolsetting', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100297, '此接口用來設定自動風控設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/setautoriskcontrolsetting', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100296, 100297]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;