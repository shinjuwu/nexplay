package api

const (
	ProtocolSubjectError        string = "error"
	ProtocolSubjectMessage      string = "message"
	ProtocolSNotification       string = "notification"
	ProtocolSubjectAnnouncement string = "announcement"
)

type ProtocolCode int8

const (
	ProtocolError ProtocolCode = iota - 1
	ProtocolMessage
	ProtocolNotification
)
