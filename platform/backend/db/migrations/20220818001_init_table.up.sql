/*
 Navicat Premium Data Transfer

 Source Server         : dev_151
 Source Server Type    : PostgreSQL
 Source Server Version : 110016 (110016)
 Source Host           : 172.30.0.151:5432
 Source Catalog        : dcc_game
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110016 (110016)
 File Encoding         : 65001

 Date: 18/08/2022 13:30:48
*/


-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin_user";
CREATE TABLE "public"."admin_user" (
  "agent_id" int4 NOT NULL DEFAULT 0,
  "username" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "nickname" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "google_auth" bool NOT NULL DEFAULT false,
  "google_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "allow_ip" varchar(1000) COLLATE "pg_catalog"."default" NOT NULL DEFAULT '0.0.0.0/0'::character varying,
  "account_type" int4 NOT NULL DEFAULT 0,
  "is_readonly" int4 NOT NULL DEFAULT 0,
  "is_enabled" int4 NOT NULL DEFAULT 1,
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "admin_user_pkey" PRIMARY KEY ("username")
)
;
COMMENT ON COLUMN "public"."admin_user"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."admin_user"."username" IS '管理者帳號';
COMMENT ON COLUMN "public"."admin_user"."password" IS '管理者密碼';
COMMENT ON COLUMN "public"."admin_user"."nickname" IS '管理者暱稱';
COMMENT ON COLUMN "public"."admin_user"."google_auth" IS '是否開啟google驗證';
COMMENT ON COLUMN "public"."admin_user"."google_key" IS 'Google密鑰';
COMMENT ON COLUMN "public"."admin_user"."allow_ip" IS '用戶IP白名單';
COMMENT ON COLUMN "public"."admin_user"."account_type" IS '管理者類型 0:一般帳號, 1:總台號';
COMMENT ON COLUMN "public"."admin_user"."is_readonly" IS '是否為唯讀帳號 1.唯讀, 0.可做簡易編輯';
COMMENT ON COLUMN "public"."admin_user"."is_enabled" IS '是否啟用中1.啟用中,0停用中';
COMMENT ON COLUMN "public"."admin_user"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "public"."admin_user"."create_time" IS '創建時間';
COMMENT ON TABLE "public"."admin_user" IS '後台帳號設置';

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'dccuser', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc user', 'f', '', '', 1, 0, 1, '2022-07-11 15:29:36.573581', '2022-07-11 15:29:36.573581');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'admin', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc admin', 'f', '', '', 1, 0, 1, '2022-07-11 15:29:36.573581', '2022-07-11 15:29:36.573581');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'chouyang', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc chouyang', 'f', '', '', 1, 0, 1, '2022-08-10 10:00:54.748602', '2022-08-10 10:00:54.748602');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'kumei', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc kumei', 'f', '', '0.0.0.0/0', 1, 0, 1, '2022-08-22 15:02:31.097548', '2022-08-22 15:02:31.097548');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'Abby', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc abby', 'f', '', '0.0.0.0/0', 1, 0, 1, '2022-08-22 15:02:31.097548', '2022-08-22 15:02:31.097548');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'kenny', 'CJZiODlaKNqtw8S58q8cJjo6DGN_0rXbkD1n6qkB08o5Ow', 'New user', 'f', '', '', 1, 0, 1, '2022-08-25 02:29:15.54664', '2022-08-25 02:29:15.54664');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time") VALUES (1, 'kinco', 'hhx8WDdqL-vXZFd2vWa26jo6kKWuOmRP21P2X3tjZiwqFw', 'dcc kinco', 'f', '', '0.0.0.0/0', 1, 0, 1, '2022-08-22 15:02:31.097548', '2022-08-22 15:02:31.097548');
-- ----------------------------
-- Table structure for admin_user_action_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin_user_action_log";
CREATE TABLE "public"."admin_user_action_log" (
  "log_time" varchar(18) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "action_code" int4 NOT NULL DEFAULT 0,
  "action_log" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "ip" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  CONSTRAINT "admin_user_action_log_pkey" PRIMARY KEY ("log_time", "username")
)
;
COMMENT ON COLUMN "public"."admin_user_action_log"."log_time" IS '紀錄時間';
COMMENT ON COLUMN "public"."admin_user_action_log"."username" IS '管理者帳號';
COMMENT ON COLUMN "public"."admin_user_action_log"."action_code" IS '操作代號';
COMMENT ON COLUMN "public"."admin_user_action_log"."action_log" IS '操作紀錄';
COMMENT ON COLUMN "public"."admin_user_action_log"."ip" IS '登入IP';
COMMENT ON TABLE "public"."admin_user_action_log" IS '後台帳號操作記錄';

