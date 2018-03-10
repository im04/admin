package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckAdminToken() gin.HandlerFunc {
	return checkToken
}