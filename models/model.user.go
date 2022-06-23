package models

import (
	"github.com/google/uuid"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"gorm.io/gorm"
)

type User struct {
	User_id         string `json:"user_id" gorm:"primary_key"`
	User_first_name string `json:"user_first_name" gorm:"type:varchar;  not null"`
	User_last_name  string `json:"user_last_name" gorm:"type:varchar; not null"`
	User_email      string `json:"user_email" gorm:"type:varchar; unique; not null"`
	User_password   string `json:"user_password" gorm:"type:varchar; not null"`
}

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
