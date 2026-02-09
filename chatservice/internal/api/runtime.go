package api

import (
	"chatservice/pkg/config"
	"chatservice/pkg/melody"
	"database/sql"
)

const (
	TEQUILA_CTX_RUNTIME_USER_ID         = "user_id"
	TEQUILA_CTX_RUNTIME_PLATFORM        = "platform"
	TEQUILA_CTX_RUNTIME_AGENT_ID        = "agent_id"
	TEQUILA_CTX_RUNTIME_USER_NAME       = "username"
	TEQUILA_CTX_RUNTIME_USER_SESSION_ID = "session_id"

	TEQUILA_RUNTIME_EVENT_CONNECT    = "tequila_connect"
	TEQUILA_RUNTIME_EVENT_DISCONNECT = "tequila_disconnect"
	TEQUILA_RUNTIME_EVENT_ERROR      = "tequila_error"
	TEQUILA_RUNTIME_EVENT_CLOSE      = "tequila_close"
	TEQUILA_RUNTIME_EVENT_RPC        = "tequila_rpc"
)

// type MelodyFunc func(*melody.Session)

type RuntimeInternalEventFunc func(*sql.DB, config.Config, SessionManager, *melody.Session, *Envelope)

type RuntimeTequilaRpcFunc func(*sql.DB, ITequila, map[string]interface{}, string) (string, error)
