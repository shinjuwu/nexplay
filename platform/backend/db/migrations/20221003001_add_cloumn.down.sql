ALTER TABLE "public"."game" 
  DROP COLUMN "room_number",
  DROP COLUMN "table_number";

ALTER TABLE "public"."admin_user" 
  DROP COLUMN "login_time";