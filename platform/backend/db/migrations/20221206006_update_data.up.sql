ALTER TABLE "public"."rp_agent_stat_day" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rp_agent_stat_hour" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rp_agent_stat_month" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rp_agent_stat_week" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rt_data_stat_day" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rt_data_stat_week" 
  DROP COLUMN "agent_name";

ALTER TABLE "public"."rt_data_stat_month" 
  DROP COLUMN "agent_name";

UPDATE rt_data_stat_day AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE rt_data_stat_week AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE rt_data_stat_month AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;
