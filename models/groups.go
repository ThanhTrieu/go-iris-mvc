package models

import (
	"time"

	"gorm.io/gorm"
)

type Groups struct {
	gorm.Model
	ID           uint
	GroupName    string `gorm:"size:100"`
	Description  string `gorm:"size:255"`
	Status       int64
	created_at time.Time
	updated_at time.Time
	deleted_at time.Time
}