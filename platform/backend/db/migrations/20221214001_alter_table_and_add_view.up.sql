ALTER TABLE "public"."user_play_log_baccarat"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_baccarat" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_fantan"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_fantan" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_colordisc"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_colordisc" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_prawncrab"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_prawncrab" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_hundredsicbo"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_hundredsicbo" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_blackjack"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_blackjack" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_sangong"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_sangong" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

ALTER TABLE "public"."user_play_log_bullbull"
  ALTER COLUMN "bet_id" TYPE varchar(30) USING "game_id" || '-' || "bet_id",
  ALTER COLUMN "bet_id" SET DEFAULT '',
  ADD COLUMN "username" varchar(100) NOT NULL DEFAULT '';
UPDATE "public"."user_play_log_bullbull" AS "upl"
  SET "username" = "gu"."original_username"
  FROM "public"."game_users" AS "gu"
  WHERE "gu"."id" = "upl"."user_id";

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
) "tmp";
