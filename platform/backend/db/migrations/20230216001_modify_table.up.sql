DROP TABLE IF EXISTS "public"."agent_game_ratio";
CREATE TABLE "public"."agent_game_ratio" (
  "id" varchar(128) NOT NULL,
  "agent_id" int4 NOT NULL,
  "game_id" int4 NOT NULL,
  "game_type" int4 NOT NULL,
  "room_type" int4 NOT NULL,
  "kill_ratio" numeric(20,4) NOT NULL DEFAULT 0,
  "info" varchar(255) NOT NULL DEFAULT ''::varchar,
  "update_time" timestamptz(6) NOT NULL DEFAULT now(),
  CONSTRAINT "agent_game_ratio_pkey" PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "public"."agent_game_ratio"."id" IS 'pKey 組合';
COMMENT ON COLUMN "public"."agent_game_ratio"."agent_id" IS '代理id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_id" IS '遊戲id';
COMMENT ON COLUMN "public"."agent_game_ratio"."game_type" IS '遊戲類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."room_type" IS '房間類型';
COMMENT ON COLUMN "public"."agent_game_ratio"."kill_ratio" IS '基礎殺率';
COMMENT ON COLUMN "public"."agent_game_ratio"."info" IS '備註';
COMMENT ON COLUMN "public"."agent_game_ratio"."update_time" IS '最後更新時間';
COMMENT ON TABLE "public"."agent_game_ratio" IS '代理遊戲平台機率設定表';