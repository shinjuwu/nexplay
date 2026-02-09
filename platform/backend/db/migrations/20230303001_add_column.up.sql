ALTER TABLE "public"."game_users"
  ADD COLUMN "last_login_time" timestamptz(6) NOT NULL DEFAULT'1970-01-01 00:00:00+00'::timestamp with time zone,
  ADD COLUMN "last_logout_time" timestamptz(6) NOT NULL DEFAULT'1970-01-01 00:00:00+00'::timestamp with time zone;

COMMENT ON COLUMN "public"."game_users"."last_login_time" IS '最後登入時間';
COMMENT ON COLUMN "public"."game_users"."last_logout_time" IS '最後登出時間';
