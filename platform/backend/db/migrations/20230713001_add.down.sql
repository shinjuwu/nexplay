-- 刪除火箭遊戲紀錄表 --
DROP TABLE "public"."user_play_log_rocket";

-- 刪除代理火箭遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" = 1008
);

-- 刪除代理火箭遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" = 1008;

-- 刪除火箭遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" = 1008;

-- 刪除火箭遊戲 --
DELETE FROM "public"."game" WHERE "id" = 1008;