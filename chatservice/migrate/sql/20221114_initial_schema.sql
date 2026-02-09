-- +migrate Up
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "platform" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "username" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "unread" int4 NOT NULL DEFAULT 0,
  "metadata" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "is_disabled" bool NOT NULL DEFAULT false,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_username_platform" UNIQUE ("username", "platform"),
  CONSTRAINT "users_pkey" PRIMARY KEY ("id", "platform", "username")
)
;

CREATE TABLE "public"."notification" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "platform" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "content" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "message_pkey" PRIMARY KEY ("id")
)
;

-- +migrate Down
DROP TABLE IF EXISTS
    users, notification;