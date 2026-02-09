INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100319, '此接口用來取得遊戲基礎設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/getgamesetting', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100320, '此接口用來設定遊戲基礎設定(只有開發商（營運商）可使用)', '/api/v1/riskcontrol/setgamesetting', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100319,100320]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;

ALTER TABLE "public"."agent_game_ratio"
  ADD COLUMN "new_kill_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "active_num" int NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."agent_game_ratio"."new_kill_ratio" IS '新手殺率';
COMMENT ON COLUMN "public"."agent_game_ratio"."active_num" IS '啟動人數';
COMMENT ON TABLE "public"."agent_game_ratio" IS '代理遊戲平台機率設定表';

UPDATE "public"."agent_game_ratio"
  SET "new_kill_ratio" = 0.03;

UPDATE "public"."agent_game_ratio"
  SET "active_num" = 1
  WHERE "game_type" = 1;

INSERT INTO "public"."storage" ("id", "key", "value", "readonly", "create_time", "update_time") VALUES ('ac832d69-fc2a-46e8-9722-889b852ae381', 'GameSettingSupportInfo', '[1001,1002,1003,1004,1005,1006,1007,1008,1009,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,3001,3002,3003,4001]', 'f', '2023-09-05 05:37:53.108477+00', '2023-09-05 05:37:53.108477+00');