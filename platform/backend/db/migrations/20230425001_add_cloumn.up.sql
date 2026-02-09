ALTER TABLE "public"."agent" 
  ADD COLUMN "api_ip_whitelist" jsonb NOT NULL DEFAULT '[]'::jsonb;

INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (	100292	,	'取得代理API IP資訊'	,	'/api/v1/agent/getagentapiipwhitelist'	,	't'	,	'後台使用'	,	't'	,	0	);
INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES (	100293	,	'設置代理API IP資訊'	,	'/api/v1/agent/setagentapiipwhitelist'	,	't'	,	'後台使用'	,	't'	,	2	);

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100292, 100293]'::jsonb, false);