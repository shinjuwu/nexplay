ALTER TABLE "public"."play_log_common" 
  ADD COLUMN "sync" bool NOT NULL DEFAULT false;

COMMENT ON COLUMN "public"."play_log_common"."sync" IS '注單服務同步旗標';

ALTER TABLE "public"."play_log_common" 
  ADD COLUMN "start_time" timestamptz NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  ADD COLUMN "end_time" timestamptz NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone;

COMMENT ON COLUMN "public"."play_log_common"."start_time" IS '遊戲開始時間';

COMMENT ON COLUMN "public"."play_log_common"."end_time" IS '遊戲結束時間';

CREATE INDEX "idx_play_log_common_sync" ON "public"."play_log_common" USING btree (
  "sync"
);

CREATE INDEX "idx_user_play_log_sangong_bet_time" ON "public"."user_play_log_sangong" USING btree (
	 "bet_time"
  );
CREATE INDEX "idx_user_play_log_rcfishing_bet_time" ON "public"."user_play_log_rcfishing" USING btree (
	 "bet_time"
  ); 
	CREATE INDEX "idx_user_play_log_fantan_bet_time" ON "public"."user_play_log_fantan" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_hundredsicbo_bet_time" ON "public"."user_play_log_hundredsicbo" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_blackjack_bet_time" ON "public"."user_play_log_blackjack" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_fruitslot_bet_time" ON "public"."user_play_log_fruitslot" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_cockfight_bet_time" ON "public"."user_play_log_cockfight" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_dogracing_bet_time" ON "public"."user_play_log_dogracing" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_rummy_bet_time" ON "public"."user_play_log_rummy" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_goldenflower_bet_time" ON "public"."user_play_log_goldenflower" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_pokdeng_bet_time" ON "public"."user_play_log_pokdeng" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_prawncrab_bet_time" ON "public"."user_play_log_prawncrab" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_colordisc_bet_time" ON "public"."user_play_log_colordisc" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_baccarat_bet_time" ON "public"."user_play_log_baccarat" USING btree (
	 "bet_time"
  ); 

	CREATE INDEX "idx_user_play_log_texas_bet_time" ON "public"."user_play_log_texas" USING btree (
	 "bet_time"
  ); 
 
	CREATE INDEX "idx_user_play_log_bullbull_bet_time" ON "public"."user_play_log_bullbull" USING btree (
	 "bet_time"
  ); 
