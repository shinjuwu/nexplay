CREATE TABLE "public"."agent_game_icon_list" (
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_admin" bool NOT NULL DEFAULT false,
  "is_default" bool NOT NULL DEFAULT true,
  "icon_list" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "update_time" timestamptz NOT NULL DEFAULT now(),
  "create_time" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("agent_id", "level_code")
)
;

COMMENT ON COLUMN "public"."agent_game_icon_list"."is_admin" IS '是否為總管理設定(預設值)';
COMMENT ON COLUMN "public"."agent_game_icon_list"."is_default" IS '是否使用總管理設定(預設值)';
COMMENT ON COLUMN "public"."agent_game_icon_list"."icon_list" IS '自定義icon list';


INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type")
VALUES
  (100290, '取得遊戲icon list', '/api/v1/game/getgameiconlist', 't', '後台使用', 't', 0),
  (100291, '設置遊戲icon list', '/api/v1/game/setgameiconlist', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100290, 100291]'::jsonb, false);