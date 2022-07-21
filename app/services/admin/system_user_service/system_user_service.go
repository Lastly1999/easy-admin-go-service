package system_user_service

import (
	"easy-admin-go-service/app/services/admin/response"
	"easy-admin-go-service/models/system_models"
	"easy-admin-go-service/pkg/gorm"
)

type UserService struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type IUserService interface {
	FindUserInfo() (*response.GetUserInfoResponse, error)
}

func (userService *UserService) FindUserInfo() (*response.GetUserInfoResponse, error) {
	getUserInfoResponse := &response.GetUserInfoResponse{}
	// 查询用户是否存在
	err := gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Where("user_name = ? AND pass_word = ?", userService.UserName, userService.PassWord).First(getUserInfoResponse).Error
	if err != nil {
		return nil, err
	}
	return getUserInfoResponse, nil
}
