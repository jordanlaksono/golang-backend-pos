package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Karyawan struct {
	ID                string `json:"id" gorm:"primary_key"`
	Karyawan_nama     string `json:"karyawan_nama" gorm:"type:varchar; not null"`
	Karyawan_alamat   string `json:"karyawan_alamat" gorm:"type:text; not null"`
	Karyawan_kota     string `json:"karyawan_kota" gorm:"type:varchar; not null"`
	Karyawan_negara   string `json:"karyawan_negara" gorm:"type:varchar; not null"`
	Karyawan_kode_pos string `json:"karyawan_kode_pos" gorm:"type:bigint; not null"`
	Karyawan_no_telp  string `json:"karyawan_no_telp" gorm:"type:varchar; not null"`
	Karyawan_email    string `json:"karyawan_email" gorm:"type:varchar; not null"`
}

func (Karyawan) TableName() string {
	return "karyawans"
}

//type MenuRole struct {
//	Menu_role_id string `json:"menu_role_id" gorm:"primary_key"`
//	MenuID       string `json:"menu_id" gorm:"index"`
//	Menu_nama    string `json:"menu_nama" gorm:"index"`
//	Role_id      string `json:"role_id" gorm:"index"`
//	User_id      string `json:"user_id" gorm:"index"`
//}

func (m *Karyawan) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	return nil
}

func (m *Karyawan) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	// m.UpdatedAt = time.Now()
	return nil
}
