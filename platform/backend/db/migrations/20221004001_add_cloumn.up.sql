ALTER TABLE "public"."admin_user" 
  ADD COLUMN "role" int2 NOT NULL DEFAULT 0,
  ADD COLUMN "info" varchar(255) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."admin_user"."role" IS '角色';

COMMENT ON COLUMN "public"."admin_user"."info" IS '備註';
