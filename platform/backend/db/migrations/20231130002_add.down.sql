-- 刪除土耳其麻將紀錄表 --
DROP TABLE "public"."user_play_log_okey";

-- 刪除代理土耳其麻將遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (2010)
);

-- 刪除代理土耳其麻將遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (2010);

-- 刪除土耳其麻將遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (2010);

-- 刪除土耳其麻將遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (2010);