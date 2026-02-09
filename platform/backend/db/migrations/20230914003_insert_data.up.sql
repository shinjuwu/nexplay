CREATE TABLE "public"."jackpot_token_log" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "token_id" varchar(100) NOT NULL DEFAULT '',
  "agent_id" int4 NOT NULL DEFAULT 0,
  "level_code" varchar(128) NOT NULL DEFAULT '',
  "user_id" int4 NOT NULL DEFAULT 0,
  "username" varchar(100) NOT NULL DEFAULT '',
  "source_game_id" int4 NOT NULL DEFAULT 0,
  "source_lognumber" varchar(100) NOT NULL DEFAULT '',
  "source_bet_id" varchar(30) NOT NULL DEFAULT '',
  "jp_bet" numeric(20,4) NOT NULL DEFAULT 0,
  "usage_count" int4 NOT NULL DEFAULT 0,
  "creator" varchar(20) NOT NULL DEFAULT '',
  "info" varchar(500) NOT NULL DEFAULT '',
  "token_create_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "status" int2 NOT NULL DEFAULT 0,
  "error_code" int4 NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "jackpot_token_log_pkey" PRIMARY KEY ("id")
)
;

CREATE INDEX "idx_jackpot_token_log_token_token_id" ON "public"."jackpot_token_log" ("token_id");
CREATE INDEX "idx_jackpot_token_log_token_agent_id" ON "public"."jackpot_token_log" ("agent_id");
CREATE INDEX "idx_jackpot_token_log_token_username" ON "public"."jackpot_token_log" ("username");
CREATE INDEX "idx_jackpot_token_log_token_create_time" ON "public"."jackpot_token_log" ("token_create_time" DESC NULLS LAST);

COMMENT ON COLUMN "public"."jackpot_token_log"."id" IS '訂單編號';
COMMENT ON COLUMN "public"."jackpot_token_log"."token_id" IS '代幣編號';
COMMENT ON COLUMN "public"."jackpot_token_log"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."jackpot_token_log"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."jackpot_token_log"."user_id" IS '玩家編號';
COMMENT ON COLUMN "public"."jackpot_token_log"."username" IS '玩家名稱';
COMMENT ON COLUMN "public"."jackpot_token_log"."source_game_id" IS '代幣來源遊戲';
COMMENT ON COLUMN "public"."jackpot_token_log"."source_lognumber" IS '代幣來源局號';
COMMENT ON COLUMN "public"."jackpot_token_log"."source_bet_id" IS '代幣來源玩家遊戲訂單號';
COMMENT ON COLUMN "public"."jackpot_token_log"."jp_bet" IS '代幣分數';
COMMENT ON COLUMN "public"."jackpot_token_log"."usage_count" IS '代幣使用次數';
COMMENT ON COLUMN "public"."jackpot_token_log"."creator" IS '代幣創建者';
COMMENT ON COLUMN "public"."jackpot_token_log"."info" IS '備註';
COMMENT ON COLUMN "public"."jackpot_token_log"."token_create_time" IS '代幣創建時間';
COMMENT ON COLUMN "public"."jackpot_token_log"."status" IS '訂單狀態';
COMMENT ON COLUMN "public"."jackpot_token_log"."error_code" IS '錯誤代碼';
COMMENT ON COLUMN "public"."jackpot_token_log"."create_time" IS '訂單創建時間';
COMMENT ON COLUMN "public"."jackpot_token_log"."update_time" IS '訂單更新時間';
COMMENT ON TABLE "public"."jackpot_token_log" IS 'jp代幣紀錄';

CREATE TABLE "public"."jackpot_log" (
  "bet_id" varchar(30) NOT NULL DEFAULT '',
  "lognumber" varchar(100) NOT NULL DEFAULT '',
  "token_id" varchar(100) NOT NULL DEFAULT '',
  "agent_id" int4 NOT NULL DEFAULT 0,
  "level_code" varchar(128) NOT NULL DEFAULT '',
  "user_id" int4 NOT NULL DEFAULT 0,
  "username" varchar(100) NOT NULL DEFAULT '',
  "jp_bet" numeric(20,4) NOT NULL DEFAULT 0,
  "token_create_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "prize_score" numeric(20,4) NOT NULL DEFAULT 0,
  "prize_item" int4 NOT NULL DEFAULT 0,
  "winning_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "show_pool" numeric(20,4) NOT NULL DEFAULT 0,
  "real_pool" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL DEFAULT 0,
  CONSTRAINT "jackpot_log_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "token_id")
)
;

CREATE INDEX "idx_jackpot_log_lognumber" ON "public"."jackpot_log" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_jackpot_log_betid" ON "public"."jackpot_log" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_jackpot_log_winning_time" ON "public"."jackpot_log" ("winning_time" DESC NULLS LAST);

COMMENT ON COLUMN "public"."jackpot_log"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."jackpot_log"."lognumber" IS '局號';
COMMENT ON COLUMN "public"."jackpot_log"."token_id" IS '代幣編號';
COMMENT ON COLUMN "public"."jackpot_log"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."jackpot_log"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."jackpot_log"."user_id" IS '玩家編號';
COMMENT ON COLUMN "public"."jackpot_log"."username" IS '玩家名稱';
COMMENT ON COLUMN "public"."jackpot_log"."jp_bet" IS '代幣分數';
COMMENT ON COLUMN "public"."jackpot_log"."token_create_time" IS '代幣創建時間';
COMMENT ON COLUMN "public"."jackpot_log"."prize_score" IS '中獎分數';
COMMENT ON COLUMN "public"."jackpot_log"."prize_item" IS '中獎項目';
COMMENT ON COLUMN "public"."jackpot_log"."winning_time" IS '中獎時間';
COMMENT ON COLUMN "public"."jackpot_log"."show_pool" IS '公告獎池';
COMMENT ON COLUMN "public"."jackpot_log"."real_pool" IS '真實獎池';
COMMENT ON COLUMN "public"."jackpot_log"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."jackpot_log"."is_robot" IS '是否為機器人';
COMMENT ON TABLE "public"."jackpot_log" IS 'jp紀錄';