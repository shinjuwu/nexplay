ALTER TABLE "public"."agent" 
  DROP COLUMN "member_count";

DROP FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar);
