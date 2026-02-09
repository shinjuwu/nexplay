CREATE TABLE "public"."marquee" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "lang" int2 NOT NULL DEFAULT 0,
  "type" int2 NOT NULL DEFAULT 0,
  "order" int2 NOT NULL DEFAULT 0,
  "freq" int2 NOT NULL DEFAULT 0,
  "is_enabled" bool NOT NULL DEFAULT false,
  "is_open" bool NOT NULL DEFAULT false,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "start_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "public"."marquee"."lang" IS '語系';
COMMENT ON COLUMN "public"."marquee"."type" IS '跑馬燈類型';
COMMENT ON COLUMN "public"."marquee"."order" IS '順序';
COMMENT ON COLUMN "public"."marquee"."freq" IS '播放頻率(每?秒播放一次)';
COMMENT ON COLUMN "public"."marquee"."is_enabled" IS '是否啟動';
COMMENT ON COLUMN "public"."marquee"."is_open" IS '是否開啟';
COMMENT ON COLUMN "public"."marquee"."content" IS '內文';
COMMENT ON COLUMN "public"."marquee"."start_time" IS '開始時間';
COMMENT ON COLUMN "public"."marquee"."end_time" IS '結束時間';
COMMENT ON COLUMN "public"."marquee"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."marquee"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."marquee" IS '跑馬燈設定表';

-- ----------------------------
-- Primary Key structure for table marquee
-- ----------------------------
ALTER TABLE "public"."marquee" ADD CONSTRAINT "marquee_pkey" PRIMARY KEY ("id");

COMMENT ON COLUMN "public"."permission_list"."is_required" IS '是否需要驗證token才可使用(true:需驗證, false:無需驗證)';

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100104, '此接口供遊戲伺服器取得跑馬燈設定列表', '/api/v1/intercom/getmarqueesetting', 't', '遊戲SERVER串接使用', 'f', '2022-10-06 06:42:06.492519+00', '2022-10-06 06:42:06.492519+00');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100222, '取得目前跑馬燈設定列表', '/api/v1/manage/getmarqueelist', 't', '後台使用', 't', '2022-10-06 06:49:11.672024+00', '2022-10-06 06:49:11.672024+00');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100223, '指定取得某筆跑馬燈設定', '/api/v1/manage/getmarquee', 't', '後台使用', 't', '2022-10-06 06:49:43.992841+00', '2022-10-06 06:49:43.992841+00');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100224, '添加跑馬燈功能', '/api/v1/manage/createmarquee', 't', '後台使用', 't', '2022-10-06 06:50:19.738799+00', '2022-10-06 06:50:19.738799+00');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100225, '編輯跑馬燈功能', '/api/v1/manage/updatemarquee', 't', '後台使用', 't', '2022-10-06 06:50:46.541694+00', '2022-10-06 06:50:46.541694+00');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required", "create_time", "update_time") VALUES (100226, '刪除跑馬燈功能', '/api/v1/manage/deletemarquee', 't', '後台使用', 't', '2022-10-06 06:51:31.765154+00', '2022-10-06 06:51:31.765154+00');


UPDATE "public"."admin_user" SET "permission" = '{"list": [100200, 100201, 100202, 100203, 100204, 100205, 100208, 100209, 100210, 100211, 100212, 100213, 100214, 100215, 100216, 100217, 100218, 100219, 100220, 100221, 100222, 100223, 100224, 100225, 100226, 200101, 200102]}' WHERE 1=1;

