ALTER TABLE "public"."wallet_ledger"
  ADD COLUMN "request" jsonb NOT NULL DEFAULT '{}'::jsonb;

-- 先備份原本的info --
UPDATE "public"."wallet_ledger"
  SET "request" = "info";

ALTER TABLE "public"."wallet_ledger"
  ALTER COLUMN "info" TYPE varchar(255) USING '',
  ALTER COLUMN "info" SET DEFAULT '';

COMMENT ON COLUMN "public"."wallet_ledger"."request" IS '請求參數';
COMMENT ON COLUMN "public"."wallet_ledger"."info" IS '備註';
