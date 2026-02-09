DROP VIEW "public"."view_game_setting";

ALTER VIEW "public"."view_game_management_game" RENAME TO "view_agent_game";

ALTER VIEW "public"."view_game_management_game_room" RENAME TO "view_agent_game_room";

ALTER VIEW "public"."view_login_admin_user" RENAME TO "view_admin_user";
