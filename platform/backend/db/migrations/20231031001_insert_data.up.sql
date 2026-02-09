UPDATE "public"."agent_permission" AS "ag1"
SET "permission" = jsonb_set("permission", '{list}', "ag3"."list", false)
FROM (
    SELECT "ag2"."id", jsonb_agg("ag2ple") AS "list"
    FROM "public"."agent_permission" AS "ag2", jsonb_array_elements("ag2"."permission"->'list') AS "ag2ple"
    WHERE "ag2ple"::int NOT IN (100313)
    GROUP BY "ag2"."id"
) AS "ag3"
WHERE "ag1"."id" = "ag3"."id" AND "ag1"."account_type" IN (2,3);

UPDATE "public"."permission_list"
SET "is_required" = false
WHERE "feature_code" = 100313;