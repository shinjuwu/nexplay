-- 刪除搶庄牛牛玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_bullbull";

-- 刪除百人骰寶玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_hundredsicbo";

-- 刪除代理百人骰寶、搶庄牛牛遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (1005, 2003)
);

-- 刪除代理百人骰寶、搶庄牛牛遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (1005, 2003);

-- 刪除百人骰寶、搶庄牛牛遊戲房間 --
DELETE FROM  "public"."game_room" WHERE "game_id" IN (1005, 2003);

-- 刪除百人骰寶、搶庄牛牛遊戲 --
DELETE FROM  "public"."game" WHERE "id" IN (1005, 2003);
