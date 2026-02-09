CREATE FUNCTION "public"."udf_create_agent_permission" ("_agent_id" int4, "_name" varchar, "_info" varchar, "_account_type" int2, "_permission" jsonb, "_check_permission" boolean)
  RETURNS json AS $$
DECLARE
  "tmp_permission" jsonb;
  "tmp_list" jsonb;
  "ret_id" uuid;
BEGIN
  -- 尋找root account type所擁有的最大權限 --
  SELECT "permission" INTO "tmp_permission"
    FROM "public"."agent_permission"
    WHERE "agent_id" = -1 AND "account_type" = "_account_type";

  -- 過濾root沒有的權限 --
  SELECT jsonb_agg("r"."elements") INTO "tmp_list"
    FROM (
      SELECT jsonb_array_elements("tmp_permission"->'list') AS "elements"
      INTERSECT
      SELECT jsonb_array_elements("_permission"->'list') AS "elements"
    ) AS "r";
  
  IF "_check_permission" THEN
    -- 尋找agent所擁有的最大權限 --
    SELECT "ap"."permission" INTO "tmp_permission"
      FROM "public"."agent_permission" AS "ap"
      INNER JOIN "public"."admin_user" AS "au" ON "au"."role" = "ap"."id"
      WHERE "au"."agent_id" = "_agent_id" AND "is_added" = false;

    -- 過濾沒有的權限 --
    SELECT jsonb_agg("r"."elements") INTO "tmp_list"
      FROM (
        SELECT jsonb_array_elements("tmp_permission"->'list') AS "elements"
        INTERSECT
        SELECT jsonb_array_elements("tmp_list") AS "elements"
      ) AS "r";

    -- 更新權限 --
    "_permission" = jsonb_set("_permission", '{list}', "tmp_list", false);
  END IF;

  INSERT INTO "public"."agent_permission" ("agent_id", "name", "info", "account_type", "permission")
    VALUES ("_agent_id", "_name", "_info", "_account_type", "_permission")
    RETURNING "id" INTO "ret_id";

  RETURN json_build_object(
    'id', "ret_id",
    'permission', "_permission"
  );
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION "public"."udf_update_agent_permission" ("_id" uuid, "_agent_id" int4, "_name" varchar, "_info" varchar, "_account_type" int2, "_permission" jsonb, "_check_permission" boolean, "_update_child_agent_permission" boolean)
  RETURNS json AS $$
DECLARE
  "tmp_permission" jsonb;
  "tmp_list" jsonb;
  "ret_id" uuid;
  "ret_permission" jsonb;
BEGIN
  -- 尋找root account type所擁有的最大權限 --
  SELECT "permission" INTO "tmp_permission"
    FROM "public"."agent_permission"
    WHERE "agent_id" = -1 AND "account_type" = "_account_type";

  -- 過濾root沒有的權限 --
  SELECT jsonb_agg("r"."elements") INTO "tmp_list"
    FROM (
      SELECT jsonb_array_elements("tmp_permission"->'list') AS "elements"
      INTERSECT
      SELECT jsonb_array_elements("_permission"->'list') AS "elements"
    ) AS "r";
  
  IF "_check_permission" THEN
    -- 尋找agent所擁有的最大權限 --
    SELECT "ap"."permission" INTO "tmp_permission"
      FROM "public"."agent_permission" AS "ap"
      INNER JOIN "public"."admin_user" AS "au" ON "au"."role" = "ap"."id"
      WHERE "au"."agent_id" = "_agent_id" AND "is_added" = false;

    -- 過濾沒有的權限 --
    SELECT jsonb_agg("r"."elements") INTO "tmp_list"
      FROM (
        SELECT jsonb_array_elements("tmp_permission"->'list') AS "elements"
        INTERSECT
        SELECT jsonb_array_elements("_permission"->'list') AS "elements"
      ) AS "r";

    -- 更新權限 --
    "_permission" = jsonb_set("_permission", '{list}', "tmp_list", false);
  END IF;

  UPDATE "public"."agent_permission"
    SET "name" = "_name",
      "info" = "_info",
      "account_type" = "_account_type",
      "permission" = "_permission"
    WHERE "id" = "_id";

  IF "_update_child_agent_permission" THEN
    CALL "public"."usp_update_agents_permission" ("_id", '');
  END IF;

  RETURN json_build_object(
    'permission', "_permission"
  );
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE "public"."usp_update_agents_permission" ("_agent_permission_id" uuid, "_level_code" varchar)
  LANGUAGE plpgsql AS
$$
BEGIN
  WITH RECURSIVE "find_agents" AS (
	  SELECT "a"."id"
	    FROM "public"."agent" AS "a"
	    INNER JOIN "public"."admin_user" AS "au" ON "au"."agent_id" = "a"."id"
	    WHERE "au"."is_added" = false AND "au"."role" = "_agent_permission_id" AND "a"."level_code" LIKE "_level_code" || '%'
    UNION ALL
	  SELECT "a"."id"
	    FROM "public"."agent" AS "a", "find_agents" AS "fa"
	    WHERE "a"."top_agent_id" = "fa"."id"
  ), "root_permission" AS (
    SELECT "permission"
      FROM "public"."agent_permission"
      WHERE "id" = "_agent_permission_id"
  )

  UPDATE "public"."agent_permission" AS "ag"
    SET "permission" = jsonb_set("permission", '{list}', "r"."list", false)
    FROM (
      SELECT "id", jsonb_agg("self") AS "list"
        FROM "public"."agent_permission"
        CROSS JOIN LATERAL jsonb_array_elements("permission"->'list') AS self("p")
		    INNER JOIN (SELECT jsonb_array_elements("permission"->'list') FROM "root_permission") AS root("p") USING ("p")
        WHERE "agent_id" IN (SELECT "id" FROM "find_agents")
        GROUP BY "id"
    ) AS "r"
    WHERE "ag"."id" = "r"."id";
END
$$;

DROP FUNCTION "public"."udf_update_agent";
CREATE FUNCTION "public"."udf_update_agent"("_agent_id" int4, "_agent_name" varchar, "_agent_info" varchar, "_agent_commission" int4, "_admin_user_username" varchar, "_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_is_enabled" int4, "_is_admin_user_role_changed" boolean)
  RETURNS json AS $$
DECLARE
  "agent_level_code" varchar;
  "ret_agent_update_time" timestamp;
  "ret_admin_user" json;
BEGIN
  UPDATE "public"."agent"
    SET "name" = "_agent_name",
      "info" = "_agent_info",
      "commission" = "_agent_commission",
      "update_time" = now()
    WHERE "id" = "_agent_id"
    RETURNING "update_time", "level_code" INTO "ret_agent_update_time", "agent_level_code";

  SELECT "public"."udf_update_admin_user"("_agent_id", "_admin_user_username", "_admin_user_role",
    "_admin_user_info", "_admin_user_is_enabled") INTO "ret_admin_user";

  IF "_is_admin_user_role_changed" THEN
    CALL "public"."usp_update_agents_permission" ("_admin_user_role", "agent_level_code");
  END IF;

  RETURN json_build_object(
    'agent', json_build_object(
      'update_time', extract(epoch from "ret_agent_update_time") * 1000000
    ),
    'admin_user', "ret_admin_user"
  );
END;
$$ LANGUAGE plpgsql;
