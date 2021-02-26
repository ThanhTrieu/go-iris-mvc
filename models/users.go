package models

import (
	"time"
)

type Users struct {
	ID         uint
	Username   string `gorm:"size:60"`
	Password   string `gorm:"size:60"`
	Email      string `gorm:"size:60"`
	Phone      string `gorm:"size:10"`
	Role       int64
	AuthenKey  string `gorm:"size:100"`
	Fullname   string `gorm:"size:60"`
	Address    string `gorm:"size:200"`
	Brirthday  *time.Time
	Status     int64
	Gender     int64
	LastLogin  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
