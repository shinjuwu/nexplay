-- 新增level_code欄位
ALTER TABLE "public"."user_play_log_baccarat" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_baccarat"."level_code" IS '代理層級碼';

ALTER TABLE "public"."user_play_log_fantan" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_fantan"."level_code" IS '代理層級碼';

ALTER TABLE "public"."user_play_log_sangong" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_sangong"."level_code" IS '代理層級碼';

ALTER TABLE "public"."user_play_log_blackjack" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_blackjack"."level_code" IS '代理層級碼';

ALTER TABLE "public"."user_play_log_prawncrab" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_prawncrab"."level_code" IS '代理層級碼';

ALTER TABLE "public"."user_play_log_colordisc" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."user_play_log_colordisc"."level_code" IS '代理層級碼';

-- 補資料
UPDATE user_play_log_baccarat AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE user_play_log_fantan AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE user_play_log_sangong AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE user_play_log_blackjack AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE user_play_log_prawncrab AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;

UPDATE user_play_log_colordisc AS up
SET level_code=aa.level_code
FROM agent AS aa
WHERE up.agent_id = aa.id;