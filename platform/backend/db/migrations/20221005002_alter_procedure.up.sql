DROP PROCEDURE "public"."procedure_update_agent_game_state_by_top_agent";
CREATE PROCEDURE "public"."procedure_update_agent_game_state_by_top_agent"(target_top_agent_id int, target_agent_level_code varchar, target_game_id int, agent_game_state int)
  LANGUAGE plpgsql AS
$$
BEGIN
  UPDATE "public"."agent_game" AS "ag"
    SET "state" = agent_game_state
    FROM "public"."agent" AS "a"
    WHERE "ag"."agent_id" = "a"."id"
      AND "a"."level_code" LIKE target_agent_level_code || '%'
      AND "ag"."game_id" = target_game_id;

  CALL "public"."procedure_update_agent_game_state_by_agent"(target_top_agent_id, target_game_id, agent_game_state);
END
$$;

DROP PROCEDURE "public"."procedure_update_agent_game_room_state_by_top_agent";
CREATE PROCEDURE "public"."procedure_update_agent_game_room_state_by_top_agent"(target_top_agent_id int, target_agent_level_code varchar, target_game_room_id int, agent_game_room_state int)
  LANGUAGE plpgsql AS
$$
BEGIN
  UPDATE "public"."agent_game_room" AS "agr"
    SET "state" = agent_game_room_state
    FROM "public"."agent" AS "a"
    WHERE "agr"."agent_id" = "a"."id"
      AND "a"."level_code" LIKE target_agent_level_code || '%'
      AND "agr"."game_room_id" = target_game_room_id;

  CALL "public"."procedure_update_agent_game_room_state_by_agent"(target_top_agent_id, target_game_room_id, agent_game_room_state);
END
$$;

