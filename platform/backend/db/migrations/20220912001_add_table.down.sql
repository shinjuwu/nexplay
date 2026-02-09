CREATE TABLE "public"."agent_group" (
  "id" serial4,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "mode" int4 NOT NULL DEFAULT 0,
  "code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_count" int4 NOT NULL DEFAULT 0,
  "commission" int4 NOT NULL DEFAULT 0,
  "info" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" int4 NOT NULL DEFAULT 1,
  "expire_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "disable_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_agent_group_code" UNIQUE ("code"),
  CONSTRAINT "agent_group_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."agent_group"."id" IS '集團編號';
COMMENT ON COLUMN "public"."agent_group"."name" IS '集團名稱';
COMMENT ON COLUMN "public"."agent_group"."mode" IS '集團模式(1.平台自營,2.系統商)';
COMMENT ON COLUMN "public"."agent_group"."code" IS '集團識別編碼-隨機生成不重複唯一碼';
COMMENT ON COLUMN "public"."agent_group"."agent_count" IS '旗下代理數 平台自營 固定為1';
COMMENT ON COLUMN "public"."agent_group"."commission" IS '分成(萬分之n)';
COMMENT ON COLUMN "public"."agent_group"."info" IS '廠商註記';
COMMENT ON COLUMN "public"."agent_group"."is_enabled" IS 'BOSS啟用中(1:啟用,0:關閉)';
COMMENT ON COLUMN "public"."agent_group"."expire_time" IS '合約到期日';
COMMENT ON COLUMN "public"."agent_group"."disable_time" IS '服務停止時間';
COMMENT ON COLUMN "public"."agent_group"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "public"."agent_group"."create_time" IS '創建時間';
COMMENT ON TABLE "public"."agent_group" IS '總代理商戶設置/集團商務設置';

-- ----------------------------
-- Records of agent_group
-- ----------------------------
INSERT INTO "public"."agent_group" VALUES (1, 'DCC GAMING', 1, '12345678', 1, 0, '', 1, '1970-01-01 00:00:00', '1970-01-01 00:00:00', '2022-08-10 11:55:47.161639', '2022-08-10 11:55:47.161639');



ALTER TABLE "public"."agent" 
  DROP CONSTRAINT "uni_agent_1",
  ADD COLUMN "agent_group_code" int4 NOT NULL DEFAULT 1,
  ADD CONSTRAINT "uni_agent_1" UNIQUE ("code", "agent_group_code"),
  ALTER COLUMN "top_agent_id" TYPE int2 USING "top_agent_id"::int2,
  ALTER COLUMN "top_agent_id" DROP NOT NULL,
  DROP COLUMN "cooperation",
  DROP COLUMN "coin_limit",
  DROP COLUMN "coin_use",
  DROP COLUMN "coin_supply_setting";

UPDATE "public"."agent" SET "top_agent_id" = NULL;

ALTER TABLE "public"."user_play_log_baccarat" 
  DROP CONSTRAINT "play_log_baccarat_pkey",
  ADD COLUMN "agent_group_id" int4 NOT NULL DEFAULT 1,
  ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "agent_group_id");