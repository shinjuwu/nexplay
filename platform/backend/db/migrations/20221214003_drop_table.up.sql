DROP TABLE IF EXISTS "public"."game_ratio";
CREATE TABLE "public"."game_ratio" (
  "game_id" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "kill_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "dive_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "up_ratio_limit" numeric(20,4) NOT NULL DEFAULT 0,
  "down_ratio_limit" numeric(20,4) NOT NULL DEFAULT 0,
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "game_ratio_pkey" PRIMARY KEY ("game_id", "room_type")
)
;
COMMENT ON COLUMN "public"."game_ratio"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."game_ratio"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."game_ratio"."kill_ratio" IS '基礎殺率';
COMMENT ON COLUMN "public"."game_ratio"."dive_ratio" IS '基礎放水率';
COMMENT ON COLUMN "public"."game_ratio"."up_ratio_limit" IS '上水線';
COMMENT ON COLUMN "public"."game_ratio"."down_ratio_limit" IS '下水線';
COMMENT ON COLUMN "public"."game_ratio"."update_time" IS '最後更新間';
COMMENT ON TABLE "public"."game_ratio" IS '遊戲平台機率設定表';

-- 同步資料
INSERT INTO game_ratio ("game_id", "room_type", "kill_ratio", "dive_ratio", "up_ratio_limit", "down_ratio_limit", "update_time")
SELECT DISTINCT(g."id"), gr.room_type, 0, 0, 0, 0, now()
FROM game g, game_room gr
WHERE g."id" > 0;

DROP TABLE IF EXISTS "public"."agent_game_ratio";
CREATE TABLE "public"."agent_game_ratio" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "game_type" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "kill_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "dive_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "up_ratio_limit" numeric(20,4) NOT NULL DEFAULT 0,
  "down_ratio_limit" numeric(20,4) NOT NULL DEFAULT 0,
  "info" varchar(255) NOT NULL DEFAULT ''::varchar,
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "agent_game_ratio_pkey" PRIMARY KEY ("agent_id", "game_id", "game_type", "room_type")
)
;
COMMENT ON COLUMN "public"."agent_game_ratio"."agent_id" IS '代理id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_type" IS '遊戲類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."kill_ratio" IS '基礎殺率';
COMMENT ON COLUMN "public"."agent_game_ratio"."dive_ratio" IS '基礎放水率';
COMMENT ON COLUMN "public"."agent_game_ratio"."up_ratio_limit" IS '上水線';
COMMENT ON COLUMN "public"."agent_game_ratio"."down_ratio_limit" IS '下水線';
COMMENT ON COLUMN "public"."agent_game_ratio"."info" IS '備註';
COMMENT ON COLUMN "public"."agent_game_ratio"."update_time" IS '最後更新時間';
COMMENT ON TABLE "public"."agent_game_ratio" IS '代理遊戲平台機率設定表';

-- 同步資料
INSERT INTO agent_game_ratio ("agent_id", "game_id", "game_type", "room_type", "kill_ratio", "dive_ratio", "up_ratio_limit", "down_ratio_limit")
SELECT DISTINCT(a."id"), g."id", g."type", gr.room_type, gra.kill_ratio, gra.dive_ratio, gra.up_ratio_limit, gra.down_ratio_limit 
FROM agent a, game g, game_room gr, game_ratio gra
WHERE g."id" > 0 AND gra.game_id=g."id" AND gra.room_type=gr.room_type;


UPDATE "public"."game" SET "type" = 2 WHERE "id" IN ( 2001,2002,2003);
UPDATE "public"."game" SET "type" = 1 WHERE "id" IN ( 1001,1002,1003,1004,1005);
