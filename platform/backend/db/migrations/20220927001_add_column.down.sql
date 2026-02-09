ALTER TABLE "public"."game"
  DROP COLUMN "type";

ALTER TABLE "public"."agent_game"
  DROP COLUMN "state";

ALTER TABLE "public"."agent_game"
  ADD COLUMN "room_list" jsonb NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN "public"."agent_game"."room_list" IS '開放房間列表';

ALTER TABLE "public"."agent_game_room"
  ADD COLUMN "agent_code" varchar(8) NOT NULL DEFAULT '',
  ADD COLUMN "agent_md5_key" varchar(32) NOT NULL DEFAULT '',
  ADD COLUMN "game_id" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "game_code" varchar(40) NOT NULL DEFAULT '',
  ADD COLUMN "room_type" int4 NOT NULL DEFAULT 0;
  
COMMENT ON COLUMN "public"."agent_game_room"."agent_code" IS '代理識別編碼(mapping agent code)';
COMMENT ON COLUMN "public"."agent_game_room"."agent_md5_key" IS '代理h5game專用md5密鑰(mapping agent md5_key)';
COMMENT ON COLUMN "public"."agent_game_room"."game_id" IS '遊戲編號(mapping game id)';
COMMENT ON COLUMN "public"."agent_game_room"."game_code" IS '遊戲識別編碼(mapping game code)';
COMMENT ON COLUMN "public"."agent_game_room"."room_type" IS '房間類型';
