CREATE TABLE "public"."admin_user_backend_action_log" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "agent_id" int4 NOT NULL DEFAULT 0,
  "username" varchar(20) NOT NULL DEFAULT '',
  "feature_code" int4 NOT NULL DEFAULT 0,
  "action_type" int2 NOT NULL DEFAULT 0,
  "error_code" int4 NOT NULL DEFAULT 0,
  "action_log" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "http_log" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "ip" varchar(40)  NOT NULL DEFAULT '',
  "method" varchar(6) NOT NULL DEFAULT '',
  "request_url" varchar(100) NOT NULL DEFAULT '',
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "admin_user_backend_action_log_pkey" PRIMARY KEY ("id")
);

COMMENT ON COLUMN "public"."admin_user_backend_action_log"."id" IS '操作記錄id';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."agent_id" IS '操作者代理id';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."username" IS '操作者';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."feature_code" IS 'api代碼';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."action_type" IS '操作類型';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."error_code" IS '錯誤代碼';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."action_log" IS '操作紀錄';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."http_log" IS 'http紀錄';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."ip" IS '登入IP';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."method" IS 'api method';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."request_url" IS '請求路徑';
COMMENT ON COLUMN "public"."admin_user_backend_action_log"."create_time" IS '操作時間';
COMMENT ON TABLE "public"."admin_user_backend_action_log" IS '後台帳號操作記錄';