-- ----------------------------
-- Records of admin_user_action_log
-- ----------------------------
INSERT INTO "public"."admin_user_action_log" VALUES ('202208101000462855', 'admin', 0, '{"cost_time": "6ms", "request_ua": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36", "request_uri": "/api/v1/user/adminuser", "request_time": 1660125646279, "response_msg": "操作失敗", "request_proto": "HTTP/1.1", "response_code": 13, "response_data": "Create new admin user wrong, please notify customer service.", "response_time": 1660125646285, "request_method": "POST", "request_referer": "http://172.30.0.150:9986/swagger/index.html", "request_client_ip": "192.168.180.185", "request_post_data": ""}', '192.168.180.185');
INSERT INTO "public"."admin_user_action_log" VALUES ('202208101000547514', 'admin', 0, '{"cost_time": "5ms", "request_ua": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36", "request_uri": "/api/v1/user/adminuser", "request_time": 1660125654746, "response_msg": "操作成功", "request_proto": "HTTP/1.1", "response_code": 0, "response_data": "Create successed.", "response_time": 1660125654751, "request_method": "POST", "request_referer": "http://172.30.0.150:9986/swagger/index.html", "request_client_ip": "192.168.180.185", "request_post_data": ""}', '192.168.180.185');

-- ----------------------------
-- Table structure for agent
-- ----------------------------
DROP TABLE IF EXISTS "public"."agent";
CREATE TABLE "public"."agent" (
  "id" serial4,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "secret_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "md5_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_group_code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "commission" int4 NOT NULL DEFAULT 0,
  "info" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" int4 NOT NULL DEFAULT 1,
  "disable_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "is_top_agent" bool NOT NULL DEFAULT false,
  "top_agent_id" int2,
  CONSTRAINT "uni_agent_1" UNIQUE ("agent_group_code", "code"),
  CONSTRAINT "agent_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."agent"."id" IS '代理編號';
COMMENT ON COLUMN "public"."agent"."name" IS '代理名稱';
COMMENT ON COLUMN "public"."agent"."code" IS '代理識別編碼-隨機生成不重複唯一碼';
COMMENT ON COLUMN "public"."agent"."secret_key" IS '代理加解密使用密鑰-隨機生成不重複唯一碼';
COMMENT ON COLUMN "public"."agent"."md5_key" IS 'h5game專用md5密鑰';
COMMENT ON COLUMN "public"."agent"."agent_group_code" IS '集團識別編碼(mapping agent_group)';
COMMENT ON COLUMN "public"."agent"."commission" IS '分成(萬分之n)';
COMMENT ON COLUMN "public"."agent"."info" IS '廠商註記';
COMMENT ON COLUMN "public"."agent"."is_enabled" IS '代理啟用中(1:啟用,0:關閉)';
COMMENT ON COLUMN "public"."agent"."disable_time" IS '服務停止時間';
COMMENT ON COLUMN "public"."agent"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "public"."agent"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."agent"."is_top_agent" IS '是否為上級代理';
COMMENT ON COLUMN "public"."agent"."top_agent_id" IS '上級代理編號';
COMMENT ON TABLE "public"."agent" IS '代理商戶設置表';

-- ----------------------------
-- Sequence default value for agent
-- ----------------------------
SELECT setval('"public"."agent_id_seq"', 1000);

-- ----------------------------
-- Records of agent
-- ----------------------------
INSERT INTO "public"."agent" VALUES (1, 'dcc master', '1qaz2wsx', '6b14314760bd9280695a95d38082478b', '', '12345678', 0, '', 1, '1970-01-01 00:00:00', '2022-08-10 11:55:47.161639', '2022-08-10 11:55:47.161639', 'f', NULL);
INSERT INTO "public"."agent" VALUES (2, 'kenny', '2qaz2wsx', '6b14314760bd9280695a95d38082478b', '', '23456789', 0, '', 1, '1970-01-01 00:00:00', '2022-08-10 13:41:49.240062', '2022-08-10 13:41:49.240062', 'f', NULL);

-- ----------------------------
-- Table structure for agent_game
-- ----------------------------
DROP TABLE IF EXISTS "public"."agent_game";
CREATE TABLE "public"."agent_game" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "room_list" jsonb NOT NULL DEFAULT '{}'::jsonb,
  CONSTRAINT "agent_game_pkey" PRIMARY KEY ("agent_id", "game_id")
)
;
COMMENT ON COLUMN "public"."agent_game"."agent_id" IS '代理id';
COMMENT ON COLUMN "public"."agent_game"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."agent_game"."room_list" IS '開放房間列表';
COMMENT ON TABLE "public"."agent_game" IS '代理開放遊戲設定';

