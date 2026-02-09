ALTER TABLE "public"."game_users" 
  ADD COLUMN "info" varchar(255) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."game_users"."info" IS '備註';