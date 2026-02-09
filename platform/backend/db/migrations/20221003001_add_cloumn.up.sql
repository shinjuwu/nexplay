ALTER TABLE "public"."game" 
  ADD COLUMN "room_number" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "table_number" int4 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."game"."room_number" IS '房間數量';

COMMENT ON COLUMN "public"."game"."table_number" IS '桌子數量';


UPDATE "public"."game" SET "room_number" = 4, "table_number" = 4 WHERE "id" = 1001;
UPDATE "public"."game" SET "room_number" = 4, "table_number" = 1 WHERE "id" = 1002;
UPDATE "public"."game" SET "room_number" = 4, "table_number" = 1 WHERE "id" = 1003;
UPDATE "public"."game" SET "room_number" = 4, "table_number" = 1 WHERE "id" = 1004;
UPDATE "public"."game" SET "room_number" = 4, "table_number" = 0 WHERE "id" = 2001;



ALTER TABLE "public"."admin_user" 
  ADD COLUMN "login_time" timestamp(6) NOT NULL DEFAULT now();

COMMENT ON COLUMN "public"."admin_user"."login_time" IS '最後登入時間';