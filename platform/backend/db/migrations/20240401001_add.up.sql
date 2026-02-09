-- 新增輪盤 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (1010, 'dev01', '輪盤', 'roulette', 1, 1, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增輪盤遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10100, '輪盤新手房', 1, 1010, 0),
  (10101, '輪盤普通房', 1, 1010, 1),
  (10102, '輪盤高級房', 1, 1010, 2),
  (10103, '輪盤大師房', 1, 1010, 3);

-- 新增代理輪盤遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (1010);

-- 新增代理輪盤遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (1010);

-- 新增輪盤遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_roulette" (
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
  "wallet_ledger_id" varchar(100) NOT NULL DEFAULT '',
  "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  "kill_level" int2 NOT NULL DEFAULT -1,
  "real_players" int2 NOT NULL DEFAULT -1,
  CONSTRAINT "play_log_roulette_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_roulette_lognumber" ON "public"."user_play_log_roulette" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_roulette_betid" ON "public"."user_play_log_roulette" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_roulette_bet_time" ON "public"."user_play_log_roulette" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_roulette"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_roulette"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_roulette"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_roulette"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_roulette"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_roulette"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_roulette"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_roulette"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_roulette"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_roulette"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_roulette"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_roulette"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_roulette"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_roulette"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_roulette"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_roulette"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_roulette"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_roulette"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_roulette"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_roulette"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_roulette"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_roulette"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_roulette"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log_roulette"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_roulette"."jp_inject_water_score" IS 'jp注水分數';
COMMENT ON COLUMN "public"."user_play_log_roulette"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_roulette"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_roulette"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_roulette"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON TABLE "public"."user_play_log_roulette" IS '輪盤玩家遊戲記錄';

-- 新增遊戲基礎設定支援遊戲 --
UPDATE "public"."storage"
  SET "value" = "value" || '[1010]'
  WHERE "key"= 'GameSettingSupportInfo';