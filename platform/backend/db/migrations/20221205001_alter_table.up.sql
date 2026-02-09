ALTER TABLE "public"."agent_wallet"
  ADD COLUMN "sum_coin_in" decimal(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "sum_coin_out" decimal(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "sum_coin_backend_in" decimal(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "sum_coin_backend_out" decimal(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."agent_wallet"."sum_coin_in" IS '累積遊戲幣轉入';
COMMENT ON COLUMN "public"."agent_wallet"."sum_coin_out" IS '累積遊戲幣轉出';
COMMENT ON COLUMN "public"."agent_wallet"."sum_coin_backend_in" IS '累積遊戲幣後台轉入';
COMMENT ON COLUMN "public"."agent_wallet"."sum_coin_backend_out" IS '累積遊戲幣後台轉出';
