package router

import (
	"github.com/gin-gonic/gin"
	adminUser "admin/admin/controller/user"
	clienUser "admin/client/controller/user"
	"admin/middleware"
)

func InitRouter (app *gin.Engine) {
	/*router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})*/
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middleware.SetHeader())
	client := app.Group("/api")
	{
		user := client.Group("/user")
		{
			user.POST("/signUp", clienUser.SignUp)
			user.POST("/signIn", clienUser.SignIn)
			user.POST("/signOut",middleware.CheckClientToken(), clienUser.SignOut)
		}
		any := client.Group("/any")
		{
			any.Use(middleware.CheckClientToken())
		}
	}
	admin := app.Group("/admin")
	{
		user := admin.Group("/user")
		{
			user.POST("/signUp", adminUser.SignUp)
			user.POST("/signIn", adminUser.SignIn)
			user.POST("/signOut", middleware.CheckAdminToken(), adminUser.SignOut)
		}
		common := admin.Group("/common")
		{
			common.Use(middleware.CheckAdminToken())
			common.GET("/getMenu", adminUser.GetMenu)
		}
	}

}
