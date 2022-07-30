package system_user_service

import (
	"easy-admin-go-service/app/services/admin/request"
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
	InsertUser(request *request.CreateUserRequest) error
	UpdateUser(updateUserRequest *request.UpdateUserRequest) error
	FindUserList() ([]*response.UserListInfoResponse, error)
}

// UpdateUser 更新用户信息
func (userService *UserService) UpdateUser(updateUserRequest *request.UpdateUserRequest) error {
	err := gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Save(updateUserRequest).Error
	if err != nil {
		return err
	}
	return nil
}

// InsertUser 创建系统用户
func (userService *UserService) InsertUser(request *request.CreateUserRequest) error {
	gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Create(request)
	return nil
}

// FindUserInfo 查询用户是否存在
func (userService *UserService) FindUserInfo() (*response.GetUserInfoResponse, error) {
	getUserInfoResponse := &response.GetUserInfoResponse{}
	// 查询用户是否存在
	err := gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Where("user_name = ? AND pass_word = ?", userService.UserName, userService.PassWord).First(getUserInfoResponse).Error
	if err != nil {
		return nil, err
	}
	return getUserInfoResponse, nil
}

// FindUserList 查询用户列表
func (userService *UserService) FindUserList() ([]*response.UserListInfoResponse, error) {
	var userListInfoResponse []*response.UserListInfoResponse
	if err := gorm.BaseDataBase.Model(&system_models.SystemUserModel{}).Find(&userListInfoResponse).Error; err != nil {
		return nil, err
	}
	return userListInfoResponse, nil
}
