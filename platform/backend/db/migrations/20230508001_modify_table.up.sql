ALTER TABLE "public"."server_info" 
  DROP COLUMN "setting";

UPDATE "public"."game" SET "server_info_code" = 'dev01' WHERE "server_info_code" = 'dev';
UPDATE "public"."server_info" SET "code" = 'dev01' WHERE "code" = 'dev';

UPDATE "public"."game" SET "server_info_code" = 'qa01' WHERE "server_info_code" = 'qa';
UPDATE "public"."server_info" SET "code" = 'qa01' WHERE "code" = 'qa';

UPDATE "public"."game" SET "server_info_code" = 'ete01' WHERE "server_info_code" = 'ete';
UPDATE "public"."server_info" SET "code" = 'ete01' WHERE "code" = 'ete';