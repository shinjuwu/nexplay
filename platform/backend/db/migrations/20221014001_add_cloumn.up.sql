ALTER TABLE "public"."agent" 
  ADD COLUMN "creator" varchar(20) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."agent"."creator" IS '創建此代理的後台帳號username';


-- 資料同步
UPDATE
    agent
SET
    creator = au.username
FROM
    admin_user as "au"
WHERE
    au.is_added = false AND agent.top_agent_id=au.agent_id AND top_agent_id  >0;
		
UPDATE
    agent
SET
    creator = 'admin'
WHERE
    top_agent_id = -1;