CREATE VIEW "public"."view_game_management_game" AS
  SELECT "ag"."agent_id",
    "ag"."game_id",
    "ag"."state",
    "a"."code" AS "agent_code",
    "a"."top_agent_id",
    "a"."is_top_agent",
    "g"."code" AS "game_code"
  FROM "public"."agent_game" AS "ag"
  INNER JOIN "public"."agent" AS "a" ON "ag"."agent_id" = "a"."id"
  INNER JOIN "public"."game" AS "g" ON "ag"."game_id" = "g"."id"
  ORDER BY
    "ag"."game_id", 
    CASE
      WHEN "a"."is_top_agent" THEN "ag"."agent_id"
      ELSE "a"."top_agent_id"
    END,
    "ag"."agent_id" ASC;

CREATE VIEW "public"."view_game_management_game_room" AS
  SELECT "agr"."agent_id",
    "agr"."game_room_id",
    "agr"."state",
    "a"."code" AS "agent_code",
    "g"."id" AS "game_id",
    "g"."code" AS "game_code",
    "gr"."room_type"
  FROM "public"."agent_game_room" AS "agr"
  INNER JOIN "public"."agent" AS "a" ON "agr"."agent_id" = "a"."id"
  INNER JOIN "public"."game_room" AS "gr" ON "agr"."game_room_id" = "gr"."id"
  INNER JOIN "public"."game" AS "g" ON "gr"."game_id" = "g"."id";