-- DROP TABLE IF EXISTS "public"."game_users_stat";
CREATE TABLE "public"."game_users_stat" (
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "game_users_id" int4 NOT NULL,
  "de" numeric(20,4) NOT NULL DEFAULT 0,
  "ya" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "play_count" int8 NOT NULL DEFAULT 0,
  "big_win_count" int4 NOT NULL DEFAULT 0,
  "first_bet_time" timestamptz(6) NOT NULL,
  "last_bet_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "game_users_stat_pkey" PRIMARY KEY ("agent_id", "level_code", "game_users_id")
)
;
COMMENT ON COLUMN "public"."game_users_stat"."de" IS '得分';
COMMENT ON COLUMN "public"."game_users_stat"."ya" IS '壓分';
COMMENT ON COLUMN "public"."game_users_stat"."vaild_ya" IS '有效壓分';
COMMENT ON COLUMN "public"."game_users_stat"."play_count" IS '下注次數';
COMMENT ON COLUMN "public"."game_users_stat"."big_win_count" IS '中大奨次數';
COMMENT ON COLUMN "public"."game_users_stat"."first_bet_time" IS '首次下注時間';
COMMENT ON COLUMN "public"."game_users_stat"."last_bet_time" IS '最後一次下注時間';
COMMENT ON COLUMN "public"."game_users_stat"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."game_users_stat" IS '統計玩家下注資料表';


-- DROP TABLE IF EXISTS "public"."game_users_stat_hour";
CREATE TABLE "public"."game_users_stat_hour" (
  "log_time" varchar(18) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "game_users_id" int4 NOT NULL,
  "de" numeric(20,4) NOT NULL DEFAULT 0,
  "ya" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "play_count" int8 NOT NULL DEFAULT 0,
  "big_win_count" int4 NOT NULL DEFAULT 0,
  CONSTRAINT "game_users_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code", "game_users_id")
)
;
COMMENT ON COLUMN "public"."game_users_stat_hour"."de" IS '得分';
COMMENT ON COLUMN "public"."game_users_stat_hour"."ya" IS '壓分';
COMMENT ON COLUMN "public"."game_users_stat_hour"."vaild_ya" IS '有效壓分';
COMMENT ON COLUMN "public"."game_users_stat_hour"."play_count" IS '下注次數';
COMMENT ON COLUMN "public"."game_users_stat_hour"."big_win_count" IS '中大奨次數';
COMMENT ON TABLE "public"."game_users_stat_hour" IS '每小時統計玩家下注資料表';


-- DROP PROCEDURE IF EXISTS "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_play_count" int4, "_big_win_count" int4, "_last_bet_time" timestamptz);
CREATE OR REPLACE PROCEDURE "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_play_count" int4, "_big_win_count" int4, "_last_bet_time" timestamptz)
 AS $BODY$
 DECLARE _log_hour varchar(12);
 BEGIN
	-- Routine body goes here...
	_log_hour := to_char(_last_bet_time,'YYYYMMDDHH24');

	EXECUTE 'INSERT INTO "public"."game_users_stat"(
		"agent_id", "level_code", "game_users_id", "de", "ya",
		"vaild_ya", "play_count", "big_win_count", "first_bet_time", "last_bet_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
			ON CONFLICT ON CONSTRAINT "game_users_stat_pkey" DO 
			UPDATE SET 	
				de = game_users_stat.de + EXCLUDED.de,
				ya = game_users_stat.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat.vaild_ya + EXCLUDED.vaild_ya,
				play_count = game_users_stat.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat.big_win_count + EXCLUDED.big_win_count,
				last_bet_time =  $11,
				update_time = now()
			;'
	USING
	_agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _play_count, _big_win_count, _last_bet_time, _last_bet_time, _last_bet_time;
				
	EXECUTE 'INSERT INTO "public"."game_users_stat_hour"(
		"log_time", "agent_id", "level_code", "game_users_id", "de",
		"ya",	"vaild_ya", "play_count", "big_win_count")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9
			ON CONFLICT ON CONSTRAINT "game_users_stat_hour_pkey" DO 
			UPDATE SET 	
				de = game_users_stat_hour.de + EXCLUDED.de,
				ya = game_users_stat_hour.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat_hour.vaild_ya + EXCLUDED.vaild_ya,
				play_count = game_users_stat_hour.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat_hour.big_win_count + EXCLUDED.big_win_count
			;'
	USING
	_log_hour, _agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _play_count, _big_win_count;

END$BODY$
  LANGUAGE plpgsql;

UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'c23621fc-a492-44d3-9873-4e4be72292ef';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = '4d7add6a-b629-4f20-87fb-f2de409d4627';
UPDATE "public"."job_scheduler" SET "is_enabled" = 'f' WHERE "id" = 'f25055d6-b3c2-4d44-a539-7daccbe8579f';