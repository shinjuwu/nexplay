-- 新增好友房紀錄表 --
CREATE TABLE "public"."friend_room_log" (
  "id" varchar(30) NOT NULL DEFAULT '',
  "agent_id" int4 NOT NULL DEFAULT -1,
  "level_code" varchar(128) NOT NULL DEFAULT '',
  "game_id" int4 NOT NULL DEFAULT -1,
  "room_id" varchar(16) NOT NULL DEFAULT '',
  "user_id" int4 NOT NULL DEFAULT -1,
  "username" varchar(100) NOT NULL DEFAULT '',
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  "taxpercent" numeric(20,4) NOT NULL DEFAULT 0,
  "detail" jsonb NOT NULL DEFAULT '{}',
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "end_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "friend_room_log_pkey" PRIMARY KEY ("id")
)
;

CREATE INDEX "idx_friend_room_log_agent_id" ON "public"."friend_room_log" ("agent_id" ASC NULLS LAST);
CREATE INDEX "idx_friend_room_log_level_code" ON "public"."friend_room_log" ("level_code" ASC NULLS LAST);
CREATE INDEX "idx_friend_room_log_game_id" ON "public"."friend_room_log" ("game_id" ASC NULLS LAST);
CREATE INDEX "idx_friend_room_log_room_id" ON "public"."friend_room_log" ("room_id" DESC NULLS LAST);
CREATE INDEX "idx_friend_room_log_create_time" ON "public"."friend_room_log" ("create_time" DESC NULLS LAST);

COMMENT ON COLUMN "public"."friend_room_log"."id" IS '訂單號';
COMMENT ON COLUMN "public"."friend_room_log"."agent_id" IS '房主代理識別號';
COMMENT ON COLUMN "public"."friend_room_log"."level_code" IS '房主代理層級碼';
COMMENT ON COLUMN "public"."friend_room_log"."game_id" IS '遊戲識別號';
COMMENT ON COLUMN "public"."friend_room_log"."room_id" IS '房間編號(邀請碼)';
COMMENT ON COLUMN "public"."friend_room_log"."user_id" IS '房主識別號';
COMMENT ON COLUMN "public"."friend_room_log"."username" IS '房主名稱';
COMMENT ON COLUMN "public"."friend_room_log"."tax" IS '房間總抽水';
COMMENT ON COLUMN "public"."friend_room_log"."taxpercent" IS '房間抽水%數';
COMMENT ON COLUMN "public"."friend_room_log"."create_time" IS '房間創建時間';
COMMENT ON COLUMN "public"."friend_room_log"."end_time" IS '房間結束時間';

-- 新增百人類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_baccarat"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_baccarat_room_id" ON "public"."user_play_log_baccarat" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_fantan"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_fantan"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_fantan_room_id" ON "public"."user_play_log_fantan" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_colordisc"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_colordisc_room_id" ON "public"."user_play_log_colordisc" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_prawncrab"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_prawncrab_room_id" ON "public"."user_play_log_prawncrab" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_hundredsicbo"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_hundredsicbo_room_id" ON "public"."user_play_log_hundredsicbo" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_cockfight"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_cockfight_room_id" ON "public"."user_play_log_cockfight" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_dogracing"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_dogracing_room_id" ON "public"."user_play_log_dogracing" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_rocket"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_rocket"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_rocket_room_id" ON "public"."user_play_log_rocket" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_andarbahar"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_andarbahar"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_andarbahar_room_id" ON "public"."user_play_log_andarbahar" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_roulette"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_roulette"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_roulette_room_id" ON "public"."user_play_log_roulette" ("room_id" DESC NULLS LAST);

-- 新增對戰類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_blackjack"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_blackjack_room_id" ON "public"."user_play_log_blackjack" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_sangong"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_sangong"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_sangong_room_id" ON "public"."user_play_log_sangong" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_bullbull"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_bullbull_room_id" ON "public"."user_play_log_bullbull" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_texas"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_texas"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_texas_room_id" ON "public"."user_play_log_texas" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_rummy"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_rummy"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_rummy_room_id" ON "public"."user_play_log_rummy" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_goldenflower"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_goldenflower_room_id" ON "public"."user_play_log_goldenflower" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_pokdeng"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_pokdeng_room_id" ON "public"."user_play_log_pokdeng" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_catte"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_catte"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_catte_room_id" ON "public"."user_play_log_catte" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_chinesepoker"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_chinesepoker_room_id" ON "public"."user_play_log_chinesepoker" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_okey"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_okey"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_okey_room_id" ON "public"."user_play_log_okey" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_teenpatti"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_teenpatti_room_id" ON "public"."user_play_log_teenpatti" ("room_id" DESC NULLS LAST);

