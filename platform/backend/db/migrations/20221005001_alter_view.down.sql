DROP VIEW "public"."view_game_management_game";
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
