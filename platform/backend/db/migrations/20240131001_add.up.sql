-- 新增印度炸金花 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (2011, 'dev01', '印度炸金花', 'teenpatti', 1, 2, 4, 1, 1, 'http://172.30.0.152/client/vue/apps');

-- 新增印度炸金花遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (20110, '印度炸金花新手房', 1, 2011, 0),
  (20111, '印度炸金花普通房', 1, 2011, 1),
  (20112, '印度炸金花高級房', 1, 2011, 2),
  (20113, '印度炸金花大師房', 1, 2011, 3);

-- 新增代理印度炸金花遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (2011);

-- 新增代理印度炸金花遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (2011);

-- 新增印度炸金花遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_teenpatti" (
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
  CONSTRAINT "play_log_teenpatti_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_teenpatti_lognumber" ON "public"."user_play_log_teenpatti" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_teenpatti_betid" ON "public"."user_play_log_teenpatti" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_teenpatti_bet_time" ON "public"."user_play_log_teenpatti" ("bet_time" DESC NULLS LAST); 

COMMENT ON COLUMN "public"."user_play_log_teenpatti"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."jp_inject_water_score" IS 'jp注水分數';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON TABLE "public"."user_play_log_teenpatti" IS '印度炸金花玩家遊戲記錄';

-- 新增遊戲基礎設定支援遊戲 --
UPDATE "public"."storage"
  SET "value" = "value" || '[2011]'
  WHERE "key"= 'GameSettingSupportInfo';