-- 新增電子遊戲類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_fruitslot"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_fruitslot_room_id" ON "public"."user_play_log_fruitslot" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_rcfishing"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_rcfishing_room_id" ON "public"."user_play_log_rcfishing" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_plinko"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_plinko"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_plinko_room_id" ON "public"."user_play_log_plinko" ("room_id" DESC NULLS LAST);

-- 新增老虎機類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_fruit777slot"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_fruit777slot_room_id" ON "public"."user_play_log_fruit777slot" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_megsharkslot"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_megsharkslot_room_id" ON "public"."user_play_log_megsharkslot" ("room_id" DESC NULLS LAST);

ALTER TABLE "public"."user_play_log_midasslot"
  ADD COLUMN "room_id" varchar(16) NOT NULL DEFAULT '';
COMMENT ON COLUMN "public"."user_play_log_midasslot"."room_id" IS '房間編號(邀請碼)';
CREATE INDEX "idx_user_play_log_midasslot_room_id" ON "public"."user_play_log_midasslot" ("room_id" DESC NULLS LAST);

-- 新增好友房大廳及德州資料 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (2, 'dev02', '好友房大廳', 'friendsroom', 1, 0, 0, 0, 1, 'http://172.30.0.166/client/vue'),
  (5001, 'dev02', '好友房德撲', 'friendstexas', 1, 5, 1, 1, 1, 'http://172.30.0.166/client/vue/apps');

-- 新增好友房德州遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (50010, '德撲好友房', 1, 5001, 0);

-- 新增代理好友房德州遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" IN (5001);

-- 新增代理好友房德州遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" IN (5001);

-- 新增好友房德州遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_friendstexas" (
  "bet_id" varchar(30) NOT NULL,
  "lognumber" varchar(100) NOT NULL,
  "agent_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) NOT NULL,
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
  "room_id" varchar(16) NOT NULL DEFAULT '',
  CONSTRAINT "play_log_friendstexas_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

CREATE INDEX "idx_user_play_log_friendstexas_lognumber" ON "public"."user_play_log_friendstexas" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_friendstexas_betid" ON "public"."user_play_log_friendstexas" ("bet_id" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_friendstexas_bet_time" ON "public"."user_play_log_friendstexas" ("bet_time" DESC NULLS LAST); 
CREATE INDEX "idx_user_play_log_friendstexas_room_id" ON "public"."user_play_log_friendstexas" ("room_id" DESC NULLS LAST);

COMMENT ON COLUMN "public"."user_play_log_friendstexas"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."username" IS '代理用戶名稱';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."game_id" IS '遊戲識別號(請參照game_setting表)';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."kill_type" IS '殺放狀態';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."bonus" IS '紅利';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."jp_inject_water_score" IS 'jp注水分數';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_friendstexas"."room_id" IS '房間編號(邀請碼)';
COMMENT ON TABLE "public"."user_play_log_friendstexas" IS '好友房德州玩家遊戲記錄';

-- 新增取得好友房建房紀錄列表權限 --
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100325, '取得好友房建房紀錄列表', '/api/v1/record/getfriendroomloglist', 't', '後台使用', 't', 0);

-- 新增代理取得好友房建房紀錄列表權限 --
UPDATE "public"."agent_permission"
  SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100325]'::jsonb, false)
  WHERE "agent_id" = -1;

ALTER TABLE "public"."agent"
  ADD COLUMN "lobby_switch_info" int2 NOT NULL DEFAULT 0;
UPDATE "public"."agent"
  SET "lobby_switch_info" = 1;

