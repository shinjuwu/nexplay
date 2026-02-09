DROP TABLE IF EXISTS "public"."permission_list";

ALTER TABLE "public"."admin_user" 
  DROP COLUMN "permission";
