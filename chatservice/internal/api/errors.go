package api

import "errors"

const (
	ErrCodeCommonSuccessed = 0   // 通用成功碼
	ErrCodeCommonFailed    = 813 // 通用錯誤碼
	ErrCodeParseFailed     = 814
	ErrCodeConnectClose    = 815
	ErrCodeHandleError     = 816
	ErrCodeDidNotInsertNew = 817
	ErrCodeDatabaseFailed  = 818
	ErrCodeIMHandleError   = 819

	ErrMessageCommonSuccessed = "successed"
	ErrMessageCommonFailed    = "failed"
	ErrMessageSendFail        = "send message failed"
	ErrMessageUnknowRPC       = "rpc id not define"
	ErrMessageUnknowIM        = "im name not find"
)

var (
	ErrSessionCreateFailed = errors.New("create session failed")
)
