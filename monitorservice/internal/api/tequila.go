package api

import (
	"database/sql"
	"monitorservice/pkg/logger"
	"strings"
)

/*
定義內部功能模組供外部使用
*/
type ITequila interface {
	RegisterRPC(rpcName string, fn RuntimeTequilaRpcFunc) error
	NotificationSend(userID string, e *Envelope)
	NotificationListSend(notifications map[string]*Envelope)
	BroadcastServer(e *Envelope)
}

type Tequila struct {
	db             *sql.DB
	sessionManager SessionManager
	rpc            map[string]RuntimeTequilaRpcFunc
	logger         ILogger
}

func NewTequila(logger *logger.RuntimeGoLogger, db *sql.DB, sessionManager SessionManager) *Tequila {
	return &Tequila{
		db:             db,
		sessionManager: sessionManager,
		rpc:            make(map[string]RuntimeTequilaRpcFunc),
		logger:         logger,
	}
}

func (t *Tequila) RegisterRPC(rpcName string, fn RuntimeTequilaRpcFunc) error {
	name := strings.ToLower(rpcName)
	t.rpc[name] = fn
	t.logger.Printf("RegisterRPC() rpc id: %s", name)
	return nil
}

func (t *Tequila) BroadcastServer(e *Envelope) {
	e.Id = strings.ToLower(e.Id)
	msgb := e.ToJsonb()
	for _, s := range t.sessionManager.GetAll() {
		s.Sendb(msgb)
	}
}

func (t *Tequila) NotificationSend(userID string, e *Envelope) {
	sObj := t.sessionManager.Get(userID)
	sObj.Sendb(e.ToJsonb())
}

func (t *Tequila) NotificationListSend(notifications map[string]*Envelope) {
	for userID, ns := range notifications {
		sObj := t.sessionManager.Get(userID)
		sObj.Sendb(ns.ToJsonb())
	}
}
