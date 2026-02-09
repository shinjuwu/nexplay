DROP TABLE IF EXISTS "public"."schedule_to_backup";
CREATE TABLE IF NOT EXISTS "public"."schedule_to_backup" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "data_keeping_table" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "data_keeping_day" int4 NOT NULL DEFAULT '-1'::integer,
  "is_enabled" bool NOT NULL DEFAULT false,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  "disable_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  "last_exec_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  CONSTRAINT "schedule_to_backup_pkey" PRIMARY KEY ("id")
)
;

INSERT INTO "public"."schedule_to_backup" ("id", "data_keeping_table", "data_keeping_day", "is_enabled", "create_time", "update_time", "disable_time", "last_exec_time") VALUES ('1078ffb1-23c7-463f-a75e-da05d63cfa97', 'user_play_log', 7, 't', '2023-05-25 06:18:49.957161+00', '2023-05-25 06:18:49.957161+00', '1970-01-01 00:00:00+00', '1970-01-01 00:00:00+00');
