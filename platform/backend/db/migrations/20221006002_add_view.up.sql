CREATE VIEW "public"."view_login_admin_user" AS
  SELECT "au"."agent_id",
    "au"."username",
    "au"."password",
    "au"."nickname",
    "au"."google_auth",
    "au"."google_key",
    "au"."allow_ip",
    "au"."account_type",
    "au"."is_readonly",
    "au"."is_enabled",
    "ap"."permission",
    "au"."is_added",
    "au"."role"
  FROM "public"."admin_user" AS "au"
  INNER JOIN "public"."agent_permission" AS "ap" ON "ap"."id" = "au"."role";
