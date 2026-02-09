ALTER TABLE "public"."user_play_log_andarbahar"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_baccarat"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_blackjack"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_bullbull"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_catte"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_chinesepoker"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_cockfight"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_colordisc"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_dogracing"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_fantan"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_fruit777slot"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_fruitslot"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_goldenflower"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_hundredsicbo"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_okey"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_plinko"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_pokdeng"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_prawncrab"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_rcfishing"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_rocket"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_rummy"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_sangong"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."user_play_log_texas"
  DROP COLUMN "wallet_ledger_id";
ALTER TABLE "public"."wallet_ledger" 
  DROP COLUMN "single_wallet_id";

DROP INDEX "idx_wallet_ledger_id";
DROP INDEX "idx_wallet_ledger_single_wallet_id";

DELETE FROM "public"."permission_list" WHERE "feature_code" = 100106 AND "api_path" = '/api/v1/intercom/singlewallet';

DROP FUNCTION "public"."udf_game_user_start_coin_in"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric, "_single_wallet_id" varchar);
CREATE OR REPLACE FUNCTION "public"."udf_game_user_start_coin_in"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "_agent_wallet_amount" numeric;
  "_insert_wallet_ledger_count" bigint;
BEGIN
  -- 檢查參數 --
  IF "_add_agent_wallet_amount" < 0 OR "_add_agent_wallet_sum_coin_out" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  -- 先扣款商戶 --
  UPDATE "public"."agent_wallet"
    SET "amount" = "amount" - "_add_agent_wallet_amount",
      "sum_coin_out" = "sum_coin_out" + "_add_agent_wallet_sum_coin_out"
    WHERE "agent_id" = "_agent_id"
    RETURNING "amount" INTO "_agent_wallet_amount";

  -- 商戶錢不夠rollback --
  IF "_agent_wallet_amount" < 0 THEN
    UPDATE "public"."agent_wallet"
      SET "amount" = "amount" + "_add_agent_wallet_amount",
        "sum_coin_out" = "sum_coin_out" - "_add_agent_wallet_sum_coin_out"
      WHERE "agent_id" = "_agent_id";

    RETURN json_build_object(
      'code', 2
    );    
  END IF;

  -- 建立紀錄 --
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
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;


DROP FUNCTION "public"."udf_game_user_start_coin_out"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_single_wallet_id" varchar);
CREATE OR REPLACE FUNCTION "public"."udf_game_user_start_coin_out"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "_insert_wallet_ledger_count" bigint;
BEGIN
  -- 建立紀錄 --
  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator", "request")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator", "_request")
    ON CONFLICT ("id") DO NOTHING;

  GET DIAGNOSTICS "_insert_wallet_ledger_count" = ROW_COUNT;
  IF "_insert_wallet_ledger_count" = 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  RETURN json_build_object(
    'code', 0
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;