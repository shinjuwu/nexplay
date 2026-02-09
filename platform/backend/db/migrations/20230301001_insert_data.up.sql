INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES 
 (100287,'此接口用來取得玩家標示設定下拉選單資料list(供下拉選單使用','/api/v1/global/getagentcustomtagsettinglist','t','後台使用','t',0);
UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100287]'::jsonb, false)
WHERE "agent_id" = -1;
