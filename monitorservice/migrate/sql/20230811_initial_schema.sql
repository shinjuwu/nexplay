-- +migrate Up
CREATE TABLE "public"."db_conn_info" (
  "id" varchar(5) COLLATE "pg_catalog"."default" NOT NULL,
  "addresses" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT false,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  "disable_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone, 
  CONSTRAINT "db_conn_info_pkey" PRIMARY KEY ("id")
)
;

CREATE TABLE "public"."service_info" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar(5) COLLATE "pg_catalog"."default" NOT NULL,
  "sub_name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "api_urls" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT true,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  "disable_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  CONSTRAINT "service_info_pkey" PRIMARY KEY ("id")
)
;

CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "top_code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "nickname" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "user_metadata" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "is_enabled" bool NOT NULL DEFAULT false,
  "last_login_time" timestamptz(6) NOT NULL DEFAULT now(),
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  "disable_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  "is_admin" bool NOT NULL DEFAULT false,
  "permission" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "users_pkey" PRIMARY KEY ("id", "top_code", "username"),
  CONSTRAINT "users_name" UNIQUE ("username")
)
;

INSERT INTO "public"."users" ("id", "top_code", "username", "password", "nickname", "user_metadata", "is_enabled", "last_login_time", "info", "create_time", "update_time", "disable_time", "is_admin", "permission") VALUES ('5b393680-c0cc-4580-b158-c1ec24438570', '0d608095', 'admin', 'Rayg38qIwQ7E0dVlIkDGGTo6D8c5YU1Jumwlt3MVlc5HEQ', 'admmmmm', '{}', 't', '2023-08-29 06:17:46.558257+00', '', '2023-08-14 10:36:44.468554+00', '2023-08-14 10:36:44.468554+00', '1970-01-01 00:00:00+00', 't', '["dev", "qa", "ete", "pro"]');

-- +migrate Down
DROP TABLE IF EXISTS
    db_conn_info, service_info, users;