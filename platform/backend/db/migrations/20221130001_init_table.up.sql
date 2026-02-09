CREATE TABLE "public"."rt_game_stat" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "de_score" numeric(20,4) NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "rt_game_stat_pkey" PRIMARY KEY ("agent_id", "game_id")
)
;
COMMENT ON COLUMN "public"."rt_game_stat"."ya_score" IS '總投注(投注)';
COMMENT ON COLUMN "public"."rt_game_stat"."vaild_ya_score" IS '總有效投注';
COMMENT ON COLUMN "public"."rt_game_stat"."de_score" IS '總派獎(玩家得分)';
COMMENT ON COLUMN "public"."rt_game_stat"."bonus" IS '總紅利';
COMMENT ON COLUMN "public"."rt_game_stat"."tax" IS '總抽水';
COMMENT ON TABLE "public"."rt_game_stat" IS '遊戲即時統計資料(realtime game stat)';

CREATE TABLE "public"."rt_game_stat_hour" (
  "log_time" varchar(12) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "vaild_ya_score" numeric(20,4) NOT NULL DEFAULT 0,
  "de_score" numeric(20,4) NOT NULL DEFAULT 0,
  "bonus" numeric(20,4) NOT NULL DEFAULT 0,
  "tax" numeric(20,4) NOT NULL DEFAULT 0,
  CONSTRAINT "rt_game_stat_hour_pkey" PRIMARY KEY ("log_time", "agent_id", "game_id")
)
;
COMMENT ON COLUMN "public"."rt_game_stat_hour"."log_time" IS 'YYYYMMDDhhmm';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."ya_score" IS '總投注(投注)';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."vaild_ya_score" IS '總有效投注';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."de_score" IS '總派獎(玩家得分)';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."bonus" IS '總紅利';
COMMENT ON COLUMN "public"."rt_game_stat_hour"."tax" IS '總抽水';
COMMENT ON TABLE "public"."rt_game_stat" IS '遊戲即時統計資料(realtime game stat)';


INSERT INTO "public"."job_scheduler" ("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date", "update_time") VALUES ('cb7b30fb-9f2c-4137-922d-c5cda719e280', '0 1 */1 * * *', '每小時的1分執行', 'job_rt_game_stat_hour', 't', 0, '', '2022-11-30 06:08:29+00');

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100259, '此接口用來取得chat service 連線資訊', '/api/v1/notify/getchatserviceconnInfo', 't', '後台使用', '2022-11-29 07:41:47.722619+00', '2022-11-29 07:41:47.722619+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100260, '此接口用來取得當天資訊總覽的資料', '/api/v1/manage/getstatdata', 't', '後台使用', '2022-11-30 09:15:53.772723+00', '2022-11-30 09:15:53.772723+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100261, '此接口用來取得今日風險玩家清單(前100名)', '/api/v1/manage/getriskuserlist', 't', '後台使用', '2022-11-30 09:16:22.089192+00', '2022-11-30 09:16:22.089192+00', 't');
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "create_time", "update_time", "is_required") VALUES (100262, '此接口用來取得今日遊戲輸贏排行榜', '/api/v1/manage/getgameleaderboards', 't', '後台使用', '2022-11-30 09:16:46.454451+00', '2022-11-30 09:16:46.454451+00', 't');

-- 可操作權限 account_type 開發者:1, 總代理:2, 子代理:3
UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100214,100215,100225,100226,100227,100228,100229,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100255,100256,100257,100258,100259,100260,100261,100262]}'
WHERE
  "account_type" = 1;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100200,100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100250,100251,100252,100253,100254,100255,100256,100257,100258,100259,100260,100261,100262]}'
WHERE
  "account_type" = 2;

UPDATE
  "public"."agent_permission"
SET
  "permission" = '{"list":[100201,100202,100203,100204,100205,100206,100207,100208,100209,100210,100211,100212,100213,100230,100231,100232,100233,100236,100237,100238,100239,100240,100241,100242,100245,100246,100247,100248,100249,100252,100253,100254,100255,100256,100257,100258,100259,100260,100261,100262]}'
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