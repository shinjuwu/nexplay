ALTER TABLE "public"."game_users" 
  ADD COLUMN "kill_dive_state" int2 NOT NULL DEFAULT 0;

COMMENT ON COLUMN "public"."game_users"."kill_dive_state" IS '殺放設定狀態(一般玩家:0、定點玩家:1、黑名單玩家:2)';

DROP FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar);
CREATE OR REPLACE FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar, "_user_metadata_default" jsonb)
  RETURNS "pg_catalog"."json" AS $BODY$
	DECLARE
	_user_id int4;
	_username varchar;
	_user_metadata jsonb;
 	_is_new bool;
	_is_enabled bool;
	_kill_dive_state int2;
	_risk_control_status varchar;
BEGIN

	SELECT 
		"id", "username", is_enabled, false, "user_metadata", "kill_dive_state", "risk_control_status" into _user_id, _username, _is_enabled, _is_new, _user_metadata, _kill_dive_state, _risk_control_status
	FROM
		"public"."game_users" 
	WHERE
		agent_id = _agent_id AND original_username = _original_username;
		
	IF NOT FOUND THEN

		INSERT INTO "public"."game_users" 
			( "agent_id", "original_username", "username", "user_metadata", "sum_coin_in", "sum_coin_out", "level_code") 
		SELECT
			_agent_id, _original_username, _trans_username, _user_metadata_default, 0, _coin, _level_code
		RETURNING "id", "username", "is_enabled", true, "kill_dive_state", "risk_control_status" into _user_id, _username, _is_enabled, _is_new, _kill_dive_state, _risk_control_status
		;
		
		UPDATE agent SET member_count = member_count+1, update_time = now() WHERE "id" = _agent_id
		;
		
		_user_metadata = _user_metadata_default;
	END IF;
	
	IF _user_metadata = '{}'::jsonb THEN
		_user_metadata = _user_metadata_default;
	END IF;
	
	
	
	RETURN json_build_object(
	'id', _user_id,
	'username', _username,
	'is_new', _is_new,
	'is_enabled', _is_enabled,
	'user_metadata', _user_metadata,
	'kill_dive_state', _kill_dive_state,
	'risk_control_status', _risk_control_status
	);

END
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;