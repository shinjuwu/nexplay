CREATE TABLE "public"."agent_wallet" (
  "agent_id" int4 NOT NULL DEFAULT 0,
  "amount" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "agent_wallet_pkey" PRIMARY KEY ("agent_id")
);

INSERT INTO "public"."agent_wallet" ("agent_id")
  SELECT "id"
  FROM "public"."agent"
  WHERE "cooperation" = 1;

CREATE TABLE "public"."agent_wallet_ledger" (
  "id" varchar(100) NOT NULL,
  "agent_id" int4 NOT NULL DEFAULT 0,
  "changeset" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) NOT NULL DEFAULT '',
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "kind" int2 NOT NULL DEFAULT 0,
  "creator" varchar(20) NOT NULL,
  CONSTRAINT "agent_wallet_ledger_pkey" PRIMARY KEY ("id")
);

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

CREATE FUNCTION "public"."udf_create_order_id" ("_type" int, "_salt" varchar)
  RETURNS varchar AS $$
DECLARE
  "ret_order_id" varchar;
  "_c" int := 0;
  "_ch" varchar;
BEGIN
  "ret_order_id" = "_type" || to_char(now(), 'YYYYMMDDHHMISSMS');

  FOREACH "_ch" IN ARRAY regexp_split_to_array(md5("_salt"), '') LOOP
    IF "_ch" >= '0' AND "_ch" <= '9' THEN
      "ret_order_id" = "ret_order_id" || "_ch";
      "_c" = "_c" + 1;

      IF "_c" = 6 THEN
        EXIT;
      END IF;
    END IF;
  END LOOP;

  IF "_c" < 6 THEN
    WHILE "_c" < 6 LOOP
      "_c" = "_c" + 1;
      "ret_order_id" = "ret_order_id" || '0';
    END LOOP;  
  END IF;

  RETURN "ret_order_id";
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE "public"."sp_create_agent_wallet_ledger" ("_agent_id" int4, "_before_coin" numeric, "_add_coin" numeric, "_after_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
  LANGUAGE plpgsql
AS $$
DECLARE
  "_id" varchar;
  "_salt" varchar;
  "_changeset" jsonb;
BEGIN
  "_changeset" = jsonb_build_object(
    'before_coin', "_before_coin",
    'add_coin', "_add_coin",
    'after_coin', "_after_coin"
  );

  "_salt" = "_kind" || '-' || "_agent_id" || '-' || "_creator" || '-' || to_char(now(), 'YYYYMMDDHHMISSMS');
  SELECT "public"."udf_create_order_id" ("_kind", "_salt") INTO "_id";

  INSERT INTO "public"."agent_wallet_ledger" ("id", "agent_id", "changeset", "info", "kind", "creator")
    VALUES ("_id", "_agent_id", "_changeset", "_info", "_kind", "_creator");
END;
$$;

--直接處理給管理者使用--
CREATE FUNCTION "public"."udf_backend_update_agent_wallect" ("_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
  RETURNS boolean AS $$
DECLARE
  "ret_result" boolean := false;
  "_agent_wallet_amount" numeric := 0;
BEGIN
  SELECT "public"."udf_update_agent_wallet"("_agent_id", "_add_coin") INTO "_agent_wallet_amount";
  IF "_agent_wallet_amount" < 0 THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", -"_add_coin");
  ELSE
    CALL "public"."sp_create_agent_wallet_ledger" ("_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", "_kind", "_creator");
    "ret_result" = true;
  END IF;

  RETURN "ret_result";
END;
$$ LANGUAGE plpgsql;

--非管理者間互相轉移使用--
--錢從From(下分)到To(上分)--
CREATE FUNCTION "public"."udf_backend_update_agent_wallects" ("_from_agent_id" int4, "_to_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_creator" varchar)
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
      CALL "public"."sp_create_agent_wallet_ledger" ("_from_agent_id", "_agent_wallet_amount" + "_add_coin", -"_add_coin", "_agent_wallet_amount", "_info", 4::int2, "_creator");

      SELECT "public"."udf_update_agent_wallet"("_to_agent_id", "_add_coin") INTO "_agent_wallet_amount";
      CALL "public"."sp_create_agent_wallet_ledger" ("_to_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", 3::int2, "_creator");

      "ret_result" = true;
    END IF;
  END IF;
  
  RETURN "ret_result";
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
