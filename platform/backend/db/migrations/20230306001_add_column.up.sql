UPDATE "public"."job_scheduler" SET "is_enabled" = 't' WHERE "id" = '89ab932a-2807-4ea8-9694-ffb66f12d567';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'c23621fc-a492-44d3-9873-4e4be72292ef';

ALTER TABLE "public"."rt_data_stat_day" 
  DROP CONSTRAINT "rt_data_stat_day_pkey",
  ADD CONSTRAINT "rt_data_stat_day_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");