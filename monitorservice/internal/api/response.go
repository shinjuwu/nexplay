package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"` //回傳編碼
	Message string      `json:"msg"`
	Data    interface{} `json:"data"` //回傳資料
}

// http code 200
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context, data interface{}) {
	Result(c, ERROR_CODE_SUCCESS, "操作成功", data)
}

func Fail(c *gin.Context, data interface{}) {
	Result(c, ERROR_CODE_FAIL, "操作失敗", data)
}

// http code 404
func StatusNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code:    ERROR_CODE_ERROR_EXCEPTION,
		Message: "Error",
		Data:    "Not Found",
	})
}

// http code 400
// func StatusBadRequest(c *gin.Context) {
// 	c.JSON(http.StatusBadRequest, Response{
// 		Code:    ERROR_CODE_ERROR_PERMISSION,
// 		Message: "Error",
// 		Data:    "Bad Request",
// 	})
// }

// http code 401
func StatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    ERROR_CODE_ERROR_UNAUTH,
		Message: "Error",
		Data:    "Authentication failed, please log in again or contact the administrator",
	})
}

// http code 500
func StatusInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    ERROR_CODE_ERROR_LOCAL,
		Message: "Error",
		Data:    "Something wrong, please contact the administrator",
	})
}
