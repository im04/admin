package middleware

import (
	"github.com/gin-gonic/gin"
	"admin/model/db"
	myError "admin/model/error"
	"admin/utils"
	"fmt"
	"time"
	"net/http"
	"log"
	"github.com/dgrijalva/jwt-go"
	"admin/config"
	"github.com/kataras/iris/core/errors"
	"encoding/json"
	//"admin/model/adminUser"
	"github.com/garyburd/redigo/redis"
	"admin/model/adminUser"
)
func checkToken(ctx *gin.Context) {
	t := time.Now()
	repErr := func(e error) {
		ctx.Abort()
		myError.SendLoginError(e.Error(), ctx)
	}
	defer checkTokenPanic(ctx)
	fmt.Println("checkToken")
	token := ctx.GetHeader("token")
	fmt.Println(token)
	if token == "" {
		repErr(errors.New("未登陆"))
		return
	}
	claims, ok := utils.ParseToken(token, config.ServerConfig.TokenSecret)
	if !ok {
		repErr(errors.New("未登陆"))
		return
	}
	content, flag := claims.(jwt.MapClaims)["content"].(string)
	fmt.Println(content)
	if !flag {
		repErr(errors.New("未登陆"))
		return
	}
	var tokenContent utils.TokenJson
 	err := json.Unmarshal([]byte(content), &tokenContent)
 	if err != nil {
 		fmt.Println("解析json错误")
 		repErr(err)
		return
	}
	ct := time.Now().UTC().Unix()
	if   ct > tokenContent.Exp  {
		repErr(errors.New("登录过期"))
		return
	}
	redisCtx := db.RedisPool.Get()
	defer redisCtx.Close()
	var userRedis adminUser.AdminUser
	userBytes, err := redis.Bytes(redisCtx.Do("GET", tokenContent.UserId))
	if err != nil {
		repErr(errors.New("用户未登陆"))
		return
	}
	bytesErr := json.Unmarshal(userBytes, &userRedis)
	if bytesErr != nil {
		repErr(err)
		return
	}
	if err != nil {
		repErr(errors.New("未登陆:123"))
		return
	}
	if userRedis.Token != token {
		repErr(errors.New("token失效"))
		return
	}
	ctx.Set("user", userRedis)
	ctx.Next()
	latency := time.Since(t)
	log.Print(latency)
}
func checkTokenPanic (ctx *gin.Context) {
	if err := recover(); err != nil {
		fmt.Println("触发panic:", err)
		e, er := err.(error)
		ctx.Abort()
		if er {
			ctx.JSON(http.StatusOK, gin.H{
				"code": myError.ErrorCode.ERROR,
				"msg": e.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": myError.ErrorCode.ERROR,
				"msg": err,
			})
		}
	}
}
