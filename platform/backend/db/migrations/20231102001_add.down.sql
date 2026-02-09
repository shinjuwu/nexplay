-- 刪除安達巴哈遊戲紀錄表 --
DROP TABLE "public"."user_play_log_andarbahar";

-- 刪除代理安達巴哈遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (1009)
);

-- 刪除代理安達巴哈遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (1009);

-- 刪除安達巴哈遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (1009);

-- 刪除安達巴哈遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (1009);