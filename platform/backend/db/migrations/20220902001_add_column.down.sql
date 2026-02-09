ALTER TABLE "public"."admin_user_action_log" RENAME COLUMN "error_code" TO "action_code";

ALTER TABLE "public"."admin_user_action_log" 
  DROP COLUMN "method",
  DROP COLUMN "request_url";
