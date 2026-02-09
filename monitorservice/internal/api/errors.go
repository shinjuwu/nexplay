package api

import "errors"

const (
	EMPTY_JSON_DATA = "{}"
)

// 內部錯誤碼
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

// api response error code
const (
	ERROR_CODE_SUCCESS         = 0
	ERROR_CODE_FAIL            = 1 // 失敗
	ERROR_CODE_ERROR_UNAUTH    = 2 // http code 401
	ERROR_CODE_ERROR_EXCEPTION = 3 // http code 404
	ERROR_CODE_ERROR_LOCAL     = 4 // http code 500
	ERROR_CODE_ERROR_JWT       = 5

	RESPONSE_CODE_BINDPARAMS_SUCCESS = 100
	RESPONSE_CODE_BINDPARAMS_FAILED  = 101
)
