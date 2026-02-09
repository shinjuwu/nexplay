ALTER TABLE "public"."wallet_ledger" 
  ADD COLUMN "agent_id" int4 NOT NULL DEFAULT 0,
  ADD COLUMN "agent_code" varchar(8) NOT NULL DEFAULT '',
  ADD COLUMN "kind" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."wallet_ledger"."agent_id" IS 'mapping id of agent';

COMMENT ON COLUMN "public"."wallet_ledger"."agent_code" IS 'mapping code of agent';

COMMENT ON COLUMN "public"."wallet_ledger"."kind" IS '上下分類型(1:上分,2:下分)';


ALTER TABLE "public"."wallet_ledger" 
  ALTER COLUMN "username" TYPE varchar(100) COLLATE "pg_catalog"."default";