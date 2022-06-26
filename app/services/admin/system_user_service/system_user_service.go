package system_user_service

import (
	"easy-admin-go-service/app/services/admin/response"
	"easy-admin-go-service/models/system_models"
	"easy-admin-go-service/pkg/gorm"
)

type User struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

func (user *User) GetUser() (error, *response.GetUserInfoResponse) {
	getUserInfoResponse := &response.GetUserInfoResponse{}
	// 查询用户是否存在
	err := gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Where("user_name = ? AND pass_word = ?", user.UserName, user.PassWord).First(getUserInfoResponse).Error
	if err != nil {
		return err, nil
	}
	return nil, getUserInfoResponse
}
