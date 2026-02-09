-- 新增安達巴哈 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (1009, 'dev01', '安达巴哈', 'andarbahar', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增安達巴哈遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10090, '安达巴哈新手房', 1, 1009, 0),
  (10091, '安达巴哈普通房', 1, 1009, 1),
  (10092, '安达巴哈高級房', 1, 1009, 2),
  (10093, '安达巴哈大師房', 1, 1009, 3);

-- 新增代理安達巴哈遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (1009);

-- 新增代理安達巴哈遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (1009);

-- 新增安達巴哈遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_andarbahar" (
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
  "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "play_log_andarbahar_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_andarbahar_lognumber" ON "public"."user_play_log_andarbahar" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_andarbahar_betid" ON "public"."user_play_log_andarbahar" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_andarbahar_bet_time" ON "public"."user_play_log_andarbahar" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_andarbahar"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."jp_inject_water_score" IS 'jp注水分數';
COMMENT ON TABLE "public"."user_play_log_andarbahar" IS '安达巴哈玩家遊戲記錄';