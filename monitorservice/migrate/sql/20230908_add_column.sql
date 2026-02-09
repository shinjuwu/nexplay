-- +migrate Up
ALTER TABLE "public"."user_play_log" 
  ADD COLUMN "room_name" varchar(40) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log"."room_name" IS '房間名稱';

-- +migrate Down
ALTER TABLE "public"."user_play_log" 
  DROP COLUMN "room_name";