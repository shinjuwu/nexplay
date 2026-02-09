INSERT INTO "public"."permission_list" ("feature_code", "name", "api_path", "is_enabled", "remark", "is_required","action_type") VALUES
 (100286, '此接口用來取得代理設定機率&遊戲輸贏結果(只有管理可以使用)' , '/api/v1/riskcontrol/getagentincomeratioandgamedata', 't', '後台使用', 't', 0);                  

UPDATE "public"."agent_permission"
SET "permission" = jsonb_set("permission", '{list}', "permission"->'list' || '[100286]'::jsonb, false)
WHERE "agent_id" = -1 AND "account_type" = 1;
