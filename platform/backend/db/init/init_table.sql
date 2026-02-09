/*
 Navicat Premium Data Transfer

 Source Server         : ete-demo-backend-pgsql
 Source Server Type    : PostgreSQL
 Source Server Version : 140004 (140004)
 Source Catalog        : dcc_game
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140004 (140004)
 File Encoding         : 65001

 Date: 09/12/2022 15:15:39
*/


-- ----------------------------
-- Sequence structure for agent_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "agent_id_seq";
CREATE SEQUENCE "agent_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for exchange_data_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "exchange_data_id_seq";
CREATE SEQUENCE "exchange_data_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for game_users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "game_users_id_seq";
CREATE SEQUENCE "game_users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_baccarat_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_baccarat_bet_id_seq";
CREATE SEQUENCE "user_play_log_baccarat_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_blackjack_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_blackjack_bet_id_seq";
CREATE SEQUENCE "user_play_log_blackjack_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_colordisc_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_colordisc_bet_id_seq";
CREATE SEQUENCE "user_play_log_colordisc_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_fantan_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_fantan_bet_id_seq";
CREATE SEQUENCE "user_play_log_fantan_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_prawncrab_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_prawncrab_bet_id_seq";
CREATE SEQUENCE "user_play_log_prawncrab_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_play_log_sangong_bet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "user_play_log_sangong_bet_id_seq";
CREATE SEQUENCE "user_play_log_sangong_bet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS "admin_user";
CREATE TABLE "admin_user" (
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
  "is_added" bool NOT NULL DEFAULT true,
  "login_time" timestamp(6) NOT NULL DEFAULT now(),
  "role" uuid NOT NULL,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "admin_user"."agent_id" IS '代理編號';
COMMENT ON COLUMN "admin_user"."username" IS '管理者帳號';
COMMENT ON COLUMN "admin_user"."password" IS '管理者密碼';
COMMENT ON COLUMN "admin_user"."nickname" IS '管理者暱稱';
COMMENT ON COLUMN "admin_user"."google_auth" IS '是否開啟google驗證';
COMMENT ON COLUMN "admin_user"."google_key" IS 'Google密鑰';
COMMENT ON COLUMN "admin_user"."allow_ip" IS '用戶IP白名單';
COMMENT ON COLUMN "admin_user"."account_type" IS '管理者類型 1:總台號,2:總代理,3:子代理';
COMMENT ON COLUMN "admin_user"."is_readonly" IS '是否為唯讀帳號 1.唯讀, 0.可做簡易編輯';
COMMENT ON COLUMN "admin_user"."is_enabled" IS '是否啟用中1.啟用中,0停用中';
COMMENT ON COLUMN "admin_user"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "admin_user"."create_time" IS '創建時間';
COMMENT ON COLUMN "admin_user"."is_added" IS '是否為分身帳號';
COMMENT ON COLUMN "admin_user"."login_time" IS '最後登入時間';
COMMENT ON COLUMN "admin_user"."role" IS '角色';
COMMENT ON COLUMN "admin_user"."info" IS '備註';
COMMENT ON TABLE "admin_user" IS '後台帳號設置';

-- ----------------------------
-- Table structure for admin_user_action_log
-- ----------------------------
DROP TABLE IF EXISTS "admin_user_action_log";
CREATE TABLE "admin_user_action_log" (
  "log_time" varchar(18) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "error_code" int4 NOT NULL DEFAULT 0,
  "action_log" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "ip" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "method" varchar(6) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "request_url" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "admin_user_action_log"."log_time" IS '紀錄時間';
COMMENT ON COLUMN "admin_user_action_log"."username" IS '管理者帳號';
COMMENT ON COLUMN "admin_user_action_log"."error_code" IS '錯誤碼';
COMMENT ON COLUMN "admin_user_action_log"."action_log" IS '操作紀錄';
COMMENT ON COLUMN "admin_user_action_log"."ip" IS '登入IP';
COMMENT ON COLUMN "admin_user_action_log"."method" IS 'api method';
COMMENT ON COLUMN "admin_user_action_log"."request_url" IS '請求路徑';
COMMENT ON TABLE "admin_user_action_log" IS '後台帳號操作記錄';

-- ----------------------------
-- Table structure for agent
-- ----------------------------
DROP TABLE IF EXISTS "agent";
CREATE TABLE "agent" (
  "id" int4 NOT NULL DEFAULT nextval('agent_id_seq'::regclass),
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "secret_key" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "md5_key" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "commission" int4 NOT NULL DEFAULT 0,
  "info" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" int4 NOT NULL DEFAULT 1,
  "disable_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "is_top_agent" bool NOT NULL DEFAULT false,
  "top_agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "cooperation" int2 NOT NULL DEFAULT 1,
  "coin_limit" numeric(20,4) NOT NULL DEFAULT 0,
  "coin_use" numeric(20,4) NOT NULL DEFAULT 0,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "aes_key" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "member_count" int4 NOT NULL DEFAULT 0,
  "creator" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ip_whitelist" jsonb NOT NULL DEFAULT '[]'::jsonb
)
;
COMMENT ON COLUMN "agent"."id" IS '代理編號';
COMMENT ON COLUMN "agent"."name" IS '代理名稱';
COMMENT ON COLUMN "agent"."code" IS '代理識別編碼-隨機生成不重複唯一碼';
COMMENT ON COLUMN "agent"."secret_key" IS '代理加解密使用密鑰-隨機生成不重複唯一碼';
COMMENT ON COLUMN "agent"."md5_key" IS 'h5game專用md5密鑰';
COMMENT ON COLUMN "agent"."commission" IS '分成(萬分之n)';
COMMENT ON COLUMN "agent"."info" IS '廠商註記';
COMMENT ON COLUMN "agent"."is_enabled" IS '代理啟用中(1:啟用,0:關閉)';
COMMENT ON COLUMN "agent"."disable_time" IS '服務停止時間';
COMMENT ON COLUMN "agent"."update_time" IS '資料更新時間';
COMMENT ON COLUMN "agent"."create_time" IS '創建時間';
COMMENT ON COLUMN "agent"."is_top_agent" IS '是否為上級代理';
COMMENT ON COLUMN "agent"."top_agent_id" IS '上級代理編號';
COMMENT ON COLUMN "agent"."cooperation" IS '合作模式(代理結帳類型, 1: 買分, 2: 信用)';
COMMENT ON COLUMN "agent"."coin_limit" IS '買分模式分數上限';
COMMENT ON COLUMN "agent"."coin_use" IS '已使用分數';
COMMENT ON COLUMN "agent"."level_code" IS '層級碼,每四碼一個層級 (admin:0000)';
COMMENT ON COLUMN "agent"."aes_key" IS 'h5game專用aes密鑰';
COMMENT ON COLUMN "agent"."member_count" IS '會員人數';
COMMENT ON COLUMN "agent"."creator" IS '創建此代理的後台帳號username';
COMMENT ON TABLE "agent" IS '代理商戶設置表';

-- ----------------------------
-- Table structure for agent_game
-- ----------------------------
DROP TABLE IF EXISTS "agent_game";
CREATE TABLE "agent_game" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "state" int2 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "agent_game"."agent_id" IS '代理id';
COMMENT ON COLUMN "agent_game"."game_id" IS '遊戲id';
COMMENT ON COLUMN "agent_game"."state" IS '房間狀態';
COMMENT ON TABLE "agent_game" IS '代理開放遊戲設定';

-- ----------------------------
-- Table structure for agent_game_room
-- ----------------------------
DROP TABLE IF EXISTS "agent_game_room";
CREATE TABLE "agent_game_room" (
  "agent_id" int2 NOT NULL,
  "game_room_id" int4 NOT NULL,
  "state" int2 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "agent_game_room"."agent_id" IS '代理編號(mapping agent id)';
COMMENT ON COLUMN "agent_game_room"."game_room_id" IS '房間編號';
COMMENT ON COLUMN "agent_game_room"."state" IS '代理遊戲房間狀態';
COMMENT ON TABLE "agent_game_room" IS '代理開放遊戲房間設定表';

-- ----------------------------
-- Table structure for agent_permission
-- ----------------------------
DROP TABLE IF EXISTS "agent_permission";
CREATE TABLE "agent_permission" (
  "agent_id" int4 NOT NULL DEFAULT 0,
  "permission" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "account_type" int2 NOT NULL DEFAULT 0,
  "id" uuid NOT NULL DEFAULT gen_random_uuid()
)
;
COMMENT ON COLUMN "agent_permission"."agent_id" IS 'mapping agent id';
COMMENT ON COLUMN "agent_permission"."permission" IS '帳號權限list';
COMMENT ON COLUMN "agent_permission"."name" IS '角色名稱';
COMMENT ON COLUMN "agent_permission"."info" IS '備註';
COMMENT ON COLUMN "agent_permission"."account_type" IS '帳號角色';
COMMENT ON COLUMN "agent_permission"."id" IS '代理權限角色編號';

-- ----------------------------
-- Table structure for agent_wallet
-- ----------------------------
DROP TABLE IF EXISTS "agent_wallet";
CREATE TABLE "agent_wallet" (
  "agent_id" int4 NOT NULL DEFAULT 0,
  "amount" numeric(20,4) NOT NULL DEFAULT 0
)
;

-- ----------------------------
-- Table structure for agent_wallet_ledger
-- ----------------------------
DROP TABLE IF EXISTS "agent_wallet_ledger";
CREATE TABLE "agent_wallet_ledger" (
  "id" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "agent_id" int4 NOT NULL DEFAULT 0,
  "changeset" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "kind" int2 NOT NULL DEFAULT 0,
  "creator" varchar(20) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Table structure for exchange_data
-- ----------------------------
DROP TABLE IF EXISTS "exchange_data";
CREATE TABLE "exchange_data" (
  "id" int4 NOT NULL DEFAULT nextval('exchange_data_id_seq'::regclass),
  "currency" varchar(3) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "to_cny" numeric(20,4) NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "exchange_data"."id" IS '貨幣識別ID';
COMMENT ON COLUMN "exchange_data"."currency" IS '原貨幣';
COMMENT ON COLUMN "exchange_data"."to_cny" IS '轉換遊戲幣數額';
COMMENT ON TABLE "exchange_data" IS '貨幣兌換遊戲幣表';

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS "game";
CREATE TABLE "game" (
  "id" int4 NOT NULL,
  "server_info_code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "state" int2 NOT NULL DEFAULT 1,
  "image" varchar COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "h5_link" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "type" int4 NOT NULL DEFAULT 0,
  "room_number" int4 NOT NULL DEFAULT 0,
  "table_number" int4 NOT NULL DEFAULT 0,
  "cal_state" int2 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "game"."id" IS '遊戲id(PK)';
COMMENT ON COLUMN "game"."server_info_code" IS '遊戲的server code';
COMMENT ON COLUMN "game"."name" IS '遊戲名稱';
COMMENT ON COLUMN "game"."code" IS '遊戲識別編碼';
COMMENT ON COLUMN "game"."state" IS '遊戲狀態(1:open,2:close)';
COMMENT ON COLUMN "game"."image" IS '遊戲圖片base64';
COMMENT ON COLUMN "game"."h5_link" IS '前端遊戲位置';
COMMENT ON COLUMN "game"."create_time" IS '創建時間';
COMMENT ON COLUMN "game"."update_time" IS '更新時間';
COMMENT ON COLUMN "game"."type" IS '遊戲類型';
COMMENT ON COLUMN "game"."room_number" IS '房間數量';
COMMENT ON COLUMN "game"."table_number" IS '桌子數量';
COMMENT ON COLUMN "game"."cal_state" IS '是否開啟計算報表(1:open,2:close)';
COMMENT ON TABLE "game" IS '遊戲列表';

-- ----------------------------
-- Table structure for game_room
-- ----------------------------
DROP TABLE IF EXISTS "game_room";
CREATE TABLE "game_room" (
  "id" int4 NOT NULL,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "state" int2 NOT NULL DEFAULT 1,
  "setting_info" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "game_id" int4 NOT NULL,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "room_type" int4 NOT NULL DEFAULT 1
)
;
COMMENT ON COLUMN "game_room"."id" IS '房間id(PK)';
COMMENT ON COLUMN "game_room"."name" IS '房間名稱';
COMMENT ON COLUMN "game_room"."state" IS '房間狀態';
COMMENT ON COLUMN "game_room"."setting_info" IS '詳細內容';
COMMENT ON COLUMN "game_room"."game_id" IS '遊戲id(FK)';
COMMENT ON COLUMN "game_room"."create_time" IS '創建時間';
COMMENT ON COLUMN "game_room"."update_time" IS '更新時間';
COMMENT ON COLUMN "game_room"."room_type" IS '房間類型';
COMMENT ON TABLE "game_room" IS '遊戲設定表';

-- ----------------------------
-- Table structure for game_users
-- ----------------------------
DROP TABLE IF EXISTS "game_users";
CREATE TABLE "game_users" (
  "id" int4 NOT NULL DEFAULT nextval('game_users_id_seq'::regclass),
  "agent_id" int4 NOT NULL,
  "original_username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "user_metadata" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "sum_coin_in" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_coin_out" numeric(20,4) NOT NULL DEFAULT 0,
  "is_enabled" bool NOT NULL DEFAULT true,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "disabled_time" timestamp(6) NOT NULL DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "game_users"."agent_id" IS '代理編號';
COMMENT ON COLUMN "game_users"."original_username" IS '用戶原平台帳號';
COMMENT ON COLUMN "game_users"."username" IS '用戶帳號(after mapping)';
COMMENT ON COLUMN "game_users"."user_metadata" IS '用戶基本資料';
COMMENT ON COLUMN "game_users"."sum_coin_in" IS '累積遊戲幣轉入';
COMMENT ON COLUMN "game_users"."sum_coin_out" IS '累積遊戲幣轉出';
COMMENT ON COLUMN "game_users"."is_enabled" IS '是否開啟';
COMMENT ON COLUMN "game_users"."create_time" IS '創建時間';
COMMENT ON COLUMN "game_users"."update_time" IS '更新時間';
COMMENT ON COLUMN "game_users"."disabled_time" IS '關閉時間(可預先設定帳號關閉時間)';
COMMENT ON COLUMN "game_users"."level_code" IS '代理層級碼';
COMMENT ON COLUMN "game_users"."info" IS '備註';
COMMENT ON TABLE "game_users" IS '遊戲用戶表';

-- ----------------------------
-- Table structure for job_scheduler
-- ----------------------------
DROP TABLE IF EXISTS "job_scheduler";
CREATE TABLE "job_scheduler" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "spec" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "trigger_func" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT false,
  "exec_limit" int2 NOT NULL DEFAULT 0,
  "last_sync_date" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "update_time" timestamptz(0) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "job_scheduler"."spec" IS '執行時間';
COMMENT ON COLUMN "job_scheduler"."info" IS '備註';
COMMENT ON COLUMN "job_scheduler"."trigger_func" IS '執行func，需要系統有支援才可使用';
COMMENT ON COLUMN "job_scheduler"."is_enabled" IS '是否開啟';
COMMENT ON COLUMN "job_scheduler"."exec_limit" IS '指定執行次數';
COMMENT ON COLUMN "job_scheduler"."last_sync_date" IS '最後同步日期辨識用字串(YYYYMMDDhhmm)';
COMMENT ON COLUMN "job_scheduler"."update_time" IS '最後更新時間';
COMMENT ON TABLE "job_scheduler" IS '排程設定表';

-- ----------------------------
-- Table structure for marquee
-- ----------------------------
DROP TABLE IF EXISTS "marquee";
CREATE TABLE "marquee" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "lang" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 0,
  "type" int2 NOT NULL DEFAULT 0,
  "order" int2 NOT NULL DEFAULT 0,
  "freq" int2 NOT NULL DEFAULT 0,
  "is_enabled" bool NOT NULL DEFAULT false,
  "is_open" bool NOT NULL DEFAULT false,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "start_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "marquee"."lang" IS '語系';
COMMENT ON COLUMN "marquee"."type" IS '跑馬燈類型';
COMMENT ON COLUMN "marquee"."order" IS '順序';
COMMENT ON COLUMN "marquee"."freq" IS '播放頻率(每?秒播放一次)';
COMMENT ON COLUMN "marquee"."is_enabled" IS '是否啟動';
COMMENT ON COLUMN "marquee"."is_open" IS '是否開啟';
COMMENT ON COLUMN "marquee"."content" IS '內文';
COMMENT ON COLUMN "marquee"."start_time" IS '開始時間';
COMMENT ON COLUMN "marquee"."end_time" IS '結束時間';
COMMENT ON COLUMN "marquee"."create_time" IS '創建時間';
COMMENT ON COLUMN "marquee"."update_time" IS '更新時間';
COMMENT ON TABLE "marquee" IS '跑馬燈設定表';

-- ----------------------------
-- Table structure for permission_list
-- ----------------------------
DROP TABLE IF EXISTS "permission_list";
CREATE TABLE "permission_list" (
  "feature_code" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "api_path" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "is_enabled" bool NOT NULL,
  "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_required" bool NOT NULL DEFAULT true,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "permission_list"."feature_code" IS '功能代碼';
COMMENT ON COLUMN "permission_list"."name" IS '功能名稱';
COMMENT ON COLUMN "permission_list"."api_path" IS 'api 路徑';
COMMENT ON COLUMN "permission_list"."is_enabled" IS '開關';
COMMENT ON COLUMN "permission_list"."remark" IS '說明';
COMMENT ON COLUMN "permission_list"."is_required" IS '是否需要驗證才可使用(true:需驗證, false:無需驗證)';
COMMENT ON TABLE "permission_list" IS '後台帳號權限表';

-- ----------------------------
-- Table structure for play_log_common
-- ----------------------------
DROP TABLE IF EXISTS "play_log_common";
CREATE TABLE "play_log_common" (
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
  "tax" numeric(20,4) NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "play_log_common"."lognumber" IS '單號';
COMMENT ON COLUMN "play_log_common"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "play_log_common"."room_type" IS '房間類型';
COMMENT ON COLUMN "play_log_common"."desk_id" IS '桌子id';
COMMENT ON COLUMN "play_log_common"."exchange" IS '一幣分值';
COMMENT ON COLUMN "play_log_common"."playlog" IS '遊戲記錄';
COMMENT ON COLUMN "play_log_common"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "play_log_common"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "play_log_common"."valid_score" IS '有效投注';
COMMENT ON COLUMN "play_log_common"."create_time" IS '記錄時間';
COMMENT ON COLUMN "play_log_common"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "play_log_common"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "play_log_common"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "play_log_common"."tax" IS '抽水';
COMMENT ON TABLE "play_log_common" IS '遊戲局記錄';

-- ----------------------------
-- Table structure for rp_agent_stat
-- ----------------------------
DROP TABLE IF EXISTS "rp_agent_stat";
CREATE TABLE "rp_agent_stat" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "log_hour" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat"."log_hour" IS 'YYYYMMDDhh';
COMMENT ON COLUMN "rp_agent_stat"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat"."sum_bonus" IS '總額外獎勵';
COMMENT ON COLUMN "rp_agent_stat"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_day
-- ----------------------------
DROP TABLE IF EXISTS "rp_agent_stat_day";
CREATE TABLE "rp_agent_stat_day" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat_day"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat_day"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat_day"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat_day"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat_day"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat_day"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat_day"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat_day"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat_day"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat_day"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "rp_agent_stat_day"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat_day" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_hour
-- ----------------------------
DROP TABLE IF EXISTS "rp_agent_stat_hour";
CREATE TABLE "rp_agent_stat_hour" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat_hour"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat_hour"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat_hour"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat_hour"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat_hour"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat_hour"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat_hour"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat_hour"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat_hour"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat_hour"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "rp_agent_stat_hour"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat_hour" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_month
-- ----------------------------
DROP TABLE IF EXISTS "rp_agent_stat_month";
CREATE TABLE "rp_agent_stat_month" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat_month"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat_month"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat_month"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat_month"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat_month"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat_month"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat_month"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat_month"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat_month"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat_month"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "rp_agent_stat_month"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat_month" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for rp_agent_stat_week
-- ----------------------------
DROP TABLE IF EXISTS "rp_agent_stat_week";
CREATE TABLE "rp_agent_stat_week" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT '-1'::integer,
  "agent_name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "bet_user" int4 NOT NULL DEFAULT 0,
  "bet_count" int4 NOT NULL DEFAULT 0,
  "sum_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_vaild_ya" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_de" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "sum_tax" numeric(20,4) NOT NULL DEFAULT 0,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "rp_agent_stat_week"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "rp_agent_stat_week"."agent_id" IS '代理商編號';
COMMENT ON COLUMN "rp_agent_stat_week"."agent_name" IS '代理商名稱';
COMMENT ON COLUMN "rp_agent_stat_week"."level_code" IS '層級碼';
COMMENT ON COLUMN "rp_agent_stat_week"."bet_user" IS '投注總人數';
COMMENT ON COLUMN "rp_agent_stat_week"."bet_count" IS '投注總單量';
COMMENT ON COLUMN "rp_agent_stat_week"."sum_ya" IS '總壓';
COMMENT ON COLUMN "rp_agent_stat_week"."sum_vaild_ya" IS '總有效壓注';
COMMENT ON COLUMN "rp_agent_stat_week"."sum_de" IS '總得';
COMMENT ON COLUMN "rp_agent_stat_week"."sum_bonus" IS '額外獎勵';
COMMENT ON COLUMN "rp_agent_stat_week"."sum_tax" IS '總抽水';
COMMENT ON TABLE "rp_agent_stat_week" IS '代理時段統計資料(record profit)';

-- ----------------------------
-- Table structure for server_info
-- ----------------------------
DROP TABLE IF EXISTS "server_info";
CREATE TABLE "server_info" (
  "code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ip" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "addresses" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT false,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now()
)
;
COMMENT ON COLUMN "server_info"."code" IS '唯一代碼';
COMMENT ON COLUMN "server_info"."ip" IS 'server address';
COMMENT ON COLUMN "server_info"."addresses" IS 'connect addresses';
COMMENT ON COLUMN "server_info"."info" IS '備註';
COMMENT ON COLUMN "server_info"."is_enabled" IS '開關';
COMMENT ON TABLE "server_info" IS '遊戲伺服器連線位置紀錄表';

-- ----------------------------
-- Table structure for user_play_log_baccarat
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_baccarat";
CREATE TABLE "user_play_log_baccarat" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_baccarat_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "user_play_log_baccarat"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_baccarat"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_baccarat"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_baccarat"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_baccarat"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_baccarat"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_baccarat"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_baccarat"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_baccarat"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_baccarat"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_baccarat"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_baccarat"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_baccarat"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_baccarat"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_baccarat"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_baccarat"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_baccarat"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_baccarat"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_baccarat"."bet_time" IS '遊戲結算時間';
COMMENT ON COLUMN "user_play_log_baccarat"."tax" IS '抽水';
COMMENT ON TABLE "user_play_log_baccarat" IS '玩家遊戲記錄';

-- ----------------------------
-- Table structure for user_play_log_blackjack
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_blackjack";
CREATE TABLE "user_play_log_blackjack" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_blackjack_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "user_play_log_blackjack"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_blackjack"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_blackjack"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_blackjack"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_blackjack"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_blackjack"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_blackjack"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_blackjack"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_blackjack"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_blackjack"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_blackjack"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_blackjack"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_blackjack"."tax" IS '抽水';
COMMENT ON COLUMN "user_play_log_blackjack"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_blackjack"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_blackjack"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_blackjack"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_blackjack"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_blackjack"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_blackjack"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "user_play_log_blackjack" IS '21點玩家遊戲記錄';

-- ----------------------------
-- Table structure for user_play_log_colordisc
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_colordisc";
CREATE TABLE "user_play_log_colordisc" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_colordisc_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "user_play_log_colordisc"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_colordisc"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_colordisc"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_colordisc"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_colordisc"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_colordisc"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_colordisc"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_colordisc"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_colordisc"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_colordisc"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_colordisc"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_colordisc"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_colordisc"."tax" IS '抽水';
COMMENT ON COLUMN "user_play_log_colordisc"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_colordisc"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_colordisc"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_colordisc"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_colordisc"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_colordisc"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_colordisc"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "user_play_log_colordisc" IS '色碟玩家遊戲記錄';

-- ----------------------------
-- Table structure for user_play_log_fantan
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_fantan";
CREATE TABLE "user_play_log_fantan" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_fantan_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "user_play_log_fantan"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_fantan"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_fantan"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_fantan"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_fantan"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_fantan"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_fantan"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_fantan"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_fantan"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_fantan"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_fantan"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_fantan"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_fantan"."tax" IS '抽水';
COMMENT ON COLUMN "user_play_log_fantan"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_fantan"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_fantan"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_fantan"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_fantan"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_fantan"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_fantan"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "user_play_log_fantan" IS '番攤玩家遊戲記錄';

-- ----------------------------
-- Table structure for user_play_log_prawncrab
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_prawncrab";
CREATE TABLE "user_play_log_prawncrab" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_prawncrab_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "user_play_log_prawncrab"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_prawncrab"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_prawncrab"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_prawncrab"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_prawncrab"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_prawncrab"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_prawncrab"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_prawncrab"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_prawncrab"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_prawncrab"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_prawncrab"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_prawncrab"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_prawncrab"."tax" IS '抽水';
COMMENT ON COLUMN "user_play_log_prawncrab"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_prawncrab"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_prawncrab"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_prawncrab"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_prawncrab"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_prawncrab"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_prawncrab"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "user_play_log_prawncrab" IS '番攤玩家遊戲記錄';

-- ----------------------------
-- Table structure for user_play_log_sangong
-- ----------------------------
DROP TABLE IF EXISTS "user_play_log_sangong";
CREATE TABLE "user_play_log_sangong" (
  "bet_id" int8 NOT NULL DEFAULT nextval('user_play_log_sangong_bet_id_seq'::regclass),
  "lognumber" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
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
  "tax" numeric(20,4) NOT NULL,
  "start_score" numeric(20,4) NOT NULL,
  "end_score" numeric(20,4) NOT NULL,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "is_robot" int4 NOT NULL,
  "is_big_win" bool NOT NULL,
  "is_issue" bool NOT NULL,
  "bet_time" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "user_play_log_sangong"."bet_id" IS '注單號';
COMMENT ON COLUMN "user_play_log_sangong"."lognumber" IS '單號';
COMMENT ON COLUMN "user_play_log_sangong"."agent_id" IS '代理識別號';
COMMENT ON COLUMN "user_play_log_sangong"."user_id" IS '代理用戶id';
COMMENT ON COLUMN "user_play_log_sangong"."game_id" IS '請參照game_setting表';
COMMENT ON COLUMN "user_play_log_sangong"."room_type" IS '房間類型';
COMMENT ON COLUMN "user_play_log_sangong"."desk_id" IS '桌子id';
COMMENT ON COLUMN "user_play_log_sangong"."seat_id" IS '座位id';
COMMENT ON COLUMN "user_play_log_sangong"."exchange" IS '一幣分值';
COMMENT ON COLUMN "user_play_log_sangong"."de_score" IS '總得遊戲分';
COMMENT ON COLUMN "user_play_log_sangong"."ya_score" IS '總壓遊戲分';
COMMENT ON COLUMN "user_play_log_sangong"."valid_score" IS '有效投注';
COMMENT ON COLUMN "user_play_log_sangong"."tax" IS '抽水';
COMMENT ON COLUMN "user_play_log_sangong"."start_score" IS '玩家壓住前遊戲分';
COMMENT ON COLUMN "user_play_log_sangong"."end_score" IS '玩家壓住後遊戲分';
COMMENT ON COLUMN "user_play_log_sangong"."create_time" IS '記錄時間';
COMMENT ON COLUMN "user_play_log_sangong"."is_robot" IS '是否為機器人';
COMMENT ON COLUMN "user_play_log_sangong"."is_big_win" IS '是否為大獎';
COMMENT ON COLUMN "user_play_log_sangong"."is_issue" IS '是否為問題單';
COMMENT ON COLUMN "user_play_log_sangong"."bet_time" IS '遊戲結算時間';
COMMENT ON TABLE "user_play_log_sangong" IS '三公玩家遊戲記錄';

-- ----------------------------
-- Table structure for wallet_ledger
-- ----------------------------
DROP TABLE IF EXISTS "wallet_ledger";
CREATE TABLE "wallet_ledger" (
  "id" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" int4 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "changeset" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  "agent_id" int4 NOT NULL DEFAULT 0,
  "kind" int2 NOT NULL DEFAULT 0,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "status" int2 NOT NULL DEFAULT 0,
  "error_code" int4 NOT NULL DEFAULT 0,
  "creator" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "wallet_ledger"."id" IS '訂單號(orderid)';
COMMENT ON COLUMN "wallet_ledger"."user_id" IS '用戶id';
COMMENT ON COLUMN "wallet_ledger"."username" IS '用戶帳號';
COMMENT ON COLUMN "wallet_ledger"."changeset" IS '變更摘要';
COMMENT ON COLUMN "wallet_ledger"."info" IS '詳細內容';
COMMENT ON COLUMN "wallet_ledger"."create_time" IS '創建時間';
COMMENT ON COLUMN "wallet_ledger"."update_time" IS '更新時間';
COMMENT ON COLUMN "wallet_ledger"."agent_id" IS 'mapping id of agent';
COMMENT ON COLUMN "wallet_ledger"."kind" IS '上下分類型(1:上分,2:下分)';
COMMENT ON COLUMN "wallet_ledger"."level_code" IS '層級碼';
COMMENT ON COLUMN "wallet_ledger"."status" IS '訂單狀態';
COMMENT ON COLUMN "wallet_ledger"."error_code" IS '錯誤代碼';
COMMENT ON COLUMN "wallet_ledger"."creator" IS '創建人';
COMMENT ON TABLE "wallet_ledger" IS '帳變資料表';

-- ----------------------------
-- Function structure for postgres_fdw_disconnect
-- ----------------------------
DROP FUNCTION IF EXISTS "postgres_fdw_disconnect"(text);
CREATE OR REPLACE FUNCTION "postgres_fdw_disconnect"(text)
  RETURNS "pg_catalog"."bool" AS '$libdir/postgres_fdw', 'postgres_fdw_disconnect'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for postgres_fdw_disconnect_all
-- ----------------------------
DROP FUNCTION IF EXISTS "postgres_fdw_disconnect_all"();
CREATE OR REPLACE FUNCTION "postgres_fdw_disconnect_all"()
  RETURNS "pg_catalog"."bool" AS '$libdir/postgres_fdw', 'postgres_fdw_disconnect_all'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for postgres_fdw_get_connections
-- ----------------------------
DROP FUNCTION IF EXISTS "postgres_fdw_get_connections"(OUT "server_name" text, OUT "valid" bool);
CREATE OR REPLACE FUNCTION "postgres_fdw_get_connections"(OUT "server_name" text, OUT "valid" bool)
  RETURNS SETOF "pg_catalog"."record" AS '$libdir/postgres_fdw', 'postgres_fdw_get_connections'
  LANGUAGE c VOLATILE STRICT
  COST 1
  ROWS 1000;

-- ----------------------------
-- Function structure for postgres_fdw_handler
-- ----------------------------
DROP FUNCTION IF EXISTS "postgres_fdw_handler"();
CREATE OR REPLACE FUNCTION "postgres_fdw_handler"()
  RETURNS "pg_catalog"."fdw_handler" AS '$libdir/postgres_fdw', 'postgres_fdw_handler'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for postgres_fdw_validator
-- ----------------------------
DROP FUNCTION IF EXISTS "postgres_fdw_validator"(_text, oid);
CREATE OR REPLACE FUNCTION "postgres_fdw_validator"(_text, oid)
  RETURNS "pg_catalog"."void" AS '$libdir/postgres_fdw', 'postgres_fdw_validator'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Procedure structure for procedure_update_agent_game_room_state_by_agent
-- ----------------------------
DROP PROCEDURE IF EXISTS "procedure_update_agent_game_room_state_by_agent"("target_agent_id" int4, "target_game_room_id" int4, "agent_game_room_state" int4);
CREATE OR REPLACE PROCEDURE "procedure_update_agent_game_room_state_by_agent"("target_agent_id" int4, "target_game_room_id" int4, "agent_game_room_state" int4)
 AS $BODY$
BEGIN
  UPDATE "public"."agent_game_room"
    SET "state" = agent_game_room_state
    WHERE "agent_id" = target_agent_id AND "game_room_id" = target_game_room_id;
END
$BODY$
  LANGUAGE plpgsql;

-- ----------------------------
-- Procedure structure for procedure_update_agent_game_room_state_by_top_agent
-- ----------------------------
DROP PROCEDURE IF EXISTS "procedure_update_agent_game_room_state_by_top_agent"("target_top_agent_id" int4, "target_agent_level_code" varchar, "target_game_room_id" int4, "agent_game_room_state" int4);
CREATE OR REPLACE PROCEDURE "procedure_update_agent_game_room_state_by_top_agent"("target_top_agent_id" int4, "target_agent_level_code" varchar, "target_game_room_id" int4, "agent_game_room_state" int4)
 AS $BODY$
BEGIN
  UPDATE "public"."agent_game_room" AS "agr"
    SET "state" = agent_game_room_state
    FROM "public"."agent" AS "a"
    WHERE "agr"."agent_id" = "a"."id"
      AND "a"."level_code" LIKE target_agent_level_code || '%'
      AND "agr"."game_room_id" = target_game_room_id;

  CALL "public"."procedure_update_agent_game_room_state_by_agent"(target_top_agent_id, target_game_room_id, agent_game_room_state);
END
$BODY$
  LANGUAGE plpgsql;

-- ----------------------------
-- Procedure structure for procedure_update_agent_game_state_by_agent
-- ----------------------------
DROP PROCEDURE IF EXISTS "procedure_update_agent_game_state_by_agent"("target_agent_id" int4, "target_game_id" int4, "agent_game_state" int4);
CREATE OR REPLACE PROCEDURE "procedure_update_agent_game_state_by_agent"("target_agent_id" int4, "target_game_id" int4, "agent_game_state" int4)
 AS $BODY$
BEGIN
  UPDATE "public"."agent_game"
    SET "state" = agent_game_state
    WHERE "agent_id" = target_agent_id AND "game_id" = target_game_id;
END
$BODY$
  LANGUAGE plpgsql;

-- ----------------------------
-- Procedure structure for procedure_update_agent_game_state_by_top_agent
-- ----------------------------
DROP PROCEDURE IF EXISTS "procedure_update_agent_game_state_by_top_agent"("target_top_agent_id" int4, "target_agent_level_code" varchar, "target_game_id" int4, "agent_game_state" int4);
CREATE OR REPLACE PROCEDURE "procedure_update_agent_game_state_by_top_agent"("target_top_agent_id" int4, "target_agent_level_code" varchar, "target_game_id" int4, "agent_game_state" int4)
 AS $BODY$
BEGIN
  UPDATE "public"."agent_game" AS "ag"
    SET "state" = agent_game_state
    FROM "public"."agent" AS "a"
    WHERE "ag"."agent_id" = "a"."id"
      AND "a"."level_code" LIKE target_agent_level_code || '%'
      AND "ag"."game_id" = target_game_id;

  CALL "public"."procedure_update_agent_game_state_by_agent"(target_top_agent_id, target_game_id, agent_game_state);
END
$BODY$
  LANGUAGE plpgsql;

-- ----------------------------
-- Procedure structure for sp_create_agent_wallet_ledger
-- ----------------------------
DROP PROCEDURE IF EXISTS "sp_create_agent_wallet_ledger"("_id" varchar, "_agent_id" int4, "_before_coin" numeric, "_add_coin" numeric, "_after_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar);
CREATE OR REPLACE PROCEDURE "sp_create_agent_wallet_ledger"("_id" varchar, "_agent_id" int4, "_before_coin" numeric, "_add_coin" numeric, "_after_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
 AS $BODY$
DECLARE
  "_changeset" jsonb;
BEGIN
  "_changeset" = jsonb_build_object(
    'before_coin', "_before_coin",
    'add_coin', "_add_coin",
    'after_coin', "_after_coin"
  );

  INSERT INTO "public"."agent_wallet_ledger" ("id", "agent_id", "changeset", "info", "kind", "creator")
    VALUES ("_id", "_agent_id", "_changeset", "_info", "_kind", "_creator");
END;
$BODY$
  LANGUAGE plpgsql;

-- ----------------------------
-- Function structure for udf_backend_update_agent_wallect
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_backend_update_agent_wallect"("_id" varchar, "_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar);
CREATE OR REPLACE FUNCTION "udf_backend_update_agent_wallect"("_id" varchar, "_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_kind" int2, "_creator" varchar)
  RETURNS "pg_catalog"."bool" AS $BODY$
DECLARE
  "ret_result" boolean := false;
  "_agent_wallet_amount" numeric := 0;
BEGIN
  SELECT "public"."udf_update_agent_wallet"("_agent_id", "_add_coin") INTO "_agent_wallet_amount";
  IF "_agent_wallet_amount" < 0 THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", -"_add_coin");
  ELSE
    CALL "public"."sp_create_agent_wallet_ledger" ("_id", "_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", "_kind", "_creator");
    "ret_result" = true;
  END IF;

  RETURN "ret_result";
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_backend_update_agent_wallects
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_backend_update_agent_wallects"("_from_id" varchar, "_from_agent_id" int4, "_to_id" varchar, "_to_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_creator" varchar);
CREATE OR REPLACE FUNCTION "udf_backend_update_agent_wallects"("_from_id" varchar, "_from_agent_id" int4, "_to_id" varchar, "_to_agent_id" int4, "_add_coin" numeric, "_info" varchar, "_creator" varchar)
  RETURNS "pg_catalog"."bool" AS $BODY$
DECLARE
  "ret_result" boolean := false;
  "_agent_wallet_amount" numeric := 0;
BEGIN
  IF "_add_coin" > 0 THEN
    SELECT "public"."udf_update_agent_wallet"("_from_agent_id", -"_add_coin") INTO "_agent_wallet_amount";
    IF "_agent_wallet_amount" < 0 THEN
      PERFORM "public"."udf_update_agent_wallet"("_from_agent_id", "_add_coin");
    ELSE
      CALL "public"."sp_create_agent_wallet_ledger" ("_from_id", "_from_agent_id", "_agent_wallet_amount" + "_add_coin", -"_add_coin", "_agent_wallet_amount", "_info", 4::int2, "_creator");

      SELECT "public"."udf_update_agent_wallet"("_to_agent_id", "_add_coin") INTO "_agent_wallet_amount";
      CALL "public"."sp_create_agent_wallet_ledger" ("_to_id", "_to_agent_id", "_agent_wallet_amount" - "_add_coin", "_add_coin", "_agent_wallet_amount", "_info", 3::int2, "_creator");

      "ret_result" = true;
    END IF;
  END IF;
  
  RETURN "ret_result";
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_check_game_users_data
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar);
CREATE OR REPLACE FUNCTION "udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar)
  RETURNS "pg_catalog"."json" AS $BODY$
	DECLARE
	_user_id int4;
	_username varchar;
 	_is_new bool;
BEGIN

	SELECT 
		"id", "username", false  into _user_id, _username, _is_new
	FROM
		"public"."game_users" 
	WHERE
		agent_id = _agent_id AND original_username = _original_username;

	IF NOT FOUND THEN

		INSERT INTO "public"."game_users" 
			( "agent_id", "original_username", "username", "user_metadata", "sum_coin_in", "sum_coin_out") 
		SELECT
			_agent_id, _original_username, _trans_username, '{}', 0, _coin
		RETURNING "id", "username", true into _user_id, _username, _is_new
		;
		
		UPDATE agent SET member_count = member_count+1, update_time = now() WHERE "id" = _agent_id
		;
	END IF;
	
	RETURN json_build_object(
	'id', _user_id,
	'username', _username,
	'is_new', _is_new
	);

END
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_create_admin_user
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_create_admin_user"("_agent_id" int4, "_username" varchar, "_password" varchar, "_nickname" varchar, "_account_type" int4, "_is_readonly" int4, "_is_added" bool, "_role" uuid, "_info" varchar);
CREATE OR REPLACE FUNCTION "udf_create_admin_user"("_agent_id" int4, "_username" varchar, "_password" varchar, "_nickname" varchar, "_account_type" int4, "_is_readonly" int4, "_is_added" bool, "_role" uuid, "_info" varchar)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
    "ret_agent_id" int4;
    "ret_username" varchar;
    "ret_password" varchar;
    "ret_nickname" varchar;
    "ret_account_type" int4;
    "ret_allow_ip" varchar;
    "ret_is_readonly" int4;
    "ret_is_enabled" int4;
    "ret_update_time" timestamp;
    "ret_create_time" timestamp;
    "ret_is_added" bool;
    "ret_login_time" timestamp;
    "ret_role" uuid;
    "ret_info" varchar;
BEGIN
    INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "account_type",
        "is_readonly", "is_added", "role", "info")
        VALUES ("_agent_id", "_username", "_password", "_nickname", "_account_type", "_is_readonly",
            "_is_added", "_role", "_info")
        RETURNING "agent_id", "username", "password", "nickname", "account_type", "allow_ip",
            "is_readonly", "is_enabled", "update_time", "create_time", "is_added", "login_time",
            "role", "info" INTO "ret_agent_id", "ret_username", "ret_password", "ret_nickname",
            "ret_account_type", "ret_allow_ip", "ret_is_readonly", "ret_is_enabled", 
            "ret_update_time", "ret_create_time", "ret_is_added", "ret_login_time", "ret_role",
            "ret_info";

    RETURN json_build_object(
        'agent_id', "ret_agent_id",
        'username', "ret_username",
        'password', "ret_password",
        'nickname', "ret_nickname",
        'account_type', "ret_account_type",
        'allow_ip', "ret_allow_ip",
        'is_readonly', "ret_is_readonly",
        'is_enabled', "ret_is_enabled",
        'update_time', extract(epoch from "ret_update_time") * 1000000,
        'create_time', extract(epoch from "ret_create_time") * 1000000,
        'is_added', "ret_is_added",
        'login_time', extract(epoch from "ret_login_time") * 1000000,
        'role', "ret_role",
        'info', "ret_info"
    );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_create_agent
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4);
CREATE OR REPLACE FUNCTION "udf_create_agent"("_agent_name" varchar, "_agent_code" varchar, "_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar, "_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_ip_whitelist" jsonb, "_agent_creator" varchar, "_agent_commission" int4, "_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool, "_admin_user_username" varchar, "_admin_user_password" varchar, "_admin_user_nickname" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4, "_admin_user_is_readonly" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "ret_agent_id" int4;
  "ret_agent_name" varchar;
  "ret_agent_code" varchar;
  "ret_agent_secret_key" varchar;
  "ret_agent_aes_key" varchar;
  "ret_agent_md5_key" varchar;
  "ret_agent_commission" int4;
  "ret_agent_info" varchar;
  "ret_agent_is_enabled" int4;
  "ret_agent_disable_time" timestamp;
  "ret_agent_update_time" timestamp;
  "ret_agent_create_time" timestamp;
  "ret_agent_is_top_agent" bool;
  "ret_agent_top_agent_id" int4;
  "ret_agent_cooperation" int2;
  "ret_agent_coin_limit" decimal;
  "ret_agent_coin_use" decimal;
  "ret_agent_level_code" varchar;
  "ret_agent_ip_whitelist" jsonb;
  "ret_agent_creator" varchar;
  "ret_admin_user" json;
  "ret_agent_games" json;
  "ret_agent_game_rooms" json;
BEGIN
  INSERT INTO "public"."agent" ("name", "code", "level_code", "info", "secret_key",
    "aes_key", "md5_key", "ip_whitelist", "creator", "commission", "cooperation",
    "top_agent_id", "is_top_agent")
    VALUES ("_agent_name", "_agent_code", "_agent_level_code", "_agent_info",
      "_agent_secret_key", "_agent_aes_key", "_agent_md5_key", "_agent_ip_whitelist",
      "_agent_creator", "_agent_commission", "_agent_cooperation", "_agent_top_agent_id", 
      "_agent_is_top_agent")
    RETURNING "id", "name", "code", "secret_key", "aes_key", "md5_key", "commission", "info",
      "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id",
      "cooperation", "coin_limit", "coin_use", "ip_whitelist", "creator" INTO
      "ret_agent_id", "ret_agent_name", "ret_agent_code", "ret_agent_secret_key",
      "ret_agent_aes_key", "ret_agent_md5_key", "ret_agent_commission", "ret_agent_info",
      "ret_agent_is_enabled", "ret_agent_disable_time", "ret_agent_update_time",
      "ret_agent_create_time", "ret_agent_is_top_agent", "ret_agent_top_agent_id",
      "ret_agent_cooperation", "ret_agent_coin_limit", "ret_agent_coin_use",
      "ret_agent_ip_whitelist", "ret_agent_creator";

  UPDATE "public"."agent"
    SET "level_code" = "level_code" || LPAD(to_hex("ret_agent_id"), 4, '0')
    WHERE "id" = "ret_agent_id"
    RETURNING "level_code" INTO "ret_agent_level_code";

  IF "ret_agent_cooperation" = 1 THEN
    INSERT INTO "public"."agent_wallet" ("agent_id")
      VALUES ("ret_agent_id");
  END IF;

  SELECT "public"."udf_create_admin_user" ("ret_agent_id", "_admin_user_username", "_admin_user_password",
    "_admin_user_nickname", "_admin_user_account_type", "_admin_user_is_readonly", false, "_admin_user_role",
    "_admin_user_info") INTO "ret_admin_user";

  SELECT "public"."udf_create_agent_games" ("ret_agent_id", "ret_agent_top_agent_id") INTO "ret_agent_games";

  SELECT "public"."udf_create_agent_game_rooms" ("ret_agent_id", "ret_agent_top_agent_id") INTO "ret_agent_game_rooms";

  RETURN json_build_object(
    'agent', json_build_object(
      'id', "ret_agent_id",
      'name', "ret_agent_name",
      'code', "ret_agent_code",
      'level_code', "ret_agent_level_code",
      'secret_key', "ret_agent_secret_key",
      'aes_key', "ret_agent_aes_key",
      'md5_key', "ret_agent_md5_key",
      'commission', "ret_agent_commission",
      'info', "ret_agent_info",
      'is_enabled', "ret_agent_is_enabled",
      'disable_time', extract(epoch from "ret_agent_disable_time") * 1000000,
      'update_time', extract(epoch from "ret_agent_update_time") * 1000000,
      'create_time', extract(epoch from "ret_agent_create_time") * 1000000,
      'is_top_agent', "ret_agent_is_top_agent",
      'top_agent_id', "ret_agent_top_agent_id",
      'cooperation', "ret_agent_cooperation",
      'coin_limit', "ret_agent_coin_limit",
      'coin_use', "ret_agent_coin_use",
      'ip_whitelist', "ret_agent_ip_whitelist",
      'creator', "ret_agent_creator"
    ),
    'admin_user', "ret_admin_user",
    'agent_games', "ret_agent_games",
    'agent_game_rooms', "ret_agent_game_rooms"
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_create_agent_game_rooms
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_create_agent_game_rooms"("_agent_id" int4, "_top_agent_id" int4);
CREATE OR REPLACE FUNCTION "udf_create_agent_game_rooms"("_agent_id" int4, "_top_agent_id" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "ret_agent_game_rooms" json;
BEGIN
  INSERT INTO "public"."agent_game_room" ("agent_id", "game_room_id", "state")
    SELECT "_agent_id" AS "agent_id", "game_room_id", "state"
      FROM "public"."agent_game_room"
      WHERE "agent_id" = "_top_agent_id";
	  
  SELECT json_agg("agent_game_rooms") INTO "ret_agent_game_rooms"
    FROM (
	  SELECT "agent_id", "game_room_id", "state"
		  FROM "public"."agent_game_room"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_game_rooms";
  
  RETURN "ret_agent_game_rooms";
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_create_agent_games
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_create_agent_games"("_agent_id" int4, "_top_agent_id" int4);
CREATE OR REPLACE FUNCTION "udf_create_agent_games"("_agent_id" int4, "_top_agent_id" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "ret_agent_games" json;
BEGIN
  INSERT INTO "public"."agent_game" ("agent_id", "game_id", "state")
    SELECT "_agent_id" AS "agent_id", "game_id", "state"
      FROM "public"."agent_game"
      WHERE "agent_id" = "_top_agent_id";
	  
  SELECT json_agg("agent_games") INTO "ret_agent_games"
    FROM (
	  SELECT "agent_id", "game_id", "state"
		  FROM "public"."agent_game"
		  WHERE "agent_id" = "_agent_id"
	) AS "agent_games";
  
  RETURN "ret_agent_games";
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_game_user_complete_transfer
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_game_user_complete_transfer"("_id" varchar, "_changeset" jsonb, "_status" int2, "_error_code" int2, "_update_agent_wallet" bool, "_agent_id" int4, "_add_agent_wallet_amount" numeric);
CREATE OR REPLACE FUNCTION "udf_game_user_complete_transfer"("_id" varchar, "_changeset" jsonb, "_status" int2, "_error_code" int2, "_update_agent_wallet" bool, "_agent_id" int4, "_add_agent_wallet_amount" numeric)
  RETURNS "pg_catalog"."json" AS $BODY$
BEGIN
  IF "_add_agent_wallet_amount" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  IF "_update_agent_wallet" THEN
    PERFORM "public"."udf_update_agent_wallet"("_agent_id", "_add_agent_wallet_amount");
  END IF;

  UPDATE "public"."wallet_ledger"
    SET "changeset" = "_changeset",
      "status" = "_status",
      "error_code" = "_error_code",
	    "update_time" = now()
    WHERE "id" = "_id";

  RETURN json_build_object(
    'code', 0
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_game_user_create_transfer
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_game_user_create_transfer"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" jsonb, "_creator" varchar, "_update_agent_wallet" bool, "_add_agent_wallet_amount" numeric);
CREATE OR REPLACE FUNCTION "udf_game_user_create_transfer"("_id" varchar, "_user_id" int4, "_username" varchar, "_agent_id" int4, "_agent_level_code" varchar, "_kind" int2, "_status" int2, "_info" jsonb, "_creator" varchar, "_update_agent_wallet" bool, "_add_agent_wallet_amount" numeric)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
  "_agent_wallet_amount" numeric;
  "_insert_wallet_ledger_count" bigint;
BEGIN
  IF "_add_agent_wallet_amount" < 0 THEN
    RETURN json_build_object(
      'code', 1
    );
  END IF;

  IF "_update_agent_wallet" THEN
    SELECT "public"."udf_update_agent_wallet"("_agent_id", -"_add_agent_wallet_amount") INTO "_agent_wallet_amount";
    IF "_agent_wallet_amount" < 0 THEN
      PERFORM "public"."udf_update_agent_wallet"("_agent_id", "_add_agent_wallet_amount");

      RETURN json_build_object(
        'code', 2
      );
    END IF;
  END IF;

  INSERT INTO "public"."wallet_ledger" ("id", "user_id", "username", "agent_id", "level_code", "kind", "status", "info", "creator")
    VALUES ("_id", "_user_id", "_username", "_agent_id", "_agent_level_code", "_kind", "_status", "_info", "_creator")
    ON CONFLICT ("id") DO NOTHING;

  GET DIAGNOSTICS "_insert_wallet_ledger_count" = ROW_COUNT;
  IF "_insert_wallet_ledger_count" = 0 THEN
    RETURN json_build_object(
      'code', 3
    );
  END IF;

  RETURN json_build_object(
    'code', 0
  );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_update_admin_user
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_update_admin_user"("_agent_id" int4, "_username" varchar, "_role" uuid, "_info" varchar, "_is_enabled" int4);
CREATE OR REPLACE FUNCTION "udf_update_admin_user"("_agent_id" int4, "_username" varchar, "_role" uuid, "_info" varchar, "_is_enabled" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
    "ret_update_time" timestamp;
BEGIN
    UPDATE "public"."admin_user"
        SET "role" = "_role",
            "info" = "_info",
            "is_enabled" = "_is_enabled",
            "update_time" = now()
        WHERE "agent_id" = "_agent_id" AND "username" = "_username"
        RETURNING "update_time" INTO "ret_update_time";

    RETURN json_build_object(
        'update_time', extract(epoch from "ret_update_time") * 1000000
    );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_update_agent
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4);
CREATE OR REPLACE FUNCTION "udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4)
  RETURNS "pg_catalog"."json" AS $BODY$
DECLARE
    "ret_agent_update_time" timestamp;
    "ret_admin_user" json;
BEGIN
    UPDATE "public"."agent"
            SET "name" = "_agent_name",
                "info" = "_agent_info",
                "commission" = "_agent_commission",
                "update_time" = now()
            WHERE "id" = "_agent_id"
            RETURNING "update_time" INTO "ret_agent_update_time";

    SELECT "public"."udf_update_admin_user"("_agent_id", "_admin_user_username", "_admin_user_role",
        "_admin_user_info", "_admin_user_is_enabled") INTO "ret_admin_user";

    RETURN json_build_object(
        'agent', json_build_object(
            'update_time', extract(epoch from "ret_agent_update_time") * 1000000
        ),
        'admin_user', "ret_admin_user"
    );
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for udf_update_agent_wallet
-- ----------------------------
DROP FUNCTION IF EXISTS "udf_update_agent_wallet"("_agent_id" int4, "_amount" numeric);
CREATE OR REPLACE FUNCTION "udf_update_agent_wallet"("_agent_id" int4, "_amount" numeric)
  RETURNS "pg_catalog"."numeric" AS $BODY$
DECLARE
  "ret_amount" numeric;
Begin
  UPDATE "public"."agent_wallet"
    SET "amount" = "amount" + "_amount"
    WHERE "agent_id" = "_agent_id"
    RETURNING "amount" INTO "ret_amount";

  RETURN "ret_amount";
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- View structure for view_admin_user
-- ----------------------------
DROP VIEW IF EXISTS "view_admin_user";
CREATE VIEW "view_admin_user" AS  SELECT au.agent_id,
    au.username,
    au.password,
    au.nickname,
    au.google_auth,
    au.google_key,
    au.allow_ip,
    au.account_type,
    au.is_readonly,
    au.is_enabled,
    ap.permission,
    au.is_added,
    au.role
   FROM admin_user au
     JOIN agent_permission ap ON ap.id = au.role;

-- ----------------------------
-- View structure for view_agent_game_room
-- ----------------------------
DROP VIEW IF EXISTS "view_agent_game_room";
CREATE VIEW "view_agent_game_room" AS  SELECT agr.agent_id,
    agr.game_room_id,
    agr.state,
    a.name AS agent_name,
    a.code AS agent_code,
    g.id AS game_id,
    g.code AS game_code,
    gr.room_type
   FROM agent_game_room agr
     JOIN agent a ON agr.agent_id = a.id
     JOIN game_room gr ON agr.game_room_id = gr.id
     JOIN game g ON gr.game_id = g.id;

-- ----------------------------
-- View structure for view_agent_game
-- ----------------------------
DROP VIEW IF EXISTS "view_agent_game";
CREATE VIEW "view_agent_game" AS  SELECT ag.agent_id,
    ag.game_id,
    ag.state,
    a.name AS agent_name,
    a.code AS agent_code,
    a.level_code AS agent_level_code,
    g.code AS game_code,
    g.state AS game_state
   FROM agent_game ag
     JOIN agent a ON ag.agent_id = a.id
     JOIN game g ON ag.game_id = g.id
  ORDER BY ag.game_id, a.level_code;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "agent_id_seq"
OWNED BY "agent"."id";
SELECT setval('"agent_id_seq"', 1000, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "exchange_data_id_seq"
OWNED BY "exchange_data"."id";
SELECT setval('"exchange_data_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "game_users_id_seq"
OWNED BY "game_users"."id";
SELECT setval('"game_users_id_seq"', 1055, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_baccarat_bet_id_seq"
OWNED BY "user_play_log_baccarat"."bet_id";
SELECT setval('"user_play_log_baccarat_bet_id_seq"', 1000000812, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_blackjack_bet_id_seq"
OWNED BY "user_play_log_blackjack"."bet_id";
SELECT setval('"user_play_log_blackjack_bet_id_seq"', 980, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_colordisc_bet_id_seq"
OWNED BY "user_play_log_colordisc"."bet_id";
SELECT setval('"user_play_log_colordisc_bet_id_seq"', 1121, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_fantan_bet_id_seq"
OWNED BY "user_play_log_fantan"."bet_id";
SELECT setval('"user_play_log_fantan_bet_id_seq"', 1098, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_prawncrab_bet_id_seq"
OWNED BY "user_play_log_prawncrab"."bet_id";
SELECT setval('"user_play_log_prawncrab_bet_id_seq"', 507, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "user_play_log_sangong_bet_id_seq"
OWNED BY "user_play_log_sangong"."bet_id";
SELECT setval('"user_play_log_sangong_bet_id_seq"', 2582, true);

-- ----------------------------
-- Primary Key structure for table admin_user
-- ----------------------------
ALTER TABLE "admin_user" ADD CONSTRAINT "admin_user_pkey" PRIMARY KEY ("username");

-- ----------------------------
-- Primary Key structure for table admin_user_action_log
-- ----------------------------
ALTER TABLE "admin_user_action_log" ADD CONSTRAINT "admin_user_action_log_pkey" PRIMARY KEY ("log_time", "username");

-- ----------------------------
-- Uniques structure for table agent
-- ----------------------------
ALTER TABLE "agent" ADD CONSTRAINT "uni_agent_1" UNIQUE ("code");

-- ----------------------------
-- Primary Key structure for table agent
-- ----------------------------
ALTER TABLE "agent" ADD CONSTRAINT "agent_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table agent_game
-- ----------------------------
ALTER TABLE "agent_game" ADD CONSTRAINT "agent_game_pkey" PRIMARY KEY ("agent_id", "game_id");

-- ----------------------------
-- Primary Key structure for table agent_game_room
-- ----------------------------
ALTER TABLE "agent_game_room" ADD CONSTRAINT "agent_game_room_pkey" PRIMARY KEY ("agent_id", "game_room_id");

-- ----------------------------
-- Primary Key structure for table agent_permission
-- ----------------------------
ALTER TABLE "agent_permission" ADD CONSTRAINT "agent_permission_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table agent_wallet
-- ----------------------------
ALTER TABLE "agent_wallet" ADD CONSTRAINT "agent_wallet_pkey" PRIMARY KEY ("agent_id");

-- ----------------------------
-- Primary Key structure for table agent_wallet_ledger
-- ----------------------------
ALTER TABLE "agent_wallet_ledger" ADD CONSTRAINT "agent_wallet_ledger_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table exchange_data
-- ----------------------------
ALTER TABLE "exchange_data" ADD CONSTRAINT "uni_exchange_data_id" UNIQUE ("id");

-- ----------------------------
-- Primary Key structure for table exchange_data
-- ----------------------------
ALTER TABLE "exchange_data" ADD CONSTRAINT "exchange_data_pkey" PRIMARY KEY ("currency");

-- ----------------------------
-- Uniques structure for table game
-- ----------------------------
ALTER TABLE "game" ADD CONSTRAINT "uni_game_code" UNIQUE ("code");

-- ----------------------------
-- Primary Key structure for table game
-- ----------------------------
ALTER TABLE "game" ADD CONSTRAINT "game_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table game_room
-- ----------------------------
ALTER TABLE "game_room" ADD CONSTRAINT "room_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table game_users
-- ----------------------------
ALTER TABLE "game_users" ADD CONSTRAINT "uni_id" UNIQUE ("id");
ALTER TABLE "game_users" ADD CONSTRAINT "uni_usrname" UNIQUE ("username");

-- ----------------------------
-- Primary Key structure for table game_users
-- ----------------------------
ALTER TABLE "game_users" ADD CONSTRAINT "game_users_pkey" PRIMARY KEY ("id", "agent_id", "original_username", "username");

-- ----------------------------
-- Primary Key structure for table job_scheduler
-- ----------------------------
ALTER TABLE "job_scheduler" ADD CONSTRAINT "job_scheduler_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table permission_list
-- ----------------------------
ALTER TABLE "permission_list" ADD CONSTRAINT "uni_feature_code" UNIQUE ("feature_code");

-- ----------------------------
-- Primary Key structure for table permission_list
-- ----------------------------
ALTER TABLE "permission_list" ADD CONSTRAINT "permission_list_pkey" PRIMARY KEY ("feature_code", "api_path");

-- ----------------------------
-- Indexes structure for table play_log_common
-- ----------------------------
CREATE INDEX "idx_play_log_common_1" ON "play_log_common" USING btree (
  "create_time" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table play_log_common
-- ----------------------------
ALTER TABLE "play_log_common" ADD CONSTRAINT "play_log_common_pkey" PRIMARY KEY ("lognumber", "game_id");

-- ----------------------------
-- Primary Key structure for table rp_agent_stat
-- ----------------------------
ALTER TABLE "rp_agent_stat" ADD CONSTRAINT "rp_agent_stat_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");

-- ----------------------------
-- Primary Key structure for table rp_agent_stat_day
-- ----------------------------
ALTER TABLE "rp_agent_stat_day" ADD CONSTRAINT "rp_agent_stat_day_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");

-- ----------------------------
-- Primary Key structure for table rp_agent_stat_hour
-- ----------------------------
ALTER TABLE "rp_agent_stat_hour" ADD CONSTRAINT "rp_agent_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");

-- ----------------------------
-- Primary Key structure for table rp_agent_stat_month
-- ----------------------------
ALTER TABLE "rp_agent_stat_month" ADD CONSTRAINT "rp_agent_stat_month_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");

-- ----------------------------
-- Primary Key structure for table rp_agent_stat_week
-- ----------------------------
ALTER TABLE "rp_agent_stat_week" ADD CONSTRAINT "rp_agent_stat_week_pkey" PRIMARY KEY ("log_time", "agent_id", "level_code");

-- ----------------------------
-- Primary Key structure for table server_info
-- ----------------------------
ALTER TABLE "server_info" ADD CONSTRAINT "server_info_pkey" PRIMARY KEY ("code");

-- ----------------------------
-- Primary Key structure for table user_play_log_baccarat
-- ----------------------------
ALTER TABLE "user_play_log_baccarat" ADD CONSTRAINT "play_log_baccarat_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table user_play_log_blackjack
-- ----------------------------
ALTER TABLE "user_play_log_blackjack" ADD CONSTRAINT "play_log_blackjack_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table user_play_log_colordisc
-- ----------------------------
ALTER TABLE "user_play_log_colordisc" ADD CONSTRAINT "play_log_colordisc_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table user_play_log_fantan
-- ----------------------------
ALTER TABLE "user_play_log_fantan" ADD CONSTRAINT "play_log_fantan_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table user_play_log_prawncrab
-- ----------------------------
ALTER TABLE "user_play_log_prawncrab" ADD CONSTRAINT "play_log_prawncrab_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table user_play_log_sangong
-- ----------------------------
ALTER TABLE "user_play_log_sangong" ADD CONSTRAINT "play_log_sangong_pkey" PRIMARY KEY ("lognumber", "agent_id", "user_id", "game_id", "room_type", "desk_id", "seat_id");

-- ----------------------------
-- Primary Key structure for table wallet_ledger
-- ----------------------------
ALTER TABLE "wallet_ledger" ADD CONSTRAINT "wallet_ledger_pkey" PRIMARY KEY ("id");
