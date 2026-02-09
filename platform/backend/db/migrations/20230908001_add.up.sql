-- 新增彈珠檯 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (2009, 'dev01', '十三水', 'chinesepoker', 1, 2, 8, 1, 1, 'http://172.30.0.152/client/vue/apps'),
  (3003, 'dev01', '彈珠檯', 'plinko', 1, 3, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增彈珠檯遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (20090, '十三水新手場', 1, 2009, 0),
  (20091, '十三水普通場', 1, 2009, 1),
  (20092, '十三水高手場', 1, 2009, 2),
  (20093, '十三水大師場', 1, 2009, 3),
  (20094, '十三水初级场', 1, 2009, 4),
  (20095, '十三水中级场', 1, 2009, 5),
  (20096, '十三水高级场', 1, 2009, 6),
  (20097, '十三水至尊场', 1, 2009, 7),
  (30030, '彈珠檯新手房', 1, 3003, 0),
  (30031, '彈珠檯普通房', 1, 3003, 1),
  (30032, '彈珠檯高級房', 1, 3003, 2),
  (30033, '彈珠檯大師房', 1, 3003, 3);

-- 新增代理十三水、彈珠檯遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (2009, 3003);

-- 新增代理十三水、彈珠檯遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (2009, 3003);

-- 新增遊戲十三水紀錄表 --
CREATE TABLE "public"."user_play_log_chinesepoker" (
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
  CONSTRAINT "play_log_chinesepoker_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_chinesepoker_lognumber" ON "public"."user_play_log_chinesepoker" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_chinesepoker_betid" ON "public"."user_play_log_chinesepoker" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_chinesepoker_bet_time" ON "public"."user_play_log_chinesepoker" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_chinesepoker" IS '十三水玩家遊戲記錄';

-- 新增彈珠檯遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_plinko" (
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
  CONSTRAINT "play_log_plinko_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_plinko_lognumber" ON "public"."user_play_log_plinko" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_plinko_betid" ON "public"."user_play_log_plinko" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_plinko_bet_time" ON "public"."user_play_log_plinko" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_plinko"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_plinko"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_plinko"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_plinko"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_plinko"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_plinko"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_plinko"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_plinko"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_plinko"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_plinko"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_plinko"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_plinko"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_plinko"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_plinko"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_plinko"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_plinko"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_plinko"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_plinko"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_plinko"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_plinko"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_plinko"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_plinko"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_plinko"."bonus" IS '紅利';
COMMENT ON TABLE "public"."user_play_log_plinko" IS '彈珠檯玩家遊戲記錄';
