
ALTER TABLE "public"."play_log_baccarat" RENAME TO "user_play_log_baccarat";
ALTER SEQUENCE "public"."play_log_baccarat_bet_id_seq" RENAME TO "user_play_log_baccarat_bet_id_seq";

-- ALTER TABLE "public"."play_log_common" RENAME TO "play_log_baccarat";
-- ALTER INDEX "public"."idx_play_log_common_1" RENAME TO "idx_play_log_baccarat_1";