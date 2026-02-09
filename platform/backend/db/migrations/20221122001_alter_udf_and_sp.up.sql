DROP FUNCTION "public"."udf_create_order_id";

DROP PROCEDURE "public"."sp_create_agent_wallet_ledger";
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