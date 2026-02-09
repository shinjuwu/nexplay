INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100303, '取得維護頁設定(只有開發商(營用商)可以使用)', '/api/v1/manage/getmaintainpagesetting', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES ( 100304, '設定維護頁設定(只有開發商(營用商)可以使用)', '/api/v1/manage/setmaintainpagesetting', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100303, 100304]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;