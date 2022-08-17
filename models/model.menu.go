package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID        string `json:"menu_id" gorm:"primary_key"`
	Menu_nama string `json:"menu_nama" gorm:"type:varchar;  not null"`
	Menu_link string `json:"menu_link" gorm:"type:varchar;  not null"`
}

func (m *Menu) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	return nil
}

func (m *Menu) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	return nil
}
