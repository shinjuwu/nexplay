DROP TABLE IF EXISTS "backend_login_log";

CREATE TABLE IF NOT EXISTS "public"."backend_login_log" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "login_time" timestamptz(6) NOT NULL DEFAULT now(),
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ip" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "error_code" int4 NOT NULL DEFAULT 0,
  CONSTRAINT "backend_login_log_pkey" PRIMARY KEY ("id")
);

CREATE INDEX "idx_backend_login_log" ON "public"."backend_login_log" ("agent_id" ASC NULLS LAST);

COMMENT ON COLUMN "public"."backend_login_log"."login_time" IS '登入時間';
COMMENT ON COLUMN "public"."backend_login_log"."agent_id" IS '代理ID';
COMMENT ON COLUMN "public"."backend_login_log"."username" IS '登入帳號';
COMMENT ON COLUMN "public"."backend_login_log"."ip" IS '登入IP';
COMMENT ON COLUMN "public"."backend_login_log"."error_code" IS '錯誤碼';




