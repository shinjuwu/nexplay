CREATE FUNCTION "public"."udf_create_admin_user"("_agent_id" int4, "_username" varchar,
    "_password" varchar, "_nickname" varchar, "_account_type" int4, "_is_readonly" int4,
    "_is_added" bool, "_role" uuid, "_info" varchar)
    RETURNS json AS $$
DECLARE
    "ret_agent_id" int4;
    "ret_username" varchar;
    "ret_password" varchar;
    "ret_nickname" varchar;
    "ret_account_type" int4;
    "ret_allow_ip" varchar;
    "ret_is_readonly" int4;
    "ret_is_enabled" int4;
    "ret_update_time" timestamp;
    "ret_create_time" timestamp;
    "ret_is_added" bool;
    "ret_login_time" timestamp;
    "ret_role" uuid;
    "ret_info" varchar;
BEGIN
    INSERT INTO "public"."admin_user" ("agent_id", "username", "password", "nickname", "account_type",
        "is_readonly", "is_added", "role", "info")
        VALUES ("_agent_id", "_username", "_password", "_nickname", "_account_type", "_is_readonly",
            "_is_added", "_role", "_info")
        RETURNING "agent_id", "username", "password", "nickname", "account_type", "allow_ip",
            "is_readonly", "is_enabled", "update_time", "create_time", "is_added", "login_time",
            "role", "info" INTO "ret_agent_id", "ret_username", "ret_password", "ret_nickname",
            "ret_account_type", "ret_allow_ip", "ret_is_readonly", "ret_is_enabled", 
            "ret_update_time", "ret_create_time", "ret_is_added", "ret_login_time", "ret_role",
            "ret_info";

    RETURN json_build_object(
        'agent_id', "ret_agent_id",
        'username', "ret_username",
        'password', "ret_password",
        'nickname', "ret_nickname",
        'account_type', "ret_account_type",
        'allow_ip', "ret_allow_ip",
        'is_readonly', "ret_is_readonly",
        'is_enabled', "ret_is_enabled",
        'update_time', extract(epoch from "ret_update_time") * 1000000,
        'create_time', extract(epoch from "ret_create_time") * 1000000,
        'is_added', "ret_is_added",
        'login_time', extract(epoch from "ret_login_time") * 1000000,
        'role', "ret_role",
        'info', "ret_info"
    );
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION "public"."udf_update_admin_user"("_agent_id" int4, "_username" varchar,
	"_role" uuid, "_info" varchar, "_is_enabled" int4)
	RETURNS json AS $$
DECLARE
    "ret_update_time" timestamp;
BEGIN
    UPDATE "public"."admin_user"
        SET "role" = "_role",
            "info" = "_info",
            "is_enabled" = "_is_enabled",
            "update_time" = now()
        WHERE "agent_id" = "_agent_id" AND "username" = "_username"
        RETURNING "update_time" INTO "ret_update_time";

    RETURN json_build_object(
        'update_time', extract(epoch from "ret_update_time") * 1000000
    );
END;
$$ LANGUAGE plpgsql;
