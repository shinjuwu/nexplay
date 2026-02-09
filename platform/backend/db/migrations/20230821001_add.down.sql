-- 刪除水果777遊戲紀錄表 --
DROP TABLE "public"."user_play_log_fruit777slot";

-- 刪除代理水果777遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" = 4001
);

-- 刪除代理水果777遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" = 4001;

-- 刪除水果777遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" = 4001;

-- 刪除水果777遊戲 --
DELETE FROM "public"."game" WHERE "id" = 4001;