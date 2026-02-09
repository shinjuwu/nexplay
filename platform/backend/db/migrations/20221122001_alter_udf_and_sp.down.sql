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

DROP PROCEDURE "public"."sp_create_agent_wallet_ledger";
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

DROP FUNCTION "public"."udf_backend_update_agent_wallect";
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

DROP FUNCTION "public"."udf_backend_update_agent_wallects";
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
