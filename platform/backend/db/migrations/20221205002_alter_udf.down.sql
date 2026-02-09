CREATE FUNCTION "public"."udf_update_agent_wallet" ("_agent_id" int4, "_amount" numeric)
  RETURNS numeric AS $$
DECLARE
  "ret_amount" numeric;
Begin
  UPDATE "public"."agent_wallet"
    SET "amount" = "amount" + "_amount"
    WHERE "agent_id" = "_agent_id"
    RETURNING "amount" INTO "ret_amount";

  RETURN "ret_amount";
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE "public"."sp_create_agent_wallet_ledger" ("_id" varchar, "_agent_id" int4, "_before_coin" numeric, "_add_coin" numeric, "_after_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
  LANGUAGE plpgsql
AS $$
DECLARE
  "_changeset" jsonb;
BEGIN
  "_changeset" = jsonb_build_object(
    'before_coin', "_before_coin",
    'add_coin', "_add_coin",
    'after_coin', "_after_coin"
  );

  INSERT INTO "public"."agent_wallet_ledger" ("id", "agent_id", "changeset", "info", "kind", "creator")
    VALUES ("_id", "_agent_id", "_changeset", "_info", "_kind", "_creator");
END;
$$;

DROP FUNCTION "public"."udf_backend_update_agent_wallect";
CREATE FUNCTION "public"."udf_backend_update_agent_wallect" ("_id" varchar, "_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
  RETURNS boolean AS $$
DECLARE
  "ret_result" boolean := false;
  "_agent_wallet_amount" numeric := 0;
BEGIN
  SELECT "public"."udf_update_agent_wallet"("_agent_id", "_add_coin") INTO "_agent_wallet_amount";
  IF "_agent_wallet_amount" < 0 THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", -"_add_coin");
  ELSE
    CALL "public"."sp_create_agent_wallet_ledger" ("_id", "_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", "_kind", "_creator");
    "ret_result" = true;
  END IF;

  RETURN "ret_result";
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION "public"."udf_backend_update_agent_wallects";
CREATE FUNCTION "public"."udf_backend_update_agent_wallects" ("_from_id" varchar, "_from_agent_id" int4, "_to_id" varchar, "_to_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_creator" varchar)
  RETURNS boolean AS $$
DECLARE
  "ret_result" boolean := false;
  "_agent_wallet_amount" numeric := 0;
BEGIN
  IF "_add_coin" > 0 THEN
    SELECT "public"."udf_update_agent_wallet"("_from_agent_id", -"_add_coin") INTO "_agent_wallet_amount";
    IF "_agent_wallet_amount" < 0 THEN
      PERFORM "public"."udf_update_agent_wallet"("_from_agent_id", "_add_coin");
    ELSE
      CALL "public"."sp_create_agent_wallet_ledger" ("_from_id", "_from_agent_id", "_agent_wallet_amount" + "_add_coin", -"_add_coin", "_agent_wallet_amount", "_info", 4::int2, "_creator");

      SELECT "public"."udf_update_agent_wallet"("_to_agent_id", "_add_coin") INTO "_agent_wallet_amount";
      CALL "public"."sp_create_agent_wallet_ledger" ("_to_id", "_to_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", 3::int2, "_creator");

      "ret_result" = true;
    END IF;
  END IF;
  
  RETURN "ret_result";
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION "public"."udf_game_user_start_coin_in";
DROP FUNCTION "public"."udf_game_user_finish_coin_in";
DROP FUNCTION "public"."udf_game_user_start_coin_out";
DROP FUNCTION "public"."udf_game_user_finish_coin_out";

CREATE FUNCTION "public"."udf_game_user_create_transfer" ("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_update_agent_wallet" boolean, "_add_agent_wallet_amount" numeric)
  RETURNS json AS $$
DECLARE
  "_agent_wallet_amount" numeric;
  "_insert_wallet_ledger_count" bigint;
BEGIN
  IF "_add_agent_wallet_amount" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  IF "_update_agent_wallet" THEN
    SELECT "public"."udf_update_agent_wallet"("_agent_id", -"_add_agent_wallet_amount") INTO "_agent_wallet_amount";
    IF "_agent_wallet_amount" < 0 THEN
      PERFORM "public"."udf_update_agent_wallet"("_agent_id", "_add_agent_wallet_amount");

      RETURN json_build_object(
        'code', 2
      );
    END IF;
  END IF;

  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator", "request")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator", "_request")
    ON CONFLICT ("id") DO NOTHING;

  GET DIAGNOSTICS "_insert_wallet_ledger_count" = ROW_COUNT;
  IF "_insert_wallet_ledger_count" = 0 THEN
    RETURN json_build_object(
      'code', 3
    );
  END IF;

  RETURN json_build_object(
    'code', 0
  );
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION "public"."udf_game_user_complete_transfer" ("_id" varchar, "_changeset" jsonb, "_status" int2, "_error_code" int2, "_update_agent_wallet" boolean, "_agent_id" int4, "_add_agent_wallet_amount" numeric, "_update_game_user_sum_coin" boolean, "_user_id" int4, "_add_game_user_sum_coin_in" numeric, "_add_game_user_sum_coin_out" numeric)
  RETURNS json AS $$
BEGIN
  IF "_add_agent_wallet_amount" < 0 OR "_add_game_user_sum_coin_in" < 0 OR "_add_game_user_sum_coin_out" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  IF "_update_agent_wallet" THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", "_add_agent_wallet_amount");
  END IF;

  IF "_update_game_user_sum_coin" THEN
    UPDATE "public"."game_users"
      SET "sum_coin_in" = "sum_coin_in" + "_add_game_user_sum_coin_in",
        "sum_coin_out" = "sum_coin_out" + "_add_game_user_sum_coin_out",
        "update_time" = now()
      WHERE "id" = "_user_id";
  END IF;

  UPDATE "public"."wallet_ledger"
    SET "changeset" = "_changeset",
      "status" = "_status",
      "error_code" = "_error_code",
	  "update_time" = now()
    WHERE "id" = "_id";

  RETURN json_build_object(
    'code', 0
  );
END;
$$ LANGUAGE plpgsql;

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

  IF "ret_agent_cooperation" = 1 THEN
    INSERT INTO "public"."agent_wallet" ("agent_id")
      VALUES ("ret_agent_id");
  END IF;

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
