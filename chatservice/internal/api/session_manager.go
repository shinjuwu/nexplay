package api

import (
	"context"

	"go.uber.org/atomic"
)

type SessionManager interface {
	Stop()
	Count() int
	Get(sessionID string) ITequilaSession
	GetAll() []ITequilaSession
	Add(session ITequilaSession)
	Remove(sessionID string)
	Disconnect(ctx context.Context, sessionID string) error
	SingleSession(ctx context.Context, userID string)
}

type LocalSessionManager struct {
	sessions     *MapOf[string, ITequilaSession]
	sessionCount *atomic.Int32
}

func NewLocalSessionManager() SessionManager {
	return &LocalSessionManager{
		sessions:     &MapOf[string, ITequilaSession]{}, // session id, session object
		sessionCount: atomic.NewInt32(0),
	}
}

func (r *LocalSessionManager) Stop() {}

func (r *LocalSessionManager) Count() int {
	return int(r.sessionCount.Load())
}

func (r *LocalSessionManager) Get(sessionID string) ITequilaSession {
	session, ok := r.sessions.Load(sessionID)
	if !ok {
		return nil
	}
	return session
}

func (r *LocalSessionManager) GetAll() []ITequilaSession {
	ret := make([]ITequilaSession, 0)
	r.sessions.Range(func(key string, value ITequilaSession) bool {
		ret = append(ret, value)
		return true
	})

	return ret
}

func (r *LocalSessionManager) Add(session ITequilaSession) {
	r.sessions.Store(session.Id(), session)
	r.sessionCount.Inc()

}

func (r *LocalSessionManager) Remove(sessionID string) {
	r.sessions.Delete(sessionID)
	r.sessionCount.Dec()
}

func (r *LocalSessionManager) Disconnect(ctx context.Context, sessionID string) error {
	session, ok := r.sessions.Load(sessionID)
	if ok {
		session.Close()
		r.Remove(sessionID)
	}
	return nil
}

func (r *LocalSessionManager) SingleSession(ctx context.Context, userID string) {
	r.sessions.Range(func(sid string, v ITequilaSession) bool {
		if userID == v.UserId() {
			r.Disconnect(ctx, sid)
		}
		return true
	})
}
