DROP TABLE "public"."miscellaneous";
CREATE TABLE "public"."storage" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "value" jsonb NOT NULL,
  "readonly" bool NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_storage_key" UNIQUE ("key"),
  CONSTRAINT "storage_pkey" PRIMARY KEY ("id")
)
;

INSERT INTO "public"."storage" ("id", "key", "value", "readonly", "create_time", "update_time") VALUES ('8dba60c6-0fbf-4e8e-943c-d4e38865be2a', 'GameServerInfo', '{"state": 1}', 'f', '2023-02-09 06:32:59.437471+00', '2023-02-09 06:32:59.437471+00');
