CREATE OR REPLACE FUNCTION "public"."udf_game_user_finish_coin_in_cancel"("_id" varchar, "_changeset" jsonb, "_status" int2, "_error_code" int2, "_agent_id" int4, "_add_agent_wallet_amount" numeric, "_add_agent_wallet_sum_coin_out" numeric, "_user_id" int4, "_add_user_sum_coin_in" numeric)
  RETURNS "pg_catalog"."json" AS $BODY$
BEGIN
  -- 檢查參數 --
  IF "_add_agent_wallet_amount" < 0 OR "_add_agent_wallet_sum_coin_out" < 0 OR "_add_user_sum_coin_in" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  -- 取消上分成功，加回已經扣款的商戶錢 --
  IF "_error_code" = 0 THEN
    UPDATE "public"."agent_wallet"
      SET "amount" = "amount" + "_add_agent_wallet_amount",
        "sum_coin_out" = "sum_coin_out" - "_add_agent_wallet_sum_coin_out"
      WHERE "agent_id" = "_agent_id";
  END IF;

  -- 更新紀錄 --
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
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;