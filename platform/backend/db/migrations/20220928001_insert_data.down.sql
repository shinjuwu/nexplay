DELETE FROM "public"."game" WHERE "id" IN (1004, 2001);

DELETE FROM "public"."game_room" WHERE "game_id" IN (1002, 1003, 1004, 2001);