DROP FUNCTION "public"."udf_create_agent";
CREATE OR REPLACE FUNCTION "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_currency" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_agent_wallet_type" int2, "_agent_wallet_conninfo" jsonb, "_agent_lobby_switch_info" int2, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "ret_agent_id" int4;
  "ret_agent_name" varchar;
  "ret_agent_code" varchar;
  "ret_agent_secret_key" varchar;
  "ret_agent_aes_key" varchar;
  "ret_agent_md5_key" varchar;
  "ret_agent_currency" varchar;
  "ret_agent_commission" int4;
  "ret_agent_info" varchar;
  "ret_agent_is_enabled" int4;
  "ret_agent_disable_time" timestamp;
  "ret_agent_update_time" timestamp;
  "ret_agent_create_time" timestamp;
  "ret_agent_is_top_agent" bool;
  "ret_agent_top_agent_id" int4;
  "ret_agent_cooperation" int2;
  "ret_agent_coin_limit" decimal;
  "ret_agent_coin_use" decimal;
  "ret_agent_level_code" varchar;
  "ret_agent_ip_whitelist" jsonb;
  "ret_agent_creator" varchar;
	"ret_agent_wallet_type" int2;
	"ret_agent_wallet_conninfo" jsonb;
	"ret_agent_jackpot_status" int4;
  "ret_agent_jackpot_start_time" timestamp;
  "ret_agent_jackpot_end_time" timestamp;
  "ret_agent_lobby_switch_info" int2;
  "ret_admin_user" json;
  "ret_agent_games" json;
  "ret_agent_game_rooms" json;
