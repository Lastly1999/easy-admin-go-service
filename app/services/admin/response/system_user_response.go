package response

import (
	"easy-admin-go-service/models"
	"easy-admin-go-service/models/system_models"
)

type GetUserInfoResponse struct {
	models.CommonModel
	UserName   string                          `json:"userName" `
	UserAvatar string                          `json:"userAvatar" `
	NikeName   string                          `json:"nikeName"`
	RoleId     uint                            `json:"roleId"`
	Status     *bool                           `json:"status"`
	Role       []system_models.SystemRoleModel `json:"role" gorm:"many2many:base_system_role_user;"` // 多对多 用户角色表
	DepId      string
}

type UserListInfoResponse struct {
	models.CommonModel
	UserName   string                          `json:"userName" `
	UserAvatar string                          `json:"userAvatar" `
	NikeName   string                          `json:"nikeName"`
	RoleId     uint                            `json:"roleId"`
	Status     *bool                           `json:"status"`
	Role       []system_models.SystemRoleModel `json:"role" gorm:"many2many:base_system_role_user;"` // 多对多 用户角色表
	DepId      string
}
