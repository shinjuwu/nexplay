INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100324, '此接口用來批次設定殺數設定資料(只有管理可以使用)', '/api/v1/riskcontrol/setincomeratios', 't', '後台使用', 't', 2);                  

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100324]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;