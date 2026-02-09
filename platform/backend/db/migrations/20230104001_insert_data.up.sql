-- 新增德州撲克 --
INSERT INTO "public"."game" ("id", "server_info_code", "name", "code", "state", "type", "room_number", "table_number", "cal_state", "h5_link")
VALUES
  (2004, 'dev', '德州撲克', 'texas', 1, 2, 4, 0, 1, 'http://172.30.0.152/client/vue/');

-- 新增德州撲克遊戲房間 --
INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (20040, '德州撲克新手房', 1, 2004, 0),
  (20041, '德州撲克普通房', 1, 2004, 1),
  (20042, '德州撲克高手房', 1, 2004, 2),
  (20043, '德州撲克大師房', 1, 2004, 3),
  (20044, '德州撲克初學乍練', 1, 2004, 4),
  (20045, '德州撲克略有小成', 1, 2004, 5),
  (20046, '德州撲克得心應手', 1, 2004, 6),
  (20047, '德州撲克爐火純青', 1, 2004, 7);

-- 新增代理德州撲克遊戲設定 --
INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" = 2004;

-- 新增代理德州撲克遊戲房間設定 --
INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr"
  WHERE "gr"."game_id" = 2004;

-- 新增德州撲克遊戲紀錄表 --
CREATE TABLE "public"."user_play_log_texas" (
  "bet_id" varchar(30),
  "lognumber" varchar(100) NOT NULL,
  "agent_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) NOT NULL DEFAULT '',
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "desk_id" int4 NOT NULL,
  "seat_id" int4 NOT NULL,
  "exchange" int4 NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "ya_score" numeric(20,4) NOT NULL,
  "valid_score" numeric(20,4) NOT NULL,
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL,
  "level_code" varchar(128) NOT NULL DEFAULT '',
  CONSTRAINT "play_log_texas_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id")
)
;

COMMENT ON COLUMN "public"."user_play_log_texas"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."user_play_log_texas"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."user_play_log_texas"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."user_play_log_texas"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."user_play_log_texas"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."user_play_log_texas"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."user_play_log_texas"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."user_play_log_texas"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."user_play_log_texas"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."user_play_log_texas"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."user_play_log_texas"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."user_play_log_texas"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."user_play_log_texas"."tax" IS '抽水';
COMMENT ON COLUMN "public"."user_play_log_texas"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."user_play_log_texas"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."user_play_log_texas"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."user_play_log_texas"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."user_play_log_texas"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."user_play_log_texas"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."user_play_log_texas"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "public"."user_play_log_texas"."level_code" IS '代理層級碼';
COMMENT ON TABLE "public"."user_play_log_texas" IS '德州撲克玩家遊戲記錄';

