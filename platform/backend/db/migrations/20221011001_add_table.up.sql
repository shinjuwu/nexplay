-- ----------------------------------------------------------------------------------------------------------------
-- 
-- ----------------------------------------------------------------------------------------------------------------
CREATE TABLE "public"."rp_agent_stat" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "log_hour" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
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
  "create_time" timestamptz NOT NULL DEFAULT now(),
  "update_time" timestamptz NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "public"."rp_agent_stat"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rp_agent_stat"."log_hour" IS 'YYYYMMDDhh';
COMMENT ON COLUMN "public"."rp_agent_stat"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rp_agent_stat"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rp_agent_stat"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rp_agent_stat"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "public"."rp_agent_stat"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "public"."rp_agent_stat"."sum_ya" IS '總壓';
COMMENT ON COLUMN "public"."rp_agent_stat"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "public"."rp_agent_stat"."sum_de" IS '總得';
COMMENT ON COLUMN "public"."rp_agent_stat"."sum_bonus" IS '總額外獎勵';
COMMENT ON COLUMN "public"."rp_agent_stat"."sum_tax" IS '總抽水';
COMMENT ON TABLE "public"."rp_agent_stat" IS '代理時段統計資料(record profit)';

ALTER TABLE "public"."rp_agent_stat" ADD CONSTRAINT "rp_agent_stat_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");
