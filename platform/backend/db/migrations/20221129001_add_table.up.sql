CREATE TABLE "public"."user_login_log" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "game_user_id" int4 NOT NULL,
  "token" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "coin" numeric(20,4) NOT NULL DEFAULT 0,
  "user_info" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "is_new" bool NOT NULL,
  "login_time" timestamptz(6) NOT NULL DEFAULT now(),
  "logout_time" timestamptz(6) NOT NULL DEFAULT '1970-01-01 00:00:00+00'::timestamp with time zone,
  CONSTRAINT "user_login_log_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."user_login_log"."token" IS '本次登錄使用token';
COMMENT ON COLUMN "public"."user_login_log"."coin" IS '本次帶入遊戲幣';
COMMENT ON COLUMN "public"."user_login_log"."user_info" IS '本次登錄用戶資料';
COMMENT ON COLUMN "public"."user_login_log"."is_new" IS '是否為新帳號';
COMMENT ON TABLE "public"."user_login_log" IS '用戶登入紀錄';

CREATE TABLE "public"."rt_data_stat_day" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "active_player" int4 NOT NULL DEFAULT 0,
  "number_bettors" int4 NOT NULL DEFAULT 0,
  "number_registrants" int4 NOT NULL DEFAULT 0,
  "odd_number" int4 NOT NULL DEFAULT 0,
  "total_betting" numeric(20,4) NOT NULL DEFAULT 0,
  "game_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_win_score" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_lose_score" numeric(20,4) NOT NULL DEFAULT 0,
  "raw_data" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "rt_data_stat_day_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code")
)
;
COMMENT ON COLUMN "public"."rt_data_stat_day"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_data_stat_day"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rt_data_stat_day"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rt_data_stat_day"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rt_data_stat_day"."active_player" IS '活躍玩家(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_day"."number_bettors" IS '註冊人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_day"."number_registrants" IS '投注人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_day"."odd_number" IS '注單數';
COMMENT ON COLUMN "public"."rt_data_stat_day"."total_betting" IS '總投注';
COMMENT ON COLUMN "public"."rt_data_stat_day"."game_tax" IS '遊戲抽水';
COMMENT ON COLUMN "public"."rt_data_stat_day"."platform_win_score" IS '平台總贏分數';
COMMENT ON COLUMN "public"."rt_data_stat_day"."platform_lose_score" IS '平台總輸分數';
COMMENT ON COLUMN "public"."rt_data_stat_day"."raw_data" IS '原始數據';
COMMENT ON TABLE "public"."rt_data_stat_day" IS '代理即時統計資料(realtime data stat)';

CREATE TABLE "public"."rt_data_stat_week" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "active_player" int4 NOT NULL DEFAULT 0,
  "number_bettors" int4 NOT NULL DEFAULT 0,
  "number_registrants" int4 NOT NULL DEFAULT 0,
  "odd_number" int4 NOT NULL DEFAULT 0,
  "total_betting" numeric(20,4) NOT NULL DEFAULT 0,
  "game_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_win_score" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_lose_score" numeric(20,4) NOT NULL DEFAULT 0,
  "raw_data" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "rt_data_stat_week_pkey" PRIMARY KEY ("log_time")
)
;
COMMENT ON COLUMN "public"."rt_data_stat_week"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_data_stat_week"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rt_data_stat_week"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rt_data_stat_week"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rt_data_stat_week"."active_player" IS '活躍玩家(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_week"."number_bettors" IS '註冊人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_week"."number_registrants" IS '投注人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_week"."odd_number" IS '注單數';
COMMENT ON COLUMN "public"."rt_data_stat_week"."total_betting" IS '總投注';
COMMENT ON COLUMN "public"."rt_data_stat_week"."game_tax" IS '遊戲抽水';
COMMENT ON COLUMN "public"."rt_data_stat_week"."platform_win_score" IS '平台總贏分數';
COMMENT ON COLUMN "public"."rt_data_stat_week"."platform_lose_score" IS '平台總輸分數';
COMMENT ON COLUMN "public"."rt_data_stat_week"."raw_data" IS '原始數據';
COMMENT ON TABLE "public"."rt_data_stat_week" IS '代理即時統計資料(realtime data stat)';

CREATE TABLE "public"."rt_data_stat_month" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "active_player" int4 NOT NULL DEFAULT 0,
  "number_bettors" int4 NOT NULL DEFAULT 0,
  "number_registrants" int4 NOT NULL DEFAULT 0,
  "odd_number" int4 NOT NULL DEFAULT 0,
  "total_betting" numeric(20,4) NOT NULL DEFAULT 0,
  "game_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_win_score" numeric(20,4) NOT NULL DEFAULT 0,
  "platform_lose_score" numeric(20,4) NOT NULL DEFAULT 0,
  "raw_data" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "rt_data_stat_month_pkey" PRIMARY KEY ("log_time")
)
;
COMMENT ON COLUMN "public"."rt_data_stat_month"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_data_stat_month"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "public"."rt_data_stat_month"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "public"."rt_data_stat_month"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."rt_data_stat_month"."active_player" IS '活躍玩家(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_month"."number_bettors" IS '註冊人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_month"."number_registrants" IS '投注人數(不重複)';
COMMENT ON COLUMN "public"."rt_data_stat_month"."odd_number" IS '注單數';
COMMENT ON COLUMN "public"."rt_data_stat_month"."total_betting" IS '總投注';
COMMENT ON COLUMN "public"."rt_data_stat_month"."game_tax" IS '遊戲抽水';
COMMENT ON COLUMN "public"."rt_data_stat_month"."platform_win_score" IS '平台總贏分數';
COMMENT ON COLUMN "public"."rt_data_stat_month"."platform_lose_score" IS '平台總輸分數';
COMMENT ON COLUMN "public"."rt_data_stat_month"."raw_data" IS '原始數據';
COMMENT ON TABLE "public"."rt_data_stat_month" IS '代理即時統計資料(realtime data stat)';

INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('e1ca6b64-5d13-4cc3-8e36-bd0ef8537a4f', '0 */1 * * * *', '每分鐘執行一次', 'job_rt_data_stat_backup_sec', 't', 0, '', '2022-11-23 02:55:36+00');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('89ab932a-2807-4ea8-9694-ffb66f12d567', '0 1 0 */1 * *', '每天的00:01分執行', 'job_rt_data_stat_backup_day', 't', 0, '', '2022-11-23 02:55:45+00');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('8bc970c5-832d-486a-8d6c-5502496179d1', '0 9 0 * * *', '每周一的00:09執行', 'job_rt_data_stat_backup_week', 't', 0, '', '2022-11-23 02:55:51+00');
INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('c340d2af-8b0d-41b4-abf5-6b22ff4bbf62', '0 14 0 1 */1 *', '每個月一號的00:14執行', 'job_rt_data_stat_backup_month', 't', 0, '', '2022-11-23 02:55:57+00');


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
  (200101, '第三方遊戲API', '/channel/channelHandle', true, '對外串接', false),
  (200102, '第三方取遊戲記錄', '/record/getRecordHandle', true, '對外串接', false);


-- 可操作權限 account_type 開發者:1, 總代理:2, 子代理:3
UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100215,100225,100226,100227,100228,100229,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100255,100256]}'
WHERE
  "account_type" = 1;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100253,100254,100255,100256]}'
WHERE
  "account_type" = 2;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100252,100253,100254,100255,100256]}'
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