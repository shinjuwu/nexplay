ALTER TABLE "public"."exchange_data" RENAME COLUMN "to_cny" TO "to_coin";

DELETE FROM "public"."exchange_data";

INSERT INTO "public"."exchange_data" ("currency", "to_coin")
VALUES
  ('CNY', 1),
  ('VND', 3200),
  ('THB', 5),
  ('MYR', 0.5),
  ('PHP', 8),
  ('INR', 10),
  ('BRL', 0.5);

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100294, '取得匯率設定', '/api/v1/system/getexchangedatalist', 't', '後台使用', 't', 0 );
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100295, '設定匯率設定', '/api/v1/system/setexchangedatalist', 't', '後台使用', 't', 2 );

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100294, 100295]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;