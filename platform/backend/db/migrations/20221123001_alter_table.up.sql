ALTER TABLE "public"."wallet_ledger"
  ADD COLUMN "status" int2 NOT NULL DEFAULT 0,
  ADD COLUMN "error_code" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "creator" varchar(20) NOT NULL DEFAULT '',
  DROP COLUMN "agent_code";

COMMENT ON COLUMN "public"."wallet_ledger"."status" IS '訂單狀態';
COMMENT ON COLUMN "public"."wallet_ledger"."error_code" IS '錯誤代碼';
COMMENT ON COLUMN "public"."wallet_ledger"."creator" IS '創建人';
