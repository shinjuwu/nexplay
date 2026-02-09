-- mv_cal_game_stat_hour刪除bonus --
DROP MATERIALIZED VIEW mv_cal_game_stat_hour;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_cal_game_stat_hour AS
SELECT date_trunc('hour'::text, tmp.bet_time) AS log_time,
    tmp.level_code,
    tmp.game_id,
    count(DISTINCT tmp.user_id) AS bet_user,
    count(tmp.lognumber) AS bet_count,
    sum(tmp.ya_score) AS sum_ya,
    sum(tmp.valid_score) AS sum_valid_ya,
    sum(tmp.de_score) AS sum_de,
    0 AS sum_bonus,
    sum(tmp.tax) AS sum_tax
   FROM ( SELECT user_play_log_baccarat.lognumber,
            user_play_log_baccarat.level_code,
            user_play_log_baccarat.user_id,
            user_play_log_baccarat.game_id,
            user_play_log_baccarat.ya_score,
            user_play_log_baccarat.valid_score,
            user_play_log_baccarat.de_score,
            user_play_log_baccarat.tax,
            user_play_log_baccarat.bet_time
           FROM user_play_log_baccarat
          WHERE ((user_play_log_baccarat.is_robot = 0) AND (user_play_log_baccarat.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_colordisc.lognumber,
            user_play_log_colordisc.level_code,
            user_play_log_colordisc.user_id,
            user_play_log_colordisc.game_id,
            user_play_log_colordisc.ya_score,
            user_play_log_colordisc.valid_score,
            user_play_log_colordisc.de_score,
            user_play_log_colordisc.tax,
            user_play_log_colordisc.bet_time
           FROM user_play_log_colordisc
          WHERE ((user_play_log_colordisc.is_robot = 0) AND (user_play_log_colordisc.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_fantan.lognumber,
            user_play_log_fantan.level_code,
            user_play_log_fantan.user_id,
            user_play_log_fantan.game_id,
            user_play_log_fantan.ya_score,
            user_play_log_fantan.valid_score,
            user_play_log_fantan.de_score,
            user_play_log_fantan.tax,
            user_play_log_fantan.bet_time
           FROM user_play_log_fantan
          WHERE ((user_play_log_fantan.is_robot = 0) AND (user_play_log_fantan.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_prawncrab.lognumber,
            user_play_log_prawncrab.level_code,
            user_play_log_prawncrab.user_id,
            user_play_log_prawncrab.game_id,
            user_play_log_prawncrab.ya_score,
            user_play_log_prawncrab.valid_score,
            user_play_log_prawncrab.de_score,
            user_play_log_prawncrab.tax,
            user_play_log_prawncrab.bet_time
           FROM user_play_log_prawncrab
          WHERE ((user_play_log_prawncrab.is_robot = 0) AND (user_play_log_prawncrab.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_blackjack.lognumber,
            user_play_log_blackjack.level_code,
            user_play_log_blackjack.user_id,
            user_play_log_blackjack.game_id,
            user_play_log_blackjack.ya_score,
            user_play_log_blackjack.valid_score,
            user_play_log_blackjack.de_score,
            user_play_log_blackjack.tax,
            user_play_log_blackjack.bet_time
           FROM user_play_log_blackjack
          WHERE ((user_play_log_blackjack.is_robot = 0) AND (user_play_log_blackjack.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_sangong.lognumber,
            user_play_log_sangong.level_code,
            user_play_log_sangong.user_id,
            user_play_log_sangong.game_id,
            user_play_log_sangong.ya_score,
            user_play_log_sangong.valid_score,
            user_play_log_sangong.de_score,
            user_play_log_sangong.tax,
            user_play_log_sangong.bet_time
           FROM user_play_log_sangong
          WHERE ((user_play_log_sangong.is_robot = 0) AND (user_play_log_sangong.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_hundredsicbo.lognumber,
            user_play_log_hundredsicbo.level_code,
            user_play_log_hundredsicbo.user_id,
            user_play_log_hundredsicbo.game_id,
            user_play_log_hundredsicbo.ya_score,
            user_play_log_hundredsicbo.valid_score,
            user_play_log_hundredsicbo.de_score,
            user_play_log_hundredsicbo.tax,
            user_play_log_hundredsicbo.bet_time
           FROM user_play_log_hundredsicbo
          WHERE ((user_play_log_hundredsicbo.is_robot = 0) AND (user_play_log_hundredsicbo.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_bullbull.lognumber,
            user_play_log_bullbull.level_code,
            user_play_log_bullbull.user_id,
            user_play_log_bullbull.game_id,
            user_play_log_bullbull.ya_score,
            user_play_log_bullbull.valid_score,
            user_play_log_bullbull.de_score,
            user_play_log_bullbull.tax,
            user_play_log_bullbull.bet_time
           FROM user_play_log_bullbull
          WHERE ((user_play_log_bullbull.is_robot = 0) AND (user_play_log_bullbull.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_texas.lognumber,
            user_play_log_texas.level_code,
            user_play_log_texas.user_id,
            user_play_log_texas.game_id,
            user_play_log_texas.ya_score,
            user_play_log_texas.valid_score,
            user_play_log_texas.de_score,
            user_play_log_texas.tax,
            user_play_log_texas.bet_time
           FROM user_play_log_texas
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_fruitslot.lognumber,
            user_play_log_fruitslot.level_code,
            user_play_log_fruitslot.user_id,
            user_play_log_fruitslot.game_id,
            user_play_log_fruitslot.ya_score,
            user_play_log_fruitslot.valid_score,
            user_play_log_fruitslot.de_score,
            user_play_log_fruitslot.tax,
            user_play_log_fruitslot.bet_time
           FROM user_play_log_fruitslot
           WHERE ((user_play_log_fruitslot.is_robot = 0) AND (user_play_log_fruitslot.bet_time > (now() - '3 mons'::interval)))
        UNION
         SELECT user_play_log_rcfishing.lognumber,
            user_play_log_rcfishing.level_code,
            user_play_log_rcfishing.user_id,
            user_play_log_rcfishing.game_id,
            user_play_log_rcfishing.ya_score,
            user_play_log_rcfishing.valid_score,
            user_play_log_rcfishing.de_score,
            user_play_log_rcfishing.tax,
            user_play_log_rcfishing.bet_time
           FROM user_play_log_rcfishing
           WHERE ((user_play_log_rcfishing.is_robot = 0) AND (user_play_log_rcfishing.bet_time > (now() - '3 mons'::interval)))) tmp
  GROUP BY (date_trunc('hour'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;

DROP VIEW "public"."view_agent_game_room_ratio_log";
CREATE OR REPLACE VIEW "public"."view_agent_game_room_ratio_log" AS  SELECT date_trunc('day'::text, tmp.bet_time) AS log_time,
    tmp.agent_id,
    tmp.level_code,
    tmp.game_id,
    tmp.room_type,
    count(DISTINCT tmp.user_id) AS bet_user,
    count(tmp.lognumber) AS bet_count,
    sum(tmp.ya_score) AS sum_ya,
    sum(tmp.valid_score) AS sum_valid_ya,
    sum(tmp.de_score) AS sum_de,
    0 AS sum_bonus,
    sum(tmp.tax) AS sum_tax
   FROM ( SELECT user_play_log_baccarat.lognumber,
            user_play_log_baccarat.agent_id,
            user_play_log_baccarat.level_code,
            user_play_log_baccarat.user_id,
            user_play_log_baccarat.game_id,
            user_play_log_baccarat.room_type,
            user_play_log_baccarat.ya_score,
            user_play_log_baccarat.valid_score,
            user_play_log_baccarat.de_score,
            user_play_log_baccarat.tax,
            user_play_log_baccarat.bet_time
           FROM user_play_log_baccarat
          WHERE user_play_log_baccarat.is_robot = 0 AND user_play_log_baccarat.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_colordisc.lognumber,
            user_play_log_colordisc.agent_id,
            user_play_log_colordisc.level_code,
            user_play_log_colordisc.user_id,
            user_play_log_colordisc.game_id,
            user_play_log_colordisc.room_type,
            user_play_log_colordisc.ya_score,
            user_play_log_colordisc.valid_score,
            user_play_log_colordisc.de_score,
            user_play_log_colordisc.tax,
            user_play_log_colordisc.bet_time
           FROM user_play_log_colordisc
          WHERE user_play_log_colordisc.is_robot = 0 AND user_play_log_colordisc.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_fantan.lognumber,
            user_play_log_fantan.agent_id,
            user_play_log_fantan.level_code,
            user_play_log_fantan.user_id,
            user_play_log_fantan.game_id,
            user_play_log_fantan.room_type,
            user_play_log_fantan.ya_score,
            user_play_log_fantan.valid_score,
            user_play_log_fantan.de_score,
            user_play_log_fantan.tax,
            user_play_log_fantan.bet_time
           FROM user_play_log_fantan
          WHERE user_play_log_fantan.is_robot = 0 AND user_play_log_fantan.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_prawncrab.lognumber,
            user_play_log_prawncrab.agent_id,
            user_play_log_prawncrab.level_code,
            user_play_log_prawncrab.user_id,
            user_play_log_prawncrab.game_id,
            user_play_log_prawncrab.room_type,
            user_play_log_prawncrab.ya_score,
            user_play_log_prawncrab.valid_score,
            user_play_log_prawncrab.de_score,
            user_play_log_prawncrab.tax,
            user_play_log_prawncrab.bet_time
           FROM user_play_log_prawncrab
          WHERE user_play_log_prawncrab.is_robot = 0 AND user_play_log_prawncrab.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_blackjack.lognumber,
            user_play_log_blackjack.agent_id,
            user_play_log_blackjack.level_code,
            user_play_log_blackjack.user_id,
            user_play_log_blackjack.game_id,
            user_play_log_blackjack.room_type,
            user_play_log_blackjack.ya_score,
            user_play_log_blackjack.valid_score,
            user_play_log_blackjack.de_score,
            user_play_log_blackjack.tax,
            user_play_log_blackjack.bet_time
           FROM user_play_log_blackjack
          WHERE user_play_log_blackjack.is_robot = 0 AND user_play_log_blackjack.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_sangong.lognumber,
            user_play_log_sangong.agent_id,
            user_play_log_sangong.level_code,
            user_play_log_sangong.user_id,
            user_play_log_sangong.game_id,
            user_play_log_sangong.room_type,
            user_play_log_sangong.ya_score,
            user_play_log_sangong.valid_score,
            user_play_log_sangong.de_score,
            user_play_log_sangong.tax,
            user_play_log_sangong.bet_time
           FROM user_play_log_sangong
          WHERE user_play_log_sangong.is_robot = 0 AND user_play_log_sangong.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_hundredsicbo.lognumber,
            user_play_log_hundredsicbo.agent_id,
            user_play_log_hundredsicbo.level_code,
            user_play_log_hundredsicbo.user_id,
            user_play_log_hundredsicbo.game_id,
            user_play_log_hundredsicbo.room_type,
            user_play_log_hundredsicbo.ya_score,
            user_play_log_hundredsicbo.valid_score,
            user_play_log_hundredsicbo.de_score,
            user_play_log_hundredsicbo.tax,
            user_play_log_hundredsicbo.bet_time
           FROM user_play_log_hundredsicbo
          WHERE user_play_log_hundredsicbo.is_robot = 0 AND user_play_log_hundredsicbo.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_bullbull.lognumber,
            user_play_log_bullbull.agent_id,
            user_play_log_bullbull.level_code,
            user_play_log_bullbull.user_id,
            user_play_log_bullbull.game_id,
            user_play_log_bullbull.room_type,
            user_play_log_bullbull.ya_score,
            user_play_log_bullbull.valid_score,
            user_play_log_bullbull.de_score,
            user_play_log_bullbull.tax,
            user_play_log_bullbull.bet_time
           FROM user_play_log_bullbull
          WHERE user_play_log_bullbull.is_robot = 0 AND user_play_log_bullbull.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_texas.lognumber,
            user_play_log_texas.agent_id,
            user_play_log_texas.level_code,
            user_play_log_texas.user_id,
            user_play_log_texas.game_id,
            user_play_log_texas.room_type,
            user_play_log_texas.ya_score,
            user_play_log_texas.valid_score,
            user_play_log_texas.de_score,
            user_play_log_texas.tax,
            user_play_log_texas.bet_time
           FROM user_play_log_texas
          WHERE user_play_log_texas.is_robot = 0 AND user_play_log_texas.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_fruitslot.lognumber,
            user_play_log_fruitslot.agent_id,
            user_play_log_fruitslot.level_code,
            user_play_log_fruitslot.user_id,
            user_play_log_fruitslot.game_id,
            user_play_log_fruitslot.room_type,
            user_play_log_fruitslot.ya_score,
            user_play_log_fruitslot.valid_score,
            user_play_log_fruitslot.de_score,
            user_play_log_fruitslot.tax,
            user_play_log_fruitslot.bet_time
           FROM user_play_log_fruitslot
          WHERE user_play_log_fruitslot.is_robot = 0 AND user_play_log_fruitslot.bet_time > (now() - '3 mons'::interval)
        UNION ALL
         SELECT user_play_log_rcfishing.lognumber,
            user_play_log_rcfishing.agent_id,
            user_play_log_rcfishing.level_code,
            user_play_log_rcfishing.user_id,
            user_play_log_rcfishing.game_id,
            user_play_log_rcfishing.room_type,
            user_play_log_rcfishing.ya_score,
            user_play_log_rcfishing.valid_score,
            user_play_log_rcfishing.de_score,
            user_play_log_rcfishing.tax,
            user_play_log_rcfishing.bet_time
           FROM user_play_log_rcfishing
          WHERE user_play_log_rcfishing.is_robot = 0 AND user_play_log_rcfishing.bet_time > (now() - '3 mons'::interval)) tmp
  GROUP BY (date_trunc('day'::text, tmp.bet_time)), tmp.agent_id, tmp.level_code, tmp.game_id, tmp.room_type;

-- 代理遊戲平台營收統計usp刪除tax及bonus --
DROP PROCEDURE IF EXISTS "public"."usp_insert_agent_game_ratio_stat"("_id" varchar, "_level_code" varchar, "_agent_id" int4, "_game_id" int4, "_game_type" int4, "_room_type" int4, "_de" numeric, "_ya" numeric, "_vaild_ya" numeric, "_tax" numeric, "_bonus" numeric, "_play_count" int4, "_bet_time" timestamptz);
CREATE OR REPLACE PROCEDURE "public"."usp_insert_agent_game_ratio_stat"("_id" varchar, "_level_code" varchar, "_agent_id" int4, "_game_id" int4, "_game_type" int4, "_room_type" int4, "_de" numeric, "_ya" numeric, "_vaild_ya" numeric, "_play_count" int4, "_bet_time" timestamptz)
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
	
	EXECUTE 'INSERT INTO "public"."'||_stat_table_name||'"("id", "level_code", "agent_id", "game_id", "game_type", "room_type", "de", "ya", "vaild_ya", "play_count", "update_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
			ON CONFLICT ON CONSTRAINT "'||_stat_table_name||'_pkey" DO 
			UPDATE SET 	
				de = '||_stat_table_name||'.de + EXCLUDED.de,
				ya = '||_stat_table_name||'.ya + EXCLUDED.ya,
				vaild_ya = '||_stat_table_name||'.vaild_ya + EXCLUDED.vaild_ya,
				play_count = '||_stat_table_name||'.play_count + EXCLUDED.play_count,
				update_time = EXCLUDED.update_time
						;'
				USING
					_id, _level_code, _agent_id, _game_id, _game_type, _room_type, _de, _ya, _vaild_ya, 1, now();

	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = _stat_month_table_name) THEN
		CALL "public"."usp_check_agent_game_ratio_stat"(_table_month);
	END IF;
	
	_log_time := to_char(_bet_time,'YYYYMMDD');
	
	EXECUTE 'INSERT INTO "public"."'||_stat_month_table_name||'"("log_time", "id", "level_code", "agent_id", "game_id", "game_type", "room_type", "de", "ya", "vaild_ya", "play_count", "update_time")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
			ON CONFLICT ON CONSTRAINT "'||_stat_month_table_name||'_pkey" DO 
			UPDATE SET 	
				de = '||_stat_month_table_name||'.de + EXCLUDED.de,
				ya = '||_stat_month_table_name||'.ya + EXCLUDED.ya,
				vaild_ya = '||_stat_month_table_name||'.vaild_ya + EXCLUDED.vaild_ya,
				play_count = '||_stat_month_table_name||'.play_count + EXCLUDED.play_count,
				update_time = EXCLUDED.update_time
						;'
				USING
					_log_time, _id, _level_code, _agent_id, _game_id, _game_type, _room_type, _de, _ya, _vaild_ya, 1, now();
	
END$BODY$
  LANGUAGE plpgsql;

-- 代理遊戲平台營收統計總表刪除tax及bonus --
DROP PROCEDURE IF EXISTS "public"."usp_check_agent_game_ratio_stat"("_thismonth" varchar);
CREATE OR REPLACE PROCEDURE "public"."usp_check_agent_game_ratio_stat"("_thismonth" varchar)
 AS $BODY$
DECLARE _table_name varchar(50);
DECLARE _comment_table_name varchar(50);
BEGIN

	IF _thismonth = '' THEN
		_table_name := 'agent_game_ratio_stat';
		_comment_table_name := '代理遊戲平台營收統計總表';
		
		IF NOT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_catalog = current_database() AND table_schema = 'public' AND table_name = _table_name) THEN
			EXECUTE '
				CREATE TABLE IF NOT EXISTS "public"."'||_table_name||'" (
					"id" varchar(128) NOT NULL DEFAULT '''',
					"level_code" varchar(128) NOT NULL DEFAULT '''',
					"agent_id" int4 NOT NULL DEFAULT 0,
					"game_id" int4 NOT NULL DEFAULT 0,
					"game_type" int4 NOT NULL DEFAULT 0,
					"room_type" int4 NOT NULL DEFAULT 0,
					"de" numeric(20,4) NOT NULL DEFAULT 0,
					"ya" numeric(20,4) NOT NULL DEFAULT 0,
					"vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
					"play_count" int8 NOT NULL DEFAULT 0,
					"update_time" timestamptz(6) NOT NULL DEFAULT now(),
					CONSTRAINT "'||_table_name||'_pkey" PRIMARY KEY ("id")
				);';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_1" ON "public"."'||_table_name||'"("update_time");';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_2" ON "public"."'||_table_name||'"("agent_id","game_id","game_type","room_type");';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."id" IS ''pKey 組合'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."level_code" IS ''層級碼'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."agent_id" IS ''代理id'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_id" IS ''遊戲id'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_type" IS ''遊戲類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."room_type" IS ''房間類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."de" IS ''得分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."ya" IS ''壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."vaild_ya" IS ''有效壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."play_count" IS ''下注次數'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."update_time" IS ''最後更新時間'';';
			EXECUTE 'COMMENT ON TABLE "public"."'||_table_name||'" IS '''||_comment_table_name||''';';
		END IF;
	ELSE
		_table_name := 'agent_game_ratio_stat'||'_'||_thismonth;
		_comment_table_name := '代理遊戲平台營收統計表';
	
		IF NOT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_catalog = current_database() AND table_schema = 'public' AND table_name = _table_name) THEN
			EXECUTE '
				CREATE TABLE IF NOT EXISTS "public"."'||_table_name||'" (
					"log_time" varchar(18) NOT NULL DEFAULT '''',
					"id" varchar(128) NOT NULL DEFAULT '''',
					"level_code" varchar(128) NOT NULL DEFAULT '''',
					"agent_id" int4 NOT NULL DEFAULT 0,
					"game_id" int4 NOT NULL DEFAULT 0,
					"game_type" int4 NOT NULL DEFAULT 0,
					"room_type" int4 NOT NULL DEFAULT 0,
					"de" numeric(20,4) NOT NULL DEFAULT 0,
					"ya" numeric(20,4) NOT NULL DEFAULT 0,
					"vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
					"play_count" int8 NOT NULL DEFAULT 0,
					"update_time" timestamptz(6) NOT NULL DEFAULT now(),
					CONSTRAINT "'||_table_name||'_pkey" PRIMARY KEY ("log_time", "id")
				);';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_1" ON "public"."'||_table_name||'"("update_time");';
			EXECUTE 'CREATE INDEX IF NOT EXISTS "idx_'||_table_name||'_2" ON "public"."'||_table_name||'"("agent_id","game_id","game_type","room_type");';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."id" IS ''pKey 組合'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."level_code" IS ''層級碼'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."agent_id" IS ''代理id'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_id" IS ''遊戲id'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."game_type" IS ''遊戲類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."room_type" IS ''房間類型'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."de" IS ''得分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."ya" IS ''壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."vaild_ya" IS ''有效壓分'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."play_count" IS ''下注次數'';';
			EXECUTE 'COMMENT ON COLUMN "public"."'||_table_name||'"."update_time" IS ''最後更新時間'';';
			EXECUTE 'COMMENT ON TABLE "public"."'||_table_name||'" IS '''||_comment_table_name||''';';
		END IF;
	END IF;

END$BODY$
  LANGUAGE plpgsql;

-- 統計玩家usp下注資料刪除bonus --
DROP PROCEDURE IF EXISTS "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_tax" float8, "_bonus" float8, "_play_count" int4, "_big_win_count" int4, "_win_count" int4, "_lose_count" int4, "_last_bet_time" timestamptz);
CREATE OR REPLACE PROCEDURE "public"."usp_game_users_stat"("_agent_id" int4, "_level_code" varchar, "_game_users_id" int4, "_de" float8, "_ya" float8, "_vaild_ya" float8, "_tax" float8, "_play_count" int4, "_big_win_count" int4, "_win_count" int4, "_lose_count" int4, "_last_bet_time" timestamptz)
 AS $BODY$
 DECLARE _log_hour varchar(12);
 BEGIN
	-- Routine body goes here...
	_log_hour := to_char(_last_bet_time,'YYYYMMDDHH24');

	EXECUTE 'INSERT INTO "public"."game_users_stat"(
		"agent_id", "level_code", "game_users_id", "de", "ya",
		"vaild_ya", "tax", "play_count", "big_win_count", "first_bet_time",
		"last_bet_time", "win_count", "lose_count")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			ON CONFLICT ON CONSTRAINT "game_users_stat_pkey" DO 
			UPDATE SET 	
				de = game_users_stat.de + EXCLUDED.de,
				ya = game_users_stat.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat.vaild_ya + EXCLUDED.vaild_ya,
				tax = game_users_stat.tax + EXCLUDED.tax,
				play_count = game_users_stat.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat.big_win_count + EXCLUDED.big_win_count,
                win_count = game_users_stat.win_count + EXCLUDED.win_count,
				lose_count = game_users_stat.lose_count + EXCLUDED.lose_count,
				last_bet_time =  $14,
				update_time = now()
			;'
	USING
	_agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _tax, _play_count, _big_win_count, _last_bet_time, _last_bet_time, _win_count, _lose_count, _last_bet_time;
				
	EXECUTE 'INSERT INTO "public"."game_users_stat_hour"(
		"log_time", "agent_id", "level_code", "game_users_id", "de",
		"ya",	"vaild_ya", "tax", "play_count", "big_win_count",
		"win_count", "lose_count")
			SELECT 
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
			ON CONFLICT ON CONSTRAINT "game_users_stat_hour_pkey" DO 
			UPDATE SET 	
				de = game_users_stat_hour.de + EXCLUDED.de,
				ya = game_users_stat_hour.ya + EXCLUDED.ya,
				vaild_ya = game_users_stat_hour.vaild_ya + EXCLUDED.vaild_ya,
				tax = game_users_stat_hour.tax + EXCLUDED.tax,
				play_count = game_users_stat_hour.play_count + EXCLUDED.play_count,
				big_win_count = game_users_stat_hour.big_win_count + EXCLUDED.big_win_count,
				win_count = game_users_stat_hour.win_count + EXCLUDED.win_count,
				lose_count = game_users_stat_hour.lose_count + EXCLUDED.lose_count
			;'
	USING
	_log_hour, _agent_id, _level_code, _game_users_id, _de, _ya , _vaild_ya, _tax, _play_count, _big_win_count, _win_count, _lose_count;

END$BODY$
  LANGUAGE plpgsql;

-- 每小時統計玩家下注資料表刪除bonus --
ALTER TABLE "public"."game_users_stat_hour" DROP COLUMN "bonus";

-- 統計玩家下注資料表刪除bonus --
ALTER TABLE "public"."game_users_stat" DROP COLUMN "bonus";

-- view_user_play_log刪除bonus --
DROP VIEW "public"."view_user_play_log";
CREATE OR REPLACE VIEW "public"."view_user_play_log" AS
 SELECT tmp.bet_id,
    tmp.lognumber,
    tmp.agent_id,
    tmp.game_id,
    tmp.room_type,
    tmp.desk_id,
    tmp.seat_id,
    tmp.exchange,
    tmp.de_score,
    tmp.ya_score,
    tmp.valid_score,
    tmp.start_score,
    tmp.end_score,
    tmp.create_time,
    tmp.is_robot,
    tmp.is_big_win,
    tmp.is_issue,
    tmp.bet_time,
    tmp.tax,
    tmp.level_code,
    tmp.username,
    tmp.kill_type
   FROM ( SELECT baccarat.bet_id,
            baccarat.lognumber,
            baccarat.agent_id,
            baccarat.game_id,
            baccarat.room_type,
            baccarat.desk_id,
            baccarat.seat_id,
            baccarat.exchange,
            baccarat.de_score,
            baccarat.ya_score,
            baccarat.valid_score,
            baccarat.start_score,
            baccarat.end_score,
            baccarat.create_time,
            baccarat.is_robot,
            baccarat.is_big_win,
            baccarat.is_issue,
            baccarat.bet_time,
            baccarat.tax,
            baccarat.level_code,
            baccarat.username,
            baccarat.kill_type
           FROM user_play_log_baccarat baccarat
        UNION
         SELECT fantan.bet_id,
            fantan.lognumber,
            fantan.agent_id,
            fantan.game_id,
            fantan.room_type,
            fantan.desk_id,
            fantan.seat_id,
            fantan.exchange,
            fantan.de_score,
            fantan.ya_score,
            fantan.valid_score,
            fantan.start_score,
            fantan.end_score,
            fantan.create_time,
            fantan.is_robot,
            fantan.is_big_win,
            fantan.is_issue,
            fantan.bet_time,
            fantan.tax,
            fantan.level_code,
            fantan.username,
            fantan.kill_type
           FROM user_play_log_fantan fantan
        UNION
         SELECT colordisc.bet_id,
            colordisc.lognumber,
            colordisc.agent_id,
            colordisc.game_id,
            colordisc.room_type,
            colordisc.desk_id,
            colordisc.seat_id,
            colordisc.exchange,
            colordisc.de_score,
            colordisc.ya_score,
            colordisc.valid_score,
            colordisc.start_score,
            colordisc.end_score,
            colordisc.create_time,
            colordisc.is_robot,
            colordisc.is_big_win,
            colordisc.is_issue,
            colordisc.bet_time,
            colordisc.tax,
            colordisc.level_code,
            colordisc.username,
            colordisc.kill_type
           FROM user_play_log_colordisc colordisc
        UNION
         SELECT prawncrab.bet_id,
            prawncrab.lognumber,
            prawncrab.agent_id,
            prawncrab.game_id,
            prawncrab.room_type,
            prawncrab.desk_id,
            prawncrab.seat_id,
            prawncrab.exchange,
            prawncrab.de_score,
            prawncrab.ya_score,
            prawncrab.valid_score,
            prawncrab.start_score,
            prawncrab.end_score,
            prawncrab.create_time,
            prawncrab.is_robot,
            prawncrab.is_big_win,
            prawncrab.is_issue,
            prawncrab.bet_time,
            prawncrab.tax,
            prawncrab.level_code,
            prawncrab.username,
            prawncrab.kill_type
           FROM user_play_log_prawncrab prawncrab
        UNION
         SELECT hundredsicbo.bet_id,
            hundredsicbo.lognumber,
            hundredsicbo.agent_id,
            hundredsicbo.game_id,
            hundredsicbo.room_type,
            hundredsicbo.desk_id,
            hundredsicbo.seat_id,
            hundredsicbo.exchange,
            hundredsicbo.de_score,
            hundredsicbo.ya_score,
            hundredsicbo.valid_score,
            hundredsicbo.start_score,
            hundredsicbo.end_score,
            hundredsicbo.create_time,
            hundredsicbo.is_robot,
            hundredsicbo.is_big_win,
            hundredsicbo.is_issue,
            hundredsicbo.bet_time,
            hundredsicbo.tax,
            hundredsicbo.level_code,
            hundredsicbo.username,
            hundredsicbo.kill_type
           FROM user_play_log_hundredsicbo hundredsicbo
        UNION
         SELECT blackjack.bet_id,
            blackjack.lognumber,
            blackjack.agent_id,
            blackjack.game_id,
            blackjack.room_type,
            blackjack.desk_id,
            blackjack.seat_id,
            blackjack.exchange,
            blackjack.de_score,
            blackjack.ya_score,
            blackjack.valid_score,
            blackjack.start_score,
            blackjack.end_score,
            blackjack.create_time,
            blackjack.is_robot,
            blackjack.is_big_win,
            blackjack.is_issue,
            blackjack.bet_time,
            blackjack.tax,
            blackjack.level_code,
            blackjack.username,
            blackjack.kill_type
           FROM user_play_log_blackjack blackjack
        UNION
         SELECT sangong.bet_id,
            sangong.lognumber,
            sangong.agent_id,
            sangong.game_id,
            sangong.room_type,
            sangong.desk_id,
            sangong.seat_id,
            sangong.exchange,
            sangong.de_score,
            sangong.ya_score,
            sangong.valid_score,
            sangong.start_score,
            sangong.end_score,
            sangong.create_time,
            sangong.is_robot,
            sangong.is_big_win,
            sangong.is_issue,
            sangong.bet_time,
            sangong.tax,
            sangong.level_code,
            sangong.username,
            sangong.kill_type
           FROM user_play_log_sangong sangong
        UNION
         SELECT bullbull.bet_id,
            bullbull.lognumber,
            bullbull.agent_id,
            bullbull.game_id,
            bullbull.room_type,
            bullbull.desk_id,
            bullbull.seat_id,
            bullbull.exchange,
            bullbull.de_score,
            bullbull.ya_score,
            bullbull.valid_score,
            bullbull.start_score,
            bullbull.end_score,
            bullbull.create_time,
            bullbull.is_robot,
            bullbull.is_big_win,
            bullbull.is_issue,
            bullbull.bet_time,
            bullbull.tax,
            bullbull.level_code,
            bullbull.username,
            bullbull.kill_type
           FROM user_play_log_bullbull bullbull
        UNION
         SELECT texas.bet_id,
            texas.lognumber,
            texas.agent_id,
            texas.game_id,
            texas.room_type,
            texas.desk_id,
            texas.seat_id,
            texas.exchange,
            texas.de_score,
            texas.ya_score,
            texas.valid_score,
            texas.start_score,
            texas.end_score,
            texas.create_time,
            texas.is_robot,
            texas.is_big_win,
            texas.is_issue,
            texas.bet_time,
            texas.tax,
            texas.level_code,
            texas.username,
            texas.kill_type
           FROM user_play_log_texas texas
        UNION
         SELECT fruitslot.bet_id,
            fruitslot.lognumber,
            fruitslot.agent_id,
            fruitslot.game_id,
            fruitslot.room_type,
            fruitslot.desk_id,
            fruitslot.seat_id,
            fruitslot.exchange,
            fruitslot.de_score,
            fruitslot.ya_score,
            fruitslot.valid_score,
            fruitslot.start_score,
            fruitslot.end_score,
            fruitslot.create_time,
            fruitslot.is_robot,
            fruitslot.is_big_win,
            fruitslot.is_issue,
            fruitslot.bet_time,
            fruitslot.tax,
            fruitslot.level_code,
            fruitslot.username,
            fruitslot.kill_type
           FROM user_play_log_fruitslot fruitslot
        UNION
         SELECT rcfishing.bet_id,
            rcfishing.lognumber,
            rcfishing.agent_id,
            rcfishing.game_id,
            rcfishing.room_type,
            rcfishing.desk_id,
            rcfishing.seat_id,
            rcfishing.exchange,
            rcfishing.de_score,
            rcfishing.ya_score,
            rcfishing.valid_score,
            rcfishing.start_score,
            rcfishing.end_score,
            rcfishing.create_time,
            rcfishing.is_robot,
            rcfishing.is_big_win,
            rcfishing.is_issue,
            rcfishing.bet_time,
            rcfishing.tax,
            rcfishing.level_code,
            rcfishing.username,
            rcfishing.kill_type
           FROM user_play_log_rcfishing rcfishing) tmp;

-- 三國捕魚紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_rcfishing" DROP COLUMN "bonus";

-- 水果機紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_fruitslot" DROP COLUMN "bonus";

-- 德州撲克紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_texas" DROP COLUMN "bonus";

-- 牛牛紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_bullbull" DROP COLUMN "bonus";

-- 三公紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_sangong" DROP COLUMN "bonus";

-- 21點紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_blackjack" DROP COLUMN "bonus";

-- 百人骰寶紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_hundredsicbo" DROP COLUMN "bonus";

-- 魚蝦蟹紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_prawncrab" DROP COLUMN "bonus";

-- 色碟紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_colordisc" DROP COLUMN "bonus";

-- 番攤紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_fantan" DROP COLUMN "bonus";

-- 百家樂紀錄刪除bonus --
ALTER TABLE "public"."user_play_log_baccarat" DROP COLUMN "bonus";

-- 局紀錄刪除bonus --
ALTER TABLE "public"."play_log_common" DROP COLUMN "bonus";