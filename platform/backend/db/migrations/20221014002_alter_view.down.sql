CREATE VIEW "public"."view_game_setting" AS
  SELECT "id", "code", "state"
  FROM "public"."game"
  WHERE "id" > 0;

ALTER VIEW "public"."view_agent_game" RENAME TO "view_game_management_game";

ALTER VIEW "public"."view_agent_game_room" RENAME TO "view_game_management_game_room";

ALTER VIEW "public"."view_admin_user" RENAME TO "view_login_admin_user";
