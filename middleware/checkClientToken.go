package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckClientToken() gin.HandlerFunc {
	return checkToken
}