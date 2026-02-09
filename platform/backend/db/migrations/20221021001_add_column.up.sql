ALTER TABLE "public"."agent" 
    ADD COLUMN "ip_whitelist" jsonb NOT NULL DEFAULT '[]'::jsonb;