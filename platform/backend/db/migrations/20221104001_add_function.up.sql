CREATE FUNCTION "public"."udf_create_agent_games"("_agent_id" int4)
  RETURNS json AS $$
DECLARE
  "ret_agent_games" json;
BEGIN
  INSERT INTO "public"."agent_game" ("agent_id", "game_id")
    SELECT "_agent_id" AS "agent_id", "id" AS "game_id"
      FROM "public"."game"
      WHERE "id" <> 0;
	  
  SELECT json_agg("agent_games") INTO "ret_agent_games"
    FROM (
	  SELECT "agent_id", "game_id", "state"
		  FROM "public"."agent_game"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_games";
  
  RETURN "ret_agent_games";
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION "public"."udf_create_agent_game_rooms"("_agent_id" int4)
  RETURNS json AS $$
DECLARE
  "ret_agent_game_rooms" json;
BEGIN
  INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
    SELECT "_agent_id" AS "agent_id", "id" AS "game_id"
      FROM "public"."game_room";
	  
  SELECT json_agg("agent_game_rooms") INTO "ret_agent_game_rooms"
    FROM (
	  SELECT "agent_id", "game_room_id", "state"
		  FROM "public"."agent_game_room"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_game_rooms";
  
  RETURN "ret_agent_game_rooms";
END;
$$ LANGUAGE plpgsql;