DROP VIEW "public"."v_user_play_log";
CREATE VIEW "public"."v_user_play_log" AS
SELECT *
FROM
(
  SELECT "baccarat"."bet_id",
    "baccarat"."lognumber",
    "baccarat"."agent_id",
    "baccarat"."game_id",
    "baccarat"."room_type",
    "baccarat"."desk_id",
    "baccarat"."seat_id",
    "baccarat"."exchange",
    "baccarat"."de_score",
    "baccarat"."ya_score",
    "baccarat"."valid_score",
    "baccarat"."start_score",
    "baccarat"."end_score",
    "baccarat"."create_time",
    "baccarat"."is_robot",
    "baccarat"."is_big_win",
    "baccarat"."is_issue",
    "baccarat"."bet_time",
    "baccarat"."tax",
    "baccarat"."level_code",
    "baccarat"."username"
    FROM "public"."user_play_log_baccarat" AS "baccarat"
  UNION
  SELECT "fantan"."bet_id",
    "fantan"."lognumber",
    "fantan"."agent_id",
    "fantan"."game_id",
    "fantan"."room_type",
    "fantan"."desk_id",
    "fantan"."seat_id",
    "fantan"."exchange",
    "fantan"."de_score",
    "fantan"."ya_score",
    "fantan"."valid_score",
    "fantan"."start_score",
    "fantan"."end_score",
    "fantan"."create_time",
    "fantan"."is_robot",
    "fantan"."is_big_win",
    "fantan"."is_issue",
    "fantan"."bet_time",
    "fantan"."tax",
    "fantan"."level_code",
    "fantan"."username"
    FROM "public"."user_play_log_fantan" AS "fantan"
  UNION
  SELECT "colordisc"."bet_id",
    "colordisc"."lognumber",
    "colordisc"."agent_id",
    "colordisc"."game_id",
    "colordisc"."room_type",
    "colordisc"."desk_id",
    "colordisc"."seat_id",
    "colordisc"."exchange",
    "colordisc"."de_score",
    "colordisc"."ya_score",
    "colordisc"."valid_score",
    "colordisc"."start_score",
    "colordisc"."end_score",
    "colordisc"."create_time",
    "colordisc"."is_robot",
    "colordisc"."is_big_win",
    "colordisc"."is_issue",
    "colordisc"."bet_time",
    "colordisc"."tax",
    "colordisc"."level_code",
    "colordisc"."username"
    FROM "public"."user_play_log_colordisc" AS "colordisc"
  UNION
  SELECT "prawncrab"."bet_id",
    "prawncrab"."lognumber",
    "prawncrab"."agent_id",
    "prawncrab"."game_id",
    "prawncrab"."room_type",
    "prawncrab"."desk_id",
    "prawncrab"."seat_id",
    "prawncrab"."exchange",
    "prawncrab"."de_score",
    "prawncrab"."ya_score",
    "prawncrab"."valid_score",
    "prawncrab"."start_score",
    "prawncrab"."end_score",
    "prawncrab"."create_time",
    "prawncrab"."is_robot",
    "prawncrab"."is_big_win",
    "prawncrab"."is_issue",
    "prawncrab"."bet_time",
    "prawncrab"."tax",
    "prawncrab"."level_code",
    "prawncrab"."username"
    FROM "public"."user_play_log_prawncrab" AS "prawncrab"
  UNION
  SELECT "hundredsicbo"."bet_id",
    "hundredsicbo"."lognumber",
    "hundredsicbo"."agent_id",
    "hundredsicbo"."game_id",
    "hundredsicbo"."room_type",
    "hundredsicbo"."desk_id",
    "hundredsicbo"."seat_id",
    "hundredsicbo"."exchange",
    "hundredsicbo"."de_score",
    "hundredsicbo"."ya_score",
    "hundredsicbo"."valid_score",
    "hundredsicbo"."start_score",
    "hundredsicbo"."end_score",
    "hundredsicbo"."create_time",
    "hundredsicbo"."is_robot",
    "hundredsicbo"."is_big_win",
    "hundredsicbo"."is_issue",
    "hundredsicbo"."bet_time",
    "hundredsicbo"."tax",
    "hundredsicbo"."level_code",
    "hundredsicbo"."username"
    FROM "public"."user_play_log_hundredsicbo" AS "hundredsicbo"
  UNION
  SELECT "blackjack"."bet_id",
    "blackjack"."lognumber",
    "blackjack"."agent_id",
    "blackjack"."game_id",
    "blackjack"."room_type",
    "blackjack"."desk_id",
    "blackjack"."seat_id",
    "blackjack"."exchange",
    "blackjack"."de_score",
    "blackjack"."ya_score",
    "blackjack"."valid_score",
    "blackjack"."start_score",
    "blackjack"."end_score",
    "blackjack"."create_time",
    "blackjack"."is_robot",
    "blackjack"."is_big_win",
    "blackjack"."is_issue",
    "blackjack"."bet_time",
    "blackjack"."tax",
    "blackjack"."level_code",
    "blackjack"."username"
    FROM "public"."user_play_log_blackjack" AS "blackjack"
  UNION
  SELECT "sangong"."bet_id",
    "sangong"."lognumber",
    "sangong"."agent_id",
    "sangong"."game_id",
    "sangong"."room_type",
    "sangong"."desk_id",
    "sangong"."seat_id",
    "sangong"."exchange",
    "sangong"."de_score",
    "sangong"."ya_score",
    "sangong"."valid_score",
    "sangong"."start_score",
    "sangong"."end_score",
    "sangong"."create_time",
    "sangong"."is_robot",
    "sangong"."is_big_win",
    "sangong"."is_issue",
    "sangong"."bet_time",
    "sangong"."tax",
    "sangong"."level_code",
    "sangong"."username"
    FROM "public"."user_play_log_sangong" AS "sangong"
  UNION
  SELECT "bullbull"."bet_id",
    "bullbull"."lognumber",
    "bullbull"."agent_id",
    "bullbull"."game_id",
    "bullbull"."room_type",
    "bullbull"."desk_id",
    "bullbull"."seat_id",
    "bullbull"."exchange",
    "bullbull"."de_score",
    "bullbull"."ya_score",
    "bullbull"."valid_score",
    "bullbull"."start_score",
    "bullbull"."end_score",
    "bullbull"."create_time",
    "bullbull"."is_robot",
    "bullbull"."is_big_win",
    "bullbull"."is_issue",
    "bullbull"."bet_time",
    "bullbull"."tax",
    "bullbull"."level_code",
    "bullbull"."username"
    FROM "public"."user_play_log_bullbull" AS "bullbull"
  UNION
  SELECT "texas"."bet_id",
    "texas"."lognumber",
    "texas"."agent_id",
    "texas"."game_id",
    "texas"."room_type",
    "texas"."desk_id",
    "texas"."seat_id",
    "texas"."exchange",
    "texas"."de_score",
    "texas"."ya_score",
    "texas"."valid_score",
    "texas"."start_score",
    "texas"."end_score",
    "texas"."create_time",
    "texas"."is_robot",
    "texas"."is_big_win",
    "texas"."is_issue",
    "texas"."bet_time",
    "texas"."tax",
    "texas"."level_code",
    "texas"."username"
    FROM "public"."user_play_log_texas" AS "texas"
) "tmp";

DROP MATERIALIZED VIEW mv_cal_game_stat_hour, mv_cal_game_stat_day, mv_cal_game_stat_week, mv_cal_game_stat_month;

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
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 mons'::interval)))) tmp
  GROUP BY (date_trunc('hour'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;

CREATE MATERIALIZED VIEW IF NOT EXISTS mv_cal_game_stat_day AS
SELECT date_trunc('day'::text, tmp.bet_time) AS log_time,
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
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 mons'::interval)))) tmp
  GROUP BY (date_trunc('day'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;

CREATE MATERIALIZED VIEW IF NOT EXISTS mv_cal_game_stat_week AS
SELECT date_trunc('week'::text, tmp.bet_time) AS log_time,
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
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 mons'::interval)))) tmp
  GROUP BY (date_trunc('week'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;

CREATE MATERIALIZED VIEW IF NOT EXISTS mv_cal_game_stat_month AS
SELECT date_trunc('month'::text, tmp.bet_time) AS log_time,
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
           WHERE ((user_play_log_texas.is_robot = 0) AND (user_play_log_texas.bet_time > (now() - '3 mons'::interval)))) tmp
  GROUP BY (date_trunc('month'::text, tmp.bet_time)), tmp.level_code, tmp.game_id;
