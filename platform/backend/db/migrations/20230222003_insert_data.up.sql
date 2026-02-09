INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type")
VALUES
  (100276, '修改個人資訊', '/api/v1/user/setpersonalinfo', 't', '後台使用', 't', 2),
  (100277, '修改個人密碼', '/api/v1/user/setpersonalpassword', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100276, 100277]'::jsonb, false);
