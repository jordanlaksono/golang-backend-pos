package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Penjualan struct {
	Penjualan_id           string    `json:"penjualan_id" gorm:"primary_key"`
	Karyawan               Karyawan  `gorm:"references:Penjualan_karyawan_id;foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Penjualan_karyawan_id  string    `json:"penjualan_karyawan_id" gorm:"index"`
	Pelanggan              Pelanggan `gorm:"references:Penjualan_pelanggan_id;foreignKey:Pelanggan_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Penjualan_pelanggan_id string    `json:"penjualan_pelanggan_id" gorm:"index"`
	//Top                         Top       `gorm:"references:Penjualan_top_id;foreignKey:Top_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Penjualan_top_id            string    `json:"penjualan_top_id" gorm:"index"`
	Penjualan_no_faktur         string    `json:"penjualan_no_faktur" gorm:"type:varchar; not null"`
	Penjualan_no_nota           string    `json:"penjualan_no_nota" gorm:"type:varchar; not null"`
	Penjualan_metode_pembayaran string    `json:"penjualan_metode_pembayaran" gorm:"type:varchar; not null"`
	Penjualan_diskon_persen     string    `json:"penjualan_metode_pembayaran" gorm:"type:varchar; not null"`
	Penjualan_diskon_rupiah     string    `json:"penjualan_diskon_rupiah" gorm:"type:varchar; not null"`
	Penjualan_total_harga       string    `json:"penjualan_total_harga" gorm:"type:varchar; not null"`
	Penjualan_biaya             string    `json:"penjualan_biaya" gorm:"type:varchar; not null"`
	Penjualan_biaya_debit       string    `json:"penjualan_biaya_debit" gorm:"type:varchar; not null"`
	Penjualan_selisih           string    `json:"penjualan_selisih" gorm:"type:varchar; not null"`
	Penjualan_potongan          string    `json:"penjualan_potongan" gorm:"type:varchar; not null"`
	Penjualan_biaya_kirim       string    `json:"penjualan_biaya_kirim" gorm:"type:varchar; not null"`
	Penjualan_jatuh_tempo       time.Time `json:"penjualan_jatuh_tempo"`
}

func (m *Penjualan) BeforeCreate(db *gorm.DB) error {
	m.Penjualan_id = uuid.NewString()
	return nil
}
