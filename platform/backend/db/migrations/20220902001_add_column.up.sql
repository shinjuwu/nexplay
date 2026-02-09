ALTER TABLE "public"."admin_user_action_log" RENAME COLUMN "action_code" TO "error_code";

ALTER TABLE "public"."admin_user_action_log" 
  ADD COLUMN "method" varchar(6) NOT NULL DEFAULT '',
  ADD COLUMN "request_url" varchar(100) NOT NULL DEFAULT '';

COMMENT ON COLUMN "public"."admin_user_action_log"."error_code" IS '錯誤碼';

COMMENT ON COLUMN "public"."admin_user_action_log"."method" IS 'api method';

COMMENT ON COLUMN "public"."admin_user_action_log"."request_url" IS '請求路徑';
