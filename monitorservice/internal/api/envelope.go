package api

import (
	"encoding/json"
	"log"
	"time"
)

/*
一般訊息通知

client --> server
server --> client
*/

type Envelope struct {
	Id      string `json:"id"`      // method id (rpc id)
	Subject string `json:"subject"` // sub method id (exec state)
	Code    int    `json:"code"`    // error code
	Payload string `json:"payload"` // send data
}

func NewEnvelope(id, payload, subject string, code int) *Envelope {
	return &Envelope{
		Id:      id,
		Subject: subject,
		Code:    code,
		Payload: payload,
	}
}

func (n *Envelope) ToJson() string {
	b, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(b)
}

func (n *Envelope) ToJsonb() []byte {
	b, err := json.Marshal(n)
	if err != nil {
		return []byte("")
	}
	return b
}

/*
A client --> server
server --> A client
       --> B client
       --> C client
*/
type Notification struct {
	// ID of the Notification.
	Id string `json:"id"`
	// Subject of the notification.
	Subject string `json:"subject"`
	// Content of the notification in JSON.
	Content string `json:"content"`
	// ID of the sender, if a user. Otherwise 'null'.
	SenderId string `json:"sender_id"`
	// The UNIX time when the notification was created.
	CreateTime int64 `json:"create_time"`
}

func NewNotification(id, content, SenderId string) *Notification {
	return &Notification{
		Id:         id,
		Subject:    ProtocolSNotification,
		Content:    content,
		SenderId:   SenderId,
		CreateTime: time.Now().UTC().Unix(),
	}
}

func (n *Notification) ToJson() string {
	b, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(b)
}

func (n *Notification) ToJsonb() []byte {
	b, err := json.Marshal(n)
	if err != nil {
		return []byte("")
	}
	return b
}

/*
錯誤訊息回傳格式
*/
type Error struct {
	Subject string `json:"subject"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(message string, code int) *Error {
	return &Error{
		Subject: ProtocolSubjectError,
		Code:    code,
		Message: message,
	}
}

func (e *Error) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(b)
}
func (e *Error) ToJsonb() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
		return b
	}
	return b
}
