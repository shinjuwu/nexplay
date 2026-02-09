package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"monitorservice/pkg/config"
	"monitorservice/pkg/melody"
	"monitorservice/pkg/utils"
)

/*
內部 func
*/

func TequilaHandleConnect(db *sql.DB, config config.Config, sessionManager SessionManager, me *melody.Session, e *Envelope) {
	userId := utils.ToString(me.Keys[TEQUILA_CTX_RUNTIME_USER_ID], "")
	platform := utils.ToString(me.Keys[TEQUILA_CTX_RUNTIME_PLATFORM], "")
	username := utils.ToString(me.Keys[TEQUILA_CTX_RUNTIME_USER_NAME], "")

	if config.GetSocket().SingleSocket {
		// Kick any other sockets for this user.
		sessionManager.SingleSession(context.Background(), userId)
	}

	mySession := NewTequilaSession(platform, userId, username, me)

	me.Keys[TEQUILA_CTX_RUNTIME_USER_SESSION_ID] = mySession.Id()

	sessionManager.Add(mySession)
}

func TequilaHandleDisconnect(db *sql.DB, config config.Config, sessionManager SessionManager, me *melody.Session, e *Envelope) {
	sessionID, ok := me.Keys["session_id"].(string)
	if ok {
		sessionManager.Disconnect(context.Background(), sessionID)
	}
}

func TequilaHandleError(db *sql.DB, config config.Config, sessionManager SessionManager, me *melody.Session, e *Envelope) {
	var err Error
	err.Subject = ProtocolSubjectError
	err.Code = ErrCodeHandleError
	err.Message = e.Payload

	js, _ := json.Marshal(e)
	me.Write(js)
}

func TequilaHandleClose(db *sql.DB, config config.Config, sessionManager SessionManager, me *melody.Session, e *Envelope) {
	sessionID, ok := me.Keys["session_id"].(string)
	if ok {
		var err Error
		err.Subject = ProtocolSubjectError
		err.Code = ErrCodeConnectClose
		err.Message = e.Payload

		js, _ := json.Marshal(e)
		me.Write(js)
		sessionManager.Disconnect(context.Background(), sessionID)
	}
}
