INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type")
VALUES
  (100289, '取得後台登入紀錄列表', '/api/v1/record/getbackendloginloglist', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100289]'::jsonb, false)
WHERE "agent_id" = -1;
