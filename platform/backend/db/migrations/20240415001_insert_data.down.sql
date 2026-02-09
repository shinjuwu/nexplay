DROP FUNCTION "public"."udf_update_agent";
CREATE OR REPLACE FUNCTION "public"."udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4, "_is_admin_user_role_changed" bool, "_wallet_conninfo" jsonb)
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

DROP FUNCTION "public"."udf_create_agent";
CREATE OR REPLACE FUNCTION "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_currency" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_agent_wallet_type" int2, "_agent_wallet_conninfo" jsonb, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4)
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
  "ret_admin_user" json;
  "ret_agent_games" json;
  "ret_agent_game_rooms" json;
BEGIN
  -- 新增代理 --
  INSERT INTO "public"."agent" ("name", "code", "level_code", "info", "secret_key",
    "aes_key", "md5_key", "currency", "ip_whitelist", "creator", "commission", "cooperation",
    "top_agent_id", "is_top_agent", "wallet_type", "wallet_conninfo")
    VALUES ("_agent_name", "_agent_code", "_agent_level_code", "_agent_info",
      "_agent_secret_key", "_agent_aes_key", "_agent_md5_key", "_agent_currency", "_agent_ip_whitelist",
      "_agent_creator", "_agent_commission", "_agent_cooperation", "_agent_top_agent_id", 
      "_agent_is_top_agent", "_agent_wallet_type", "_agent_wallet_conninfo")
    RETURNING "id", "name", "code", "secret_key", "aes_key", "md5_key", "currency", "commission", "info",
      "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id",
      "cooperation", "coin_limit", "coin_use", "ip_whitelist", "creator", "wallet_type", "wallet_conninfo" INTO
      "ret_agent_id", "ret_agent_name", "ret_agent_code", "ret_agent_secret_key",
      "ret_agent_aes_key", "ret_agent_md5_key", "ret_agent_currency", "ret_agent_commission", "ret_agent_info",
      "ret_agent_is_enabled", "ret_agent_disable_time", "ret_agent_update_time",
      "ret_agent_create_time", "ret_agent_is_top_agent", "ret_agent_top_agent_id",
      "ret_agent_cooperation", "ret_agent_coin_limit", "ret_agent_coin_use",
      "ret_agent_ip_whitelist", "ret_agent_creator", "ret_agent_wallet_type", "ret_agent_wallet_conninfo";

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
      'jackpot_end_time', extract(epoch from "ret_agent_jackpot_end_time") * 1000000
    ),
    'admin_user', "ret_admin_user",
    'agent_games', "ret_agent_games",
    'agent_game_rooms', "ret_agent_game_rooms"
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

ALTER TABLE "public"."agent"
  DROP COLUMN "lobby_switch_info";

-- 刪除代理取得好友房建房紀錄列表權限 --
UPDATE "public"."agent_permission" AS "ag1"
SET "permission" = jsonb_set("permission", '{list}', "ag3"."list", false)
FROM (
    SELECT "ag2"."id", jsonb_agg("ag2ple") AS "list"
    FROM "public"."agent_permission" AS "ag2", jsonb_array_elements("ag2"."permission"->'list') AS "ag2ple"
    WHERE "ag2ple"::int NOT IN (100325)
    GROUP BY "ag2"."id"
) AS "ag3"
WHERE "ag1"."id" = "ag3"."id";

-- 刪除取得好友房建房紀錄列表權限 --
DELETE FROM "public"."permission_list" WHERE feature_code IN (100325);

-- 刪除好友房德州紀錄表 --
DROP TABLE "public"."user_play_log_friendstexas";

-- 刪除代理好友房德州遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (5001)
);

-- 刪除代理好友房德州遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (5001);

-- 刪除好友房德州遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (5001);

-- 刪除新增好友房大廳及好友房德州遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (2, 5001);

-- 刪除老虎機類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_fruit777slot" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_megsharkslot" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_midasslot" DROP COLUMN "room_id";

-- 刪除電子遊戲類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_fruitslot" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_rcfishing" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_plinko" DROP COLUMN "room_id";

-- 刪除對戰類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_blackjack" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_sangong" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_bullbull" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_texas" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_rummy" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_goldenflower" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_pokdeng" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_catte" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_chinesepoker" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_okey" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_teenpatti" DROP COLUMN "room_id";

-- 刪除百人類型遊戲紀錄表房間編號(邀請碼)欄位 --
ALTER TABLE "public"."user_play_log_baccarat" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_fantan" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_colordisc" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_prawncrab" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_hundredsicbo" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_cockfight" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_dogracing" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_rocket" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_andarbahar" DROP COLUMN "room_id";
ALTER TABLE "public"."user_play_log_roulette" DROP COLUMN "room_id";

-- 刪除好友房紀錄表 --
DROP TABLE "public"."friend_room_log";
