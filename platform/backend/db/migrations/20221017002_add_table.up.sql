CREATE TABLE "public"."job_scheduler" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "spec" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "trigger_func" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT false,
  CONSTRAINT "job_scheduler_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."job_scheduler"."spec" IS '執行時間';
COMMENT ON COLUMN "public"."job_scheduler"."info" IS '備註';
COMMENT ON COLUMN "public"."job_scheduler"."trigger_func" IS '執行func，需要系統有支援才可使用';
COMMENT ON COLUMN "public"."job_scheduler"."is_enabled" IS '是否開啟';
COMMENT ON TABLE "public"."job_scheduler" IS '排程設定表';

INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled") VALUES ('b74347a2-1c87-468c-837e-a05c4939b8ff', '0 2 */1 * * *', '每小時的2分執行', 'job_rp_agent_stat_hour', 't');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled") VALUES ('7b528ce3-60fe-4744-8183-2b203fec2636', '0 5 0 */1 * *', '每天的00:05分執行', 'job_rp_agent_stat_day', 't');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled") VALUES ('230398a5-254b-4f53-bc1c-5df990f89da6', '0 15 0 1 */1 *', '每個月一號的凌晨00:15執行', 'job_rp_agent_stat_month', 't');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled") VALUES ('fa284152-a747-4754-ac3c-65ad811ccb77', '0 10 0 * * 1', '每周一的凌晨00:10執行', 'job_rp_agent_stat_week', 't');

-- ----------------------------
-- Table structure for rp_agent_stat_day
-- ----------------------------
CREATE TABLE "public"."rp_agent_stat_day" (
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
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "rp_agent_stat_day_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code")
)
;
COMMENT ON COLUMN "public"."rp_agent_stat_day"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."sum_ya" IS '總壓';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."sum_de" IS '總得';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "public"."rp_agent_stat_day"."sum_tax" IS '總抽水';
COMMENT ON TABLE "public"."rp_agent_stat_day" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_hour
-- ----------------------------
CREATE TABLE "public"."rp_agent_stat_hour" (
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
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "rp_agent_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code")
)
;
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."sum_ya" IS '總壓';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."sum_de" IS '總得';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "public"."rp_agent_stat_hour"."sum_tax" IS '總抽水';
COMMENT ON TABLE "public"."rp_agent_stat_hour" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_month
-- ----------------------------
CREATE TABLE "public"."rp_agent_stat_month" (
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
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "rp_agent_stat_month_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code")
)
;
COMMENT ON COLUMN "public"."rp_agent_stat_month"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."sum_ya" IS '總壓';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."sum_de" IS '總得';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "public"."rp_agent_stat_month"."sum_tax" IS '總抽水';
COMMENT ON TABLE "public"."rp_agent_stat_month" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_week
-- ----------------------------
CREATE TABLE "public"."rp_agent_stat_week" (
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
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "rp_agent_stat_week_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code")
)
;
COMMENT ON COLUMN "public"."rp_agent_stat_week"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."sum_ya" IS '總壓';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."sum_de" IS '總得';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "public"."rp_agent_stat_week"."sum_tax" IS '總抽水';
COMMENT ON TABLE "public"."rp_agent_stat_week" IS '代理時段統計資料(record profit)';
