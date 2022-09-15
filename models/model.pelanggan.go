package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pelanggan struct {
	Pelanggan_id        string `json:"pelanggan_id" gorm:"primary_key"`
	Pelanggan_kode      string `json:"pelanggan_kode" gorm:"type:varchar; unique; not null"`
	Pelanggan_nama      string `json:"pelanggan_nama" gorm:"type:varchar; not null"`
	Pelanggan_alamat    string `json:"pelanggan_alamat" gorm:"type:varchar; not null"`
	Pelanggan_no_telp_1 string `json:"pelanggan_no_telp_1" gorm:"type:varchar; not null"`
	Pelanggan_no_telp_2 string `json:"pelanggan_no_telp_2" gorm:"type:varchar; not null"`
	Pelanggan_email     string `json:"pelanggan_email" gorm:"type:varchar; not null"`
	Pelanggan_kota      string `json:"pelanggan_kota" gorm:"type:varchar; not null"`
	Pelanggan_kode_pos  string `json:"pelanggan_kode_pos" gorm:"type:varchar; not null"`
	Pelanggan_catatan   string `json:"pelanggan_catatan" gorm:"type:varchar; not null"`
	Top                 Top    `gorm:"references:Pelanggan_top_id;foreignKey:Top_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pelanggan_top_id    string `json:"pelanggan_top_id" gorm:"index"`
	Pelanggan_status    string `json:"pelanggan_status" gorm:"type:varchar; not null"`
	Penjualan_saldo     string `json:"penjualan_saldo" gorm:"type:varchar; not null;"`
}

func (m *Pelanggan) BeforeCreate(db *gorm.DB) error {
	m.Pelanggan_id = uuid.NewString()
	return nil
}
