UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100313]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" IN (2,3);

UPDATE "public"."permission_list"
SET "is_required" = true
WHERE "feature_code" = 100313;