CREATE TABLE "public"."game_ratio" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "base_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "up_ratio_limit" numeric(20,4) NOT NULL,
  "down_ratio_limit" numeric(20,4) NOT NULL,
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "game_ratio_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."game_ratio"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."game_ratio"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."game_ratio"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."game_ratio"."base_ratio" IS '基礎殺率';
COMMENT ON COLUMN "public"."game_ratio"."up_ratio_limit" IS '上水線';
COMMENT ON COLUMN "public"."game_ratio"."down_ratio_limit" IS '下水線';
COMMENT ON COLUMN "public"."game_ratio"."update_time" IS '最後更新間';
COMMENT ON TABLE "public"."game_ratio" IS '遊戲平台機率設定表';


DELETE FROM "public"."permission_list" WHERE "feature_code" = 100260;
DELETE FROM "public"."permission_list" WHERE "feature_code" = 100261;
DELETE FROM "public"."permission_list" WHERE "feature_code" = 100262;

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required") VALUES (100260,'此接口用來取得當前平台機率設定列表(只有管理可以使用)','/api/v1/riskcontrol/getincomeratiolist','t','後台使用','t');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required") VALUES (100261,'此接口用來取得指定id平台機率設定資料(只有管理可以使用)','/api/v1/riskcontrol/getincomeratio','t','後台使用','t');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required") VALUES (100262,'此接口用來設定指定id平台機率設定資料(只有管理可以使用)','/api/v1/riskcontrol/setincomeratio','t','後台使用','t');

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100260,100261,100262]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;
