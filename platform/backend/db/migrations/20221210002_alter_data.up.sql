UPDATE "public"."game_users" AS "gu"
  SET "level_code" = "a"."level_code"
  FROM "public"."agent" AS "a"
  WHERE "a"."id" = "gu"."agent_id";