package models

import (
	"time"
)

type MemberFolder struct {
	ID       uint
	LeaderId   int64
	MemberName  string `gorm:"size:64"`
	MemberEmail string `gorm:"size:100"`
	MemberTelegram string `gorm:"size:100"`
	MemberFolder string `gorm:"size:1024"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}