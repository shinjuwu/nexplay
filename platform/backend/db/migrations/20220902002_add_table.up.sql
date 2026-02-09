CREATE TABLE "public"."permission_list" (
  "feature_code" int4 NOT NULL,
  "name" varchar(32) NOT NULL,
  "api_path" varchar(128) NOT NULL,
  "is_enabled" bool NOT NULL,
  "remark" varchar(255) NOT NULL DEFAULT '',
  "is_required" bool NOT NULL DEFAULT true,
  "create_time" timestamptz NOT NULL DEFAULT now(),
  "update_time" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("feature_code", "api_path"),
  CONSTRAINT "uni_feature_code" UNIQUE ("feature_code")
)
;

COMMENT ON COLUMN "public"."permission_list"."feature_code" IS '功能代碼';
COMMENT ON COLUMN "public"."permission_list"."name" IS '功能名稱';
COMMENT ON COLUMN "public"."permission_list"."is_enabled" IS '開關';
COMMENT ON COLUMN "public"."permission_list"."remark" IS '說明';
COMMENT ON COLUMN "public"."permission_list"."api_path" IS 'api 路徑';
COMMENT ON COLUMN "public"."permission_list"."is_required" IS '是否需要驗證才可使用(true:需驗證, false:無需驗證)';
COMMENT ON TABLE "public"."permission_list" IS '後台帳號權限表';

-- init data
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100001, '檢查伺服器是否活著', '/api/v1/example/health', 't', 'test', '2022-09-05 01:14:27.916512+00', '2022-09-05 01:14:27.916512+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100100, 'client取得驗證token', '/api/v1/intercom/getlogintoken', 't', '遊戲SERVER串接使用', '2022-09-05 06:20:22.082328+00', '2022-09-05 06:20:22.082328+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100101, '驗證token並登入遊戲', '/api/v1/intercom/logingame', 't', '遊戲SERVER串接使用', '2022-09-05 06:21:17.960166+00', '2022-09-05 06:21:17.960166+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100102, '遊戲用戶登出', '/api/v1/intercom/logoutgame', 't', '遊戲SERVER串接使用', '2022-09-05 06:22:18.683415+00', '2022-09-05 06:22:18.683415+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100103, '創建每局遊戲紀錄', '/api/v1/intercom/creategamerecord', 't', '遊戲SERVER串接使用', '2022-09-05 06:23:17.964023+00', '2022-09-05 06:23:17.964023+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100200, '取得代理遊戲房間列表', '/api/v1/agent/agentgameroom', 't', '後台使用', '2022-09-05 06:25:19.126548+00', '2022-09-05 06:25:19.126548+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100201, '遊戲設置列表Config', '/api/v1/game/gamesettinglistconfig', 't', '後台使用', '2022-09-05 06:26:33.284588+00', '2022-09-05 06:26:33.284588+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100202, '取得遊戲設置列表', '/api/v1/game/gamesettinglist', 't', '後台使用', '2022-09-05 06:26:50.645539+00', '2022-09-05 06:26:50.645539+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100203, '遊戲設置列表Config', '/api/v1/game/setgamestate', 't', '後台使用', '2022-09-05 06:27:53.908613+00', '2022-09-05 06:27:53.908613+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100204, '查詢遊戲用戶列表', '/api/v1/game/gameuserlist', 't', '後台使用', '2022-09-05 06:28:17.399453+00', '2022-09-05 06:28:17.399453+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100205, '此接口用來重新載入本地端資料', '/api/v1/global/roloadglobaldata', 't', '後台使用', '2022-09-05 06:28:40.943294+00', '2022-09-05 06:28:40.943294+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100206, '取得驗證碼', '/api/v1/login/captcha', 't', '後台使用', '2022-09-05 06:29:07.287572+00', '2022-09-05 06:29:07.287572+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100207, '用戶登錄', '/api/v1/login/login', 't', '後台使用', '2022-09-05 06:29:27.025051+00', '2022-09-05 06:29:27.025051+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100208, '輸贏報表Config', '/api/v1/record/winlosereportconfig', 't', '後台使用', '2022-09-05 06:30:02.547576+00', '2022-09-05 06:30:02.547576+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100209, '輸贏報表', '/api/v1/record/winlosereport', 't', '後台使用', '2022-09-05 06:30:20.296554+00', '2022-09-05 06:30:20.296554+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100210, '遊戲日誌解析', '/api/v1/record/gamelog', 't', '後台使用', '2022-09-05 06:30:41.126336+00', '2022-09-05 06:30:41.126336+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100211, '上下分紀錄Config', '/api/v1/record/walletledgerconfig', 't', '後台使用', '2022-09-05 06:31:00.711452+00', '2022-09-05 06:31:00.711452+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100212, '上下分紀錄', '/api/v1/record/walletledger', 't', '後台使用', '2022-09-05 06:31:23.990416+00', '2022-09-05 06:31:23.990416+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100213, '取得目前已產生的有效 token list', '/api/v1/user/getalivetokenlist', 't', '後台使用', '2022-09-05 06:31:53.799861+00', '2022-09-05 06:31:53.799861+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100214, '將用戶 token 列入黑名單(主動使登入token 失效)', '/api/v1/user/blacktoken', 't', '後台使用', '2022-09-05 06:32:12.464255+00', '2022-09-05 06:32:12.464255+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100215, '創建後台帳號', '/api/v1/user/adminuser', 't', '後台使用', '2022-09-05 06:32:33.078566+00', '2022-09-05 06:32:33.078566+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100216, 'ping(must verify)', '/api/v1/user/ping', 't', '後台使用', '2022-09-05 06:33:23.536493+00', '2022-09-05 06:33:23.536493+00', 'f');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (200101, '第三方遊戲API', '/channel/channelHandle', 't', '對外串接', '2022-09-05 06:16:59.774567+00', '2022-09-05 06:16:59.774567+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (200102, '第三方取遊戲記錄', '/record/getRecordHandle', 't', '對外串接', '2022-09-05 06:15:56.333726+00', '2022-09-05 06:15:56.333726+00', 't');


ALTER TABLE "public"."admin_user" 
  ADD COLUMN "permission" jsonb NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN "public"."admin_user"."permission" IS '帳號權限list';

-- set all
UPDATE admin_user SET permission='{"list":[100200,100201,100202,100203,100204,100205,100208,100209,100210,100211,100212,100213,100214,100215,100216,200101,200102]}';
