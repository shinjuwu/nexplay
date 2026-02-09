ALTER TABLE "public"."admin_user"
  ALTER COLUMN "role" DROP DEFAULT,
  ALTER COLUMN "role" TYPE uuid USING gen_random_uuid(),
  DROP COLUMN "permission";

ALTER TABLE "public"."agent_permission"
  DROP COLUMN "level",
  ADD COLUMN "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("id");

COMMENT ON COLUMN "public"."agent_permission"."id" IS '代理權限角色編號';
