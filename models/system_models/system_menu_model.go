package system_models

import "easy-admin-go-service/models"

type SystemMenuModel struct {
	models.CommonModel
	Name     string            `json:"label" gorm:"column:name;comment:权限菜单名称"`
	Path     string            `json:"path" gorm:"column:path;comment:权限菜单路径"`
	ParentId uint              `json:"pId" gorm:"column:parentId;comment:根菜单id"`
	Icon     string            `json:"icon" gorm:"column:icon;comment:图标名称"`
	OrderNum uint              `json:"orderNum" gorm:"column:order_num;comment:排序"`
	Type     uint              `json:"type" gorm:"column:type;comment:类型 0：目录 1：菜单 2：按钮"`
	Role     []SystemRoleModel `gorm:"many2many:base_system_role_menu;"` // 多对多 用户角色表
}

func (SystemMenuModel) TableName() string {
	return "base_system_menu"
}
