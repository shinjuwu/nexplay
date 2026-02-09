ALTER TABLE "public"."agent_permission" 
  DROP CONSTRAINT "agent_permission_pkey",
  ADD COLUMN "name" varchar(20) NOT NULL DEFAULT ''::character varying,
  ADD COLUMN "info" varchar(255) NOT NULL DEFAULT ''::character varying,
  ADD COLUMN "account_type" int2 NOT NULL DEFAULT 0,
  ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("agent_id", "level", "account_type");

COMMENT ON COLUMN "public"."agent_permission"."name" IS '角色名稱';

COMMENT ON COLUMN "public"."agent_permission"."info" IS '備註';

COMMENT ON COLUMN "public"."agent_permission"."account_type" IS '帳號角色';