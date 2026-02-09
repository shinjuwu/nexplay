-- 統計玩家下注資料表刪除遊戲id 調整pk key --
ALTER TABLE "public"."game_users_stat" DROP COLUMN "game_id";
ALTER TABLE "public"."game_users_stat" ADD CONSTRAINT "game_users_stat_pkey" PRIMARY KEY ("agent_id", "level_code", "game_users_id");

-- 每小時統計玩家下注資料刪除遊戲id 調整pk key --
ALTER TABLE "public"."game_users_stat_hour" DROP COLUMN "game_id";
ALTER TABLE "public"."game_users_stat_hour" ADD CONSTRAINT "game_users_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code", "game_users_id");

-- 統計玩家usp下注資料新增bonus --
DROP PROCEDURE IF EXISTS "public"."usp_game_users_stat";
CREATE OR REPLACE PROCEDURE "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_tax" float8, "_bonus" float8, "_play_count" int4, "_big_win_count" int4, "_win_count" int4, "_lose_count" int4, "_last_bet_time" timestamptz)
 AS $BODY$
 DECLARE _log_hour varchar(12);
 BEGIN
	-- Routine body goes here...
	_log_hour := to_char(_last_bet_time,'YYYYMMDDHH24');

	EXECUTE 'INSERT INTO "public"."game_users_stat"(
		"agent_id", "level_code", "game_users_id", "de", "ya",
		"vaild_ya", "tax", "bonus", "play_count", "big_win_count",
        "first_bet_time", "last_bet_time", "win_count", "lose_count")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
			ON CONFLICT ON CONSTRAINT "game_users_stat_pkey" DO 
			UPDATE SET 	
				de = game_users_stat.de + EXCLUDED.de,
				ya = game_users_stat.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat.vaild_ya + EXCLUDED.vaild_ya,
				tax = game_users_stat.tax + EXCLUDED.tax,
                bonus = game_users_stat.bonus + EXCLUDED.bonus,
				play_count = game_users_stat.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat.big_win_count + EXCLUDED.big_win_count,
                win_count = game_users_stat.win_count + EXCLUDED.win_count,
				lose_count = game_users_stat.lose_count + EXCLUDED.lose_count,
				last_bet_time =  $15,
				update_time = now()
			;'
	USING
	_agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _tax, _bonus, _play_count, _big_win_count, _last_bet_time, _last_bet_time, _win_count, _lose_count, _last_bet_time;
				
	EXECUTE 'INSERT INTO "public"."game_users_stat_hour"(
		"log_time", "agent_id", "level_code", "game_users_id", "de",
		"ya", "vaild_ya", "tax", "bonus", "play_count",
        "big_win_count", "win_count", "lose_count")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			ON CONFLICT ON CONSTRAINT "game_users_stat_hour_pkey" DO 
			UPDATE SET 	
				de = game_users_stat_hour.de + EXCLUDED.de,
				ya = game_users_stat_hour.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat_hour.vaild_ya + EXCLUDED.vaild_ya,
				tax = game_users_stat_hour.tax + EXCLUDED.tax,
				bonus = game_users_stat_hour.bonus + EXCLUDED.bonus,
				play_count = game_users_stat_hour.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat_hour.big_win_count + EXCLUDED.big_win_count,
				win_count = game_users_stat_hour.win_count + EXCLUDED.win_count,
				lose_count = game_users_stat_hour.lose_count + EXCLUDED.lose_count
			;'
	USING
	_log_hour, _agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _tax, _bonus, _play_count, _big_win_count, _win_count, _lose_count;

END$BODY$
  LANGUAGE plpgsql;