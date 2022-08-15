package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Role_id   string `json:"role_id" gorm:"primary_key"`
	Role_name string `json:"role_name" gorm:"type:varchar;  not null"`
}

func (m *Role) BeforeCreate(db *gorm.DB) error {
	m.Role_id = uuid.NewString()
	return nil
}

func (m *Role) BeforeUpdate(db *gorm.DB) error {
	m.Role_id = uuid.NewString()
	return nil
}
