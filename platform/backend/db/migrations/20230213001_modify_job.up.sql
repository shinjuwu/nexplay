UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = '230398a5-254b-4f53-bc1c-5df990f89da6';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'fa284152-a747-4754-ac3c-65ad811ccb77';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = '7b528ce3-60fe-4744-8183-2b203fec2636';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'b74347a2-1c87-468c-837e-a05c4939b8ff';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = '89ab932a-2807-4ea8-9694-ffb66f12d567';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'c340d2af-8b0d-41b4-abf5-6b22ff4bbf62';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = '8bc970c5-832d-486a-8d6c-5502496179d1';

INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('0f696d8a-b889-49fa-8a32-cfea6ae02e20', '30 0,15,30,45 * * * *', '每小時的0,15,30,45分30秒執行', 'job_rp_agent_stat_15min', 't', 0, '', '2023-02-13 05:43:42+00');

CREATE TABLE "rp_agent_stat_15min" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat_15min"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat_15min"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat_15min"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat_15min"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat_15min"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat_15min"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat_15min"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat_15min"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat_15min"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat_15min"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "rp_agent_stat_15min"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat_15min" IS '代理時段統計資料(record profit)';