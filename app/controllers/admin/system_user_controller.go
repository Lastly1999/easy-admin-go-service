package admin

import (
	"easy-admin-go-service/app/services/admin/response"
	"easy-admin-go-service/global"
	"easy-admin-go-service/models/system_models"
	"easy-admin-go-service/pkg/gorm"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (userApi *UserApi) GetUser(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	getUserInfoResponse := &response.GetUserInfoResponse{}
	gorm.BaseDataBase.Model(&system_models.SystemUserModel{
		UserName: "admin",
		PassWord: "1234",
	}).First(&getUserInfoResponse)
	rep.Success(getUserInfoResponse)
}
