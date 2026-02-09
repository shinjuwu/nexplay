
-- DROP TABLE
DROP TABLE IF EXISTS "public"."agent_group";

-- UPDATE
UPDATE "public"."agent" SET "top_agent_id" = -1;

-- ALTER TABLE
ALTER TABLE "public"."agent" 
  DROP CONSTRAINT "uni_agent_1",
  DROP COLUMN "agent_group_code",
  ADD CONSTRAINT "uni_agent_1" UNIQUE ("code"),
  ALTER COLUMN "top_agent_id" TYPE int4 USING "top_agent_id"::int4,
  ALTER COLUMN "top_agent_id" SET NOT NULL,
  ALTER COLUMN "top_agent_id" SET DEFAULT -1,
  ADD COLUMN "cooperation" int2 NOT NULL DEFAULT 1,
  ADD COLUMN "coin_limit" decimal(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "coin_use" decimal(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "coin_supply_setting" jsonb NOT NULL DEFAULT '{}'::jsonb,
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT '',
  ADD COLUMN "aes_key" varchar(16) NOT NULL DEFAULT ''::character varying,
  ALTER COLUMN "md5_key" TYPE varchar(16) COLLATE "pg_catalog"."default";

COMMENT ON COLUMN "public"."agent"."cooperation" IS '合作模式(代理結帳類型, 1: 買分, 2: 信用)';
COMMENT ON COLUMN "public"."agent"."coin_limit" IS '買分模式分數上限';
COMMENT ON COLUMN "public"."agent"."coin_use" IS '已使用分數';
COMMENT ON COLUMN "public"."agent"."coin_supply_setting" IS '自動補分設定';
COMMENT ON COLUMN "public"."agent"."level_code" IS '層級碼,每四碼一個層級 (admin:0000)';
COMMENT ON COLUMN "public"."agent"."aes_key" IS 'h5game專用aes密鑰';

ALTER TABLE "public"."user_play_log_baccarat" 
  DROP CONSTRAINT "play_log_baccarat_pkey",
  DROP COLUMN "agent_group_id",
  ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id");


COMMENT ON COLUMN "public"."admin_user"."account_type" IS '管理者類型 1:總台號,2:總代理,3:子代理';
