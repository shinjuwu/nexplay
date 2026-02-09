
-- 新增巨齒鯊 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (4002, 'dev01', '巨齒鯊', 'megsharkslot', 1, 4, 1, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增巨齒鯊遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (40020, '巨齒鯊單一房', 1, 4002, 0);

-- 新增代理巨齒鯊遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (4002);

-- 新增代理巨齒鯊遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (4002);

-- 新增巨齒鯊遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_megsharkslot" (
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
  CONSTRAINT "play_log_megsharkslot_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_megsharkslot_lognumber" ON "public"."user_play_log_megsharkslot" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_megsharkslot_betid" ON "public"."user_play_log_megsharkslot" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_megsharkslot_bet_time" ON "public"."user_play_log_megsharkslot" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."jp_inject_water_score" IS 'jp注水分數';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON TABLE "public"."user_play_log_megsharkslot" IS '巨齒鯊玩家遊戲記錄';

-- 新增遊戲基礎設定支援遊戲 --
UPDATE "public"."storage"
  SET "value" = "value" || '[4002]'
  WHERE "key"= 'GameSettingSupportInfo';