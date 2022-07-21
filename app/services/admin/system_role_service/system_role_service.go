package system_role_service

import (
	"easy-admin-go-service/global"
	"easy-admin-go-service/models/system_models"
)

type RoleService struct {
}

type IRoleService interface {
	GetUserRoleIds(userId int) (error, []int)
}

func (roleService *RoleService) GetUserRoleIds(userId int) (error, []int) {
	var roleIds []int
	if err := global.BaseDataBase.Model(&system_models.SystemRoleMenuModel{}).Select("role_id").Find(&roleIds).Error; err != nil {
		return err, nil
	}
	return nil, roleIds
}
