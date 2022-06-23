package system_models

import "easy-admin-go-service/models"

type SystemRoleUserModel struct {
	models.CommonModel
	UserId uint `json:"userId"`
	RoleId uint `json:"roleId"`
}

func (SystemRoleUserModel) TableName() string {
	return "base_system_role_user"
}
