package admin

import (
	"easy-admin-go-service/app/services/admin/system_auth_service"
	"easy-admin-go-service/global"
	"easy-admin-go-service/pkg/gorm"
	"easy-admin-go-service/pkg/redis_store"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

// LoginAction 授权登录
// @Summary 用户登录
// @Description 授权登录
// @Tags 授权
// @Accept application/json
// @Produce application/json
// @Router /auth/login [post]
func (authApi *AuthApi) LoginAction(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	service := system_auth_service.Auth{
		UserName: "pass",
	}
	ret := service.CheckLogin()
	rep.Success(&ret)
}

// GetGraphicCode 获取图形验证码
// @Summary 获取图形验证码
// @Description 获取登录所需的图形验证码
// @Tags 授权
// @Accept application/json
// @Produce application/json
// @Router /auth/graphic [get]
func (authApi *AuthApi) GetGraphicCode(ctx *gin.Context) {
	gorm.BaseDataBase.Table("sys_user")
	redis_store.RedisDb.Time(ctx)
}
