package request

import "easy-admin-go-service/models"

type CreateUserRequest struct {
	UserName   string `json:"userName" gorm:"column:user_name;comment:用户名"`
	PassWord   string `json:"passWord" gorm:"column:pass_word;comment:账户密码"`
	UserAvatar string `json:"userAvatar" gorm:"column:user_avatar;comment:用户头像"`
	NikeName   string `json:"nikeName" gorm:"column:nike_name;comment:用户昵称"`
}

type UpdateUserRequest struct {
	models.CommonModel
	UserName   string `json:"userName" gorm:"column:user_name;comment:用户名"`
	UserAvatar string `json:"userAvatar" gorm:"column:user_avatar;comment:用户头像"`
	NikeName   string `json:"nikeName" gorm:"column:nike_name;comment:用户昵称"`
}
