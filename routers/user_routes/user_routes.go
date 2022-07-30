package user_routes

import (
	"easy-admin-go-service/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

// InitUserRoutes 用户模块路由组
func InitUserRoutes(router *gin.RouterGroup) {
	userApi := admin.UserApi{}
	userRouter := router.Group("/user")
	{
		// 创建用户
		userRouter.PUT("/user", userApi.Create)
		// 更新用户
		userRouter.PATCH("/user", userApi.Update)
		// 用户列表
		userRouter.POST("/user", userApi.List)
	}
}
