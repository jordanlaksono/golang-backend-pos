package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelompok struct {
	Kelompok_id   string `json:"kelompok_id" gorm:"primary_key"`
	Kelompok_kode string `json:"kelompok_kode" gorm:"type:varchar; not null"`
	Kelompok_nama string `json:"kelompok_nama" gorm:"type:varchar; not null"`
}

func (m *Kelompok) BeforeCreate(db *gorm.DB) error {
	m.Kelompok_id = uuid.NewString()
	return nil
}

func (m *Kelompok) BeforeUpdate(db *gorm.DB) error {

	return nil
}
