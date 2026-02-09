package middleware

import "time"

/*
CREATE TABLE "public"."admin_user_action_log"(
	"log_time" varchar(18) NOT NULL DEFAULT '',
	"username" varchar(100) NOT NULL DEFAULT '',
	"action_code" integer NOT NULL DEFAULT 0,
	"action_log" jsonb NOT NULL DEFAULT '{}',
	"ip" varchar(40) NOT NULL DEFAULT '',
	CONSTRAINT "admin_user_action_log_pkey" PRIMARY KEY("log_time","username")
);
COMMENT ON TABLE "public"."admin_user_action_log" IS '後台帳號操作記錄';

COMMENT ON COLUMN "public"."admin_user_action_log"."log_time" IS '紀錄時間';
COMMENT ON COLUMN "public"."admin_user_action_log"."username" IS '管理者帳號';
COMMENT ON COLUMN "public"."admin_user_action_log"."action_code" IS '操作代號';
COMMENT ON COLUMN "public"."admin_user_action_log"."action_log" IS '操作紀錄';
COMMENT ON COLUMN "public"."admin_user_action_log"."ip" IS '登入IP';
*/

const (
	// 一般操作 (get)
	OPERATE_GET = 0
	// 更新操作 (post, put, update, delete)
	OPERATE_POST   = 1
	OPERATE_PUT    = 2
	OPERATE_UPDATE = 3
	OPERATE_DELETE = 4
)

type AdminUserActionLog struct {
	LogTime   string `json:"log_time"`
	Username  string `json:"username"`
	ErrorCode int    `json:"error_code"`
	ActionLog string `json:"action_log"`
	Ip        string `json:"ip"`
}

type AdminUserBackendActionLog struct {
	AgentId     int       `json:"agent_id"`
	Username    string    `json:"username"`
	ErrorCode   int       `json:"error_code"`
	Ip          string    `json:"ip"`
	FeatureCode int       `json:"feature_code"`
	ActionType  int       `json:"action_type"`
	HttpLog     string    `json:"http_log"`
	ActionLog   string    `json:"action_log"`
	CreateTime  time.Time `json:"create_time"`
}

type AdminUserLoginLog struct {
	AgentId			int				`json:"agent_id"`
	Username  	string 		`json:"username"`
	Ip        	string 		`json:"ip"`
	ErrorCode 	int    		`json:"error_code"`
	LogTime   	time.Time 		`json:"log_time"`
}
