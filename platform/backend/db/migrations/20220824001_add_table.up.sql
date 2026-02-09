ALTER TABLE "public"."agent" DROP constraint IF EXISTS "agent_fkey_top_agent_id";
ALTER TABLE "public"."agent_game" DROP constraint IF EXISTS "agent_id";
ALTER TABLE "public"."agent_game" DROP constraint IF EXISTS "game_id";
ALTER TABLE "public"."agent_game_room" DROP constraint IF EXISTS "agent_game_room_fkey_agent_id";
ALTER TABLE "public"."agent_game_room" DROP constraint IF EXISTS "agent_game_room_fkey_game_id";
ALTER TABLE "public"."agent_game_room" DROP constraint IF EXISTS "agent_game_room_fkey_game_room_id";
ALTER TABLE "public"."game_room" DROP constraint IF EXISTS "room_fkey_game_id";

-- ----------------------------
-- Table structure for server_info
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."server_info";
CREATE TABLE "public"."server_info" (
  "code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ip" varchar(100) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "addresses" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "info" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_enabled" bool NOT NULL DEFAULT false,
  "create_time" timestamptz(6) NOT NULL DEFAULT now(),
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "server_info_pkey" PRIMARY KEY ("code")
)
;
COMMENT ON COLUMN "public"."server_info"."code" IS '唯一代碼';
COMMENT ON COLUMN "public"."server_info"."ip" IS 'server address';
COMMENT ON COLUMN "public"."server_info"."addresses" IS 'connect addresses';
COMMENT ON COLUMN "public"."server_info"."info" IS '備註';
COMMENT ON COLUMN "public"."server_info"."is_enabled" IS '開關';
COMMENT ON TABLE "public"."server_info" IS '遊戲伺服器連線位置紀錄表';

-- ----------------------------
-- Records of server_info
-- ----------------------------
INSERT INTO "public"."server_info" ("code", "info", "is_enabled", "create_time", "update_time", "addresses", "ip") VALUES ('dev', '', 't', '2022-08-24 02:57:02.390628+00', '2022-08-24 02:57:02.390628+00', '{"notification": "http://172.30.0.154:9642/"}', '172.30.0.154');
INSERT INTO "public"."server_info" ("code", "info", "is_enabled", "create_time", "update_time", "addresses", "ip") VALUES ('qa', '', 'f', '2022-08-24 03:10:47.459456+00', '2022-08-24 03:10:47.459456+00', '{"notification": "http://172.30.0.164:9642/"}', '172.30.0.164');

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS "public"."game";
CREATE TABLE "public"."game" (
  "id" int4 NOT NULL,
  "server_info_code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "code" varchar(40) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "state" int2 NOT NULL DEFAULT 1,
  "image" varchar COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "api" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "sup_lang" jsonb NOT NULL DEFAULT '{}'::jsonb,
  "h5_link" varchar(500) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "create_time" timestamp(6) NOT NULL DEFAULT now(),
  "update_time" timestamp(6) NOT NULL DEFAULT now(),
  CONSTRAINT "uni_game_code" UNIQUE ("code"),
  CONSTRAINT "game_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."game"."id" IS '遊戲id(PK)';
COMMENT ON COLUMN "public"."game"."server_info_code" IS '遊戲的server code';
COMMENT ON COLUMN "public"."game"."name" IS '遊戲名稱';
COMMENT ON COLUMN "public"."game"."code" IS '遊戲識別編碼';
COMMENT ON COLUMN "public"."game"."state" IS '遊戲狀態';
COMMENT ON COLUMN "public"."game"."image" IS '遊戲圖片base64';
COMMENT ON COLUMN "public"."game"."api" IS 'game server api list';
COMMENT ON COLUMN "public"."game"."sup_lang" IS '支援語系';
COMMENT ON COLUMN "public"."game"."h5_link" IS '前端遊戲位置';
COMMENT ON COLUMN "public"."game"."create_time" IS '創建時間';
COMMENT ON COLUMN "public"."game"."update_time" IS '更新時間';
COMMENT ON TABLE "public"."game" IS '遊戲列表';

-- ----------------------------
-- Records of game
-- ----------------------------
INSERT INTO "public"."game" VALUES (0, 'dev', '大廳', 'lobby', 1, '', '{}', '{}', 'http://172.30.0.152/dev/lobby/', '2022-08-11 13:58:05.831006', '2022-08-11 13:58:05.831006');
INSERT INTO "public"."game" VALUES (1001, 'dev', '百家樂', 'baccarat', 1, '', '{}', '{}', 'http://172.30.0.152/dev/lobby/', '2022-08-10 03:23:50.310687', '2022-08-10 03:23:50.310687');
INSERT INTO "public"."game" VALUES (1002, 'dev', '測試用馬娘', 'ma', 1, '', '{}', '{}', 'http://172.30.0.152/dev/ma/', '2022-08-17 14:54:26.68037', '2022-08-17 14:54:26.68037');
