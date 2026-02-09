-- 刪除3公玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_sangong";

-- 刪除代理3公遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" = 2002
);

-- 刪除代理3公遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" = 2002;

-- 刪除3公遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" = 2002;

-- 刪除3公遊戲 --
DELETE FROM "public"."game" WHERE "id" = 2002;
