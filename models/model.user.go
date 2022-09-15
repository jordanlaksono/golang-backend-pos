package models

import (
	"github.com/google/uuid"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"gorm.io/gorm"
)

type User struct {
	User_id        string   `json:"user_id" gorm:"primary_key"`
	User_username  string   `json:"user_username" gorm:"type:varchar; unique; not null"`
	User_password  string   `json:"user_password" gorm:"type:varchar; not null"`
	Karyawan       Karyawan `gorm:"references:UserKaryawanID;foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserKaryawanID string   `json:"user_karyawan_id" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

//type MenuRole struct {
//	Menu_role_id string `json:"menu_role_id" gorm:"primary_key"`
//	MenuID       string `json:"menu_id" gorm:"index"`
//	Menu_nama    string `json:"menu_nama" gorm:"index"`
//	Role_id      string `json:"role_id" gorm:"index"`
//	User_id      string `json:"user_id" gorm:"index"`
//}

func (m *User) BeforeCreate(db *gorm.DB) error {
	m.User_id = uuid.NewString()
	m.User_password = pkg.HashPassword(m.User_password)
	return nil
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
	m.User_id = uuid.NewString()
	m.User_password = pkg.HashPassword(m.User_password)
	// m.UpdatedAt = time.Now()
	return nil
}
