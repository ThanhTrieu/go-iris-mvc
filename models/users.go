package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID         uint
	Username   string `gorm:"size:60"`
	Password   string `gorm:"size:60"`
	Email      string `gorm:"size:60"`
	Phone      string `gorm:"size:10"`
	AuthenKey  string `gorm:"size:100"`
	Fullname   string `gorm:"size:60"`
	Address    string `gorm:"size:200"`
	Brirthday  *time.Time
	Status     int64
	Gender     int64
	LastLogin  time.Time
	created_at time.Time
	updated_at time.Time
	deleted_at time.Time
}
