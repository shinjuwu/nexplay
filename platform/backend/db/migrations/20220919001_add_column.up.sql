-- 遊戲局紀錄新增抽水欄位 --
ALTER TABLE "public"."play_log_common"
  ADD COLUMN "tax" numeric(20,4) NOT NULL DEFAULT(0);

COMMENT ON COLUMN "public"."play_log_common"."tax" IS '抽水';


-- 百家樂遊戲紀錄新增抽水欄位 --
ALTER TABLE "public"."user_play_log_baccarat"
  ADD COLUMN "tax" numeric(20,4) NOT NULL DEFAULT(0);

COMMENT ON COLUMN "public"."user_play_log_baccarat"."tax" IS '抽水';
