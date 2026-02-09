-- 刪除印度炸金花紀錄表 --
DROP TABLE "public"."user_play_log_teenpatti";

-- 刪除代理印度炸金花遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (2011)
);

-- 刪除代理印度炸金花遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (2011);

-- 刪除印度炸金花遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (2011);

-- 刪除印度炸金花遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (2011);

-- 刪除遊戲基礎設定支援遊戲 --
UPDATE "public"."storage"
  SET "value" = (
	  SELECT jsonb_agg("game_id")
	  FROM (
		  SELECT jsonb_array_elements("value") AS "game_id"
            FROM "public"."storage"
            WHERE "key" = 'GameSettingSupportInfo'
	  )
	  WHERE "game_id"::int NOT IN (2011)
  )
  WHERE "key"= 'GameSettingSupportInfo';
