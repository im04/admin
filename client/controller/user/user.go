package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "SignUpOk",
	})
}

func SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "SignInOk",
	})
}

func SignOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "SignOutOk",
	})
}