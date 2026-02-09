ALTER TABLE "public"."admin_user"
  ALTER COLUMN "role" TYPE int2 USING 0,
  ALTER COLUMN "role" SET DEFAULT 0,
  ADD COLUMN "permission" jsonb NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN "public"."admin_user"."permission" IS '帳號權限list';

ALTER TABLE "public"."agent_permission"
  DROP COLUMN "id",
  ADD COLUMN "level" int2 NOT NULL DEFAULT 0,
  ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("agent_id", "level", "account_type");

COMMENT ON COLUMN "public"."agent_permission"."level" IS '預設後台帳號層級權限(defult:0) ';
