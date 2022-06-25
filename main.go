package main

import (
	"github.com/anthoturc/go/controllers"
	"github.com/anthoturc/go/database"
	"github.com/anthoturc/go/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:@tcp(localhost:3306)/ottotestjunior?parseTime=true")
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/login", controllers.Signin)
		api.Use()
		securedAth := api.Group("/check").Use(middlewares.Auth())
		{
			securedAth.GET("/profile", controllers.GetProfile)
			securedAth.GET("/wallet", controllers.GetWallet)
		}
	}
	return router
}
