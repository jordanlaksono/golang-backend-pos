package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	Supplier_id       string `json:"supplier_id" gorm:"primary_key"`
	Supplier_kode     string `json:"supplier_kode" gorm:"type:varchar; unique; not null"`
	Supplier_nama     string `json:"supplier_nama" gorm:"type:varchar; not null"`
	Supplier_alamat   string `json:"supplier_alamat" gorm:"type:varchar; not null"`
	Supplier_kota     string `json:"supplier_kota" gorm:"type:varchar; not null"`
	Supplier_kode_pos string `json:"supplier_kode_pos" gorm:"type:varchar; not null"`
	Supplier_email    string `json:"supplier_email" gorm:"type:varchar; not null"`
	Supplier_no_telp  string `json:"supplier_no_telp" gorm:"type:varchar; not null"`
	Top               Top    `gorm:"references:Supplier_top_id;foreignKey:Top_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Supplier_top_id   string `json:"supplier_top_id" gorm:"index"`
}

func (m *Supplier) BeforeCreate(db *gorm.DB) error {
	m.Supplier_id = uuid.NewString()
	return nil
}

func (m *Supplier) BeforeUpdate(db *gorm.DB) error {
	return nil
}
