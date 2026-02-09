DROP TABLE IF EXISTS "public"."game_users_stat";
DROP TABLE IF EXISTS "public"."game_users_stat_hour";
DROP PROCEDURE IF EXISTS "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_play_count" int4, "_big_win_count" int4, "_last_bet_time" timestamptz);