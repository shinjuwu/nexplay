-- 刪除彈珠檯遊戲紀錄表 --
DROP TABLE "public"."user_play_log_plinko";

-- 刪除十三水遊戲紀錄表 --
DROP TABLE "public"."user_play_log_chinesepoker";

-- 刪除代理十三水、彈珠檯遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (2009, 3003)
);

-- 刪除代理十三水、彈珠檯遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (2009, 3003);

-- 刪除十三水、彈珠檯遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (2009, 3003);

-- 刪除十三水、彈珠檯遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (2009, 3003);