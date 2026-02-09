CREATE OR REPLACE PROCEDURE "public"."usp_insert_agent_game_ratio_stat"("_id" varchar, "_level_code" varchar, "_agent_id" int4, "_game_id" int4, "_game_type" int4, "_room_type" int4, "_de" numeric, "_ya" numeric, "_vaild_ya" numeric, "_tax" numeric, "_bonus" numeric, "_play_count" int4, "_bet_time" timestamptz)
 AS $BODY$
 DECLARE _table_month varchar(6);
 DECLARE _stat_table_name varchar(50);
 DECLARE _stat_month_table_name varchar(50);
 DECLARE _log_time varchar(8);
 BEGIN
	
	_table_month := to_char(_bet_time,'YYYYMM');
	_stat_table_name := 'agent_game_ratio_stat';
	_stat_month_table_name := _stat_table_name||'_'||_table_month;

	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = _stat_table_name) THEN
		CALL "public"."usp_check_agent_game_ratio_stat"('');
	END IF;
	
	EXECUTE 'INSERT INTO "public"."'||_stat_table_name||'"("id", "level_code", "agent_id", "game_id", "game_type", "room_type", "de", "ya", "vaild_ya", "tax", "bonus", "play_count", "update_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			ON CONFLICT ON CONSTRAINT "'||_stat_table_name||'_pkey" DO 
			UPDATE SET 	
				de = '||_stat_table_name||'.de + EXCLUDED.de,
				ya = '||_stat_table_name||'.ya + EXCLUDED.ya,
				vaild_ya = '||_stat_table_name||'.vaild_ya + EXCLUDED.vaild_ya,
				tax = '||_stat_table_name||'.tax + EXCLUDED.tax,
				bonus = '||_stat_table_name||'.bonus + EXCLUDED.bonus,
				play_count = '||_stat_table_name||'.play_count + EXCLUDED.play_count,
				update_time = EXCLUDED.update_time
						;'
				USING
					_id, _level_code, _agent_id, _game_id, _game_type, _room_type, _de, _ya, _vaild_ya, _tax, _bonus, _play_count, now();

	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = _stat_month_table_name) THEN
		CALL "public"."usp_check_agent_game_ratio_stat"(_table_month);
	END IF;
	
	_log_time := to_char(_bet_time,'YYYYMMDD');
	
	EXECUTE 'INSERT INTO "public"."'||_stat_month_table_name||'"("log_time", "id", "level_code", "agent_id", "game_id", "game_type", "room_type", "de", "ya", "vaild_ya", "tax", "bonus", "play_count", "update_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
			ON CONFLICT ON CONSTRAINT "'||_stat_month_table_name||'_pkey" DO 
			UPDATE SET 	
				de = '||_stat_month_table_name||'.de + EXCLUDED.de,
				ya = '||_stat_month_table_name||'.ya + EXCLUDED.ya,
				vaild_ya = '||_stat_month_table_name||'.vaild_ya + EXCLUDED.vaild_ya,
            tax = '||_stat_month_table_name||'.tax + EXCLUDED.tax,
				bonus = '||_stat_month_table_name||'.bonus + EXCLUDED.bonus,
				play_count = '||_stat_month_table_name||'.play_count + EXCLUDED.play_count,
				update_time = EXCLUDED.update_time
						;'
				USING
					_log_time, _id, _level_code, _agent_id, _game_id, _game_type, _room_type, _de, _ya, _vaild_ya, _tax, _bonus, _play_count, now();
	
END$BODY$
  LANGUAGE plpgsql;