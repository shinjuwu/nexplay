DELETE FROM "public"."permission_list" WHERE "feature_code" >= 100260 AND "feature_code" <= 100262;

UPDATE "public"."agent_permission" AS "ag1"
SET "permission" = jsonb_set("permission", '{list}', "ag3"."list", false)
FROM (
    SELECT "ag2"."id", jsonb_agg("ag2ple") AS "list"
    FROM "public"."agent_permission" AS "ag2", jsonb_array_elements("ag2"."permission"->'list') AS "ag2ple"
    WHERE "ag2ple"::int NOT IN (100260, 100261, 100262)
    GROUP BY "ag2"."id"
) AS "ag3"
WHERE "ag1"."id" = "ag3"."id";

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES (100266,	'取得後台操作紀錄列表',	'/api/v1/record/getbackendactionloglist', 't', '後台使用', 't');

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100266]'::jsonb, false)
WHERE "agent_id" = -1;
