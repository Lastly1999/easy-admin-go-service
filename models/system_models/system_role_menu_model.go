package system_models

import "easy-admin-go-service/models"

type SystemRoleMenuModel struct {
	models.CommonModel
	RoleId uint `json:"roleId"`
	MenuId uint `json:"menuId"`
}

func (SystemRoleMenuModel) TableName() string {
	return "base_system_role_menu"
}
