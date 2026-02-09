CREATE TABLE "public"."miscellaneous" (
    "key" varchar(32) NOT NULL,
    "data" jsonb NOT NULL DEFAULT '{}'::jsonb,
    CONSTRAINT "miscellaneous_pkey" PRIMARY KEY ("key")
);

INSERT INTO "public"."miscellaneous" ("key", "data")
  VALUES ('GameServerInfo', '{"state":1}');

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100267, '取得遊戲server狀態', '/api/v1/game/getgameserverstate', 't', '後台使用', 't', 0);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100268, '設置遊戲server狀態', '/api/v1/game/setgameserverstate', 't', '後台使用', 't', 2);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100269, '創建更新遊戲相關設定(遊戲server維護中才可以使用)', '/api/v1/game/notifygameserver', 't', '後台使用', 't', 2);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100267, 100268, 100269]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;
