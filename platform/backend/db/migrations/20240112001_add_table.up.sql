ALTER TABLE "public"."user_play_log_andarbahar"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_baccarat"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_blackjack"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_bullbull"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_catte"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_chinesepoker"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_cockfight"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_colordisc"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_dogracing"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_fantan"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_fruit777slot"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_fruitslot"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_goldenflower"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_hundredsicbo"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_okey"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_plinko"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_pokdeng"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_prawncrab"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_rcfishing"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_rocket"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_rummy"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_sangong"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;
ALTER TABLE "public"."user_play_log_texas"
  ADD COLUMN "wallet_ledger_id" varchar(100) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_andarbahar"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_catte"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_fantan"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_okey"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_plinko"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_rocket"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_rummy"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_sangong"."wallet_ledger_id" IS '對應單一錢包上下分';
COMMENT ON COLUMN "public"."user_play_log_texas"."wallet_ledger_id" IS '對應單一錢包上下分';


ALTER TABLE "public"."wallet_ledger" 
  ADD COLUMN "single_wallet_id" varchar(36) NOT NULL DEFAULT ''::character varying,
  ALTER COLUMN "info" SET DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."wallet_ledger"."single_wallet_id" IS '單一錢包上下分群組識別碼';

CREATE INDEX "idx_wallet_ledger_id" ON "public"."wallet_ledger" (
  "id"
);
CREATE INDEX "idx_wallet_ledger_single_wallet_id" ON "public"."wallet_ledger" (
  "single_wallet_id"
);

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required", "action_type") VALUES (100106, '此接口用於單一錢包 Api 整合', '/api/v1/intercom/singlewallet', 't', '遊戲SERVER串接使用', now(),  now(), 'f', 0);

DROP FUNCTION "public"."udf_game_user_start_coin_in"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric);
CREATE OR REPLACE FUNCTION "public"."udf_game_user_start_coin_in"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric, "_single_wallet_id" varchar)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "_agent_wallet_amount" numeric;
  "_insert_wallet_ledger_count" bigint;
	"_error_code_value_not_enough" int4;
	"_changeset" jsonb;
BEGIN

_error_code_value_not_enough = 22;
_changeset =json_build_object(
      'add_coin', _add_agent_wallet_amount,
			'after_coin', 0,
			'before_coin', 0,
			'to_coin', 0,
			'currency',''
    );
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

	_status = '1';

		INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator", "request", 			"single_wallet_id", "error_code", "changeset")
			VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator", "_request", "_single_wallet_id", "_error_code_value_not_enough", "_changeset")
			ON CONFLICT ("id") DO NOTHING;

  GET DIAGNOSTICS "_insert_wallet_ledger_count" = ROW_COUNT;
  IF "_insert_wallet_ledger_count" = 0 THEN
    RETURN json_build_object(
      'code', 3
    );
  END IF;

    RETURN json_build_object(
      'code', 2
    );    
  END IF;

  -- 建立紀錄 --
  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator", "request", "single_wallet_id")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator", "_request", "_single_wallet_id")
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

DROP FUNCTION "public"."udf_game_user_start_coin_out"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb);
CREATE OR REPLACE FUNCTION "public"."udf_game_user_start_coin_out"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" varchar, "_creator" varchar, "_request" jsonb, "_single_wallet_id" varchar)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "_insert_wallet_ledger_count" bigint;
BEGIN
  -- 建立紀錄 --
  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator", "request", "single_wallet_id")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator", "_request", "_single_wallet_id")
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
