ALTER TABLE "public"."game_users" RENAME COLUMN "temporary_coin" TO "sum_coin_in";

ALTER TABLE "public"."game_users" RENAME COLUMN "coin" TO "sum_coin_out";

COMMENT ON COLUMN "public"."game_users"."sum_coin_in" IS '累積遊戲幣轉入';

COMMENT ON COLUMN "public"."game_users"."sum_coin_out" IS '累積遊戲幣轉出';

UPDATE game_users SET sum_coin_in=0, sum_coin_out=0;

ALTER TABLE "public"."game_users" 
  ADD COLUMN "level_code" varchar(128) NOT NULL DEFAULT ''::character varying;

COMMENT ON COLUMN "public"."game_users"."level_code" IS '代理層級碼';

-- update game_user level_code from agent
UPDATE game_users
SET level_code= aa.level_code
FROM (SELECT "id", level_code FROM agent) as aa
WHERE agent_id = aa."id";

ALTER TABLE "public"."agent" 
  ADD COLUMN "member_count" int4 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."agent"."member_count" IS '會員人數';

-- update agent member_count from game_users
UPDATE agent
SET member_count= count
FROM (SELECT agent_id, COUNT(*) as "count" FROM game_users GROUP BY agent_id) as gs
WHERE "id" = gs.agent_id;

CREATE OR REPLACE FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar)
  RETURNS TABLE("user_id" int4, "is_new" bool) AS $BODY$
	DECLARE
	_user_id int4;
 	_is_new bool;
BEGIN

	SELECT 
		"id", false  into _user_id, _is_new
	FROM
		"public"."game_users" 
	WHERE
		agent_id = _agent_id AND original_username = _original_username;

	IF NOT FOUND THEN

		INSERT INTO "public"."game_users" 
			( "agent_id", "original_username", "username", "user_metadata", "temporary_coin", "coin") 
		SELECT
			_agent_id, _original_username, _trans_username, '{}', 0, _coin
		RETURNING "id", true into _user_id, _is_new
		;
		
		UPDATE agent SET member_count = member_count+1, update_time = now() WHERE "id" = _agent_id
		;
	END IF;
	
	RETURN QUERY
		SELECT _user_id ,_is_new ;

END
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
  ROWS 1
;