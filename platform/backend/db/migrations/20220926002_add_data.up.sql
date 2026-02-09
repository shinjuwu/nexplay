-- 新增權限時，需要將其他相關的表也一併做更新，
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100221, '修改指定代理補分相關資料設定', '/api/v1/agent/setagentcoinsupplyinfo', 't', '後台使用', '2022-09-26 06:02:42.486643+00', '2022-09-26 06:02:42.486643+00', 't');

UPDATE "public"."agent_permission" SET "permission"='{"list":[100200,100201,100202,100203,100204,100205,100208,100209,100210,100211,100212,100213,100214,100215,100216,100217,100218,100219,100220,100221,200101,200102]}' WHERE "level" = 0;

UPDATE "public"."admin_user" SET "permission"='{"list":[100200,100201,100202,100203,100204,100205,100208,100209,100210,100211,100212,100213,100214,100215,100216,100217,100218,100219,100220,100221,200101,200102]}';

