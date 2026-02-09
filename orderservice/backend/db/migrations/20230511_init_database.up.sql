DROP TABLE IF EXISTS "public"."user_play_log";
CREATE TABLE "public"."user_play_log" (
  "bet_id" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "room_type" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "desk_id" int4 NOT NULL,
  "seat_id" int4 NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "ya_score" numeric(20,4) NOT NULL,
  "valid_score" numeric(20,4) NOT NULL,
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "bet_time" timestamptz(6) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "start_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  "end_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  CONSTRAINT "play_log_pkey" PRIMARY KEY ("bet_id", "username")
)
;
COMMENT ON COLUMN "public"."user_play_log"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log"."start_time" IS '遊戲開始時間';
COMMENT ON COLUMN "public"."user_play_log"."end_time" IS '遊戲結束時間';
COMMENT ON TABLE "public"."user_play_log" IS '玩家遊戲記錄';

-- ----------------------------
-- Indexes structure for table user_play_log
-- ----------------------------
CREATE INDEX "idx_user_play_log_lognumber" ON "public"."user_play_log" USING btree (
  "lognumber" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_ops" DESC NULLS LAST
);
CREATE INDEX "idx_user_play_log_username" ON "public"."user_play_log" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_ops" DESC NULLS LAST
);
CREATE INDEX "idx_user_play_log_bet_id" ON "public"."user_play_log" USING btree (
  "bet_id" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_ops" DESC NULLS LAST
);
CREATE INDEX "idx_user_play_log_game_id" ON "public"."user_play_log" USING btree (
  "game_id" "pg_catalog"."int4_ops" DESC NULLS LAST
);
CREATE INDEX "idx_user_play_log_bet_time" ON "public"."user_play_log" USING btree (
  "bet_time" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);