ALTER TABLE "public"."user_play_log_baccarat" 
  DROP CONSTRAINT "play_log_baccarat_pkey",
  ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id");

ALTER TABLE "public"."user_play_log_fantan" 
  DROP CONSTRAINT "play_log_fantan_pkey",
  ADD CONSTRAINT "play_log_fantan_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id");

ALTER TABLE "public"."user_play_log_colordisc" 
  DROP CONSTRAINT "play_log_colordisc_pkey",
  ADD CONSTRAINT "play_log_colordisc_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id");
