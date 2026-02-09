DROP VIEW "public"."view_game_management_game";
CREATE VIEW "public"."view_game_management_game" AS
  SELECT "ag"."agent_id",
    "ag"."game_id",
    "ag"."state",
    "a"."code" AS "agent_code",
    "a"."level_code" AS "agent_level_code",
    "g"."code" AS "game_code",
    "g"."state" AS "game_state"
  FROM "public"."agent_game" AS "ag"
  INNER JOIN "public"."agent" AS "a" ON "ag"."agent_id" = "a"."id"
  INNER JOIN "public"."game" AS "g" ON "ag"."game_id" = "g"."id"
  ORDER BY "ag"."game_id" ASC, "a"."level_code" ASC;
