package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Barang struct {
	Barang_id          string   `json:"barang_id" gorm:"primary_key"`
	Barang_kode        string   `json:"barang_kode" gorm:"type:varchar; not null"`
	Barang_nama        string   `json:"barang_nama" gorm:"type:varchar; not null"`
	Barang_image       string   `json:"barang_image" gorm:"type:varchar; not null"`
	Kelompok           Kelompok `gorm:"references:Barang_kelompok_id;foreignKey:Kelompok_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Barang_kelompok_id string   `json:"barang_kelompok_id" gorm:"index"`
	Supplier           Supplier `gorm:"references:Barang_supplier_id;foreignKey:Supplier_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Barang_supplier_id string   `json:"barang_supplier_id" gorm:"index"`
	Barang_keterangan  string   `json:"barang_keterangan" gorm:"type:varchar; not null"`
}

func (m *Barang) BeforeCreate(db *gorm.DB) error {
	m.Barang_id = uuid.NewString()
	return nil
}

func (m *Barang) BeforeUpdate(db *gorm.DB) error {
	return nil
}
