ALTER TABLE "public"."wallet_ledger"
  ALTER COLUMN "info" SET DEFAULT '{}'::jsonb;

ALTER TABLE "public"."wallet_ledger"
  ALTER COLUMN "info" TYPE jsonb USING "request",
  DROP COLUMN "request";

COMMENT ON COLUMN "public"."wallet_ledger"."info" IS '詳細內容';
