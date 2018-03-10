package middleware

import "github.com/gin-gonic/gin"

/*cache-control →no-store, no-cache, must-revalidate, post-check=0, pre-check=0
connection →keep-alive
content-type →application/json; charset=utf-8*/
func SetHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("cache-control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		ctx.Header("connection", "keep-alive")
		ctx.Next()
	}
}