package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lokasi struct {
	Lokasi_id   string `json:"lokasi_id" gorm:"primary_key"`
	Lokasi_kode string `json:"lokasi_kode" gorm:"type:varchar; not null"`
	Lokasi_nama string `json:"lokasi_nama" gorm:"type:varchar; not null"`
}

func (m *Lokasi) BeforeCreate(db *gorm.DB) error {
	m.Lokasi_id = uuid.NewString()

	return nil
}

func (m *Lokasi) BeforeUpdate(db *gorm.DB) error {
	//m.ID = uuid.NewString()
	//m.Satuan_updated_at = time.Now()
	return nil
}
