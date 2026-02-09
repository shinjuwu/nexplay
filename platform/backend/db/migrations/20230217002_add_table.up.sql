CREATE TABLE "public"."announcement" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "type" int2 NOT NULL DEFAULT 0,
  "is_open" bool NOT NULL DEFAULT false,
  "subject" varchar(100) NOT NULL DEFAULT '',
  "content" varchar(1000) NOT NULL DEFAULT '',
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "announcement_pkey" PRIMARY KEY ("id")
);

COMMENT ON COLUMN "public"."announcement"."id" IS '公告uuid';
COMMENT ON COLUMN "public"."announcement"."type" IS '公告類型';
COMMENT ON COLUMN "public"."announcement"."is_open" IS '是否開啟';
COMMENT ON COLUMN "public"."announcement"."subject" IS '主旨';
COMMENT ON COLUMN "public"."announcement"."content" IS '內文';
COMMENT ON COLUMN "public"."announcement"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."announcement"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."announcement" IS '公告設定表';

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100270, '取得目前後台公告列表', '/api/v1/manage/getannouncementlist', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100271, '指定取得某筆後台公告設定', '/api/v1/manage/getannouncement', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100272, '添加後台公告功能', '/api/v1/manage/createannouncement', 't', '後台使用', 't', 1);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100273, '編輯後台公告功能', '/api/v1/manage/updateannouncement', 't', '後台使用', 't', 2);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100274, '刪除後台公告功能', '/api/v1/manage/deleteannouncement', 't', '後台使用', 't', 3);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100272, 100273, 100274]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100270, 100271]'::jsonb, false)
WHERE "agent_id" = -1;
