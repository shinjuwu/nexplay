-- 新增鬥雞、賽狗、拉密、炸金花、泰式博丁 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (1006, 'dev', '鬥雞', 'cockfight', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/apps'),
  (1007, 'dev', '賽狗', 'dogracing', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/apps'),
  (2005, 'dev', '拉密', 'rummy', 1, 2, 4, 1, 1, 'http://172.30.0.152/client/vue/apps'),
  (2006, 'dev', '炸金花', 'goldenflower', 1, 2, 4, 1, 1, 'http://172.30.0.152/client/vue/apps'),
  (2007, 'dev', '泰式博丁', 'pokdeng', 1, 2, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增鬥雞、賽狗、拉密、炸金花、泰式博丁遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10060, '鬥雞新手房', 1, 1006, 0),
  (10061, '鬥雞普通房', 1, 1006, 1),
  (10062, '鬥雞高手房', 1, 1006, 2),
  (10063, '鬥雞大師房', 1, 1006, 3),
  (10070, '賽狗新手房', 1, 1007, 0),
  (10071, '賽狗普通房', 1, 1007, 1),
  (10072, '賽狗高手房', 1, 1007, 2),
  (10073, '賽狗大師房', 1, 1007, 3),
  (20050, '拉密新手房', 1, 2005, 0),
  (20051, '拉密普通房', 1, 2005, 1),
  (20052, '拉密高手房', 1, 2005, 2),
  (20053, '拉密大師房', 1, 2005, 3),
  (20060, '炸金花新手房', 1, 2006, 0),
  (20061, '炸金花普通房', 1, 2006, 1),
  (20062, '炸金花高手房', 1, 2006, 2),
  (20063, '炸金花大師房', 1, 2006, 3),
  (20070, '泰式博丁新手房', 1, 2007, 0),
  (20071, '泰式博丁普通房', 1, 2007, 1),
  (20072, '泰式博丁高手房', 1, 2007, 2),
  (20073, '泰式博丁大師房', 1, 2007, 3);

-- 新增代理鬥雞、賽狗、拉密、炸金花、泰式博丁遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (1006, 1007, 2005, 2006, 2007);

-- 新增代理鬥雞、賽狗、拉密、炸金花、泰式博丁遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (1006, 1007, 2005, 2006, 2007);

-- 新增鬥雞遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_cockfight" (
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
  CONSTRAINT "play_log_cockfight_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_cockfight_lognumber" ON "public"."user_play_log_cockfight" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_cockfight_betid" ON "public"."user_play_log_cockfight" ("bet_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_cockfight"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_cockfight" IS '鬥雞玩家遊戲記錄';

-- 新增賽狗遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_dogracing" (
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
  CONSTRAINT "play_log_dogracing_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_dogracing_lognumber" ON "public"."user_play_log_dogracing" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_dogracing_betid" ON "public"."user_play_log_dogracing" ("bet_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_dogracing"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_dogracing" IS '賽狗玩家遊戲記錄';

-- 新增拉密遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_rummy" (
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
  CONSTRAINT "play_log_rummy_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_rummy_lognumber" ON "public"."user_play_log_rummy" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_rummy_betid" ON "public"."user_play_log_rummy" ("bet_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_rummy"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_rummy"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_rummy"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_rummy"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_rummy"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_rummy"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_rummy"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_rummy"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_rummy"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_rummy"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rummy"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rummy"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_rummy"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_rummy"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rummy"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_rummy"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_rummy"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_rummy"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_rummy"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_rummy"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_rummy"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_rummy"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_rummy"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_rummy" IS '拉密玩家遊戲記錄';

-- 新增炸金花遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_goldenflower" (
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
  CONSTRAINT "play_log_goldenflower_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_goldenflower_lognumber" ON "public"."user_play_log_goldenflower" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_goldenflower_betid" ON "public"."user_play_log_goldenflower" ("bet_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_goldenflower"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_goldenflower" IS '炸金花玩家遊戲記錄';

-- 新增泰式博丁遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_pokdeng" (
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
  CONSTRAINT "play_log_pokdeng_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_pokdeng_lognumber" ON "public"."user_play_log_pokdeng" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_pokdeng_betid" ON "public"."user_play_log_pokdeng" ("bet_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_pokdeng"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_pokdeng" IS '泰式博丁玩家遊戲記錄';

-- view_user_play_log新增鬥雞、賽狗、拉密、炸金花、泰式博丁 --
DROP VIEW "public"."view_user_play_log";
CREATE OR REPLACE VIEW "public"."view_user_play_log" AS
 SELECT tmp.bet_id,
    tmp.lognumber,
    tmp.agent_id,
    tmp.game_id,
    tmp.room_type,
    tmp.desk_id,
    tmp.seat_id,
    tmp.exchange,
    tmp.de_score,
    tmp.ya_score,
    tmp.valid_score,
    tmp.start_score,
    tmp.end_score,
    tmp.create_time,
    tmp.is_robot,
    tmp.is_big_win,
    tmp.is_issue,
    tmp.bet_time,
    tmp.tax,
    tmp.bonus,
    tmp.level_code,
    tmp.username,
    tmp.kill_type
   FROM ( SELECT baccarat.bet_id,
            baccarat.lognumber,
            baccarat.agent_id,
            baccarat.game_id,
            baccarat.room_type,
            baccarat.desk_id,
            baccarat.seat_id,
            baccarat.exchange,
            baccarat.de_score,
            baccarat.ya_score,
            baccarat.valid_score,
            baccarat.start_score,
            baccarat.end_score,
            baccarat.create_time,
            baccarat.is_robot,
            baccarat.is_big_win,
            baccarat.is_issue,
            baccarat.bet_time,
            baccarat.tax,
            baccarat.level_code,
            baccarat.username,
            baccarat.kill_type,
            baccarat.bonus
           FROM user_play_log_baccarat baccarat
           WHERE baccarat.is_robot = 0 AND baccarat.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT fantan.bet_id,
            fantan.lognumber,
            fantan.agent_id,
            fantan.game_id,
            fantan.room_type,
            fantan.desk_id,
            fantan.seat_id,
            fantan.exchange,
            fantan.de_score,
            fantan.ya_score,
            fantan.valid_score,
            fantan.start_score,
            fantan.end_score,
            fantan.create_time,
            fantan.is_robot,
            fantan.is_big_win,
            fantan.is_issue,
            fantan.bet_time,
            fantan.tax,
            fantan.level_code,
            fantan.username,
            fantan.kill_type,
            fantan.bonus
           FROM user_play_log_fantan fantan
           WHERE fantan.is_robot = 0 AND fantan.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT colordisc.bet_id,
            colordisc.lognumber,
            colordisc.agent_id,
            colordisc.game_id,
            colordisc.room_type,
            colordisc.desk_id,
            colordisc.seat_id,
            colordisc.exchange,
            colordisc.de_score,
            colordisc.ya_score,
            colordisc.valid_score,
            colordisc.start_score,
            colordisc.end_score,
            colordisc.create_time,
            colordisc.is_robot,
            colordisc.is_big_win,
            colordisc.is_issue,
            colordisc.bet_time,
            colordisc.tax,
            colordisc.level_code,
            colordisc.username,
            colordisc.kill_type,
            colordisc.bonus
           FROM user_play_log_colordisc colordisc
           WHERE colordisc.is_robot = 0 AND colordisc.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT prawncrab.bet_id,
            prawncrab.lognumber,
            prawncrab.agent_id,
            prawncrab.game_id,
            prawncrab.room_type,
            prawncrab.desk_id,
            prawncrab.seat_id,
            prawncrab.exchange,
            prawncrab.de_score,
            prawncrab.ya_score,
            prawncrab.valid_score,
            prawncrab.start_score,
            prawncrab.end_score,
            prawncrab.create_time,
            prawncrab.is_robot,
            prawncrab.is_big_win,
            prawncrab.is_issue,
            prawncrab.bet_time,
            prawncrab.tax,
            prawncrab.level_code,
            prawncrab.username,
            prawncrab.kill_type,
            prawncrab.bonus
           FROM user_play_log_prawncrab prawncrab
           WHERE prawncrab.is_robot = 0 AND prawncrab.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT hundredsicbo.bet_id,
            hundredsicbo.lognumber,
            hundredsicbo.agent_id,
            hundredsicbo.game_id,
            hundredsicbo.room_type,
            hundredsicbo.desk_id,
            hundredsicbo.seat_id,
            hundredsicbo.exchange,
            hundredsicbo.de_score,
            hundredsicbo.ya_score,
            hundredsicbo.valid_score,
            hundredsicbo.start_score,
            hundredsicbo.end_score,
            hundredsicbo.create_time,
            hundredsicbo.is_robot,
            hundredsicbo.is_big_win,
            hundredsicbo.is_issue,
            hundredsicbo.bet_time,
            hundredsicbo.tax,
            hundredsicbo.level_code,
            hundredsicbo.username,
            hundredsicbo.kill_type,
            hundredsicbo.bonus
           FROM user_play_log_hundredsicbo hundredsicbo
           WHERE hundredsicbo.is_robot = 0 AND hundredsicbo.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT blackjack.bet_id,
            blackjack.lognumber,
            blackjack.agent_id,
            blackjack.game_id,
            blackjack.room_type,
            blackjack.desk_id,
            blackjack.seat_id,
            blackjack.exchange,
            blackjack.de_score,
            blackjack.ya_score,
            blackjack.valid_score,
            blackjack.start_score,
            blackjack.end_score,
            blackjack.create_time,
            blackjack.is_robot,
            blackjack.is_big_win,
            blackjack.is_issue,
            blackjack.bet_time,
            blackjack.tax,
            blackjack.level_code,
            blackjack.username,
            blackjack.kill_type,
            blackjack.bonus
           FROM user_play_log_blackjack blackjack
           WHERE blackjack.is_robot = 0 AND blackjack.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT sangong.bet_id,
            sangong.lognumber,
            sangong.agent_id,
            sangong.game_id,
            sangong.room_type,
            sangong.desk_id,
            sangong.seat_id,
            sangong.exchange,
            sangong.de_score,
            sangong.ya_score,
            sangong.valid_score,
            sangong.start_score,
            sangong.end_score,
            sangong.create_time,
            sangong.is_robot,
            sangong.is_big_win,
            sangong.is_issue,
            sangong.bet_time,
            sangong.tax,
            sangong.level_code,
            sangong.username,
            sangong.kill_type,
            sangong.bonus
           FROM user_play_log_sangong sangong
           WHERE sangong.is_robot = 0 AND sangong.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT bullbull.bet_id,
            bullbull.lognumber,
            bullbull.agent_id,
            bullbull.game_id,
            bullbull.room_type,
            bullbull.desk_id,
            bullbull.seat_id,
            bullbull.exchange,
            bullbull.de_score,
            bullbull.ya_score,
            bullbull.valid_score,
            bullbull.start_score,
            bullbull.end_score,
            bullbull.create_time,
            bullbull.is_robot,
            bullbull.is_big_win,
            bullbull.is_issue,
            bullbull.bet_time,
            bullbull.tax,
            bullbull.level_code,
            bullbull.username,
            bullbull.kill_type,
            bullbull.bonus
           FROM user_play_log_bullbull bullbull
           WHERE bullbull.is_robot = 0 AND bullbull.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT texas.bet_id,
            texas.lognumber,
            texas.agent_id,
            texas.game_id,
            texas.room_type,
            texas.desk_id,
            texas.seat_id,
            texas.exchange,
            texas.de_score,
            texas.ya_score,
            texas.valid_score,
            texas.start_score,
            texas.end_score,
            texas.create_time,
            texas.is_robot,
            texas.is_big_win,
            texas.is_issue,
            texas.bet_time,
            texas.tax,
            texas.level_code,
            texas.username,
            texas.kill_type,
            texas.bonus
           FROM user_play_log_texas texas
           WHERE texas.is_robot = 0 AND texas.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT fruitslot.bet_id,
            fruitslot.lognumber,
            fruitslot.agent_id,
            fruitslot.game_id,
            fruitslot.room_type,
            fruitslot.desk_id,
            fruitslot.seat_id,
            fruitslot.exchange,
            fruitslot.de_score,
            fruitslot.ya_score,
            fruitslot.valid_score,
            fruitslot.start_score,
            fruitslot.end_score,
            fruitslot.create_time,
            fruitslot.is_robot,
            fruitslot.is_big_win,
            fruitslot.is_issue,
            fruitslot.bet_time,
            fruitslot.tax,
            fruitslot.level_code,
            fruitslot.username,
            fruitslot.kill_type,
            fruitslot.bonus
           FROM user_play_log_fruitslot fruitslot
           WHERE fruitslot.is_robot = 0 AND fruitslot.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT rcfishing.bet_id,
            rcfishing.lognumber,
            rcfishing.agent_id,
            rcfishing.game_id,
            rcfishing.room_type,
            rcfishing.desk_id,
            rcfishing.seat_id,
            rcfishing.exchange,
            rcfishing.de_score,
            rcfishing.ya_score,
            rcfishing.valid_score,
            rcfishing.start_score,
            rcfishing.end_score,
            rcfishing.create_time,
            rcfishing.is_robot,
            rcfishing.is_big_win,
            rcfishing.is_issue,
            rcfishing.bet_time,
            rcfishing.tax,
            rcfishing.level_code,
            rcfishing.username,
            rcfishing.kill_type,
            rcfishing.bonus
           FROM user_play_log_rcfishing rcfishing
           WHERE rcfishing.is_robot = 0 AND rcfishing.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT cockfight.bet_id,
            cockfight.lognumber,
            cockfight.agent_id,
            cockfight.game_id,
            cockfight.room_type,
            cockfight.desk_id,
            cockfight.seat_id,
            cockfight.exchange,
            cockfight.de_score,
            cockfight.ya_score,
            cockfight.valid_score,
            cockfight.start_score,
            cockfight.end_score,
            cockfight.create_time,
            cockfight.is_robot,
            cockfight.is_big_win,
            cockfight.is_issue,
            cockfight.bet_time,
            cockfight.tax,
            cockfight.level_code,
            cockfight.username,
            cockfight.kill_type,
            cockfight.bonus
           FROM user_play_log_cockfight cockfight
           WHERE cockfight.is_robot = 0 AND cockfight.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT dogracing.bet_id,
            dogracing.lognumber,
            dogracing.agent_id,
            dogracing.game_id,
            dogracing.room_type,
            dogracing.desk_id,
            dogracing.seat_id,
            dogracing.exchange,
            dogracing.de_score,
            dogracing.ya_score,
            dogracing.valid_score,
            dogracing.start_score,
            dogracing.end_score,
            dogracing.create_time,
            dogracing.is_robot,
            dogracing.is_big_win,
            dogracing.is_issue,
            dogracing.bet_time,
            dogracing.tax,
            dogracing.level_code,
            dogracing.username,
            dogracing.kill_type,
            dogracing.bonus
           FROM user_play_log_dogracing dogracing
           WHERE dogracing.is_robot = 0 AND dogracing.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT rummy.bet_id,
            rummy.lognumber,
            rummy.agent_id,
            rummy.game_id,
            rummy.room_type,
            rummy.desk_id,
            rummy.seat_id,
            rummy.exchange,
            rummy.de_score,
            rummy.ya_score,
            rummy.valid_score,
            rummy.start_score,
            rummy.end_score,
            rummy.create_time,
            rummy.is_robot,
            rummy.is_big_win,
            rummy.is_issue,
            rummy.bet_time,
            rummy.tax,
            rummy.level_code,
            rummy.username,
            rummy.kill_type,
            rummy.bonus
           FROM user_play_log_rummy rummy
           WHERE rummy.is_robot = 0 AND rummy.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT goldenflower.bet_id,
            goldenflower.lognumber,
            goldenflower.agent_id,
            goldenflower.game_id,
            goldenflower.room_type,
            goldenflower.desk_id,
            goldenflower.seat_id,
            goldenflower.exchange,
            goldenflower.de_score,
            goldenflower.ya_score,
            goldenflower.valid_score,
            goldenflower.start_score,
            goldenflower.end_score,
            goldenflower.create_time,
            goldenflower.is_robot,
            goldenflower.is_big_win,
            goldenflower.is_issue,
            goldenflower.bet_time,
            goldenflower.tax,
            goldenflower.level_code,
            goldenflower.username,
            goldenflower.kill_type,
            goldenflower.bonus
           FROM user_play_log_goldenflower goldenflower
           WHERE goldenflower.is_robot = 0 AND goldenflower.bet_time > (now() - '1 mon'::interval)
        UNION
         SELECT pokdeng.bet_id,
            pokdeng.lognumber,
            pokdeng.agent_id,
            pokdeng.game_id,
            pokdeng.room_type,
            pokdeng.desk_id,
            pokdeng.seat_id,
            pokdeng.exchange,
            pokdeng.de_score,
            pokdeng.ya_score,
            pokdeng.valid_score,
            pokdeng.start_score,
            pokdeng.end_score,
            pokdeng.create_time,
            pokdeng.is_robot,
            pokdeng.is_big_win,
            pokdeng.is_issue,
            pokdeng.bet_time,
            pokdeng.tax,
            pokdeng.level_code,
            pokdeng.username,
            pokdeng.kill_type,
            pokdeng.bonus
           FROM user_play_log_pokdeng pokdeng
           WHERE pokdeng.is_robot = 0 AND pokdeng.bet_time > (now() - '1 mon'::interval)) tmp;

-- mv_cal_game_stat_hour新增鬥雞、賽狗、拉密、炸金花、泰式博丁 --
DROP MATERIALIZED VIEW mv_cal_game_stat_hour;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_cal_game_stat_hour AS
SELECT date_trunc('hour'::text, tmp.bet_time) AS log_time,
    tmp.level_code,
    tmp.game_id,
    count(DISTINCT tmp.user_id) AS bet_user,
    count(tmp.lognumber) AS bet_count,
    sum(tmp.ya_score) AS sum_ya,
    sum(tmp.valid_score) AS sum_valid_ya,
    sum(tmp.de_score) AS sum_de,
    0 AS sum_bonus,
    sum(tmp.tax) AS sum_tax
   FROM ( SELECT user_play_log_baccarat.lognumber,
            user_play_log_baccarat.level_code,
            user_play_log_baccarat.user_id,
            user_play_log_baccarat.game_id,
            user_play_log_baccarat.ya_score,
            user_play_log_baccarat.valid_score,
            user_play_log_baccarat.de_score,
            user_play_log_baccarat.tax,
            user_play_log_baccarat.bonus,
            user_play_log_baccarat.bet_time
           FROM user_play_log_baccarat
          WHERE ((user_play_log_baccarat.is_robot = 0) AND (user_play_log_baccarat.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_colordisc.lognumber,
            user_play_log_colordisc.level_code,
            user_play_log_colordisc.user_id,
            user_play_log_colordisc.game_id,
            user_play_log_colordisc.ya_score,
            user_play_log_colordisc.valid_score,
            user_play_log_colordisc.de_score,
            user_play_log_colordisc.tax,
            user_play_log_colordisc.bonus,
            user_play_log_colordisc.bet_time
           FROM user_play_log_colordisc
          WHERE ((user_play_log_colordisc.is_robot = 0) AND (user_play_log_colordisc.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_fantan.lognumber,
            user_play_log_fantan.level_code,
            user_play_log_fantan.user_id,
            user_play_log_fantan.game_id,
            user_play_log_fantan.ya_score,
            user_play_log_fantan.valid_score,
            user_play_log_fantan.de_score,
            user_play_log_fantan.tax,
            user_play_log_fantan.bonus,
            user_play_log_fantan.bet_time
           FROM user_play_log_fantan
          WHERE ((user_play_log_fantan.is_robot = 0) AND (user_play_log_fantan.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_prawncrab.lognumber,
            user_play_log_prawncrab.level_code,
            user_play_log_prawncrab.user_id,
            user_play_log_prawncrab.game_id,
            user_play_log_prawncrab.ya_score,
            user_play_log_prawncrab.valid_score,
            user_play_log_prawncrab.de_score,
            user_play_log_prawncrab.tax,
            user_play_log_prawncrab.bonus,
            user_play_log_prawncrab.bet_time
           FROM user_play_log_prawncrab
          WHERE ((user_play_log_prawncrab.is_robot = 0) AND (user_play_log_prawncrab.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_blackjack.lognumber,
            user_play_log_blackjack.level_code,
            user_play_log_blackjack.user_id,
            user_play_log_blackjack.game_id,
            user_play_log_blackjack.ya_score,
            user_play_log_blackjack.valid_score,
            user_play_log_blackjack.de_score,
            user_play_log_blackjack.tax,
            user_play_log_blackjack.bonus,
            user_play_log_blackjack.bet_time
           FROM user_play_log_blackjack
          WHERE ((user_play_log_blackjack.is_robot = 0) AND (user_play_log_blackjack.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_sangong.lognumber,
            user_play_log_sangong.level_code,
            user_play_log_sangong.user_id,
            user_play_log_sangong.game_id,
            user_play_log_sangong.ya_score,
            user_play_log_sangong.valid_score,
            user_play_log_sangong.de_score,
            user_play_log_sangong.tax,
            user_play_log_sangong.bonus,
            user_play_log_sangong.bet_time
           FROM user_play_log_sangong
          WHERE ((user_play_log_sangong.is_robot = 0) AND (user_play_log_sangong.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_hundredsicbo.lognumber,
            user_play_log_hundredsicbo.level_code,
            user_play_log_hundredsicbo.user_id,
            user_play_log_hundredsicbo.game_id,
            user_play_log_hundredsicbo.ya_score,
            user_play_log_hundredsicbo.valid_score,
            user_play_log_hundredsicbo.de_score,
            user_play_log_hundredsicbo.tax,
            user_play_log_hundredsicbo.bonus,
            user_play_log_hundredsicbo.bet_time
           FROM user_play_log_hundredsicbo
          WHERE ((user_play_log_hundredsicbo.is_robot = 0) AND (user_play_log_hundredsicbo.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_bullbull.lognumber,
            user_play_log_bullbull.level_code,
            user_play_log_bullbull.user_id,
            user_play_log_bullbull.game_id,
            user_play_log_bullbull.ya_score,
            user_play_log_bullbull.valid_score,
            user_play_log_bullbull.de_score,
            user_play_log_bullbull.tax,
            user_play_log_bullbull.bonus,
            user_play_log_bullbull.bet_time
           FROM user_play_log_bullbull
          WHERE ((user_play_log_bullbull.is_robot = 0) AND (user_play_log_bullbull.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_texas.lognumber,
            user_play_log_texas.level_code,
            user_play_log_texas.user_id,
            user_play_log_texas.game_id,
            user_play_log_texas.ya_score,
            user_play_log_texas.valid_score,
            user_play_log_texas.de_score,
            user_play_log_texas.tax,
            user_play_log_texas.bonus,
            user_play_log_texas.bet_time
           FROM user_play_log_texas
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_fruitslot.lognumber,
            user_play_log_fruitslot.level_code,
            user_play_log_fruitslot.user_id,
            user_play_log_fruitslot.game_id,
            user_play_log_fruitslot.ya_score,
            user_play_log_fruitslot.valid_score,
            user_play_log_fruitslot.de_score,
            user_play_log_fruitslot.tax,
            user_play_log_fruitslot.bonus,
            user_play_log_fruitslot.bet_time
           FROM user_play_log_fruitslot
           WHERE ((user_play_log_fruitslot.is_robot = 0) AND (user_play_log_fruitslot.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_rcfishing.lognumber,
            user_play_log_rcfishing.level_code,
            user_play_log_rcfishing.user_id,
            user_play_log_rcfishing.game_id,
            user_play_log_rcfishing.ya_score,
            user_play_log_rcfishing.valid_score,
            user_play_log_rcfishing.de_score,
            user_play_log_rcfishing.tax,
            user_play_log_rcfishing.bonus,
            user_play_log_rcfishing.bet_time
           FROM user_play_log_rcfishing
           WHERE ((user_play_log_rcfishing.is_robot = 0) AND (user_play_log_rcfishing.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_cockfight.lognumber,
            user_play_log_cockfight.level_code,
            user_play_log_cockfight.user_id,
            user_play_log_cockfight.game_id,
            user_play_log_cockfight.ya_score,
            user_play_log_cockfight.valid_score,
            user_play_log_cockfight.de_score,
            user_play_log_cockfight.tax,
            user_play_log_cockfight.bonus,
            user_play_log_cockfight.bet_time
           FROM user_play_log_cockfight
           WHERE ((user_play_log_cockfight.is_robot = 0) AND (user_play_log_cockfight.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_dogracing.lognumber,
            user_play_log_dogracing.level_code,
            user_play_log_dogracing.user_id,
            user_play_log_dogracing.game_id,
            user_play_log_dogracing.ya_score,
            user_play_log_dogracing.valid_score,
            user_play_log_dogracing.de_score,
            user_play_log_dogracing.tax,
            user_play_log_dogracing.bonus,
            user_play_log_dogracing.bet_time
           FROM user_play_log_dogracing
           WHERE ((user_play_log_dogracing.is_robot = 0) AND (user_play_log_dogracing.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_rummy.lognumber,
            user_play_log_rummy.level_code,
            user_play_log_rummy.user_id,
            user_play_log_rummy.game_id,
            user_play_log_rummy.ya_score,
            user_play_log_rummy.valid_score,
            user_play_log_rummy.de_score,
            user_play_log_rummy.tax,
            user_play_log_rummy.bonus,
            user_play_log_rummy.bet_time
           FROM user_play_log_rummy
           WHERE ((user_play_log_rummy.is_robot = 0) AND (user_play_log_rummy.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_goldenflower.lognumber,
            user_play_log_goldenflower.level_code,
            user_play_log_goldenflower.user_id,
            user_play_log_goldenflower.game_id,
            user_play_log_goldenflower.ya_score,
            user_play_log_goldenflower.valid_score,
            user_play_log_goldenflower.de_score,
            user_play_log_goldenflower.tax,
            user_play_log_goldenflower.bonus,
            user_play_log_goldenflower.bet_time
           FROM user_play_log_goldenflower
           WHERE ((user_play_log_goldenflower.is_robot = 0) AND (user_play_log_goldenflower.bet_time > (now() - '3 days'::interval)))
        UNION
         SELECT user_play_log_pokdeng.lognumber,
            user_play_log_pokdeng.level_code,
            user_play_log_pokdeng.user_id,
            user_play_log_pokdeng.game_id,
            user_play_log_pokdeng.ya_score,
            user_play_log_pokdeng.valid_score,
            user_play_log_pokdeng.de_score,
            user_play_log_pokdeng.tax,
            user_play_log_pokdeng.bonus,
            user_play_log_pokdeng.bet_time
           FROM user_play_log_pokdeng
           WHERE ((user_play_log_pokdeng.is_robot = 0) AND (user_play_log_pokdeng.bet_time > (now() - '3 days'::interval)))) tmp
  GROUP BY (date_trunc('hour'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;

-- view_agent_game_room_ratio_log新增鬥雞、賽狗、拉密、炸金花、泰式博丁 --
DROP VIEW "public"."view_agent_game_room_ratio_log";
CREATE OR REPLACE VIEW "public"."view_agent_game_room_ratio_log" AS  SELECT date_trunc('day'::text, tmp.bet_time) AS log_time,
    tmp.agent_id,
    tmp.level_code,
    tmp.game_id,
    tmp.room_type,
    count(DISTINCT tmp.user_id) AS bet_user,
    count(tmp.lognumber) AS bet_count,
    sum(tmp.ya_score) AS sum_ya,
    sum(tmp.valid_score) AS sum_valid_ya,
    sum(tmp.de_score) AS sum_de,
    sum(tmp.bonus) AS sum_bonus,
    sum(tmp.tax) AS sum_tax
   FROM ( SELECT user_play_log_baccarat.lognumber,
            user_play_log_baccarat.agent_id,
            user_play_log_baccarat.level_code,
            user_play_log_baccarat.user_id,
            user_play_log_baccarat.game_id,
            user_play_log_baccarat.room_type,
            user_play_log_baccarat.ya_score,
            user_play_log_baccarat.valid_score,
            user_play_log_baccarat.de_score,
            user_play_log_baccarat.tax,
            user_play_log_baccarat.bonus,
            user_play_log_baccarat.bet_time
           FROM user_play_log_baccarat
          WHERE user_play_log_baccarat.is_robot = 0 AND user_play_log_baccarat.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_colordisc.lognumber,
            user_play_log_colordisc.agent_id,
            user_play_log_colordisc.level_code,
            user_play_log_colordisc.user_id,
            user_play_log_colordisc.game_id,
            user_play_log_colordisc.room_type,
            user_play_log_colordisc.ya_score,
            user_play_log_colordisc.valid_score,
            user_play_log_colordisc.de_score,
            user_play_log_colordisc.tax,
            user_play_log_colordisc.bonus,
            user_play_log_colordisc.bet_time
           FROM user_play_log_colordisc
          WHERE user_play_log_colordisc.is_robot = 0 AND user_play_log_colordisc.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_fantan.lognumber,
            user_play_log_fantan.agent_id,
            user_play_log_fantan.level_code,
            user_play_log_fantan.user_id,
            user_play_log_fantan.game_id,
            user_play_log_fantan.room_type,
            user_play_log_fantan.ya_score,
            user_play_log_fantan.valid_score,
            user_play_log_fantan.de_score,
            user_play_log_fantan.tax,
            user_play_log_fantan.bonus,
            user_play_log_fantan.bet_time
           FROM user_play_log_fantan
          WHERE user_play_log_fantan.is_robot = 0 AND user_play_log_fantan.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_prawncrab.lognumber,
            user_play_log_prawncrab.agent_id,
            user_play_log_prawncrab.level_code,
            user_play_log_prawncrab.user_id,
            user_play_log_prawncrab.game_id,
            user_play_log_prawncrab.room_type,
            user_play_log_prawncrab.ya_score,
            user_play_log_prawncrab.valid_score,
            user_play_log_prawncrab.de_score,
            user_play_log_prawncrab.tax,
            user_play_log_prawncrab.bonus,
            user_play_log_prawncrab.bet_time
           FROM user_play_log_prawncrab
          WHERE user_play_log_prawncrab.is_robot = 0 AND user_play_log_prawncrab.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_blackjack.lognumber,
            user_play_log_blackjack.agent_id,
            user_play_log_blackjack.level_code,
            user_play_log_blackjack.user_id,
            user_play_log_blackjack.game_id,
            user_play_log_blackjack.room_type,
            user_play_log_blackjack.ya_score,
            user_play_log_blackjack.valid_score,
            user_play_log_blackjack.de_score,
            user_play_log_blackjack.tax,
            user_play_log_blackjack.bonus,
            user_play_log_blackjack.bet_time
           FROM user_play_log_blackjack
          WHERE user_play_log_blackjack.is_robot = 0 AND user_play_log_blackjack.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_sangong.lognumber,
            user_play_log_sangong.agent_id,
            user_play_log_sangong.level_code,
            user_play_log_sangong.user_id,
            user_play_log_sangong.game_id,
            user_play_log_sangong.room_type,
            user_play_log_sangong.ya_score,
            user_play_log_sangong.valid_score,
            user_play_log_sangong.de_score,
            user_play_log_sangong.tax,
            user_play_log_sangong.bonus,
            user_play_log_sangong.bet_time
           FROM user_play_log_sangong
          WHERE user_play_log_sangong.is_robot = 0 AND user_play_log_sangong.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_hundredsicbo.lognumber,
            user_play_log_hundredsicbo.agent_id,
            user_play_log_hundredsicbo.level_code,
            user_play_log_hundredsicbo.user_id,
            user_play_log_hundredsicbo.game_id,
            user_play_log_hundredsicbo.room_type,
            user_play_log_hundredsicbo.ya_score,
            user_play_log_hundredsicbo.valid_score,
            user_play_log_hundredsicbo.de_score,
            user_play_log_hundredsicbo.tax,
            user_play_log_hundredsicbo.bonus,
            user_play_log_hundredsicbo.bet_time
           FROM user_play_log_hundredsicbo
          WHERE user_play_log_hundredsicbo.is_robot = 0 AND user_play_log_hundredsicbo.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_bullbull.lognumber,
            user_play_log_bullbull.agent_id,
            user_play_log_bullbull.level_code,
            user_play_log_bullbull.user_id,
            user_play_log_bullbull.game_id,
            user_play_log_bullbull.room_type,
            user_play_log_bullbull.ya_score,
            user_play_log_bullbull.valid_score,
            user_play_log_bullbull.de_score,
            user_play_log_bullbull.tax,
            user_play_log_bullbull.bonus,
            user_play_log_bullbull.bet_time
           FROM user_play_log_bullbull
          WHERE user_play_log_bullbull.is_robot = 0 AND user_play_log_bullbull.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_texas.lognumber,
            user_play_log_texas.agent_id,
            user_play_log_texas.level_code,
            user_play_log_texas.user_id,
            user_play_log_texas.game_id,
            user_play_log_texas.room_type,
            user_play_log_texas.ya_score,
            user_play_log_texas.valid_score,
            user_play_log_texas.de_score,
            user_play_log_texas.tax,
            user_play_log_texas.bonus,
            user_play_log_texas.bet_time
           FROM user_play_log_texas
          WHERE user_play_log_texas.is_robot = 0 AND user_play_log_texas.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_fruitslot.lognumber,
            user_play_log_fruitslot.agent_id,
            user_play_log_fruitslot.level_code,
            user_play_log_fruitslot.user_id,
            user_play_log_fruitslot.game_id,
            user_play_log_fruitslot.room_type,
            user_play_log_fruitslot.ya_score,
            user_play_log_fruitslot.valid_score,
            user_play_log_fruitslot.de_score,
            user_play_log_fruitslot.tax,
            user_play_log_fruitslot.bonus,
            user_play_log_fruitslot.bet_time
           FROM user_play_log_fruitslot
          WHERE user_play_log_fruitslot.is_robot = 0 AND user_play_log_fruitslot.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_rcfishing.lognumber,
            user_play_log_rcfishing.agent_id,
            user_play_log_rcfishing.level_code,
            user_play_log_rcfishing.user_id,
            user_play_log_rcfishing.game_id,
            user_play_log_rcfishing.room_type,
            user_play_log_rcfishing.ya_score,
            user_play_log_rcfishing.valid_score,
            user_play_log_rcfishing.de_score,
            user_play_log_rcfishing.tax,
            user_play_log_rcfishing.bonus,
            user_play_log_rcfishing.bet_time
           FROM user_play_log_rcfishing
          WHERE user_play_log_rcfishing.is_robot = 0 AND user_play_log_rcfishing.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_cockfight.lognumber,
            user_play_log_cockfight.agent_id,
            user_play_log_cockfight.level_code,
            user_play_log_cockfight.user_id,
            user_play_log_cockfight.game_id,
            user_play_log_cockfight.room_type,
            user_play_log_cockfight.ya_score,
            user_play_log_cockfight.valid_score,
            user_play_log_cockfight.de_score,
            user_play_log_cockfight.tax,
            user_play_log_cockfight.bonus,
            user_play_log_cockfight.bet_time
           FROM user_play_log_cockfight
          WHERE user_play_log_cockfight.is_robot = 0 AND user_play_log_cockfight.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_dogracing.lognumber,
            user_play_log_dogracing.agent_id,
            user_play_log_dogracing.level_code,
            user_play_log_dogracing.user_id,
            user_play_log_dogracing.game_id,
            user_play_log_dogracing.room_type,
            user_play_log_dogracing.ya_score,
            user_play_log_dogracing.valid_score,
            user_play_log_dogracing.de_score,
            user_play_log_dogracing.tax,
            user_play_log_dogracing.bonus,
            user_play_log_dogracing.bet_time
           FROM user_play_log_dogracing
          WHERE user_play_log_dogracing.is_robot = 0 AND user_play_log_dogracing.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_rummy.lognumber,
            user_play_log_rummy.agent_id,
            user_play_log_rummy.level_code,
            user_play_log_rummy.user_id,
            user_play_log_rummy.game_id,
            user_play_log_rummy.room_type,
            user_play_log_rummy.ya_score,
            user_play_log_rummy.valid_score,
            user_play_log_rummy.de_score,
            user_play_log_rummy.tax,
            user_play_log_rummy.bonus,
            user_play_log_rummy.bet_time
           FROM user_play_log_rummy
          WHERE user_play_log_rummy.is_robot = 0 AND user_play_log_rummy.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_goldenflower.lognumber,
            user_play_log_goldenflower.agent_id,
            user_play_log_goldenflower.level_code,
            user_play_log_goldenflower.user_id,
            user_play_log_goldenflower.game_id,
            user_play_log_goldenflower.room_type,
            user_play_log_goldenflower.ya_score,
            user_play_log_goldenflower.valid_score,
            user_play_log_goldenflower.de_score,
            user_play_log_goldenflower.tax,
            user_play_log_goldenflower.bonus,
            user_play_log_goldenflower.bet_time
           FROM user_play_log_goldenflower
          WHERE user_play_log_goldenflower.is_robot = 0 AND user_play_log_goldenflower.bet_time > (now() - '3 days'::interval)
        UNION ALL
         SELECT user_play_log_pokdeng.lognumber,
            user_play_log_pokdeng.agent_id,
            user_play_log_pokdeng.level_code,
            user_play_log_pokdeng.user_id,
            user_play_log_pokdeng.game_id,
            user_play_log_pokdeng.room_type,
            user_play_log_pokdeng.ya_score,
            user_play_log_pokdeng.valid_score,
            user_play_log_pokdeng.de_score,
            user_play_log_pokdeng.tax,
            user_play_log_pokdeng.bonus,
            user_play_log_pokdeng.bet_time
           FROM user_play_log_pokdeng
          WHERE user_play_log_pokdeng.is_robot = 0 AND user_play_log_pokdeng.bet_time > (now() - '3 days'::interval)) tmp
  GROUP BY (date_trunc('day'::text, tmp.bet_time)), tmp.agent_id, tmp.level_code, tmp.game_id, tmp.room_type;