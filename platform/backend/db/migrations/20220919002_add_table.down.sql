-- 刪除番攤、色碟的遊戲資料 --
DELETE FROM "public"."game" WHERE "id" IN (1002,1003);

-- 刪除番攤玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_fantan";

-- 刪除色碟玩家遊戲紀錄表 --
DROP TABLE "public"."user_play_log_colordisc";
