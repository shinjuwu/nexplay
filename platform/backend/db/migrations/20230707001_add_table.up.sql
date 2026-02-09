CREATE TABLE "public"."agent_game_users_stat_min" (
  "log_time" varchar(18) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "agent_id" int4 NOT NULL DEFAULT -1,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "online_game_user_count" int4 NOT NULL DEFAULT 0,
  "create_time" timestamptz NOT NULL DEFAULT now(),
  "update_time" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("log_time", "agent_id", "level_code")
);

COMMENT ON COLUMN "public"."agent_game_users_stat_min"."log_time" IS '紀錄時間';
COMMENT ON COLUMN "public"."agent_game_users_stat_min"."agent_id" IS '代理編號';
COMMENT ON COLUMN "public"."agent_game_users_stat_min"."level_code" IS '層級碼';
COMMENT ON COLUMN "public"."agent_game_users_stat_min"."online_game_user_count" IS '線上玩家人數';
COMMENT ON COLUMN "public"."agent_game_users_stat_min"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."agent_game_users_stat_min"."update_time" IS '資料更新時間';
COMMENT ON TABLE "public"."agent_game_users_stat_min" IS '代理玩家每分鐘統計表';


INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100306, '此接口用來取得今日昨日的各時段線上人數', '/api/v1/manage/getintervalrealtimeuserdata', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100306]'::jsonb, false)
WHERE "agent_id" = -1;