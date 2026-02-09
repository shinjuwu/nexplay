ALTER TABLE "public"."admin_user" 
  ADD COLUMN "is_added" bool NOT NULL DEFAULT true;

COMMENT ON COLUMN "public"."admin_user"."is_added" IS '是否為分身帳號';

-- update data
UPDATE "public"."admin_user" SET "is_added"=false WHERE "username"='admin';

-- set all
UPDATE "public"."agent_permission" SET "permission"='{"list":[100200,100201,100202,100203,100204,100205,100208,100209,100210,100211,100212,100213,100214,100215,100216,100217,100218,100219,100220,200101,200102]}' WHERE "level" = 0;
