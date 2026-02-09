ALTER TABLE "public"."agent" 
  ADD COLUMN "child_agent_count" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "jackpot_status" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "jackpot_start_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  ADD COLUMN "jackpot_end_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone;

COMMENT ON COLUMN "public"."agent"."child_agent_count" IS '子代理數量';
COMMENT ON COLUMN "public"."agent"."jackpot_status" IS 'jackpot狀態';
COMMENT ON COLUMN "public"."agent"."jackpot_start_time" IS 'jackpot開始時間';
COMMENT ON COLUMN "public"."agent"."jackpot_end_time" IS 'jackpot結束時間';

UPDATE "public"."agent" AS "a"
  SET "child_agent_count" = "tmp"."child_agent_count"
  FROM (
    SELECT
      "top"."id" AS "id",
      COUNT("child"."id") AS "child_agent_count"
   FROM "public"."agent" AS "top"
   INNER JOIN "public"."agent" AS "child" ON "top"."id" = "child"."top_agent_id"
   GROUP BY "top"."id"
  ) AS "tmp"
  WHERE "a"."id" = "tmp"."id";

DROP FUNCTION IF EXISTS "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_currency" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4);
CREATE OR REPLACE FUNCTION "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_currency" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4)
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
    "top_agent_id", "is_top_agent")
    VALUES ("_agent_name", "_agent_code", "_agent_level_code", "_agent_info",
      "_agent_secret_key", "_agent_aes_key", "_agent_md5_key", "_agent_currency", "_agent_ip_whitelist",
      "_agent_creator", "_agent_commission", "_agent_cooperation", "_agent_top_agent_id", 
      "_agent_is_top_agent")
    RETURNING "id", "name", "code", "secret_key", "aes_key", "md5_key", "currency", "commission", "info",
      "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id",
      "cooperation", "coin_limit", "coin_use", "ip_whitelist", "creator" INTO
      "ret_agent_id", "ret_agent_name", "ret_agent_code", "ret_agent_secret_key",
      "ret_agent_aes_key", "ret_agent_md5_key", "ret_agent_currency", "ret_agent_commission", "ret_agent_info",
      "ret_agent_is_enabled", "ret_agent_disable_time", "ret_agent_update_time",
      "ret_agent_create_time", "ret_agent_is_top_agent", "ret_agent_top_agent_id",
      "ret_agent_cooperation", "ret_agent_coin_limit", "ret_agent_coin_use",
      "ret_agent_ip_whitelist", "ret_agent_creator";

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

INSERT INTO "public"."storage" ("id", "key", "value", "readonly", "create_time", "update_time") VALUES ('4f0b1215-0070-4e6d-b91d-2f4f80c92e0f', 'JackpotSetting', '{"jackpot_switch": true}', 'f', '2023-09-05 05:37:53.108477+00', '2023-09-05 05:37:53.108477+00');

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100105, '此接口用來創建jackpot紀錄(不做參數檢查)', '/api/v1/intercom/createjackpotrecord', 't', '遊戲SERVER串接使用', 'f', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100307, '取得總代理jackpot設定列表(只有開發商(營運商)可使用)', '/api/v1/jackpot/getagentjackpotlist', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100308, '設定總代理jackpot(只有開發商(營運商)可使用)', '/api/v1/jackpot/setagentjackpot', 't', '後台使用', 't', 2);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100309, '取得平台jackpot設定(只有開發商(營運商)可使用)', '/api/v1/jackpot/getjackpotsetting', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100310, '設定平台jackpot設定(只有開發商(營運商)可使用)', '/api/v1/jackpot/setjackpotsetting', 't', '後台使用', 't', 2);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100311, '取得jackpot代幣列表(只有開發商(營運商)可使用)', '/api/v1/jackpot/getjackpottokenlist', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100312, '建立jackpot代幣(只有開發商(營運商)可使用)', '/api/v1/jackpot/createjackpottoken', 't', '後台使用', 't', 1);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100313, '取得jackpot紀錄列表', '/api/v1/jackpot/getjackpotlist', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100314, '取得jackpot獎池資訊(只有開發商(營運商)可使用)', '/api/v1/jackpot/getjackpotpooldata', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100315, 'server同步jackpot資訊(只有開發商(營運商)可使用)', '/api/v1/jackpot/notifygameserveragentjackpotinfo', 't', '後台使用', 't', 2);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100316, '取得jackpot玩家貢獻度(只有開發商(營運商)可使用)', '/api/v1/jackpot/getjackpotleaderboard', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100317, '取得指定總代理jackpot設定(只有開發商(營運商)可使用)', '/api/v1/jackpot/getagentjackpot', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100307,100308,100309,100310,100311,100312,100314,100315,100316,100317]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100313]'::jsonb, false)
WHERE "agent_id" = -1;