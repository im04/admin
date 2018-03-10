package main

import (
	routerSet "admin/router"
	"github.com/gin-gonic/gin"
	"admin/config"
	"fmt"
)

func main() {
	router := gin.Default()
	routerSet.InitRouter(router)
	router.Run(":" + fmt.Sprintf("%d", config.ServerConfig.Port))
}
