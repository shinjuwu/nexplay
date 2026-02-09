CREATE TABLE "public"."resend_data" (
  "id" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "request_type" int2 NOT NULL DEFAULT 0,
  "is_resend" bool NOT NULL DEFAULT false,
  "retries" int4 NOT NULL DEFAULT 0,
  "request_from" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "request_to" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

COMMENT ON COLUMN "public"."resend_data"."id" IS '訂單號(orderid)';
COMMENT ON COLUMN "public"."resend_data"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."resend_data"."level_code" IS '層級碼,每四碼一個層級 (admin:0000)';
COMMENT ON COLUMN "public"."resend_data"."request_type" IS 'type of api request';
COMMENT ON COLUMN "public"."resend_data"."is_resend" IS '重試狀態';
COMMENT ON COLUMN "public"."resend_data"."retries" IS '重試次數';
COMMENT ON COLUMN "public"."resend_data"."request_from" IS '原始請求參數';
COMMENT ON COLUMN "public"."resend_data"."request_to" IS '回調請求參數';
COMMENT ON COLUMN "public"."resend_data"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."resend_data"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."resend_data" IS '訊息重送紀錄表';

CREATE INDEX "idx_resend_data_id" ON "public"."resend_data" USING btree (
  "id"
);
