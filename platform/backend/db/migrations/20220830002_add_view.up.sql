CREATE VIEW "public"."view_game_setting" AS
  SELECT "id", "code", "state"
  FROM "public"."game"
  WHERE "id" > 0
