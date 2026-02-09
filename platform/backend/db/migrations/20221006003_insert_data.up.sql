DELETE FROM "public"."permission_list";

INSERT INTO "public"."permission_list"("feature_code", "name", "api_path", "is_enabled", "remark", "is_required")
VALUES
  (100001, '檢查伺服器是否活著', '/api/v1/example/health', true, 'test', false),

  (100100, 'client取得驗證token', '/api/v1/intercom/getlogintoken', true, '遊戲SERVER串接使用', false),
  (100101, '驗證token並登入遊戲', '/api/v1/intercom/logingame', true, '遊戲SERVER串接使用', false),
  (100102, '遊戲用戶登出', '/api/v1/intercom/logoutgame', true, '遊戲SERVER串接使用', false),
  (100103, '創建每局遊戲紀錄', '/api/v1/intercom/creategamerecord', true, '遊戲SERVER串接使用', false),
  (100104, '此接口供遊戲伺服器取得跑馬燈設定列表', '/api/v1/intercom/getmarqueesetting', true, '遊戲SERVER串接使用', false),

  (100200, '創建代理帳號', '/api/v1/agent/createagent', true, '後台使用', true),
  (100201, '取得代理底下所有代理資料', '/api/v1/agent/getagentlist', true, '後台使用', true),
  (100202, '秘鑰資訊顯示', '/api/v1/agent/getagentsecretkey', true, '後台使用', true),
  (100203, '取得指定代理補分相關資料設定', '/api/v1/agent/getagentcoinsupplyinfo', true, '後台使用', true),
  (100204, '修改指定代理補分相關資料設定', '/api/v1/agent/setagentcoinsupplyinfo', true, '後台使用', true),
  (100205, '取得遊戲管理Config', '/api/v1/agent/getgamemanagementconfig', true, '後台使用', true),
  (100206, '取得遊戲管理代理遊戲列表', '/api/v1/agent/getgamemanagementgamelist', true, '後台使用', true),
  (100207, '設置代理遊戲狀態', '/api/v1/agent/setagentgamestate', true, '後台使用', true),
  (100208, '取得遊戲管理代理遊戲房間列表', '/api/v1/agent/getgamemanagementgameroomlist', true, '後台使用', true),
  (100209, '設置代理遊戲房間狀態', '/api/v1/agent/setagentgameroomstate', true, '後台使用', true),
  (100210, '取得角色管理權限樣板', '/api/v1/agent/getrolemanagementtemplateinfo', true, '後台使用', true),
  (100211, '取得代理角色列表', '/api/v1/agent/getrolemanagementlist', true, '後台使用', true),
  (100212, '創建代理角色', '/api/v1/agent/createagentpermission', true, '後台使用', true),
  (100213, '修改代理角色', '/api/v1/agent/setagentpermission', true, '後台使用', true),
  (100214, '刪除代理角色', '/api/v1/agent/deleteagentpermission', true, '後台使用', true),
  (100215, '取得遊戲設置Config', '/api/v1/game/getgamesettingconfig', true, '後台使用', true),
  (100216, '取得遊戲設置列表', '/api/v1/game/getgamesettinglist', true, '後台使用', true),
  (100217, '修改遊戲狀態', '/api/v1/game/setgamestate', true, '後台使用', true),
  (100218, '此接口用來重新載入本地端資料', '/api/v1/global/roloadglobaldata', true, '後台使用', true),
  (100219, '此接口用來檢查並重設 game_room setting', '/api/v1/global/checkgameroomsetting', true, '後台使用', true),
  (100220, '用戶登錄', '/api/v1/login/login', true, '後台使用', false),
  (100221, '取得驗證碼', '/api/v1/login/captcha', true, '後台使用', false),
  (100222, '取得目前跑馬燈設定列表', '/api/v1/manage/getmarqueelist', true, '後台使用', true),
  (100223, '指定取得某筆跑馬燈設定', '/api/v1/manage/getmarquee', true, '後台使用', true),
  (100224, '添加跑馬燈功能', '/api/v1/manage/createmarquee', true, '後台使用', true),
  (100225, '編輯跑馬燈功能', '/api/v1/manage/updatemarquee', true, '後台使用', true),
  (100226, '刪除跑馬燈功能', '/api/v1/manage/deletemarquee', true, '後台使用', true),
  (100227, '取得輸贏報表Config', '/api/v1/record/getwinlosereportconfig', true, '後台使用', true),
  (100228, '取得輸贏報表', '/api/v1/record/getwinlosereport', true, '後台使用', true),
  (100229, '取得遊戲日誌解析', '/api/v1/record/getgamelog', true, '後台使用', true),
  (100230, '取得上下分紀錄Config', '/api/v1/record/getwalletledgerconfig', true, '後台使用', true),
  (100231, '取得上下分紀錄', '/api/v1/record/getwalletledger', true, '後台使用', true),
  (100232, 'ping', '/api/v1/user/ping', true, '後台使用', true),
  (100233, '創建後台帳號(只能創建自己的後台帳號)', '/api/v1/user/createadminuser', true, '後台使用', true),
  (100234, '依照查詢者角色權限列出自身權限下的子帳號列表', '/api/v1/user/getadminusers', true, '後台使用', true),
  (100235, '指定查詢某後台帳號狀態', '/api/v1/user/getadminuserinfo', true, '後台使用', true),
  (100236, '指定設定某後台帳號狀態', '/api/v1/user/updateadminuserinfo', true, '後台使用', true),
  (100237, '依照查詢者角色權限列出遊戲會員帳號清單', '/api/v1/user/getgameusers', true, '後台使用', true),
  (100238, '指定查詢某遊戲會員帳號信息', '/api/v1/user/getgameuserinfo', true, '後台使用', true),
  (100239, '指定修改某遊戲會員帳號信息', '/api/v1/user/updategameuserinfo', true, '後台使用', true),

  (200101, '第三方遊戲API', '/channel/channelHandle', true, '對外串接', false),
  (200102, '第三方取遊戲記錄', '/record/getRecordHandle', true, '對外串接', false);

