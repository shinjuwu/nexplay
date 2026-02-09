-- +migrate Up
CREATE TABLE "public"."wallet_ledger" (
  "platform" varchar(6) COLLATE "pg_catalog"."default" NOT NULL,
  "id" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL DEFAULT 0,
  "agent_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "kind" int2 NOT NULL DEFAULT 0,
  "status" int2 NOT NULL DEFAULT 0,
  "changeset" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "wallet_ledger_pkey" PRIMARY KEY ("platform", "id", "user_id")
)
;
COMMENT ON COLUMN "public"."wallet_ledger"."platform" IS '平台碼(6碼)';
COMMENT ON COLUMN "public"."wallet_ledger"."id" IS '訂單號(orderid)';
COMMENT ON COLUMN "public"."wallet_ledger"."agent_id" IS 'mapping id of agent';
COMMENT ON COLUMN "public"."wallet_ledger"."user_id" IS '用戶id';
COMMENT ON COLUMN "public"."wallet_ledger"."username" IS '用戶帳號';
COMMENT ON COLUMN "public"."wallet_ledger"."kind" IS '上下分類型(1:上分,2:下分)';
COMMENT ON COLUMN "public"."wallet_ledger"."status" IS '訂單狀態';
COMMENT ON COLUMN "public"."wallet_ledger"."changeset" IS '變更摘要';
COMMENT ON COLUMN "public"."wallet_ledger"."create_time" IS '創建時間';
COMMENT ON TABLE "public"."wallet_ledger" IS '玩家上下分記錄表';

CREATE INDEX "idx_wallet_ledger_platform" ON "public"."wallet_ledger" USING btree (
  "platform" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" DESC NULLS LAST
);

CREATE TABLE "public"."user_play_log" (
  "platform" varchar(6) COLLATE "pg_catalog"."default" NOT NULL,
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "agent_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "game_id" int4 NOT NULL,
  "game_name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "room_type" int4 NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "bet_id" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_time" timestamptz(6) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "user_play_log_pkey" PRIMARY KEY ("platform", "lognumber", "bet_id")
)
;
COMMENT ON COLUMN "public"."user_play_log"."platform" IS '平台碼(6碼)';
COMMENT ON COLUMN "public"."user_play_log"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log"."game_name" IS '遊戲名稱';
COMMENT ON COLUMN "public"."user_play_log"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log"."create_time" IS '記錄時間';
COMMENT ON TABLE "public"."user_play_log" IS '玩家輸贏記錄';


CREATE INDEX "idx_user_play_log_bet_time" ON "public"."user_play_log" USING btree (
  "bet_time" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_user_play_log_platform" ON "public"."user_play_log" USING btree (
  "platform" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" DESC NULLS LAST
);

-- +migrate Down
DROP TABLE IF EXISTS
    wallet_ledger, user_play_log;