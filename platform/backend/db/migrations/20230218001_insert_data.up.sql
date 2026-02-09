INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100275, '指定時間區段下層代理總和的資料', '/api/v1/cal/getperformancereportlist', 't', '後台使用', 't', 0 );

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100275]'::jsonb, false)
WHERE "agent_id" = -1;
