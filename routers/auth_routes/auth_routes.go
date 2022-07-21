package auth_routes

import (
	v1 "easy-admin-go-service/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

// InitAuthRoutes Auth模块
func InitAuthRoutes(router *gin.RouterGroup) {
	authRoutes := router.Group("/auth")
	authApi := v1.AuthApi{}
	{
		// 获取图形验证码
		authRoutes.GET("/graphic", authApi.GetGraphicCode)
		// 授权登录
		authRoutes.POST("/login", authApi.LoginAction)
		// 获取授权菜单
		authRoutes.GET("/menus", authApi.GetRoleMenus)
	}
}
