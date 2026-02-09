-- 調整其他玩家遊戲紀錄表PK條件 --
ALTER TABLE "public"."user_play_log_baccarat" 
  DROP CONSTRAINT "play_log_baccarat_pkey",
  ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id");

ALTER TABLE "public"."user_play_log_fantan" 
  DROP CONSTRAINT "play_log_fantan_pkey",
  ADD CONSTRAINT "play_log_fantan_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id");

ALTER TABLE "public"."user_play_log_colordisc" 
  DROP CONSTRAINT "play_log_colordisc_pkey",
  ADD CONSTRAINT "play_log_colordisc_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id");


-- 新增魚蝦蟹玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_prawncrab" (
  "bet_id" serial8,
  "lognumber" varchar(100) NOT NULL,
  "agent_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "desk_id" int4 NOT NULL,
  "seat_id" int4 NOT NULL,
  "exchange" int4 NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "ya_score" numeric(20,4) NOT NULL,
  "valid_score" numeric(20,4) NOT NULL,
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL,
  CONSTRAINT "play_log_prawncrab_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_prawncrab"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."user_play_log_prawncrab" IS '番攤玩家遊戲記錄';
