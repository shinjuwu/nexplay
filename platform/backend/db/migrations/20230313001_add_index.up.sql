CREATE INDEX "idx_play_log_common_lognumber" ON "public"."play_log_common" ("lognumber" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_baccarat_lognumber" ON "public"."user_play_log_baccarat" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_baccarat_betid" ON "public"."user_play_log_baccarat" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_fantan_lognumber" ON "public"."user_play_log_fantan" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_fantan_betid" ON "public"."user_play_log_fantan" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_colordisc_lognumber" ON "public"."user_play_log_colordisc" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_colordisc_betid" ON "public"."user_play_log_colordisc" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_prawncrab_lognumber" ON "public"."user_play_log_prawncrab" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_prawncrab_betid" ON "public"."user_play_log_prawncrab" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_hundredsicbo_lognumber" ON "public"."user_play_log_hundredsicbo" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_hundredsicbo_betid" ON "public"."user_play_log_hundredsicbo" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_blackjack_lognumber" ON "public"."user_play_log_blackjack" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_blackjack_betid" ON "public"."user_play_log_blackjack" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_sangong_lognumber" ON "public"."user_play_log_sangong" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_sangong_betid" ON "public"."user_play_log_sangong" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_bullbull_lognumber" ON "public"."user_play_log_bullbull" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_bullbull_betid" ON "public"."user_play_log_bullbull" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_texas_lognumber" ON "public"."user_play_log_texas" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_texas_betid" ON "public"."user_play_log_texas" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_fruitslot_lognumber" ON "public"."user_play_log_fruitslot" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_fruitslot_betid" ON "public"."user_play_log_fruitslot" ("bet_id" DESC NULLS LAST);

CREATE INDEX "idx_user_play_log_rcfishing_lognumber" ON "public"."user_play_log_rcfishing" ("lognumber" DESC NULLS LAST);
CREATE INDEX "idx_user_play_log_rcfishing_betid" ON "public"."user_play_log_rcfishing" ("bet_id" DESC NULLS LAST);
