DROP VIEW "public"."view_agent_game";
CREATE VIEW "public"."view_agent_game" AS
  SELECT "ag"."agent_id",
    "ag"."game_id",
    "ag"."state",
    "a"."name" AS "agent_name",
    "a"."code" AS "agent_code",
    "a"."level_code" AS "agent_level_code",
    "g"."code" AS "game_code",
    "g"."state" AS "game_state"
  FROM "public"."agent_game" AS "ag"
  INNER JOIN "public"."agent" AS "a" ON "ag"."agent_id" = "a"."id"
  INNER JOIN "public"."game" AS "g" ON "ag"."game_id" = "g"."id"
  ORDER BY "ag"."game_id" ASC, "a"."level_code" ASC;

DROP VIEW "public"."view_agent_game_room";
CREATE VIEW "public"."view_agent_game_room" AS
  SELECT "agr"."agent_id",
    "agr"."game_room_id",
    "agr"."state",
    "a"."name" AS "agent_name",
    "a"."code" AS "agent_code",
    "g"."id" AS "game_id",
    "g"."code" AS "game_code",
    "gr"."room_type"
  FROM "public"."agent_game_room" AS "agr"
  INNER JOIN "public"."agent" AS "a" ON "agr"."agent_id" = "a"."id"
  INNER JOIN "public"."game_room" AS "gr" ON "agr"."game_room_id" = "gr"."id"
  INNER JOIN "public"."game" AS "g" ON "gr"."game_id" = "g"."id";