DELETE FROM "public"."agent_permission";

INSERT INTO "public"."agent_permission"("agent_id", "name", "account_type", "permission")
VALUES
  (-1, '管理後台權限', 1, '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100215,100216,100217,100218,100219,100222,100223,100224,100225,100226,100227,100228,100229,100230,100231,100232,100233,100234,100235,100236,100237,100238,100239]}'),
  (-1, '總代後台權限', 2, '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100222,100223,100227,100228,100229,100230,100231,100232,100233,100234,100235,100236,100237,100238,100239]}'),
  (-1, '子代後台帳號', 3, '{"list":[100201,100202,100203,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100222,100223,100227,100228,100229,100230,100231,100232,100233,100234,100235,100236,100237,100238,100239]}');

INSERT INTO "public"."agent_permission"("agent_id", "name", "account_type", "permission")
SELECT "au"."agent_id",
  '后台权限群组' AS "name",
  "au"."account_type",
  "ag"."permission"
  FROM "public"."admin_user" AS "au"
  INNER JOIN "public"."agent_permission" AS "ag" ON "ag"."account_type" = "au"."account_type"
  WHERE "au"."agent_id" IN (
    SELECT "agent_id" FROM "public"."admin_user"
    WHERE "is_added" = true
    GROUP BY "agent_id"
  )
  GROUP BY "au"."agent_id", "au"."account_type", "ag"."permission";

INSERT INTO "public"."agent_permission"("agent_id", "name", "account_type", "permission")
SELECT "a"."top_agent_id" AS "agent_id",
  '代理后台权限群组' AS "name",
  "au"."account_type",
  "ag"."permission"
  FROM "public"."admin_user" AS "au"
  INNER JOIN "public"."agent_permission" AS "ag" ON "ag"."account_type" = "au"."account_type"
  INNER JOIN "public"."agent" AS "a" ON "a"."id" = "au"."agent_id"
  WHERE "au"."agent_id" IN (
    SELECT "agent_id" FROM "public"."admin_user"
    WHERE "is_added" = false AND "agent_id" <> 1
    GROUP BY "agent_id"
  )
  GROUP BY "a"."top_agent_id", "au"."account_type", "ag"."permission";

UPDATE "public"."admin_user" AS "au"
SET "role" = "ap"."id"
FROM "public"."agent_permission" AS "ap"
WHERE "au"."agent_id" = "ap"."agent_id" AND "au"."account_type" = "ap"."account_type" AND "au"."is_added" = true;

UPDATE "public"."admin_user" AS "au"
SET "role" = "ap"."id"
FROM "public"."admin_user" AS "auu"
INNER JOIN "public"."agent" AS "a" ON "a"."id" = "auu"."agent_id"
INNER JOIN "public"."agent_permission" AS "ap" ON "a"."top_agent_id" = "ap"."agent_id"
WHERE "au"."agent_id" = "auu"."agent_id" AND "au"."username" = "auu"."username" AND "auu"."is_added" = false AND "auu"."account_type" = "ap"."account_type";

