UPDATE "public"."agent" SET "md5_key" = 'hongkong3345678', "aes_key" = '1234567890123456', "level_code" = '0001' WHERE "id" = 1;
UPDATE "public"."agent" SET "md5_key" = 'hongkong3345678', "aes_key" = '1234567890123456', "level_code" = '0002' WHERE "id" = 2;

-- 增加權限
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100217, '創建代理帳號', '/api/v1/agent/createagent', 't', '後台使用', '2022-09-19 02:01:12.930327+00', '2022-09-19 02:01:12.930327+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100218, '取得代理底下所有代理資料', '/api/v1/agent/getagentlist', 't', '後台使用', '2022-09-19 02:01:50.641097+00', '2022-09-19 02:01:50.641097+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100219, '秘鑰資訊顯示', '/api/v1/agent/getagentsecretkey', 't', '後台使用', '2022-09-19 02:02:30.499902+00', '2022-09-19 02:02:30.499902+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100220, '取得指定代理補分相關資料設定', '/api/v1/agent/getagentcoinsupplyinfo', 't', '後台使用', '2022-09-19 02:02:59.835716+00', '2022-09-19 02:02:59.835716+00', 't');


-- set all
UPDATE admin_user SET permission='{"list":[100200,100201,100202,100203,100204,100205,100208,100209,100210,100211,100212,100213,100214,100215,100216,100217,100218,100219,100220,200101,200102]}';
