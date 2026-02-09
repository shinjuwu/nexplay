ALTER TABLE "public"."user_play_log_andarbahar"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_baccarat"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_blackjack"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_bullbull"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_catte"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_chinesepoker"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_cockfight"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_colordisc"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_dogracing"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_fantan"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_fruit777slot"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_fruitslot"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_goldenflower"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_hundredsicbo"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_megsharkslot"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_midasslot"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_okey"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_plinko"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_pokdeng"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_prawncrab"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_rcfishing"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_rocket"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_rummy"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_sangong"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_teenpatti"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;
ALTER TABLE "public"."user_play_log_texas"
  ADD COLUMN "kill_prob" numeric(20,4) NOT NULL DEFAULT 0,
  ADD COLUMN "kill_level" int2 NOT NULL DEFAULT -1,
  ADD COLUMN "real_players" int2 NOT NULL DEFAULT -1;

COMMENT ON COLUMN "public"."user_play_log_andarbahar"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_catte"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_fantan"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_midasslot"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_okey"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_plinko"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_rocket"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_rummy"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_sangong"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."kill_prob" IS '殺放設定機率';
COMMENT ON COLUMN "public"."user_play_log_texas"."kill_prob" IS '殺放設定機率';

COMMENT ON COLUMN "public"."user_play_log_andarbahar"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_catte"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fantan"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_midasslot"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_okey"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_plinko"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rocket"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rummy"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_sangong"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."kill_level" IS '殺放層級 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_texas"."kill_level" IS '殺放層級 (預設值-1)';

COMMENT ON COLUMN "public"."user_play_log_andarbahar"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_baccarat"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_blackjack"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_bullbull"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_catte"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_chinesepoker"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_cockfight"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_colordisc"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_dogracing"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fantan"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fruit777slot"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_fruitslot"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_goldenflower"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_hundredsicbo"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_megsharkslot"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_midasslot"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_okey"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_plinko"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_pokdeng"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_prawncrab"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rcfishing"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rocket"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_rummy"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_sangong"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_teenpatti"."real_players" IS '真實玩家人數 (預設值-1)';
COMMENT ON COLUMN "public"."user_play_log_texas"."real_players" IS '真實玩家人數 (預設值-1)';

