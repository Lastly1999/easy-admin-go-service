package system_models

import "easy-admin-go-service/models"

type SystemDepartmentModel struct {
	models.CommonModel
	DepName  string            `json:"depName" gorm:"column:dep_name;comment:部门名称"`
	DepPid   uint              `json:"depPid" gorm:"column:dep_pid;comment:根部门id"`
	OrderNum uint              `json:"orderNum" gorm:"column:order_num;comment:排序"`
	Users    []SystemUserModel `gorm:"ForeignKey:DepId;AssociationForeignKey:ID" json:"users"`
}

func (SystemDepartmentModel) TableName() string {
	return "base_system_department"
}
