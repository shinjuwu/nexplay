ALTER TABLE "public"."user_play_log_baccarat"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_baccarat"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_fantan"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_fantan"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_fantan"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_colordisc"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_colordisc"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_prawncrab"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_prawncrab"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_hundredsicbo"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_cockfight"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_cockfight"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_dogracing"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_dogracing"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_rocket"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_rocket"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_rocket"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_blackjack"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_blackjack"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_sangong"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_sangong"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_sangong"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_bullbull"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_bullbull"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_texas"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_texas"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_texas"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_rummy"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_rummy"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_rummy"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_goldenflower"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_goldenflower"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_pokdeng"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_pokdeng"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_catte"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_catte"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_catte"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_chinesepoker"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_fruitslot"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_fruitslot"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_rcfishing"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_rcfishing"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_plinko"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_plinko"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_plinko"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."user_play_log_fruit777slot"
  ADD COLUMN "jp_inject_water_rate" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."jp_inject_water_rate" IS 'jp注水%數';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."jp_inject_water_score" IS 'jp注水分數';

ALTER TABLE "public"."rp_agent_stat_15min"
  ADD COLUMN "jackpot_user" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "jackpot_count" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "sum_jp_inject_water_score" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "sum_jp_prize_score" numeric(20,4) NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."rp_agent_stat_15min"."jackpot_user" IS 'jackpot中獎總人數';
COMMENT ON COLUMN "public"."rp_agent_stat_15min"."jackpot_count" IS 'jackpot中獎總單量';
COMMENT ON COLUMN "public"."rp_agent_stat_15min"."sum_jp_inject_water_score" IS '總jp注水分數';
COMMENT ON COLUMN "public"."rp_agent_stat_15min"."sum_jp_prize_score" IS '總jp中獎分數';
