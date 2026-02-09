INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100288, '此接口用來取得server相關設定的參數', '/api/v1/system/getserversetting', 't', '後台使用', 'f', 0);

ALTER TABLE "public"."server_info" 
  ADD COLUMN "setting" jsonb NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN "public"."server_info"."setting" IS 'server相關設定(放置參數)';

UPDATE "public"."server_info"
  SET "setting" = "setting" || '{"common_report_time_range":31,"common_report_time_before_days":90,"common_report_time_minute_increment":5,"winlose_report_time_before_days":7,"earning_report_time_minute_increment":15}'
  WHERE "code" IN ('local', 'dev', 'qa', 'ete');

UPDATE "public"."server_info"
  SET "setting" = "setting" || '{"winlose_report_time_range":1}'
  WHERE "code" = 'ete';

UPDATE "public"."server_info"
  SET "setting" = "setting" || '{"winlose_report_time_range":12}'
  WHERE "code" IN ('local', 'dev', 'qa');