-- DROP TABLE IF EXISTS "public"."agent_permission";
CREATE TABLE "public"."agent_permission" (
  "agent_id" int4 NOT NULL DEFAULT 0,
  "level" int2 NOT NULL DEFAULT 0,
  "permission" jsonb NOT NULL DEFAULT '{}'::jsonb
)
;
COMMENT ON COLUMN "public"."agent_permission"."agent_id" IS 'mapping agent id';
COMMENT ON COLUMN "public"."agent_permission"."level" IS '預設後台帳號層級權限(defult:0) ';
COMMENT ON COLUMN "public"."agent_permission"."permission" IS '帳號權限list';

-- ----------------------------
-- Primary Key structure for table agent_permission
-- ----------------------------
ALTER TABLE "public"."agent_permission" ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("agent_id", "level");


INSERT INTO "public"."agent_permission" ("agent_id", "level", "permission") 
SELECT "id", 0, '{}' FROM "public"."agent";


UPDATE "public"."agent_permission"
SET "permission"=(
SELECT "public"."admin_user"."permission" 
FROM "public"."admin_user"
WHERE "public"."admin_user".agent_id=1 LIMIT 1);