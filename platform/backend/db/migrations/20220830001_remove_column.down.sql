ALTER TABLE "public"."game"
  ADD COLUMN "api" jsonb NOT NULL DEFAULT '{}'::jsonb,
  ADD COLUMN "sup_lang" jsonb NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN "public"."game"."api" IS 'game server api list';
COMMENT ON COLUMN "public"."game"."sup_lang" IS '支援語系';
