-- 新增3公遊戲 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (2002, 'dev', '三公', 'sangong', 1, 2, 4, 0, 1, 'http://172.30.0.152/client/vue/');

-- 新增3公遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (20020, '三公新手房',     1, 2002, 0),
  (20021, '三公普通房',     1, 2002, 1),
  (20022, '三公高手房',     1, 2002, 2),
  (20023, '三公大師房',     1, 2002, 3);

-- 新增代理3公遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" = 2002;

-- 新增代理3公遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" = 2002;

-- 新增3公玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_sangong" (
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
  CONSTRAINT "play_log_sangong_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_sangong"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_sangong"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_sangong"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_sangong"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_sangong"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_sangong"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_sangong"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_sangong"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_sangong"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_sangong"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_sangong"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_sangong"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_sangong"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_sangong"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_sangong"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_sangong"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_sangong"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_sangong"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_sangong"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_sangong"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."user_play_log_sangong" IS '三公玩家遊戲記錄';
