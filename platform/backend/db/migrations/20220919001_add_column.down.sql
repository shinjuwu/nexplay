-- 遊戲局紀錄刪除抽水欄位 --
ALTER TABLE "public"."play_log_common"
  DROP COLUMN "tax";

-- 百家樂遊戲紀錄刪除抽水欄位 --
ALTER TABLE "public"."user_play_log_baccarat"
  DROP COLUMN "tax";