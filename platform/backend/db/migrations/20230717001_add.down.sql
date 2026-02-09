-- 刪除越南Catte遊戲紀錄表 --
DROP TABLE "public"."user_play_log_catte";

-- 刪除代理越南Catte遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" = 2008
);

-- 刪除代理越南Catte遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" = 2008;

-- 刪除越南Catte遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" = 2008;

-- 刪除越南Catte遊戲 --
DELETE FROM "public"."game" WHERE "id" = 2008;