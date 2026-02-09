ALTER TABLE "public"."wallet_ledger"
  DROP COLUMN "status",
  DROP COLUMN "error_code",
  DROP COLUMN "creator",
  ADD COLUMN "agent_code" varchar(8) NOT NULL DEFAULT '';

COMMENT ON COLUMN "public"."wallet_ledger"."agent_code" IS 'mapping code of agent';

UPDATE "public"."wallet_ledger" AS "wl"
  SET "agent_code" = "a"."code"
  FROM "agent" AS "a"
  WHERE "a"."id" = "wl"."agent_id"
