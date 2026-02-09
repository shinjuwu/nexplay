ALTER TABLE "public"."agent_permission" 
  DROP COLUMN "name" ,
  DROP COLUMN "info" ,
  DROP COLUMN "account_type";

ALTER TABLE "public"."agent_permission" 
  ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("agent_id", "level");