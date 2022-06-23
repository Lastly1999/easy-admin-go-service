package system_models

import (
	"easy-admin-go-service/helps"
	"gorm.io/gorm"
)

type SystemRoleModel struct {
	Id        uint              `json:"roleId" gorm:"column:id;primary_key;comment:权限id"`
	RoleName  string            `json:"roleName" gorm:"column:role_name;comment:角色名称"`
	Describe  string            `json:"describe" gorm:"column:describe;comment:角色别名"`
	Status    *bool             `json:"status" gorm:"column:status;default:1;comment:状态"`
	User      []SystemUserModel `json:"user" gorm:"many2many:base_system_role_user"`              // 多对多 用户表
	Users     []SystemUserModel `json:"users" gorm:"ForeignKey:RoleId;AssociationForeignKey:ID" ` // 一对多 用户表
	Menus     []SystemMenuModel `json:"baseMenu" gorm:"many2many:base_system_role_menu"`          // 多对多 菜单表
	CreatedAt helps.LocalTime   `json:"createdAt"`
	UpdatedAt helps.LocalTime   `json:"updatedAt"`
	DeletedAt gorm.DeletedAt    `json:"-" gorm:"index"`
}

func (SystemRoleModel) TableName() string {
	return "base_system_role"
}
