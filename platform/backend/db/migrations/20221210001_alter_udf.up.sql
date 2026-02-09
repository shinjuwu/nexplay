DROP FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar);

CREATE OR REPLACE FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar)
  RETURNS "pg_catalog"."json" AS $BODY$
	DECLARE
	_user_id int4;
	_username varchar;
 	_is_new bool;
BEGIN

	SELECT 
		"id", "username", false  into _user_id, _username, _is_new
	FROM
		"public"."game_users" 
	WHERE
		agent_id = _agent_id AND original_username = _original_username;

	IF NOT FOUND THEN

		INSERT INTO "public"."game_users" 
			( "agent_id", "original_username", "username", "user_metadata", "sum_coin_in", "sum_coin_out", "level_code") 
		SELECT
			_agent_id, _original_username, _trans_username, '{}', 0, _coin, _level_code
		RETURNING "id", "username", true into _user_id, _username, _is_new
		;
		
		UPDATE agent SET member_count = member_count+1, update_time = now() WHERE "id" = _agent_id
		;
	END IF;
	
	RETURN json_build_object(
	'id', _user_id,
	'username', _username,
	'is_new', _is_new
	);

END
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;