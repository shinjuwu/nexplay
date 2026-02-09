ALTER TABLE "public"."agent" 
  ADD COLUMN "kill_switch" bool NOT NULL DEFAULT false,
  ADD COLUMN "kill_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_info" varchar(500) NOT NULL DEFAULT '',
  ADD COLUMN "kill_update_time" timestamp(6) NOT NULL DEFAULT now();

COMMENT ON COLUMN "public"."agent"."kill_switch" IS '總代殺放開關(非總代一律false)';
COMMENT ON COLUMN "public"."agent"."kill_ratio" IS '總代基礎殺率';
COMMENT ON COLUMN "public"."agent"."kill_info" IS '總代殺放備註';
COMMENT ON COLUMN "public"."agent"."kill_update_time" IS '總代殺放更新時間';

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type")
VALUES
  (100278, '此接口用來取得當前總代理風控設定資料列表(只有管理可以使用)', '/api/v1/riskcontrol/getagentincomeratiolist', 't', '後台使用', 't', 0),
  (100279, '此接口用來取得指定id總代理風控設定資料(只有管理可以使用)', '/api/v1/riskcontrol/getagentincomeratio', 't', '後台使用', 't', 0),
  (100280, '此接口用來設定指定id總代理風控設定資料(只有管理可以使用)', '/api/v1/riskcontrol/setagentincomeratio', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100278, 100279, 100280]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;