package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRole struct {
	Menu_role_id string `json:"menu_role_id" gorm:"primary_key"`
	MenuID       string `json:"menu_id" gorm:"index"`
	Menu_nama    string `json:"menu_nama" gorm:"type:varchar;  not null"`
	Menu_link    string `json:"menu_link" gorm:"type:varchar;  not null"`
	Role_id      string `json:"role_id" gorm:"index"`
	User_id      string `json:"user_id" gorm:"index"`
	Menu_r       string `json:"menu_r" gorm:"type:varchar;  not null"`
}

func (m *MenuRole) BeforeCreate(db *gorm.DB) error {
	m.Menu_role_id = uuid.NewString()
	return nil
}

func (m *MenuRole) BeforeUpdate(db *gorm.DB) error {
	m.Menu_role_id = uuid.NewString()
	return nil
}
