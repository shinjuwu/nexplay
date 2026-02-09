INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100321, '此接口用來取得代理父帳號權限list(供下拉選單使用)', '/api/v1/global/getagentadminuserpermissionlist', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100321]'::jsonb, false)
WHERE "permission"->'list' @> '100266';