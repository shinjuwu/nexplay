ALTER TABLE "public"."game" 
  ADD COLUMN "cal_state" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."game"."state" IS '遊戲狀態(1:open,2:close)';

COMMENT ON COLUMN "public"."game"."cal_state" IS '是否開啟計算報表(1:open,2:close)';

-- update all
UPDATE game SET cal_state=1 WHERE 1=1;