package api

import (
	"chatservice/pkg/config"
	"chatservice/pkg/melody"
	"chatservice/pkg/utils"
	"database/sql"
)

type Pipeline struct {
	db                *sql.DB
	config            config.Config
	sessionManager    SessionManager
	internalEventFunc map[string]RuntimeInternalEventFunc
	tequila           *Tequila
	logger            ILogger
}

func NewPipeline(logger ILogger, db *sql.DB, config config.Config, sm SessionManager, t *Tequila) *Pipeline {
	iefn := make(map[string]RuntimeInternalEventFunc, 0)

	iefn[TEQUILA_RUNTIME_EVENT_CONNECT] = TequilaHandleConnect
	iefn[TEQUILA_RUNTIME_EVENT_DISCONNECT] = TequilaHandleDisconnect
	iefn[TEQUILA_RUNTIME_EVENT_ERROR] = TequilaHandleError
	iefn[TEQUILA_RUNTIME_EVENT_CLOSE] = TequilaHandleClose

	return &Pipeline{
		db:                db,
		config:            config,
		sessionManager:    sm,
		internalEventFunc: iefn,
		tequila:           t,
		logger:            logger,
	}
}

/*
內部動作事件觸發定義
連線、斷線、發生錯誤、關閉連線
*/
func (p *Pipeline) ProcessInternalRequest(cmd string, s *melody.Session, e *Envelope) {

	sessionId := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_SESSION_ID])
	userId := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_ID])
	username := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_NAME])

	p.logger.Info("ProcessInternalRequest(), sessionId:%s, userId:%s, username:%s ,id:%s, payload: %s", sessionId, userId, username, cmd, e.Payload)
	switch cmd {
	case TEQUILA_RUNTIME_EVENT_CONNECT:
		fallthrough
	case TEQUILA_RUNTIME_EVENT_DISCONNECT:
		fallthrough
	case TEQUILA_RUNTIME_EVENT_ERROR:
		fallthrough
	case TEQUILA_RUNTIME_EVENT_CLOSE:
		fn, ok := p.internalEventFunc[cmd]
		if ok {
			fn(p.db, p.config, p.sessionManager, s, e)
		} else {
			p.logger.Printf("ProcessInternalRequest() unknow cmd: %s", cmd)
		}
	}
}

func (p *Pipeline) ProcessRequest(cmd string, s *melody.Session, e *Envelope) (bool, error) {

	sessionId := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_SESSION_ID])
	userId := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_ID])
	username := utils.ToString(s.Keys[TEQUILA_CTX_RUNTIME_USER_NAME])

	var result string
	var err error
	p.logger.Info("ProcessRequest(), sessionId:%s, userId:%s, username:%s ,id:%s, payload: %s", sessionId, userId, username, cmd, e.Payload)
	// TODO: 加入自定義事件 func
	fn, ok := p.tequila.rpc[cmd]
	if ok {
		result, err = fn(p.db, p.tequila, s.Keys, e.Payload)
	} else {
		p.logger.Printf("ProcessRequest() unknow cmd: %s", cmd)
		return false, nil
	}

	if err != nil {
		p.logger.Printf("ProcessRequest() has error: %v", err)
		return true, err
	} else {

		tmp := &Envelope{
			Id:      e.Id,
			Subject: e.Subject,
			Code:    ErrCodeCommonSuccessed,
			Payload: result,
		}

		s.Write(tmp.ToJsonb())
	}

	return true, nil
}
