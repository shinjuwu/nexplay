CREATE TABLE "public"."auto_risk_control_log" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "agent_id" int4 NOT NULL DEFAULT 0,
  "user_id" int4 NOT NULL DEFAULT 0,
  "username" varchar(100) NOT NULL DEFAULT '',
  "level_code" varchar(128) NOT NULL DEFAULT '',
  "risk_code" int2 NOT NULL DEFAULT 0,
  "create_time" timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT "auto_risk_control_log_pkey" PRIMARY KEY ("id")
)
;

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (100300, '取得自動風控列表(只有開發商（營運商）可使用)', '/api/v1/record/getautoriskcontrolloglist', 't', '後台使用', 't', 0);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100300]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;