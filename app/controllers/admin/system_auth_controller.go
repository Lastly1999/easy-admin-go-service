package admin

import (
	"easy-admin-go-service/app/services/admin/system_auth_service"
	"easy-admin-go-service/global"
	"github.com/gin-gonic/gin"
	"net/http"
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
	userName := rep.Ctx.GetHeader("username")
	passWord := rep.Ctx.GetHeader("password")
	smsUuid := rep.Ctx.GetHeader("smsuuid")
	smsCode := rep.Ctx.GetHeader("smscode")
	service := system_auth_service.Auth{
		UserName: userName,
		PassWord: passWord,
		SmsUuid:  smsUuid,
		SmsCode:  smsCode,
	}
	err, checkLoginResponse := service.CheckLogin()
	if err != nil {
		rep.Error(http.StatusInternalServerError, err.Error())
		return
	}
	rep.Success(&checkLoginResponse)
}

// GetGraphicCode 获取图形验证码
// @Summary 获取图形验证码
// @Description 获取登录所需的图形验证码
// @Tags 授权
// @Accept application/json
// @Produce application/json
// @Router /auth/graphic [get]
// @Success 200 {object} response.GenerateGraphicCodeResponse
func (authApi *AuthApi) GetGraphicCode(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	auth := system_auth_service.Auth{}
	err, resp := auth.GenerateGraphicCode()
	if err != nil {
		rep.Error(http.StatusInternalServerError, err.Error())
		return
	}
	rep.Success(&resp)
}
