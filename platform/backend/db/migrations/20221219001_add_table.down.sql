DROP TABLE "public"."miscellaneous";

DELETE FROM "public"."permission_list" WHERE feature_code IN (100267, 100268, 100269);

UPDATE "public"."agent_permission" AS "ag1"
SET "permission" = jsonb_set("permission", '{list}', "ag3"."list", false)
FROM (
    SELECT "ag2"."id", jsonb_agg("ag2ple") AS "list"
    FROM "public"."agent_permission" AS "ag2", jsonb_array_elements("ag2"."permission"->'list') AS "ag2ple"
    WHERE "ag2"."agent_id" = -1 AND "ag2"."account_type" = 1 AND "ag2ple"::int NOT IN (100267, 100268, 100269)
    GROUP BY "ag2"."id"
) AS "ag3"
WHERE "ag1"."id" = "ag3"."id";
