ALTER TABLE "public"."server_info" DROP COLUMN "setting";

DELETE FROM "public"."permission_list" WHERE "feature_code" = 100228;