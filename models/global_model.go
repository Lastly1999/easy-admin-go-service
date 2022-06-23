package models

import (
	"easy-admin-go-service/helps"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt helps.LocalTime `json:"createdAt"`
	UpdatedAt helps.LocalTime `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
