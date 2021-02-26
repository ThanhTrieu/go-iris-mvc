package models

import (
	"time"
)

type Groups struct {
	ID           uint
	GroupName    string `gorm:"size:100"`
	Description  string `gorm:"size:255"`
	Status       int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}