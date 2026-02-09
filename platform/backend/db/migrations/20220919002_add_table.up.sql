-- 刪除存在的遊戲資料 --
DELETE FROM "public"."game" WHERE "id" IN (1002,1003);

INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "h5_link")
VALUES (1002, 'dev', '番攤', 'fantan', 1, 'http://172.30.0.152/dev/ma/'),
  (1003, 'dev', '色碟', 'colordisc', 1, 'http://172.30.0.152/dev/ma/');


-- 新增番攤玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_fantan" (
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
  CONSTRAINT "play_log_fantan_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_fantan"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_fantan"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_fantan"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_fantan"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_fantan"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_fantan"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_fantan"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_fantan"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_fantan"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_fantan"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_fantan"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_fantan"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_fantan"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_fantan"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_fantan"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_fantan"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_fantan"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_fantan"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_fantan"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_fantan"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."user_play_log_fantan" IS '番攤玩家遊戲記錄';


-- 新增色碟玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_colordisc" (
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
  CONSTRAINT "play_log_colordisc_pkey" PRIMARY KEY("lognumber", "agent_id", "user_id", "game_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_colordisc"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."user_play_log_colordisc" IS '色碟玩家遊戲記錄';
