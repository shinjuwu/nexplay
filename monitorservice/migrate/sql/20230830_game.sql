-- +migrate Up
CREATE TABLE "public"."game" (
  "id" varchar(6) NOT NULL,
  "list" jsonb NOT NULL,
  CONSTRAINT "game_pkey" PRIMARY KEY ("id")
)
;

COMMENT ON TABLE "public"."game" IS '遊戲列表';


INSERT INTO "public"."game" ("id", "list") VALUES ('dev', '[{"id": 1008, "code": "rocket", "name": "火箭", "type": 1}, {"id": 2008, "code": "catte", "name": "越南Catte", "type": 2}, {"id": 4001, "code": "fruit777slot", "name": "水果777", "type": 4}, {"id": 1005, "code": "hundredsicbo", "name": "百人骰寶", "type": 1}, {"id": 2004, "code": "texas", "name": "德州撲克", "type": 2}, {"id": 2001, "code": "blackjack", "name": "21點", "type": 2}, {"id": 2002, "code": "sangong", "name": "三公", "type": 2}, {"id": 2003, "code": "bullbull", "name": "搶庄牛牛", "type": 2}, {"id": 1001, "code": "baccarat", "name": "百家樂", "type": 1}, {"id": 1004, "code": "prawncrab", "name": "魚蝦蟹", "type": 1}, {"id": 1002, "code": "fantan", "name": "翻攤", "type": 1}, {"id": 1003, "code": "colordisc", "name": "色碟", "type": 1}, {"id": 3001, "code": "fruitslot", "name": "水果機", "type": 3}, {"id": 3002, "code": "rcfishing", "name": "三國捕魚", "type": 3}, {"id": 1006, "code": "cockfight", "name": "鬥雞", "type": 1}, {"id": 1007, "code": "dogracing", "name": "賽狗", "type": 1}, {"id": 2005, "code": "rummy", "name": "拉密", "type": 2}, {"id": 2006, "code": "goldenflower", "name": "炸金花", "type": 2}, {"id": 2007, "code": "pokdeng", "name": "泰式博丁", "type": 2}]');

-- +migrate Down
DROP TABLE IF EXISTS
    game;