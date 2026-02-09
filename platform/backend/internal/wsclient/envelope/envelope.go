package envelope

import "encoding/json"

const (
	ProtocolSubjectError   string = "error"
	ProtocolSubjectMessage string = "message"
	ProtocolSNotification  string = "notification"
)

type Message struct {
	Id      string `json:"id"`      // method id (rpc id)
	Subject string `json:"subject"` // sub method id (exec state)
	Code    int    `json:"code"`    // error code
	Payload string `json:"payload"` // send data
}

func NewMessage(id, payload string, code int) *Message {
	return &Message{
		Id:      id,
		Subject: ProtocolSubjectMessage,
		Code:    code,
		Payload: payload,
	}
}

func (n *Message) ToJson() string {
	b, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(b)
}

func (n *Message) ToJsonb() []byte {
	b, err := json.Marshal(n)
	if err != nil {
		return []byte("")
	}
	return b
}