BEGIN
  -- 新增代理 --
  INSERT INTO "public"."agent" ("name", "code", "level_code", "info", "secret_key",
    "aes_key", "md5_key", "currency", "ip_whitelist", "creator", "commission", "cooperation",
    "top_agent_id", "is_top_agent", "wallet_type", "wallet_conninfo", "lobby_switch_info")
    VALUES ("_agent_name", "_agent_code", "_agent_level_code", "_agent_info",
      "_agent_secret_key", "_agent_aes_key", "_agent_md5_key", "_agent_currency", "_agent_ip_whitelist",
      "_agent_creator", "_agent_commission", "_agent_cooperation", "_agent_top_agent_id", 
      "_agent_is_top_agent", "_agent_wallet_type", "_agent_wallet_conninfo", "_agent_lobby_switch_info")
    RETURNING "id", "name", "code", "secret_key", "aes_key", "md5_key", "currency", "commission", "info",
      "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id",
      "cooperation", "coin_limit", "coin_use", "ip_whitelist", "creator", "wallet_type", "wallet_conninfo",
      "lobby_switch_info" INTO
      "ret_agent_id", "ret_agent_name", "ret_agent_code", "ret_agent_secret_key",
      "ret_agent_aes_key", "ret_agent_md5_key", "ret_agent_currency", "ret_agent_commission", "ret_agent_info",
      "ret_agent_is_enabled", "ret_agent_disable_time", "ret_agent_update_time",
      "ret_agent_create_time", "ret_agent_is_top_agent", "ret_agent_top_agent_id",
      "ret_agent_cooperation", "ret_agent_coin_limit", "ret_agent_coin_use",
      "ret_agent_ip_whitelist", "ret_agent_creator", "ret_agent_wallet_type", "ret_agent_wallet_conninfo",
      "ret_agent_lobby_switch_info";

  -- 更新上級代理子代理數量
  UPDATE "public"."agent"
    SET "child_agent_count" = "child_agent_count" + 1
    WHERE "id" = "_agent_top_agent_id";

  -- 更新 level code --
  UPDATE "public"."agent"
    SET "level_code" = "level_code" || LPAD(to_hex("ret_agent_id"), 4, '0')
    WHERE "id" = "ret_agent_id"
    RETURNING "level_code" INTO "ret_agent_level_code";

  -- 更新 jackpot time --
  SELECT "jackpot_status", "jackpot_start_time", "jackpot_end_time" INTO "ret_agent_jackpot_status", "ret_agent_jackpot_start_time", "ret_agent_jackpot_end_time"  FROM "public"."agent" WHERE "id" = "_agent_top_agent_id";
  UPDATE "public"."agent"
    SET "jackpot_status" = "ret_agent_jackpot_status",
      "jackpot_start_time" = "ret_agent_jackpot_start_time",
      "jackpot_end_time" = "ret_agent_jackpot_end_time"
    WHERE "id" = "ret_agent_id";
  
  -- 新增錢包 --
  INSERT INTO "public"."agent_wallet" ("agent_id")
    VALUES ("ret_agent_id");

  -- 新增 admin user --
  SELECT "public"."udf_create_admin_user" ("ret_agent_id", "_admin_user_username", "_admin_user_password",
    "_admin_user_nickname", "_admin_user_account_type", "_admin_user_is_readonly", false, "_admin_user_role",
    "_admin_user_info") INTO "ret_admin_user";

  -- 新增遊戲設定 --
  INSERT INTO "public"."agent_game" ("agent_id", "game_id", "state")
    SELECT "ret_agent_id" AS "agent_id", "game_id", "state"
      FROM "public"."agent_game"
      WHERE "agent_id" = "ret_agent_top_agent_id";
    
  SELECT json_agg("agent_games") INTO "ret_agent_games"
    FROM (
    SELECT "agent_id", "game_id", "state"
      FROM "public"."agent_game"
      WHERE "agent_id" = "ret_agent_id"
  ) AS "agent_games";

  -- 新增遊戲房間設定 --
  INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id", "state")
    SELECT "ret_agent_id" AS "agent_id", "game_room_id", "state"
      FROM "public"."agent_game_room"
      WHERE "agent_id" = "ret_agent_top_agent_id";
    
  SELECT json_agg("agent_game_rooms") INTO "ret_agent_game_rooms"
    FROM (
    SELECT "agent_id", "game_room_id", "state"
      FROM "public"."agent_game_room"
      WHERE "agent_id" = "ret_agent_id"
  ) AS "agent_game_rooms";

  RETURN json_build_object(
    'agent', json_build_object(
      'id', "ret_agent_id",
      'name', "ret_agent_name",
      'code', "ret_agent_code",
      'level_code', "ret_agent_level_code",
      'secret_key', "ret_agent_secret_key",
      'aes_key', "ret_agent_aes_key",
      'md5_key', "ret_agent_md5_key",
      'currency', "ret_agent_currency",
      'commission', "ret_agent_commission",
      'info', "ret_agent_info",
      'is_enabled', "ret_agent_is_enabled",
      'disable_time', extract(epoch from "ret_agent_disable_time") * 1000000,
      'update_time', extract(epoch from "ret_agent_update_time") * 1000000,
      'create_time', extract(epoch from "ret_agent_create_time") * 1000000,
      'is_top_agent', "ret_agent_is_top_agent",
      'top_agent_id', "ret_agent_top_agent_id",
      'cooperation', "ret_agent_cooperation",
      'coin_limit', "ret_agent_coin_limit",
      'coin_use', "ret_agent_coin_use",
      'ip_whitelist', "ret_agent_ip_whitelist",
      'creator', "ret_agent_creator",
			'wallet_type',  "ret_agent_wallet_type",
			'wallet_conninfo',  "ret_agent_wallet_conninfo",
      'jackpot_status', "ret_agent_jackpot_status",
      'jackpot_start_time', extract(epoch from "ret_agent_jackpot_start_time") * 1000000,
      'jackpot_end_time', extract(epoch from "ret_agent_jackpot_end_time") * 1000000,
      'lobby_switch_info', "ret_agent_lobby_switch_info"
    ),
    'admin_user', "ret_admin_user",
    'agent_games', "ret_agent_games",
    'agent_game_rooms', "ret_agent_game_rooms"
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

DROP FUNCTION "public"."udf_update_agent";
CREATE OR REPLACE FUNCTION "public"."udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4, "_is_admin_user_role_changed" bool, "_wallet_conninfo" jsonb, "_agent_lobby_switch_info" int2)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "agent_level_code" varchar;
  "ret_agent_update_time" timestamp;
  "ret_admin_user" json;
BEGIN
  UPDATE "public"."agent"
    SET "name" = "_agent_name",
      "info" = "_agent_info",
      "commission" = "_agent_commission",
			"wallet_conninfo" = "_wallet_conninfo",
      "update_time" = now()
    WHERE "id" = "_agent_id"
    RETURNING "update_time", "level_code" INTO "ret_agent_update_time", "agent_level_code";

  SELECT "public"."udf_update_admin_user"("_agent_id", "_admin_user_username", "_admin_user_role",
    "_admin_user_info", "_admin_user_is_enabled") INTO "ret_admin_user";

  IF "_is_admin_user_role_changed" THEN
    CALL "public"."usp_update_agents_permission" ("_admin_user_role", "agent_level_code");
  END IF;

  UPDATE "public"."agent"
    SET "lobby_switch_info" = "_agent_lobby_switch_info"
    WHERE "level_code" LIKE "agent_level_code" || '%';

  RETURN json_build_object(
    'agent', json_build_object(
      'update_time', extract(epoch from "ret_agent_update_time") * 1000000
    ),
    'admin_user', "ret_admin_user"
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;