DROP FUNCTION "public"."udf_create_agent_permission";

DROP FUNCTION "public"."udf_update_agent_permission";

DROP PROCEDURE "public"."usp_update_agents_permission";

DROP FUNCTION "public"."udf_update_agent";
CREATE OR REPLACE FUNCTION "public"."udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4)
  RETURNS json AS $$
DECLARE
  "ret_agent_update_time" timestamp;
  "ret_admin_user" json;
BEGIN
  UPDATE "public"."agent"
    SET "name" = "_agent_name",
      "info" = "_agent_info",
      "commission" = "_agent_commission",
      "update_time" = now()
    WHERE "id" = "_agent_id"
    RETURNING "update_time" INTO "ret_agent_update_time";

  SELECT "public"."udf_update_admin_user"("_agent_id", "_admin_user_username", "_admin_user_role",
    "_admin_user_info", "_admin_user_is_enabled") INTO "ret_admin_user";

  RETURN json_build_object(
    'agent', json_build_object(
      'update_time', extract(epoch from "ret_agent_update_time") * 1000000
    ),
    'admin_user', "ret_admin_user"
  );
END;
$$ LANGUAGE plpgsql;