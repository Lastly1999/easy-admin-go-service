package routers

import (
	_ "easy-admin-go-service/docs"
	"easy-admin-go-service/middlewares"
	"easy-admin-go-service/routers/auth_routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() *gin.Engine {
	gin.ForceConsoleColor()
	app := gin.Default()
	app.Use(middlewares.Cors())
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiV1 := app.Group("/v1")
	{
		// 权限路由组
		auth_routes.InitAuthRoutes(apiV1)
	}

	return app
}
