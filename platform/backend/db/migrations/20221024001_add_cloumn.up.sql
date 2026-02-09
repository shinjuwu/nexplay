ALTER TABLE "public"."job_scheduler" 
  ADD COLUMN "exec_limit" int2 NOT NULL DEFAULT 0,
  ADD COLUMN "last_sync_date" varchar(12) NOT NULL DEFAULT ''::character varying,
  ADD COLUMN "update_time" timestamptz(0) NOT NULL DEFAULT now();

COMMENT ON COLUMN "public"."job_scheduler"."exec_limit" IS '指定執行次數';
COMMENT ON COLUMN "public"."job_scheduler"."last_sync_date" IS '最後同步日期辨識用字串(YYYYMMDDhhmm)';
COMMENT ON COLUMN "public"."job_scheduler"."update_time" IS '最後更新時間';

INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit") VALUES ('1699256e-c294-437b-ba63-106259056491', '*/1 * * * * *', '每次開機檢查報表插入時間是否為最新', 'job_rp_agent_stat_check', 'f', 1);
