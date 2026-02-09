-- 調整其他玩家遊戲紀錄表PK條件 --
ALTER TABLE "public"."user_play_log_baccarat" 
  DROP CONSTRAINT "play_log_baccarat_pkey",
  ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type");

ALTER TABLE "public"."user_play_log_fantan" 
  DROP CONSTRAINT "play_log_fantan_pkey",
  ADD CONSTRAINT "play_log_fantan_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type");

ALTER TABLE "public"."user_play_log_colordisc" 
  DROP CONSTRAINT "play_log_colordisc_pkey",
  ADD CONSTRAINT "play_log_colordisc_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type");

-- 刪除魚蝦蟹玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_prawncrab";
