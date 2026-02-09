DROP FUNCTION "public"."udf_create_agent_games";
CREATE FUNCTION "public"."udf_create_agent_games"("_agent_id" int4, "_top_agent_id" int4)
  RETURNS json AS $$
DECLARE
  "ret_agent_games" json;
BEGIN
  INSERT INTO "public"."agent_game" ("agent_id", "game_id", "state")
    SELECT "_agent_id" AS "agent_id", "game_id", "state"
      FROM "public"."agent_game"
      WHERE "agent_id" = "_top_agent_id";
	  
  SELECT json_agg("agent_games") INTO "ret_agent_games"
    FROM (
	  SELECT "agent_id", "game_id", "state"
		  FROM "public"."agent_game"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_games";
  
  RETURN "ret_agent_games";
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION "public"."udf_create_agent_game_rooms";
CREATE FUNCTION "public"."udf_create_agent_game_rooms"("_agent_id" int4, "_top_agent_id" int4)
  RETURNS json AS $$
DECLARE
  "ret_agent_game_rooms" json;
BEGIN
  INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id", "state")
    SELECT "_agent_id" AS "agent_id", "game_room_id", "state"
      FROM "public"."agent_game_room"
      WHERE "agent_id" = "_top_agent_id";
	  
  SELECT json_agg("agent_game_rooms") INTO "ret_agent_game_rooms"
    FROM (
	  SELECT "agent_id", "game_room_id", "state"
		  FROM "public"."agent_game_room"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_game_rooms";
  
  RETURN "ret_agent_game_rooms";
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION "public"."udf_create_agent";
CREATE FUNCTION "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4)
  RETURNS json AS $$
DECLARE
    "ret_agent_id" int4;
    "ret_agent_name" varchar;
    "ret_agent_code" varchar;
    "ret_agent_secret_key" varchar;
    "ret_agent_aes_key" varchar;
    "ret_agent_md5_key" varchar;
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
    "ret_admin_user" json;
    "ret_agent_games" json;
    "ret_agent_game_rooms" json;
BEGIN
    INSERT INTO "public"."agent" ("name", "code", "level_code", "info", "secret_key",
        "aes_key", "md5_key", "ip_whitelist", "creator", "commission", "cooperation",
        "top_agent_id", "is_top_agent")
        VALUES ("_agent_name", "_agent_code", "_agent_level_code", "_agent_info",
            "_agent_secret_key", "_agent_aes_key", "_agent_md5_key", "_agent_ip_whitelist",
            "_agent_creator", "_agent_commission", "_agent_cooperation", "_agent_top_agent_id", 
            "_agent_is_top_agent")
        RETURNING "id", "name", "code", "secret_key", "aes_key", "md5_key", "commission", "info",
            "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id",
            "cooperation", "coin_limit", "coin_use", "ip_whitelist", "creator" INTO
            "ret_agent_id", "ret_agent_name", "ret_agent_code", "ret_agent_secret_key",
            "ret_agent_aes_key", "ret_agent_md5_key", "ret_agent_commission", "ret_agent_info",
            "ret_agent_is_enabled", "ret_agent_disable_time", "ret_agent_update_time",
            "ret_agent_create_time", "ret_agent_is_top_agent", "ret_agent_top_agent_id",
            "ret_agent_cooperation", "ret_agent_coin_limit", "ret_agent_coin_use",
            "ret_agent_ip_whitelist", "ret_agent_creator";

    UPDATE "public"."agent"
        SET "level_code" = "level_code" || LPAD(to_hex("ret_agent_id"), 4, '0')
        WHERE "id" = "ret_agent_id"
        RETURNING "level_code" INTO "ret_agent_level_code";

    SELECT "public"."udf_create_admin_user" ("ret_agent_id", "_admin_user_username", "_admin_user_password",
        "_admin_user_nickname", "_admin_user_account_type", "_admin_user_is_readonly", false, "_admin_user_role",
        "_admin_user_info") INTO "ret_admin_user";

    SELECT "public"."udf_create_agent_games" ("ret_agent_id", "ret_agent_top_agent_id") INTO "ret_agent_games";

    SELECT "public"."udf_create_agent_game_rooms" ("ret_agent_id", "ret_agent_top_agent_id") INTO "ret_agent_game_rooms";


    RETURN json_build_object(
        'agent', json_build_object(
            'id', "ret_agent_id",
            'name', "ret_agent_name",
            'code', "ret_agent_code",
            'level_code', "ret_agent_level_code",
            'secret_key', "ret_agent_secret_key",
            'aes_key', "ret_agent_aes_key",
            'md5_key', "ret_agent_md5_key",
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
            'creator', "ret_agent_creator"
        ),
        'admin_user', "ret_admin_user",
        'agent_games', "ret_agent_games",
        'agent_game_rooms', "ret_agent_game_rooms"
    );
END;
$$ LANGUAGE plpgsql;
