DELETE FROM "public"."permission_list" WHERE "feature_code" = 100217;

DELETE FROM "public"."game_room";

INSERT INTO "public"."game_room" ("id", "name", "state", "game_id", "room_type")
VALUES
  (10010, '百家樂新手房',   1, 1001, 0),
  (10011, '百家樂普通房',   1, 1001, 1),
  (10012, '百家樂高手房',   1, 1001, 2),
  (10013, '百家樂大師房',   1, 1001, 3),
  (10020, '番攤新手房',     1, 1002, 0),
  (10021, '番攤普通房',     1, 1002, 1),
  (10022, '番攤高手房',     1, 1002, 2),
  (10023, '番攤大師房',     1, 1002, 3),
  (10030, '色碟新手房',     1, 1003, 0),
  (10031, '色碟普通房',     1, 1003, 1),
  (10032, '色碟高手房',     1, 1003, 2),
  (10033, '色碟大師房',     1, 1003, 3),
  (10040, '魚蝦蟹新手房',   1, 1004, 0),
  (10041, '魚蝦蟹普通房',   1, 1004, 1),
  (10042, '魚蝦蟹高手房',   1, 1004, 2),
  (10043, '魚蝦蟹大師房',   1, 1004, 3),
  (20010, '21點新手房',     1, 2001, 0),
  (20011, '21點普通房',     1, 2001, 1),
  (20012, '21點高手房',     1, 2001, 2),
  (20013, '21點大師房',     1, 2001, 3);

DELETE FROM "public"."agent_game";

INSERT INTO "public"."agent_game" ("agent_id", "game_id")
SELECT "a"."id" AS "agent_id",
  "g"."id" AS "game_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game" AS "g"
  WHERE "g"."id" <> 0;

DELETE FROM "public"."agent_game_room";

INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id")
SELECT "a"."id" AS "agent_id",
  "gr"."id" AS "game_room_id"
  FROM "public"."agent" AS "a"
  CROSS JOIN "public"."game_room" AS "gr";
