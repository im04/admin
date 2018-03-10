package error

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type errorCode struct {
	SUCCESS      int
	ERROR        int
	NotFound     int
	LoginError   int
	LoginTimeout int
	InActive     int
}

// ErrorCode 错误码
var ErrorCode = errorCode{
	SUCCESS      : 0,
	ERROR        : 1,
	NotFound     : 404,
	LoginError   : 1000, //用户名或密码错误
	LoginTimeout : 1001, //登录超时
	InActive     : 1002, //未激活账号
}

func sendError(msg string, c *gin.Context, code int, arg ... interface{}) {
	var data gin.H
	if len(arg) > 0 {
		e, er := arg[0].(error)
		if er {
			data = gin.H{
				"errorMsg": e.Error(),
			}
		} else {
			data = gin.H{
				"errorMsg": e,
			}
		}
	} else {
		data = nil
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":   msg,
		"data":  data,
	})
}

func SendError(msg string, c *gin.Context, arg ... interface{}) {
	sendError(msg, c, ErrorCode.ERROR, arg ...)
}
func SendLoginError(msg string, c *gin.Context, arg ... interface{}) {
	sendError(msg, c, ErrorCode.LoginError, arg ...)
}
func SendLoginTimeout(msg string, c *gin.Context, arg ... interface{}) {
	sendError(msg, c, ErrorCode.LoginTimeout, arg ...)
}
