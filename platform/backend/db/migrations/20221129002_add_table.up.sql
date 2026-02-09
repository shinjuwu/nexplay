DELETE FROM "public"."permission_list";

-- 更新權限列表
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
  (100205, '取得代理遊戲列表', '/api/v1/agent/getagentgamelist', true, '後台使用', true),
  (100206, '設置代理遊戲狀態', '/api/v1/agent/setagentgamestate', true, '後台使用', true),
  (100207, '取得代理遊戲房間列表', '/api/v1/agent/getagentgameroomlist', true, '後台使用', true),
  (100208, '設置代理遊戲房間狀態', '/api/v1/agent/setagentgameroomstate', true, '後台使用', true),
  (100209, '取得代理權限群組權限樣板', '/api/v1/agent/getagentpermissiontemplateinfo', true, '後台使用', true),
  (100210, '取得代理權限群組列表', '/api/v1/agent/getagentpermissionlist', true, '後台使用', true),
  (100211, '創建代理權限群組', '/api/v1/agent/createagentpermission', true, '後台使用', true),
  (100212, '修改代理權限群組', '/api/v1/agent/setagentpermission', true, '後台使用', true),
  (100213, '刪除代理權限群組', '/api/v1/agent/deleteagentpermission', true, '後台使用', true),
  (100214, '取得遊戲列表', '/api/v1/game/getgamelist', true, '後台使用', true),
  (100215, '修改遊戲狀態', '/api/v1/game/setgamestate', true, '後台使用', true),
  (100216, '此接口用來重新載入本地端資料', '/api/v1/global/roloadglobaldata', true, '後台使用', true),
  (100218, '取得全部代理商list(供下拉選單使用)', '/api/v1/global/getagentlist', true, '後台使用', false),
  (100219, '取得全部遊戲list(供下拉選單使用)', '/api/v1/global/getallgamelist', true, '後台使用', false),
  (100220, '取得上線及維護中的遊戲list(供下拉選單使用)', '/api/v1/global/getgamelist', true, '後台使用', false),
  (100221, '取得房間類型list(供下拉選單使用)', '/api/v1/global/getroomtypelist', true, '後台使用', false),
  (100222, '取得權限群組層級list(供下拉選單使用)', '/api/v1/global/getagentpermissionlist', true, '後台使用', false),
  (100223, '用戶登錄', '/api/v1/login/login', true, '後台使用', false),
  (100224, '取得驗證碼', '/api/v1/login/captcha', true, '後台使用', false),
  (100225, '取得目前跑馬燈設定列表', '/api/v1/manage/getmarqueelist', true, '後台使用', true),
  (100226, '指定取得某筆跑馬燈設定', '/api/v1/manage/getmarquee', true, '後台使用', true),
  (100227, '添加跑馬燈功能', '/api/v1/manage/createmarquee', true, '後台使用', true),
  (100228, '編輯跑馬燈功能', '/api/v1/manage/updatemarquee', true, '後台使用', true),
  (100229, '刪除跑馬燈功能', '/api/v1/manage/deletemarquee', true, '後台使用', true),
  (100230, '取得個人遊戲紀錄列表', '/api/v1/record/getuserplayloglist', true, '後台使用', true),
  (100231, '取得遊戲局記錄', '/api/v1/record/getplaylogcommon', true, '後台使用', true),
  (100232, '取得帳變資料列表', '/api/v1/record/getwalletledgerlist', true, '後台使用', true),
  (100233, 'ping', '/api/v1/user/ping', true, '後台使用', true),
  (100234, '取得目前已產生的有效 token list', '/api/v1/user/getalivetoken', true, '後台使用', true),
  (100235, '將用戶 token 列入黑名單(主動使登入token 失效)', '/api/v1/user/blacktoken', true, '後台使用', true),
  (100236, '創建後台帳號(只能創建自己的後台帳號)', '/api/v1/user/createadminuser', true, '後台使用', true),
  (100237, '依照查詢者角色權限列出自身權限下的子帳號列表', '/api/v1/user/getadminusers', true, '後台使用', true),
  (100238, '指定查詢某後台帳號狀態', '/api/v1/user/getadminuserinfo', true, '後台使用', true),
  (100239, '指定設定某後台帳號狀態', '/api/v1/user/updateadminuserinfo', true, '後台使用', true),
  (100240, '依照查詢者角色權限列出遊戲會員帳號清單', '/api/v1/user/getgameusers', true, '後台使用', true),
  (100241, '指定查詢某遊戲會員帳號信息', '/api/v1/user/getgameuserinfo', true, '後台使用', true),
  (100242, '指定修改某遊戲會員帳號信息', '/api/v1/user/updategameuserinfo', true, '後台使用', true),(100243, '此接口用來重新計算業績報表', '/api/v1/cal/calperformancereport', true, '後台使用', true),
  (100244, '此接口用來取得當前 job 的資訊清單', '/api/v1/servicestatus/getjobshedulerlist', true, '後台使用',true),
  (100245, '依照指定報表類型取得指定時間區段的資料', '/api/v1/cal/getperformancereport', true, '後台使用', true),
  (100246, '取得代理後台IP資訊列表', '/api/v1/agent/getagentipwhitelistlist', true, '後台使用', true),
  (100247, '取得代理後台IP資訊', '/api/v1/agent/getagentipwhitelist', true, '後台使用', true),
  (100248, '設置代理後台IP資訊', '/api/v1/agent/setagentipwhitelist', true, '後台使用', true),
  (100249, '取得代理權限群組', '/api/v1/agent/getagentpermission', true, '後台使用', true),
  (100250, '取得代理錢包餘額列表', '/api/v1/agent/getagentwalletlist', true, '後台使用', true),
  (100251, '設置代理錢包餘額', '/api/v1/agent/setagentwallet', true, '後台使用', true),
  (100252, '取得代理分數紀錄列表', '/api/v1/record/getagentwalletledgerlist', true, '後台使用', true),
  (100253, '取得玩家錢包餘額列表', '/api/v1/user/getgameuserwalletlist', true, '後台使用', true),
  (100254, '設置玩家錢包餘額', '/api/v1/user/setgameuserwallet', true, '後台使用', true),
  (100255, '更新玩家分數紀錄狀態', '/api/v1/record/confirmwalletledger', true, '後台使用', true),
  (100256, '此接口用來取得當天資訊總覽的資料', '/api/v1/manage/getstatdata', true, '後台使用', true),
  (100257, '此接口用來取得今日風險玩家清單(前100名)', '/api/v1/manage/getriskuserlist', true, '後台使用', true),
  (100258, '此接口用來取得今日遊戲輸贏排行榜', '/api/v1/manage/getgameleaderboards', true, '後台使用', true),
  (200101, '第三方遊戲API', '/channel/channelHandle', true, '對外串接', false),
  (200102, '第三方取遊戲記錄', '/record/getRecordHandle', true, '對外串接', false);


-- 可操作權限 account_type 開發者:1, 總代理:2, 子代理:3
UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100215,100225,100226,100227,100228,100229,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100255,100256,100257,100258]}'
WHERE
  "account_type" = 1;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100253,100254,100255,100256,100257,100258]}'
WHERE
  "account_type" = 2;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100252,100253,100254,100255,100256,100257,100258]}'
WHERE
  "account_type" = 3;

UPDATE
  "public"."agent_permission"
SET
  "permission" = jsonb_set(
    "permission",
    '{list}',
    "permission" -> 'list' || '[100216]' :: jsonb,
    false
  )
WHERE
  "agent_id" = -1 AND "account_type" = 1;