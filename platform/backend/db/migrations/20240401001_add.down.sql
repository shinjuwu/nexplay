-- 刪除輪盤紀錄表 --
DROP TABLE "public"."user_play_log_roulette";

-- 刪除代理輪盤遊戲房間設定 --
DELETE FROM "public"."agent_game_room" WHERE "game_room_id" IN (
  SELECT "id" FROM "public"."game_room" WHERE "game_id" IN (1010)
);

-- 刪除代理輪盤遊戲設定 --
DELETE FROM "public"."agent_game" WHERE "game_id" IN (1010);

-- 刪除輪盤遊戲房間 --
DELETE FROM "public"."game_room" WHERE "game_id" IN (1010);

-- 刪除輪盤遊戲 --
DELETE FROM "public"."game" WHERE "id" IN (1010);

-- 刪除遊戲基礎設定支援遊戲 --
UPDATE "public"."storage"
  SET "value" = (
	  SELECT jsonb_agg("game_id")
	  FROM (
		  SELECT jsonb_array_elements("value") AS "game_id"
            FROM "public"."storage"
            WHERE "key" = 'GameSettingSupportInfo'
	  )
	  WHERE "game_id"::int NOT IN (1010)
  )
  WHERE "key"= 'GameSettingSupportInfo';