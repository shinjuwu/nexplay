-- 新增百人骰寶、搶庄牛牛遊戲 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (1005, 'dev', '百人骰寶', 'hundredsicbo', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/'),
  (2003, 'dev', '搶庄牛牛', 'bullbull', 1, 2, 4, 0, 1, 'http://172.30.0.152/client/vue/');

-- 新增百人骰寶、搶庄牛牛遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10050, '百人骰寶新手房', 1, 1005, 0),
  (10051, '百人骰寶普通房', 1, 1005, 1),
  (10052, '百人骰寶高手房', 1, 1005, 2),
  (10053, '百人骰寶大師房', 1, 1005, 3),
  (20030, '搶庄牛牛新手房', 1, 2003, 0),
  (20031, '搶庄牛牛普通房', 1, 2003, 1),
  (20032, '搶庄牛牛高手房', 1, 2003, 2),
  (20033, '搶庄牛牛大師房', 1, 2003, 3);

-- 新增代理百人骰寶、搶庄牛牛遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (1005, 2003);

-- 新增代理百人骰寶、搶庄牛牛遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (1005, 2003);

-- 新增百人骰寶玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_hundredsicbo" (
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
  "level_code" varchar(128) NOT NULL DEFAULT '',
  CONSTRAINT "play_log_hundredsicbo_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."level_code" IS '代理層級碼';
COMMENT ON TABLE "public"."user_play_log_hundredsicbo" IS '百人骰寶玩家遊戲記錄';

-- 新增搶庄牛牛玩家遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_bullbull" (
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
  "level_code" varchar(128) NOT NULL DEFAULT '',
  CONSTRAINT "play_log_bullbull_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_bullbull"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."level_code" IS '代理層級碼';
COMMENT ON TABLE "public"."user_play_log_bullbull" IS '搶庄牛牛玩家遊戲記錄';