-- ----------------------------
-- Records of agent_game
-- ----------------------------
INSERT INTO "public"."agent_game" VALUES (1, 1001, '{"roomlist": [1, 2, 3, 4]}');

-- ----------------------------
-- Table structure for agent_game_room
-- ----------------------------
DROP TABLE IF EXISTS "public"."agent_game_room";
CREATE TABLE "public"."agent_game_room" (
  "agent_id" int2 NOT NULL,
  "agent_code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_md5_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "game_id" int4 NOT NULL,
  "game_code" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "game_room_id" int4 NOT NULL,
  "state" int2 NOT NULL DEFAULT 1,
  "room_type" int4 NOT NULL,
  CONSTRAINT "agent_game_room_pkey" PRIMARY KEY ("agent_id", "game_room_id")
)
;
COMMENT ON COLUMN "public"."agent_game_room"."agent_id" IS '代理編號(mapping agent id)';
COMMENT ON COLUMN "public"."agent_game_room"."agent_code" IS '代理識別編碼(mapping agent code)';
COMMENT ON COLUMN "public"."agent_game_room"."agent_md5_key" IS '代理h5game專用md5密鑰(mapping agent md5_key)';
COMMENT ON COLUMN "public"."agent_game_room"."game_id" IS '遊戲編號(mapping game id)';
COMMENT ON COLUMN "public"."agent_game_room"."game_code" IS '遊戲識別編碼(mapping game code)';
COMMENT ON COLUMN "public"."agent_game_room"."game_room_id" IS '房間編號';
COMMENT ON COLUMN "public"."agent_game_room"."state" IS '代理遊戲房間狀態';
COMMENT ON COLUMN "public"."agent_game_room"."room_type" IS '房間類型';
COMMENT ON TABLE "public"."agent_game_room" IS '代理開放遊戲房間設定表';

-- ----------------------------
-- Records of agent_game_room
-- ----------------------------
INSERT INTO "public"."agent_game_room" VALUES (1, '1qaz2wsx', '6b14314760bd9280695a95d38082478b', 1001, 'baccarat', 1, 1, 0);
INSERT INTO "public"."agent_game_room" VALUES (1, '1qaz2wsx', '6b14314760bd9280695a95d38082478b', 1001, 'baccarat', 2, 1, 1);
INSERT INTO "public"."agent_game_room" VALUES (1, '1qaz2wsx', '6b14314760bd9280695a95d38082478b', 1001, 'baccarat', 3, 1, 2);
INSERT INTO "public"."agent_game_room" VALUES (1, '1qaz2wsx', '6b14314760bd9280695a95d38082478b', 1001, 'baccarat', 4, 1, 3);

