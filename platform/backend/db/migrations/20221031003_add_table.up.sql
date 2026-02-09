-- 新增21點玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_blackjack" (
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
  CONSTRAINT "play_log_blackjack_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_blackjack"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."user_play_log_blackjack" IS '21點玩家遊戲記錄';
