package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Top struct {
	Top_id     string `json:"top_id" gorm:"primary_key"`
	Top_kode   string `json:"top_kode" gorm:"type:varchar; not null"`
	Top_nama   string `json:"top_nama" gorm:"type:varchar; not null"`
	Top_jumlah string `json:"top_jumlah" gorm:"type:bigint; not null"`
}

func (m *Top) BeforeCreate(db *gorm.DB) error {
	m.Top_id = uuid.NewString()
	return nil
}

func (m *Top) BeforeUpdate(db *gorm.DB) error {
	return nil
}
