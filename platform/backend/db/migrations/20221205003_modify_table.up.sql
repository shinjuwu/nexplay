DROP TABLE IF EXISTS "public"."rt_game_stat";
CREATE TABLE "public"."rt_game_stat" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "de_score" numeric(20,4) NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "rt_game_stat_pkey" PRIMARY KEY ("agent_id", "game_id")
)
;
COMMENT ON COLUMN "public"."rt_game_stat"."bet_count" IS '總注單數';
COMMENT ON COLUMN "public"."rt_game_stat"."ya_score" IS '總投注(投注)';
COMMENT ON COLUMN "public"."rt_game_stat"."vaild_ya_score" IS '總有效投注';
COMMENT ON COLUMN "public"."rt_game_stat"."de_score" IS '總派獎(玩家得分)';
COMMENT ON COLUMN "public"."rt_game_stat"."bonus" IS '總紅利';
COMMENT ON COLUMN "public"."rt_game_stat"."tax" IS '總抽水';
COMMENT ON TABLE "public"."rt_game_stat" IS '遊戲即時統計資料(realtime game stat)';

DROP TABLE IF EXISTS "public"."rt_game_stat_hour";
CREATE TABLE "public"."rt_game_stat_hour" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "de_score" numeric(20,4) NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "rt_game_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "game_id")
)
;
COMMENT ON COLUMN "public"."rt_game_stat_hour"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."ya_score" IS '總投注(投注)';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."vaild_ya_score" IS '總有效投注';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."de_score" IS '總派獎(玩家得分)';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."bonus" IS '總紅利';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."tax" IS '總抽水';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."bet_user" IS '不重複投注人數';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."bet_count" IS '總注單數';
COMMENT ON TABLE "public"."rt_game_stat_hour" IS '遊戲即時統計資料(realtime game stat)';