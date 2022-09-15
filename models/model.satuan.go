package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Satuan struct {
	Satuan_id         string    `json:"satuan_id" gorm:"primary_key"`
	Satuan_kode       string    `json:"satuan_kode" gorm:"type:varchar; not null"`
	Satuan_nama       string    `json:"satuan_nama" gorm:"type:varchar; not null"`
	Satuan_created_at time.Time `json:"satuan_created_at"`
	Satuan_updated_at time.Time `json:"satuan_updated_at"`
}

func (m *Satuan) BeforeCreate(db *gorm.DB) error {
	m.Satuan_id = uuid.NewString()
	m.Satuan_created_at = time.Now()
	return nil
}

func (m *Satuan) BeforeUpdate(db *gorm.DB) error {
	//m.ID = uuid.NewString()
	//m.Satuan_updated_at = time.Now()
	return nil
}
