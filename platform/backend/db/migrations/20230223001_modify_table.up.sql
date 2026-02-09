UPDATE "public"."job_scheduler" SET "is_enabled" = 't' WHERE "id" = 'c23621fc-a492-44d3-9873-4e4be72292ef';

ALTER TABLE "public"."game_users" 
  ADD COLUMN "kill_dive_value" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "is_risk" bool NOT NULL DEFAULT false,
  ADD COLUMN "custom_status" varchar(8) NOT NULL DEFAULT '00000000';

COMMENT ON COLUMN "public"."game_users"."is_risk" IS '是否為高風險用戶';
COMMENT ON COLUMN "public"."game_users"."custom_status" IS '自定義狀態碼';
COMMENT ON COLUMN "public"."game_users"."kill_dive_value" IS '殺放定點數值設定';

CREATE TABLE "public"."agent_custom_tag_info" (
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) NOT NULL,
  "custom_tag_info" jsonb NOT NULL DEFAULT '{}',
  "update_time" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("agent_id")
)
;

COMMENT ON COLUMN "public"."agent_custom_tag_info"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."agent_custom_tag_info"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."agent_custom_tag_info"."custom_tag_info" IS '自定義標籤資訊(array in json)';
COMMENT ON TABLE "public"."agent_custom_tag_info" IS '代理自定義玩家標籤';

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type")
VALUES
  (100281, '此接口用來取得指定id總代理風控設定資料(只有代理可以用)', '/api/v1/riskcontrol/getagentcustomtagsettinglist', 't', '後台使用', 't', 0),
  (100282, '此接口用來設定指定id總代理風控設定資料(只有代理可以用)', '/api/v1/riskcontrol/setagentcustomtagsettinglist', 't', '後台使用', 't', 2),
  (100283, '取得玩家標示資料列表', '/api/v1/riskcontrol/getgameuserscustomtaglist', 't', '後台使用', 't', 0),
  (100284, '取得標示玩家資料(只有代理可以用)', '/api/v1/riskcontrol/getgameuserscustomtag', 't', '後台使用', 't', 0),
  (100285, '設定標示玩家資料(只有代理可以用)', '/api/v1/riskcontrol/setgameuserscustomtag', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100281, 100282, 100284, 100285]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" > 1;

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100283]'::jsonb, false)
WHERE "agent_id" = -1;
