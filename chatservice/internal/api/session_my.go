package api

import (
	"chatservice/pkg/melody"

	"github.com/gofrs/uuid"
)

type ITequilaSession interface {
	Id() string
	Platform() string
	UserId() string
	Username() string
	Close()
	Session() *melody.Session
	Send(msg string) error
	Sendb(msg []byte) error
}

type TequilaSession struct {
	id       string // session id
	platform string
	userId   string // user id
	username string // username
	session  *melody.Session
}

func NewTequilaSession(platform string, userId, username string, ms *melody.Session) ITequilaSession {
	uuidGen := uuid.NewGen()

	sessionId, _ := uuidGen.NewV4()
	return &TequilaSession{
		id:       sessionId.String(),
		platform: platform,
		userId:   userId,
		username: username,
		session:  ms,
	}
}

func (p *TequilaSession) Id() string {
	return p.id
}

func (p *TequilaSession) Platform() string {
	return p.platform
}

func (p *TequilaSession) UserId() string {
	return p.userId
}

func (p *TequilaSession) Username() string {
	return p.username
}

func (p *TequilaSession) Close() {}

func (p *TequilaSession) Session() *melody.Session {
	return p.session
}

func (p *TequilaSession) Send(msg string) error {
	return p.session.Write([]byte(msg))
}

func (p *TequilaSession) Sendb(msg []byte) error {
	return p.session.Write(msg)
}
