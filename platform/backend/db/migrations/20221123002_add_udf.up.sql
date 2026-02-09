CREATE FUNCTION "public"."udf_game_user_create_transfer" ("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" jsonb, "_creator" varchar, "_update_agent_wallet" boolean, "_add_agent_wallet_amount" numeric)
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

  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator")
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

CREATE FUNCTION "public"."udf_game_user_complete_transfer" ("_id" varchar, "_changeset" jsonb, "_status" int2, "_error_code" int2, "_update_agent_wallet" boolean, "_agent_id" int4, "_add_agent_wallet_amount" numeric)
  RETURNS json AS $$
BEGIN
  IF "_add_agent_wallet_amount" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  IF "_update_agent_wallet" THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", "_add_agent_wallet_amount");
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
