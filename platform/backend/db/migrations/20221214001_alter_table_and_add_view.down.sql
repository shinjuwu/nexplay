DROP VIEW "public"."v_user_play_log";

ALTER TABLE "public"."user_play_log_baccarat"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_baccarat_bet_id_seq');
ALTER TABLE "public"."user_play_log_baccarat"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_baccarat_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_fantan"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_fantan_bet_id_seq');
ALTER TABLE "public"."user_play_log_fantan"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_fantan_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_colordisc"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_colordisc_bet_id_seq');
ALTER TABLE "public"."user_play_log_colordisc"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_colordisc_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_prawncrab"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_prawncrab_bet_id_seq');
ALTER TABLE "public"."user_play_log_prawncrab"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_prawncrab_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_hundredsicbo"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_hundredsicbo_bet_id_seq');
ALTER TABLE "public"."user_play_log_hundredsicbo"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_hundredsicbo_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_blackjack"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_blackjack_bet_id_seq');
ALTER TABLE "public"."user_play_log_blackjack"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_blackjack_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_sangong"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_sangong_bet_id_seq');
ALTER TABLE "public"."user_play_log_sangong"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_sangong_bet_id_seq'),
  DROP COLUMN "username";

ALTER TABLE "public"."user_play_log_bullbull"
  ALTER COLUMN "bet_id" SET DEFAULT nextval('user_play_log_bullbull_bet_id_seq');
ALTER TABLE "public"."user_play_log_bullbull"
  ALTER COLUMN "bet_id" TYPE int8 USING nextval('user_play_log_bullbull_bet_id_seq'),
  DROP COLUMN "username";
