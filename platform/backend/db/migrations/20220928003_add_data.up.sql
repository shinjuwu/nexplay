ALTER TABLE "public"."wallet_ledger" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."wallet_ledger"."level_code" IS '層級碼';

-- 新增預設總/子代理
INSERT INTO "public"."agent" ("id", "name", "code", "secret_key", "md5_key", "commission", "info", "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id", "cooperation", "coin_limit", "coin_use", "coin_supply_setting", "level_code", "aes_key") 
VALUES (3, 'test1', '52078174', '5dfdb4ea61c520fd76847d764a510fa0', 'f7c637cd39679e99', 0, '總代理', 1, '1970-01-01 00:00:00', '2022-09-14 01:40:35.446957', '2022-09-14 01:40:35.446957', 'f', 1, 1, '0.0000', '0.0000', '{}', '00010003', 'ddxbst648uf7hdbc');
INSERT INTO "public"."agent" ("id", "name", "code", "secret_key", "md5_key", "commission", "info", "is_enabled", "disable_time", "update_time", "create_time", "is_top_agent", "top_agent_id", "cooperation", "coin_limit", "coin_use", "coin_supply_setting", "level_code", "aes_key") 
VALUES (4, 'test2', '8d366578', '2daa5dc3a4697aa599be104cda1e255e', '17b11ea73bd9e4a6', 0, '子代理', 1, '1970-01-01 00:00:00', '2022-09-14 05:43:33.753453', '2022-09-14 05:43:33.753453', 'f', 3, 1, '0.0000', '0.0000', '{}', '000100030004', 'js73hf6sgyehvg3s');

-- 新增預設總/子代理後臺帳號
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time", "permission", "is_added") 
VALUES (3, 'test1', 'XaQEuBzacH2SojHn5cqhxjo62bW2K4Z1N1Rjj4dGjZRerQ', '', 'f', '', '0.0.0.0/0', 2, 0, 1, '2022-09-14 01:40:37.442888', '2022-09-14 01:40:37.442888', '{"list": [100200, 100201, 100202, 100203, 100204, 100205, 100208, 100209, 100210, 100211, 100212, 100213, 100214, 100215, 100216, 100217, 100218, 100219, 100220, 100221, 200101, 200102]}', 'f');
INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "google_auth", "google_key", "allow_ip", "account_type", "is_readonly", "is_enabled", "update_time", "create_time", "permission", "is_added") 
VALUES (4, 'test2', 'wxCWBc9W1v3zxmTCtXiKJjo69LddLtVDsiah00sNI4exEA', '', 'f', '', '0.0.0.0/0', 3, 0, 1, '2022-09-14 05:43:33.761367', '2022-09-14 05:43:33.761367', '{"list": [100200, 100201, 100202, 100203, 100204, 100205, 100208, 100209, 100210, 100211, 100212, 100213, 100214, 100215, 100216, 100217, 100218, 100219, 100220, 100221, 200101, 200102]}', 'f');

-- 將所有玩家都歸屬到預設總代理之下
UPDATE game_users
SET agent_id = 3;

-- 修改所有遊戲紀錄的代理id
UPDATE user_play_log_fantan
SET agent_id = 3;

UPDATE user_play_log_baccarat
SET agent_id = 3;

UPDATE user_play_log_colordisc
SET agent_id = 3;

-- 修改帳變紀錄
UPDATE wallet_ledger
SET agent_id = 3, agent_code = '52078174', level_code = '00010003';