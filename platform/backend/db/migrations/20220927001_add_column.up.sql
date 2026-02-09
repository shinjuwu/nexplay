ALTER TABLE "public"."game"
  ADD COLUMN "type" int NOT NULL DEFAULT 0;

UPDATE "public"."game"
  SET "type" = 1 WHERE "id" IN (1001, 1002, 1003);

UPDATE "public"."game_room"
  SET setting_info = '{"table_list":[100101,100102,100103,100104]}'
  WHERE "game_id" = 1001 AND "room_type" = 0;
UPDATE "public"."game_room"
  SET setting_info = '{"table_list":[100111,100112,100113,100114]}'
  WHERE "game_id" = 1001 AND "room_type" = 1;
UPDATE "public"."game_room"
  SET setting_info = '{"table_list":[100121,100122,100123,100124]}'
  WHERE "game_id" = 1001 AND "room_type" = 2;
UPDATE "public"."game_room"
  SET setting_info = '{"table_list":[100131,100132,100133,100134]}'
  WHERE "game_id" = 1001 AND "room_type" = 3;

COMMENT ON COLUMN "public"."game"."type" IS '遊戲類型';
ALTER TABLE "public"."agent_game"
  DROP COLUMN "room_list";

ALTER TABLE "public"."agent_game"
  ADD COLUMN "state" int2 NOT NULL DEFAULT 1;

COMMENT ON COLUMN "public"."agent_game"."state" IS '房間狀態';

ALTER TABLE "public"."agent_game_room"
  DROP COLUMN "agent_code",
  DROP COLUMN "agent_md5_key",
  DROP COLUMN "game_id",
  DROP COLUMN "game_code",
  DROP COLUMN "room_type";
