ALTER TABLE "public"."game_users" ADD COLUMN "risk_control_status" varchar(4) NOT NULL DEFAULT '0000';
COMMENT ON COLUMN "public"."game_users"."risk_control_status" IS '風控狀態碼';

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100298, '此接口用來取得指定玩家的處置設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/getgameuserriskcontroltag', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100299, '此接口用來設定指定玩家的處置設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/setgameuserriskcontroltag', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100298, 100299]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;