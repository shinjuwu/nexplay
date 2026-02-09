DROP TABLE IF EXISTS "public"."rt_data_stat_week";
DROP TABLE IF EXISTS "public"."rt_data_stat_month";
DROP TABLE IF EXISTS "public"."rp_agent_stat_day";
DROP TABLE IF EXISTS "public"."rp_agent_stat_hour";
DROP TABLE IF EXISTS "public"."rp_agent_stat_month";
DROP TABLE IF EXISTS "public"."rp_agent_stat_week";


CREATE TABLE "public"."rt_data_stat_hour" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "active_player" int4 NOT NULL DEFAULT 0,
  "number_bettors" int4 NOT NULL DEFAULT 0,
  "number_registrants" int4 NOT NULL DEFAULT 0,
  "odd_number" int4 NOT NULL DEFAULT 0,
  "total_betting" numeric(20,4) NOT NULL DEFAULT 0,
  "game_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_win_score" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_lose_score" numeric(20,4) NOT NULL DEFAULT 0,
  "raw_data" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "rt_data_stat_hour_pkey" PRIMARY KEY ("log_time")
)
;
COMMENT ON COLUMN "public"."rt_data_stat_hour"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."active_player" IS '活躍玩家(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."number_bettors" IS '註冊人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."number_registrants" IS '投注人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."odd_number" IS '注單數';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."total_betting" IS '總投注';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."game_tax" IS '遊戲抽水';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."platform_win_score" IS '平台總贏分數';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."platform_lose_score" IS '平台總輸分數';
COMMENT ON COLUMN "public"."rt_data_stat_hour"."raw_data" IS '原始數據';
COMMENT ON TABLE "public"."rt_data_stat_hour" IS '代理即時統計資料(realtime data stat)';


DELETE FROM "public"."job_scheduler" WHERE "id" = 'c340d2af-8b0d-41b4-abf5-6b22ff4bbf62';
DELETE FROM "public"."job_scheduler" WHERE "id" = '8bc970c5-832d-486a-8d6c-5502496179d1';
DELETE FROM "public"."job_scheduler" WHERE "id" = '1699256e-c294-437b-ba63-106259056491';
DELETE FROM "public"."job_scheduler" WHERE "id" = '7b528ce3-60fe-4744-8183-2b203fec2636';
DELETE FROM "public"."job_scheduler" WHERE "id" = 'b74347a2-1c87-468c-837e-a05c4939b8ff';
DELETE FROM "public"."job_scheduler" WHERE "id" = '230398a5-254b-4f53-bc1c-5df990f89da6';
DELETE FROM "public"."job_scheduler" WHERE "id" = 'fa284152-a747-4754-ac3c-65ad811ccb77';
DELETE FROM "public"."job_scheduler" WHERE "id" = 'f25055d6-b3c2-4d44-a539-7daccbe8579f';
DELETE FROM "public"."job_scheduler" WHERE "id" = '4d7add6a-b629-4f20-87fb-f2de409d4627';
DELETE FROM "public"."job_scheduler" WHERE "id" = 'c23621fc-a492-44d3-9873-4e4be72292ef';
DELETE FROM "public"."job_scheduler" WHERE "id" = '89ab932a-2807-4ea8-9694-ffb66f12d567';
UPDATE "public"."job_scheduler" SET "trigger_func" = 'job_rt_data_stat_backup_min' WHERE "id" = 'e1ca6b64-5d13-4cc3-8e36-bd0ef8537a4f';
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('6ec15b7a-12a1-4c31-a045-9709353d5369', '0 1 */1 * * *', '每小時的1分執行', 'job_rt_data_stat_backup_hour', 't', 0, '', '2023-03-09 06:27:27+00');