-- ----------------------------
-- Table structure for agent_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."agent_group";
CREATE TABLE "public"."agent_group" (
  "id" serial4,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "mode" int4 NOT NULL DEFAULT 0,
  "code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_count" int4 NOT NULL DEFAULT 0,
  "commission" int4 NOT NULL DEFAULT 0,
  "info" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" int4 NOT NULL DEFAULT 1,
  "expire_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "disable_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_agent_group_code" UNIQUE ("code"),
  CONSTRAINT "agent_group_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."agent_group"."id" IS '集團編號';
COMMENT ON COLUMN "public"."agent_group"."name" IS '集團名稱';
COMMENT ON COLUMN "public"."agent_group"."mode" IS '集團模式(1.平台自營,2.系統商)';
COMMENT ON COLUMN "public"."agent_group"."code" IS '集團識別編碼-隨機生成不重複唯一碼';
COMMENT ON COLUMN "public"."agent_group"."agent_count" IS '旗下代理數 平台自營 固定為1';
COMMENT ON COLUMN "public"."agent_group"."commission" IS '分成(萬分之n)';
COMMENT ON COLUMN "public"."agent_group"."info" IS '廠商註記';
COMMENT ON COLUMN "public"."agent_group"."is_enabled" IS 'BOSS啟用中(1:啟用,0:關閉)';
COMMENT ON COLUMN "public"."agent_group"."expire_time" IS '合約到期日';
COMMENT ON COLUMN "public"."agent_group"."disable_time" IS '服務停止時間';
COMMENT ON COLUMN "public"."agent_group"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "public"."agent_group"."create_time" IS '創建時間';
COMMENT ON TABLE "public"."agent_group" IS '總代理商戶設置/集團商務設置';

-- ----------------------------
-- Records of agent_group
-- ----------------------------
INSERT INTO "public"."agent_group" VALUES (1, 'DCC GAMING', 1, '12345678', 1, 0, '', 1, '1970-01-01 00:00:00', '1970-01-01 00:00:00', '2022-08-10 11:55:47.161639', '2022-08-10 11:55:47.161639');

-- ----------------------------
-- Table structure for exchange_data
-- ----------------------------
DROP TABLE IF EXISTS "public"."exchange_data";
CREATE TABLE "public"."exchange_data" (
  "id" serial4,
  "currency" varchar(3) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "to_cny" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "uni_exchange_data_id" UNIQUE ("id"),
  CONSTRAINT "exchange_data_pkey" PRIMARY KEY ("currency")
)
;
COMMENT ON COLUMN "public"."exchange_data"."id" IS '貨幣識別ID';
COMMENT ON COLUMN "public"."exchange_data"."currency" IS '原貨幣';
COMMENT ON COLUMN "public"."exchange_data"."to_cny" IS '轉換遊戲幣數額';
COMMENT ON TABLE "public"."exchange_data" IS '貨幣兌換遊戲幣表';

-- ----------------------------
-- Records of exchange_data
-- ----------------------------
INSERT INTO "public"."exchange_data" VALUES (1, 'CNY', 1.0000);

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS "public"."game";
CREATE TABLE "public"."game" (
  "id" int4 NOT NULL,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "state" int2 NOT NULL DEFAULT 1,
  "image" varchar COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "api" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "sup_lang" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "h5_link" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_game_code" UNIQUE ("code"),
  CONSTRAINT "game_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."game"."id" IS '遊戲id(PK)';
COMMENT ON COLUMN "public"."game"."name" IS '遊戲名稱';
COMMENT ON COLUMN "public"."game"."code" IS '遊戲識別編碼';
COMMENT ON COLUMN "public"."game"."state" IS '遊戲狀態';
COMMENT ON COLUMN "public"."game"."image" IS '遊戲圖片base64';
COMMENT ON COLUMN "public"."game"."api" IS 'game server api list';
COMMENT ON COLUMN "public"."game"."sup_lang" IS '支援語系';
COMMENT ON COLUMN "public"."game"."h5_link" IS '前端遊戲位置';
COMMENT ON COLUMN "public"."game"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."game"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."game" IS '遊戲列表';

-- ----------------------------
-- Records of game
-- ----------------------------
INSERT INTO "public"."game" VALUES (0, '大廳', 'lobby', 1, '', '{}', '{}', 'http://172.30.0.152/dev/lobby/', '2022-08-11 13:58:05.831006', '2022-08-11 13:58:05.831006');
INSERT INTO "public"."game" VALUES (1001, '百家樂', 'baccarat', 1, '', '{}', '{}', 'http://172.30.0.152/dev/lobby/', '2022-08-10 03:23:50.310687', '2022-08-10 03:23:50.310687');

-- ----------------------------
-- Table structure for game_room
-- ----------------------------
DROP TABLE IF EXISTS "public"."game_room";
CREATE TABLE "public"."game_room" (
  "id" int4 NOT NULL,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "state" int2 NOT NULL DEFAULT 1,
  "setting_info" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "game_id" int4 NOT NULL,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "room_type" int4 NOT NULL DEFAULT 1,
  CONSTRAINT "room_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."game_room"."id" IS '房間id(PK)';
COMMENT ON COLUMN "public"."game_room"."name" IS '房間名稱';
COMMENT ON COLUMN "public"."game_room"."state" IS '房間狀態';
COMMENT ON COLUMN "public"."game_room"."setting_info" IS '詳細內容';
COMMENT ON COLUMN "public"."game_room"."game_id" IS '遊戲id(FK)';
COMMENT ON COLUMN "public"."game_room"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."game_room"."update_time" IS '更新時間';
COMMENT ON COLUMN "public"."game_room"."room_type" IS '房間類型';
COMMENT ON TABLE "public"."game_room" IS '遊戲設定表';

-- ----------------------------
-- Records of game_room
-- ----------------------------
INSERT INTO "public"."game_room" VALUES (4, '百家樂', 1, '{}', 1001, '2022-08-18 10:14:25.895875', '2022-08-18 10:14:25.895875', 4);
INSERT INTO "public"."game_room" VALUES (3, '百家樂', 1, '{}', 1001, '2022-08-18 10:14:25.895875', '2022-08-18 10:14:25.895875', 3);
INSERT INTO "public"."game_room" VALUES (2, '百家樂', 1, '{}', 1001, '2022-08-18 10:14:25.895875', '2022-08-18 10:14:25.895875', 2);
INSERT INTO "public"."game_room" VALUES (1, '百家樂', 1, '{}', 1001, '2022-08-18 10:14:25.895875', '2022-08-18 10:14:25.895875', 1);

-- ----------------------------
-- Table structure for game_users
-- ----------------------------
DROP TABLE IF EXISTS "public"."game_users";
CREATE TABLE "public"."game_users" (
  "id" serial4,
  "agent_id" int4 NOT NULL,
  "original_username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "user_metadata" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "temporary_coin" numeric(20,4) NOT NULL DEFAULT 0,
  "coin" numeric(20,4) NOT NULL DEFAULT 0,
  "is_enabled" bool NOT NULL DEFAULT true,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "disabled_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  CONSTRAINT "game_users_pkey" PRIMARY KEY ("id", "agent_id", "original_username", "username"),
  CONSTRAINT "uni_id" UNIQUE ("id"),
  CONSTRAINT "uni_usrname" UNIQUE ("username")
)
;
COMMENT ON COLUMN "public"."game_users"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."game_users"."original_username" IS '用戶原平台帳號';
COMMENT ON COLUMN "public"."game_users"."username" IS '用戶帳號(after mapping)';
COMMENT ON COLUMN "public"."game_users"."user_metadata" IS '用戶基本資料';
COMMENT ON COLUMN "public"."game_users"."temporary_coin" IS '暫存遊戲幣(update in runtime)';
COMMENT ON COLUMN "public"."game_users"."coin" IS '遊戲幣';
COMMENT ON COLUMN "public"."game_users"."is_enabled" IS '是否開啟';
COMMENT ON COLUMN "public"."game_users"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."game_users"."update_time" IS '更新時間';
COMMENT ON COLUMN "public"."game_users"."disabled_time" IS '關閉時間(可預先設定帳號關閉時間)';
COMMENT ON TABLE "public"."game_users" IS '遊戲用戶表';

-- ----------------------------
-- Sequence default value for game_users
-- ----------------------------
SELECT setval('"public"."game_users_id_seq"', 1000);

-- ----------------------------
-- Records of game_users
-- ----------------------------
INSERT INTO "public"."game_users" VALUES (1, 1,'aaaa', 'dyg_454200d4d33f80c4', '{}', 0.0000, 100.0000, 't', '2022-08-09 17:18:45.812561', '2022-08-09 17:18:45.812561', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (2, 1,'kinco', 'dyg_d1dddec0e53379d8', '{}', 0.0000, 100.0000, 't', '2022-08-10 17:55:50.11068', '2022-08-10 17:55:50.11068', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (3, 1,'kinco900', 'dyg_d9feb5405361698e', '{}', 0.0000, 100.0000, 't', '2022-08-11 11:19:56.817023', '2022-08-11 11:19:56.817023', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (4, 1,'kkkkkhhh', 'dyg_ae63c83009149ef1', '{}', 0.0000, 100.0000, 't', '2022-08-10 09:10:59.228478', '2022-08-10 09:10:59.228478', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (5, 1,'simon', 'dyg_371c686298d32281', '{}', 0.0000, 10.0000, 't', '2022-08-11 10:35:10.564054', '2022-08-11 10:35:10.564054', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (6, 1,'simon1', 'dyg_6dbf57fed3b08565', '{}', 0.0000, 10.0000, 't', '2022-08-11 10:37:23.625108', '2022-08-11 10:37:23.625108', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (7, 1,'simon2', 'dyg_a11278841be8409a', '{}', 0.0000, 10.0000, 't', '2022-08-11 10:39:53.46264', '2022-08-11 10:39:53.46264', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (8, 1,'kinco901', 'dyg_645b06d41b29cd75', '{}', 0.0000, 0.0000, 't', '2022-08-11 11:46:08.317071', '2022-08-11 11:46:08.317071', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (9, 1,'simon3', 'dyg_7b506af9ab16f4d5', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:39:22.97836', '2022-08-11 13:39:22.97836', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (10, 1,'simon4', 'dyg_792bb0b8f668b877', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:47:45.818761', '2022-08-11 13:47:45.818761', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (11, 1,'simon5', 'dyg_41e228693259303b', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:50:24.388907', '2022-08-11 13:50:24.388907', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (12, 1,'simon6', 'dyg_16932c22d8627bd3', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:52:37.909104', '2022-08-11 13:52:37.909104', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (13, 1,'simon7', 'dyg_261561ac560b6794', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:55:21.806341', '2022-08-11 13:55:21.806341', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (14, 1,'simon8', 'dyg_575e02aeb2f52275', '{}', 0.0000, 10.0000, 't', '2022-08-11 13:59:35.811736', '2022-08-11 13:59:35.811736', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (15, 1,'simon10', 'dyg_6887137c4a2c40e8', '{}', 0.0000, 10.0000, 't', '2022-08-11 14:03:49.618964', '2022-08-11 14:03:49.618964', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (16, 1,'simon11', 'dyg_0242b9d5926b547e', '{}', 0.0000, 10.0000, 't', '2022-08-11 14:06:39.617852', '2022-08-11 14:06:39.617852', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (17, 1,'simon12', 'dyg_6d91130812c2d235', '{}', 0.0000, 10.0000, 't', '2022-08-11 14:27:17.405687', '2022-08-11 14:27:17.405687', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (18, 1,'aabb', 'dyg_dfac81c1e7dddcaf', '{}', 0.0000, 0.0000, 't', '2022-08-12 10:47:29.313053', '2022-08-12 10:47:29.313053', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (19, 1,'shinjuwu', 'dyg_e11ab02d3d83c674', '{}', 0.0000, 10.0000, 't', '2022-08-12 17:55:33.908531', '2022-08-12 17:55:33.908531', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (20, 1,'shinjuwu2', 'dyg_38f393750246cd49', '{}', 0.0000, 10.0000, 't', '2022-08-12 18:10:50.85413', '2022-08-12 18:10:50.85413', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (21, 1,'shinjuwu3', 'dyg_e4dd20e1b19b7ac9', '{}', 0.0000, 10.0000, 't', '2022-08-12 18:16:15.291914', '2022-08-12 18:16:15.291914', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (22, 1,'ggyy', 'dyg_25f96824b464dba1', '{}', 0.0000, 100.0000, 't', '2022-08-12 18:17:07.476154', '2022-08-12 18:17:07.476154', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (23, 1,'shinjuwu1', 'dyg_d6be9d579e69030f', '{}', 0.0000, 10.0000, 't', '2022-08-12 18:17:14.182939', '2022-08-12 18:17:14.182939', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (24, 1,'5566', 'dyg_16c7e8214fd659e4', '{}', 0.0000, 100.0000, 't', '2022-08-12 18:23:02.986277', '2022-08-12 18:23:02.986277', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (25, 1,'jarga', 'dyg_8b4433af7aaaa433', '{}', 0.0000, 5566.0000, 't', '2022-08-12 18:29:49.914323', '2022-08-12 18:29:49.914323', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (30, 1,'Tengyi', 'dyg_30e98658ca3dd4fb', '{}', 0.0000, 987654321.0000, 't', '2022-08-15 10:14:46.991383', '2022-08-15 10:14:46.991383', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (32, 1,'gt10001', 'dyg_b884eaf538a71d75', '{}', 0.0000, 1000.0000, 't', '2022-08-15 10:26:41.103391', '2022-08-15 10:26:41.103391', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (33, 1,'wilson', 'dyg_ba55577590736ef6', '{}', 0.0000, 100.0000, 't', '2022-08-15 10:26:41.774418', '2022-08-15 10:26:41.774418', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (38, 1,'Tenyi', 'dyg_3877ba5e677f21b3', '{}', 0.0000, 987654321.0000, 't', '2022-08-15 11:31:17.118093', '2022-08-15 11:31:17.118093', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (39, 1,'0', 'dyg_95d565ef66e7dff9', '{}', 0.0000, 0.0000, 't', '2022-08-15 16:51:34.885225', '2022-08-15 16:51:34.885225', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (40, 1,'gene001', 'dyg_e208af0a2f39ab79', '{}', 0.0000, 0.0000, 't', '2022-08-15 16:55:15.850209', '2022-08-15 16:55:15.850209', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (42, 1,'gt10002', 'dyg_6493f71ed2f1a97e', '{}', 0.0000, 12.0000, 't', '2022-08-15 17:10:53.429671', '2022-08-15 17:10:53.429671', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (43, 1,'gt1002', 'dyg_31291a67cf0ec33b', '{}', 0.0000, 12.0000, 't', '2022-08-15 17:16:33.51651', '2022-08-15 17:16:33.51651', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (49, 1,'kumei02', 'dyg_254b41ab697d1692', '{}', 0.0000, 1200.0000, 't', '2022-08-16 14:46:00.438716', '2022-08-16 14:46:00.438716', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (28, 1,'abby', 'dyg_98162eb6743f1728', '{}', 0.0000, 1100.0000, 't', '2022-08-15 09:38:42.492219', '2022-08-15 09:38:42.492219', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (29, 1,'kumei', 'dyg_f1889fe827e7084d', '{}', 0.0000, 10200.0000, 't', '2022-08-15 10:14:15.882714', '2022-08-15 10:14:15.882714', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (27, 1,'kenny', 'dyg_8d375a112998beac', '{}', 0.0000, 50.0000, 't', '2022-08-12 18:51:48.924314', '2022-08-12 18:51:48.924314', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (48, 1,'gt1001', 'dyg_aae974c9e700dd76', '{}', 0.0000, 0.0000, 't', '2022-08-16 14:32:52.325764', '2022-08-16 14:32:52.325764', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (54, 1,'kenny01', 'dyg_5d29cfdd8d64b651', '{}', 0.0000, 0.0000, 't', '2022-08-17 10:11:40.696881', '2022-08-17 10:11:40.696881', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (56, 1,'gt10003', 'dyg_994522ddca719cbd', '{}', 0.0000, 0.0000, 't', '2022-08-17 10:25:10.411841', '2022-08-17 10:25:10.411841', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (52, 1,'Action', 'dyg_a40003140292e973', '{}', 0.0000, 1555555.0000, 't', '2022-08-16 17:18:48.824034', '2022-08-16 17:18:48.824034', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (35, 1,'alan', 'dyg_324e7c4f269c6982', '{}', 0.0000, 250000.0000, 't', '2022-08-15 10:30:58.081229', '2022-08-15 10:30:58.081229', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (53, 1,'terry', 'dyg_2b04d4420477ee60', '{}', 0.0000, 3000.0000, 't', '2022-08-16 17:20:12.790198', '2022-08-16 17:20:12.790198', '1970-01-01 00:00:00');
INSERT INTO "public"."game_users" VALUES (47, 1,'chouyang', 'dyg_8e4fff5954157c38', '{}', 0.0000, 1000000.0000, 't', '2022-08-15 17:22:56.792261', '2022-08-15 17:22:56.792261', '1970-01-01 00:00:00');

-- ----------------------------
-- Table structure for play_log_baccarat
-- ----------------------------
DROP TABLE IF EXISTS "public"."play_log_baccarat";
CREATE TABLE "public"."play_log_baccarat" (
  "bet_id" serial8,
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_group_id" varchar(8) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "desk_id" int4 NOT NULL,
  "seat_id" int4 NOT NULL,
  "exchange" int4 NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "ya_score" numeric(20,4) NOT NULL,
  "valid_score" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL,
  CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("bet_id")
)
;
COMMENT ON COLUMN "public"."play_log_baccarat"."bet_id" IS '注單號';
COMMENT ON COLUMN "public"."play_log_baccarat"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."play_log_baccarat"."agent_group_id" IS '集團識別號';
COMMENT ON COLUMN "public"."play_log_baccarat"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "public"."play_log_baccarat"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "public"."play_log_baccarat"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."play_log_baccarat"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."play_log_baccarat"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."play_log_baccarat"."seat_id" IS '座位id';
COMMENT ON COLUMN "public"."play_log_baccarat"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."play_log_baccarat"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."play_log_baccarat"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."play_log_baccarat"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."play_log_baccarat"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "public"."play_log_baccarat"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "public"."play_log_baccarat"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."play_log_baccarat"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "public"."play_log_baccarat"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."play_log_baccarat"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."play_log_baccarat"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."play_log_baccarat" IS '玩家遊戲記錄';

-- ----------------------------
-- Sequence default value for play_log_baccarat
-- ----------------------------
SELECT setval('"public"."play_log_baccarat_bet_id_seq"', 1000000000, true);

-- ----------------------------
-- Table structure for play_log_common
-- ----------------------------
DROP TABLE IF EXISTS "public"."play_log_common";
CREATE TABLE "public"."play_log_common" (
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "desk_id" int4 NOT NULL,
  "exchange" int4 NOT NULL,
  "playlog" jsonb NOT NULL,
  "de_score" numeric(20,4) NOT NULL,
  "ya_score" numeric(20,4) NOT NULL,
  "valid_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL,
  CONSTRAINT "play_log_common_pkey" PRIMARY KEY ("lognumber", "game_id")
)
;
COMMENT ON COLUMN "public"."play_log_common"."lognumber" IS '單號';
COMMENT ON COLUMN "public"."play_log_common"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "public"."play_log_common"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."play_log_common"."desk_id" IS '桌子id';
COMMENT ON COLUMN "public"."play_log_common"."exchange" IS '一幣分值';
COMMENT ON COLUMN "public"."play_log_common"."playlog" IS '遊戲記錄';
COMMENT ON COLUMN "public"."play_log_common"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "public"."play_log_common"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "public"."play_log_common"."valid_score" IS '有效投注';
COMMENT ON COLUMN "public"."play_log_common"."create_time" IS '記錄時間';
COMMENT ON COLUMN "public"."play_log_common"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "public"."play_log_common"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "public"."play_log_common"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "public"."play_log_common" IS '遊戲局記錄';

-- ----------------------------
-- Table structure for wallet_ledger
-- ----------------------------
DROP TABLE IF EXISTS "public"."wallet_ledger";
CREATE TABLE "public"."wallet_ledger" (
  "id" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "changeset" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "wallet_ledger_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."wallet_ledger"."id" IS '訂單號(orderid)';
COMMENT ON COLUMN "public"."wallet_ledger"."user_id" IS '用戶id';
COMMENT ON COLUMN "public"."wallet_ledger"."username" IS '用戶帳號';
COMMENT ON COLUMN "public"."wallet_ledger"."changeset" IS '變更摘要';
COMMENT ON COLUMN "public"."wallet_ledger"."info" IS '詳細內容';
COMMENT ON COLUMN "public"."wallet_ledger"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."wallet_ledger"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."wallet_ledger" IS '帳變資料表';

-- ----------------------------
-- Indexes structure for table play_log_common
-- ----------------------------
CREATE INDEX "idx_play_log_common_1" ON "public"."play_log_common" USING btree (
  "create_time" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);


