ALTER TABLE "public"."user_play_log_hundredsicbo"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_texas"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_texas"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_blackjack"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_blackjack"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_sangong"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_sangong"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_bullbull"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_bullbull"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_baccarat"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_baccarat"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_prawncrab"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_prawncrab"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_fantan"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_fantan"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_colordisc"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_colordisc"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_fruitslot"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_fruitslot"."kill_type" IS '殺放狀態';

ALTER TABLE "public"."user_play_log_rcfishing"
ADD COLUMN "kill_type" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_rcfishing"."kill_type" IS '殺放狀態';

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