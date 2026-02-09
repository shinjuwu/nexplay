-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE PROCEDURE "public"."usp_check_game_ratio_stat"("_thismonth" varchar, "_platform" varchar)
 AS $BODY$
DECLARE _table_name varchar(50);
DECLARE _comment_table_name varchar(50);
BEGIN

		_table_name := _platform||'_'||'game_ratio_stat'||'_'||_thismonth;
		_comment_table_name := '代理遊戲平台營收統計表';
	
		IF NOT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_catalog = current_database() AND table_schema = 'public' AND table_name = _table_name) THEN
			EXECUTE '
				CREATE TABLE IF NOT EXISTS "public"."'||_table_name||'" (
					"log_time" varchar(18) NOT NULL DEFAULT '''',
					"game_id" int4 NOT NULL DEFAULT 0,
					"game_type" int4 NOT NULL DEFAULT 0,
					"room_type" int4 NOT NULL DEFAULT 0,
					"de" numeric(20,4) NOT NULL DEFAULT 0,
					"ya" numeric(20,4) NOT NULL DEFAULT 0,
					"vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
					"tax" numeric(20,4) NOT NULL DEFAULT 0,
					"bonus" numeric(20,4) NOT NULL DEFAULT 0,
					"play_count" int8 NOT NULL DEFAULT 0,
					"update_time" timestamptz(6) NOT NULL DEFAULT now(),
					CONSTRAINT "'||_table_name||'_pkey" PRIMARY KEY ("log_time", "game_id")
				);';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_1" ON "public"."'||_table_name||'"("update_time");';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_2" ON "public"."'||_table_name||'"("game_id","game_type","room_type");';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_id" IS ''遊戲id'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_type" IS ''遊戲類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."room_type" IS ''房間類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."de" IS ''得分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."ya" IS ''壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."vaild_ya" IS ''有效壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."tax" IS ''抽水'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."bonus" IS ''紅利'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."play_count" IS ''下注次數'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."update_time" IS ''最後更新時間'';';
			EXECUTE 'COMMENT ON TABLE "public"."'||_table_name||'" IS '''||_comment_table_name||''';';
		END IF;

END
$BODY$
  LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE "public"."usp_insert_game_ratio_stat"("_platform" varchar, "_game_id" int4, "_game_type" int4, "_room_type" int4, "_de" numeric, "_ya" numeric, "_vaild_ya" numeric, "_tax" numeric, "_bonus" numeric, "_play_count" int4, "_bet_time" timestamptz)
 AS $BODY$
 DECLARE _table_month varchar(6);
 DECLARE _stat_table_name varchar(50);
 DECLARE _stat_month_table_name varchar(50);
 DECLARE _log_time varchar(8);
 BEGIN
	
	_table_month := to_char(_bet_time,'YYYYMM');
	_stat_table_name := 'game_ratio_stat';
	_stat_month_table_name := _platform||'_'||_stat_table_name||'_'||_table_month;

	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = _stat_month_table_name) THEN
		CALL "public"."usp_check_game_ratio_stat"(_table_month, _platform);
	END IF;
	
	_log_time := to_char(_bet_time,'YYYYMMDD');
	
	EXECUTE 'INSERT INTO "public"."'||_stat_month_table_name||'"("log_time", "game_id", "game_type", "room_type", "de", "ya", "vaild_ya", "tax", "bonus", "play_count", "update_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
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
					_log_time, _game_id, _game_type, _room_type, _de, _ya, _vaild_ya, _tax, _bonus, _play_count, now();
	
END;
$BODY$
  LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP PROCEDURE "public"."usp_check_agent_game_ratio_stat"("_thismonth" varchar, "_platform" varchar);
DROP PROCEDURE "public"."usp_insert_game_ratio_stat"("_platform" varchar, "_game_id" int4, "_game_type" int4, "_room_type" int4, "_de" numeric, "_ya" numeric, "_vaild_ya" numeric, "_tax" numeric, "_bonus" numeric, "_play_count" int4, "_bet_time" timestamptz);
