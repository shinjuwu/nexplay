-- 新增火箭 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (1008, 'dev01', '火箭', 'rocket', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增火箭遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10080, '火箭新手房', 1, 1008, 0),
  (10081, '火箭普通房', 1, 1008, 1),
  (10082, '火箭高手房', 1, 1008, 2),
  (10083, '火箭大師房', 1, 1008, 3);

-- 新增代理火箭遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" = 1008;

-- 新增代理火箭遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" = 1008;

-- 新增火箭遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_rocket" (
  "bet_id" varchar(30),
  "lognumber" varchar(100) NOT NULL,
  "agent_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) NOT NULL DEFAULT '',
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
  "level_code" varchar(128) NOT NULL DEFAULT '',
  "kill_type" int2 NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "play_log_rocket_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_rocket_lognumber" ON "public"."user_play_log_rocket" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_rocket_betid" ON "public"."user_play_log_rocket" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_rocket_bet_time" ON "public"."user_play_log_rocket" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_rocket"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_rocket"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_rocket"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_rocket"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_rocket"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_rocket"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_rocket"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_rocket"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_rocket"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_rocket"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rocket"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rocket"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_rocket"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_rocket"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rocket"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rocket"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_rocket"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_rocket"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_rocket"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_rocket"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_rocket"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_rocket"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_rocket"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_rocket" IS '火箭玩家遊戲記錄';
