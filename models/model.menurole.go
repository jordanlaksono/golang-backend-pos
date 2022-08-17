package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRole struct {
	Menu_role_id string `json:"menu_role_id" gorm:"primary_key"`
	Menu         Menu
	MenuID       string `json:"menu_id" gorm:"index"`
	Role         Role
	Role_id      string `json:"role_id" gorm:"index"`
}

func (m *MenuRole) BeforeCreate(db *gorm.DB) error {
	m.Menu_role_id = uuid.NewString()
	return nil
}

func (m *MenuRole) BeforeUpdate(db *gorm.DB) error {
	m.Menu_role_id = uuid.NewString()
	return nil
}
