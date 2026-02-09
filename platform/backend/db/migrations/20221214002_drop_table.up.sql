ALTER TABLE "public"."game_ratio" 
  DROP CONSTRAINT "game_ratio_pkey",
  DROP COLUMN "id",
  DROP COLUMN "level_code",
  ADD CONSTRAINT "game_ratio_pkey" PRIMARY KEY ("game_id", "room_type");

-- 同步資料
INSERT INTO game_ratio ("game_id", "room_type", "base_ratio", "up_ratio_limit", "down_ratio_limit", "update_time")
SELECT DISTINCT(g."id"), gr.room_type, 0, 0, 0, now()
FROM game g, game_room gr
WHERE g."id" > 0;


CREATE TABLE "public"."agent_game_ratio" (
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "game_type" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "base_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "up_ratio_limit" numeric(20,4) NOT NULL,
  "down_ratio_limit" numeric(20,4) NOT NULL,
  "info" varchar(255) NOT NULL DEFAULT ''::varchar,
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "agent_game_ratio_pkey" PRIMARY KEY ("agent_id", "game_id", "game_type", "room_type")
)
;
COMMENT ON COLUMN "public"."agent_game_ratio"."agent_id" IS '代理id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_type" IS '遊戲類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."base_ratio" IS '基礎殺率';
COMMENT ON COLUMN "public"."agent_game_ratio"."up_ratio_limit" IS '上水線';
COMMENT ON COLUMN "public"."agent_game_ratio"."down_ratio_limit" IS '下水線';
COMMENT ON COLUMN "public"."agent_game_ratio"."info" IS '備註';
COMMENT ON COLUMN "public"."agent_game_ratio"."update_time" IS '最後更新時間';
COMMENT ON TABLE "public"."agent_game_ratio" IS '代理遊戲平台機率設定表';

-- 同步資料
INSERT INTO agent_game_ratio ("agent_id", "game_id", "game_type", "room_type", "base_ratio", "up_ratio_limit", "down_ratio_limit")
SELECT DISTINCT(a."id"), g."id", g."type", gr.room_type, gra.base_ratio, gra.up_ratio_limit, gra.down_ratio_limit 
FROM agent a, game g, game_room gr, game_ratio gra
WHERE g."id" > 0 AND gra.game_id=g."id" AND gra.room_type=gr.room_type;