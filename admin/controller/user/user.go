package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "admin/model/error"
	"admin/model/db"
	"admin/utils"
	"strconv"
	"admin/model/adminUser"
	"admin/config"
	"fmt"
	"encoding/json"
)


func SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": ErrorCode.SUCCESS,
		"msg": "SignUpOk",
	})
}

type SignInRequest struct {
	Account string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(ctx *gin.Context) {
	var requestParams SignInRequest
	if err := ctx.BindJSON(&requestParams); err != nil {
		SendError("登陆参数错误", ctx, err)
		return
	}
	var user adminUser.AdminUser
	if _, err := db.DB.Where("account = ?", requestParams.Account).Get(&user); err != nil {
		SendError("账号不存在", ctx, err)
		return
	}
	if err := user.CheckPassword(requestParams.Password); err != nil {
		SendError(err.Error(), ctx)
		return
	}
	content := strconv.Itoa(user.Id)
	key := config.ServerConfig.TokenSecret
	roundKey := ctx.ClientIP()
	tokenString, expTime := utils.CreateToken(key, content, roundKey, config.ServerConfig.TokenMaxAge)
	user.Token = tokenString
	redisCtx := db.RedisPool.Get()
	defer redisCtx.Close()
	fmt.Println("userId：", content)
	userBytes, err := json.Marshal(user)
	if err != nil {
		panic(err.Error())
	}
	if _, err := redisCtx.Do("SET", content, string(userBytes), "EX", expTime); err != nil {
		SendError(err.Error(), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": ErrorCode.SUCCESS,
		"msg": "登陆成功",
		"data": gin.H{
			"token": tokenString,
		},
	})
}

func SignOut(ctx *gin.Context) {
	user, flag := ctx.MustGet("userId").(adminUser.AdminUser)
	if !flag {
		SendError("程序错误", ctx)
		return
	}
	redisCtx := db.RedisPool.Get()
	defer redisCtx.Close()
	status, err := redisCtx.Do("DEL", strconv.Itoa(user.Id))
	if err != nil {
		SendError("程序错误", ctx, err)
		return
	}
	if status.(int64) == 0 {
		SendLoginError("用户未登陆", ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": ErrorCode.SUCCESS,
		"msg": "登出成功",
		"data": nil,
	})
	/*var tokenState string
	if ok {
		fmt.Println(claims)
		oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
		ct := time.Now().UTC().Unix()
		if  ct > oldT {
			ok = false
			tokenState = "Token 已过期"
		} else {
			tokenState = "Token 正常"
		}
	}else {
		tokenState = "token 无效"
	}
	fmt.Println(tokenState)*/

}

func GetMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": ErrorCode.SUCCESS,
		"msg": "获取菜单成功",
	})
}